package main

import (
	"fmt"
	"github.com/lqqyt2423/go-mitmproxy/proxy"
	"io/ioutil"
	"net/http"
	"strings"
)

type RewriteHost struct {
	proxy.BaseAddon
}

func (a *RewriteHost) Requestheaders(f *proxy.Flow) {
	if strings.Contains(f.Request.URL.Host, "naver.com") {
		f.Request.URL.Host = "localhost:51231"
		f.Request.URL.Scheme = "http"
	}
}

func runWebBackend() {
	http.HandleFunc("/", handleRequest)
	if err := http.ListenAndServe(":51231", nil); err != nil {
		panic(err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	printRequestDetails(r)
	writeResponseDetails(w, r)
}

func printRequestDetails(r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Printf("%v: %v\n", name, h)
		}
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}
	defer r.Body.Close()

	fmt.Println("Body:", string(bodyBytes))
}

func writeResponseDetails(w http.ResponseWriter, r *http.Request) {
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	w.Write([]byte(fmt.Sprintf("Body: %s\n\nURL: %s\n\nHeaders:\n", bodyBytes, r.URL)))
	for name, headers := range r.Header {
		for _, h := range headers {
			w.Write([]byte(fmt.Sprintf("%v: %v\n", name, h)))
		}
	}
}

func main() {
	go runWebBackend()

	opts := &proxy.Options{
		Addr:              ":9080",
		StreamLargeBodies: 1024 * 1024 * 5,
	}

	p, err := proxy.NewProxy(opts)
	if err != nil {
		panic(err)
	}

	p.AddAddon(&RewriteHost{})
	p.AddAddon(&proxy.LogAddon{})

	p.Start()
}
