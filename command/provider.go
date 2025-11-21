package command

import "github.com/spf13/cobra"

type Provider struct {
	name    string
	baseUrl string
	models  []string
	apiKey  string
}

type ProviderArray []Provider

func (arr *ProviderArray) Add(provider Provider) {}
func (arr *ProviderArray) Del(provider Provider) {}
func (arr *ProviderArray) List()                 {}

var ProviderCmd = &cobra.Command{
	Use: "provider",
}

func init() {
	providers := ProviderArray{
		Provider{
			name:    "openrouter",
			baseUrl: "https://openrouter.ai/api",
			models:  []string{},
			apiKey:  "",
		},
		Provider{
			name:    "aihubmix",
			baseUrl: "https://aihubmix.com",
			models:  []string{},
			apiKey:  "",
		},
		Provider{
			name:    "gptsapi",
			baseUrl: "https://api.gptsapi.net",
			models:  []string{},
			apiKey:  "",
		},
		Provider{
			name:    "dashscope",
			baseUrl: "https://dashscope.aliyuncs.com/compatible-mode",
			models:  []string{},
			apiKey:  "",
		},
		Provider{
			name:    "ark",
			baseUrl: "https://ark.cn-beijing.volces.com/api",
			models:  []string{},
			apiKey:  "",
		},
	}

	ProviderCmd.AddCommand(&cobra.Command{
		Use:  "list",
		Long: "List all the model providers",
		Run: func(cmd *cobra.Command, args []string) {
			providers.List()
		},
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
