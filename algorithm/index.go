package algorithm

import (
	co "../algorithm/Codification"
	fi "../algorithm/FitnessFunction"
	in "../algorithm/InitialPopulation"
	c "../algorithm/crossover"
	m "../algorithm/mutation"
	se "../algorithm/selection"
	"../lib"
	"fmt"
	"math/rand"
	"time"
)

func RunBasicGA(f *int, e *int, gen *int, p *int, ti *float64, tc *float64, tm *float64) ([]lib.Genome, []lib.Evaluation) {
	/*
		Validate values for inputs
	*/
	// Parse the constants and functions for the selected option
	option, err := lib.ParseOption(f)
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil
	}
	// Validate the range of the threshold for population generation
	err = lib.ValidateThreshold(ti)
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil
	}
	// Validate the min population for generation
	err = lib.ValidatePopulation(p)
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil
	}
	// Validate the number of generations
	err = lib.ValidateGenerations(gen)
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil
	}
	// Validate the elitism against the population
	err = lib.ValidateElitism(e, p)
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil
	}
	// Validate the range of the threshold for crossovers
	err = lib.ValidateThreshold(tc)
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil
	}
	// Validate the range of the threshold for mutations
	err = lib.ValidateThreshold(tm)
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil
	}

	// start timer on algorithm process
	//start := time.Now()
	/*
		Mutation algorithm begins.
	*/
	var variables [][]lib.Nucleotide
	variables = append(variables, co.Codificacion(&option.X))
	variables = append(variables, co.Codificacion(&option.Y))
	// The seed for random numbers must be declared at top level
	rand.Seed(time.Now().UnixNano())
	// Declare empty slices for the current and next generations
	currentGen := make([]lib.Genome, *p)
	nextGen := make([]lib.Genome, *p)
	// Build of the initial population
	currentGen = in.BuildInitialGenotype(p, variables, ti)
	// One iteration for every desired generation
	for g := 0; g < *gen; g++ {
		// Fitness function
		evaluation := fi.EvalGenotype(option, currentGen)
		// Selection
		selections := se.Elitism(evaluation, e)
		se.MoveToNextGeneration(currentGen, nextGen, selections)
		/* Identify how many elements are missing from the
		   next generation after the elitism selection */
		remaining := *p - *e
		for i := 0; i < remaining; i = i + 2 {
			// Crossover
			a, b := c.Crossover(currentGen, evaluation, tc, variables)
			// Mutation
			m.Mutate(a, tm)
			m.Mutate(b, tm)
			// add elements to the next generations
			nextGen[len(selections) + i] = *a
			nextGen[len(selections) + i + 1] = *b
		}
		// restructure the generations
		for i := 0; i < len(currentGen); i++ {
			currentGen[i] = nextGen[i]
		}
	}
	// Evaluate the last generation
	evaluation := fi.EvalGenotype(option, currentGen)

	//elapsed := time.Since(start)
	//fmt.Printf("Tiempo en completar el algoritmo %s\n", elapsed)

	return currentGen, evaluation
}

