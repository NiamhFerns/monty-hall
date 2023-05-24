package main

import (
	"fmt"
	"math/rand"
)

func run() bool {
	vals := [3]bool{false, false, false}
	choice := rand.Intn(3)
	vals[choice] = true

	choice = rand.Intn(3)
	// We always switch and can assume an incorrect one is taken so...
	// If you're correct initially you'll always switch to the wrong one.
	if vals[choice] {
		return false
	}

	// If you're wrong initially you'll always switch to the correct one.
	return !vals[choice]
}

const ITERATIONS = 1000000

func main() {
	var correct int
	var incorrect int
	for i := 0; i < ITERATIONS; i++ {
		if run() {
			correct++
		} else {
			incorrect++
		}
	}

	fmt.Printf(
		"You got %d correct and %d incorrect meaning you're correct %f%% of the time.",
		correct,
		incorrect,
		float64(correct)/ITERATIONS*100,
	)
}
