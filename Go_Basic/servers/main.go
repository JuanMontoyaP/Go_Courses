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

	i := 0
	for i < 9 {

		for _, server := range servers {
			go checkServer(server, chanel)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-chanel)
		i++
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
