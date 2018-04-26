package indicators

import (
	"fmt"
	
	"testing"
)

func TestBBands(t *testing.T) {
	source := []float64{
		big.NewFloat(4.0),
		big.NewFloat(5.0),
		big.NewFloat(6.0),
		big.NewFloat(4.0),
		big.NewFloat(5.0),
		big.NewFloat(6.0),
	}

	outputs, err := BBands(3, 1, source)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v \n", outputs)

	t.Errorf("bye")
}
