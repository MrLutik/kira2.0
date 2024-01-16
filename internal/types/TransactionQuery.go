package types

// TxData represents transaction structs for `sekaid query tx`
type TxData struct {
	Height    string `json:"height"`
	Txhash    string `json:"txhash"`
	Code      int    `json:"code"`
	Data      string `json:"data"`
	RawLog    string `json:"raw_log"`
	Logs      []any  `json:"logs"` // No data provided, assuming a slice of empty interface
	Info      string `json:"info"`
	GasWanted string `json:"gas_wanted"`
	GasUsed   string `json:"gas_used"`
	Tx        any    `json:"tx"` // No data provided, using an empty interface
	Timestamp string `json:"timestamp"`
	Events    []any  `json:"events"` // No data provided, assuming a slice of empty interface
}
