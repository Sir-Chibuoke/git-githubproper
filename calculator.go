// How to calculate in Go Language
package main

import "fmt"

func main() {
	var num1 float32
	var num2 float32
	var operator string
	for {
		

			fmt.Printf("\nPlease kindly input the values that you want to calculate")
			fmt.Printf("\nEnter the first Number\n")
			fmt.Scan(&num1)
			fmt.Printf("\nEnter the second Number\n")
			fmt.Scan(&num2)
			fmt.Printf("\nEnter the operator (* or / or + or -)\n")
			fmt.Scan(&operator)
			switch operator {
			case "+":
				fmt.Printf("\nThe sum of %v and %v is: %v ", num1, num2, num1+num2)
				break
			case "-":
				fmt.Printf("\nThe difference of %v and %v is: %v ", num1, num2, num1-num2)
				break

			case "*":
				fmt.Printf("\nThe product of %v and %v is: %v ", num1, num2, num1*num2)
				break
			case "/":
				fmt.Printf("\nThe dividend of %.3v and %.3v is: %.3v ", num1, num2, num1/num2)
				break
			default:
				fmt.Println("Invalid operator")
			}

		}
	}
