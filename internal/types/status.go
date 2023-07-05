package types

// curl http://168.119.228.165:36657/status
// {
// 	"jsonrpc": "2.0",
// 	"id": -1,
// 	"result": {
// 	  "node_info": {
// 	    "protocol_version": {
// 		 "p2p": "8",
// 		 "block": "11",
// 		 "app": "0"
// 	    },
// 	    "id": "213267558cc0febde6f1b8149f92fecf16d48b38",
// 	    "listen_addr": "tcp://168.119.228.165:36656",
// 	    "network": "localnet-4",
// 	    "version": "0.34.22",
// 	    "channels": "40202122233038606100",
// 	    "moniker": "KIRA VALIDATOR NODE",
// 	    "other": {
// 		 "tx_index": "on",
// 		 "rpc_address": "tcp://0.0.0.0:26657"
// 	    }
// 	  },
// 	  "sync_info": {
// 	    "latest_block_hash": "296BFEC53E828B5265590CBB2062F1EA7506D76F1EF7E27152A0056207348E64",
// 	    "latest_app_hash": "D76F5BCC13CFBED74793BB6BC786F93E8C132807EA0ED2DFE1CF9448DC3B982E",
// 	    "latest_block_height": "334267",
// 	    "latest_block_time": "2023-07-05T11:14:09.02975714Z",
// 	    "earliest_block_hash": "26AD6395E9AB39AE545D53F64B2DFF4C664F55448A5A96F9963336386CFE2B3C",
// 	    "earliest_app_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
// 	    "earliest_block_height": "1",
// 	    "earliest_block_time": "2023-05-25T12:56:59.153143882Z",
// 	    "catching_up": false
// 	  },
// 	  "validator_info": {
// 	    "address": "AFC8EBD65CE1E7DD38E1E4DD514E9B03A0085E98",
// 	    "pub_key": {
// 		 "type": "tendermint/PubKeyEd25519",
// 		 "value": "4q2MwShsxSt2UBXmEv/rLXevigs6uln9J74uxA/XbRE="
// 	    },
// 	    "voting_power": "1"
// 	  }
// 	}
//    }
type Status struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  Result `json:"result"`
}

type Result struct {
	NodeInfo      NodeInfo      `json:"node_info"`
	SyncInfo      SyncInfo      `json:"sync_info"`
	ValidatorInfo ValidatorInfo `json:"validator_info"`
}

type NodeInfo struct {
	ProtocolVersion ProtocolVersion `json:"protocol_version"`
	ID              string          `json:"id"`
	ListenAddr      string          `json:"listen_addr"`
	Network         string          `json:"network"`
	Version         string          `json:"version"`
	Channels        string          `json:"channels"`
	Moniker         string          `json:"moniker"`
	Other           Other           `json:"other"`
}

type Other struct {
	TxIndex    string `json:"tx_index"`
	RPCAddress string `json:"rpc_address"`
}

type ProtocolVersion struct {
	P2P   string `json:"p2p"`
	Block string `json:"block"`
	App   string `json:"app"`
}

type SyncInfo struct {
	LatestBlockHash     string `json:"latest_block_hash"`
	LatestAppHash       string `json:"latest_app_hash"`
	LatestBlockHeight   string `json:"latest_block_height"`
	LatestBlockTime     string `json:"latest_block_time"`
	EarliestBlockHash   string `json:"earliest_block_hash"`
	EarliestAppHash     string `json:"earliest_app_hash"`
	EarliestBlockHeight string `json:"earliest_block_height"`
	EarliestBlockTime   string `json:"earliest_block_time"`
	CatchingUp          bool   `json:"catching_up"`
}

type ValidatorInfo struct {
	Address     string `json:"address"`
	PubKey      PubKey `json:"pub_key"`
	VotingPower string `json:"voting_power"`
}
type PubKey struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
