Mortgage Calculator

A simple mortgage calculator CLI application written in Go. This tool calculates the monthly payment for a mortgage based on the loan principal, annual interest rate, and loan term provided by the user.

Features

	•	Takes user input for:
	•	Loan Principal
	•	Annual Interest Rate
	•	Loan Term (in years)
	•	Validates user input to ensure positive values are provided.
	•	Calculates the monthly payment using the standard mortgage formula.
	•	Provides clear error messages and prompts for retrying invalid input.

Prerequisites

	•	Go (version 1.20 or higher recommended)

Getting Started

1. Clone the Repository

git clone https://github.com/your-username/mortgage-calculator.git
cd mortgage-calculator

2. Run the Program

go run main.go

Usage

	1.	Launch the program by running go run main.go.
	2.	Follow the prompts to enter:
	•	Loan principal (e.g., 100000 for £100,000).
	•	Annual interest rate (e.g., 5 for 5%).
	•	Loan term in years (e.g., 30 for 30 years).
	3.	The program will calculate and display the monthly payment.

Example

Input:

Enter the loan principal (e.g., 100000): 100000
Enter the annual interest rate (e.g., 5): 5
Enter the loan term in years (e.g., 30): 30

Output:

Your monthly payment is: £536.82

Formula Used

The monthly mortgage payment is calculated using the formula:
￼

Where:
	•	￼: Monthly payment
	•	￼: Loan principal
	•	￼: Monthly interest rate (annual rate divided by 12 and converted to decimal)
	•	￼: Total number of payments (loan term in years multiplied by 12)

Code Highlights

	•	Input Handling:
	•	Functions IntReader and FloatReader handle user input with validation and retry logic.
	•	Modular Design:
	•	The payment struct encapsulates mortgage data and calculation logic.
	•	The NewPayment function initializes a payment struct with user input.
	•	Error Handling:
	•	Invalid inputs are caught, and the user is prompted to re-enter values.

Future Improvements

	•	Add support for an amortization schedule to show interest vs. principal breakdowns over time.
	•	Save user input and results to a file.
	•	Build a web interface for easier usage.

License

This project is licensed under the MIT License. Feel free to use, modify, and distribute.

