package lib

type Nucleotide struct {
	Position string
	Value float64
}

type Genome struct {
	Chromosomes []Chromosome
}

type Chromosome struct {
	Gen    []bool
	Allele float64
}

type Evaluation struct {
	Codomain    float64  // Result of the function given the genome variables
	Probability float64
	Cumulative  float64
	Element     int      // Index of the genome in the genotype
}