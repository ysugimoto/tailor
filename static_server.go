package main

import (
	"fmt"
	"net/http"
	"strings"
)

type StaticServer struct {
}

type StaticResponse struct {
	Content    []byte
	MimeType   string
	StatusCode int
}

func (s StaticServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	info, err := s.FindStaticFile(req.URL.Path)
	if err != nil {
		info.Content, _ = Asset("assets/notfound.html")
		info.MimeType = "text/html"
	}

	resp.Header().Set("Content-Type", fmt.Sprintf("%s; charset=UTF-8", info.MimeType))
	resp.Header().Set("Content-Length", fmt.Sprint(len(info.Content)))
	resp.WriteHeader(info.StatusCode)
	resp.Write(info.Content)
}

func (s StaticServer) FindStaticFile(path string) (info StaticResponse, err error) {
	if path == "/" {
		path = "assets/index.html"
	} else {
		path = "assets" + path
	}

	info.Content, err = Asset(path)
	if err != nil {
		return info, err
	}
	info.StatusCode = http.StatusOK
	if strings.Contains(path, ".html") {
		info.MimeType = "text/html"
	} else if strings.Contains(path, ".css") {
		info.MimeType = "text/css"
	} else if strings.Contains(path, ".js") {
		info.MimeType = "text/javascript"
	}

	return info, nil
}
