package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"os"
	"strings"
	"sync"
	"time"
)

var bufferStack []string
var lock = new(sync.Mutex)

// Tail library wrapper
// this function should be called by goroutine
func startTail(fileName string, readCallback func(Payload)) error {
	conf := tail.Config{
		Location: &tail.SeekInfo{
			Whence: os.SEEK_END,
		},
		Follow:    true,
		Poll:      true,
		MustExist: false,
		Logger:    tail.DiscardingLogger,
	}

	host, _ := os.Hostname()
	if t, err := tail.TailFile(fileName, conf); err != nil {
		fmt.Println("[ERROR]", err)
	} else {
		go func() {
			for line := range t.Lines {
				lines := appendStack(line.Text)
				if lines > 5 {
					readCallback(Payload{
						Message: getStack(),
						Host:    host,
						Time:    time.Now().Format("2006-01-02 15:03:04"),
					})
				}
			}
		}()
		timer := time.NewTicker(5 * time.Second)
		for {
			select {
			case <-timer.C:
				buf := getStack()
				if buf != "" {
					readCallback(Payload{
						Message: buf,
						Host:    host,
						Time:    time.Now().Format("2006-01-02 15:03:04"),
					})
				}
			}
		}
	}

	return nil
}

func appendStack(buffer string) int {
	lock.Lock()
	bufferStack = append(bufferStack, buffer)
	lock.Unlock()

	return len(bufferStack)
}

func getStack() (buf string) {
	lock.Lock()
	buf = strings.Join(bufferStack, "\n")
	bufferStack = []string{}
	lock.Unlock()

	return
}
