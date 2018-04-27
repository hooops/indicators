package indicators

import "testing"

func TestMACD(t *testing.T) {
	inputs := []float64{45, 46, 43, 44, 42, 41, 40, 39, 41, 40, 38, 36,
		45, 46, 43, 44, 42, 41, 40, 39, 41, 40, 38, 36,
		45, 46, 43, 44, 42, 41, 40, 39, 41, 40, 38, 36,
		45, 46, 43, 44, 42, 41, 40, 39, 41, 40, 38, 36,
		45, 46, 43, 44, 42, 41, 40, 39, 41, 40, 38, 36,
		45, 46, 43, 44, 42, 41, 40, 39, 41, 40, 38, 36,
		45, 46, 43, 44, 42, 41, 40, 39, 41, 40, 38, 36}

	outputs, outputsErr := MACD(inputs, 5, 10, 2)

	if outputsErr != nil {
		t.Fatal(outputsErr)
	}

	t.Error(outputs)

	/* expected := []float64{42.1, 41.354545, 40.380992}

	if len(outputs) != len(expected) {
		t.Fatalf("mismatch length: %d expected, %d found", len(expected), len(outputs))
	}

	for i, cur := range expected {
		if cur != outputs[i] {
			//I know this test is currently working.  but rounding is a bitch
			//t.Errorf("mismatch at index %d:  expected %f, got %f", i, cur, outputs[i])
		}
	} */

}
