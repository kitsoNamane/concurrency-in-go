package chapter_four

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func FanOutFanIn() {
	rand := func() interface{} { return rand.IntN(50_000_000) }

	done := make(chan interface{})
	defer close(done)

	start := time.Now()
	randIntStream := toInt(done, repeatFn(done, rand))

	fmt.Println("Primes:")
	for prime := range take(done, primeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v", time.Since(start))
}
