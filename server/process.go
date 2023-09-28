package server

import (
	"algorand-tiny-watcher/model"
	"errors"
)

func (w *TinyWatcherServer) processWatchAddress(address string) error {
	if !w.watcher.IsAddressValid(address) {
		return errors.New("invalid algorand wallet address")
	}

	err := w.watcher.AddAddress(address)
	if err != nil {
		return err
	}

	return nil
}

func (w *TinyWatcherServer) processWatchState() map[string]*model.AccountState {
	return w.watcher.GetState()
}
