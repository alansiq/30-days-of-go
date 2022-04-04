package main

import (
	"fmt"
)

func fizzBuzz(n int32) {

	for i := int32(1); i <= n; i++ {
		result := ""
		if i%3 == 0 {
			result += "Fizz"
		}

		if i%5 == 0 {
			result += "Buzz"
		}

		if len([]rune(result)) < 1 {
			fmt.Println(i)
			continue
		}
		fmt.Println(result)
	}

}

func main() {
	fizzBuzz(50)
	fizzBuzz(80)
	fizzBuzz(5)
}
