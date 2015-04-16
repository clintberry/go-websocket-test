package main

import (
	"flag"
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

const (
	ORIGIN = "http://go-websocket-test"
)

func main() {
	url := flag.String("url", "", "url of the websocket server")
	connections := flag.Int("connections", 10000, "number of connections")

	flag.Parse()
	connectionChan := make(chan *websocket.Conn)

	wsMap := make(map[int]*websocket.Conn)

	// Go routine to read connections from my connection channel (thread safe)
	go func() {
		i := 0
		for {
			select {
			case ws := <-connectionChan:
				wsMap[i] = ws
				i++
			}
		}
	}()

	for i := 0; i < *connections; i++ {
		go func() {
			var err error
			ws, err := websocket.Dial(*url, "", ORIGIN)
			if err != nil {
				fmt.Println("Error connecting...")
				log.Fatal(err)
			}
			connectionChan <- ws
		}()
	}

	quit := make(chan bool)
	<-quit
}
