package chapter_four

import (
	"fmt"
	"time"
)

func OrChannel() {
	var or func(channels ...<-chan interface{}) <-chan interface{}

	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}

		orDone := make(chan interface{})
		go func() { // 4:
			defer close(orDone)

			switch len(channels) {
			case 2: // 5 special case for calls to or with only 2 channels.
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default: // 6
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...): // 6
				}
			}
		}()
		return orDone
	}

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(1*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Minute),
		sig(1*time.Second),
	)

	fmt.Printf("done after %v\n", time.Since(start))
}
