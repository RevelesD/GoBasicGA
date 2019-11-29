package FitnessFunction

import (
	"../../lib"
	"math"
)
// This function has hardcoded variables and operations
// because of the options given on the project
func EvalGenotype(o *lib.Opcion, g []lib.Genome) []lib.Evaluation {
	evaluations := make([]lib.Evaluation, len(g))

	var total float64
	for i, v := range g {
		// hardcoded indexes!
		codomain := o.Operation(v.Chromosomes[0].Allele, v.Chromosomes[1].Allele)
		total += math.Abs(codomain)
		evaluations[i].Codomain = codomain
		evaluations[i].Element = i
	}

	var cumulative float64
	for i, v := range evaluations {
		probability := math.Abs(v.Codomain/total)
		cumulative += probability
		evaluations[i].Probability = probability
		evaluations[i].Cumulative = cumulative
	}
	return evaluations
}
