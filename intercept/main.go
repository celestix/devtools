package main

import (
	"os"
	"fmt"
	"io"
	"net/http"
)

var domain = os.Getenv("INTERCEPT_DOMAIN") 

func main() {
	fmt.Println("DOMAIN:", domain)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.RequestURI
		method := r.Method
		uri := domain + path
		fmt.Println("URI:", uri)
		req, err := http.NewRequest(method, uri, nil)
		if err != nil {
			return
		}
		for k, v := range r.Header {
			req.Header.Set(k, v[0])
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		for k, v := range resp.Header {
			if k == "Content-Length" {
				continue
			}
			w.Header().Set(k, v[0])
		}
		w.WriteHeader(resp.StatusCode)
		_, err = io.Copy(w, resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	http.ListenAndServe(":4744", mux)
}
