package main

import (
	"ec/tsptw"
	"log"
	"math/rand"
	"time"
)

func main() {
	var seed int64
	seed = int64(time.Now().Unix())
	rand.Seed(seed)
	start := time.Now()
	data := tsptw.RC201_2
	go func() {
		for {
			select {
			case <-time.Tick(10 * time.Second):
				log.Println("processing")
			}
		}
	}()
	tsptw.ACO(data, seed)
	finish := time.Since(start)
	tsptw.Conclusion(seed, data, finish)
}
