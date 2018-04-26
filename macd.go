package indicators

import (
	"fmt"
)

//MACD port of tulipindicators MACD method
func MACD(inputs []float64, short, long, signal int) ([][]float64, error) {
	offset := (long - 1)

	if len(inputs) < offset {
		return nil, fmt.Errorf("not enough inputs")
	}

	/* macdLine := make([]float64, len(inputs)-offset)
	signalLine := make([]float64, len(inputs)-offset)
	hist := make([]float64, len(inputs)-offset) */

	if short < 1 || long < 2 || long < short || signal < 1 {
		return nil, fmt.Errorf("invalid params")
	}

	var (
		shortEMA, longEMA       []float64
		shortEMAErr, longEMAErr error
	)

	if shortEMA, shortEMAErr = EMA(inputs, short); shortEMAErr != nil {
		return nil, shortEMAErr
	}

	if longEMA, longEMAErr = EMA(inputs, long); longEMAErr != nil {
		return nil, longEMAErr
	}

	/**
	Okay, because this is confusing naming, I'll explain what's going on.
	The LONG EMA is the EMA list for the LONG period (26, usually).  However,
	the resulting EMA slice from the inputs for that period is going to be
	SHORTERthan the EMA slice for the SHORT period, because it has a bigger
	offset before values are calculated.  TL;DR: the EMA slices are named
	for the param they're the slices FOR, not the expected length of that
	slice.
	*/
	periodDiff := len(shortEMA) - len(longEMA)

	macdLine := make([]float64, len(longEMA))

	for i := 0; i < len(longEMA); i++ {
		macdLine[i] = shortEMA[i+periodDiff] - longEMA[i]
	}

	signalLine, signalErr := EMA(macdLine, signal)

	if signalErr != nil {
		return nil, signalErr
	}

	histogramDiff := len(macdLine) - len(signalLine)

	histogramLine := make([]float64, len(signalLine))

	for i := 0; i < len(signalLine); i++ {
		histogramLine[i] = macdLine[i+histogramDiff] - signalLine[i]
	}

	return [][]float64{
		macdLine[histogramDiff-1:],
		signalLine,
		histogramLine,
	}, nil
}
