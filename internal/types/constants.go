package types

const (
	KiraVersion = "v0.0.50"

	ValidatorAccountName = "validator"
	GenesisFileName      = "genesis.json"

	// constants of validatorStatus
	// warning validator status from sekaid comes in upper case format
	Active   = "active"
	Paused   = "paused"
	Inactive = "inactive"
	Waiting  = "waiting"
	Jailed   = "jailed"
)
