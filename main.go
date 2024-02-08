package main

import (
	"log"
	"myfirstgo/calculator"
)

func main() {
	log.Printf("My Calculator")

	log.Printf("1+1=%v", calculator.NewTheBestCalculator("D ja").Add(1, 1))

	// log.Printf("1+1=%v", calculator.Add(1, 1))
	// log.Printf("1.0+1.0=%v", 1.0+1.0)
	// log.Printf("1.0*1.0=%.2f", 1.0*1.0)
	// log.Printf("1.0/1.0=%+v", 1.0/1.0)

}
