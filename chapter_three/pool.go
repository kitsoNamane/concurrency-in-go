package chapter_three

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func Pool() {
	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated += 1
			mem := make([]byte, 1024)
			return &mem
		},
	}

	// seed the pool with 4KB
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()

			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)

			// Assume something interesting, but quick is being done with this memory
		}()
	}

	wg.Wait()
	fmt.Printf("%d calculators were created\n", numCalcsCreated)
}

func ConnectToService() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{}
}

func StartNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		server, err := net.Listen("tcp", "localhost:9090")
		if err != nil {
			log.Fatalf("cannot listen: %v\n", err)
		}
		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v\n", err)
				continue
			}
			ConnectToService()
			fmt.Fprintf(conn, "")
			conn.Close()
		}
	}()

	return &wg
}
