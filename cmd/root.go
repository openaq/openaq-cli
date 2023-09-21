/*
Copyright © OpenAQ <dev@openaq.org>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/openaq/openaq-cli/cmd/about"
	"github.com/openaq/openaq-cli/cmd/config"

	"github.com/openaq/openaq-cli/cmd/countries"
	"github.com/openaq/openaq-cli/cmd/instruments"
	"github.com/openaq/openaq-cli/cmd/locations"
	"github.com/openaq/openaq-cli/cmd/manufacturers"
	"github.com/openaq/openaq-cli/cmd/measurements"
	"github.com/openaq/openaq-cli/cmd/owners"
	"github.com/openaq/openaq-cli/cmd/parameters"
	"github.com/openaq/openaq-cli/cmd/providers"
	"github.com/openaq/openaq-cli/cmd/settings"
	"github.com/openaq/openaq-cli/cmd/version"
)

var (
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
	cobra.OnInitialize(config.InitConfig)
	rootCmd.PersistentFlags().StringVarP(&config.CfgFile, "config", "c", "", "config file (default is $HOME/.openaq.toml)")

	rootCmd.AddCommand(settings.SettingsCmd)
	rootCmd.AddCommand(countries.CountriesCmd)
	rootCmd.AddCommand(instruments.InstrumentsCmd)
	rootCmd.AddCommand(locations.LocationsCmd)
	rootCmd.AddCommand(manufacturers.ManufacturersCmd)
	rootCmd.AddCommand(measurements.MeasurementsCmd)
	rootCmd.AddCommand(parameters.ParametersCmd)
	rootCmd.AddCommand(providers.ProvidersCmd)
	rootCmd.AddCommand(owners.OwnersCmd)
	rootCmd.AddCommand(version.VersionCmd)
	rootCmd.AddCommand(about.AboutCmd)
}
