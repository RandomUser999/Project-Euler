// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
//
// What is the 10 001st prime number?

package main

import (
  "fmt"
)

func main() {
  primNumList := []int{2,3,5,7,11,13}
  isPrime := false

  for i := 2;; i++ {
    for _, value := range primNumList {
      if i % value == 0 {
        //not a prime number
        isPrime = false
        break
      } else {
        isPrime = true
      }
    }
    if isPrime {
     //is a prime number
     primNumList = append(primNumList, i)
     if len(primNumList) == 10001 {
       break
     }
   }
  }
  fmt.Println(primNumList[len(primNumList)-1:len(primNumList)])
}
