package types

// Transaction structs for `sekaid query tx`

type TxData struct {
	Height    string `json:"height"`
	Txhash    string `json:"txhash"`
	Code      int    `json:"code"`
	Data      string `json:"data"`
	RawLog    string `json:"raw_log"`
	Logs      []any  `json:"logs"` // no data provided, assuming a slice of empty interface
	Info      string `json:"info"`
	GasWanted string `json:"gas_wanted"`
	GasUsed   string `json:"gas_used"`
	Tx        any    `json:"tx"` // no data provided, using an empty interface
	Timestamp string `json:"timestamp"`
	Events    []any  `json:"events"` // no data provided, assuming a slice of empty interface
}
