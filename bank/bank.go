package bank

import (
	"fmt"

	fileHandling "example.hu/app/utility/filehandling"
)

const defaultBalance float64 = 1500.5

func Banking() {
	accountBalance, fileReadingError := fileHandling.ReadFloatFromFile("balance.txt")
	if fileReadingError != nil {
		fmt.Println(fmt.Sprintf("Error reading balance: %v\n", fileReadingError))
		fmt.Println("Using the default value for the balance.")
		accountBalance = defaultBalance
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		printMenuOptions()
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Balance: %.2f\n", accountBalance)
			break
		case 2:
			accountBalance = depositAmount(accountBalance)
			fileHandling.WriteFloatToFile("balance.txt", accountBalance)
			break
		case 3:
			withDrawAmount(&accountBalance)
			fileHandling.WriteFloatToFile("balance.txt", accountBalance)
			break
		default:
			fmt.Println("Good bye!")
			return
		}
	}
}

func printMenuOptions() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance.")
	fmt.Println("2. Deposit")
	fmt.Println("3. Widthdraw")
	fmt.Println("4. Exit")
}

func depositAmount(balance float64) float64 {
	fmt.Println("Amount of deposit:")
	var depositAmount float64
	fmt.Scan(&depositAmount)
	balance += depositAmount

	return balance
}

func withDrawAmount(accountBalance *float64) {
	fmt.Println("Amount of withdraw:")
	var amount float64
	fmt.Scan(&amount)
	(*accountBalance) -= amount
}
