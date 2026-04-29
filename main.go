package main

import (
	//"fmt"
	//"time"
	"lab/worker"
)

func main() {
	cn := make(chan int)
	qtdWorkers := 3

	for y := range qtdWorkers{
		go worker.Worker(y, cn)
	}

	//fmt.Println(<-cn)

	//time.Sleep(time.Millisecond)
	//fmt.Println("pong")

	for y := range 95 {
		cn <- y
	}
}
