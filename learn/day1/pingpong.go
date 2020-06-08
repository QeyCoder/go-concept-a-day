package main

import (
	"fmt"
	"time"
)

func main() {
	//i:=1000
	pingChan := make(chan string)
	pongChan := make(chan string)

	for {
		go ping(pingChan)
		go pong(pongChan)
		fmt.Println(<-pingChan)
		fmt.Println(<-pongChan)
	}
}

func ping(ch chan string) {

	time.Sleep(time.Second)
	ch <- "Ping"

}

func pong(ch chan string) {
	time.Sleep(time.Second)
	ch <- "Pong"
}
