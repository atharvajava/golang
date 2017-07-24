/*
notFoundHandler returns 404
RedirectHandler redirects to different path and returns its status
Strip Prefix handler : it is kind of like middleware , it decorates another handler
Timeout Handler : if it takes too long to load it times out and responds the given mesage
FileServer handler : returns file with an optional err

we have type Dir and use filesystem interface where filesystem is gonna serve from
*/

package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("public" + r.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		defer f.Close()
		io.Copy(w, f)
	})

	http.ListenAndServe(":3000", nil)
}
