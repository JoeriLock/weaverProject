package energy

import (
	"time"
)

// TODO: use const instead of bool
//takes input, checks type and convert to kwh
func convertToKwh(input float64, isGass bool) float64 {
	if isGass {
		return input * 9.769
	}
	return input / 1000

}

// checks if timestamp is from weekend and night
// TODO: make sure timezone is correct
func parseTime(timeSt int64) (bool, bool) {
	timeStamp := time.Unix(timeSt, 0)
	hr := timeStamp.Hour()
	day := timeStamp.Weekday()
	isWeekend := (day == time.Sunday || day == time.Saturday)
	isNight := (hr < 7 && hr > 23)
	return isNight, isWeekend
}

//Multiply kwh by correct price (depends on date)
func EnergyToCost(energy float64, isGass bool, timeSt int64) float64 {
	isNight, isWeekend := parseTime(timeSt)
	kwh := convertToKwh(energy, isGass)

	if isGass {
		return kwh * 0.06
	}
	if !isNight && !isWeekend {
		return kwh * 0.20
	}
	return kwh * 0.18
}
