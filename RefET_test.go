package RefET

import (
	"math"
	"testing"
)

type dailyWeather struct {
	tmax float64
	tmin float64
	ea   float64
	rs   float64
	ws   float64
	date string
}

type stationData struct {
	wz  float64
	z   float64
	lat float64
}

func Test_calculateRefET(t *testing.T) {
	const etTolerance = 0.025

	// daily weather data from Appendix C Table C-3
	jul1 := dailyWeather{32.4, 10.9, 1.27, 22.4, 1.94, "07-01-2000"}
	jul2 := dailyWeather{33.6, 12.2, 1.19, 26.8, 2.14, "07-02-2000"}
	jul3 := dailyWeather{32.6, 14.8, 1.40, 23.3, 2.06, "07-03-2000"}
	jul4 := dailyWeather{33.8, 11.8, 1.18, 29.0, 1.97, "07-04-2000"}
	jul5 := dailyWeather{32.7, 15.9, 1.59, 27.9, 2.98, "07-05-2000"}
	dailyData := []dailyWeather{jul1, jul2, jul3, jul4, jul5}

	greeley := stationData{wz: 3.0, z: 1462.4, lat: 40.41}

	wantShort := []float64{5.71, 6.71, 5.98, 6.86, 7.03}
	wantTall := []float64{7.34, 8.68, 7.65, 8.73, 9.07}

	for i, d := range dailyData {
		doy, _ := DateToDOY(d.date)
		lat, _ := DegreesToRad(greeley.lat)

		etShort, etTall := calculateRefET(d.tmax, d.tmin, d.ea, d.rs, d.ws, greeley.wz, greeley.z, lat, doy)

		if math.Abs(wantShort[i]-etShort) > etTolerance {
			t.Errorf("want %f, got %f", wantShort[i], etShort)
		}

		if math.Abs(wantTall[i]-etTall) > etTolerance {
			t.Errorf("want %f, got %f", wantTall[i], etTall)
		}
	}
}

type dailyClimate struct {
	Tmax Input
	Tmin Input
	Ea   EaInput
	Rs   Input
	Ws   Input
	Date Input
}

type stationDt struct {
	wz  Input
	z   Input
	lat Input
}

func Test_RefET(t *testing.T) {
	const etTolerance = 0.025
	greeley := stationDt{wz: Input{3.0, "m"}, z: Input{1462.4, "m"}, lat: Input{40.41, "degrees"}}

	jul1 := dailyClimate{Tmax: Input{32.4, "C"}, Tmin: Input{10.9, "C"}, Ea: EaInput{Input: Input{1.27, "KPA"}, Method: 1}, Rs: Input{22.4, "MJ m-2 d-1"}, Ws: Input{1.94, "m s-1"}, Date: Input{"07-01-2000", "date"}}
	jul2 := dailyClimate{Tmax: Input{33.6, "C"}, Tmin: Input{12.2, "C"}, Ea: EaInput{Input: Input{1.19, "KPA"}, Method: 1}, Rs: Input{26.8, "MJ m-2 d-1"}, Ws: Input{2.14, "m s-1"}, Date: Input{"07-02-2000", "Date"}}
	jul3 := dailyClimate{Tmax: Input{32.6, "C"}, Tmin: Input{14.8, "C"}, Ea: EaInput{Input: Input{1.40, "KPA"}, Method: 1}, Rs: Input{23.3, "MJ m-2 d-1"}, Ws: Input{2.06, "m s-1"}, Date: Input{"07-03-2000", "date"}}
	jul4 := dailyClimate{Tmax: Input{33.8, "C"}, Tmin: Input{11.8, "C"}, Ea: EaInput{Input: Input{1.18, "KPA"}, Method: 1}, Rs: Input{29.0, "MJ m-2 d-1"}, Ws: Input{1.97, "m s-1"}, Date: Input{"07-04-2000", "date"}}
	jul5 := dailyClimate{Tmax: Input{32.7, "C"}, Tmin: Input{15.9, "C"}, Ea: EaInput{Input: Input{1.59, "KPA"}, Method: 1}, Rs: Input{27.9, "MJ m-2 d-1"}, Ws: Input{2.98, "m s-1"}, Date: Input{"07-05-2000", "date"}}
	dailyData := []dailyClimate{jul1, jul2, jul3, jul4, jul5}

	wantShort := []float64{5.71, 6.71, 5.98, 6.86, 7.03}
	wantTall := []float64{7.34, 8.68, 7.65, 8.73, 9.07}

	for i, d := range dailyData {
		etShort, etTall, err := RefET(d.Tmin, d.Tmax, d.Ea, d.Rs, d.Ws, greeley.wz, greeley.z, greeley.lat, d.Date, "mm")
		if err != nil {
			t.Fatalf("Error in conversion or RefET Method: %s", err)
		}

		if math.Abs(wantShort[i]-etShort) > etTolerance {
			t.Errorf("want %f, got %f", wantShort[i], etShort)
		}

		if math.Abs(wantTall[i]-etTall) > etTolerance {
			t.Errorf("want %f, got %f", wantTall[i], etTall)
		}
	}

}
