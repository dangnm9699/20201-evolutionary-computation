package tsp

import (
	"ec/utils"
	"fmt"
	"math/rand"
	"os"
	"sort"
)

func improvedHeuristicCrossover(p1, p2 []int) (c []int) {
	if len(p1) != len(p2) {
		panic("[PAN] Not the same length")
	}
	randCity := rand.Intn(NumCities) + 1
	indexInP1, err := utils.GetIndex(p1, randCity)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	indexInP2, err := utils.GetIndex(p2, randCity)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	c = make([]int, 0)
	a := make([]int, len(p1))
	b := make([]int, len(p2))
	copy(a, p1)
	copy(b, p2)
	prev := 0
	c = append(c, randCity)
	rightCircumvolve(p1, indexInP1, 0)
	rightCircumvolve(p2, indexInP2, 0)
	for i := 1; i < NumCities; i++ {
		prev += 1
		if isNeighborhood(c[prev-1], p1[prev], a) &&
			isNeighborhood(c[prev-1], p2[prev], b) {
			if DistanceMatrix[c[prev-1]][p1[i]] < DistanceMatrix[c[prev-1]][p2[i]] {
				c = append(c, p1[i])
				index, err := utils.GetIndex(p2, p1[i])
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				rightCircumvolve(p2, index, i)
			} else {
				c = append(c, p2[i])
				index, err := utils.GetIndex(p1, p2[i])
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				rightCircumvolve(p1, index, i)
			}
		} else {
			others := getCityNotIn(c)
			sort.Slice(others, func(i, j int) bool {
				return DistanceMatrix[c[prev-1]][others[i]] < DistanceMatrix[c[prev-1]][others[j]]
			})
			c = append(c, others[0])
			indexA, err := utils.GetIndex(p1, others[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
			indexB, err := utils.GetIndex(p2, others[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
			rightCircumvolve(p1, indexA, i)
			rightCircumvolve(p2, indexB, i)
		}
	}
	return
}

func isNeighborhood(cityCenter, cityB int, p []int) bool {
	indexCenter, err := utils.GetIndex(p, cityCenter)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	indexB, err := utils.GetIndex(p, cityB)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	if !isNear(indexCenter, indexB) {
		return false
	}
	return true
}

func isValid(i int) bool {
	if i < 0 || i >= NumCities {
		return false
	}
	return true
}

func isNear(center, b int) bool {
	if !isValid(center) || !isValid(b) {
		panic("[PAN] Invalid index")
	}
	if center == 0 {
		if b != 1 {
			return false
		}
	} else {
		if center == NumCities-1 {
			if b != NumCities-2 {
				return false
			}
		} else {
			if b < center-1 || b > center+1 {
				return false
			}
		}
	}
	return true
}

func getCityNotIn(c []int) []int {
	li := make([]int, NumCities+1)
	for i := 0; i < len(c); i++ {
		li[c[i]] = 1
	}
	res := make([]int, 0)
	for i := 1; i <= NumCities; i++ {
		if li[i] != 1 {
			res = append(res, i)
		}
	}
	return res
}
