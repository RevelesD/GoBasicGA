package selection

import "../../lib"

/**
* @Param(c) current generation
* @Param(n) next generation
* @Param(s) selection by elitism
*/
func MoveToNextGeneration(c []lib.Genome, n []lib.Genome, s []int)  {
	for index, value := range s {
		n[index] = c[value]
	}
}
