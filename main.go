package main

import (
	"chproxy/proxy"
	"log"
)

func main() {
	server := proxy.NewServer()

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
