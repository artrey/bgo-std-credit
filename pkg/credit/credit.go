package credit

import "math"

func Calculate(sum uint64, periods uint16, percent uint32) (monthlyPayment, overpay, totalSum uint64) {
	// divide by months, percents scale (3 decimal places) and percents
	// more accuracy variant: monthlyPercent := float64(percent) / 12 / 1000 / 100;
	monthlyPercent := math.Round(float64(percent) * 10 / 12 / 100) / 10000;
	increaseKoef := math.Pow(1 + monthlyPercent, float64(periods))
	k := monthlyPercent * increaseKoef / (increaseKoef - 1)
	monthlyPayment = uint64(math.Round(k * float64(sum)))
	totalSum = monthlyPayment * uint64(periods)
	overpay = totalSum - sum
	return
}
