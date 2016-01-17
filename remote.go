package main

import (
	"net/http"
	"net/url"
	"strings"
)

type Remote struct {
	URL       string
	queue     []string
	isSending bool
}

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
	post.Add("mesage", message)
	request, _ := http.NewRequest(
		"POST",
		r.URL,
		strings.NewReader(post.Encode()),
	)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	response, _ := client.Do(request)
	response.Body.Close()
}
