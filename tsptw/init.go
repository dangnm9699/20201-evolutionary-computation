package tsptw

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

const (
	DEFAULT = "tsptw/my_best"
)

var (
	unvisited   []int
	constructed bool = false
	path        []int
)

func rand_feasible(seed int64) (res []int, float642 float64) {
	unvisited = new_unvisited()
	path = new_path()
	log.Println("=====> Start generate a feasible solution")
	start := time.Now()
	try(1, 0)
	log.Println("==+", time.Since(start))
	log.Printf("==+ Time:%7.2f; Path: %v\n", BestTime, BestTour)
	log.Println("=====> Done")
	return BestTour, BestTime
}

func try(k int, val float64) {
	fmt.Println(path)
	if k == NC {
		constructed = true
		BestTime = val
		copy(BestTour, path)
	}
	for i := 0; i < NC-1; i++ {
		curr := path[len(path)-1]
		next := unvisited[i]
		if constructed {
			return
		}
		if next != 0 && val+DM[curr][next] <= TW[next][1] {
			unvisited[i] = 0
			path = append(path, next)
			try(k+1, math.Max(val+DM[curr][next], TW[next][0]))
			unvisited[i] = next
			path = path[:len(path)-1]
		}
	}
}

func is_feasible(sol []int) (bool, float64) {
	timer := 0.0
	for i := 0; i < NC; i++ {
		j := (i + 1) % NC
		cur := sol[i]
		nex := sol[j]
		timer = math.Max(timer+DM[cur][nex], TW[nex][0])
		if timer > TW[nex][1] {
			return false, 0
		}
	}
	return true, timer
}

// Check if Best solution is feasible
// and print information
func Conclusion(seed int64, data BSPB, dur time.Duration) {
	fmt.Println("======== RESULT ========")
	fmt.Printf("+ BestTour: %v\n", BestTour)
	fmt.Printf("+ BestTime: %6.2f\n", BestTime+DM[BestTour[len(BestTour)-1]][0])
	//
	total := 0.0
	timer := 0.0
	ok := true
	fmt.Println("= CHECKING FEASIBILITY =")
	for i := 0; i < NC; i++ {
		ok = true
		j := (i + 1) % NC
		cur := BestTour[i]
		nex := BestTour[j]
		timer = math.Max(timer+DM[cur][nex], TW[nex][0])
		if timer > TW[nex][1] {
			ok = false
		}
		total += DM[cur][nex]
		if !ok {
			fmt.Printf("+ From: %2d ; To: %2d ; Cost: %6.2f ; Time Window: %3v | %v\n", cur, nex, timer, TW[nex], ok)
		} else {
			fmt.Printf("+ From: %2d ; To: %2d ; Cost: %6.2f ; Time Window: %3v\n", cur, nex, timer, TW[nex])
		}
	}
	fmt.Println("Path:", BestTour)
	fmt.Printf("Cost: %.2f, IsFeasible: %v\n", timer, ok)
	payload := fmt.Sprintf("%15d|%20v|%7.2f| %v\n", seed, dur, timer, BestTour)
	save(data, payload)
}

func save(bspb BSPB, payload string) {
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

func new_ant() Ant {
	return Ant{
		Position:  0,
		Path:      new_path(),
		Cost:      0.0,
		Time:      0.0,
		Unvisited: new_unvisited(),
		Delta:     new_delta(),
	}
}

func new_delta() [][]float64 {
	ma := make([][]float64, NC)
	for i := 0; i < NC; i++ {
		ma[i] = make([]float64, NC)
		for j := 0; j < NC; j++ {
			ma[i][j] = 1.0
		}
	}
	return ma
}

func new_path() []int {
	return []int{0}
}

func new_unvisited() []int {
	u := make([]int, 0)
	for i := 1; i < NC; i++ {
		u = append(u, i)
	}
	return u
}

func remove_index(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
