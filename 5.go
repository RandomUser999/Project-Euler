// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
//
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

package main

import (
  "fmt"
)

func main() {

  var (
    upperLimit int
    isDivisible bool
  )

  upperLimit = factorial(20)

  for i:= 20; i < upperLimit; i += 20 {
    for j:= 20; j > 0; j-- {
      if i % j == 0 {
        isDivisible = true
      } else {
        isDivisible = false
        break
      }
    }
    if isDivisible {
      fmt.Println(i)
      break
    }
  }
}

func factorial(i int) int {
  if i == 0 {
    return 1
  } else {
    return i * factorial(i-1)
  }
}
