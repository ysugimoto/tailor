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
	queue []Payload

	// State of request is sending
	isSending bool
}

// Send the message to server
func (r *Remote) Send(p Payload) {
	if r.isSending {
		r.queue = append(r.queue, p)
		return
	}
	r.isSending = true
	defer func() {
		r.isSending = false
		if len(r.queue) > 0 {
			r.Send(r.queue[0])
		}
	}()
	post := url.Values{}
	post.Set("message", p.Message)
	post.Set("host", p.Host)
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
