package watcher

import (
	"algorand-tiny-watcher/config"
	"algorand-tiny-watcher/model"
	"algorand-tiny-watcher/pkg/persist"
	provider "algorand-tiny-watcher/providers"
	"algorand-tiny-watcher/providers/algorand"
	"errors"
	"sync"
	"time"
)

type Address string

type Watcher struct {
	sync.RWMutex
	config    *config.WatcherConfig
	provider  provider.AlgorandProvider
	addresses map[string]*model.AccountState
}

func NewWatcher(config *config.WatcherConfig) *Watcher {
	algorand := algorand.NewAlgorandClient(config.AlgorandProviderUrl)
	addresses, err := persist.LoadAddresses()
	if err != nil {
		logger.Err(err).Msg("Failed to load addresses file")
	}
	if addresses == nil {
		initialAddressesCapacity := 100
		addresses = make(map[string]*model.AccountState, initialAddressesCapacity)
	}
	return &Watcher{
		provider:  algorand,
		config:    config,
		addresses: addresses,
	}
}

func (w *Watcher) AddAddress(address string) error {
	w.RWMutex.Lock()
	defer w.RWMutex.Unlock()

	_, exists := w.addresses[address]
	if exists {
		return errors.New("address already being watched")
	}

	w.addresses[address] = nil

	persist.SaveAddresses(w.addresses)

	return nil
}

func (w *Watcher) RemoveAddress(address string) error {
	w.RWMutex.Lock()
	defer w.RWMutex.Unlock()

	_, exists := w.addresses[address]
	if !exists {
		return errors.New("address not found in watcher")
	}

	delete(w.addresses, address)

	return nil
}

func (w *Watcher) GetState() map[string]*model.AccountState {
	w.RWMutex.RLock()
	defer w.RWMutex.RUnlock()

	copiedAddresses := make(map[string]*model.AccountState, len(w.addresses))
	for k, v := range w.addresses {
		if v != nil {
			copiedAddresses[k] = v.Clone()
		} else {
			copiedAddresses[k] = nil
		}
	}

	return copiedAddresses
}

func (w *Watcher) IsAddressValid(address string) bool {
	return w.provider.IsAddressValid(address)
}

func (w *Watcher) watch() {
	w.RWMutex.Lock()
	defer w.RWMutex.Unlock()

	for addr, state := range w.addresses {
		updatedState, err := w.provider.AccountState(addr)
		if err != nil {
			logger.Err(err).Msg("Error fetching account state from provider")
		}

		isNewState := state == nil
		hasChanged := (state != nil && state.HasChanged(updatedState))

		if isNewState || hasChanged {
			w.addresses[addr] = updatedState
			printAccountState(updatedState)
			err := persist.SaveAddresses(w.addresses)
			if err != nil {
				logger.Err(err).Msg("Error saving addresses cache file")
			}
		}
	}
}

func (w *Watcher) Run() error {
	initLogger()

	duration := time.Duration(w.config.UpdateTimeoutInSeconds) * time.Second
	ticker := time.NewTicker(duration)

	for {
		logger.Info().Msg("Running a watcher update")
		w.watch()
		<-ticker.C
	}
}
