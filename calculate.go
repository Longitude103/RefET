package RefET

import "math"

// atmosPressure is a function to calculate mean Atmospheric Pressure from elevation (meters) (Eq. 3)
func atmosPressure(z float64) float64 {
	paren := (293.0 - 0.0065*z) / 293.0
	paren = math.Pow(paren, 5.26)
	paren *= 101.3

	return paren
}

func psyConst(p float64) float64 {
	return p * 0.000665
}

func meanT(max float64, min float64) float64 {
	return (max + min) / 2
}

func esSlope(tmean float64) float64 {
	e := (17.27 * tmean) / (tmean + 237.3)
	num := 2503.0 * math.Exp(e)
	denom := math.Pow((tmean + 237.3), 2)

	return num / denom
}

func satVP(max float64, min float64) float64 {
	return (eo(max) + eo(min)) / 2
}

func eo(t float64) float64 {
	return 0.6108 * math.Exp((17.27*t)/(t+237.3))
}
