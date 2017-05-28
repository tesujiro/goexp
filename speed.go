package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

const DEBUG = 0

func dprintf(format string, a ...interface{}) {
	if DEBUG != 0 {
		fmt.Fprintf(os.Stderr, format, a...)
	}
}

func main() {
	ctx := context.Background()

	var speed *int = flag.Int("bandwidth", 0, "Bytes Per Sec.")
	flag.Parse()

	limitedPipe(ctx, os.Stdin, os.Stdout, *speed)
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

func limitedPipe(ctx context.Context, in io.Reader, out io.Writer, speed int) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)

	rbchan := make(chan readbuf, 1)
	done := make(chan struct{}, 1)
	go func() {
		read(in, rbchan)
		done <- struct{}{}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	mon := newMonitor(ctx, cancel)
	go func() {
		mon.run()
		wg.Done()
	}()

	tick := time.NewTicker(time.Millisecond * time.Duration(200)).C

	sk := NewSpeedKeeper(time.Now(), speed)
	readBytes := 0
L:
	for {
		select {
		case rb := <-rbchan:
			readBytes += rb.length
			out.Write(rb.buf)
			sk.killTime(readBytes)
		case <-tick:
			mon.msg <- fmt.Sprintf("\r\033[K[%s] %dBytes\t@ %dKBps",
				time.Now().Format("2006/01/02 15:04:05.000 MST"),
				readBytes,
				sk.currentSpeed(readBytes)/1024)
		case <-done:
			mon.msg <- fmt.Sprintf("\n")
			cancel()
			break L
		case <-ctx.Done():
			break L
		}
	}
	//<-ctx.Done()
	wg.Wait()
}

type speedKeeper struct {
	start      time.Time
	bytePerSec int
}

func NewSpeedKeeper(s time.Time, b int) *speedKeeper {
	return &speedKeeper{
		start:      s,
		bytePerSec: b,
	}
}

func (sk *speedKeeper) killTime(curBytes int) {
	if sk.bytePerSec <= 0 {
		return
	}
	//target_duration := time.Duration(float64(curBytes/sk.bytePerSec)) * time.Second //NG
	target_duration := time.Duration(float64(curBytes*1000/sk.bytePerSec)) * time.Millisecond
	//target_duration := time.Duration(float64(curBytes*1e9/sk.bytePerSec)) * time.Nanosecond
	current_duration := time.Since(sk.start)
	wait := target_duration - current_duration
	dprintf(" target=%s current=%s\n", target_duration, current_duration)
	if wait > 0 {
		dprintf("Sleep %s\n", wait)
		time.Sleep(wait)
	}
}

func (sk *speedKeeper) currentSpeed(curBytes int) int {
	//d := int(time.Since(sk.start).Seconds())
	d := int(time.Since(sk.start).Nanoseconds())
	if d == 0 {
		return 0
	}
	return curBytes * 1e9 / d
}

type monitor struct {
	ctx    context.Context
	cancel func()
	tty    io.Writer
	//tty *os.File
	msg chan string
}

func getTty() *os.File {
	device := "/dev/tty"
	tty, err := os.Create(device)
	if err != nil {
		// Openエラー処理
		fmt.Printf("File Open Error device:%s error:%v\n", device, err)
	}
	return tty
}

func newMonitor(ctx context.Context, cancel func()) *monitor {
	tty := getTty()
	return &monitor{
		ctx:    ctx,
		cancel: cancel,
		tty:    tty,
		msg:    make(chan string),
	}
}

func (mon *monitor) print(s string) {
	fmt.Fprint(mon.tty, s)
}

func (mon *monitor) run() {
L:
	for {
		select {
		case m := <-mon.msg:
			mon.print(m)
		case <-mon.ctx.Done():
			break L
		}
	}
}
