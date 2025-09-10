package main

import (
	"fmt"

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
	fmt.Println("API call successful:", message)
}
