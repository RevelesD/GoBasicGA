package crossover

import (
	in "github.com/RevelesD/GoBasicGA/algorithm/InitialPopulation"
	"github.com/RevelesD/GoBasicGA/lib"
	"math/rand"
)

/**
	Note
	this version of the AGT is going to differ slightly from
	the one studied on class, this is because it's not clear
	if elitism happens before or after the crossover. If it happens
	before the crossover and the number of elements passed
	by elitism is an odd number it's going to cause problems when
	we try to insert an even number of new elements to the next
	generation and because there are an odd number of elements
	passed by elitism the generation is going to either fall short
	or overflow on individuals.
	To prevent this there are two options,
		1) Make the elitism after the crossover and replace k random
		   individuals from the new generation with the elitism selection.
		2) Make the elitism before the crossover and Accept only even
		   numbers for elitism.
	In this implementation we are going with option 2.
	For the first implementation the next lecture is recommended
	https://www.ncbi.nlm.nih.gov/pmc/articles/PMC5597564/
 */


/**
* @param(c)    Current generation
* @param(eval) Current generation evaluation
* @param(ct)   Crossover threshold
*/
func Crossover(c []lib.Genome, eval []lib.Evaluation, ct *float64, n [][]lib.Nucleotide) (*lib.Genome, *lib.Genome) {
	a, b := selectPair(eval)
	x, y := crossoverPair(&c[a], &c[b], ct, n)
	return x, y
}
/**
* param(eval) Evaluations of the current generation
* returns()   A pair of indexes from different elements
*/
func selectPair(eval []lib.Evaluation) (int, int) {
	random := rand.Float64()
	var current = 1.0
	var firstIndex int
	var secondIndex int
	for _, value := range eval {
		if random < value.Cumulative && value.Cumulative < current {
			current = value.Cumulative
			firstIndex = value.Element
		}
	}
	current = 1.0
	for {
		random = rand.Float64()
		for _, value := range eval {
			if random < value.Cumulative && value.Cumulative < current {
				current = value.Cumulative
				secondIndex = value.Element
			}
		}
		if secondIndex != firstIndex {
			break
		}
		current = 1.0
	}
	return firstIndex, secondIndex
}
/**
* param(ct) Crossover threshold
* param(f)  The first element selected to be breed
* param(s)  The second element selected to be breed
* returns() The progeny result of the crossover
*/
func crossoverPair(f *lib.Genome, s *lib.Genome, ct *float64, n [][]lib.Nucleotide) (*lib.Genome, *lib.Genome) {
	if rand.Float64() <= *ct {
		return f, s
	}
	var half int
	chromosomesA := make([]lib.Chromosome, len(f.Chromosomes))
	chromosomesB := make([]lib.Chromosome, len(f.Chromosomes))
	a := lib.Genome{Chromosomes: chromosomesA}
	b := lib.Genome{Chromosomes: chromosomesB}
	for i, value := range f.Chromosomes {
		half = len(value.Gen) / 2
		chromosomesA[i].Gen = append(f.Chromosomes[i].Gen[:half], s.Chromosomes[i].Gen[half:]...)
		chromosomesA[i].Allele = in.EvaluateGen(chromosomesA[i].Gen, n[i])
		chromosomesB[i].Gen = append(s.Chromosomes[i].Gen[:half], f.Chromosomes[i].Gen[half:]...)
		chromosomesB[i].Allele = in.EvaluateGen(chromosomesB[i].Gen, n[i])
	}
	return &a, &b
}
