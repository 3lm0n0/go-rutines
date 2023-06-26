package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	userName := getUsername()
	fmt.Printf("userName: %s and took: %v \n", userName, time.Since(start))

	likes := getUserLikes(userName)
	match := matchUserName(userName)

	fmt.Println("likes: ", likes)
	fmt.Println("match: ", match)
	fmt.Println("main took: ", time.Since(start))
}

func getUsername() string {
	time.Sleep(time.Millisecond * 100)

	return "Diego"
}

func getUserLikes(userName string) int {
	time.Sleep(time.Millisecond * 150)

	return 11

} 

func matchUserName(userName string) string {
	time.Sleep(time.Millisecond * 100)

	return "Lionel"
}