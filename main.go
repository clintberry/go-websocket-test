// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	go h.run()
	serverMap := make(map[int]*http.ServeMux)

	for i := 8001; i <= 8050; i++ {
		port := i
		serverMap[port] = http.NewServeMux()
		serverMap[port].HandleFunc("/ws", serveWs)
		go func(server *http.ServeMux, port int) {
			fmt.Println("Listening for websocket connections on port " + strconv.Itoa(port))
			err := http.ListenAndServe("localhost:"+strconv.Itoa(port), server)
			if err != nil {
				log.Fatal("ListenAndServe: ", err)
			}
		}(serverMap[port], port)
	}

	quit := make(chan bool)
	<-quit

}
