package types

import (
	"encoding/json"
	"time"
)

// ResponseSekaidStatus represents the JSON response structure for the Sekaid `/status` query.
type ResponseSekaidStatus struct {
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
type ResponseChunkedGenesis struct {
	Result struct {
		Chunk json.Number `json:"chunk"`
		Total json.Number `json:"total"`
		Data  string      `json:"data"`
	} `json:"result"`
}

// ResponseBlock represents the JSON response structure for the Sekaid `/block` query.
type ResponseBlock struct {
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
