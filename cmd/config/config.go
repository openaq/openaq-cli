package config

import (
	"fmt"
	"os"
	"path"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var CfgFile string

func InitConfig() {
	if CfgFile != "" {
		viper.SetConfigFile(CfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		if _, err := os.Stat(path.Join(home, ".openaq.toml")); err != nil {
			fmt.Printf("Config file not found in home directory creating now...\n")
			os.Create(path.Join(home, ".openaq.toml"))
		}
		viper.AddConfigPath(home)
		viper.SetConfigType("toml")
		viper.SetConfigName(".openaq")
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}

type DefaultFormatEnumError struct{}

func (d *DefaultFormatEnumError) Error() string {
	f := color.New(color.FgHiRed).Add(color.Underline)
	return f.Sprint("WARNING: Configuration value format must be one of values 'json', 'csv', ignoring value\n")
}

func isValidFormat(defaultFormat string) bool {
	switch defaultFormat {
	case "":
		return true
	case
		"json",
		"csv":
		return true
	}
	return false
}

func validateConfig() error {
	InitConfig()
	defaultFormat := viper.GetString("defaults.format")
	if valid := isValidFormat(defaultFormat); !valid {
		return &DefaultFormatEnumError{}
	}
	return nil
}

var JSONConfig bool
var CSVConfig bool
var PrettyConfig bool

func ReadConfigValues() error {
	InitConfig()
	err := validateConfig()
	if err != nil {
		return err
	}
	defaultFormat := viper.GetString("defaults.format")
	switch defaultFormat {
	case "json":
		JSONConfig = true
	case "csv":
		CSVConfig = true
	}
	defaultPretty := viper.GetBool("defaults.pretty")
	if defaultPretty {
		PrettyConfig = true
	}
	return nil
}

func init() {
	err := ReadConfigValues()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
