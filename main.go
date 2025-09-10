package main

import (
	"fmt"
	"strconv"

	"currency/forms"
	"currency/request"
)

func main() {
	fmt.Println("Welcome to my currency converter!")
	formData := forms.GetFormData()
	err := formData.InitializeForm()
	if err != nil {
		fmt.Println("Error initializing form:", err)
		return
	}
	fmt.Println("Form submitted successfully!")
	// You can call your API function here after form submission
	message, err := request.ApiCall()
	if err != nil {
		fmt.Println("Error making API call:", err)
		return
	}

	currencyFrom := getRate(message, formData.CurrencyFrom)
	currencyTo := getRate(message, formData.CurrencyTo)
	amount, err := strconv.ParseFloat(formData.Amount, 64)
	if err != nil {
		fmt.Println("Invalid amount entered:", err)
		return
	}

	result := amount * (currencyTo / currencyFrom)

	fmt.Printf("Amount %v from currency %v to currency %v is %v \n \n ", amount, formData.CurrencyFrom, formData.CurrencyTo, result)
	fmt.Printf("My logic is that %v to USD is %v and %v to USD is %v", formData.CurrencyFrom, currencyFrom, formData.CurrencyTo, currencyTo)
}

func getRate(result map[string]interface{}, key string) float64 {
	rates, ok := result["rates"].(map[string]interface{})
	if !ok {
		fmt.Println("Error in accessing rates from api response")
		return 0
	}

	rate, ok := rates[key].(float64)
	if !ok {
		fmt.Printf("Error in accessing %v rate\n", key)
		return 0
	}
	return rate
}
