// The prime factors of 13195 are 5, 7, 13 and 29.
//
// What is the largest prime factor of the number 600851475143 ?


package main

import (
  "fmt"
)

func main(){
  var isPrime bool
  myPrimeFactors := []int64{}
  mainNum := int64(600851475143)
  largestFactorLimit := int64(10000)
  for i:= int64(2); i < largestFactorLimit; i++ {
    isPrime = false
    if mainNum % i == 0 {
      for j:= int64(2); j < i; j++ {
        if i % j == 0 {
          //Factor is not prime
          isPrime = false
          break
        } else {
          isPrime = true
        }
      }
      if isPrime {
        myPrimeFactors = append(myPrimeFactors, i)
      }
    }
  }
  fmt.Println(myPrimeFactors[len(myPrimeFactors)-1:len(myPrimeFactors)])
}
