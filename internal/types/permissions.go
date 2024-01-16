package types

const (
	// No-op permission
	PermZero = iota

	// The permission that allows to Set Permissions to other actors
	PermSetPermissions

	// Permission that allows to Claim a validator Seat
	PermClaimValidator

	// Permission that allows to Claim a Councilor Seat
	PermClaimCouncilor

	// # permission to create proposals to whitelist account permission
	PermWhitelistAccountPermissionProposal

	// Permission to vote on a proposal to whitelist account permission
	PermVoteWhitelistAccountPermissionProposal

	// Permission to upsert token alias
	PermUpsertTokenAlias

	// Permission to change transaction fees - execution fee and fee range
	PermChangeTxFee

	// Permission to upsert token rates
	PermUpsertTokenRate

	// Permission to add, modify and assign roles
	PermUpsertRole

	// Permission to create a proposal to change the Data Registry
	PermCreateUpsertDataRegistryProposal

	// Permission to vote on a proposal to change the Data Registry
	PermVoteUpsertDataRegistryProposal

	// Permission to create proposals for setting network property
	PermCreateSetNetworkPropertyProposal

	// Permission to vote a proposal to set network property
	PermVoteSetNetworkPropertyProposal

	// Permission to create proposals to upsert token alias
	PermCreateUpsertTokenAliasProposal

	// Permission to vote proposals to upsert token alias
	PermVoteUpsertTokenAliasProposal

	// Permission to create proposals for setting poor network messages
	PermCreateSetPoorNetworkMessagesProposal

	// Permission to vote proposals to set poor network messages
	PermVoteSetPoorNetworkMessagesProposal

	// Permission to create proposals to upsert token rate
	PermCreateUpsertTokenRateProposal

	// Permission to vote propsals to upsert token rate
	PermVoteUpsertTokenRateProposal

	// Permission to create a proposal to unjail a validator
	PermCreateUnjailValidatorProposal

	// Permission to vote a proposal to unjail a validator
	PermVoteUnjailValidatorProposal

	// Permission to create a proposal to create a role
	PermCreateRoleProposal

	// Permission to vote a proposal to create a role
	PermVoteCreateRoleProposal

	// Permission to create a proposal to change blacklist/whitelisted tokens
	PermCreateTokensWhiteBlackChangeProposal

	// Permission to vote a proposal to change blacklist/whitelisted tokens
	PermVoteTokensWhiteBlackChangeProposal

	// Permission needed to create a proposal to reset whole validator rank
	PermCreateResetWholeValidatorRankProposal

	// Permission needed to vote on reset whole validator rank proposal
	PermVoteResetWholeValidatorRankProposal

	// Permission needed to create a proposal for software upgrade
	PermCreateSoftwareUpgradeProposal

	// Permission needed to vote on software upgrade proposal
	PermVoteSoftwareUpgradeProposal

	// Permission that allows to Set ClaimValidatorPermission to other actors
	PermSetClaimValidatorPermission

	// Permission needed to create a proposal to set proposal duration
	PermCreateSetProposalDurationProposal

	// Permission needed to vote a proposal to set proposal duration
	PermVoteSetProposalDurationProposal

	// Permission needed to create proposals for blacklisting an account permission.
	PermBlacklistAccountPermissionProposal

	// Permission that an actor must have in order to vote a Proposal to blacklist account permission.
	PermVoteBlacklistAccountPermissionProposal

	// Permission needed to create proposals for removing whitelisted permission from an account.
	PermRemoveWhitelistedAccountPermissionProposal

	// Permission that an actor must have in order to vote a proposal to remove a whitelisted account permission
	PermVoteRemoveWhitelistedAccountPermissionProposal

	// Permission needed to create proposals for removing blacklisted permission from an account.
	PermRemoveBlacklistedAccountPermissionProposal

	// Permission that an actor must have in order to vote a proposal to remove a blacklisted account permission.
	PermVoteRemoveBlacklistedAccountPermissionProposal

	// Permission needed to create proposals for whitelisting an role permission.
	PermWhitelistRolePermissionProposal

	// Permission that an actor must have in order to vote a proposal to whitelist role permission.
	PermVoteWhitelistRolePermissionProposal

	// Permission needed to create proposals for blacklisting an role permission.
	PermBlacklistRolePermissionProposal

	// Permission that an actor must have in order to vote a proposal to blacklist role permission.
	PermVoteBlacklistRolePermissionProposal

	// Permission needed to create proposals for removing whitelisted permission from a role.
	PermRemoveWhitelistedRolePermissionProposal

	// Permission that an actor must have in order to vote a proposal to remove a whitelisted role permission.
	PermVoteRemoveWhitelistedRolePermissionProposal

	// Permission needed to create proposals for removing blacklisted permission from a role.
	PermRemoveBlacklistedRolePermissionProposal

	// Permission that an actor must have in order to vote a proposal to remove a blacklisted role permission.
	PermVoteRemoveBlacklistedRolePermissionProposal

	// Permission needed to create proposals to assign role to an account
	PermAssignRoleToAccountProposal

	// Permission that an actor must have in order to vote a proposal to assign role to an account
	PermVoteAssignRoleToAccountProposal

	// Permission needed to create proposals to unassign role from an account
	PermUnassignRoleFromAccountProposal

	// Permission that an actor must have in order to vote a proposal to unassign role from an account
	PermVoteUnassignRoleFromAccountProposal

	// Permission needed to create a proposal to remove a role.
	PermRemoveRoleProposal

	// Permission needed to vote a proposal to remove a role.
	PermVoteRemoveRoleProposal

	// Permission needed to create proposals to upsert ubi
	PermCreateUpsertUBIProposal

	// Permission that an actor must have in order to vote a proposal to upsert ubi
	PermVoteUpsertUBIProposal

	// Permission needed to create a proposal to remove ubi.
	PermCreateRemoveUBIProposal

	// Permission needed to vote a proposal to remove ubi.
	PermVoteRemoveUBIProposal

	// Permission needed to create a proposal to slash validator.
	PermCreateSlashValidatorProposal

	// Permission needed to vote a proposal to slash validator.
	PermVoteSlashValidatorProposal

	// Permission needed to create a proposal related to basket.
	PermCreateBasketProposal

	// Permission needed to vote a proposal related to basket.
	PermVoteBasketProposal

	// Permission needed to handle emergency issues on basket.
	PermHandleBasketEmergency

	// Permission needed to create a proposal to reset whole councilor rank
	PermCreateResetWholeCouncilorRankProposal

	// Permission needed to vote on reset whole councilor rank proposal
	PermVoteResetWholeCouncilorRankProposal

	// Permission needed to create a proposal to jail councilors
	PermCreateJailCouncilorProposal

	// Permission needed to vote on jail councilors proposal
	PermVoteJailCouncilorProposal
)
