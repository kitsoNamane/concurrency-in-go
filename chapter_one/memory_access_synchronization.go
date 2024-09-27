package chapter_one

import (
	"fmt"
	"sync"
)

func MemoryAccessSync() {
	var data int
	var memoryAccess sync.Mutex

	go func() {
		memoryAccess.Lock()
		data++
		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()
	if data == 0 {
		fmt.Println("value of data is: 0")
	} else {
		fmt.Printf("value of data is: %v\n", data)
	}
	memoryAccess.Unlock()
}
