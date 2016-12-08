// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
//
// Find the largest palindrome made from the product of two 3-digit numbers.

package main

import (
	"fmt"
	"strconv"
)

func main() {

	var reverseStr string
	var intToStr string
	var strToInt int
	var largestPalindrom = 0

	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			intToStr = strconv.Itoa(i * j)
				reverseStr = reverseString(intToStr)
				if intToStr == reverseStr {
					strToInt, _ = strconv.Atoi(intToStr)
					if largestPalindrom < strToInt {
						largestPalindrom = strToInt
					}
				}
		  }
	  }

	fmt.Println(largestPalindrom)

}

func reverseString(s string) string {
	var result string
	for _, value := range s {
		result = string(value) + result
	}
	return result
}
