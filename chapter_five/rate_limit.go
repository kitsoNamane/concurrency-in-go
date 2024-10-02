package chapter_five

import (
	"context"
	"log"
	"os"
	"sync"
)

type APIConnection struct{}

func Open() *APIConnection {
	return &APIConnection{}
}

func (a *APIConnection) DeadFile(ctx context.Context) error {
	// Pretend we do work here
	return nil
}

func (a *APIConnection) ResolveAddress(ctx context.Context) error {
	// Pretend we do work here
	return nil
}

func RateLimit() {
	defer log.Println("Done...")

	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	apiConnection := Open()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			err := apiConnection.DeadFile(context.Background())
			if err != nil {
				log.Printf("cannot read file: %v\n", err)
			}
			log.Println("read file")
		}()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			err := apiConnection.ResolveAddress(context.Background())
			if err != nil {
				log.Printf("cannot resolve address %v\n", err)
			}
			log.Println("resolved address")
		}()
	}
	wg.Wait()
}
