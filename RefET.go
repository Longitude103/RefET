package RefET

// This is the entry function that returns the daily Short and Tall RefET for the day using ASCE method and data provided.
// Input units can very for each item and will be converted if required. Use the Exposed "Input" struct to put the data into the function.
// These are the required data and the associated units that can be used, ea has it's own "EaInput" struct to use since there are a variety
// of methods and combinations that can be used to calculate it.  Please see the many tests for example usage.
// tmin: float or int
//     Temperature dialy Min in C or F
// tmax: float or int
// 	   Temperature daily max in C or F
// ea : float or int
//    Actual measured vapor pressure [kPa] or Pa.
//    or alternate methods using method 1, 2, 5, 6, 7
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
// OutUnits:
// two choices are "mm" (default) or "in" for the standard millimeters or a conversion to inches
func RefET(tmin Input, tmax Input, ea EaInput, rs Input, ws Input, zw Input, z Input, lat Input, date Input, outUnits interface{}) (etShort float64, etTall float64, err error) {

	tm, err := tmin.convertTemp()
	if err != nil {
		return 0.0, 0.0, err
	}

	tma, err := tmax.convertTemp()
	if err != nil {
		return 0.0, 0.0, err
	}

	e, err := ea.Ea()
	if err != nil {
		return 0.0, 0.0, err
	}

	r, err := rs.convertRS()
	if err != nil {
		return 0.0, 0.0, err
	}

	w, err := ws.convertWS()
	if err != nil {
		return 0.0, 0.0, err
	}

	wz, err := zw.convertWZ()
	if err != nil {
		return 0.0, 0.0, err
	}

	ele, err := z.convertZ()
	if err != nil {
		return 0.0, 0.0, err
	}

	l, err := lat.convertLat()
	if err != nil {
		return 0.0, 0.0, err
	}

	d, err := date.convertDate()
	if err != nil {
		return 0.0, 0.0, err
	}

	etShort, etTall = calculateRefET(tma, tm, e, r, w, wz, ele, l, d)

	if outUnits == "in" {
		etShort *= 0.03937
		etTall *= 0.03937
	}

	return etShort, etTall, nil
}

// calculateRefET is the raw function that can be called if all units are in the correct format and everything is assured to have data. The
// assumption is that you will format and put all the data in correctly, no error checking or conversions.
func calculateRefET(tmax float64, tmin float64, ea float64, rs float64, ws float64, wz float64, z float64, lat float64, doy int) (etShort float64, etTall float64) {
	const lambda = 0.408 // Latent Heat of Vaporization
	const G = 0          // soil heat flux daily is zero

	// atmospheric pressure
	atmosP := atmosPressure(z)

	// psycometric constant
	gamma := psyConst(atmosP)

	// mean temperature
	tmean := meanT(tmax, tmin)

	// slope of vapor pressure
	delta := esSlope(tmean)

	// saturated vapor pressure
	es := satVP(tmax, tmin)

	// // Actual Vapor Pressure from Tdew
	// eaFromDew := eo(ea)
	// fmt.Println("EA from Dew", eaFromDew)

	ra := calcRA(lat, doy)
	rso := calcRSO(ra, z)
	fcd := calcFCD(rs, rso)
	rnl := calcRNL(fcd, ea, tmax, tmin)

	rns := calcRNS(rs)
	rn := calcRN(rns, rnl)

	adjWS := calcWS(ws, wz)

	etShortNum := lambda*delta*(rn-G) + gamma*(900/(tmean+273))*adjWS*(es-ea)
	etShortDenom := delta + gamma*(1+0.34*adjWS)

	etTallNum := lambda*delta*(rn-G) + gamma*(1600/(tmean+273))*adjWS*(es-ea)
	etTallDenom := delta + gamma*(1+0.38*adjWS)

	return etShortNum / etShortDenom, etTallNum / etTallDenom
}
