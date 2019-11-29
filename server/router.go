package server

import (
	al "../algorithm"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"sort"

	//"sort"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/startga", func(c *gin.Context) {
		type Entries struct {
			Func        int     `json:"func"`
			Elitism     int     `json:"elitism"`
			Generations int     `json:"generations"`
			Population  int     `json:"population"`
			Umbral      float64 `json:"umbral"`
			Crossover   float64 `json:"crossover"`
			Mutation    float64 `json:"mutation"`
		}
		var j Entries
		c.BindJSON(&j)

		gen, evals := al.RunBasicGA(&j.Func, &j.Elitism, &j.Generations, &j.Population, &j.Umbral, &j.Crossover, &j.Mutation)
		type Individual struct {
			X float64
			Y float64
			Value float64
		}
		if len(gen) != len(evals) {
			panic("Lenghts of las generations doesn't match")
		}
		generation := make([]Individual, len(gen))
		for key, value := range evals {
			generation[key].Value = value.Codomain
			generation[key].X = gen[value.Element].Chromosomes[0].Allele
			generation[key].Y = gen[value.Element].Chromosomes[1].Allele
		}
		sort.Slice(generation, func(i, j int) bool {
			return generation[i].Value > generation[j].Value
		})
		c.JSON(200, generation)
	})
	return r
}
