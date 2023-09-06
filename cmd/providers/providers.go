package providers

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

	ProvidersCmd.AddCommand(listCmd)
	ProvidersCmd.AddCommand(getCmd)

	flags.AddFormat(ProvidersCmd)
}

var ProvidersCmd = &cobra.Command{
	Use:   "providers",
	Short: "OpenAQ providers",
	Long:  `OpenAQ providers`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := internal.CheckAPIKey()
		if err != nil {
			return err
		}
		return nil
	},
}

func parseFlags(flags *pflag.FlagSet) (*openaq.ProvidersArgs, error) {
	providersArgs := &openaq.ProvidersArgs{}
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
	providersArgs.BaseArgs = baseArgs
	return providersArgs, nil
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists OpenAQ providers",
	Long:  `Lists OpenAQ providers"`,
	PreRun: func(cmd *cobra.Command, args []string) {

	},
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := internal.SetupClient()
		if err != nil {
			fmt.Println("cannot initialize client")
		}
		ctx := context.Background()
		providersArgs, err := parseFlags(cmd.Flags())
		if err != nil {
			panic(err)
		}
		providers, err := client.GetProviders(ctx, *providersArgs)
		if err != nil {
			return internal.ErrorCheck(err)
		}
		res := internal.FormatResult(providers, cmd.Flags())
		fmt.Println(res)
		return nil
	},
}

var getCmd = &cobra.Command{
	Use:   "get [providersID]",
	Short: "Get a single provider by providers ID",
	Long:  `Get a single provider by providers ID`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		providersID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			panic(err)
		}
		client, err := internal.SetupClient()

		if err != nil {
			fmt.Println("cannot initialize client")
		}
		ctx := context.Background()
		provider, err := client.GetProvider(ctx, providersID)
		if err != nil {
			panic(err)
		}
		res := internal.FormatResult(provider, cmd.Flags())
		fmt.Println(res)
	},
}
