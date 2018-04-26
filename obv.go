package indicators

import (
	"fmt"
)

//OBV port of obv indicator from tulipindicators c lib.
//@todo migrate this, and other ported go versions to my tulipindicators go lib.
func OBV(close, volume []float64) ([]float64, error) {

	if len(close) != len(volume) {
		return nil, fmt.Errorf("inputs must be same length")
	}

	outputs := make([]float64, len(close))

	var sum float64

	for index, curClose := range close {
		if index == 0 {
			outputs[index] = sum
			continue
		}

		prev := close[index-1]

		if curClose > prev /*close[index] > prev*/ {
			sum += volume[index]
		} else if curClose < prev /*close[index] < prev*/ {
			sum -= volume[index]
		}

		outputs[index] = sum
	}

	return outputs, nil
}
