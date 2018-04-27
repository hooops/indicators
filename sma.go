package indicators

import (
	"bitbonk/utils"
	"errors"
	"fmt"
)

//SMA Simple Moving Average
func SMA(inputs []float64, period int) ([]float64, error) {

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

	window := utils.NewSliceWindow(period)

	outputs := make([]float64, len(inputs)-offset)

	fmt.Println("output len", len(inputs), len(outputs))

	for i, curInput := range inputs {
		window.PushBack(curInput)

		if i < offset {
			continue
		}

		fmt.Println("i", i)

		outputs[i-offset] = window.Mean()
	}

	return outputs, nil
}
