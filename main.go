package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var currencyMapping = map[string]float64{
    "USD": 0.151,
    "EUR": 0.137,
    "JPY": 16.29,
    "GBP": 0.13,
    "CHF": 0.1402,
    "AUD": 0.2712,
    "CAD": 0.2374,
    "CNY": 1.251,
    "HKD": 1.326,
    "NZD": 0.2922,
    "SEK": 1.655,
    "NOK": 1.806,
    "DKK": 1.122,
    "SGD": 0.2249,
    "KRW": 242.97,
    "ZAR": 3.239,
    "MXN": 3.454,
    "INR": 14.71,
    "ILS": 0.63,
    "THB": 5.74,
    "IDR": 2875.0,
    "MYR": 0.754,
    "PHP": 9.74,
    "PLN": 0.644,
    "CZK": 3.77,
    "HUF": 61.59,
    "TRY": 6.49,
    "BGN": 0.293,
    "RON": 0.746,
}

type CurrencyConverter struct {
	money float64
	rate float64
}

type TaxJson struct {
	Rates map[string]float64 `json:"rates"`
}

func (cc *CurrencyConverter) SetValues(money string, currency string) (*CurrencyConverter, error) {
	moneyValue, err := strconv.ParseFloat(money, 64)
	if err != nil {
		return nil, errors.New("Invalid money value.")
	}

	file, err := os.ReadFile("tax2.json")
	if err != nil {
		rate, ok := currencyMapping[strings.ToUpper(currency)]
		if !ok {
			return nil, errors.New("Unsupported currency code.")
		}
		cc.money = moneyValue
		cc.rate = rate
		return cc, nil
	}
	
	var result TaxJson
	json.Unmarshal(file, &result)
	rate, ok := result.Rates[strings.ToUpper(currency)]
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
	_, err := currency.SetValues(os.Args[1], os.Args[2])
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
