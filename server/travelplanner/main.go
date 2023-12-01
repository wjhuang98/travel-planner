package main

import (
	"fmt"
	"main/config"
	"os"
	"os/signal"
)

func main() {
	notifyCh := make(chan os.Signal, 1)
	listenCh := make(chan error)
	signal.Notify(notifyCh, os.Interrupt)

	tp := TravelPlanner{
		config: config.GetConfig(),
	}

	go func() {
		listenCh <- tp.Run()
	}()

	select {
	case <-notifyCh:
		fmt.Printf("recieved interrupt signal")
	case err := <-listenCh:
		fmt.Printf("recieved error: %v", err)
	}
}
