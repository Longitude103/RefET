package RefET

import "errors"

func MakeFloat(c interface{}) (float64, error) {
	var value float64

	switch t := c.(type) {
	case int64:
		value = float64(t)
	case int32:
		value = float64(t)
	case int:
		value = float64(t)
	case float64:
		value = t
	case float32:
		value = float64(t)
	default:
		return 0.0, errors.New("failed to convert celsius to fehrenheit")
	}

	return value, nil
}

func CToF(c interface{}) (float64, error) {
	value, err := MakeFloat(c)
	if err != nil {
		return 0.0, err
	}

	value = value*9/5 + 32
	return value, nil
}

func FToC(f interface{}) (float64, error) {
	value, err := MakeFloat(f)
	if err != nil {
		return 0.0, err
	}

	value = (value - 32) * 5 / 9
	return value, nil
}

func PaToKpa(pa interface{}) (float64, error) {
	value, err := MakeFloat(pa)
	if err != nil {
		return 0.0, err
	}

	value = value / 1000
	return value, nil
}

func LangToMJ(lang interface{}) (float64, error) {
	value, err := MakeFloat(lang)
	if err != nil {
		return 0.0, err
	}

	value = value * 0.04184
	return value, nil
}

func WattToMJ(watt interface{}) (float64, error) {
	value, err := MakeFloat(watt)
	if err != nil {
		return 0.0, err
	}

	value = value * 0.0864
	return value, nil
}
