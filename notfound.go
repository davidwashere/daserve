// Ref: https://stackoverflow.com/questions/47285119/how-to-custom-handle-a-file-not-being-found-when-using-go-static-file-server
package main

import (
	"log"
	"net/http"
)

type NotFoundRedirectRespWr struct {
	http.ResponseWriter // We embed http.ResponseWriter
	status              int
}

func (w *NotFoundRedirectRespWr) WriteHeader(status int) {
	w.status = status // Store the status for our own use
	if status != http.StatusNotFound {
		w.ResponseWriter.WriteHeader(status)
	}
}

func (w *NotFoundRedirectRespWr) Write(p []byte) (int, error) {
	if w.status != http.StatusNotFound {
		return w.ResponseWriter.Write(p)
	}
	return len(p), nil // Lie that we successfully written it
}

func wrapHandler(h http.Handler, redirectDest string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nfrw := &NotFoundRedirectRespWr{ResponseWriter: w}
		h.ServeHTTP(nfrw, r)

		// destWithQuery := redirectDest + "?" + r.URL.RawQuery
		// r.

		// log.Printf("Query: ?%s", r.URL.RawQuery)
		if nfrw.status == 404 {
			log.Printf("404: Redirecting %s to %s.", r.RequestURI, redirectDest)
			// http.Redirect(w, r, destWithQuery, http.StatusFound)
			// w.WriteHeader(http.StatusOK)
			// w.Write([]byte("Hello"))
			w.Header().Set("Content-Type", "text/html")

			http.ServeFile(w, r, redirectDest)
		}
	}
}
