package RefET

import (
	"errors"
	"strings"
)

//  EA has several methods in ASCE Standarized, we support many but not all
// Methods supported:
// 1 - Ea directly measured by station
// 2 - measured or computed dew point
// 5 - max, min relative humidity (add RHmax, RHmin this struct, add temp in struct as well)
// 6 - daily maximum relative humidty (put in Value, add Tmin)
// 7 - daily minimum relative humidty (put in Value, add Tmax)
type EaInput struct {
	Input
	Method int
	RHmax  interface{}
	RHmin  interface{}
	Tmax   Input
	Tmin   Input
}

// ea is a method to return the ea that can be used in the various parts of the app
func (e EaInput) Ea() (ea float64, err error) {
	switch e.Method {
	case 1:
		ea, err = e.convertEA()
		if err != nil {
			return 0.0, err
		}
	case 2:
		ea, err = e.convertFromTdew()
		if err != nil {
			return 0.0, err
		}
	case 5:
		ea, err = e.convertMinMaxRH()
		if err != nil {
			return 0.0, err
		}
	case 6:
		ea, err = e.convertRHmax()
		if err != nil {
			return 0.0, err
		}
	case 7:
		ea, err = e.convertRHmin()
		if err != nil {
			return 0.0, err
		}
	}

	return ea, nil
}

// convertET is an input methodd to convert vapor pressure from pascals to kilopascals
func (e EaInput) convertEA() (v float64, err error) {
	iU := strings.ToLower(e.Units)
	if iU[0] == 'p' {
		v, err = PaToKpa(e.Value)
		if err != nil {
			return 0.0, err
		}
		return v, nil
	}

	v, err = MakeFloat(e.Value)
	if err != nil {
		return 0.0, err
	}
	return v, nil
}

func (e EaInput) convertFromTdew() (ea float64, err error) {
	v, err := e.Input.convertTemp()
	if err != nil {
		return 0.0, err
	}

	ea = eo(v) // Eq. 8
	return ea, nil
}

func (e EaInput) convertMinMaxRH() (ea float64, err error) {
	TmaxV, err := e.Tmax.convertTemp()
	if err != nil {
		return 0.0, errors.New("RH must have valid Tmax")
	}

	TminV, err := e.Tmin.convertTemp()
	if err != nil {
		return 0.0, errors.New("RH must have valid Tmin")
	}

	rhMax, err := MakeFloat(e.RHmax)
	if err != nil {
		return 0.0, errors.New("RHmax must have valid value")
	}

	rhMin, err := MakeFloat(e.RHmin)
	if err != nil {
		return 0.0, errors.New("RHmin must have valid value")
	}

	if rhMax > 1 {
		rhMax = rhMax / 100
	}

	if rhMin > 1 {
		rhMin = rhMin / 100
	}

	ea = ((eo(TminV) * rhMax) + (eo(TmaxV) * rhMin)) / 2 // Eq. 11
	return ea, nil
}

func (e EaInput) convertRHmax() (ea float64, err error) {
	TminV, err := e.Tmin.convertTemp()
	if err != nil {
		return 0.0, errors.New("RHmax must have valid Tmin")
	}

	rhMax, err := MakeFloat(e.RHmax)
	if err != nil {
		return 0.0, errors.New("RHmax must have valid value")
	}

	if rhMax > 1 {
		rhMax = rhMax / 100
	}

	ea = eo(TminV) * rhMax // Eq. 12
	return ea, nil
}

func (e EaInput) convertRHmin() (ea float64, err error) {
	TmaxV, err := e.Tmax.convertTemp()
	if err != nil {
		return 0.0, errors.New("RHmin must have valid Tmax")
	}

	rhMin, err := MakeFloat(e.RHmin)
	if err != nil {
		return 0.0, errors.New("RHmin must have valid value")
	}

	if rhMin > 1 {
		rhMin = rhMin / 100
	}

	ea = eo(TmaxV) * rhMin // Eq. 13
	return ea, nil
}
