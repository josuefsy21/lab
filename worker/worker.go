package worker

import (
	"fmt"
	"time"
)

func Worker(workerId int, data chan int) {
	for x := range data {
		fmt.Sprintf("Worker %d recebeu %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}