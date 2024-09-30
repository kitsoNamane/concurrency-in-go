package chapter_four

import (
	"bytes"
	"fmt"
	"sync"
)

func AdHocConfinement() {
	data := make([]int, 4)

	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			data[i] = i + 1
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}

func LexicalConfinement() {
	chanOwner := func(data []int) <-chan int {
		results := make(chan int, 4)
		go func() {
			defer close(results)
			for i := range data {
				data[i] = i + 1
				results <- data[i]
			}
		}()

		return results
	}

	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
	}

	data := make([]int, 4)

	results := chanOwner(data)
	consumer(results)
}

func NonConcorrentSafeConfinement() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()

		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])

	wg.Wait()
}
