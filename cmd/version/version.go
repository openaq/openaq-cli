package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

// set versionNumber at build with ldflags
var Version string

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of the OpenAQ CLI",
	Long:  `Print the version number of the OpenAQ CLI`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("openaq v.%s\n", Version)
	},
}
