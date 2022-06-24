package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	currentTime := time.Now()

	chanel := make(chan string)

	servers := []string{
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, server := range servers {
		go checkServer(server, chanel)
	}

	for i := 0; i < len(servers); i++ {
		fmt.Println(<-chanel)
	}

	pastTime := time.Since(currentTime)
	fmt.Printf("Execution time: %s\n", pastTime)
}

func checkServer(server string, chanel chan string) {
	_, err := http.Get(server)

	if err != nil {
		chanel <- server + " is not available"
	} else {
		chanel <- server + " is available"
	}
}
