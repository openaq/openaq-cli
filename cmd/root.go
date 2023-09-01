/*
Copyright © OpenAQ <dev@openaq.org>
*/
package cmd

import (
	"fmt"
	"path"

	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/openaq/openaq-cli/cmd/configure"
	"github.com/openaq/openaq-cli/cmd/countries"
	"github.com/openaq/openaq-cli/cmd/locations"
	"github.com/openaq/openaq-cli/cmd/measurements"
	"github.com/openaq/openaq-cli/cmd/parameters"
	"github.com/openaq/openaq-cli/cmd/providers"
	"github.com/openaq/openaq-cli/cmd/version"
)

var (
	// Used for flags.
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "openaq",
		Short: "Command Line Interface for the OpenAQ API",
		Long:  ``,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.openaq.toml)")

	addLimit(countries.CountriesCmd)
	addPage(countries.CountriesCmd)
	addFormat(countries.CountriesCmd)

	addLimit(locations.LocationsCmd)
	addPage(locations.LocationsCmd)
	addFormat(locations.LocationsCmd)
	addCountries(locations.LocationsCmd)
	addProviders(locations.LocationsCmd)

	addLimit(measurements.MeasurementsCmd)
	addPage(measurements.MeasurementsCmd)
	addFormat(measurements.MeasurementsCmd)
	addFromDate(measurements.MeasurementsCmd)
	addToDate(measurements.MeasurementsCmd)
	addPeriodName(measurements.MeasurementsCmd)
	addMini(measurements.MeasurementsCmd)
	addParameters(measurements.MeasurementsCmd)

	addLimit(parameters.ParametersCmd)
	addPage(parameters.ParametersCmd)
	addFormat(parameters.ParametersCmd)
	addParametersType(parameters.ParametersCmd)

	addLimit(providers.ProvidersCmd)
	addPage(providers.ProvidersCmd)
	addFormat(providers.ProvidersCmd)
	addParametersType(providers.ProvidersCmd)

	rootCmd.AddCommand(configure.ConfigureCmd)
	rootCmd.AddCommand(countries.CountriesCmd)
	rootCmd.AddCommand(locations.LocationsCmd)
	rootCmd.AddCommand(measurements.MeasurementsCmd)
	rootCmd.AddCommand(parameters.ParametersCmd)
	rootCmd.AddCommand(providers.ProvidersCmd)
	rootCmd.AddCommand(version.VersionCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
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
