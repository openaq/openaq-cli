package internal

import (
	"github.com/spf13/viper"
)

func CheckAPIKey() error {
	apiKey := viper.Get("api-key")
	if apiKey == "" || apiKey == nil {
		return &APIKeyNotSetError{}
	}
	return nil
}
