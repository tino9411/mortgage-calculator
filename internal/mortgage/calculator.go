package mortgage

import (
	"fmt"
	"math"
)

func (p *Payment) CalculateMortgage() (float64, error) {
	if p.LoanPrincipal <= 0 || p.InterestRate <= 0 || p.LoanTerm <= 0 {
		return 0, fmt.Errorf("Invalid data: LoanPrincipal=%.2f, InterestRate=%.2f, LoanTerm=%d", 
                                            p.LoanPrincipal, p.InterestRate, p.LoanTerm)
	}
	monthlyInterestRate := p.InterestRate / (12 * 100)
	numberOfPayments := float64(p.LoanTerm * 12)
	monthlyPayment := p.LoanPrincipal * monthlyInterestRate * math.Pow(1+monthlyInterestRate, numberOfPayments) / (math.Pow(1+monthlyInterestRate, numberOfPayments) - 1)

	return monthlyPayment, nil
}
