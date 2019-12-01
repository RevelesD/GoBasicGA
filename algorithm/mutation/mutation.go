package mutation

import (
	"github.com/RevelesD/GoBasicGA/lib"
	"math/rand"
)

/**
* param(e)  New element to be added to the next generation
* param(mt) Mutation threshold
 */
func Mutate(e *lib.Genome, mt *float64) {
	for i := 0; i < len(e.Chromosomes); i++ {
		for j := 0; j < len(e.Chromosomes[i].Gen); j++ {
			if rand.Float64() > *mt {
				e.Chromosomes[i].Gen[j] = !e.Chromosomes[i].Gen[j]
			}
		}
	}
}
