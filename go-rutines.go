package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	
	userName := getUsername()
	fmt.Printf("userName: %s and took: %v \n", userName, time.Since(start))

	responseChannel := make(chan any, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go getUserLikes(userName, responseChannel, wg)
	go matchUserName(userName, responseChannel, wg)

	wg.Wait() // blocks until 2 wg.Done()
	close(responseChannel)

	for response := range responseChannel {
		fmt.Println("response: ", response)
	}
	
	fmt.Println("main took: ", time.Since(start))
}

func getUsername() string {
	time.Sleep(time.Millisecond * 100)

	return "Diego"
}

func getUserLikes(userName string, responseChannel chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)

	responseChannel <- 11
	wg.Done()
} 

func matchUserName(userName string, responseChannel chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)

	responseChannel <- "Lionel"
	wg.Done()
}