package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasChanged(t *testing.T) {
	baseState := &AccountState{
		Amount: 1000,
		Assets: []Asset{
			{AssetID: 1, Name: "Asset1", Amount: 500},
			{AssetID: 2, Name: "Asset2", Amount: 300},
		},
	}

	t.Run("identical states", func(t *testing.T) {
		identicalState := &AccountState{
			Amount: 1000,
			Assets: []Asset{
				{AssetID: 1, Name: "Asset1", Amount: 500},
				{AssetID: 2, Name: "Asset2", Amount: 300},
			},
		}
		assert.False(t, baseState.HasChanged(identicalState))
	})

	t.Run("different amounts", func(t *testing.T) {
		diffAmountState := &AccountState{
			Amount: 999,
			Assets: []Asset{
				{AssetID: 1, Name: "Asset1", Amount: 500},
				{AssetID: 2, Name: "Asset2", Amount: 300},
			},
		}
		assert.True(t, baseState.HasChanged(diffAmountState))
	})

	t.Run("different assets", func(t *testing.T) {
		diffAssetsState := &AccountState{
			Amount: 1000,
			Assets: []Asset{
				{AssetID: 1, Name: "Asset1Modified", Amount: 500},
				{AssetID: 2, Name: "Asset2", Amount: 300},
			},
		}
		assert.True(t, baseState.HasChanged(diffAssetsState))
	})

	t.Run("additional asset in the new state", func(t *testing.T) {
		additionalAssetState := &AccountState{
			Amount: 1000,
			Assets: []Asset{
				{AssetID: 1, Name: "Asset1", Amount: 500},
				{AssetID: 2, Name: "Asset2", Amount: 300},
				{AssetID: 3, Name: "Asset3", Amount: 400},
			},
		}
		assert.True(t, baseState.HasChanged(additionalAssetState))
	})

	t.Run("missing asset in the new state", func(t *testing.T) {
		missingAssetState := &AccountState{
			Amount: 1000,
			Assets: []Asset{
				{AssetID: 2, Name: "Asset2", Amount: 300},
			},
		}
		assert.True(t, baseState.HasChanged(missingAssetState))
	})
}

func TestAccountStateClone(t *testing.T) {
	orig := &AccountState{
		Address:                     "TestAddress",
		Amount:                      1000,
		AmountWithoutPendingRewards: 900,
		PendingRewards:              50,
		Rewards:                     150,
		Status:                      "Online",
		Assets: []Asset{
			{AssetID: 1, Name: "Asset1", Amount: 500},
			{AssetID: 2, Name: "Asset2", Amount: 300},
		},
	}

	clone := orig.Clone()

	t.Run("ensure values match", func(t *testing.T) {
		assert.Equal(t, orig.Address, clone.Address)
		assert.Equal(t, orig.Amount, clone.Amount)
		assert.Equal(t, orig.AmountWithoutPendingRewards, clone.AmountWithoutPendingRewards)
		assert.Equal(t, orig.PendingRewards, clone.PendingRewards)
		assert.Equal(t, orig.Rewards, clone.Rewards)
		assert.Equal(t, orig.Status, clone.Status)
		assert.Equal(t, orig.Assets, clone.Assets)
	})

	t.Run("modify clone value and ensure original is not modified", func(t *testing.T) {
		clone.Address = "ModifiedAddress"
		clone.Amount = 999
		clone.Assets[0].Name = "ModifiedAssetName"
		clone.Assets[0].Amount = 999

		assert.NotEqual(t, orig.Address, clone.Address)
		assert.NotEqual(t, orig.Amount, clone.Amount)
		assert.NotEqual(t, orig.Assets[0].Name, clone.Assets[0].Name)
		assert.NotEqual(t, orig.Assets[0].Amount, clone.Assets[0].Amount)
	})
}
