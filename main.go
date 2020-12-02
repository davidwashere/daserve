package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

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
	fmt.Fprintf(os.Stderr, msg)
	flag.PrintDefaults()
}

func main() {
	var host string
	var port string
	var dir string
	var gui bool

	flag.Usage = usage

	flag.StringVar(&host, "h", "127.0.0.1", "Host address to bind")
	flag.StringVar(&port, "p", "9080", "Port to listen on")
	flag.BoolVar(&gui, "g", false, "Opens default browser on launch")
	flag.Parse()

	dir = "./static"
	if len(flag.Args()) == 1 {
		dir = flag.Args()[0]
	}

	address := fmt.Sprintf("%s:%s", host, port)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)

	fmt.Printf("Listening on %v serving from %v\n\n", address, dir)
	if gui {
		go browser.OpenURL(fmt.Sprintf("http://%s", address))
	}

	http.ListenAndServe(address, handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
