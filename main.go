package main

import "github.com/kitsoNamane/concurrency-in-go/chapter_one"

func main() {

	for range 100 {
		chapter_one.DeadLock()
	}
}
