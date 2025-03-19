package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const MAX_BODY_SIZE = 1024 * 1024 // 1MB

func main() {
	log.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("<====================================>")
		fmt.Println("Time: ", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println("Received request from", r.RemoteAddr)
		fmt.Println("Request method:", r.Method)
		fmt.Println("Request URI:", r.RequestURI)
		fmt.Println(r.RemoteAddr, "requested", r.RequestURI)
		fmt.Println("Request headers:")
		for header := range r.Header {
			fmt.Println("Header", "\""+header+"\"", ":", r.Header.Get(header))
		}
		fmt.Println("Request content length:", r.ContentLength)
		if r.ContentLength == 0 {
			fmt.Println("<====================================>")
			return
		}
		fmt.Println("Request body:")
		defer r.Body.Close()
		var bufLen = r.ContentLength
		if bufLen > MAX_BODY_SIZE {
			log.Println("req_bod_too_larg... found:", bufLen, "; expected:", MAX_BODY_SIZE)
			bufLen = MAX_BODY_SIZE
			return
		}
		buf := make([]byte, bufLen)
		_, err := r.Body.Read(buf)
		if err != nil {
			log.Println("Error: ", err.Error())
			return
		}
		fmt.Println(string(buf))
		fmt.Println("<====================================>")
		w.WriteHeader(http.StatusOK)
	}))
}
