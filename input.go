package RefET

import "strings"

type Input struct {
	Value interface{}
	Units string
}

// convertTemp is a method on the input to convert the temperature from Feherheit to Celsius
func (i Input) convertTemp() (v float64, err error) {
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

// convertRS is an input method to convert short wave radiation from langleys or watts to megajoules per day
func (i Input) convertRS() (v float64, err error) {
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

// convertWS is an input method to convert wind speed from miles per hour into meters per second
func (i Input) convertWS() (v float64, err error) {
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

// convertWZ is an input method to convert windspeed height from feet into meters
func (i Input) convertWZ() (v float64, err error) {
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

// convertZ is an input method to convert elevation from feet into meters
func (i Input) convertZ() (v float64, err error) {
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

// convertLat is a method that converts the latitude from degrees into radians
func (i Input) convertLat() (v float64, err error) {
	v, err = DegreesToRad(i.Value)
	if err != nil {
		return 0.0, err
	}

	return v, nil
}

// convertDate is a method that converts the date in format ("07-01-2000") into a day of the year
func (i Input) convertDate() (v int, err error) {
	v, err = DateToDOY(i.Value)
	if err != nil {
		return 0, err
	}

	return v, nil
}
