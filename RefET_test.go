package RefET_test

import (
	"math"
	"testing"

	"github.com/Longitude103/RefET"
)

func Test_RefET(t *testing.T) {
	t.Parallel()
	want := 0.0
	got := RefET.RefET(RefET.Input{Value: 5, Units: "F"}, RefET.Input{Value: 40, Units: "F"}, RefET.Input{Value: 5, Units: "F"}, RefET.Input{Value: 5, Units: "F"}, RefET.Input{Value: 5, Units: "F"}, RefET.Input{Value: 5, Units: "F"}, 5, 5, 5)

	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}

}

func Test_ConvertTemp(t *testing.T) {
	const tolerance = .00001

	i := RefET.Input{Value: 0, Units: "C"}
	i2 := RefET.Input{Value: 32, Units: "F"}
	i3 := RefET.Input{Value: 100, Units: "F"}
	testGots := []RefET.Input{i, i2, i3}

	want := []float64{0.0, 0.0, 37.77778}

	for j := 0; j < len(testGots); j++ {
		got, err := testGots[j].ConvertTemp()
		if err != nil {
			t.Fatal("Error in getting conversions for temp values")
		}

		if math.Abs(want[j]-got) > tolerance {
			t.Errorf("want %f, got %f", want[j], got)
		}
	}
}

func Test_ConvertEA(t *testing.T) {
	const tolerance = 0.0001

	i := RefET.Input{Value: 1200, Units: "PA"}
	i2 := RefET.Input{Value: 1.2, Units: "KPA"}
	i3 := RefET.Input{Value: 3.2, Units: "kpa"}
	i4 := RefET.Input{Value: 2853, Units: "pa"}
	testGots := []RefET.Input{i, i2, i3, i4}

	want := []float64{1.2, 1.2, 3.2, 2.853}

	for j := 0; j < len(testGots); j++ {
		got, err := testGots[j].ConvertEA()
		if err != nil {
			t.Fatal("Error in getting conversions for ea values")
		}

		if math.Abs(want[j]-got) > tolerance {
			t.Errorf("want %f, got %f", want[j], got)
		}
	}
}

func Test_CovertRA(t *testing.T) {
	const tolerance = 0.0001

	i := RefET.Input{Value: 1, Units: "Langley"}
	i2 := RefET.Input{Value: 1, Units: "w / m-2"}
	i3 := RefET.Input{Value: 1.0, Units: "w/m2"}
	i4 := RefET.Input{Value: 1.0, Units: "MJ/m2/d"}
	testGots := []RefET.Input{i, i2, i3, i4}

	want := []float64{0.04184, 0.0864, 0.0864, 1.0}

	for j := 0; j < len(testGots); j++ {
		got, err := testGots[j].ConvertRS()
		if err != nil {
			t.Fatal("Error in getting conversions for ea values")
		}

		if math.Abs(want[j]-got) > tolerance {
			t.Errorf("want %f, got %f", want[j], got)
		}
	}
}
