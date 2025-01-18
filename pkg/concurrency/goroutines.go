package concurrency

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func returnType(url string){
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error %s", err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("%v -> %v\n", url, resp.Header.Get("Content-Type"))
}

func SiteSerial(urls []string){
	for _, url := range urls {
		returnType(url)
	}
}

func SiteConcurrent(urls []string){
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			returnType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func WorkingWithSerialExecutionUrlFetch(){
	urls := []string{"http://www.google.com", "http://www.yahoo.com", "http://www.bing.com"}
	start := time.Now()
	SiteSerial(urls)
	fmt.Println("Time taken for serial:", time.Since(start))
}

func WorkingWithGoRoutinesUrlFetch(){
	urls := []string{"http://www.google.com", "http://www.yahoo.com", "http://www.bing.com"}
	start := time.Now()
	SiteConcurrent(urls)
	fmt.Println("Time taken for concurrent url content header fetch:", time.Since(start))
}