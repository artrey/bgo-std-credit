package credit_test

import (
	"fmt"
	"github.com/artrey/bgo-std-credit/pkg/credit"
)

func ExampleCalculate() {
	fmt.Println(credit.Calculate(1_000_000_00, 36, 20_000))
	fmt.Println(credit.Calculate(10_000_00, 36, 20_000))
	// Output:
	// 3718397 33862292 133862292
	// 37184 338624 1338624
}
