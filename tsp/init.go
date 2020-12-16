package tsp

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var (
	PopSize           int     = 1000
	NumGen            int     = 500
	Eil51             string  = "data/tsp/eil51.txt"  // optimal = 426
	Eil76             string  = "data/tsp/eil76.txt"  // optimal = 538
	Eil101            string  = "data/tsp/eil101.txt" // optimal = 629
	Att48             string  = "data/tsp/att48.txt"  // optimal = none
	Chn31             string  = "data/tsp/chn31.txt"  // optimal = none
	Chn144            string  = "data/tsp/chn144.txt" // optimal = none
	MutationThreshold float64 = 0.2
	NumCities         int
	LocationData      [][]int
	DistanceMatrix    [][]float64
	Data              string
)

func init() {
	Data = Eil76
	getDataFromFile(Data)
}

func getDataFromFile(fileName string) {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fileString := string(fileBytes)
	slices := strings.Split(fileString, "\n")
	NumCities = len(slices)
	LocationData = make([][]int, NumCities)
	var x, y int
	for i, slice := range slices {
		LocationData[i] = make([]int, 2)
		el := strings.Split(slice, " ")
		x, _ = strconv.Atoi(el[1])
		y, _ = strconv.Atoi(el[2])
		LocationData[i][0] = x
		LocationData[i][1] = y
	}
	calculateDistanceMatrix()
}

func calculateDistanceMatrix() {
	DistanceMatrix = make([][]float64, NumCities+1)
	for i := 1; i <= NumCities; i++ {
		DistanceMatrix[i] = make([]float64, NumCities+1)
		for j := 1; j <= NumCities; j++ {
			if i == j {
				continue
			}
			dX := LocationData[i-1][0] - LocationData[j-1][0]
			dY := LocationData[i-1][1] - LocationData[j-1][1]
			DistanceMatrix[i][j] = distance(float64(dX), float64(dY))
		}
	}
}
