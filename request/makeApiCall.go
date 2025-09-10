package request

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func ApiCall() (map[string]interface{}, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	apiKey := os.Getenv("CURRENCY_APP_ID")
	apiURL := os.Getenv("CURRENCY_API_URL")

	apiUrl := apiURL + "?app_id=" + apiKey

	// Make the API call using the formData and the env vars
	// ...
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{
		// CheckRedirect: redirectPolicyFunc,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close() // Done to close the response body

	var result map[string]interface{}
	// Handle the response
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
