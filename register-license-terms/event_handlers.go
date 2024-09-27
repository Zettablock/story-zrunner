package main

import (
	"context"
	"math/big"
	"strconv"

	"story-zrunner-demo/dao"

	"github.com/Zettablock/zsource/dao/ethereum"
	"github.com/Zettablock/zsource/utils"

	geth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const ABI = "pil_license_template.json"

func HandleLicenseTemplateRegistered(log ethereum.Log, deps *utils.Deps) (bool, error) {
	address := log.ArgumentValues[0]

	client, err := ethclient.Dial("YOUR_RPC_URL")
	if err != nil {
		return false, err
	}

	name, err := getName(client, address, deps)
	if err != nil {
		return false, err
	}

	metadataURI, err := getMetadataURI(client, address, deps)
	if err != nil {
		return false, err
	}

	licenseTemplate := dao.LicenseTemplate{
		BlockNumber: log.BlockNumber,
		BlockTime:   log.BlockTime.Unix(),
		ID:          address,
		Name:        name,
		MetadataURI: metadataURI,
	}

	deps.DestinationDB.Save(&licenseTemplate)

	err = deps.SaveTemplate("LicenseTemplate", address)
	if err != nil {
		return false, err
	}

	txn := &ethereum.Transaction{}
	if err = deps.SourceDB.Where("hash = ?", log.TransactionHash).Take(txn).Error; err != nil {
		return false, err
	}

	daoTx := &dao.Transaction{
		BlockNumber:      log.BlockNumber,
		BlockTime:        log.BlockTime.Unix(),
		ID:               log.TransactionHash,
		TxHash:           log.TransactionHash,
		TransactionIndex: log.TransactionIndex,
		LogIndex:         log.LogIndex,
		Initiator:        txn.FromAddress,
		CreatedAt:        log.BlockTime.Unix(),
		ResourceID:       log.ContractAddress,
		IPID:             "",
		ActionType:       "Register",
		ResourceType:     "LicenseTemplate",
	}

	if err = deps.DestinationDB.Save(daoTx).Error; err != nil {
		return false, err
	}

	return false, nil
}

func HandleLicenseTermsRegistered(log ethereum.Log, deps *utils.Deps) (bool, error) {
	client, err := ethclient.Dial("YOUR_RPC_URL")
	if err != nil {
		return false, err
	}

	licenseTermsId := log.ArgumentValues[0]
	licenseTemplate := log.ArgumentValues[1]

	json, err := getJson(client, licenseTemplate, licenseTermsId, deps)
	if err != nil {
		return false, err
	}

	licenseTerm := dao.LicenseTerm{
		BlockNumber:     log.BlockNumber,
		BlockTime:       log.BlockTime.Unix(),
		ID:              licenseTermsId,
		LicenseTemplate: licenseTemplate,
		JSON:            json,
	}

	if err = deps.DestinationDB.Save(&licenseTerm).Error; err != nil {
		return false, err
	}

	txn := &ethereum.Transaction{}
	if err = deps.SourceDB.Where("hash = ?", log.TransactionHash).Take(txn).Error; err != nil {
		return false, err
	}

	daoTx := &dao.Transaction{
		BlockNumber:      log.BlockNumber,
		BlockTime:        log.BlockTime.Unix(),
		ID:               log.TransactionHash,
		TxHash:           log.TransactionHash,
		TransactionIndex: log.TransactionIndex,
		LogIndex:         log.LogIndex,
		Initiator:        txn.FromAddress,
		CreatedAt:        log.BlockTime.Unix(),
		ResourceID:       log.ContractAddress,
		IPID:             "",
		ActionType:       "Register",
		ResourceType:     "LicenseTerm",
	}

	if err = deps.DestinationDB.Save(daoTx).Error; err != nil {
		return false, err
	}
	return false, nil
}

func getMetadataURI(client *ethclient.Client, address string, deps *utils.Deps) (string, error) {
	// Define the contract address and ABI
	contractAddress := common.HexToAddress(address)

	parsedABI, err := deps.LoadABIByName(ABI)
	if err != nil {
		return "", err
	}

	// Prepare the call
	callData, err := parsedABI.Pack("getMetadataURI")
	if err != nil {
		return "", err
	}

	// Call the contract function
	msg := geth.CallMsg{
		To:   &contractAddress,
		Data: callData,
	}

	res, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return "", err
	}

	// Unpack the result
	var result string
	err = parsedABI.UnpackIntoInterface(&result, "getMetadataURI", res)
	if err != nil {
		return "", err
	}

	return result, nil
}

func getName(client *ethclient.Client, address string, deps *utils.Deps) (string, error) {

	// Define the contract address and ABI
	contractAddress := common.HexToAddress(address)

	parsedABI, err := deps.LoadABIByName(ABI)
	if err != nil {
		return "", err
	}
	// Prepare the call
	callData, err := parsedABI.Pack("name")
	if err != nil {
		return "", err
	}

	// Call the contract function
	msg := geth.CallMsg{
		To:   &contractAddress,
		Data: callData,
	}

	res, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return "", err
	}

	// Unpack the result
	var result string
	err = parsedABI.UnpackIntoInterface(&result, "name", res)
	if err != nil {
		return "", err
	}

	return result, nil
}

func getJson(client *ethclient.Client, address string, licenseTermID string, deps *utils.Deps) (string, error) {

	// Define the contract address and ABI
	contractAddress := common.HexToAddress(address)

	parsedABI, err := deps.LoadABIByName(ABI)
	if err != nil {
		return "", err
	}

	id, err := strconv.ParseInt(licenseTermID, 10, 64)
	if err != nil {
		return "", err
	}

	// Prepare the call
	callData, err := parsedABI.Pack("toJson", big.NewInt(id))
	if err != nil {
		return "", err
	}

	// Call the contract function
	msg := geth.CallMsg{
		To:   &contractAddress,
		Data: callData,
	}

	res, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return "", err
	}

	// Unpack the result
	var result string
	err = parsedABI.UnpackIntoInterface(&result, "toJson", res)
	if err != nil {
		return "", err
	}

	return result, nil
}
