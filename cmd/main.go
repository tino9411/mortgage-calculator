package main

import (
	"fmt"
	"mortgage-calculator/internal/mortgage"
)

func main() {

	fmt.Println("Mortgage Calculator")
	payment := mortgage.NewPayment()
	result, err := payment.CalculateMortgage()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Your monthly payment is: £%.2f\n", result)

}
