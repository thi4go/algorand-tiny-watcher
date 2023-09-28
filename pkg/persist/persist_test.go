package persist

import (
	"algorand-tiny-watcher/model"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveAndLoadAddresses(t *testing.T) {
	addresses := map[string]*model.AccountState{
		"addr1": {
			Address: "addr1",
			Amount:  1000,
		},
		"addr2": {
			Address: "addr2",
			Amount:  500,
		},
	}

	defer os.Remove(addressesFile)

	t.Run("Save and Load addresses", func(t *testing.T) {
		err := SaveAddresses(addresses)
		assert.NoError(t, err)

		loadedAddresses, err := LoadAddresses()
		assert.NoError(t, err)
		assert.Equal(t, addresses, loadedAddresses, "Loaded addresses do not match saved ones")
	})

	t.Run("Load from non-existent file", func(t *testing.T) {
		os.Remove(addressesFile) // Ensure the file does not exist
		loadedAddresses, err := LoadAddresses()
		assert.NoError(t, err, "Error loading from a non-existent file")
		assert.Nil(t, loadedAddresses, "Loaded addresses should be nil for a non-existent file")
	})

	t.Run("Error handling", func(t *testing.T) {
		// This test checks that we handle errors properly, such as having a corrupted file.

		// Create a corrupted file
		err := ioutil.WriteFile(addressesFile, []byte("corrupted data"), 0644)
		assert.NoError(t, err, "Failed to create a corrupted file")

		loadedAddresses, err := LoadAddresses()
		assert.Error(t, err, "Expected an error loading from a corrupted file")
		assert.Nil(t, loadedAddresses, "Loaded addresses should be nil when reading a corrupted file")
	})
}
