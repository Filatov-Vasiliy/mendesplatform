package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"main/internal/storage"
)

func main() {
	fmt.Println("hello")
	// TODO init config

	// TODO init logger

	// TODO init Kafka producer

	// TODO init kafka consumer

	// TODO init database
	storage.New("")
	fmt.Println("hello")
	// TODO init server
}
