# DaServe
`daserve` An extremely basic web server for serving static content

## Install
Clone the repo and run `go install` from inside the cloned directory

## Usage
To start with default settings (and content), simply run:

`daserve`

## Help

```
$ daserve --help
Runs a basic static content web-server, configure using the following flags:

  -d string
        Directory to serve (default "./static")
  -h string
        Host address to bind (default "127.0.0.1")
  -p string
        Port to listen on (default "9080")
```