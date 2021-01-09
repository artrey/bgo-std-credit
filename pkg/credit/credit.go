package credit

import (
	"math"
)

// Calculate returns credit stats: monthly payment, overpay and total sum.
func Calculate(sum uint64, periods uint16, percent uint16) (monthlyPayment, overpay, totalSum uint64) {
	// divide by months, percents scale (1 decimal place) and percents
	monthlyPercent := float64(percent) / 12 / 10 / 100
	increaseFactor := math.Pow(1 + monthlyPercent, float64(periods))
	k := monthlyPercent * increaseFactor / (increaseFactor - 1)
	monthlyPayment = uint64(math.Round(k * float64(sum)))
	totalSum = monthlyPayment * uint64(periods)
	overpay = totalSum - sum
	return
}
