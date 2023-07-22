package types

import "time"

// ResponseInterxStatus represents the JSON response structure for the InterxStatus API.
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
