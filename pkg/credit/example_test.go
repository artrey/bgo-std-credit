package credit_test

import (
	"fmt"
	"github.com/artrey/bgo-std-credit/pkg/credit"
)

func ExampleCalculate() {
	fmt.Println(credit.Calculate(1_000_000_00, 36, 20_0))
	fmt.Println(credit.Calculate(10_000_00, 36, 20_0))
	fmt.Println(credit.Calculate(4_000_000_00, 84, 8_9))
	// Output:
	// 3716358 33788888 133788888
	// 37164 337904 1337904
	// 6415350 138889400 538889400
}
