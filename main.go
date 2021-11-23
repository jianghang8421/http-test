// github.com/bigwhite/experiments/blob/master/http-client/default-client/client.go

package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(256)
	for i := 0; i < 256; i++ {
		go func() {
			defer wg.Done()
			resp, err := http.Get("http://172.31.26.230:32083/healthz")
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			fmt.Println(string(body))
		}()
	}
	wg.Wait()
}
