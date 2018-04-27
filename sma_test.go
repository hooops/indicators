package indicators

import "testing"

func TestSMA(t *testing.T) {
	inputs := []float64{45, 46, 43, 44, 42, 41, 40, 39, 41, 40, 38, 36}

	outputs, outputsErr := SMA(inputs, 10)

	if outputsErr != nil {
		t.Fatal(outputsErr)
	}

	expected := []float64{42.1, 41.4, 40.4}

	if len(outputs) != len(expected) {
		t.Fatalf("mismatch length: %d expected, %d found", len(expected), len(outputs))
	}

	for i, cur := range expected {
		if cur != outputs[i] {
			t.Errorf("mismatch at index %d:  expected %f, got %f", i, cur, outputs[i])
		}
	}

}
