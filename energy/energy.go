package energy

// TODO: use const instead of bool
func convertToKwh(input float64, isGass bool) float64 {
	if isGass {
		return input * 9.769
	}
	return input / 1000

}

// TODO: use date not bool
func EnergyToCost(energy float64, isGass bool, isNight bool, isWeekend bool) float64 {
	kwh := convertToKwh(energy, isGass)
	if isGass {
		return kwh * 0.06
	}
	if !isNight && !isWeekend {
		return kwh * 0.20
	}
	return kwh * 0.18
}
