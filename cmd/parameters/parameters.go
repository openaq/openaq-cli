package parameters

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
	ParametersCmd.AddCommand(listCmd)
	ParametersCmd.AddCommand(getCmd)
}

var ParametersCmd = &cobra.Command{
	Use:   "parameters",
	Short: "OpenAQ parameters",
	Long:  `OpenAQ parameters`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := internal.CheckAPIKey()
		if err != nil {
			return err
		}
		return nil
	},
}

func parseFlags(flags *pflag.FlagSet) (*openaq.ParametersArgs, error) {
	parametersArgs := &openaq.ParametersArgs{}
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
	parametersArgs.BaseArgs = baseArgs
	parameterType, err := flags.GetString("type")
	if err != nil {
		return nil, err
	}
	parametersArgs.ParameterType = parameterType
	return parametersArgs, nil
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists OpenAQ parameters",
	Long:  `Lists OpenAQ parameters"`,
	PreRun: func(cmd *cobra.Command, args []string) {

	},
	RunE: func(cmd *cobra.Command, args []string) error {
		config := openaq.Config{
			APIKey: viper.GetString("api-key"),
		}
		client, err := openaq.NewClient(config)
		if err != nil {
			fmt.Println("cannot initialize client")
		}
		ctx := context.Background()
		parametersArgs, err := parseFlags(cmd.Flags())
		if err != nil {
			panic(err)
		}
		parameters, err := client.GetParameters(ctx, *parametersArgs)
		if err != nil {
			return internal.ErrorCheck(err)
		}
		res := internal.FormatResult(parameters, cmd.Flags())
		fmt.Println(res)
		return nil
	},
}

var getCmd = &cobra.Command{
	Use:   "get [parametersID]",
	Short: "Get a single parameter by parameters ID",
	Long:  `Get a single parameter by parameters ID`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		parametersID, err := strconv.ParseInt(args[0], 10, 64)
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
		parameter, err := client.GetParameter(ctx, parametersID)
		if err != nil {
			panic(err)
		}
		res := internal.FormatResult(parameter, cmd.Flags())
		fmt.Println(res)
	},
}
