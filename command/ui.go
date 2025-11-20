package command

import "github.com/spf13/cobra"

type UiCmdFlags struct {
	port int
	host string
}

var UiCmd = &cobra.Command{
	Use: "ui",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	flags := UiCmdFlags{}

	UiCmd.Flags().IntVarP(&flags.port, "port", "", 5713, "")
	UiCmd.Flags().StringVarP(&flags.host, "host", "", "localhost", "")
}
