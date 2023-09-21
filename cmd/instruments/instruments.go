package instruments

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

	InstrumentsCmd.AddCommand(listCmd)
	InstrumentsCmd.AddCommand(getCmd)

	flags.AddFormat(InstrumentsCmd)
}

var InstrumentsCmd = &cobra.Command{
	Use:   "instruments",
	Short: "OpenAQ instruments",
	Long:  `OpenAQ instruments`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := internal.CheckAPIKey()
		if err != nil {
			return err
		}
		return nil
	},
}

func parseFlags(flags *pflag.FlagSet) (*openaq.InstrumentArgs, error) {
	instrumentsArgs := &openaq.InstrumentArgs{}
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
	instrumentsArgs.BaseArgs = baseArgs
	return instrumentsArgs, nil
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists OpenAQ instruments",
	Long:  `Lists OpenAQ instruments"`,
	PreRun: func(cmd *cobra.Command, args []string) {

	},
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := internal.SetupClient()
		if err != nil {
			fmt.Println("cannot initialize client")
		}
		ctx := context.Background()
		instrumentsArgs, err := parseFlags(cmd.Flags())
		if err != nil {
			panic(err)
		}
		instruments, err := client.GetInstruments(ctx, *instrumentsArgs)
		if err != nil {
			return internal.ErrorCheck(err)
		}
		res := internal.FormatResult(instruments, cmd.Flags())
		fmt.Println(res)
		return nil
	},
}

var getCmd = &cobra.Command{
	Use:   "get [instrumentsID]",
	Short: "Get a single instrument by instruments ID",
	Long:  `Get a single instrument by instruments ID`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		instrumentsID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			panic(err)
		}
		client, err := internal.SetupClient()

		if err != nil {
			fmt.Println("cannot initialize client")
		}
		ctx := context.Background()
		instrument, err := client.GetInstrument(ctx, instrumentsID)
		if err != nil {
			panic(err)
		}
		res := internal.FormatResult(instrument, cmd.Flags())
		fmt.Println(res)
	},
}
