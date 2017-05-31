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

func main() {
	ctx := context.Background()

	var speed *int = flag.Int("bandwidth", 0, "Bytes Per Sec.")
	flag.Parse()

	switch len(flag.Args()) {
	case 0:
		limitedPipe(ctx, os.Stdin, os.Stdout, *speed, 0)
	case 1:
		file, fileinfo, err := openfile(flag.Args()[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "\n\nFile Open Error :%v\n", err)
			os.Exit(9)
		}
		defer file.Close()
		limitedPipe(ctx, file, os.Stdout, *speed, int(fileinfo.Size()))
	default:
		fmt.Fprintf(os.Stderr, "\n\nParameter Error\n") // Todo read more than one file at once
		os.Exit(9)

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
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "\n\nFile Read Error :%v\n", err)
			os.Exit(9)
		} else {
			rb <- readbuf{length: n, buf: buf}
		}
	}
}

func limitedPipe(ctx context.Context, in io.Reader, out io.Writer, speed int, size int) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)

	rbchan := make(chan readbuf, 1)
	done_reading := make(chan struct{}, 1)
	go func() {
		read(in, rbchan)
		done_reading <- struct{}{}
	}()

	wg := &sync.WaitGroup{}

	//sk := NewSpeedKeeper(time.Now(), speed, size)
	sk := NewSpeedKeeper(ctx, cancel, time.Now(), speed, size)
	wg.Add(1)
	go func() {
		sk.run()
		wg.Done()
	}()

	mon := newMonitor(ctx, cancel, sk)
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
		case <-done_reading:
			mon.progress <- struct{}{}
			cancel()
			break L
		case <-ctx.Done():
			break L
		}
	}
	//<-ctx.Done()
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
	outchan    chan struct{}
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
	progress chan struct{}
	sk       *speedKeeper
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
		progress: make(chan struct{}),
		sk:       sk,
	}
}

func (mon *monitor) printProgress() {
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

func (mon *monitor) run() {
L:
	for {
		select {
		case <-mon.progress:
			mon.printProgress()
			//if mon.sk.current == mon.sk.size {
			//mon.cancel()
			//}
		case <-mon.ctx.Done():
			fmt.Fprintf(mon.tty, "\n")
			break L
		}
	}
}
