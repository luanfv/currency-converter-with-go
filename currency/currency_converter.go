package currency

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type CurrencyConverter struct {
	money float64
	rate float64
}

func CreateCurrencyConverter(money string, currency string) (*CurrencyConverter, error) {
	cc := &CurrencyConverter{}

	moneyValue, err := strconv.ParseFloat(money, 64)
	if err != nil {
		return nil, errors.New("Invalid money value.")
	}

	file, err := os.ReadFile("currency/tax.json")
	if err != nil {
		rate, ok := CurrencyMapping[strings.ToUpper(currency)]
		if !ok {
			return nil, errors.New("Unsupported currency code.")
		}
		cc.money = moneyValue
		cc.rate = rate
		return cc, nil
	}
	
	var result taxJson
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
