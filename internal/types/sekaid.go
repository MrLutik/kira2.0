package types

import "time"

// ResponseSekaidStatus represents the JSON response structure for the Sekaid status query.
type ResponseSekaidStatus struct {
	Result struct {
		NodeInfo struct {
			ID string `json:"id"`
		} `json:"node_info"`
		SyncInfo struct {
			LatestBlockHeight string    `json:"latest_block_height"`
			LatestBlockTime   time.Time `json:"latest_block_time"`
			CatchingUp        bool      `json:"catching_up"`
		} `json:"sync_info"`
	} `json:"result"`
}
