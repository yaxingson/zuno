package command

import "github.com/spf13/cobra"

var ProviderCmd = &cobra.Command{
	Use: "provider",
}

func init() {
	ProviderCmd.AddCommand(&cobra.Command{
		Use:  "list",
		Long: "List all the model providers",
		Run:  func(cmd *cobra.Command, args []string) {},
	})

	ProviderCmd.AddCommand(&cobra.Command{
		Use:  "current",
		Long: "Show current provider name",
		Run:  func(cmd *cobra.Command, args []string) {},
	})

	ProviderCmd.AddCommand(&cobra.Command{
		Use:  "use",
		Long: "Change current provider",
		Run:  func(cmd *cobra.Command, args []string) {},
	})

	ProviderCmd.AddCommand(&cobra.Command{
		Use:  "add",
		Long: "Add other provider",
		Run:  func(cmd *cobra.Command, args []string) {},
	})

	ProviderCmd.AddCommand(&cobra.Command{
		Use:  "cost",
		Long: "Show current cost for specific or all provider",
		Run:  func(cmd *cobra.Command, args []string) {},
	})

}
