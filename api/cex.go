package api

import (
	"encoding/json"
	"fmt"
	"go_crud/datatypes"
	"io"
	"net/http"
	"strings"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	if len(currency) != 3 {
		return nil, fmt.Errorf("currency must be 3 characters, received %d", len(currency))
	}

	currency = strings.ToUpper(currency)
	resp, err := http.Get(fmt.Sprintf(apiUrl, currency))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var response Response

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(body, &response)

		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("status code received: %v", resp.StatusCode)
	}

	rate := datatypes.Rate{
		Currency: currency,
		Price:    response.Bid,
	}

	return &rate, nil
}
