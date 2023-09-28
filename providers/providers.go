package provider

import "algorand-tiny-watcher/model"

type AlgorandProvider interface {
	IsAddressValid(address string) bool
	AccountState(address string) (*model.AccountState, error)
}
