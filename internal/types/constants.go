package types

const KiraVersion = "v0.0.50"

const (
	ValidatorAccountName = "validator"
	GenesisFileName      = "genesis.json"
)

// constants of validatorStatus
// warning validator status from sekaid comes in uper case format
const (
	Active   = "active"
	Paused   = "paused"
	Inactive = "inactive"
	Waiting  = "waiting"
	Jailed   = "jailed"
)
