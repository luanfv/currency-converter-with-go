package main

import (
	"currencyConverter/currency"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Error: Money value is required.")
		return
	}
	if len(os.Args) <= 2 {
		fmt.Println("Error: Currency code is required.")
		return
	}

	currencyConverter, err := currency.CreateCurrencyConverter(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	convertedValue, err := currencyConverter.Convert()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("%.2f\n", convertedValue)
}
