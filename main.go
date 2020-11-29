package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

var host string
var port string
var dir string

func usage() {
	fmt.Fprintf(os.Stderr, "Runs a basic static content web-server, configure using the following flags:\n\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage

	flag.StringVar(&host, "h", "127.0.0.1", "Host address to bind")
	flag.StringVar(&port, "p", "9080", "Port to listen on")
	flag.StringVar(&dir, "d", "./static", "Directory to serve")
	flag.Parse()

	address := fmt.Sprintf("%s:%s", host, port)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)

	fmt.Printf("Listening on %v serving from %v\n\n", address, dir)
	http.ListenAndServe(address, handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
