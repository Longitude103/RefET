package RefET

import (
	"strings"
)

type Input struct {
	Value interface{}
	Units string
}

// ConvertTemp is a method on the input to convert the temperature from Feherheit to Celsius
func (i Input) ConvertTemp() (v float64, err error) {
	iU := strings.ToLower(i.Units)
	if iU[0] == 'f' {
		v, err = FToC(i.Value)
		if err != nil {
			return 0.0, err
		}
		return v, nil
	}

	v, err = MakeFloat(i.Value)
	if err != nil {
		return 0.0, err
	}

	return v, nil
}

// ConvertET is an input methodd to convert vapor pressure from pascals to kilopascals
func (i Input) ConvertEA() (v float64, err error) {
	iU := strings.ToLower(i.Units)
	if iU[0] == 'p' {
		v, err = PaToKpa(i.Value)
		if err != nil {
			return 0.0, err
		}
		return v, nil
	}

	v, err = MakeFloat(i.Value)
	if err != nil {
		return 0.0, err
	}
	return v, nil
}

// ConvertRS is an input method to convert short wave radiation from langleys or watts to megajoules per day
func (i Input) ConvertRS() (v float64, err error) {
	iU := strings.ToLower(i.Units)
	if iU[0] == 'l' {
		// langleys
		v, err = LangToMJ(i.Value)
		if err != nil {
			return 0.0, err
		}
		return v, nil
	} else if iU[0] == 'w' {
		// w/m-2
		v, err = WattToMJ(i.Value)
		if err != nil {
			return 0.0, err
		}
		return v, nil
	}

	v, err = MakeFloat(i.Value)
	if err != nil {
		return 0.0, err
	}

	return v, nil
}

// ConvertWS is an input method to convert wind speed from miles per hour into meters per second
func (i Input) ConvertWS() (v float64, err error) {
	iU := strings.ToLower(i.Units)
	if iU == "mph" {
		// miles per hour
		v, err = MphToMS(i.Value)
		if err != nil {
			return 0.0, err
		}

		return v, nil
	}

	v, err = MakeFloat(i.Value)
	if err != nil {
		return 0.0, err
	}

	return v, nil
}

// ConvertZ is an input method to convert elevation from feet into meters
func (i Input) ConvertZ() (v float64, err error) {
	iU := strings.ToLower(i.Units)
	if iU[0] == 'f' {
		// feet
		v, err = FeetToMeters(i.Value)
		if err != nil {
			return 0.0, err
		}

		return v, nil
	}

	v, err = MakeFloat(i.Value)
	if err != nil {
		return 0.0, err
	}

	return v, nil
}

func (i Input) ConvertLat() (v float64, err error) {
	v, err = DegreesToRad(i.Value)
	if err != nil {
		return 0.0, err
	}

	return v, nil
}

func (i Input) ConvertDate() (v int, err error) {
	v, err = DateToDOY(i.Value)
	if err != nil {
		return 0, err
	}

	return v, nil
}

// This is the entry function that returns the daily RefET for the day using ASCE method and data provided. Input units can be either F or C
// and will be converted if required
// tmin: float or int
//     Temperature dialy Min in C or F
// tmax: float or int
// 	   Temperature daily max in C or F
// ea : float or int
//    Actual vapor pressure [kPa] or Pa.
// rs : float or int
// 	Incoming shortwave solar radiation [MJ m-2 day-1].
// uz : float or int
// 	Wind speed [m s-1].
// zw : float or int
// 	Wind speed height [m].
// z: float or int
// 	Elevation [m].
// lat : float
// 	Latitude [degrees].
// date : date
// 	date of calculation in the form of mm-dd-year. and month 1 must be 01 and day must be 02
func RefET(tmin Input, tmax Input, ea Input, rs Input, ws Input, z Input, lat Input, date Input) (float64, error) {

	tm, err := tmin.ConvertTemp()
	tma, err := tmax.ConvertTemp()
	e, err := ea.ConvertEA()
	r, err := rs.ConvertRS()
	w, err := ws.ConvertWS()
	ele, err := z.ConvertZ()
	l, err := lat.ConvertLat()
	d, err := date.ConvertDate()

	if err != nil {
		return 0.0, err
	}

	return calculateRefET(tm, tma, e, r, w, ele, l, d), nil
}

func calculateRefET(tmin float64, tmax float64, ea float64, rs float64, ws float64, z float64, lat float64, doy int) float64 {

	atmosP := atmosPressure(z)
	psy := psyConst(atmosP)

	tmean := meanT(tmax, tmin)
	esS := esSlope(tmean)

	esP := satVP(tmax, tmin)

	_ = psy
	_ = esS
	_ = esP

	return 0.0
}
