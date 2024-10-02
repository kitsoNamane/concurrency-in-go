package chapter_five

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func doWork(done <-chan interface{}) (<-chan interface{}, <-chan int) {
	heartbeatStream := make(chan interface{}, 1)
	workStream := make(chan int)

	go func() {
		defer close(heartbeatStream)
		defer close(workStream)

		for i := 0; i < 10; i++ {
			select {
			case heartbeatStream <- struct{}{}:
			default:
			}

			select {
			case <-done:
				return
			case workStream <- rand.IntN(10):
			}
		}

	}()

	return heartbeatStream, workStream
}

func DoWork(done <-chan interface{}, nums ...int) (<-chan interface{}, <-chan int) {
	heartbeat := make(chan interface{}, 1)
	intStream := make(chan int)

	go func() {
		defer close(heartbeat)
		defer close(intStream)

		time.Sleep(2 * time.Second)

		for _, n := range nums {
			select {
			case heartbeat <- struct{}{}:
			default:
			}

			select {
			case <-done:
				return
			case intStream <- n:
			}
		}
	}()

	return heartbeat, intStream
}

func HeartBeat() {
	done := make(chan interface{})
	defer close(done)

	heartbeat, results := doWork(done)
	for {
		select {
		case _, ok := <-heartbeat: // 4
			if ok == false {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-results: // 5
			if ok == false {
				return
			}
			fmt.Printf("results %v\n", r)
		}
	}
}
