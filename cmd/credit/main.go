package main

import (
	"fmt"
	"github.com/artrey/bgo-std-credit/pkg/credit"
	"math"
)

func cents2rub(cents uint64) uint64 {
	return uint64(math.Round(float64(cents) / 100))
}

func main() {
	payment, overpay, total := credit.Calculate(1_000_000_00, 36, 20_0)
	fmt.Printf("Monthly payment is %d rub\nOverpay is %d rub\nTotal sum is %d rub",
		cents2rub(payment), cents2rub(overpay), cents2rub(total))
}
