package owners

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

	OwnersCmd.AddCommand(listCmd)
	OwnersCmd.AddCommand(getCmd)

	flags.AddFormat(OwnersCmd)
}

var OwnersCmd = &cobra.Command{
	Use:   "owners",
	Short: "OpenAQ owners",
	Long:  `OpenAQ owners`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := internal.CheckAPIKey()
		if err != nil {
			return err
		}
		return nil
	},
}

func parseFlags(flags *pflag.FlagSet) (*openaq.OwnersArgs, error) {
	ownersArgs := &openaq.OwnersArgs{}
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
	ownersArgs.BaseArgs = baseArgs
	return ownersArgs, nil
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists OpenAQ owners",
	Long:  `Lists OpenAQ owners"`,
	PreRun: func(cmd *cobra.Command, args []string) {

	},
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := internal.SetupClient()
		if err != nil {
			fmt.Println("cannot initialize client")
		}
		ctx := context.Background()
		ownersArgs, err := parseFlags(cmd.Flags())
		if err != nil {
			panic(err)
		}
		owners, err := client.GetOwners(ctx, *ownersArgs)
		if err != nil {
			return internal.ErrorCheck(err)
		}
		res := internal.FormatResult(owners, cmd.Flags())
		fmt.Println(res)
		return nil
	},
}

var getCmd = &cobra.Command{
	Use:   "get [ownersID]",
	Short: "Get a single owner by owners ID",
	Long:  `Get a single owner by owners ID`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ownersID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			panic(err)
		}
		client, err := internal.SetupClient()

		if err != nil {
			fmt.Println("cannot initialize client")
		}
		ctx := context.Background()
		owner, err := client.GetOwner(ctx, ownersID)
		if err != nil {
			panic(err)
		}
		res := internal.FormatResult(owner, cmd.Flags())
		fmt.Println(res)
	},
}
