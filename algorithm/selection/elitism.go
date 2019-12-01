package selection

import (
	"github.com/RevelesD/GoBasicGA/lib"
	"sort"
)

func Elitism(evals []lib.Evaluation, n *int) []int {
	picks := make([]int, *n)
	// sort by probability
	sort.Slice(evals, func(i, j int) bool {
		return evals[i].Codomain > evals[j].Codomain
		//return evals[i].Probability > evals[j].Probability
	})
	// select n top elements
	for i := 0; i < *n; i++ {
		 picks[i] = evals[i].Element
	}
	return picks
}
