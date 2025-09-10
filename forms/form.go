package forms

import (
	"github.com/charmbracelet/huh"
)

type FormData struct {
	currencyFrom string
	currencyTo   string
	amount       string
}

var formData = &FormData{
	currencyFrom: "",
	currencyTo:   "",
	amount:       "",
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
				Value(&formData.currencyFrom),

			// Let the user select multiple toppings.
			huh.NewSelect[string]().
				Title("Choose the currency you want the amount in").
				Options(
					huh.NewOption("USD", "USD"),
					huh.NewOption("EUR", "EUR"),
					huh.NewOption("GBP", "GBP"),
					huh.NewOption("INR", "INR"),
				).
				Value(&formData.currencyTo),

			huh.NewInput().
				Title("How much is the amount you want to convert ").
				Prompt("Enter a number").
				Value(&formData.amount),
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
