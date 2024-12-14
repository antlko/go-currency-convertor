package internal

import (
	"fmt"
	"github.com/antlko/go-currency-convertor/internal/convertor"
	"github.com/antlko/go-currency-convertor/internal/currencyfreaks"
)

func InitService(cfg AppConfig) error {
	currencyApi := currencyfreaks.NewApi(cfg.CurrencyFreaks.ApiKey)

	latestCurrRates, err := currencyApi.GetLatestCurrencies()
	if err != nil {
		return fmt.Errorf("get latest currencies: %w", err)
	}

	if err := convertor.Start(latestCurrRates); err != nil {
		return fmt.Errorf("convertor: %w")
	}

	return nil
}
