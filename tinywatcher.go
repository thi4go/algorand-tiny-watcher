package main

import (
	"algorand-tiny-watcher/server"
	"log"
)

func main() {
	tinyWatcher, err := server.New()
	if err != nil {
		log.Fatal(err)
	}

	tinyWatcher.Start()
}
