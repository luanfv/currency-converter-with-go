package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var currencyMapping = map[string]float64{
	"USD": 0.151,
	"EUR": 0.137,
}

type CurrencyConverter struct {
	money float64
	rate float64
}

func (cc *CurrencyConverter) SetValues(money string, currency string) (*CurrencyConverter, error) {
	moneyValue, err := strconv.ParseFloat(money, 64)
	if err != nil {
		return nil, errors.New("Invalid money value.")
	}
	rate, ok := currencyMapping[strings.ToUpper(currency)]
	if !ok {
		return nil, errors.New("Unsupported currency code.")
	}
	cc.money = moneyValue
	cc.rate = rate
	return cc, nil
}

func (cc *CurrencyConverter) ok() bool {
	if cc.money <= 0 && cc.rate <= 0 {
		return false
	}
	return true
}

func (cc *CurrencyConverter) Convert() (float64, error) {
	if !cc.ok() {
		return -1, errors.New("Invalid conversion values.")
	}
	return cc.money * cc.rate, nil
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Error: Money value is required.")
		return
	}
	if len(os.Args) <= 2 {
		fmt.Println("Error: Currency code is required.")
		return
	}

	currency := CurrencyConverter{}
	_, err :=currency.SetValues(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	convertedValue, err := currency.Convert()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("%.2f\n", convertedValue)
}
