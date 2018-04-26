//THIS INDICATOR IS NOT PART OF TULIP
//https://www.investopedia.com/terms/a/accumulationdistribution.asp

package indicators

import (
	"fmt"
)

//CMF  AKA Chaikin Money Flow.  Simplified from
//equations found on https://www.investopedia.com/terms/a/accumulationdistribution.asp
//http://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:accumulation_distribution_line
//https://www.tradingview.com/wiki/Chaikin_Money_Flow_(CMF)
//tulipindicators does not include this. (MFI) is a different indicator.
func CMF(highs, lows, volumes []float64, period int) ([]float64, error) {
	offset := period - 1
	if len(highs) < period {
		return nil, fmt.Errorf("insufficient inputs")
	}

	if (len(highs) | len(lows) | len(volumes)) != len(lows) {
		return nil, fmt.Errorf("all inputs must be of equal length")
	}

	var (
		rangeSum      float64
		rangeDiff     float64
		multiplier    float64
		flowVolume    float64
		sumFlowVolume float64
		sumVolume     float64
	)

	outputs := make([]float64, len(highs)-offset)

	flowVolumes := make([]float64, len(highs))

	for i := 0; i < len(highs); i++ {
		/*
			Money Flow Multiplier:
			((Close - Low) - (High - Close)) / (High - Low)
			Simplifies to:
			(Low + High) / (Low - High)
		*/

		rangeSum = lows[i] + highs[i]
		rangeDiff = lows[i] - highs[i]

		multiplier = rangeSum / rangeDiff

		/*
			Money Flow volume
			Money Flow Multiplier x Volume
		*/
		flowVolume = multiplier * volumes[i]

		//store the flow volume for later.
		flowVolumes[i] = flowVolume

		//gotta get the sums goin.
		sumVolume += volumes[i]
		sumFlowVolume += flowVolume

		//if we haven't hit the offset yet, keep goin.
		if i < offset {
			continue
		}

		//if we're higher than the offset, remove old volumes from sums
		if i > offset {
			sumVolume -= volumes[i-offset]
			sumFlowVolume -= flowVolumes[i-offset]
		}

		/*
			CMF
			period sum of Money Flow Volume / period sum of volume = period CMF
		*/

		outputs[i-offset] = sumFlowVolume / sumVolume

	}

	return outputs, nil
}
