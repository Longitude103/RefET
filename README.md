# RefET
RefET Equation using ASCE Standardized in GO from "The ASCE Standardized Reference Evapotranspiration Equation" by Richard Allen and Ivan Walter et. al. published by the Environemental Water Resources Institute of the Americal Society of Civil Engineers (2005). This package does not implement the hourly calculation method.

## What is RefET
This is a Go package that was created based on the ASCE Standardized Reference Evapotranspriration Equation (RefET). This package takes in a daily weather information and returns two values and and error. The first value is the daily "Short crop RefET" value and the second is the daily "Tall crop RefET" value. The application can also return an error if there is a conversion problem or missing data. 

### Reference Crops
The short crop is a reference to a crop similar to clipped grass whereas a tall crop is similar to full-cover alfalfa. The function in this package returns both of these values for each daily weather input.

### Output
The output of the application includes the two values and each are in mm/day. The output units can be changed to yeild anwers in in/day as well. There are no hourly calculations preformed by this package, only the daily calculations.

## Input Data
The input data required for this package is daily climate data. The climate data required includes Min and Max Temperature, measured solar radiation, average wind speed, and date. A measurement of vapor pressure is preferred, but can be calculated with other methods. The methods implemented here are not all the methods within the ASCE standardized equation. They are as follows:

| Method No. |                                   Method                                   | Preference Ranking | Equation(s) | Required Data                           |
| :--------: | :------------------------------------------------------------------------: | :----------------: | :---------: | --------------------------------------- |
|     1      |                     ea averaged over the daily period                      |         1          |      7      | ea, units                               |
|     2      | Measured or computed dew point  temperature averaged over the daily period |         1          |      8      | dew point, units                        |
|     5      |                Daily maximum and minimum relative humidity                 |         2          |    7, 11    | RHmin, RHmax, Tmin, Tmax, units of each |
|     6      |                      Daily maximum relative humidity                       |         3          |    7, 12    | RHmax, Tmin, units of each              |
|     7      |                      Daily maximum relative humidity                       |         3          |    7, 13    | RHmin, Tmax, units of each              |

The station information requirements are the wind speed height measurement (standard is 2 meters), the station elevation, and the station latitude in degrees. 

The final portion of the function is how to output the units either in "mm" or "in".

## Package Use
Th entry function that returns the daily short crop and tall crop RefET for the day using ASCE method is `RefET()`. The input units can very for each argument in the function and will be converted if required. Use the Exposed `Input` struct to format the data correctly to be used by the function. These are the required arguments for the function and associated units that can be used, ea has it's own `EaInput` struct to use since there are a variety of methods and combinations that can be used to calculate it.  Please see the many tests for example usage.

Daily Climate requirements: 
- tmin: float or int => Temperature dialy Min in C or F
- tmax: float or int => Temperature daily max in C or F
- ea: float or int => Actual measured vapor pressure (kPa or Pa) or alternate methods using methods 1, 2, 5, 6, 7 in ASCE reference book, the "Method" is referenced as part of the struct and corresponds to the above table of methods
- rs: float or int => Incoming shortwave solar radiation [MJ m-2 day-1].
- uz: float or int => Wind speed (m s-1, mph).
- date: date => date of calculation in the form of mm-dd-year (month 1 must be 01 and day must be 02)

Daily Climate Station Information Requirements:
- zw: float or int => Wind speed height (m or ft).
- z: float or int => Elevation (m or ft).
- lat: float => Latitude (degrees).


Output Format
- outUnits: "mm" or "in" => units that will be output from the equation, mm = millimeters / day || in = inches / day

There are many tests that show how to use this method including the "ea" functions that are included in the "eaInput_test.go" file. This will be of help on how to format the struct.
