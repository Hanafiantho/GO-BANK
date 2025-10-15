package main

import (
	"fmt"
	"os"
)

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main() {
	var menu int
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		fmt.Print("Your choice: ")
		fmt.Scan(&menu)

		// handle condition for check balance
		// if menu == 1 {
		// 	fmt.Println("Your balance is", accountBalance)
		// }

		// handle condition for deposit money
		// if menu == 2 {
		// 	var depositAmount float64

		// 	fmt.Print("Input deposit amount: ")
		// 	fmt.Scan(&depositAmount)

		// 	// deposit amount validation
		// 	if depositAmount <= 0 {
		// 		fmt.Println("Invalid amount! deposit amount must be greater than 0")
		// 		continue // this will makes read code start from the top of loop
		// 	}

		// 	accountBalance += depositAmount // equal to accountBalance = accountBalance + deposit
		// 	fmt.Println("Your updated balance is", accountBalance)
		// 	writeBalanceToFile(accountBalance)
		// }

		// handle condition for withdraw money
		// if menu == 3 {
		// 	var withdrawalAmount float64

		// 	fmt.Print("Input withdrawal amount: ")
		// 	fmt.Scan(&withdrawalAmount)

		// 	// withdrawal amount validation (zero value)
		// 	if withdrawalAmount <= 0 {
		// 		fmt.Println("Invalid amount! withdrawal amount must be greater than 0")
		// 		continue // this will makes read code start from the top of loop
		// 	}

		// 	// withdrawal amount validation (withdrawal amount greater that account balance)
		// 	if withdrawalAmount > accountBalance {
		// 		fmt.Println("Invalid amount! withdrawal amount must be less than account balance")
		// 		continue // this will makes read code start from the top of loop
		// 	}

		// 	accountBalance -= withdrawalAmount // equal to accountBalance = accountBalance - withdrawalAmount
		// 	fmt.Println("Your updated balance is", accountBalance)
		// 	writeBalanceToFile(accountBalance)
		// }

		// handle condition for exit
		// if menu == 4 {
		// 	fmt.Println("Goodbye!")
		// 	break // this will makes read code out of the loop
		// }

		// alternative of using if statement is using switch
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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for choosing our bank!")
			return // in switch break does not mean that it will close from the loop proccess instead of the swith case
		}
	}

	// fmt.Println("Thank you for choosing our bank!")
}
