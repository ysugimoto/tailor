package main

import (
	"github.com/hpcloud/tail"
	"os"
)

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
	t, err := tail.TailFile(fileName, conf)
	if err != nil {
		return err
	}

	for line := range t.Lines {
		readCallback(line.Text)
	}

	return nil
}
