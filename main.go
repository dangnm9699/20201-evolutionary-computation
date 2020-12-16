package main

import (
	"ec/tsptw"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

const (
	DEFAULT = "tsptw/my_best"
)

func save(bspb tsptw.BSPB, payload string) {
	name := string(bspb)
	path := DEFAULT + "/" + name
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	_, err = file.WriteString(payload)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	var seed int64
	seed = int64(time.Now().Unix())
	//seed = int64(1608000695)
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
	tsptw.ACO(data)
	finish := time.Since(start)
	fmt.Println("Start at  :", start.Format(time.RFC3339))
	fmt.Println("Runtime at:", finish)
	fmt.Println("Seed =", seed)
	fmt.Println("======== RESULT ========")
	fmt.Printf("+ BestCost: %6.2f\n", tsptw.BestCost)
	fmt.Printf("+ BestTour: %v\n", tsptw.BestTour)
	fmt.Printf("+ BestTime: %6.2f\n", tsptw.BestTime+tsptw.DM[tsptw.BestTour[len(tsptw.BestTour)-1]][0])
	total := 0.0
	timer := 0.0
	ok := true
	fmt.Println("= CHECKING FEASIBILITY =")
	for i := 0; i < tsptw.NC; i++ {
		j := (i + 1) % tsptw.NC
		cur := tsptw.BestTour[i]
		nex := tsptw.BestTour[j]
		timer = math.Max(timer+tsptw.DM[cur][nex], tsptw.TW[nex][0])
		if timer > tsptw.TW[nex][1] {
			ok = false
		}
		total += tsptw.DM[cur][nex]
		fmt.Printf("+ From: %2d ; To: %2d ; Cost: %6.2f ; Time: %6.2f ; Time Window: %v\n", cur, nex, total, timer, tsptw.TW[nex])
	}
	fmt.Println("Path:", tsptw.BestTour)
	fmt.Printf("Cost: %.2f, Time: %.2f, IsFeasible: %v\n", total, timer, ok)
	fmt.Println("Duration:", finish)
	payload := fmt.Sprintf("%15d|%20v|%7.2f|%7.2f| %v\n", seed, finish, total, timer, tsptw.BestTour)
	save(data, payload)
}
