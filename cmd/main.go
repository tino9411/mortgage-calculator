package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tino9411/mortgage-calculator/internal/mortgage"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {

    if shouldQuit(reader) {
      fmt.Print("Exiting the program. Goodbye!")
      break
    }
		fmt.Println("Mortgage Calculator")
		payment := mortgage.NewPayment(reader)
		result, err := payment.CalculateMortgage()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Your monthly payment is: £%.2f\n", result)
	}
}

func shouldQuit(reader *bufio.Reader) bool {
	fmt.Print("Type 'q' to exit or press Enter to continue: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))
	return input == "q"

}
