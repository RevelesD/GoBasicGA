package lib

import (
	"github.com/pkg/errors"
)

func ValidateThreshold(thresholdFlag *float64) error {
	if *thresholdFlag > 1 || *thresholdFlag < 0 {
		return errors.New("Threshold cannot be lower than 0 or bigger than 1")
	}
	return nil
}

func ValidateElitism(elitism *int, population *int) error {
	if *elitism >= *population {
		return errors.New("If the number of elements selected by elitism" +
			" is equal or bigger than the population number, no evolution can be made")
	}
	if *elitism % 2 != 0 {
		return errors.New("Elitism must come in pairs, for more information see ...")
	}
	return nil
}

func ParseOption(optionFlag *int) (*Opcion, error) {
	option := Opcion{}
	switch *optionFlag {
	case 1:
		option.X = -720
		option.Y = -720
		option.Operation = OptionOne
		break
	case 2:
		option.X = -720
		option.Y = -720
		option.Operation = OptionTwo
		break
	case 3:
		option.X = -360
		option.Y = -360
		option.Operation = OptionThree
		break
	case 4:
		option.X = -360
		option.Y = -360
		option.Operation = OptionFour
		break
	default:
		return &option, errors.New("Invalid function option")
	}
	return &option, nil
}

func ValidatePopulation(f *int) error {
	if *f < 6 {
		return errors.New("Population must be bigger at least 6 elements")
	}
	if *f % 2 != 0 {
		return errors.New("Population size must be an even number")
	}
	return nil
}

func ValidateGenerations(g *int) error {
	if *g < 1 {
		return errors.New("Cannot have less than 1 generation")
	}
	return nil
}