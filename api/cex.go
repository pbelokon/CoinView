package api

import (
	"coinView/go/crypto/rates"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*rates.Rate, error) {
	if len(currency) != 3 {
		return nil, fmt.Errorf("3 characters required: %d received", len(currency))
	}
	
	// http and https are the same here with ssl
	upCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))

	if err != nil { 
		return nil, err
	} 
	
	var response CEXResponse

	if res.StatusCode == http.StatusOK { 
		// wait for all chunks 
		data, err := io.ReadAll(res.Body)

		if err != nil { 
			return nil, err
		}


		err = json.Unmarshal(data, &response)

		if err != nil { 
			return nil, err
		}


	} else { 
		return nil, fmt.Errorf("status code received %v", res.StatusCode)
	}

	rate := rates.Rate{Currency: currency, Price: response.Ask}

	return &rate, nil
}