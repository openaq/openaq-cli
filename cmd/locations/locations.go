package locations

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
	LocationsCmd.AddCommand(listCmd)
	LocationsCmd.AddCommand(getCmd)
}

var LocationsCmd = &cobra.Command{
	Use:   "locations",
	Short: "OpenAQ locations",
	Long:  "OpenAQ locations",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := internal.CheckAPIKey()
		if err != nil {
			return err
		}
		return nil
	},
}

func parseFlags(flags *pflag.FlagSet) (*openaq.LocationArgs, error) {
	locationArgs := &openaq.LocationArgs{}
	baseArgs := &openaq.BaseArgs{}
	limit, err := flags.GetInt64("limit")
	if err != nil {
		return nil, err
	}
	page, err := flags.GetInt64("page")
	if err != nil {
		return nil, err
	}
	baseArgs.Page = page
	baseArgs.Limit = limit
	locationArgs.BaseArgs = baseArgs
	var countries openaq.Countries
	var providers openaq.Providers

	countries_ids, err := flags.GetInt64Slice("countries")
	if err != nil {
		return nil, err
	}
	if len(countries_ids) > 0 {
		countries = openaq.Countries{
			IDs: countries_ids,
		}
		locationArgs.Countries = &countries
	}
	providers_ids, err := flags.GetInt64Slice("providers")
	if err != nil {
		return nil, err
	}
	if len(providers_ids) > 0 {
		providers = openaq.Providers{
			IDs: providers_ids,
		}
		locationArgs.Providers = &providers
	}

	isoCode, err := flags.GetString("iso")
	if err != nil {
		return nil, err
	}
	if isoCode != "" {
		locationArgs.IsoCode = isoCode
	}

	return locationArgs, nil
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists OpenAQ locations",
	Long:  "Lists OpenAQ locations",
	Run: func(cmd *cobra.Command, args []string) {
		config := openaq.Config{
			APIKey: viper.GetString("api-key"),
		}
		client, err := openaq.NewClient(config)
		if err != nil {
			fmt.Println("cannot initialize client")
		}
		ctx := context.Background()
		locationArgs, err := parseFlags(cmd.Flags())
		if err != nil {
			panic(err)
		}
		locations, err := client.GetLocations(ctx, *locationArgs)
		if err != nil {
			panic(err)
		}
		res := internal.FormatResult(locations, cmd.Flags())
		fmt.Println(res)
	},
}

var getCmd = &cobra.Command{
	Use:   "get [locationsID]",
	Short: "Get a single location by location ID",
	Long:  `Get a single location by location ID`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		locationId, err := strconv.ParseInt(args[0], 10, 64)
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
		location, err := client.GetLocation(ctx, locationId)
		if err != nil {
			panic(err)
		}
		res := internal.FormatResult(location, cmd.Flags())
		fmt.Println(res)
	},
}
