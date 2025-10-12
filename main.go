// PLEASE RUN CODE BY THE FOLLOWING COMMAND:
// $ go run .
//
// To Build to an .exe:
// $ go build -o exename.exe

package main

import (
	"fmt"
)

func main() {
	var x int
	var y int
	sleepTime := 8        //! Change The Time of the sleep at the end.
	var mode string

	fmt.Print("1. Adding\n")
	fmt.Print("2. Subtraction\n")
	fmt.Print("3. Multiplying")
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
	default:
		fmt.Print("An error has occured: unknown mode")
	}
	fmt.Print("\n")
	Wait(sleepTime)
}