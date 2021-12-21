package RefET

import "math"

// atmosPressure is a function to calculate mean Atmospheric Pressure from elevation (meters) (Eq. 3)
func atmosPressure(z float64) float64 {
	paren := (293.0 - 0.0065*z) / 293.0
	paren = math.Pow(paren, 5.26)
	paren *= 101.3

	return paren
}

// psyConst is a function to calculate the Psychrometric Constant proportional to the mean atmospheric pressure (Eq. 4)
func psyConst(p float64) float64 {
	return p * 0.000665
}

// meanT is a function to calculate the mean temperature (Eq. 2)
func meanT(max float64, min float64) float64 {
	return (max + min) / 2
}

// esSlope is a function to caluclate the slope of the Saturation Vapor Pressure-Temperature Curve (delta) (Eq. 5)
func esSlope(tmean float64) float64 {
	e := (17.27 * tmean) / (tmean + 237.3)
	num := 2503.0 * math.Exp(e)
	denom := math.Pow((tmean + 237.3), 2)

	return num / denom
}

// satVP is a function to calculate the daily Saturation Vapor Pressure (Eq. 6)
func satVP(max float64, min float64) float64 {
	return (eo(max) + eo(min)) / 2
}

// eo is a function to calculate the saturation vapor pressure at a given temperature (Eq. 7)
func eo(t float64) float64 {
	return 0.6108 * math.Exp((17.27*t)/(t+237.3))
}

func inverseRelDistFactor(doy int) float64 {
	return 1 + 0.033*math.Cos((2*math.Pi/365)*float64(doy)) // Eq. 23
}

func solarDeclin(doy int) float64 {
	return 0.409 * math.Sin((2*math.Pi/365)*float64(doy)-1.39) // Eq. 24
}

func sunsetHourAngle(lat float64, delta float64) float64 {
	return math.Acos(-1 * math.Tan(lat) * math.Tan(delta)) // Eq. 27
}

// calcRA is a function to calculate the Extraterrestrial Radiation for 24-Hour Periods, phi is latitude, doy is day-of-year
func calcRA(phi float64, doy int) float64 {
	dr := inverseRelDistFactor(doy)
	delta := solarDeclin(doy)
	omega := sunsetHourAngle(phi, delta)

	ra := 24 / math.Pi * 4.92 * dr * (omega*math.Sin(phi)*math.Sin(delta) + math.Cos(phi)*math.Cos(delta)*math.Sin(omega)) // Eq. 21

	return ra
}

// calcRSO is the function to compute clear-sky solar radition, z is station elevation in meters
func calcRSO(ra float64, z float64) float64 {
	return (0.75 + 2e-5*z) * ra // Eq. 19
}

// calcFCD is a function to compute cloudiness function
func calcFCD(rs float64, rso float64) float64 {
	relSR := rs / rso

	if relSR < 0.3 {
		relSR = 0.3
	}

	if relSR > 1.0 {
		relSR = 1.0
	}

	return 1.35*(relSR) - 0.35 // Eq. 18
}

// calcRNL is a function to compute net long-wave radiation
func calcRNL(fcd float64, ea float64, tmax float64, tmin float64) float64 {
	const sigma = 4.901e-9

	return sigma * fcd * (0.34 - 0.14*math.Sqrt(ea)) * ((math.Pow(tmax+273.16, 4) + math.Pow(tmin+273.16, 4)) / 2) // Eq. 17
}

// calcRNS is a function to compute net solar or short-wave radiation
func calcRNS(rs float64) float64 {
	const alpha = 0.23
	return (1 - alpha) * rs // Eq. 16
}

func calcRN(rns float64, rnl float64) float64 {
	return rns - rnl // Eq. 15
}

// calcWS is a function to calculate the wind speed and adjust for the standard 2m height, ws is windspeed at wz height in meters
func calcWS(ws float64, wz float64) float64 {
	if wz == 2.0 {
		return ws
	}

	return ws * (4.87 / math.Log(67.8*wz-5.42)) // Eq. 33
}
