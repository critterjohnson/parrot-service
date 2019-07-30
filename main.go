package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// copy headers
		headers := w.Header()
		for key, val := range r.Header {
			headers.Add(key, val[0])
		}

		// copy body
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalf("Parrot: error %v", err)
		}
		w.Write(body)
	})

	http.ListenAndServe(":3000", nil)
}
