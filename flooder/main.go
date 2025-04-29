package main

import "log"
import "net/http"
import "fmt"
import "sync"


func makeRequest() {
        resp, err := http.Get("https://nginx.wlan0.in/login")
        if err != nil {
                log.Println("something wrong with you bitch: ", err.Error())
                return
        }
        defer resp.Body.Close()
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
                                makeRequest()
                        }()
                }
                wg.Wait()
                n++
                fmt.Println(n, "batch done")
        }
}
