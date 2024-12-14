package convertor

import (
	"fmt"
	"github.com/antlko/go-currency-convertor/internal/currencyfreaks"
	"github.com/charmbracelet/huh"
	"strconv"
)

func Start(currencyRates currencyfreaks.Latest) error {
	var (
		currFrom string
		currTo   string
		amount   string
	)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose your currency from:").
				Options(
					huh.NewOption("USD", "USD"),
					huh.NewOption("EUR", "EUR"),
					huh.NewOption("UAH", "UAH"),
					huh.NewOption("HUF", "HUF"),
				).
				Value(&currFrom),
			huh.NewSelect[string]().
				Title("Choose your currency to:").
				Options(
					huh.NewOption("USD", "USD"),
					huh.NewOption("EUR", "EUR"),
					huh.NewOption("UAH", "UAH"),
					huh.NewOption("HUF", "HUF"),
				).
				Value(&currTo),
			huh.NewInput().
				Title("Write an amount:").
				Prompt(">").
				Validate(isNumber).
				Value(&amount),
		),
	)

	if err := form.Run(); err != nil {
		return fmt.Errorf("form run: %w", err)
	}

	amountNumber, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return fmt.Errorf("parse amount float: %w", err)
	}
	rateFromNumber, err := strconv.ParseFloat(currencyRates.Rates[currFrom], 64)
	if err != nil {
		return fmt.Errorf("parse amount float: %w", err)
	}
	rateToNumber, err := strconv.ParseFloat(currencyRates.Rates[currTo], 64)
	if err != nil {
		return fmt.Errorf("parse amount float: %w", err)
	}

	// Calculate and print result
	fromInUsd := (amountNumber / rateFromNumber) * rateToNumber
	resultForm := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title(fmt.Sprintf("%s -> %s: %0.2f", currFrom, currTo, fromInUsd)),
		),
	)
	if err := resultForm.Run(); err != nil {
		return fmt.Errorf("result form run: %w", err)
	}
	return nil
}

func isNumber(s string) error {
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("incorrect amount format, should be like: 100, 1.5, 45.12")
	}
	return nil
}
