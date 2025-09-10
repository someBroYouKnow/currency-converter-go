package forms

import (
	"github.com/charmbracelet/huh"
)

type FormData struct {
	CurrencyFrom string
	CurrencyTo   string
	Amount       string
}

var formData = &FormData{
	CurrencyFrom: "",
	CurrencyTo:   "",
	Amount:       "",
}

func (formData *FormData) InitializeForm() error {

	form := huh.NewForm(
		huh.NewGroup(
			// Ask the user for a base burger and toppings.
			huh.NewSelect[string]().
				Title("Choose the currency you want to convert to another one").
				Options(
					huh.NewOption("USD", "USD"),
					huh.NewOption("EUR", "EUR"),
					huh.NewOption("GBP", "GBP"),
					huh.NewOption("INR", "INR"),
				).
				Value(&formData.CurrencyFrom),

			// Let the user select multiple toppings.
			huh.NewSelect[string]().
				Title("Choose the currency you want the amount in").
				Options(
					huh.NewOption("USD", "USD"),
					huh.NewOption("EUR", "EUR"),
					huh.NewOption("GBP", "GBP"),
					huh.NewOption("INR", "INR"),
				).
				Value(&formData.CurrencyTo),

			huh.NewInput().
				Title("How much is the amount you want to convert ").
				Prompt("Enter a number").
				Value(&formData.Amount),
		),
	)

	err := form.Run()
	if err != nil {
		panic(err)
	}
	return nil
}

func GetFormData() *FormData {
	return formData
}
