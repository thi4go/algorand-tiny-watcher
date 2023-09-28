package algorand

import (
	"algorand-tiny-watcher/model"
	"context"
	"log"
	"strings"

	"github.com/algorand/go-algorand-sdk/client/v2/algod"
)

type Algorand struct {
	ctx    context.Context
	Client *algod.Client
}

func NewAlgorandClient(providerUrl string) *Algorand {
	algodClient, err := algod.MakeClient(
		providerUrl,
		strings.Repeat("a", 64),
	)

	if err != nil {
		log.Fatalf("Failed to create algod client: %s", err)
	}

	return &Algorand{
		ctx:    context.Background(),
		Client: algodClient,
	}
}

func (algo *Algorand) IsAddressValid(address string) bool {
	_, err := algo.Client.AccountInformation(address).Do(algo.ctx)
	return err == nil
}

func (algo *Algorand) AccountState(address string) (*model.AccountState, error) {
	res, err := algo.Client.AccountInformation(address).Do(algo.ctx)
	if err != nil {
		return nil, err
	}

	var assets []model.Asset
	for _, asset := range res.Assets {

		res2, err := algo.Client.GetAssetByID(asset.AssetId).Do(algo.ctx)
		if err != nil {
			return nil, err
		}

		assets = append(assets, model.Asset{
			AssetID: asset.AssetId,
			Name:    res2.Params.Name,
			Amount:  asset.Amount,
		})

	}

	accountState := &model.AccountState{
		Address:                     address,
		Amount:                      res.Amount,
		AmountWithoutPendingRewards: res.AmountWithoutPendingRewards,
		PendingRewards:              res.PendingRewards,
		Rewards:                     res.Rewards,
		Status:                      res.Status,
		Assets:                      assets,
	}

	return accountState, nil
}
