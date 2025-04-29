package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

func makeRequest() error {
	resp, err := http.Get(os.Getenv("TARGET_URL"))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func main() {
	numGoroutines := 100
	n := 0
	for {
		wg := sync.WaitGroup{}
		wg.Add(numGoroutines)
		for i := 0; i < numGoroutines; i++ {
			go func() {
				defer wg.Done()
				err := makeRequest()
				if err != nil {
					log.Println("something wrong with you: ", err.Error())
				} else {
					log.Println("successfull req sent")
				}
			}()
		}
		wg.Wait()
		n++
		fmt.Println(n, "batch done")
	}
}
