package main

import (
	"flag"
	"fmt"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/handlers"
	"github.com/pkg/browser"
)

func usage() {
	msg := `Runs a basic web-server for serving static content

'dir' is the directory to serve (default "./static")

Usage:
  daserve [flags] [dir]

Flags:
`
	fmt.Fprint(os.Stderr, msg)
	flag.PrintDefaults()
}

func main() {
	var host string
	var port string
	var dir string
	var gui bool
	var redirect404 string
	var redirt404ToIndex bool

	flag.Usage = usage

	flag.StringVar(&host, "h", "127.0.0.1", "Host address to bind")
	flag.StringVar(&port, "p", "9080", "Port to listen on\n")
	flag.StringVar(&redirect404, "404", "", "On 404, serve up this page (and change status to 200), ie: \"/index.html\"")
	flag.BoolVar(&redirt404ToIndex, "404i", false, "When set will serve up /index.html on 404s, does nothing when -404 also set\nUseful when history.pushState used instead of # for SPA paths (history mode)")
	flag.BoolVar(&gui, "g", false, "Opens default browser on launch")
	flag.Parse()

	// ref: https://github.com/golang/go/issues/32350
	mime.AddExtensionType(".js", "application/javascript")

	dir = "./static"
	if len(flag.Args()) == 1 {
		dir = flag.Args()[0]
	}

	address := fmt.Sprintf("%s:%s", host, port)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	if redirect404 != "" {
		path := filepath.Join(dir, redirect404)
		http.HandleFunc("/", wrapHandler(http.FileServer(http.Dir(dir)), path))
	} else if redirt404ToIndex {
		path := filepath.Join(dir, "/index.html")
		http.HandleFunc("/", wrapHandler(http.FileServer(http.Dir(dir)), path))
	} else {
		http.Handle("/", http.FileServer(http.Dir(dir)))
	}

	fmt.Printf("Listening on %v serving from %v\n\n", address, dir)
	if gui {
		go browser.OpenURL(fmt.Sprintf("http://%s", address))
	}

	http.ListenAndServe(address, handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
