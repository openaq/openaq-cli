package countries

import (
	"context"
	"fmt"
	"strconv"

	"github.com/openaq/openaq-go"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	internal "github.com/openaq/openaq-cli/cmd/internal"
)

func init() {
	CountriesCmd.AddCommand(listCmd)
	CountriesCmd.AddCommand(getCmd)
}

func parseFlags(flags *pflag.FlagSet) (*openaq.CountryArgs, error) {
	baseArgs := openaq.BaseArgs{}
	limit, err := flags.GetInt64("limit")
	if err != nil {
		return nil, err
	}
	baseArgs.Limit = limit
	page, err := flags.GetInt64("page")
	if err != nil {
		return nil, err
	}
	baseArgs.Page = page
	countriesArgs := &openaq.CountryArgs{
		BaseArgs: baseArgs,
	}

	return countriesArgs, nil
}

var CountriesCmd = &cobra.Command{
	Use:   "countries",
	Short: "Countries",
	Long:  `Countries`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := internal.CheckAPIKey()
		if err != nil {
			return err
		}
		return nil
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List countries",
	Long:  `List countries`,
	Run: func(cmd *cobra.Command, args []string) {
		config := openaq.Config{
			APIKey: viper.GetString("api-key"),
		}
		client, err := openaq.NewClient(config)
		if err != nil {
			fmt.Println("cannot initialize client")
		}
		ctx := context.Background()
		countriesArgs, err := parseFlags(cmd.Flags())
		if err != nil {
			panic(err)
		}
		countries, err := client.GetCountries(ctx, *countriesArgs)
		if err != nil {
			panic(err)
		}
		res := internal.FormatResult(countries, cmd.Flags())
		fmt.Println(res)
	},
}

var getCmd = &cobra.Command{
	Use:   "get [countriesID]",
	Short: "Get a single country by countries ID",
	Long:  `Get a single country by countries ID`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		countriesID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			panic(err)
		}
		config := openaq.Config{
			APIKey: viper.GetString("api-key"),
		}
		client, err := openaq.NewClient(config)
		if err != nil {
			fmt.Println("cannot initialize client")
		}
		ctx := context.Background()
		country, err := client.GetCountry(ctx, countriesID)
		if err != nil {
			panic(err)
		}
		res := internal.FormatResult(country, cmd.Flags())
		fmt.Println(res)
	},
}
