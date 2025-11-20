package main

type any = interface{}
type object = map[string]any

type Provider struct {
	Name    string
	BaseUrl string
	Models  []string
	ApiKey  string
}

type ProviderArray []Provider

func (arr *ProviderArray) Add(provider Provider) {}
func (arr *ProviderArray) Del(provider Provider) {}
func (arr *ProviderArray) List()                 {}
