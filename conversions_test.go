package RefET_test

import (
	"math"
	"testing"

	"github.com/Longitude103/RefET"
)

func Test_CToF(t *testing.T) {
	want := 32.0
	got, err := RefET.CToF(0)
	if err != nil {
		t.Fatalf("Received and error during conversion: %s", err)
	}

	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func Test_FToC(t *testing.T) {
	want := 0.0
	got, err := RefET.FToC(32)
	if err != nil {
		t.Fatalf("Received and error during conversion: %s", err)
	}

	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func Test_makeFloat(t *testing.T) {
	want := 1.0
	got, err := RefET.MakeFloat(1)
	if err != nil {
		t.Fatalf("Received and error during conversion: %s", err)
	}

	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func Test_PaToKpa(t *testing.T) {
	const tolerance = 0.0001

	want := 1.0
	got, err := RefET.PaToKpa(1000)
	if err != nil {
		t.Fatalf("Recieved an error during conversion: %s", err)
	}

	if math.Abs(want-got) > tolerance {
		t.Errorf("want %f, got %f", want, got)
	}
}

func Test_LangToMJ(t *testing.T) {
	const tolerance = 0.0001

	want := 0.04184
	got, err := RefET.LangToMJ(1)
	if err != nil {
		t.Fatalf("Received an error during conversion: %s", err)
	}

	if math.Abs(want-got) > tolerance {
		t.Errorf("want %f, got %f", want, got)
	}
}

func Test_WattToMJ(t *testing.T) {
	const tolerance = 0.0001

	want := 0.0864
	got, err := RefET.WattToMJ(1)
	if err != nil {
		t.Fatalf("Received an error during conversion: %s", err)
	}

	if math.Abs(want-got) > tolerance {
		t.Errorf("want %f, got %f", want, got)
	}
}
