package watcher

import (
	"algorand-tiny-watcher/model"
)

func printAccountState(account *model.AccountState) {
	// To log amounts, convert from microalgos to algos
	logger.Info().
		Str("Address", account.Address).
		Float64("Amount", float64(account.Amount)/1000000).
		Float64("AmountWithoutPendingRewards", float64(account.AmountWithoutPendingRewards)/1000000).
		Float64("PendingRewards", float64(account.PendingRewards)/1000000).
		Float64("Rewards", float64(account.Rewards)/1000000).
		Str("Status", account.Status).
		Msg(" ")

	for _, asset := range account.Assets {
		logger.Info().
			Str("Name", asset.Name).
			Uint64("Asset", asset.AssetID).
			Float64("Amount", float64(asset.Amount)/1000000).Msg("   ")
	}
}
