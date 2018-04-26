package indicators

/* func normalize(input []float64) []float64 {
	result := make([]float64, len(input))

	max := big.NewFloat(-1.0)
	min := big.NewFloat(9999999999.0)

	for _, curInput := range input {
		if curInput.Cmp(max) == 1 {
			max.Copy(curInput)
		}

		if curInput.Cmp(min) == -1 {
			min.Copy(curInput)
		}
	}

	dataRange := big.NewFloat(0.0).Sub(max, min)

	for index, curInput := range input {
		temp := big.NewFloat(0.0).Sub(curInput, min)

		result[index] = temp.Quo(temp, dataRange)
	}

	return result
} */
