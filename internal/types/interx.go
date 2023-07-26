package types

import "time"

// ResponseInterxStatus represents the JSON response structure for the Interx `/api/status` query.
type ResponseInterxStatus struct {
	NodeInfo struct {
		ID string `json:"id"`
	} `json:"node_info"`
	SyncInfo struct {
		LatestBlockHeight string    `json:"latest_block_height"`
		LatestBlockTime   time.Time `json:"latest_block_time"`
		CatchingUp        bool      `json:"catching_up"`
	} `json:"sync_info"`
}

// ResponseCheckSum represents the JSON response structure for the Interx `/api/gensum` query.
type ResponseCheckSum struct {
	Checksum string `json:"checksum"`
}
