package RefET

import (
	"math"
	"testing"
)

func Test_convertEA(t *testing.T) {
	i := EaInput{Input: Input{Value: 1200, Units: "PA"}, Method: 1}
	i2 := EaInput{Input: Input{Value: 1.2, Units: "KPA"}, Method: 1}
	i3 := EaInput{Input: Input{Value: 3.2, Units: "kpa"}, Method: 1}
	i4 := EaInput{Input: Input{Value: 2853, Units: "pa"}, Method: 1}
	testGots := []EaInput{i, i2, i3, i4}

	want := []float64{1.2, 1.2, 3.2, 2.853}

	for j := 0; j < len(testGots); j++ {
		got, err := testGots[j].convertEA()
		if err != nil {
			t.Fatal("Error in getting conversions for ea values")
		}

		if math.Abs(want[j]-got) > tolerance {
			t.Errorf("want %f, got %f", want[j], got)
		}
	}
}

func Test_convertFromTdew(t *testing.T) {
	i := EaInput{Input: Input{Value: 10, Units: "C"}, Method: 2}
	i2 := EaInput{Input: Input{Value: 65.0, Units: "F"}, Method: 2}

	testGots := []EaInput{i, i2}

	want := []float64{1.228, 2.107}
	for j := 0; j < len(testGots); j++ {
		got, err := testGots[j].convertFromTdew()
		if err != nil {
			t.Fatal("Error in getting conversions for ea values")
		}

		if math.Abs(want[j]-got) > tolerance {
			t.Errorf("want %f, got %f", want[j], got)
		}
	}
}

func Test_convertMinMaxRH(t *testing.T) {
	i := EaInput{Input: Input{}, Tmax: Input{Value: 32.0, Units: "C"}, Tmin: Input{Value: 25, Units: "C"}, RHmax: 75, RHmin: 45, Method: 5}
	i2 := EaInput{Input: Input{}, Tmax: Input{Value: 29.0, Units: "C"}, Tmin: Input{Value: 20, Units: "C"}, RHmax: 85, RHmin: 65, Method: 5}

	testGots := []EaInput{i, i2}
	want := []float64{2.258, 2.296}

	for j := 0; j < len(testGots); j++ {
		got, err := testGots[j].convertMinMaxRH()
		if err != nil {
			t.Fatal("Error in getting conversions for ea values")
		}

		if math.Abs(want[j]-got) > tolerance {
			t.Errorf("want %f, got %f", want[j], got)
		}
	}
}

func Test_convertRHmax(t *testing.T) {
	i := EaInput{Input: Input{}, Tmax: Input{}, Tmin: Input{Value: 25, Units: "C"}, RHmax: 75, Method: 6}
	i2 := EaInput{Input: Input{}, Tmax: Input{}, Tmin: Input{Value: 20, Units: "C"}, RHmax: 85, Method: 6}

	testGots := []EaInput{i, i2}
	want := []float64{2.376, 1.988}

	for j := 0; j < len(testGots); j++ {
		got, err := testGots[j].convertRHmax()
		if err != nil {
			t.Fatal("Error in getting conversions for ea values")
		}

		if math.Abs(want[j]-got) > tolerance {
			t.Errorf("want %f, got %f", want[j], got)
		}
	}
}

func Test_convertRHmin(t *testing.T) {
	i := EaInput{Input: Input{}, Tmax: Input{Value: 32.0, Units: "C"}, Tmin: Input{}, RHmin: 45, Method: 7}
	i2 := EaInput{Input: Input{}, Tmax: Input{Value: 29.0, Units: "C"}, Tmin: Input{}, RHmin: 65, Method: 7}

	testGots := []EaInput{i, i2}
	want := []float64{2.140, 2.604}

	for j := 0; j < len(testGots); j++ {
		got, err := testGots[j].convertRHmin()
		if err != nil {
			t.Fatal("Error in getting conversions for ea values")
		}

		if math.Abs(want[j]-got) > tolerance {
			t.Errorf("want %f, got %f", want[j], got)
		}
	}
}

func Test_Ea(t *testing.T) {
	i := EaInput{Input: Input{Value: 1.2, Units: "KPA"}, Method: 1}
	i2 := EaInput{Input: Input{Value: 10, Units: "C"}, Method: 2}
	i3 := EaInput{Input: Input{}, Tmax: Input{Value: 32.0, Units: "C"}, Tmin: Input{Value: 25, Units: "C"}, RHmax: 75, RHmin: 45, Method: 5}
	i4 := EaInput{Input: Input{}, Tmax: Input{}, Tmin: Input{Value: 25, Units: "C"}, RHmax: 75, Method: 6}
	i5 := EaInput{Input: Input{}, Tmax: Input{Value: 32.0, Units: "C"}, Tmin: Input{}, RHmin: 45, Method: 7}

	testGots := []EaInput{i, i2, i3, i4, i5}
	want := []float64{1.2, 1.228, 2.258, 2.376, 2.140}

	for j := 0; j < len(testGots); j++ {
		got, err := testGots[j].Ea()
		if err != nil {
			t.Fatal("Error in getting conversions for ea values")
		}

		if math.Abs(want[j]-got) > tolerance {
			t.Errorf("want %f, got %f", want[j], got)
		}
	}
}
