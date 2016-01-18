package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"os"
)

// Tail library wrapper
// this function should be called by goroutine
func startTail(fileName string, readCallback func(string)) error {
	if _, err := os.Stat(fileName); err != nil {
		return err
	}

	conf := tail.Config{
		Follow: true,
		Poll:   true,
		Logger: tail.DiscardingLogger,
	}

	host, _ := os.Hostname()
	if t, err := tail.TailFile(fileName, conf); err != nil {
		return err
	} else {
		for line := range t.Lines {
			msg := fmt.Sprintf("[%s] %s: %s", host, fileName, line.Text)
			fmt.Println(msg)
			readCallback(msg)
		}
	}

	return nil
}
