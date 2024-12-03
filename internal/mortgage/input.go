package mortgage

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func IntReader(prompt string, reader *bufio.Reader) int {
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid input. Please try again")
			continue
		}

		input = strings.TrimSpace(input)
		result, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			fmt.Println("Error parsing input: ", err)
		}
		return int(result)
	}
}

func FloatReader(prompt string, reader *bufio.Reader) float64 {
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid input. Please try again")
			continue
		}
		input = strings.TrimSpace(input)
		result, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Error parsing input: ", err)
		}
		return result

	}
}

func NewPayment(reader *bufio.Reader) *Payment {
	p := &Payment{}
	p.LoanPrincipal = FloatReader("Enter the loan principal (e.g., 100000): ", reader)
	p.InterestRate = FloatReader("Enter the annual interest rate (e.g., 5): ", reader)
	p.LoanTerm = IntReader("Enter the loan term in years (e.g., 30): ", reader)
	fmt.Printf("\nSummary:\nLoan Principal: £%.2f\nInterest Rate: %.2f%%\nLoan Term: %d years\n\n", p.LoanPrincipal, p.InterestRate, p.LoanTerm)
	return p
}
