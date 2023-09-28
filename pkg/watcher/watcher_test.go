package watcher

import (
	"algorand-tiny-watcher/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockAlgorandProvider struct{}

func (m *MockAlgorandProvider) AccountState(address string) (*model.AccountState, error) {
	return &model.AccountState{
		Address:                     address,
		Amount:                      100000,
		AmountWithoutPendingRewards: 95000,
		PendingRewards:              5000,
		Rewards:                     2500,
		Status:                      "offline",
		Assets: []model.Asset{
			{
				AssetID: 123123,
				Name:    "Lamp",
				Amount:  2500,
			},
		},
	}, nil
}

func (m *MockAlgorandProvider) IsAddressValid(address string) bool {
	return len(address) == 58
}

var testAddress = "ZIFNC42KV67E2ZJVEYRER6J4LVHLMJTR22G3YMW54SLQ3CZDIDURDAOMR4"

func TestWatcher(t *testing.T) {
	provider := &MockAlgorandProvider{}

	t.Run("add and remove address", func(t *testing.T) {
		w := &Watcher{
			provider:  provider,
			addresses: make(map[string]*model.AccountState),
		}

		err := w.AddAddress(testAddress)
		assert.NoError(t, err)
		assert.Contains(t, w.GetState(), testAddress)

		err = w.RemoveAddress(testAddress)
		assert.NoError(t, err)
		assert.NotContains(t, w.GetState(), testAddress)
	})

	t.Run("is address valid", func(t *testing.T) {
		w := &Watcher{provider: provider}
		assert.True(t, w.IsAddressValid(testAddress))
		assert.False(t, w.IsAddressValid("invalidAddress"))
	})

	t.Run("runs watch and updates state", func(t *testing.T) {
		addresses := make(map[string]*model.AccountState, 1)
		addresses[testAddress] = nil

		w := &Watcher{
			provider:  provider,
			addresses: addresses,
		}

		w.watch()

		assert.NotNil(t, w.addresses[testAddress])
	})

}
