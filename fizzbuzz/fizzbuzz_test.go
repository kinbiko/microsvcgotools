package fizzbuzz_test

import (
	"github.com/kinbiko/microsvcgotools/fizzbuzz"
	"testing"
)

func Test1(t *testing.T)  { test(t, 1, "1") }
func Test3(t *testing.T)  { test(t, 3, "fizz") }
func Test4(t *testing.T)  { test(t, 4, "4") }
func Test5(t *testing.T)  { test(t, 5, "buzz") }
func Test6(t *testing.T)  { test(t, 6, "fizz") }
func Test10(t *testing.T) { test(t, 10, "buzz") }

func test(t *testing.T, in int, exp string) {
	if got := fizzbuzz.Fizzbuzz(in); got != exp {
		t.Errorf("Given %d expect to get '%s' back", in, exp)
	}
}
