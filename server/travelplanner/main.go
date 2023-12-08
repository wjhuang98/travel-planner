package main

import (
	"context"
	"fmt"
	"main/config"
	"os"
	"os/signal"

	"github.com/redis/go-redis/v9"
)

func main() {
	notifyCh := make(chan os.Signal, 1)
	listenCh := make(chan error)
	signal.Notify(notifyCh, os.Interrupt)

	context := context.Background()
	cache := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	tp := TravelPlanner{
		config:  config.GetConfig(),
		cache:   cache,
		context: context,
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
