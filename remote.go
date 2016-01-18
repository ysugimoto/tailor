package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Remote request struct
// support message queueing
type Remote struct {

	// Request URL
	URL string

	// Queueing message stack
	queue []string

	// State of request is sending
	isSending bool
}

// Send the message to server
func (r *Remote) Send(message string) {
	if r.isSending {
		r.queue = append(r.queue, message)
		return
	}
	r.isSending = true
	defer func() {
		r.isSending = false
		if len(r.queue) > 0 {
			r.Send(strings.Join(r.queue, "\n"))
		}
	}()
	post := url.Values{}
	post.Set("message", message)
	request, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/remote", r.URL),
		strings.NewReader(post.Encode()),
	)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
}
