package model

import (
	"reflect"
)

type AccountState struct {
	Address                     string  `json:"address"`
	Amount                      uint64  `json:"amount"`
	AmountWithoutPendingRewards uint64  `json:"amountWithoutPendingRewards"`
	PendingRewards              uint64  `json:"pendingRewards"`
	Rewards                     uint64  `json:"rewards"`
	Status                      string  `json:"status"`
	Assets                      []Asset `json:"assets"`
}

type Asset struct {
	AssetID uint64 `json:"assetID"`
	Name    string `json:"name"`
	Amount  uint64 `json:"amount"`
}

func (s *AccountState) HasChanged(newState *AccountState) bool {
	if s.Amount != newState.Amount {
		return true
	}
	if !reflect.DeepEqual(s.Assets, newState.Assets) {
		return true
	}

	return false
}

func (s *AccountState) Clone() *AccountState {
	clonedAssets := make([]Asset, len(s.Assets))
	for i, asset := range s.Assets {
		clonedAssets[i] = Asset{
			AssetID: asset.AssetID,
			Name:    asset.Name,
			Amount:  asset.Amount,
		}
	}

	return &AccountState{
		Address:                     s.Address,
		Amount:                      s.Amount,
		AmountWithoutPendingRewards: s.AmountWithoutPendingRewards,
		PendingRewards:              s.PendingRewards,
		Rewards:                     s.Rewards,
		Status:                      s.Status,
		Assets:                      clonedAssets,
	}
}
