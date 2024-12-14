package currencyfreaks

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Latest struct {
	Date  string            `json:"date"`
	Base  string            `json:"base"`
	Rates map[string]string `json:"rates"`
}

type Config struct {
	ApiKey string `env:"CURRENCYFREAKS_API_KEY"`
}

const getLatestCurrencyUrl = "https://api.currencyfreaks.com/latest?apikey=%s"

type Api struct {
	apikey string
}

func NewApi(apikey string) Api {
	return Api{apikey: apikey}
}

func (a Api) GetLatestCurrencies() (Latest, error) {
	resp, err := http.Get(fmt.Sprintf(getLatestCurrencyUrl, a.apikey))
	if err != nil {
		return Latest{}, fmt.Errorf("get latest currencies: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return Latest{}, fmt.Errorf("not ok status code: %d", resp.StatusCode)
	}
	if resp.Body == nil {
		return Latest{}, fmt.Errorf("body not exist")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Latest{}, fmt.Errorf("read body: %w", err)
	}
	if len(body) == 0 {
		return Latest{}, fmt.Errorf("body is empty")
	}

	var response Latest
	if err := json.Unmarshal(body, &response); err != nil {
		return Latest{}, fmt.Errorf("unmarshalling: %w", err)
	}
	return response, nil
}
