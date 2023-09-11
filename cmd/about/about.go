package about

import (
	"fmt"

	"github.com/spf13/cobra"
)

var aboutHeader = `
  ___                      _    ___     ____ _     ___ 
 / _ \ _ __   ___ _ __    / \  / _ \   / ___| |   |_ _|
| | | | '_ \ / _ | '_ \  / _ \| | | | | |   | |    | | 
| |_| | |_) |  __| | | |/ ___ | |_| | | |___| |___ | | 
 \___/| .__/ \___|_| |_/_/   \_\__\_\  \____|_____|___|
      |_|                                              

The OpenAQ CLI is a command line interface tool to access the OpenAQ API.

The source code is open source and available on github at:

https://github.com/openaq/openaq-cli
`

var AboutCmd = &cobra.Command{
	Use:   "about",
	Short: "About the OpenAQ CLI",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(aboutHeader)
	},
}
