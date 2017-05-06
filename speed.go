package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

const DEBUG = 0

func dprintf(format string, a ...interface{}) {
	if DEBUG != 0 {
		fmt.Fprintf(os.Stderr, format, a...)
	}
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

func main() {
	var speed *int = flag.Int("bandwidth", 0, "Bytes Per Sec.")
	flag.Parse()

	tty := getTty()
	defer tty.Close()

	limitedPipe(os.Stdin, os.Stdout, tty, *speed)
}

const BUFSIZE = 4096

func limitedPipe(in io.Reader, out io.Writer, tty io.Writer, speed int) {
	sk := NewSpeedKeeper(time.Now(), speed)
	reader := bufio.NewReader(in)
	buf := make([]byte, BUFSIZE)
	readBytes := 0
	for {
		n, err := reader.Read(buf)
		//fmt.Fprintf(os.Stderr, "buffer len :%d\n", n)
		if n == 0 {
			break
		}
		if err != nil {
			break
		}
		out.Write(buf)
		readBytes += n
		if speed > 0 {
			sk.killTime(readBytes)
		}
		//dprintf("\r[%s] %d Bytes\t@%2d KBps\n",
		fmt.Fprintf(tty, "\r[%s] %d Bytes\t@%d KBps",
			time.Now().Format("2006/01/02 15:04:05.000 MST"),
			readBytes,
			sk.currentSpeed(readBytes)/1024)
	}
	fmt.Fprintf(os.Stderr, "\n")
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
