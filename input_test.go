package RefET

import (
	"math"
	"testing"
)

func Test_ConvertTemp(t *testing.T) {
	i := Input{Value: 0, Units: "C"}
	i2 := Input{Value: 32, Units: "F"}
	i3 := Input{Value: 100, Units: "F"}
	testGots := []Input{i, i2, i3}

	want := []float64{0.0, 0.0, 37.77778}

	for j := 0; j < len(testGots); j++ {
		got, err := testGots[j].convertTemp()
		if err != nil {
			t.Fatal("Error in getting conversions for temp values")
		}

		if math.Abs(want[j]-got) > tolerance {
			t.Errorf("want %f, got %f", want[j], got)
		}
	}
}

func Test_CovertRA(t *testing.T) {
	i := Input{Value: 1, Units: "Langley"}
	i2 := Input{Value: 1, Units: "w / m-2"}
	i3 := Input{Value: 1.0, Units: "w/m2"}
	i4 := Input{Value: 1.0, Units: "MJ/m2/d"}
	testGots := []Input{i, i2, i3, i4}

	want := []float64{0.04184, 0.0864, 0.0864, 1.0}

	for j := 0; j < len(testGots); j++ {
		got, err := testGots[j].convertRS()
		if err != nil {
			t.Fatal("Error in getting conversions for ea values")
		}

		if math.Abs(want[j]-got) > tolerance {
			t.Errorf("want %f, got %f", want[j], got)
		}
	}
}
