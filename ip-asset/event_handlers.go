package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"story-zrunner-demo/dao"

	"github.com/Zettablock/zsource/dao/ethereum"
	"github.com/Zettablock/zsource/utils"
	"gorm.io/gorm"
)

func HandlerIPRegistered(log ethereum.Log, deps *utils.Deps) (bool, error) {
	blockNumber := log.BlockNumber
	blockTime := log.BlockTime

	chainID, err := strconv.ParseInt(log.ArgumentValues[2], 0, 64)
	if err != nil {
		return false, err
	}

	tokenContract := log.ArgumentValues[3]
	uri := log.ArgumentValues[5]

	ipAsset := &dao.IPAsset{
		BlockNumber:   blockNumber,
		BlockTime:     blockTime.Unix(),
		ID:            log.ArgumentValues[0],
		IPID:          log.ArgumentValues[1],
		ChainID:       chainID,
		TokenContract: tokenContract,
		Metadata:      nil,
		ChildIPIDs:    nil,
		ParentIPIDs:   nil,
		RootIPIDs:     nil,
		NftName:       log.ArgumentValues[4],
		NftTokenURI:   uri,
		NftImageURL:   "",
	}
	

	if uri != "" {
		response, err := http.Get(uri)
		if err != nil {
			deps.Logger.Error("Error fetching image from %s: %s", uri, err)
			return false, err
		}

		if response.StatusCode == 200 {
			jsonResponse := make(map[string]interface{})
			err := json.NewDecoder(response.Body).Decode(&jsonResponse)
			if err != nil {
				deps.Logger.Error("Error decoding response body: %v", err)
				return false, err
			}

			if image, ok := jsonResponse["image"]; ok {
				ipAsset.NftImageURL = image.(string)
			}
		}
	}

	if err := deps.DestinationDB.Save(ipAsset).Error; err != nil {
		return false, err
	}

	collection := &dao.Collection{}
	if err := deps.DestinationDB.Where("id = ?", tokenContract).Take(collection).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	if collection.ID == "" {
		collection = &dao.Collection{
			ID:                    tokenContract,
			AssetCount:            1,
			RaisedDisputeCount:    0,
			CancelledDisputeCount: 0,
			ResolvedDisputeCount:  0,
			JudgedDisputeCount:    0,
			LicensesCount:         0,
		}
	} else {
		collection.AssetCount++
		collection.BlockNumber = blockNumber
		collection.BlockTime = blockTime.Unix()
	}
	if err = deps.DestinationDB.Save(collection).Error; err != nil {
		return false, err
	}

	txn := &ethereum.Transaction{}
	if err = deps.SourceDB.Where("hash = ?", log.TransactionHash).Take(txn).Error; err != nil {
		return false, err
	}

	daoTx := &dao.Transaction{
		BlockNumber:      blockNumber,
		BlockTime:        blockTime.Unix(),
		ID:               log.TransactionHash,
		TxHash:           log.TransactionHash,
		TransactionIndex: log.TransactionIndex,
		LogIndex:         log.LogIndex,
		Initiator:        txn.FromAddress,
		CreatedAt:        blockTime.Unix(),
		ResourceID:       log.ContractAddress,
		IPID:             log.ArgumentValues[1],
		ActionType:       "Register",
		ResourceType:     "IPAsset",
	}

	if err = deps.DestinationDB.Save(daoTx).Error; err != nil {
		return false, err
	}
	return false, nil
}
