package RefET

import (
	"strings"
)

type Input struct {
	Value interface{}
	Units string
}

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

// This is the entry function that returns the daily RefET for the day using ASCE method and data provided. Input units can be either F or C
// and will be converted if required
// tmin: float or int
//     Temperature dialy Min in C or F
// tmax:
// 	   Temperature daily max in C or F
// ea : float or int
//    Actual vapor pressure [kPa] or Pa.
// rs : float or int
// 	Incoming shortwave solar radiation [MJ m-2 day-1].
// uz : ndarray
// 	Wind speed [m s-1].
// zw : float
// 	Wind speed height [m].
// elev : ndarray
// 	Elevation [m].
// lat : ndarray
// 	Latitude [degrees].
// doy : ndarray
// 	Day of year.
func RefET(tmin Input, tmax Input, ea Input, rs Input, w Input, elev Input, lat interface{}, date interface{}, rso interface{}) float64 {

	return 0.0
}
