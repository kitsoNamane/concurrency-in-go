package chapter_one

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

// THE COFFMAN CONDITIONS
// These are conditions that must be present for a deadlock to occur. They were enumerated by Edgar Coffman:
// Mutual Exclusion - A concurrent process holds exclusive rights to a resource at any one time.
// Wait For Conditions - A concurrent process must simultaneously hold a a resource and be waiting for
// an additional resource.
// No Preemption - A resource held by a concurrent process can only be released by that process, so it
// fulfills this condition.
// Circular Wait - A concurrent process (P1) must be waiting on a chain of other concurrent process (P2),
// which are in turn waiting on it (P1), so it fulfills this final condition too
func DeadLock() {
	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()

		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)

		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}

func LiveLock() {
	cadence := sync.NewCond(&sync.Mutex{})
	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast()
		}
	}()

	takeStep := func() {
		cadence.L.Lock()
		cadence.Wait()
		cadence.L.Unlock()
	}

	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, " %v", dirName)
		atomic.AddInt32(dir, 1)
		takeStep()
		if atomic.LoadInt32(dir) == 1 {
			fmt.Fprint(out, ". Success")
			return true
		}
		takeStep()
		atomic.AddInt32(dir, -1)
		return false
	}

	var left, right int32
	tryLeft := func(out *bytes.Buffer) bool { return tryDir("left", &left, out) }
	tryRight := func(out *bytes.Buffer) bool { return tryDir("right", &right, out) }

	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer func() { fmt.Println(out.String()) }()
		defer walking.Done()

		fmt.Fprintf(&out, "%v is trying to scoot:", name)
		for i := 0; i < 5; i++ {
			if tryLeft(&out) || tryRight(&out) {
				return
			}
		}
		fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation!", name)
	}

	var peopleInHallway sync.WaitGroup
	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "Alice")
	go walk(&peopleInHallway, "Barbara")
	peopleInHallway.Wait()
}
