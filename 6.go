// The sum of the squares of the first ten natural numbers is,
// 12 + 22 + ... + 102 = 385
//
// The square of the sum of the first ten natural numbers is,
// (1 + 2 + ... + 10)2 = 552 = 3025
//
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
//
// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

package main

import (
  "fmt"
)

func main() {
  numTarget := 100
  sum := 0

  sum = squareNum(squareOfSum(numTarget), 2) - sumOfSquare(numTarget)

  fmt.Println(sum)

}

func sumOfSquare (SoS int) int {
  if SoS == 0 {
    return 1
  }
  return SoS * SoS + (sumOfSquare(SoS-1))
}

func squareOfSum (SoS int) int {
  if SoS == 0 {
    return 0
  }
  return SoS + squareOfSum(SoS-1)
}

func squareNum (A int, B int) int {
  C := 1
  if B == 0 {
    return 1
  }
  for i := B; i > 0; i-- {
    C *= A
  }
  return C
}
