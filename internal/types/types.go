package types

const (
	//  no-op permission
	PermZero = iota

	// the permission that allows to Set Permissions to other actors
	PermSetPermissions

	// permission that allows to Claim a validator Seat
	PermClaimValidator

	// permission that allows to Claim a Councilor Seat
	PermClaimCouncilor

	//# permission to create proposals to whitelist account permission
	PermWhitelistAccountPermissionProposal

	//  permission to vote on a proposal to whitelist account permission
	PermVoteWhitelistAccountPermissionProposal

	//  permission to upsert token alias
	PermUpsertTokenAlias

	//  permission to change transaction fees - execution fee and fee range
	PermChangeTxFee

	//  permission to upsert token rates
	PermUpsertTokenRate

	//  permission to add, modify and assign roles
	PermUpsertRole

	//  permission to create a proposal to change the Data Registry
	PermCreateUpsertDataRegistryProposal

	//  permission to vote on a proposal to change the Data Registry
	PermVoteUpsertDataRegistryProposal

	//  permission to create proposals for setting network property
	PermCreateSetNetworkPropertyProposal

	//  permission to vote a proposal to set network property
	PermVoteSetNetworkPropertyProposal

	//  permission to create proposals to upsert token alias
	PermCreateUpsertTokenAliasProposal

	//  permission to vote proposals to upsert token alias
	PermVoteUpsertTokenAliasProposal

	//  permission to create proposals for setting poor network messages
	PermCreateSetPoorNetworkMessagesProposal

	//  permission to vote proposals to set poor network messages
	PermVoteSetPoorNetworkMessagesProposal

	//  permission to create proposals to upsert token rate
	PermCreateUpsertTokenRateProposal

	//  permission to vote propsals to upsert token rate
	PermVoteUpsertTokenRateProposal

	//  permission to create a proposal to unjail a validator
	PermCreateUnjailValidatorProposal

	//  permission to vote a proposal to unjail a validator
	PermVoteUnjailValidatorProposal

	//  permission to create a proposal to create a role
	PermCreateRoleProposal

	//  permission to vote a proposal to create a role
	PermVoteCreateRoleProposal

	//  permission to create a proposal to change blacklist/whitelisted tokens
	PermCreateTokensWhiteBlackChangeProposal

	//  permission to vote a proposal to change blacklist/whitelisted tokens
	PermVoteTokensWhiteBlackChangeProposal

	//  permission needed to create a proposal to reset whole validator rank
	PermCreateResetWholeValidatorRankProposal

	//  permission needed to vote on reset whole validator rank proposal
	PermVoteResetWholeValidatorRankProposal

	//  permission needed to create a proposal for software upgrade
	PermCreateSoftwareUpgradeProposal

	//  permission needed to vote on software upgrade proposal
	PermVoteSoftwareUpgradeProposal

	//  permission that allows to Set ClaimValidatorPermission to other actors
	PermSetClaimValidatorPermission

	//  permission needed to create a proposal to set proposal duration
	PermCreateSetProposalDurationProposal

	//  permission needed to vote a proposal to set proposal duration
	PermVoteSetProposalDurationProposal

	//  permission needed to create proposals for blacklisting an account permission.
	PermBlacklistAccountPermissionProposal

	//  permission that an actor must have in order to vote a Proposal to blacklist account permission.
	PermVoteBlacklistAccountPermissionProposal

	//  permission needed to create proposals for removing whitelisted permission from an account.
	PermRemoveWhitelistedAccountPermissionProposal

	//  permission that an actor must have in order to vote a proposal to remove a whitelisted account permission
	PermVoteRemoveWhitelistedAccountPermissionProposal

	//  permission needed to create proposals for removing blacklisted permission from an account.
	PermRemoveBlacklistedAccountPermissionProposal

	//  permission that an actor must have in order to vote a proposal to remove a blacklisted account permission.
	PermVoteRemoveBlacklistedAccountPermissionProposal

	//  permission needed to create proposals for whitelisting an role permission.
	PermWhitelistRolePermissionProposal

	// permission that an actor must have in order to vote a proposal to whitelist role permission.
	PermVoteWhitelistRolePermissionProposal

	// permission needed to create proposals for blacklisting an role permission.
	PermBlacklistRolePermissionProposal

	//  permission that an actor must have in order to vote a proposal to blacklist role permission.
	PermVoteBlacklistRolePermissionProposal

	//  permission needed to create proposals for removing whitelisted permission from a role.
	PermRemoveWhitelistedRolePermissionProposal

	//  permission that an actor must have in order to vote a proposal to remove a whitelisted role permission.
	PermVoteRemoveWhitelistedRolePermissionProposal

	//  permission needed to create proposals for removing blacklisted permission from a role.
	PermRemoveBlacklistedRolePermissionProposal

	//  permission that an actor must have in order to vote a proposal to remove a blacklisted role permission.
	PermVoteRemoveBlacklistedRolePermissionProposal

	//  permission needed to create proposals to assign role to an account
	PermAssignRoleToAccountProposal

	//  permission that an actor must have in order to vote a proposal to assign role to an account
	PermVoteAssignRoleToAccountProposal

	//  permission needed to create proposals to unassign role from an account
	PermUnassignRoleFromAccountProposal

	//  permission that an actor must have in order to vote a proposal to unassign role from an account
	PermVoteUnassignRoleFromAccountProposal

	//  permission needed to create a proposal to remove a role.
	PermRemoveRoleProposal

	//  permission needed to vote a proposal to remove a role.
	PermVoteRemoveRoleProposal

	//  permission needed to create proposals to upsert ubi
	PermCreateUpsertUBIProposal

	//  permission that an actor must have in order to vote a proposal to upsert ubi
	PermVoteUpsertUBIProposal

	//  permission needed to create a proposal to remove ubi.
	PermCreateRemoveUBIProposal

	//  permission needed to vote a proposal to remove ubi.
	PermVoteRemoveUBIProposal

	//  permission needed to create a proposal to slash validator.
	PermCreateSlashValidatorProposal

	//  permission needed to vote a proposal to slash validator.
	PermVoteSlashValidatorProposal

	//  permission needed to create a proposal related to basket.
	PermCreateBasketProposal

	//  permission needed to vote a proposal related to basket.
	PermVoteBasketProposal

	//  permission needed to handle emergency issues on basket.
	PermHandleBasketEmergency

	//  permission needed to create a proposal to reset whole councilor rank
	PermCreateResetWholeCouncilorRankProposal

	//  permission needed to vote on reset whole councilor rank proposal
	PermVoteResetWholeCouncilorRankProposal

	//  permission needed to create a proposal to jail councilors
	PermCreateJailCouncilorProposal

	//  permission needed to vote on jail councilors proposal
	PermVoteJailCouncilorProposal
)

type Test struct {
	Text string
}

type SekaidKey struct {
	Address string `yaml:"address"`
}

type AddressPermissions struct {
	BlackList []int `json:"blacklist"`
	WhiteList []int `json:"whitelist"`
}
