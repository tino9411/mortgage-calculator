package main

import (
	"fmt"
	"math"
)

type payment struct {
	loanPrincipal float64
	interestRate  float64
	loanTerm      int
}

func (p *payment) CalculateMortgage() (float64, error) {
	if p.loanPrincipal <= 0 || p.interestRate <= 0 || p.loanTerm <= 0 {
		return 0, fmt.Errorf("Invalid data: loanPrincipal=%.2f, interestRate=%.2f, loanTerm=%d", p.loanPrincipal, p.interestRate, p.loanTerm)
	}
	monthlyInterestRate := p.interestRate / (12 * 100)
	numberOfPayments := float64(p.loanTerm * 12)
	monthlyPayment := p.loanPrincipal * monthlyInterestRate * math.Pow(1+monthlyInterestRate, numberOfPayments) / (math.Pow(1+monthlyInterestRate, numberOfPayments) - 1)

	return monthlyPayment, nil
}

func IntReader(prompt string) int {
	for {
		var value int
		fmt.Print(prompt)
		_, err := fmt.Scanln(&value)
		if err != nil || value <= 0 {
			fmt.Println("Invalid input. Please try again")
			continue
		}
		return value

	}
}

func FloatReader(prompt string) float64 {
	for {
		var value float64
		fmt.Print(prompt)
		_, err := fmt.Scanln(&value)
		if err != nil || value <= 0 {
			fmt.Println("Invalid input. Please try again")
			continue
		}
		return value

	}
}

func NewPayment() *payment {
	p := &payment{}
	p.loanPrincipal = FloatReader("Enter the loan principal (e.g., 100000): ")
	p.interestRate = FloatReader("Enter the annual interest rate (e.g., 5): ")
	p.loanTerm = IntReader("Enter the loan term in years (e.g., 30): ")
	fmt.Printf("\nSummary:\nLoan Principal: £%.2f\nInterest Rate: %.2f%%\nLoan Term: %d years\n\n", p.loanPrincipal, p.interestRate, p.loanTerm)
	return p
}
func main() {

	fmt.Println("Mortgage Calculator")

	payment := NewPayment()
	result, err := payment.CalculateMortgage()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Your monthly payment is: £%.2f\n", result)

}
