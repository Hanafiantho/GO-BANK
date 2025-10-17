package main

import (
	"fmt"
	"go-bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	var menu int
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("____________")
		// panic(err) // it will show the error message and stop the application
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		bankMenu()

		fmt.Print("Your choice: ")
		fmt.Scan(&menu)

		switch menu {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			var depositAmount float64

			fmt.Print("Input deposit amount: ")
			fmt.Scan(&depositAmount)

			// deposit amount validation
			if depositAmount <= 0 {
				fmt.Println("Invalid amount! deposit amount must be greater than 0")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Your updated balance is", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			var withdrawalAmount float64

			fmt.Print("Input withdrawal amount: ")
			fmt.Scan(&withdrawalAmount)

			// withdrawal amount validation (zero value)
			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount! withdrawal amount must be greater than 0")
				continue
			}

			// withdrawal amount validation (withdrawal amount greater that account balance)
			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount! withdrawal amount must be less than account balance")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Your updated balance is", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for choosing our bank!")
			return // in switch, break does not mean that it will close from the loop proccess instead of the swith case
		}
	}
}
