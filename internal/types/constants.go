package types

const (
	KiraVersion = "v0.0.50"

	ValidatorAccountName = "validator"
	GenesisFileName      = "genesis.json"

	// Constants of validatorStatus
	// Warning validator status from sekaid comes in upper case format
	Active   = "active"
	Paused   = "paused"
	Inactive = "inactive"
	Waiting  = "waiting"
	Jailed   = "jailed"
)
