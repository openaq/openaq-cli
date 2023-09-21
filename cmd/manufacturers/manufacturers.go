package manufacturers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/openaq/openaq-go"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/openaq/openaq-cli/cmd/flags"
	internal "github.com/openaq/openaq-cli/cmd/internal"
)

func init() {

	flags.AddLimit(listCmd)
	flags.AddPage(listCmd)

	ManufacturersCmd.AddCommand(listCmd)
	ManufacturersCmd.AddCommand(getCmd)

	flags.AddFormat(ManufacturersCmd)
}

var ManufacturersCmd = &cobra.Command{
	Use:   "manufacturers",
	Short: "OpenAQ manufacturers",
	Long:  `OpenAQ manufacturers`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := internal.CheckAPIKey()
		if err != nil {
			return err
		}
		return nil
	},
}

func parseFlags(flags *pflag.FlagSet) (*openaq.ManufacturerArgs, error) {
	manufacturerArgs := &openaq.ManufacturerArgs{}
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
	manufacturerArgs.BaseArgs = baseArgs
	return manufacturerArgs, nil
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists OpenAQ manufacturers",
	Long:  `Lists OpenAQ manufacturers"`,
	PreRun: func(cmd *cobra.Command, args []string) {

	},
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := internal.SetupClient()
		if err != nil {
			fmt.Println("cannot initialize client")
		}
		ctx := context.Background()
		manufacturerArgs, err := parseFlags(cmd.Flags())
		if err != nil {
			panic(err)
		}
		manufacturers, err := client.GetManufacturers(ctx, *manufacturerArgs)
		if err != nil {
			return internal.ErrorCheck(err)
		}
		res := internal.FormatResult(manufacturers, cmd.Flags())
		fmt.Println(res)
		return nil
	},
}

var getCmd = &cobra.Command{
	Use:   "get [manufacturersID]",
	Short: "Get a single provider by manufacturers ID",
	Long:  `Get a single provider by manufacturers ID`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		manufacturersID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			panic(err)
		}
		client, err := internal.SetupClient()

		if err != nil {
			fmt.Println("cannot initialize client")
		}
		ctx := context.Background()
		manufacturers, err := client.GetManufacturer(ctx, manufacturersID)
		if err != nil {
			panic(err)
		}
		res := internal.FormatResult(manufacturers, cmd.Flags())
		fmt.Println(res)
	},
}
