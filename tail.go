package main

import (
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
		ReOpen: true,
		Poll:   true,
		Logger: tail.DiscardingLogger,
	}

	if t, err := tail.TailFile(fileName, conf); err != nil {
		return err
	} else {
		for line := range t.Lines {
			readCallback(line.Text)
		}
	}

	return nil
}
