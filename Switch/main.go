package main

import "fmt"
import "time"

func main() {
	numero := 3
	fmt.Print("O numero ", numero, " se escreve assim:")
	switch numero {
	case 1:
		fmt.Println("um .")
	case 2:
		fmt.Println("dois.")
	case 3:
		fmt.Println("três.")
	}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Pode dormir!")
	default:
		fmt.Println("Vamos trabalhar!")
	}

	numero = 9
	fmt.Println("Este numero cabe num digito?")
	switch {
	case numero < 10:
		fmt.Println("Sim")
	case numero >= 10 && numero < 100:
		fmt.Println("Serve dois digítos")
	case numero >= 100:
		fmt.Println("Nope, número grande!")
	}
}
