package settings

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	setCmd.AddCommand(setAPIKeyCmd)
	setCmd.AddCommand(setDefaultFormatCmd)
	setCmd.AddCommand(setDefaultPrettyCmd)

	getCmd.AddCommand(getAPIKeyCmd)
	getCmd.AddCommand(getDefaultFormatCmd)
	getCmd.AddCommand(getDefaultPrettyCmd)

	SettingsCmd.AddCommand(setCmd)
	SettingsCmd.AddCommand(getCmd)
}

var SettingsCmd = &cobra.Command{
	Use:   "settings",
	Short: "Modify settings",
	Long:  ``,
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set a value in the configuration file",
	Long:  ``,
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "set a value in the configuration file",
	Long:  ``,
}

var setAPIKeyCmd = &cobra.Command{
	Use:   "api-key",
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

var getAPIKeyCmd = &cobra.Command{
	Use:   "api-key",
	Short: "Get the api-key configuration value",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("api-key")
		fmt.Println(apiKey)
	},
}

var setDefaultFormatCmd = &cobra.Command{
	Use:   "format",
	Short: `Set the default output format`,
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "json", "csv":
			viper.Set("defaults.format", args[0])
		case "none":
			viper.Set("defaults.format", "")
		default:
			fmt.Println(`Default format value must be "json","csv", or "none"`)
			os.Exit(1)
		}
		err := viper.WriteConfig()
		if err != nil {
			fmt.Println("Cannot find configuration file")
		}
		fmt.Printf("Default output format set to %s\n", args[0])
	},
}

var getDefaultFormatCmd = &cobra.Command{
	Use:   "format",
	Short: `Get the default output format`,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		defaultFormat := viper.Get("defaults.format")
		fmt.Println(defaultFormat)
	},
}

var setDefaultPrettyCmd = &cobra.Command{
	Use:   "pretty",
	Short: `Set the default for pretty print`,
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "false", "true":
			viper.Set("defaults.pretty", args[0])
		default:
			fmt.Println(`Default pretty value must be either true or false`)
			os.Exit(1)
		}
		err := viper.WriteConfig()
		if err != nil {
			fmt.Println("Cannot find configuration file")
		}
		fmt.Printf("Default output pretty print set to %s\n", args[0])
	},
}

var getDefaultPrettyCmd = &cobra.Command{
	Use:   "pretty",
	Short: `Get the default for pretty print`,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		defaultPretty := viper.Get("defaults.pretty")
		fmt.Println(defaultPretty)
	},
}
