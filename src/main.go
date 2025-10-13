// 		PLEASE RUN CODE BY THE FOLLOWING COMMAND:
// 		$ go run .
// 		OR:
//		$ go run main.go
//
// 		To Build to the .exe:
//		$ go build -o funcgo.exe

package main

import (
	"fmt"
)

func main() {
	
	var x int
	var y int
	var run bool = true
	var mode string

	fmt.Print("1. Addition\n")
	fmt.Print("2. Subtraction\n")
	fmt.Print("3. Multiplication\n")
	fmt.Print("4. Division\n")
	fmt.Print("0. Exit")
	
	for run {
		fmt.Print("\n> ")
		fmt.Scanln(&mode)

		switch mode {
		case "1":
			fmt.Print("\nEnter a number: ")

			_, err1 := fmt.Scanln(&x)
			if err1 != nil {
				fmt.Print("An error has occured: ", err1)
				return
			}

			fmt.Print("Enter another number: ")

			_, err2 := fmt.Scanln(&y)
			if err2 != nil {
				fmt.Print("An error has occured: ", err2)
				return
			}

			fmt.Print(x, " + ", y ," = ", x + y)
		case "2":
			fmt.Print("\nEnter a number: ")

			_, err1 := fmt.Scanln(&x)
			if err1 != nil {
				fmt.Print("An error has occured: ", err1)
				return
			}

			fmt.Print("Enter another number: ")

			_, err2 := fmt.Scanln(&y)
			if err2 != nil {
				fmt.Print("An error has occured: ", err2)
				return
			}

			fmt.Print(x, " - ", y ," = ", x - y)
		case "3":
			fmt.Print("\nEnter a number: ")

			_, err1 := fmt.Scanln(&x)
			if err1 != nil {
				fmt.Print("An error has occured: ", err1)
				return
			}

			fmt.Print("Enter another number: ")

			_, err2 := fmt.Scanln(&y)
			if err2 != nil {
				fmt.Print("An error has occured: ", err2)
				return
			}

			fmt.Print(x, " * ", y ," = ", x * y)
		case "4":
			fmt.Print("\nEnter a number: ")

			_, err1 := fmt.Scanln(&x)
			if err1 != nil {
				fmt.Print("An error has occured: ", err1)
				return
			} else if x == 0 {
				fmt.Print("An error has occured: division by 0")
			}

			fmt.Print("Enter another number: ")

			_, err2 := fmt.Scanln(&y)
			if err2 != nil {
				fmt.Print("An error has occured: ", err2)
				return
			} else if y == 0 {
				fmt.Print("An error has occured: division by 0")
			}

			fmt.Print(x, " / ", y ," = ", x / y, "\n")
			fmt.Print(x, " % ", y ," = ", x % y)
		case "0":
			fmt.Print("Exiting...")
			run = false
		default:
			fmt.Print("An error has occured: unknown mode")
			fmt.Print("\n")
		}
	}
}