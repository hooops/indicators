package indicators

import "testing"

func TestAroon(t *testing.T) {
	highs := []float64{
		20, 19, 18, 17, 16, 15, 16, 17, 18, 19, 20,
		15, 16, 17, 18, 19, 20, 19, 18, 17, 16, 15,
	}

	lows := []float64{
		1, 2, 3, 4, 5, 6, 5, 4, 3, 2, 1,
		6, 5, 4, 3, 2, 1, 2, 3, 4, 5, 6,
	}

	output, err := Aroon(highs, lows, 14)

	if err != nil {
		t.Fatal(err)
	}

	if len(output) != 2 {
		t.Fatalf("wtf num outputs change? Expected %d, got %d", 2, len(output))
	}

	t.Error(output)
}

func TestAroonOsc(t *testing.T) {
	highs := []float64{
		20, 19, 18, 17, 16, 15, 16, 17, 18, 19, 20,
		15, 16, 17, 18, 19, 20, 19, 18, 17, 16, 15,
	}

	lows := []float64{
		1, 2, 3, 4, 5, 6, 5, 4, 3, 2, 1,
		6, 5, 4, 3, 2, 1, 2, 3, 4, 5, 6,
	}

	output, err := AroonOsc(highs, lows, 14)

	if err != nil {
		t.Fatal(err)
	}

	if len(output) != len(highs)-13 {
		t.Fatalf("wtf num outputs change? Expected %d, got %d", len(highs)-13, len(output))
	}

	t.Error(output)
}
