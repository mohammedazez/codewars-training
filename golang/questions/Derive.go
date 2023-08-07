package questions

import "fmt"

func Derive(coefficient, exponent int) string {
	return fmt.Sprintf("%v%s%v", coefficient*exponent, "x^", exponent-1)
}
