package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"story-zrunner-demo/dao"

	"github.com/Zettablock/zsource/dao/ethereum"
	"github.com/Zettablock/zsource/utils"
)

func HandlerIPRegistered(log ethereum.Log, deps *utils.Deps) (bool, error) {
	blockNumber := log.BlockNumber
	blockTime := log.BlockTime

	chainID, err := strconv.ParseInt(log.ArgumentValues[2], 0, 64)
	if err != nil {
		return false, err
	}

	ipAsset := &dao.IPAsset{
		BlockNumber:   blockNumber,
		BlockTime:     blockTime.Unix(),
		ID:            log.ArgumentValues[0],
		IPID:          log.ArgumentValues[1],
		ChainID:       chainID,
		TokenContract: log.ArgumentValues[3],
		Metadata:      nil,
		ChildIPIDs:    nil,
		ParentIPIDs:   nil,
		RootIPIDs:     nil,
		NftName:       log.ArgumentValues[4],
		NftTokenURI:   log.ArgumentValues[5],
		NftImageURL:   "",
	}

	// try:
	//     if event.uri and len(event.uri) > 0:
	//         response = requests.get(event.uri)
	//         if response.status_code == 200:
	//             json_response = response.json()
	//             if "image" in json_response:
	//                 ip_asset["nft_image_url"] = json_response["image"]
	// except Exception as e:
	//     logging.error(f"Error fetching image from {event.uri}: {e}")

	if log.ArgumentValues[5] != "" {
		response, err := http.Get(log.ArgumentValues[5])
		if err != nil {
			deps.Logger.Error("Error fetching image from %s: %s", log.ArgumentValues[5], err)
			return false, err
		}

		if response.StatusCode == 200 {
			jsonResponse := make(map[string]interface{})
			err := json.NewDecoder(response.Body).Decode(&jsonResponse)
			if err != nil {
				deps.Logger.Error("Error decoding response body: %s", err)
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

	
}
