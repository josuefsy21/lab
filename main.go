package main

import "github.com/lab/worker"

func main() {
	cn := make(chan int)

	go worker.Worker(5, cn)
}