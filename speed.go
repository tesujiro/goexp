package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const DEBUG = 0

func dprintf(format string, a ...interface{}) {
	if DEBUG != 0 {
		fmt.Fprintf(os.Stderr, format, a...)
	}
}

func openfile(filename string) (*os.File, os.FileInfo, error) {
	cur, err := os.Getwd()
	if err != nil {
		return nil, nil, err
	}
	//filename := flag.Args()[0]
	filePath := filepath.Join(cur, filename)
	fileinfo, err := os.Stat(filePath)
	if err != nil {
		return nil, nil, err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	return file, fileinfo, nil
}

type option struct {
	speed    int
	silent   bool
	graph    bool
	filename string
}

func getOption() *option {
	// Todo: split Option parsing
	var speed *int = flag.Int("bandwidth", 0, "Bytes Per Sec.")
	var silent *bool = flag.Bool("silent", false, "Silent Mode")
	var graph *bool = flag.Bool("graph", false, "Graphic Mode")
	flag.Parse()

	var filename string
	switch len(flag.Args()) {
	case 0:
		filename = ""
	case 1:
		filename = flag.Args()[0]
	default:
		fmt.Fprintf(os.Stderr, "\n\nParameter Error\n") // Todo read more than one file at once
		os.Exit(9)
	}

	return &option{
		speed:    *speed,
		silent:   *silent,
		graph:    *graph,
		filename: filename,
	}
}

func main() {

	option := getOption()

	if option.filename == "" {
		limitedPipe(os.Stdin, os.Stdout, 0, option)
	} else {
		file, fileinfo, err := openfile(option.filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\n\nFile Open Error :%v\n", err)
			os.Exit(9)
		}
		defer file.Close()
		limitedPipe(file, os.Stdout, int(fileinfo.Size()), option)
	}
}

const BUFSIZE = 4096

type readbuf struct {
	length int
	buf    []byte
}

func read(in io.Reader, rb chan readbuf) {
	reader := bufio.NewReader(in)
	buf := make([]byte, BUFSIZE)
	for {
		if n, err := reader.Read(buf); n == 0 || err == io.EOF {
			rb <- readbuf{length: n, buf: buf}
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "\n\nFile Read Error :%v\n", err)
			os.Exit(9)
		} else {
			rb <- readbuf{length: n, buf: buf}
		}
	}
}

func limitedPipe(in io.Reader, out io.Writer, size int, option *option) {
	ctx := context.Background()
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)

	rbchan := make(chan readbuf, 1)
	reading_done := make(chan struct{}, 1)
	go func() {
		read(in, rbchan)
		reading_done <- struct{}{}
	}()

	wg := &sync.WaitGroup{}

	//sk := NewSpeedKeeper(time.Now(), speed, size)
	sk := NewSpeedKeeper(ctx, cancel, time.Now(), option.speed, size)
	wg.Add(1)
	go func() {
		sk.run()
		wg.Done()
	}()

	mon := newMonitor(ctx, cancel, sk)
	if option.silent == true {
		mon.setMode("silent")
	}
	if option.graph == true {
		mon.setMode("graph")
	}

	wg.Add(1)
	go func() {
		mon.run()
		wg.Done()
	}()

	tick := time.NewTicker(time.Millisecond * time.Duration(50)).C
	readBytes := 0
L:
	for {
		select {
		case rb := <-rbchan:
			readBytes += rb.length
			out.Write(rb.buf)
			sk.curchan <- readBytes
			<-sk.killTime()
		case <-tick:
			mon.progress <- struct{}{}
		case <-reading_done:
			mon.progress <- struct{}{}
			cancel()
			break L
		case <-ctx.Done():
			break L
		}
	}
	wg.Wait()
}

//
// The speedKeeper is a goroutine which keeps input/output speed.
//

type speedKeeper struct {
	ctx        context.Context
	cancel     func()
	start      time.Time
	bytePerSec int
	size       int
	current    int
	curchan    chan int
}

func NewSpeedKeeper(ctx context.Context, cancel func(), s time.Time, b int, size int) *speedKeeper {
	return &speedKeeper{
		ctx:        ctx,
		cancel:     cancel,
		start:      s,
		bytePerSec: b,
		size:       size,
		current:    0,
		curchan:    make(chan int),
	}
}

func (sk *speedKeeper) run() {
L:
	for {
		select {
		case curBytes := <-sk.curchan:
			sk.current = curBytes
		case <-sk.ctx.Done():
			break L
		}
	}
}

func (sk *speedKeeper) killTime() <-chan struct{} {
	outchan := make(chan struct{})
	go func() {
		if sk.bytePerSec > 0 {
			//target_duration := time.Duration(float64(curBytes/sk.bytePerSec)) * time.Second //NG
			//target_duration := time.Duration(float64(curBytes*1e9/sk.bytePerSec)) * time.Nanosecond
			//target_duration := time.Duration(float64(curBytes*1000/sk.bytePerSec)) * time.Millisecond
			target_duration := time.Duration(float64(sk.current*1000/sk.bytePerSec)) * time.Millisecond
			current_duration := time.Since(sk.start)
			wait := target_duration - current_duration
			dprintf(" target=%s current=%s\n", target_duration, current_duration)
			if wait > 0 {
				dprintf("Sleep %s\n", wait)
				time.Sleep(wait)
			}
			dprintf(" wait finished %d\n", wait)
		}
		outchan <- struct{}{}
	}()
	return outchan
}

func (sk *speedKeeper) currentSpeed() int {
	//d := int(time.Since(sk.start).Seconds())
	d := int(time.Since(sk.start).Nanoseconds())
	if d == 0 {
		return 0
	}
	return sk.current * 1e9 / d
}

//
// The monitor is a progress monitoring goroutine.
//

type monitor struct {
	ctx      context.Context
	cancel   func()
	tty      io.Writer
	sk       *speedKeeper
	mode     string // Monitor Mode : Standard, Silent, Graphical,
	progress chan struct{}
}

func getTty() *os.File {
	device := "/dev/tty"
	tty, err := os.Create(device)
	if err != nil {
		fmt.Printf("File Open Error device:%s error:%v\n", device, err)
	}
	return tty
}

func newMonitor(ctx context.Context, cancel func(), sk *speedKeeper) *monitor {
	return &monitor{
		ctx:      ctx,
		cancel:   cancel,
		tty:      getTty(),
		sk:       sk,
		progress: make(chan struct{}),
	}
}

func (mon *monitor) setMode(mode string) {
	mon.mode = mode
}

func (mon *monitor) standardProgress() {
	p := ""
	if mon.sk.size > 0 {
		p = fmt.Sprintf("(%3d%%)", int(mon.sk.current*100/mon.sk.size))
	}
	fmt.Fprintf(mon.tty, "\r\033[K[%s] %dBytes%s\t@ %dKBps",
		time.Now().Format("2006/01/02 15:04:05.000 MST"),
		mon.sk.current,
		p,
		mon.sk.currentSpeed()/1024)
}

func (mon *monitor) getGraphProgress() func() {
	var bar string

	return func() {
		p := ""
		if mon.sk.size > 0 {
			p = fmt.Sprintf("(%3d%%)", int(mon.sk.current*100/mon.sk.size))
			bar = bar + "*"
		} else {
			bar = bar + "*"
		}

		fmt.Fprintf(mon.tty, "\r\033[K[%s] %dBytes%s\t@ %dKBps\t%s",
			time.Now().Format("2006/01/02 15:04:05.000 MST"),
			mon.sk.current,
			p,
			mon.sk.currentSpeed()/1024,
			bar,
		)
	}
}

type progesser interface {
	initFunc()
	pFunc()
	endFunc()
}

type kara struct{}

func (mon *monitor) run() {
	var initFunc, pFunc, endFunc func()
	//var p progresser
	switch mon.mode {
	case "silent":
		//p = silentProgresser
		p = &struct{}{
			initFunc: func() {},
			pFunc:    func() {},
			endFunc:  func() {},
		}
	case "graph":
		p = &struct{}{
			initFunc: func() {},
			pFunc:    mon.getGraphProgress(),
			endFunc: func() {
				fmt.Fprintf(mon.tty, "\n")
			},
		}
		//initFunc = func() {}
		//pFunc = mon.getGraphProgress()
		//endFunc = func() {
		//fmt.Fprintf(mon.tty, "\n")
		//}
	default:
		p = &struct{}{
			initFunc: func() {},
			pFunc:    mon.standardProgress,
			endFunc: func() {
				fmt.Fprintf(mon.tty, "\n")
			},
		}
		//initFunc = func() {}
		//pFunc = mon.standardProgress
		//endFunc = func() {
		//fmt.Fprintf(mon.tty, "\n")
		//}
	}
L:
	p.initFunc()
	for {
		select {
		case <-mon.progress:
			p.pFunc()
			//pFunc()
			//if mon.sk.current == mon.sk.size {
			//mon.cancel()
			//}
		case <-mon.ctx.Done():
			break L
		}
	}
	p.endFunc()
}
