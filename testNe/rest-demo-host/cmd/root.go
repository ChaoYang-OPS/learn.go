package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var vers bool

var RootCmd = &cobra.Command{
	Use: "demo-api",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println("version info")
			return nil
		}
		return nil
	},
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "....")
}
