package Codification

import (
	"fmt"
	"github.com/RevelesD/GoBasicGA/lib"
	"math"
)

func Codificacion(v *float64) []lib.Nucleotide {
	var negative bool
	if *v < 0 {
		negative = true
	}
	var integers = math.Floor(math.Abs(*v))
	var decimals = math.Abs(*v) - integers

	intChromo := codeIntegers(&integers)
	if negative {
		intChromo = append(intChromo, lib.Nucleotide{
			Position: "+/-",
			Value:    0,
		})
	}
	floatChromo := codeDecimals(&decimals)
	floatChromo = append(floatChromo, intChromo...)
	return floatChromo
}

func codeIntegers(i *float64) []lib.Nucleotide {
	exp := 0.0
	suma := 0.0
	var chromo []lib.Nucleotide

	for true {
		aux := math.Pow(2, exp)
		if suma + aux > *i {
			break
		}
		gen := lib.Nucleotide{
			Position: "2^" + fmt.Sprintf("%g", exp),
			Value:    aux,
		}
		chromo = append(chromo, gen)
		suma += aux
		exp++
	}

	offset := *i - suma
	gen := lib.Nucleotide{
		Position: fmt.Sprintf("%g", offset),
		Value:    offset,
	}

	for i, v := range chromo {
		if gen.Value < v.Value {
			chromo = insert(chromo, &i, &gen)
			break
		}
	}
	return chromo
}

func codeDecimals(d *float64) []lib.Nucleotide {
	exp := -1.0
	resto := *d
	var chromo []lib.Nucleotide

	for i := 0; i < 5; i++ {
		if resto == 0 {
			break
		}
		aux := math.Pow(2, exp)
		if aux > resto {
			exp--
			continue
		}
		gen := lib.Nucleotide{
			Position: "2^" + fmt.Sprintf("%g", exp),
			Value:    aux,
		}
		chromo = append(chromo, gen)
		resto -= aux
		exp--
	}

	if resto != 0 {
		chromo = append(chromo, lib.Nucleotide{
			Position: fmt.Sprintf("%g", resto),
			Value:    resto,
		})
	}
	return chromo
}

func insert(slice []lib.Nucleotide, index *int, element *lib.Nucleotide) []lib.Nucleotide {
	slice = append(slice, lib.Nucleotide{})  // Increase the slice size with one empty object
	copy(slice[*index + 1:], slice[*index:]) // Shift all the the elements from index to end
	slice[*index] = *element 				 // Insert the new element in the duplicate position
	return slice							 // Return the new slice header
}