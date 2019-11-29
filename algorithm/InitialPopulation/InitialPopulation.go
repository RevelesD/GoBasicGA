package InitialPopulation

import (
	"../../lib"
	"math/rand"
)

func EvaluateGen(gen []bool, dna []lib.Nucleotide) float64 {
	var allele float64
	if len(gen) != len(dna) {
		panic("length of DNA and Gen codification doesn't match")
	}
	for i := 0; i < len(gen); i++ {
		if gen[i] {
			if dna[i].Position == "+/-" {
				allele *= -1
			} else {
				allele += dna[i].Value
			}
		}
	}
	return allele
}

func fillGeneration(n []lib.Nucleotide, t *float64) []bool {
	genotype := make([]bool, len(n))
	for i := 0; i < len(n); i++ {
		if rand.Float64() > *t {
			genotype[i] = true
		} else {
			genotype[i] = false
		}
	}
	return genotype
}

func buildInitialGenome(n [][]lib.Nucleotide, t *float64) lib.Genome {
	vars := len(n)
	chromosomes := make([]lib.Chromosome, vars)
	for i := 0; i < vars; i++ { // nucleotides (aka 2^x...)
		var c lib.Chromosome
		c.Gen = fillGeneration(n[i], t)
		c.Allele = EvaluateGen(c.Gen, n[i])
		chromosomes[i] = c
	}
	genotype := lib.Genome{Chromosomes: chromosomes}
	return genotype
}
/**
	@Params(qty)
	@Params(v)
	@Params(t)
*/
func BuildInitialGenotype(qty *int, v [][]lib.Nucleotide, t *float64) []lib.Genome {
	genotype := make([]lib.Genome, *qty) // how many population
	for i := 0; i < *qty; i++ {
		// generate n elements
		genotype[i] = buildInitialGenome(v, t)
	}
	return genotype
}

/*	Procedure for populate with only a float64 tri-dimensional arrays
	func GenerateInitialPopulation(qty *int, v [][]Nucleotide, t *float64) [][][]float64 {
	// first  [] - elements in generation
	// second [] - variables x, y, z, ...
	// third  [] - genotype for every var
	population := make([][][]float64, *qty)
	for i := 0; i < *qty; i++{ // first
		tmp := make([][]float64, len(v))
		population[i] = tmp
		for j := 0; j < len(v); j++ { // second
			population[i][j] = fillGeneration(v[j], t) // third
		}
	}
	return population
}*/
