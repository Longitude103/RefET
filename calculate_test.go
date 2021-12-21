package RefET

import (
	"math"
	"testing"
)

const tolerance = 0.01

// These tests us the data from Appendix C of the manual to ensure that the test data that is used is comparable to the authors data.

func Test_atmosPressure(t *testing.T) {
	elev := 1462.4

	want := 85.1666
	got := atmosPressure(elev)

	if math.Abs(want-got) > tolerance {
		t.Errorf("want %f, got %f", want, got)
	}
}

func Test_psyConst(t *testing.T) {
	want := 0.056635
	got := psyConst(85.1666)

	if math.Abs(want-got) > tolerance {
		t.Errorf("want %f, got %f", want, got)
	}
}

func Test_meanT(t *testing.T) {
	want := 21.65
	got := meanT(32.4, 10.9)

	if math.Abs(want-got) > tolerance {
		t.Errorf("want %f, got %f", want, got)
	}
}

func Test_eaSlope(t *testing.T) {
	want := 0.1585
	got := esSlope(21.7)

	if math.Abs(want-got) > tolerance {
		t.Errorf("want %f, got %f", want, got)
	}
}

func Test_satVP(t *testing.T) {
	want := 3.09
	got := satVP(32.4, 10.9)

	if math.Abs(want-got) > tolerance {
		t.Errorf("want %f, got %f", want, got)
	}
}

func Test_eo(t *testing.T) {
	want := 4.86 // test data was close at 4.88
	got := eo(32.4)

	if math.Abs(want-got) > tolerance {
		t.Errorf("want %f, got %f", want, got)
	}

	want = 1.30401 // test data was close at 1.31
	got = eo(10.9)

	if math.Abs(want-got) > tolerance {
		t.Errorf("want %f, got %f", want, got)
	}
}

func Test_calcWS(t *testing.T) {
	want := 1.79
	got := calcWS(1.94, 3.0)

	if math.Abs(want-got) > tolerance {
		t.Errorf("want %f, got %f", want, got)
	}
}

func Test_inverseRelDistFactor(t *testing.T) {
	doy, _ := DateToDOY("07-01-2000")

	want := 0.967
	got := inverseRelDistFactor(doy)

	if math.Abs(want-got) > tolerance {
		t.Errorf("want %f, got %f", want, got)
	}
}

func Test_solarDeclin(t *testing.T) {
	doy, _ := DateToDOY("07-01-2000")

	want := 0.4017
	got := solarDeclin(doy)

	if math.Abs(want-got) > tolerance {
		t.Errorf("want %f, got %f", want, got)
	}
}

func Test_sunsetHourAngle(t *testing.T) {
	lat, _ := DegreesToRad(40.41)
	doy, _ := DateToDOY("07-01-2000")

	delta := solarDeclin(doy)

	want := 1.941
	got := sunsetHourAngle(lat, delta)
	if math.Abs(want-got) > tolerance {
		t.Errorf("want %f, got %f", want, got)
	}
}

func Test_calcRA(t *testing.T) {
	doy, _ := DateToDOY("07-01-2000")
	lat, _ := DegreesToRad(40.41)

	want := 41.63
	got := calcRA(lat, doy)

	if math.Abs(want-got) > tolerance {
		t.Errorf("want %f, got %f", want, got)
	}
}

func Test_calcRSO(t *testing.T) {
	ra := 41.63
	elev := 1462.4

	want := 32.44 // close to test data of 32.43
	got := calcRSO(ra, elev)
	if math.Abs(want-got) > tolerance {
		t.Errorf("want %f, got %f", want, got)
	}
}

func Test_calcFCD(t *testing.T) {
	// this isn't in the manual, but calculated by hand using the test data
	ra := 41.63
	elev := 1462.4
	rs := 22.4
	rso := calcRSO(ra, elev)

	want := 0.5822
	got := calcFCD(rs, rso)
	if math.Abs(want-got) > tolerance {
		t.Errorf("want %f, got %f", want, got)
	}
}

func Test_calcRNL(t *testing.T) {
	ra := 41.63
	elev := 1462.4
	rs := 22.4
	rso := calcRSO(ra, elev)
	fcd := calcFCD(rs, rso)

	want := 3.96
	got := calcRNL(fcd, 1.27, 32.4, 10.9)
	if math.Abs(want-got) > tolerance {
		t.Errorf("want %f, got %f", want, got)
	}
}

func Test_calcRNS(t *testing.T) {
	rs := 22.4

	want := 17.25
	got := calcRNS(rs)
	if math.Abs(want-got) > tolerance {
		t.Errorf("want %f, got %f", want, got)
	}
}

func Test_calcRN(t *testing.T) {
	ra := 41.63
	elev := 1462.4
	rs := 22.4
	rso := calcRSO(ra, elev)
	fcd := calcFCD(rs, rso)
	rnl := calcRNL(fcd, 1.27, 32.4, 10.9)

	want := 13.29 // very close to test data of 13.31
	got := calcRN(17.25, rnl)
	if math.Abs(want-got) > tolerance {
		t.Errorf("want %f, got %f", want, got)
	}
}
