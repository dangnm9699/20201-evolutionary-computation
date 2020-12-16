package tsp

import (
	"math/rand"
	"sort"
	"sync"
)

func Run() (bestFitness float64, bestSolution []int) {
	Pop := genPop()
	sort.Slice(Pop, func(i, j int) bool {
		return fitness(Pop[i]) < fitness(Pop[j])
	})
	for curGen := 0; curGen < NumGen; curGen++ {
		//
		Parents := make([][]int, len(Pop))
		copy(Parents, Pop)
		//
		fitnessArray := make([]float64, 0)
		fitnessTotal := 0.0
		for _, ind := range Pop {
			fitnessTotal += fitness(ind)
			fitnessArray = append(fitnessArray, fitnessTotal)
		}
		for i := 0; i < PopSize; i++ {
			fitnessArray[i] /= fitnessTotal
		}
		//
		Children := make([][]int, PopSize)
		var wg sync.WaitGroup
		for i := 0; i < PopSize; i++ {
			//
			index := i
			wg.Add(1)
			go func(index int) {
				p1 := rouletteWheelSelection(fitnessArray, Pop)
				p2 := rouletteWheelSelection(fitnessArray, Pop)
				//
				Children[index] = improvedHeuristicCrossover(p1, p2)
				mutationProb := rand.Float64()
				swapMutation(Children[index], mutationProb)
				wg.Done()
				//c := heuristicCrossover(p1, p2)
			}(index)
		}
		wg.Wait()
		Pop = append(Parents, Children...)
		sort.Slice(Pop, func(i, j int) bool {
			return fitness(Pop[i]) < fitness(Pop[j])
		})
		Pop = Pop[:PopSize]
	}
	return fitness(Pop[0]), Pop[0]
}

func rouletteWheelSelection(fitnessPosProb []float64, pop [][]int) (p []int) {
	randomNumber := rand.Float64()
	it := 0
	for fitnessPosProb[it] < randomNumber {
		it += 1
	}
	p = make([]int, len(pop[it]))
	copy(p, pop[it])
	return
}
