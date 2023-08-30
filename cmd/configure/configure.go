package configure

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	apiKeyCmd.AddCommand(setApiKeyCmd)
	ConfigureCmd.AddCommand(apiKeyCmd)
}

var ConfigureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Modify configuration file",
	Long:  ``,
}

var apiKeyCmd = &cobra.Command{
	Use:   "api-key",
	Short: "Set the api-key configuration value",
	Long:  ``,
}

var setApiKeyCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the api-key configuration value",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("api-key", args[0])
		err := viper.WriteConfig()
		if err != nil {
			fmt.Println("Cannot find configuration file")
		}
		fmt.Println("API Key configuration set")
	},
}
