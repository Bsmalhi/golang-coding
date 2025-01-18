package concurrency

import (
	"fmt"
	"net/http"
	"time"
)

/*
	Great Article on GoLang buffered and unbuffered channels and there use cases
	https://everythingcoding.in/channels-in-golang/
*/

func WorkingWithChannelsDeadlockExample(){
	ch := make(chan int) // unbuffered channel
	ch <- 5 // this will cause deadlock as main routine is waiting for value to be read from channel
	val := <- ch
	fmt.Println("Value from channel is ", val)
}

func WorkingWithChannelGoRoutine(){
	ch := make(chan int) // unbuffered channel
	go func(){
		for i := 0; i < 5; i++ {
			fmt.Printf("Sending %v to channel\n", i)
			ch <- i
		}
	}() 
	for i := 0; i < 5; i++ {
		val := <- ch
		fmt.Printf("Received %v from channel\n", val)
	}
	close(ch)
}

func returnTypeWithChannel(url string, out chan string){
	resp, err := http.Get(url)
	if err != nil {
		out <- fmt.Sprintf("Error %s", err)
		return
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("Content-Type")
	out <- fmt.Sprintf("%v -> %v\n", url, ctype)
}

func SiteConcurrentWithChannels(urls []string){
	ch := make(chan string)
	for _, url := range urls {
		go returnTypeWithChannel(url, ch)
	}
	for range urls {
		fmt.Printf("URL output %v \n", <- ch)
	}
}

func WorkingWithChannelsGoRoutinesUrlFetch(){
	urls := []string{"http://www.google.com", "http://www.yahoo.com", "http://www.bing.com"}
	start := time.Now()
	SiteConcurrentWithChannels(urls)
	fmt.Println("Time taken for concurrent url content header fetch:", time.Since(start))
}