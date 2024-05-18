package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const apiUrl = "https://v6.exchangerate-api.com/v6/%s/latest/AUD"

type ExchangeRateResponse struct {
	BaseCode        string             `json:"base_code"`
	ConversionRates map[string]float64 `json:"conversion_rates"`
}

func main() {
	apiKey := os.Getenv("EXCHANGE_RATE_API_KEY") // Get the API key from the environment // export EXCHANGE_RATE_API_KEY="YOUR_API_KEY"
	if apiKey == "" {
		fmt.Println("Error: EXCHANGE_RATE_API_KEY environment variable is not set")
		os.Exit(1)
	}

	url := fmt.Sprintf(apiUrl, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received non-200 response code: %d\n", resp.StatusCode)
		os.Exit(1)
	}

	var exchangeRateResponse ExchangeRateResponse
	if err := json.NewDecoder(resp.Body).Decode(&exchangeRateResponse); err != nil {
		fmt.Printf("Error parsing response: %v\n", err)
		os.Exit(1)
	}

	usdRate, exists := exchangeRateResponse.ConversionRates["USD"]
	if !exists {
		fmt.Println("Error: USD rate not found in response")
		os.Exit(1)
	}

	audToUsd := usdRate
	usdToAud := 1 / usdRate

	fmt.Printf("The exchange rate from AUD to USD is: %.4f\n", audToUsd)
	fmt.Printf("The exchange rate from USD to AUD is: %.4f\n", usdToAud)
}
