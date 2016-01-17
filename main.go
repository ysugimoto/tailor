package main

import (
	"bufio"
	"fmt"
	"github.com/ysugimoto/go-cliargs"
	"golang.org/x/net/websocket"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type AppHandler struct {
	connections map[string]*Connection
}

func (a *AppHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	var handler websocket.Handler

	switch req.URL.Path {
	case "/remote":
		a.handleRemoteRequest(resp, req)
		return
	case "/proxy":
		handler = createProxyClient(a)
	default:
		handler = createReader(a)
	}
	ws := websocket.Server{Handler: handler}
	ws.ServeHTTP(resp, req)
}

func (a *AppHandler) handleRemoteRequest(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Content-Type", "text/plain")
	req.ParseForm()
	if msg := req.Form.Get("message"); msg == "" {
		a.Broadcast(msg)
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte("Message has accepted."))
	} else {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte("Message is empty."))
	}
}

func (a *AppHandler) Broadcast(line string) {
	for _, c := range a.connections {
		if c.Type == READER {
			c.Send(line)
		}
	}
}

func sendRemote(proxyURL, message string) {
	post := url.Values{}
	post.Add("message", message)
	request, _ := http.NewRequest(
		"POST",
		proxyURL,
		strings.NewReader(post.Encode()),
	)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	response, _ := client.Do(request)
	response.Body.Close()
}

func main() {
	args := cliarg.NewArguments()
	args.Alias("", "stdin", nil)
	args.Alias("", "proxy", "")
	args.Alias("p", "port", "9000")
	args.Alias("", "proxy-server", nil)
	args.Parse()

	if _, ok := args.GetOption("stdin"); ok {
		fmt.Println("Read from stdin")
		scanner := bufio.NewScanner(os.Stdin)
		proxy, _ := args.GetOptionAsString("proxy")
		for scanner.Scan() {
			if proxy != "" {
				sendRemote(proxy, scanner.Text())
			} else {
				os.Stdout.WriteString(scanner.Text() + "\n")
			}
		}
		return
	}

	app := &AppHandler{
		connections: make(map[string]*Connection),
	}
	if _, ok := args.GetOption("proxy-server"); !ok {
		file, _ := args.GetCommandAt(1)
		go startTail(file, app.Broadcast)
	}

	http.Handle("/", app)
	if err := http.ListenAndServe(":9000", nil); err != nil {
		panic(err)
	}

}
