package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"os"
)

// Tail library wrapper
// this function should be called by goroutine
func startTail(fileName string, readCallback func(Payload)) error {
	conf := tail.Config{
		Location: &tail.SeekInfo{
			Whence: os.SEEK_END,
		},
		Follow:    true,
		Poll:      true,
		MustExist: true,
		Logger:    tail.DiscardingLogger,
	}

	host, _ := os.Hostname()
	if t, err := tail.TailFile(fileName, conf); err != nil {
		fmt.Println("[ERROR]", err)
	} else {
		for line := range t.Lines {
			msg := fmt.Sprintf("%s: %s", fileName, line.Text)
			fmt.Println(msg)
			readCallback(Payload{
				Message: msg,
				Host:    host,
			})
		}
	}

	return nil
}
