package internal

import (
	"github.com/openaq/openaq-go"
	"github.com/spf13/viper"
)

// https://freshman.tech/snippets/go/concatenate-slices/
func appendMany[T any](slices [][]T) []T {
	var totalLen int

	for _, s := range slices {
		totalLen += len(s)
	}

	result := make([]T, totalLen)

	var i int

	for _, s := range slices {
		i += copy(result[i:], s)
	}
	return result
}

func SetupClient() (*openaq.Client, error) {
	config := openaq.Config{
		APIKey:    viper.GetString("api-key"),
		UserAgent: "openaq-cli",
	}
	client, err := openaq.NewClient(config)
	if err != nil {
		return nil, err
	}
	return client, nil
}
