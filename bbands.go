package indicators

import (
	"fmt"
	"math"

	"github.com/technicalviking/sliceWindow"
)

//BBands port of bband indicator from tulipindicators c lib.
//@todo migrate this, and other ported go versions to my tulipindicators go lib.
func BBands(inputs []float64, period int, multiplier float64) ([][]float64, error) {

	if period < 1 {
		return nil, fmt.Errorf("period value must be greater than zero.  %d provided", period)
	}

	bbandOffset := int(period) - 1
	outputLength := len(inputs) - int(bbandOffset)

	if outputLength <= 0 {
		return nil, fmt.Errorf("Insufficient inputs length for indicator")
	}

	outputs := [][]float64{
		make([]float64, outputLength), //lower
		make([]float64, outputLength), //middle aka SMA
		make([]float64, outputLength), //upper
	}

	window := sliceWindow.New(period)

	for inputIndex, curInput := range inputs {
		window.PushBack(curInput)

		if inputIndex < bbandOffset {
			continue
		}

		//SMA
		curMean := window.Mean()

		scaledWindow := window.Map(func(input float64) float64 {
			result := input - curMean
			result *= result
			return result
		})

		//work out the mean of the squared differences, and take the square root of that mean.
		standardDev := scaledWindow.Mean()

		standardDev = math.Sqrt(standardDev)

		output := standardDev * multiplier

		//lower
		outputs[0][inputIndex-bbandOffset] = curMean - output
		//sma
		outputs[1][inputIndex-bbandOffset] = curMean
		//average
		outputs[2][inputIndex-bbandOffset] = curMean + output
	}

	return outputs, nil
}
