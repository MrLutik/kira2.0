package types

import (
	"encoding/json"
	"time"
)

type (
	// ResponseSekaidStatus represents the JSON response structure for the Sekaid `/status` query.
	ResponseSekaidStatus struct {
		Result struct {
			NodeInfo struct {
				ID      string `json:"id"`
				Network string `json:"network"`
			} `json:"node_info"`
			SyncInfo struct {
				LatestBlockHeight string    `json:"latest_block_height"`
				LatestBlockTime   time.Time `json:"latest_block_time"`
				CatchingUp        bool      `json:"catching_up"`
			} `json:"sync_info"`
		} `json:"result"`
	}

	// ResponseChunkedGenesis represents the JSON response structure for the Sekaid `/chunked_genesis` query.
	ResponseChunkedGenesis struct {
		Result struct {
			Chunk json.Number `json:"chunk"`
			Total json.Number `json:"total"`
			Data  string      `json:"data"`
		} `json:"result"`
	}

	// ResponseBlock represents the JSON response structure for the Sekaid `/block` query.
	ResponseBlock struct {
		Result struct {
			BlockID struct {
				Hash string `json:"hash"`
			} `json:"block_id"`
			Block struct {
				Header struct {
					Height string `json:"height"`
				} `json:"header"`
			} `json:"block"`
		} `json:"result"`
	}

	// ValidatorStatus represents output of json structure from command:
	// sekaid query customstaking validator --addr=kira19p8h9kwvrwgeu80c89ctvhwx7w3fc7r7rh32an --output json
	ValidatorStatus struct {
		ValKey string                `json:"val_key"`
		PubKey validatorStatusPubKey `json:"pub_key"`
		Status string                `json:"status"`
		Rank   string                `json:"rank"`
		Streak string                `json:"streak"`
	}

	validatorStatusPubKey struct {
		Type string `json:"@type"`
		Key  string `json:"key"`
	}
)
