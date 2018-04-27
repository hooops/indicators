package indicators

import (
	"bitbonk/utils"
	"errors"
	"fmt"
)

//EMA Calculate Exponential Moving Average for a given set of inputs
func EMA(inputs []float64, period int) ([]float64, error) {
	if len(inputs) < 2 {
		return nil, errors.New("input length must be >= 2")
	}

	if period < 2 {
		return nil, errors.New("must have period of >= 2 to be useful")
	}

	if len(inputs) < period {
		return nil, errors.New("period must be <= provided input length")
	}

	offset := period - 1

	multiplier := 2.0 / float64(period+1)

	fmt.Println("multiplier", multiplier)

	window := utils.NewSliceWindow(period)

	outputs := make([]float64, len(inputs)-offset)

	for i, curInput := range inputs {
		window.PushBack(curInput)

		if i < offset {
			continue
		}

		//((Current price - SMA) × k) + SMA
		if i == offset {
			outputs[i-offset] = /* ((inputs[i] - window.Mean()) * multiplier) + */ window.Mean()
			continue
		}
		//((Current price - Previous EMA) × k) + Previous EMA
		outputs[i-offset] = ((inputs[i] - outputs[i-offset-1]) * multiplier) + outputs[i-offset-1]

	}

	return outputs, nil
}
