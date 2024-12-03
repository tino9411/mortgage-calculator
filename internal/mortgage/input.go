package mortgage

import (
	"bufio"
	"fmt"
	"os"
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

func NewPayment() *Payment {
  reader := bufio.NewReader(os.Stdin)
	p := &Payment{}
	p.LoanPrincipal = FloatReader("Enter the loan principal (e.g., 100000): ")
	p.InterestRate = FloatReader("Enter the annual interest rate (e.g., 5): ")
	p.LoanTerm = IntReader("Enter the loan term in years (e.g., 30): ", reader )
	fmt.Printf("\nSummary:\nLoan Principal: £%.2f\nInterest Rate: %.2f%%\nLoan Term: %d years\n\n", p.LoanPrincipal, p.InterestRate, p.LoanTerm)
	return p
}
