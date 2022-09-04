// main principal package
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	log.Println("run to port:80")
	http.ListenAndServe("0.0.0.0:80", nil)
}
