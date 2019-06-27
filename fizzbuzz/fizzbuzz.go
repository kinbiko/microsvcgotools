package fizzbuzz

import "fmt"

func Fizzbuzz(i int64) string {
	if i%3 == 0 && i%5 == 0 {
		return "fizzbuzz"
	} else if i%3 == 0 {
		return "fizz"
	} else if i%5 == 0 {
		return "buzz"
	}
	return fmt.Sprintf("%d", i)
}
