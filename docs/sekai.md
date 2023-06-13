# Sekai
`ver.: v0.3.16`

- [Sekai](#sekai)
  - [Context](#context)
    - [1. add-genesis-account](#1-add-genesis-account)
    - [2. collect-gentxs](#2-collect-gentxs)
    - [3. config](#3-config)
    - [4. debug](#4-debug)
      - [4.1 addr](#41-addr)
      - [4.2 pubkey](#42-pubkey)
      - [4.3 raw-bytes](#43-raw-bytes)
    - [5. export](#5-export)
    - [6. export-metadata](#6-export-metadata)
    - [7. export-minimized-genesis](#7-export-minimized-genesis)
    - [8. gentx](#8-gentx)
    - [9. gentx-claim](#9-gentx-claim)
    - [10. help](#10-help)
    - [11. init](#11-init)
    - [12. keys](#12-keys)
      - [12.1 add](#121-add)
      - [12.2 delete](#122-delete)
      - [12.3 export](#123-export)
      - [12.4 import](#124-import)
      - [12.5 list](#125-list)
      - [12.6 migrate](#126-migrate)
      - [12.7 mnemonic](#127-mnemonic)
      - [12.8 parse](#128-parse)
      - [12.9 show](#129-show)
    - [13. new-genesis-from-exported](#13-new-genesis-from-exported)
    - [14. query](#14-query)
      - [14.1 account](#141-account)
      - [14.2 auth](#142-auth)
        - [14.2.1 account](#1421-account)
        - [14.2.2 accounts](#1422-accounts)
        - [14.2.3 module-account](#1423-module-account)
        - [14.2.4 params](#1424-params)
      - [14.3 bank](#143-bank)
        - [14.3.1 balances](#1431-balances)
        - [14.3.2 denom-metadata](#1432-denom-metadata)
        - [14.3.3 total](#1433-total)
      - [14.4 basket](#144-basket)
      - [14.5 block](#145-block)
      - [14.6 collectives](#146-collectives)
      - [14.7 custody](#147-custody)
      - [14.8 customevidence](#148-customevidence)
      - [14.9 customgov](#149-customgov)
        - [14.9.1 all-execution-fees](#1491-all-execution-fees)
        - [14.9.2 all-identity-record-verify-requests](#1492-all-identity-record-verify-requests)
        - [14.9.3 all-proposal-durations](#1493-all-proposal-durations)
        - [14.9.4 all-roles](#1494-all-roles)
        - [14.9.5 blacklisted-permission-addresses](#1495-blacklisted-permission-addresses)
        - [14.9.6 council-registry](#1496-council-registry)
        - [14.9.7 councilors](#1497-councilors)
        - [14.9.8 data-registry](#1498-data-registry)
        - [14.9.9 data-registry-keys](#1499-data-registry-keys)
        - [14.9.10 execution-fee](#14910-execution-fee)
        - [14.9.11 identity-record](#14911-identity-record)
        - [14.9.12 identity-record-verify-request](#14912-identity-record-verify-request)
        - [14.9.13 identity-record-verify-requests-by-approver](#14913-identity-record-verify-requests-by-approver)
        - [14.9.14 identity-record-verify-requests-by-requester](#14914-identity-record-verify-requests-by-requester)
        - [14.9.15 identity-records](#14915-identity-records)
        - [14.9.16 identity-records-by-addr](#14916-identity-records-by-addr)
        - [14.9.17 network-properties](#14917-network-properties)
        - [14.9.18 non-councilors](#14918-non-councilors)
        - [14.9.19 permissions](#14919-permissions)
        - [14.9.20 poll-votes](#14920-poll-votes)
        - [14.9.21 polls](#14921-polls)
        - [14.9.22 poor-network-messages](#14922-poor-network-messages)
        - [14.9.23 proposal](#14923-proposal)
        - [14.9.24 proposal-duration](#14924-proposal-duration)
        - [14.9.25 proposals](#14925-proposals)
        - [14.9.26 proposer\_voters\_count](#14926-proposer_voters_count)
        - [14.9.27 role](#14927-role)
        - [14.9.28 roles](#14928-roles)
        - [14.9.29 vote](#14929-vote)
        - [14.9.30 voters](#14930-voters)
        - [14.9.31 votes](#14931-votes)
        - [14.9.32 whitelisted-permission-addresses](#14932-whitelisted-permission-addresses)
        - [14.9.33 whitelisted-role-addresses](#14933-whitelisted-role-addresses)
      - [14.10 customslashing](#1410-customslashing)
      - [14.11 customstaking](#1411-customstaking)
      - [14.12 distributor](#1412-distributor)
      - [14.13 layer2](#1413-layer2)
      - [14.14 multistaking](#1414-multistaking)
      - [14.15 params](#1415-params)
      - [14.16 recovery](#1416-recovery)
      - [14.17 spending](#1417-spending)
      - [14.18 tendermint-validator-set](#1418-tendermint-validator-set)
      - [14.19 tokens](#1419-tokens)
        - [14.19.1 alias](#14191-alias)
        - [14.19.2 aliases-by-denom](#14192-aliases-by-denom)
        - [14.19.3 all-aliases](#14193-all-aliases)
        - [14.19.4 all-rates](#14194-all-rates)
        - [14.19.5 rate](#14195-rate)
        - [14.19.6 rates-by-denom](#14196-rates-by-denom)
        - [14.19.7 token-black-whites](#14197-token-black-whites)
      - [14.20 tx](#1420-tx)
      - [14.21 txs](#1421-txs)
      - [14.22 ubi](#1422-ubi)
      - [14.23 upgrade](#1423-upgrade)
    - [15. rollback](#15-rollback)
    - [16. rosetta](#16-rosetta)
    - [17. start](#17-start)
    - [18. status](#18-status)
    - [19. tendermint](#19-tendermint)
    - [20. testnet](#20-testnet)
    - [21. tx](#21-tx)
      - [21.1 bank](#211-bank)
        - [21.1.1 send](#2111-send)
      - [21.2 basket](#212-basket)
      - [21.3 broadcast](#213-broadcast)
      - [21.4 collectives](#214-collectives)
      - [21.5 custody](#215-custody)
      - [21.6 customevidence](#216-customevidence)
      - [21.7 customgov](#217-customgov)
        - [21.7.1 cancel-identity-records-verify-request](#2171-cancel-identity-records-verify-request)
        - [21.7.2 councilor](#2172-councilor)
        - [21.7.3 delete-identity-records](#2173-delete-identity-records)
        - [21.7.4 handle-identity-records-verify-request](#2174-handle-identity-records-verify-request)
        - [21.7.5 permission](#2175-permission)
          - [21.7.5.1 blacklist](#21751-blacklist)
          - [21.7.5.2 remove-blacklisted](#21752-remove-blacklisted)
          - [21.7.5.3 remove-whitelisted](#21753-remove-whitelisted)
          - [21.7.5.4 whitelist](#21754-whitelist)
        - [21.7.6 poll](#2176-poll)
        - [21.7.7 proposal](#2177-proposal)
          - [21.7.7.1 account](#21771-account)
          - [21.7.7.1.1 assign-role](#217711-assign-role)
          - [21.7.7.1.2 blacklist-permission](#217712-blacklist-permission)
          - [21.7.7.1.3 remove-blacklisted-permission](#217713-remove-blacklisted-permission)
          - [21.7.7.1.4 remove-whitelisted-permission](#217714-remove-whitelisted-permission)
          - [21.7.7.1.5 unassign-role](#217715-unassign-role)
          - [21.7.7.1.6 whitelist-permission](#217716-whitelist-permission)
          - [21.7.7.2 proposal-jail-councilor](#21772-proposal-jail-councilor)
          - [21.7.7.3 proposal-reset-whole-councilor-rank](#21773-proposal-reset-whole-councilor-rank)
          - [21.7.7.4 role](#21774-role)
          - [21.7.7.4.1 blacklist-permission](#217741-blacklist-permission)
          - [21.7.7.4.2 create](#217742-create)
          - [21.7.7.4.3 remove](#217743-remove)
          - [21.7.7.4.4 remove-blacklisted-permission](#217744-remove-blacklisted-permission)
          - [21.7.7.4.5 remove-whitelisted-permission](#217745-remove-whitelisted-permission)
          - [21.7.7.4.6 whitelist-permission](#217746-whitelist-permission)
          - [21.7.7.5 set-network-property](#21775-set-network-property)
          - [21.7.7.6 set-poor-network-msgs](#21776-set-poor-network-msgs)
          - [21.7.7.7 set-proposal-durations-proposal](#21777-set-proposal-durations-proposal)
          - [21.7.7.8 upsert-data-registry](#21778-upsert-data-registry)
          - [21.7.7.9 vote](#21779-vote)
        - [21.7.8 register-identity-records](#2178-register-identity-records)
        - [21.7.9 request-identity-record-verify](#2179-request-identity-record-verify)
        - [21.7.10 role](#21710-role)
        - [21.7.11 set-execution-fee](#21711-set-execution-fee)
        - [21.7.12 set-network-properties](#21712-set-network-properties)
      - [21.8 customslashing](#218-customslashing)
      - [21.9 customstaking](#219-customstaking)
      - [21.10 decode](#2110-decode)
      - [21.11 distributor](#2111-distributor)
      - [21.12 encode](#2112-encode)
      - [21.13 layer2](#2113-layer2)
      - [21.14 multisign](#2114-multisign)
      - [21.15 multistaking](#2115-multistaking)
      - [21.16 recovery](#2116-recovery)
      - [21.17 sign](#2117-sign)
      - [21.18 sign-batch](#2118-sign-batch)
      - [21.19 spending](#2119-spending)
      - [21.20 tokens](#2120-tokens)
      - [21.21 ubi](#2121-ubi)
      - [21.22 upgrade](#2122-upgrade)
      - [21.23 validate-signatures](#2123-validate-signatures)
    - [22. val-address](#22-val-address)
    - [23. valcons-address](#23-valcons-address)
    - [24. validate-genesis](#24-validate-genesis)
    - [25. version](#25-version)

## Context

### 1. add-genesis-account
Add a genesis account to genesis.json. The provided account must specify the account address or key name and a list of initial coins. If a key name is given, the address will be looked up in the local Keybase. The list of initial tokens must contain valid denominations. Accounts may optionally be supplied with vesting parameters.
Usage:
```
sekaid add-genesis-account [address_or_key_name] [coin][,[coin]] [flags]
```
| Flags                       | Description                                                                                    | Work  |
| --------------------------- | ---------------------------------------------------------------------------------------------- | ----- |
| `--height int`              | Use a specific height to query state at (this can error if the node is pruning state)          | ❌ ?no |
| `-h, --help`                | help for add-genesis-account                                                                   | ✅ yes |
| ` --keyring-backend string` | Select keyring's backend (`os\|file\|kwallet\|pass\|test`) (default `"os"`)                    | ✅ yes |
| `--node string`             | \<host>:\<port> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ❌?no  |
| ` -o, --output string`      | Output format (`text\|json`) (default `"text`")                                                | ❌ no  |
| `--vesting-amount string`   | amount of coins for vesting accounts                                                           | ❌ ?   |
| `--vesting-end-time int`    | schedule end time (unix epoch) for vesting accounts                                            | ❌ ?   |
| `--vesting-start-time int`  | schedule start time (unix epoch) for vesting accounts                                          | ❌ ?   |

| Global flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌?no  |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌?no  |
| `--trace`             | Print out full stack trace on errors                                                   | ❌?no  |
```

# sekaid add-genesis-account --help
Add a genesis account to genesis.json. The provided account must specify
the account address or key name and a list of initial coins. If a key name is given,
the address will be looked up in the local Keybase. The list of initial tokens must
contain valid denominations. Accounts may optionally be supplied with vesting parameters.

Usage:
  sekaid add-genesis-account [address_or_key_name] [coin][,[coin]] [flags]

Flags:
      --height int               Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help                     help for add-genesis-account
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test) (default "os")
      --node string              <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string            Output format (text|json) (default "text")
      --vesting-amount string    amount of coins for vesting accounts
      --vesting-end-time int     schedule end time (unix epoch) for vesting accounts
      --vesting-start-time int   schedule start time (unix epoch) for vesting accounts

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid add-genesis-account validator 150000000000000ukex,300000000000000test,2000000000000000000000000000samolean,1000000lol --keyring-backend=test
```

[Return to top](#sekai)

### 2. collect-gentxs

Adds validator into the genesis set

Usage:

```
sekaid gentx-claim [key_name] [flags]
```

| Flags                    | Description                                                             | Work  |
| ------------------------ | ----------------------------------------------------------------------- | ----- |
| -h, --help               | help for gentx-claim                                                    | ✅ yes |
| --keyring-backend string | Select keyring's backend (os\|file\|kwallet\|pass\|test) (default "os") | ✅ yes |
| --moniker string         | the Moniker                                                             | ✅ yes |
| --pubkey string          | the public key                                                          | ❌ ?   |



| Global Flags         |                                                                                        |       |
| -------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--home string`      | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format`       | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ no  |
| `--log_level string` | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`            | print out full stack trace on errors                                                   | ❌ ?   |

```
# sekaid gentx-claim --help
Adds validator into the genesis set

Usage:
  sekaid gentx-claim [key_name] [flags]

Flags:
  -h, --help                     help for gentx-claim
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test) (default "os")
      --moniker string           the Moniker
      --pubkey string            the public key

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
# sekaid gentx-claim validator --keyring-backend=test --moniker="GENESIS VALIDATOR" --home=$SEKAID_HOME
genesis state updated to include validator
```

[Return to top](#sekai)

### 3. config
Create or query an application CLI configuration file

Usage:
```
sekaid config <key> [value] [flags]
```

| Flags  | Description      | Work |
| ------ | ---------------- | ---- |
| --help | help for version | ✅yes |



| Global Flags         |                                                                                        |     |
| -------------------- | -------------------------------------------------------------------------------------- | --- |
| `--home string`      | directory for config and data (default `"/root/.sekaid"`)                              | ❌no |
| `--log_format`       | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌no |
| `--log_level string` | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌?  |
| `--trace`            | print out full stack trace on errors                                                   | ❌?  |

```
#sekaid config --help

Create or query an application CLI configuration file

Usage:
  sekaid config <key> [value] [flags]

Flags:
  -h, --help   help for config

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

#if send ```sekaid config``` printing current config state of the node

```
# sekaid config
{
        "chain-id": "",
        "keyring-backend": "os",
        "output": "text",
        "node": "tcp://localhost:26657",
        "broadcast-mode": "sync"
}
```

#if keyring-backend value will be wrong it will break entire sekaid installation

```
sekaid config keyring-backend test
```

[Return to top](#sekai)

### 4. debug


Tool for helping with debugging your application

Usage:

```
  sekaid debug [flags]
```

```
  sekaid debug [command]
```

| Subcommand                   | Description                                          |
| ---------------------------- | ---------------------------------------------------- |
| [`addr`](#41-addr)           | Convert an address between hex and bech32            |
| [`pubkey`](#42-pubkey)       | Decode a pubkey from proto JSON                      |
| [`raw-bytes`](#43-raw-bytes) | Convert raw bytes output (eg. [10 21 13 255]) to hex |



| Flags        | Description   | Work  |
| ------------ | ------------- | ----- |
| `-h, --help` | Help for keys | ✅ yes |




| Global Flags          | Description                                                                            | Work |
| --------------------- | -------------------------------------------------------------------------------------- | ---- |
| `--home string`       | Directory for config and data (default `"/root/.sekaid"`)                              | ❌ no |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?  |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?  |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?  |

```
# sekaid debug -h
Tool for helping with debugging your application

Usage:
  sekaid debug [flags]
  sekaid debug [command]

Available Commands:
  addr        Convert an address between hex and bech32
  pubkey      Decode a pubkey from proto JSON
  raw-bytes   Convert raw bytes output (eg. [10 21 13 255]) to hex

Flags:
  -h, --help   help for debug

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors

Use "sekaid debug [command] --help" for more information about a command.
```

[Return to top](#sekai)

#### 4.1 addr

Convert an address between hex encoding and bech32.

Usage:

```
sekaid debug addr [address] [flags]
```

| Flags  | Description      | Work |
| ------ | ---------------- | ---- |
| --help | help for version | ✅yes |



| Global Flags         |                                                                                        |     |
| -------------------- | -------------------------------------------------------------------------------------- | --- |
| `--home string`      | directory for config and data (default `"/root/.sekaid"`)                              | ❌no |
| `--log_format`       | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌no |
| `--log_level string` | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌?  |
| `--trace`            | print out full stack trace on errors                                                   | ❌?  |


```
# sekaid debug addr -h
Convert an address between hex encoding and bech32.

Example:
$ sekaid debug addr cosmos1e0jnq2sun3dzjh8p2xq95kk0expwmd7shwjpfg

Usage:
  sekaid debug addr [address] [flags]

Flags:
  -h, --help   help for addr

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
# sekaid debug addr kira1wh5zkqgg87tes5r6ycp5r56qgyvzglhply5yu4
Address: [117 232 43 1 8 63 151 152 80 122 38 3 65 211 64 65 24 36 126 225]
Address (hex): 75E82B01083F9798507A260341D3404118247EE1
Bech32 Acc: kira1wh5zkqgg87tes5r6ycp5r56qgyvzglhply5yu4
Bech32 Val: kiravaloper1wh5zkqgg87tes5r6ycp5r56qgyvzglhpvzg8ye
```

[Return to top](#sekai)

#### 4.2 pubkey

Decode a pubkey from proto JSON and display it's address.

Usage:

```
sekaid debug pubkey [pubkey] [flags]
```

| Flags  | Description      | Work |
| ------ | ---------------- | ---- |
| --help | help for version | ✅yes |



| Global Flags         |                                                                                        |     |
| -------------------- | -------------------------------------------------------------------------------------- | --- |
| `--home string`      | directory for config and data (default `"/root/.sekaid"`)                              | ❌no |
| `--log_format`       | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌no |
| `--log_level string` | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌?  |
| `--trace`            | print out full stack trace on errors                                                   | ❌?  |



```
# sekaid debug pubkey  -h
Decode a pubkey from proto JSON and display it's address.

Example:
$ sekaid debug pubkey '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AurroA7jvfPd1AadmmOvWM2rJSwipXfRf8yD6pLbA2DJ"}'

Usage:
  sekaid debug pubkey [pubkey] [flags]

Flags:
  -h, --help   help for pubkey

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
# sekaid debug pubkey  '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A5O/lBKffZNUZN2rJLbu48mOX/Q/QKHBOjXGPlFn/jJ6"}'
Address: 03C5F71626A9F436387F232823DFE73AFBD46ABC
PubKey Hex: 0393bf94129f7d935464ddab24b6eee3c98e5ff43f40a1c13a35c63e5167fe327a
```

[Return to top](#sekai)

#### 4.3 raw-bytes

Convert raw-bytes to hex.

⚠️ maybe not working properly or the description is wrong

Usage:

```
sekaid debug addr [address] [flags]
```

| Flags  | Description      | Work |
| ------ | ---------------- | ---- |
| --help | help for version | ✅yes |



| Global Flags         |                                                                                        |     |
| -------------------- | -------------------------------------------------------------------------------------- | --- |
| `--home string`      | directory for config and data (default `"/root/.sekaid"`)                              | ❌no |
| `--log_format`       | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌no |
| `--log_level string` | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌?  |
| `--trace`            | print out full stack trace on errors                                                   | ❌?  |



```
# sekaid debug raw-bytes -h
Convert raw-bytes to hex.

Example:
$ sekaid debug raw-bytes [72 101 108 108 111 44 32 112 108 97 121 103 114 111 117 110 100]

Usage:
  sekaid debug raw-bytes [raw-bytes] [flags]

Flags:
  -h, --help   help for raw-bytes

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
# sekaid debug raw-bytes '72 101 108 108 111 44 32 112 108 97 121 103 114 111 117 110 100'
48656C6C6F2C20706C617967726F756E64
```

[Return to top](#sekai)

### 5. export

Export state to JSON

Usage:

```
sekaid export [flags]
```

| Flags                        | Description                                                                 | Work   |
| ---------------------------- | --------------------------------------------------------------------------- | ------ |
| --for-zero-height            | Export state to start at height zero (perform preproccessing)               | ❌ ? no |
| --height int                 | Export state from a particular height (-1 means latest height) (default -1) | ❌ ? no |
| -h, --help                   | help for export                                                             | ✅ yes  |
| --jail-allowed-addrs strings | Comma-separated list of operator addresses of jailed validators to unjail   | ❌ ?    |



| Global Flags       |                                                                                    |       |
| ------------------ | ---------------------------------------------------------------------------------- | ----- |
| --home string      | directory for config and data (default "/root/.sekaid")                            | ✅ yes |
| --log_format       | The logging format (json\|plain) (default "plain")                                 | ❌ no  |
| --log_level string | The logging level (trace\|debug\|info\|warn\|error\|fatal\|panic) (default "info") | ❌ ?   |
| --trace            | print out full stack trace on errors                                               | ❌ ?   |

```
# sekaid export -h
Export state to JSON

Usage:
  sekaid export [flags]

Flags:
      --for-zero-height              Export state to start at height zero (perform preproccessing)
      --height int                   Export state from a particular height (-1 means latest height) (default -1)
  -h, --help                         help for export
      --jail-allowed-addrs strings   Comma-separated list of operator addresses of jailed validators to unjail

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid export
```

<details>
  <summary>Check export out</summary>
  
```
# sekaid export | jq
{
  "app_hash": "",
  "app_state": {
    "auth": {
      "accounts": [
        {
          "@type": "/cosmos.auth.v1beta1.ModuleAccount",
          "base_account": {
            "account_number": "5",
            "address": "kira1yckd06xnz444h9c38gmnepx4x2yak3fsgpdq5w",
            "pub_key": null,
            "sequence": "0"
          },
          "name": "spending",
          "permissions": []
        },
        {
          "@type": "/cosmos.auth.v1beta1.BaseAccount",
          "account_number": "0",
          "address": "kira1tcyq0y66mmdpydepexc8lwrmemp9wzgmg2rqhm",
          "pub_key": {
            "@type": "/cosmos.crypto.secp256k1.PubKey",
            "key": "Ask+LXUBzB8JEPA9fyuCBM+163QLcl2kwH2dHcHxtLNU"
          },
          "sequence": "29"
        },
        {
          "@type": "/cosmos.auth.v1beta1.BaseAccount",
          "account_number": "2",
          "address": "kira1hcj5m0y3gmr6d80funkn6zj8e4amye3wk5jdmk",
          "pub_key": {
            "@type": "/cosmos.crypto.secp256k1.PubKey",
            "key": "A2iLCs+68myGXWJjrucYAtvMOSFoLRC7ILCNsjDc15HN"
          },
          "sequence": "1"
        },
        {
          "@type": "/cosmos.auth.v1beta1.BaseAccount",
          "account_number": "1",
          "address": "kira1m2m98e2fahrzvlhdc6rvf645plzt2e626knhn8",
          "pub_key": {
            "@type": "/cosmos.crypto.secp256k1.PubKey",
            "key": "Ag/kOuIB+9YTzgLFPoQcEJQXoD4BuO7lMX8eX7DYtloV"
          },
          "sequence": "1"
        },
        {
          "@type": "/cosmos.auth.v1beta1.ModuleAccount",
          "base_account": {
            "account_number": "4",
            "address": "kira1m3h30wlvsf8llruxtpukdvsy0km2kum8c7dkxd",
            "pub_key": null,
            "sequence": "0"
          },
          "name": "mint",
          "permissions": [
            "minter"
          ]
        },
        {
          "@type": "/cosmos.auth.v1beta1.ModuleAccount",
          "base_account": {
            "account_number": "3",
            "address": "kira17xpfvakm2amg962yls6f84z3kell8c5lqkfw2s",
            "pub_key": null,
            "sequence": "0"
          },
          "name": "fee_collector",
          "permissions": []
        }
      ],
      "params": {
        "max_memo_characters": "256",
        "sig_verify_cost_ed25519": "590",
        "sig_verify_cost_secp256k1": "1000",
        "tx_sig_limit": "7",
        "tx_size_cost_per_byte": "10"
      }
    },
    "bank": {
      "balances": [
        {
          "address": "kira1yckd06xnz444h9c38gmnepx4x2yak3fsgpdq5w",
          "coins": [
            {
              "amount": "500000000000",
              "denom": "ukex"
            }
          ]
        },
        {
          "address": "kira1tcyq0y66mmdpydepexc8lwrmemp9wzgmg2rqhm",
          "coins": [
            {
              "amount": "1000000",
              "denom": "lol"
            },
            {
              "amount": "2000000000000000000000000000",
              "denom": "samolean"
            },
            {
              "amount": "29999780000000000",
              "denom": "test"
            },
            {
              "amount": "299998799997100",
              "denom": "ukex"
            }
          ]
        },
        {
          "address": "kira1hcj5m0y3gmr6d80funkn6zj8e4amye3wk5jdmk",
          "coins": [
            {
              "amount": "1000000",
              "denom": "lol"
            },
            {
              "amount": "3000000000000000000000000000",
              "denom": "samolean"
            },
            {
              "amount": "200000000000",
              "denom": "test"
            },
            {
              "amount": "999999900",
              "denom": "ukex"
            }
          ]
        },
        {
          "address": "kira1m2m98e2fahrzvlhdc6rvf645plzt2e626knhn8",
          "coins": [
            {
              "amount": "10000000000",
              "denom": "test"
            },
            {
              "amount": "99999900",
              "denom": "ukex"
            }
          ]
        },
        {
          "address": "kira17xpfvakm2amg962yls6f84z3kell8c5lqkfw2s",
          "coins": [
            {
              "amount": "13616387924",
              "denom": "ukex"
            }
          ]
        }
      ],
      "denom_metadata": [],
      "params": {
        "default_send_enabled": true,
        "send_enabled": []
      },
      "supply": [
        {
          "amount": "2000000",
          "denom": "lol"
        },
        {
          "amount": "5000000000000000000000000000",
          "denom": "samolean"
        },
        {
          "amount": "29999990000000000",
          "denom": "test"
        },
        {
          "amount": "300513516384824",
          "denom": "ukex"
        }
      ]
    },
    "basket": {
      "baskets": [],
      "historical_burns": [],
      "historical_mints": [],
      "historical_swaps": []
    },
    "collectives": {
      "collectives": [],
      "contributers": []
    },
    "custody": null,
    "customevidence": {
      "evidence": []
    },
    "customgov": {
      "data_registry": {},
      "execution_fees": [
        {
          "default_parameters": "0",
          "execution_fee": "100",
          "failure_fee": "1000",
          "timeout": "10",
          "transaction_type": "activate"
        },
        {
          "default_parameters": "0",
          "execution_fee": "100",
          "failure_fee": "1",
          "timeout": "10",
          "transaction_type": "claim-councilor"
        },
        {
          "default_parameters": "0",
          "execution_fee": "100",
          "failure_fee": "1",
          "timeout": "10",
          "transaction_type": "claim-proposal-type-x"
        },
        {
          "default_parameters": "0",
          "execution_fee": "100",
          "failure_fee": "1",
          "timeout": "10",
          "transaction_type": "claim-validator"
        },
        {
          "default_parameters": "0",
          "execution_fee": "100",
          "failure_fee": "100",
          "timeout": "10",
          "transaction_type": "pause"
        },
        {
          "default_parameters": "0",
          "execution_fee": "10",
          "failure_fee": "1",
          "timeout": "10",
          "transaction_type": "submit-proposal-type-x"
        },
        {
          "default_parameters": "0",
          "execution_fee": "100",
          "failure_fee": "100",
          "timeout": "10",
          "transaction_type": "unpause"
        },
        {
          "default_parameters": "0",
          "execution_fee": "100",
          "failure_fee": "1",
          "timeout": "10",
          "transaction_type": "upsert-token-alias"
        },
        {
          "default_parameters": "0",
          "execution_fee": "100",
          "failure_fee": "1",
          "timeout": "10",
          "transaction_type": "veto-proposal-type-x"
        },
        {
          "default_parameters": "0",
          "execution_fee": "100",
          "failure_fee": "1",
          "timeout": "10",
          "transaction_type": "vote-proposal-type-x"
        }
      ],
      "id_records_verify_requests": [],
      "identity_records": [
        {
          "address": "kira1tcyq0y66mmdpydepexc8lwrmemp9wzgmg2rqhm",
          "date": "2023-05-20T08:48:01.608922330Z",
          "id": "1",
          "key": "moniker",
          "value": "GENESIS VALIDATOR",
          "verifiers": []
        },
        {
          "address": "kira1tcyq0y66mmdpydepexc8lwrmemp9wzgmg2rqhm",
          "date": "2023-05-20T08:49:44.599380563Z",
          "id": "2",
          "key": "description",
          "value": "This is genesis validator account of the KIRA Team",
          "verifiers": []
        },
        {
          "address": "kira1tcyq0y66mmdpydepexc8lwrmemp9wzgmg2rqhm",
          "date": "2023-05-20T08:49:54.615435843Z",
          "id": "3",
          "key": "social",
          "value": "https://tg.kira.network,twitter.kira.network",
          "verifiers": []
        },
        {
          "address": "kira1tcyq0y66mmdpydepexc8lwrmemp9wzgmg2rqhm",
          "date": "2023-05-20T08:50:04.631054986Z",
          "id": "4",
          "key": "contact",
          "value": "https://support.kira.network",
          "verifiers": []
        },
        {
          "address": "kira1tcyq0y66mmdpydepexc8lwrmemp9wzgmg2rqhm",
          "date": "2023-05-20T08:50:14.664373675Z",
          "id": "5",
          "key": "website",
          "value": "https://kira.network",
          "verifiers": []
        },
        {
          "address": "kira1tcyq0y66mmdpydepexc8lwrmemp9wzgmg2rqhm",
          "date": "2023-05-20T08:50:24.692282875Z",
          "id": "6",
          "key": "username",
          "value": "KIRA",
          "verifiers": []
        },
        {
          "address": "kira1tcyq0y66mmdpydepexc8lwrmemp9wzgmg2rqhm",
          "date": "2023-05-20T08:50:34.717593161Z",
          "id": "7",
          "key": "logo",
          "value": "https://kira-network.s3-eu-west-1.amazonaws.com/assets/img/tokens/kex.svg",
          "verifiers": []
        },
        {
          "address": "kira1tcyq0y66mmdpydepexc8lwrmemp9wzgmg2rqhm",
          "date": "2023-05-20T08:50:44.750398780Z",
          "id": "8",
          "key": "avatar",
          "value": "https://kira-network.s3-eu-west-1.amazonaws.com/assets/img/tokens/kex.svg",
          "verifiers": []
        },
        {
          "address": "kira1tcyq0y66mmdpydepexc8lwrmemp9wzgmg2rqhm",
          "date": "2023-05-20T08:50:54.786173458Z",
          "id": "9",
          "key": "pentest1",
          "value": "<iframe src=javascript:alert(1)>",
          "verifiers": []
        },
        {
          "address": "kira1tcyq0y66mmdpydepexc8lwrmemp9wzgmg2rqhm",
          "date": "2023-05-20T08:51:04.821495895Z",
          "id": "10",
          "key": "pentest2",
          "value": "<img/src=x a='' onerror=alert(2)>",
          "verifiers": []
        },
        {
          "address": "kira1tcyq0y66mmdpydepexc8lwrmemp9wzgmg2rqhm",
          "date": "2023-05-20T08:51:14.876204916Z",
          "id": "11",
          "key": "pentest3",
          "value": "<img src=1 onerror=alert(3)>",
          "verifiers": []
        },
        {
          "address": "kira1tcyq0y66mmdpydepexc8lwrmemp9wzgmg2rqhm",
          "date": "2023-05-20T08:51:24.930421584Z",
          "id": "12",
          "key": "validator_node_id",
          "value": "7c681cf7ea7b891e45207b3c3245926043df0f4f",
          "verifiers": []
        },
        {
          "address": "kira1m2m98e2fahrzvlhdc6rvf645plzt2e626knhn8",
          "date": "2023-05-20T08:51:34.975587365Z",
          "id": "13",
          "key": "username",
          "value": "test",
          "verifiers": []
        },
        {
          "address": "kira1hcj5m0y3gmr6d80funkn6zj8e4amye3wk5jdmk",
          "date": "2023-05-20T08:51:44.995161060Z",
          "id": "14",
          "key": "username",
          "value": "faucet",
          "verifiers": []
        }
      ],
      "last_id_record_verify_request_id": "0",
      "last_identity_record_id": "14",
      "network_actors": [
        {
          "address": "kira1tcyq0y66mmdpydepexc8lwrmemp9wzgmg2rqhm",
          "permissions": {
            "blacklist": [],
            "whitelist": [
              4,
              35,
              14,
              28,
              5,
              36,
              15,
              29
            ]
          },
          "roles": [
            "1"
          ],
          "skin": "1",
          "status": "ACTIVE",
          "votes": [
            "VOTE_OPTION_YES",
            "VOTE_OPTION_ABSTAIN",
            "VOTE_OPTION_NO",
            "VOTE_OPTION_NO_WITH_VETO"
          ]
        }
      ],
      "network_properties": {
        "abstention_rank_decrease_amount": "1",
        "dapp_auto_denounce_time": "60",
        "dapp_bond_duration": "604800",
        "dapp_inactive_rank_decrease_percent": "10",
        "dapp_liquidation_period": "0",
        "dapp_liquidation_threshold": "0",
        "dapp_max_mischance": "10",
        "dapp_mischance_rank_decrease_amount": "1",
        "dapp_pool_slippage_default": "0.100000000000000000",
        "dapp_verifier_bond": "0.001000000000000000",
        "enable_foreign_fee_payments": true,
        "enable_token_blacklist": true,
        "enable_token_whitelist": false,
        "inactive_rank_decrease_percent": "0.500000000000000000",
        "inflation_period": "31557600",
        "inflation_rate": "0.180000000000000000",
        "max_abstention": "2",
        "max_annual_inflation": "0.350000000000000000",
        "max_collective_outputs": "10",
        "max_custody_buffer_size": "10",
        "max_custody_tx_size": "8192",
        "max_dapp_bond": "10000000",
        "max_delegators": "100",
        "max_jailed_percentage": "0.250000000000000000",
        "max_mischance": "50",
        "max_proposal_checksum_size": "128",
        "max_proposal_description_size": "1024",
        "max_proposal_poll_option_count": "128",
        "max_proposal_poll_option_size": "64",
        "max_proposal_reference_size": "512",
        "max_proposal_title_size": "128",
        "max_slashing_percentage": "0.010000000000000000",
        "max_tx_fee": "1000000",
        "min_collective_bond": "100000",
        "min_collective_bonding_time": "86400",
        "min_collective_claim_period": "14400",
        "min_custody_reward": "200",
        "min_dapp_bond": "1000000",
        "min_delegation_pushout": "10",
        "min_identity_approval_tip": "200",
        "min_proposal_enactment_blocks": "1",
        "min_proposal_end_blocks": "2",
        "min_tx_fee": "100",
        "min_validators": "1",
        "minimum_proposal_end_time": "360",
        "minting_ft_fee": "100000000000000",
        "minting_nft_fee": "100000000000000",
        "mischance_confidence": "25",
        "mischance_rank_decrease_amount": "1",
        "poor_network_max_bank_send": "1000000",
        "proposal_enactment_time": "300",
        "slashing_period": "3600",
        "ubi_hardcap": "6000000",
        "unique_identity_keys": "moniker,username",
        "unjail_max_time": "1209600",
        "unstaking_period": "2629800",
        "validator_recovery_bond": "300000",
        "validators_fee_share": "0.500000000000000000",
        "vote_quorum": "33"
      },
      "next_role_id": "3",
      "poor_network_messages": {
        "messages": [
          "submit-proposal",
          "set-network-properties",
          "vote-proposal",
          "claim-councilor",
          "whitelist-permissions",
          "blacklist-permissions",
          "create-role",
          "assign-role",
          "remove-role",
          "whitelist-role-permission",
          "blacklist-role-permission",
          "remove-whitelist-role-permission",
          "remove-blacklist-role-permission",
          "claim-validator",
          "activate",
          "pause",
          "unpause",
          "register-identity-records",
          "edit-identity-record",
          "request-identity-records-verify",
          "handle-identity-records-verify-request",
          "cancel-identity-records-verify-request"
        ]
      },
      "proposal_durations": {},
      "proposals": [
        {
          "content": {
            "@type": "/kira.tokens.ProposalUpsertTokenAlias",
            "decimals": 6,
            "denoms": [
              "ukex"
            ],
            "icon": "http://kira-network.s3-eu-west-1.amazonaws.com/assets/img/tokens/kex.svg",
            "invalidated": false,
            "name": "KIRA",
            "symbol": "KEX"
          },
          "description": "Initial Setup From KIRA Manager",
          "enactment_end_time": "2023-05-20T09:02:55.050874780Z",
          "exec_result": "executed successfully",
          "min_enactment_end_block_height": "44",
          "min_voting_end_block_height": "25",
          "proposal_id": "1",
          "result": "VOTE_RESULT_PASSED",
          "submit_time": "2023-05-20T08:51:55.050874780Z",
          "title": "Upsert KEX icon URL link",
          "voting_end_time": "2023-05-20T08:57:55.050874780Z"
        },
        {
          "content": {
            "@type": "/kira.tokens.ProposalUpsertTokenAlias",
            "decimals": 8,
            "denoms": [
              "test"
            ],
            "icon": "http://kira-network.s3-eu-west-1.amazonaws.com/assets/img/tokens/test.svg",
            "invalidated": false,
            "name": "Test TestCoin",
            "symbol": "TEST"
          },
          "description": "Initial Setup From KIRA Manager",
          "enactment_end_time": "2023-05-20T09:03:15.104192359Z",
          "exec_result": "executed successfully",
          "min_enactment_end_block_height": "44",
          "min_voting_end_block_height": "27",
          "proposal_id": "2",
          "result": "VOTE_RESULT_PASSED",
          "submit_time": "2023-05-20T08:52:15.104192359Z",
          "title": "Upsert Test TestCoin icon URL link",
          "voting_end_time": "2023-05-20T08:58:15.104192359Z"
        },
        {
          "content": {
            "@type": "/kira.tokens.ProposalUpsertTokenAlias",
            "decimals": 18,
            "denoms": [
              "samolean"
            ],
            "icon": "http://kira-network.s3-eu-west-1.amazonaws.com/assets/img/tokens/samolean.svg",
            "invalidated": false,
            "name": "Samolean TestCoin",
            "symbol": "SAMO"
          },
          "description": "Initial Setup From KIRA Manager",
          "enactment_end_time": "2023-05-20T09:03:35.177132478Z",
          "exec_result": "executed successfully",
          "min_enactment_end_block_height": "44",
          "min_voting_end_block_height": "29",
          "proposal_id": "3",
          "result": "VOTE_RESULT_PASSED",
          "submit_time": "2023-05-20T08:52:35.177132478Z",
          "title": "Upsert Samolean TestCoin icon URL link",
          "voting_end_time": "2023-05-20T08:58:35.177132478Z"
        },
        {
          "content": {
            "@type": "/kira.upgrade.ProposalSoftwareUpgrade",
            "instate_upgrade": true,
            "max_enrolment_duration": "666",
            "memo": "Genesis Setup Plan",
            "name": "genesis",
            "new_chain_id": "localnet-47",
            "old_chain_id": "localnet-47",
            "reboot_required": true,
            "resources": [
              {
                "checksum": "",
                "id": "kira",
                "url": "https://github.com/KiraCore/kira/releases/download/v0.11.21/kira.zip",
                "version": ""
              }
            ],
            "rollback_checksum": "genesis",
            "skip_handler": true,
            "upgrade_time": "1684573615"
          },
          "description": "",
          "enactment_end_time": "2023-05-20T09:03:55.214931655Z",
          "exec_result": "execution failed",
          "min_enactment_end_block_height": "44",
          "min_voting_end_block_height": "31",
          "proposal_id": "4",
          "result": "VOTE_RESULT_PASSED",
          "submit_time": "2023-05-20T08:52:55.214931655Z",
          "title": "",
          "voting_end_time": "2023-05-20T08:58:55.214931655Z"
        }
      ],
      "role_permissions": {
        "1": {
          "blacklist": [],
          "whitelist": [
            1,
            2,
            3,
            6,
            8,
            9,
            12,
            13,
            10,
            11,
            14,
            15,
            18,
            19,
            20,
            21,
            22,
            23,
            31,
            32,
            24,
            25,
            16,
            17,
            4,
            5,
            26,
            27,
            28,
            29,
            30,
            33,
            34,
            35,
            36,
            37,
            38,
            39,
            40,
            41,
            42,
            43,
            44,
            45,
            46,
            47,
            48,
            49,
            50,
            51,
            52,
            53,
            54,
            55,
            56,
            57,
            58,
            59,
            60,
            61,
            62,
            63,
            64,
            65,
            66
          ]
        },
        "2": {
          "blacklist": [],
          "whitelist": [
            2
          ]
        }
      },
      "roles": [
        {
          "description": "Sudo role",
          "id": 1,
          "sid": "sudo"
        },
        {
          "description": "Validator role",
          "id": 2,
          "sid": "validator"
        }
      ],
      "starting_proposal_id": "5",
      "votes": [
        {
          "option": "VOTE_OPTION_YES",
          "proposal_id": "1",
          "slash": "0.010000000000000000",
          "voter": "kira1tcyq0y66mmdpydepexc8lwrmemp9wzgmg2rqhm"
        },
        {
          "option": "VOTE_OPTION_YES",
          "proposal_id": "2",
          "slash": "0.010000000000000000",
          "voter": "kira1tcyq0y66mmdpydepexc8lwrmemp9wzgmg2rqhm"
        },
        {
          "option": "VOTE_OPTION_YES",
          "proposal_id": "3",
          "slash": "0.010000000000000000",
          "voter": "kira1tcyq0y66mmdpydepexc8lwrmemp9wzgmg2rqhm"
        },
        {
          "option": "VOTE_OPTION_YES",
          "proposal_id": "4",
          "slash": "0.010000000000000000",
          "voter": "kira1tcyq0y66mmdpydepexc8lwrmemp9wzgmg2rqhm"
        }
      ]
    },
    "customslashing": {
      "params": {
        "downtime_inactive_duration": "600s"
      },
      "signing_infos": [
        {
          "address": "kiravalcons1dcu5arfvrfvwe9wxrdgtfap0ry3yru02s8rj76",
          "validator_signing_info": {
            "address": "kiravalcons1dcu5arfvrfvwe9wxrdgtfap0ry3yru02s8rj76",
            "inactive_until": "1970-01-01T00:00:00Z",
            "last_present_block": "3973",
            "mischance": "0",
            "mischance_confidence": "0",
            "missed_blocks_counter": "0",
            "produced_blocks_counter": "3972",
            "start_height": "0"
          }
        }
      ]
    },
    "customstaking": {
      "validators": [
        {
          "pub_key": {
            "@type": "/cosmos.crypto.ed25519.PubKey",
            "key": "T6DNclgBnpmdXM4mduAoVwP5T8U2E+fg749Roig9wPI="
          },
          "rank": "3972",
          "status": "ACTIVE",
          "streak": "3972",
          "val_key": "kiravaloper1tcyq0y66mmdpydepexc8lwrmemp9wzgmmvlr0h"
        }
      ]
    },
    "distributor": {
      "fees_collected": [],
      "fees_treasury": [
        {
          "amount": "13616387924",
          "denom": "ukex"
        }
      ],
      "previous_proposer": "kiravalcons1dcu5arfvrfvwe9wxrdgtfap0ry3yru02s8rj76",
      "snap_period": "1000",
      "validator_votes": [],
      "year_start_snapshot": {
        "snapshot_amount": "300499900000000",
        "snapshot_time": "1684572480"
      }
    },
    "feeprocessing": null,
    "genutil": {
      "gen_txs": []
    },
    "layer2": {
      "bridge": {
        "accounts": [],
        "helper": null,
        "tokens": [],
        "xams": []
      },
      "dapps": []
    },
    "multistaking": {
      "info": [],
      "pools": [],
      "rewards": [],
      "undelegations": []
    },
    "params": null,
    "recovery": {
      "recovery_records": [],
      "recovery_tokens": [],
      "rewards": [],
      "rotations": []
    },
    "spending": {
      "claims": [],
      "pools": [
        {
          "balances": [
            {
              "amount": "500000000000",
              "denom": "ukex"
            }
          ],
          "beneficiaries": {
            "accounts": [],
            "roles": [
              {
                "role": "2",
                "weight": "1"
              }
            ]
          },
          "claim_end": "0",
          "claim_expiry": "0",
          "claim_start": "0",
          "dynamic_rate": false,
          "dynamic_rate_period": "0",
          "last_dynamic_rate_calc_time": "0",
          "name": "ValidatorBasicRewardsPool",
          "owners": {
            "owner_accounts": [],
            "owner_roles": [
              "2"
            ]
          },
          "rates": [
            {
              "amount": "385.000000000000000000",
              "denom": "ukex"
            }
          ],
          "vote_enactment": "300",
          "vote_period": "300",
          "vote_quorum": "33"
        }
      ]
    },
    "tokens": {
      "aliases": [
        {
          "decimals": 6,
          "denoms": [
            "ukex"
          ],
          "icon": "http://kira-network.s3-eu-west-1.amazonaws.com/assets/img/tokens/kex.svg",
          "invalidated": false,
          "name": "KIRA",
          "symbol": "KEX"
        },
        {
          "decimals": 18,
          "denoms": [
            "samolean"
          ],
          "icon": "http://kira-network.s3-eu-west-1.amazonaws.com/assets/img/tokens/samolean.svg",
          "invalidated": false,
          "name": "Samolean TestCoin",
          "symbol": "SAMO"
        },
        {
          "decimals": 8,
          "denoms": [
            "test"
          ],
          "icon": "http://kira-network.s3-eu-west-1.amazonaws.com/assets/img/tokens/test.svg",
          "invalidated": false,
          "name": "Test TestCoin",
          "symbol": "TEST"
        }
      ],
      "rates": [
        {
          "denom": "frozen",
          "fee_payments": true,
          "fee_rate": "0.100000000000000000",
          "invalidated": false,
          "stake_cap": "0.000000000000000000",
          "stake_min": "1",
          "stake_token": false
        },
        {
          "denom": "ubtc",
          "fee_payments": true,
          "fee_rate": "10.000000000000000000",
          "invalidated": false,
          "stake_cap": "0.250000000000000000",
          "stake_min": "1",
          "stake_token": true
        },
        {
          "denom": "ukex",
          "fee_payments": true,
          "fee_rate": "1.000000000000000000",
          "invalidated": false,
          "stake_cap": "0.500000000000000000",
          "stake_min": "1",
          "stake_token": true
        },
        {
          "denom": "xeth",
          "fee_payments": true,
          "fee_rate": "0.100000000000000000",
          "invalidated": false,
          "stake_cap": "0.100000000000000000",
          "stake_min": "1",
          "stake_token": false
        }
      ],
      "tokenBlackWhites": {
        "blacklisted": [
          "frozen"
        ],
        "whitelisted": [
          "ukex"
        ]
      }
    },
    "ubi": {
      "ubi_records": [
        {
          "amount": "500000",
          "distribution_end": "0",
          "distribution_last": "1684572480",
          "distribution_start": "0",
          "dynamic": true,
          "name": "ValidatorBasicRewardsPoolUBI",
          "period": "2592000",
          "pool": "ValidatorBasicRewardsPool"
        }
      ]
    },
    "upgrade": {
      "current_plan": null,
      "next_plan": null,
      "version": "v0.1.22.11"
    }
  },
  "chain_id": "localnet-47",
  "consensus_params": {
    "block": {
      "max_bytes": "22020096",
      "max_gas": "-1",
      "time_iota_ms": "1000"
    },
    "evidence": {
      "max_age_duration": "172800000000000",
      "max_age_num_blocks": "100000",
      "max_bytes": "1048576"
    },
    "validator": {
      "pub_key_types": [
        "ed25519"
      ]
    },
    "version": {}
  },
  "genesis_time": "2023-05-20T08:48:00.848282749Z",
  "initial_height": "3974"
}
```
</details>


[Return to top](#sekai)

### 6. export-metadata

Get metadata for client interaction to the node

Usage:

```
sekaid export-metadata [flags]
```

| Flags      | Description   | Work  |
| ---------- | ------------- | ----- |
| -h, --help | help for init | ✅ yes |



| Global Flags       |                                                                                    |      |
| ------------------ | ---------------------------------------------------------------------------------- | ---- |
| --home string      | directory for config and data (default "/root/.sekaid")                            | ❌ no |
| --log_format       | The logging format (json\|plain) (default "plain")                                 | ❌ no |
| --log_level string | The logging level (trace\|debug\|info\|warn\|error\|fatal\|panic) (default "info") | ❌ ?  |
| --trace            | print out full stack trace on errors                                               | ❌ ?  |

```
# sekaid export-metadata -h
Get metadata for client interaction to the node

Usage:
  sekaid export-metadata [flags]

Flags:
  -h, --help   help for export-metadata

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors

```

```
# sekaid export-metadata
```
<details>
  <summary>Check export-metadata out</summary>

```
# sekaid export-metadata | jq    
{
  "permissions": [
    {
      "id": 1,
      "name": "PERMISSION_SET_PERMISSIONS",
      "module": "gov",
      "description": "the permission that allows to Set Permissions to other actors"
    },
    {
      "id": 2,
      "name": "PERMISSION_CLAIM_VALIDATOR",
      "module": "customstaking",
      "description": "the permission that allows to Claim a validator Seat"
    },
    {
      "id": 3,
      "name": "PERMISSION_CLAIM_COUNCILOR",
      "module": "gov",
      "description": "the permission that allows to Claim a Councilor Seat"
    },
    {
      "id": 4,
      "name": "PERMISSION_WHITELIST_ACCOUNT_PERMISSION_PROPOSAL",
      "module": "gov",
      "description": "the permission needed to create proposals for whitelisting an account permission."
    },
    {
      "id": 5,
      "name": "PERMISSION_VOTE_WHITELIST_ACCOUNT_PERMISSION_PROPOSAL",
      "module": "gov",
      "description": "the permission that an actor must have in order to vote a proposal to whitelist account permission"
    },
    {
      "id": 6,
      "name": "PERMISSION_UPSERT_TOKEN_ALIAS",
      "module": "tokens",
      "description": "the permission to upsert token alias"
    },
    {
      "id": 7,
      "name": "PERMISSION_CHANGE_TX_FEE",
      "module": "tokens",
      "description": "Permission to change tx fee"
    },
    {
      "id": 8,
      "name": "PERMISSION_UPSERT_TOKEN_RATE",
      "module": "tokens",
      "description": "Permission to upsert token rate"
    },
    {
      "id": 9,
      "name": "PERMISSION_UPSERT_ROLE",
      "module": "gov",
      "description": "Permission to upsert a role"
    },
    {
      "id": 10,
      "name": "PERMISSION_CREATE_UPSERT_DATA_REGISTRY_PROPOSAL",
      "module": "gov",
      "description": "makes possible to create a proposal to change the Data Registry"
    },
    {
      "id": 11,
      "name": "PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL",
      "module": "gov",
      "description": "makes possible to vote on a proposal to change the Data Registry."
    },
    {
      "id": 12,
      "name": "PERMISSION_CREATE_SET_NETWORK_PROPERTY_PROPOSAL",
      "module": "gov",
      "description": "the permission needed to create proposals for setting network property"
    },
    {
      "id": 13,
      "name": "PERMISSION_VOTE_SET_NETWORK_PROPERTY_PROPOSAL",
      "module": "gov",
      "description": "the permission that an actor must have in order to vote a proposal to set network property."
    },
    {
      "id": 14,
      "name": "PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL",
      "module": "tokens",
      "description": "the permission needed to create proposals for upsert token Alias."
    },
    {
      "id": 15,
      "name": "PERMISSION_VOTE_UPSERT_TOKEN_ALIAS_PROPOSAL",
      "module": "tokens",
      "description": "the permission needed to vote proposals for upsert token."
    },
    {
      "id": 16,
      "name": "PERMISSION_CREATE_SET_POOR_NETWORK_MESSAGES",
      "module": "tokens",
      "description": "the permission needed to create proposals for setting poor network messages."
    },
    {
      "id": 17,
      "name": "PERMISSION_VOTE_SET_POOR_NETWORK_MESSAGES_PROPOSAL",
      "module": "tokens",
      "description": "the permission needed to vote proposals to set poor network messages."
    },
    {
      "id": 18,
      "name": "PERMISSION_CREATE_UPSERT_TOKEN_RATE_PROPOSAL",
      "module": "tokens",
      "description": "the permission needed to create proposals for upsert token rate."
    },
    {
      "id": 19,
      "name": "PERMISSION_VOTE_UPSERT_TOKEN_RATE_PROPOSAL",
      "module": "tokens",
      "description": "the permission needed to vote proposals for upsert token rate."
    },
    {
      "id": 20,
      "name": "PERMISSION_CREATE_UNJAIL_VALIDATOR_PROPOSAL",
      "module": "slashing",
      "description": "the permission needed to create a proposal to unjail a validator."
    },
    {
      "id": 21,
      "name": "PERMISSION_VOTE_UNJAIL_VALIDATOR_PROPOSAL",
      "module": "slashing",
      "description": "the permission needed to vote a proposal to unjail a validator."
    },
    {
      "id": 22,
      "name": "PERMISSION_CREATE_CREATE_ROLE_PROPOSAL",
      "module": "gov",
      "description": "the permission needed to create a proposal to create a role."
    },
    {
      "id": 23,
      "name": "PERMISSION_VOTE_CREATE_ROLE_PROPOSAL",
      "module": "gov",
      "description": "the permission needed to vote a proposal to create a role."
    },
    {
      "id": 24,
      "name": "PERMISSION_CREATE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL",
      "module": "tokens",
      "description": "the permission needed to create a proposal to blacklist/whitelisted tokens"
    },
    {
      "id": 25,
      "name": "PERMISSION_VOTE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL",
      "module": "tokens",
      "description": "the permission needed to vote on blacklist/whitelisted tokens proposal"
    },
    {
      "id": 26,
      "name": "PERMISSION_CREATE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL",
      "module": "customstaking",
      "description": "the permission needed to create a proposal to reset whole validator rank"
    },
    {
      "id": 27,
      "name": "PERMISSION_VOTE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL",
      "module": "customstaking",
      "description": "the permission needed to vote on reset whole validator rank proposal"
    },
    {
      "id": 28,
      "name": "PERMISSION_CREATE_SOFTWARE_UPGRADE_PROPOSAL",
      "module": "upgrade",
      "description": "the permission needed to create a proposal for software upgrade"
    },
    {
      "id": 28,
      "name": "PermVoteSoftwareUpgradeProposal",
      "module": "upgrade",
      "description": "the permission needed to vote on software upgrade proposal"
    },
    {
      "id": 30,
      "name": "PERMISSION_SET_CLAIM_VALIDATOR_PERMISSION",
      "module": "customstaking",
      "description": "the permission that allows to Set ClaimValidatorPermission to other actors."
    },
    {
      "id": 31,
      "name": "PERMISSION_CREATE_SET_PROPOSAL_DURATION_PROPOSAL",
      "module": "gov",
      "description": "the permission needed to create a proposal to set proposal duration."
    },
    {
      "id": 32,
      "name": "PERMISSION_VOTE_SET_PROPOSAL_DURATION_PROPOSAL",
      "module": "gov",
      "description": "the permission needed to vote a proposal to set proposal duration."
    },
    {
      "id": 33,
      "name": "PERMISSION_BLACKLIST_ACCOUNT_PERMISSION_PROPOSAL",
      "module": "gov",
      "description": "the permission needed to create proposals for blacklisting an account permission."
    },
    {
      "id": 34,
      "name": "PERMISSION_VOTE_BLACKLIST_ACCOUNT_PERMISSION_PROPOSAL",
      "module": "gov",
      "description": "the permission that an actor must have in order to vote a proposal to blacklist account permission."
    },
    {
      "id": 35,
      "name": "PERMISSION_REMOVE_WHITELISTED_ACCOUNT_PERMISSION_PROPOSAL",
      "module": "gov",
      "description": "the permission needed to create proposals for removing whitelisted permission from an account."
    },
    {
      "id": 36,
      "name": "PERMISSION_VOTE_REMOVE_WHITELISTED_ACCOUNT_PERMISSION_PROPOSAL",
      "module": "gov",
      "description": "the permission that an actor must have in order to vote a proposal to remove a whitelisted account permission."
    },
    {
      "id": 37,
      "name": "PERMISSION_REMOVE_BLACKLISTED_ACCOUNT_PERMISSION_PROPOSAL",
      "module": "gov",
      "description": "the permission needed to create proposals for removing blacklisted permission from an account."
    },
    {
      "id": 38,
      "name": "PERMISSION_VOTE_REMOVE_BLACKLISTED_ACCOUNT_PERMISSION_PROPOSAL",
      "module": "gov",
      "description": "the permission that an actor must have in order to vote a proposal to remove a blacklisted account permission."
    },
    {
      "id": 39,
      "name": "PERMISSION_WHITELIST_ROLE_PERMISSION_PROPOSAL",
      "module": "gov",
      "description": "the permission needed to create proposals for whitelisting an role permission."
    },
    {
      "id": 40,
      "name": "PERMISSION_VOTE_WHITELIST_ROLE_PERMISSION_PROPOSAL",
      "module": "gov",
      "description": "the permission that an actor must have in order to vote a proposal to whitelist role permission."
    },
    {
      "id": 41,
      "name": "PERMISSION_BLACKLIST_ROLE_PERMISSION_PROPOSAL",
      "module": "gov",
      "description": "the permission needed to create proposals for blacklisting an role permission."
    },
    {
      "id": 42,
      "name": "PERMISSION_VOTE_BLACKLIST_ROLE_PERMISSION_PROPOSAL",
      "module": "gov",
      "description": "the permission that an actor must have in order to vote a proposal to blacklist role permission."
    },
    {
      "id": 43,
      "name": "PERMISSION_REMOVE_WHITELISTED_ROLE_PERMISSION_PROPOSAL",
      "module": "gov",
      "description": "the permission needed to create proposals for removing whitelisted permission from a role."
    },
    {
      "id": 44,
      "name": "PERMISSION_VOTE_REMOVE_WHITELISTED_ROLE_PERMISSION_PROPOSAL",
      "module": "gov",
      "description": "the permission that an actor must have in order to vote a proposal to remove a whitelisted role permission."
    },
    {
      "id": 45,
      "name": "PERMISSION_REMOVE_BLACKLISTED_ROLE_PERMISSION_PROPOSAL",
      "module": "gov",
      "description": "the permission needed to create proposals for removing blacklisted permission from a role."
    },
    {
      "id": 46,
      "name": "PERMISSION_VOTE_REMOVE_BLACKLISTED_ROLE_PERMISSION_PROPOSAL",
      "module": "gov",
      "description": "the permission that an actor must have in order to vote a proposal to remove a blacklisted role permission."
    },
    {
      "id": 47,
      "name": "PERMISSION_ASSIGN_ROLE_TO_ACCOUNT_PROPOSAL",
      "module": "gov",
      "description": "the permission needed to create proposals to assign role to an account."
    },
    {
      "id": 48,
      "name": "PERMISSION_VOTE_ASSIGN_ROLE_TO_ACCOUNT_PROPOSAL",
      "module": "gov",
      "description": "the permission that an actor must have in order to vote a proposal to assign role to an account"
    },
    {
      "id": 49,
      "name": "PERMISSION_UNASSIGN_ROLE_FROM_ACCOUNT_PROPOSAL",
      "module": "gov",
      "description": "the permission needed to create proposals to unassign role from an account"
    },
    {
      "id": 50,
      "name": "PERMISSION_VOTE_UNASSIGN_ROLE_FROM_ACCOUNT_PROPOSAL",
      "module": "gov",
      "description": "the permission that an actor must have in order to vote a proposal to unassign role from an account"
    },
    {
      "id": 51,
      "name": "PERMISSION_CREATE_REMOVE_ROLE_PROPOSAL",
      "module": "gov",
      "description": "the permission needed to create a proposal to remove a role."
    },
    {
      "id": 52,
      "name": "PERMISSION_VOTE_REMOVE_ROLE_PROPOSAL",
      "module": "gov",
      "description": "the permission needed to vote a proposal to remove a role."
    },
    {
      "id": 53,
      "name": "PERMISSION_CREATE_UPSERT_UBI_PROPOSAL",
      "module": "ubi",
      "description": "the permission needed to create proposals to upsert ubi."
    },
    {
      "id": 54,
      "name": "PERMISSION_VOTE_UPSERT_UBI_PROPOSAL",
      "module": "ubi",
      "description": "the permission that an actor must have in order to vote a proposal to upsert ubi"
    },
    {
      "id": 55,
      "name": "PERMISSION_CREATE_REMOVE_UBI_PROPOSAL",
      "module": "ubi",
      "description": "the permission needed to create a proposal to remove ubi"
    },
    {
      "id": 56,
      "name": "PERMISSION_VOTE_REMOVE_UBI_PROPOSAL",
      "module": "ubi",
      "description": "the permission needed to vote a proposal to remove ubi."
    },
    {
      "id": 57,
      "name": "PERMISSION_CREATE_SLASH_VALIDATOR_PROPOSAL",
      "module": "slashing",
      "description": "the permission needed to create a proposal to slash validator"
    },
    {
      "id": 58,
      "name": "PERMISSION_VOTE_SLASH_VALIDATOR_PROPOSAL",
      "module": "slashing",
      "description": "the permission needed to vote a proposal to slash validator"
    },
    {
      "id": 59,
      "name": "PERMISSION_CREATE_BASKET_PROPOSAL",
      "module": "basket",
      "description": "the permission needed to create a proposal related to basket"
    },
    {
      "id": 60,
      "name": "PERMISSION_VOTE_BASKET_PROPOSAL",
      "module": "basket",
      "description": "the permission needed to vote a proposal related to basket"
    },
    {
      "id": 61,
      "name": "PERMISSION_HANDLE_BASKET_EMERGENCY",
      "module": "basket",
      "description": "the permission needed to handle emergency issues on basket"
    },
    {
      "id": 62,
      "name": "PERMISSION_CREATE_RESET_WHOLE_COUNCILOR_RANK_PROPOSAL",
      "module": "gov",
      "description": "the permission needed to create a proposal to reset whole councilor rank"
    },
    {
      "id": 63,
      "name": "PERMISSION_VOTE_RESET_WHOLE_COUNCILOR_RANK_PROPOSAL",
      "module": "gov",
      "description": "the permission needed to vote on reset whole councilor rank proposal"
    },
    {
      "id": 64,
      "name": "PERMISSION_CREATE_JAIL_COUNCILOR_PROPOSAL",
      "module": "gov",
      "description": "the permission needed to create a proposal to jail councilors"
    },
    {
      "id": 65,
      "name": "PERMISSION_VOTE_JAIL_COUNCILOR_PROPOSAL",
      "module": "gov",
      "description": "the permission needed to vote on jail councilors proposal"
    }
  ],
  "properties": [
    {
      "name": "MinTxFee",
      "format": "uint64",
      "description": "Minimum transaction fee on the network"
    },
    {
      "name": "MaxTxFee",
      "format": "uint64",
      "description": "Maximum transaction fee on the network"
    },
    {
      "name": "VoteQuorum",
      "format": "uint64",
      "description": "Vote quorum to meet for a proposal to pass."
    },
    {
      "name": "MinimumProposalEndTime",
      "format": "uint64",
      "description": "Minimum proposal voting duration for a proposal should be live for voting."
    },
    {
      "name": "ProposalEnactmentTime",
      "format": "uint64",
      "description": "Proposal enactment time."
    },
    {
      "name": "MinProposalEndBlocks",
      "format": "uint64",
      "description": "Minimum number of blocks a proposal should be live for voting."
    },
    {
      "name": "MinProposalEnactmentBlocks",
      "format": "uint64",
      "description": "Minimum number of blocks a proposal should be in enactment."
    },
    {
      "name": "EnableForeignFeePayments",
      "format": "bool",
      "description": "Flag that describes foreign token (non-KEX) is enabled as fees."
    },
    {
      "name": "MischanceRankDecreaseAmount",
      "format": "uint64",
      "description": "Rank decrease amount when a validator miss a block."
    },
    {
      "name": "MaxMischance",
      "format": "uint64",
      "description": "Maximun number of sequencial miss on blocks before penalties."
    },
    {
      "name": "MischanceConfidence",
      "format": "uint64",
      "description": "Number of missed blocks accepted before increasing misschance."
    },
    {
      "name": "InactiveRankDecreasePercent",
      "format": "decimal",
      "description": "Percentage of rank decrease when a validator node become inactive."
    },
    {
      "name": "MinValidators",
      "format": "uint64",
      "description": "Number of active validators to be an active network."
    },
    {
      "name": "PoorNetworkMaxBankSend",
      "format": "uint64",
      "description": "Maximum number of tokens transferrable on poor network."
    },
    {
      "name": "UnjailMaxTime",
      "format": "uint64",
      "description": "Maximum time a validator can unjail after jail."
    },
    {
      "name": "EnableTokenWhitelist",
      "format": "bool",
      "description": "Flag to let only whitelisted tokens are transferrable"
    },
    {
      "name": "EnableTokenBlacklist",
      "format": "bool",
      "description": "Flag to prevent transfer of blacklisted tokens"
    },
    {
      "name": "MinIdentityApprovalTip",
      "format": "uint64",
      "description": "Minimum amount of tokens to be given for an identity record approval"
    },
    {
      "name": "UniqueIdentityKeys",
      "format": "string",
      "description": "Comma separated list of identity keys that should be unique across all identity keys"
    },
    {
      "name": "UbiHardcap",
      "format": "uint64",
      "description": "The maximum amount of tokens that can be allocated for sum of ubi records"
    },
    {
      "name": "ValidatorsFeeShare",
      "format": "decimal",
      "description": "The portion of fees to be given to the validator from the block fees"
    },
    {
      "name": "InflationRate",
      "format": "decimal",
      "description": "The rate of inflation during the inflation period"
    },
    {
      "name": "InflationPeriod",
      "format": "uint64",
      "description": "The duration unit for InflationRate"
    },
    {
      "name": "UnstakingPeriod",
      "format": "uint64",
      "description": "The unstaking duration on multistaking module"
    },
    {
      "name": "MaxDelegators",
      "format": "uint64",
      "description": "The maximum number of pool delegators for a single pool"
    },
    {
      "name": "MinDelegationPushout",
      "format": "uint64",
      "description": "The multiplier to push out a min delegation user when the maximum number of delegators filled in"
    },
    {
      "name": "SlashingPeriod",
      "format": "uint64",
      "description": "The period to take colluders on slash proposal"
    },
    {
      "name": "MaxJailedPercentage",
      "format": "decimal",
      "description": "The percentage of jails acceptable before slash proposal happens"
    },
    {
      "name": "MaxSlashingPercentage",
      "format": "decimal",
      "description": "The maximum slash percentage to for jail"
    },
    {
      "name": "MinCustodyReward",
      "format": "uint64",
      "description": "The minimum custody reward"
    },
    {
      "name": "MaxCustodyBufferSize",
      "format": "uint64",
      "description": "The minimum custody buffer size"
    },
    {
      "name": "MaxCustodyTxSize",
      "format": "uint64",
      "description": "The minimum custody transaction size"
    },
    {
      "name": "AbstentionRankDecreaseAmount",
      "format": "uint64",
      "description": "Rank decrease amount when a councilor does not participate in voting"
    },
    {
      "name": "MaxAbstention",
      "format": "uint64",
      "description": "The maximum absention count on voting for an active councilor"
    },
    {
      "name": "MinCollectiveBond",
      "format": "uint64",
      "description": "The minimum size of collective to be bootstrapped within bonding period"
    },
    {
      "name": "MinCollectiveBondingTime",
      "format": "uint64",
      "description": "The time to bootstrap minimum collectives bonds"
    },
    {
      "name": "MaxCollectiveOutputs",
      "format": "uint64",
      "description": "The maximum number of outputs a bonding pool could have"
    },
    {
      "name": "MinCollectiveClaimPeriod",
      "format": "uint64",
      "description": "The minimum acceptable collective claim period"
    }
  ],
  "transactions": {
    "Activate": {
      "function_id": 35,
      "description": "MsgActivate defines a message to activate an inactive validator.",
      "parameters": {
        "validator_addr": {
          "type": "string",
          "optional": false,
          "description": "bech32 format of validator address. e.g. kiravaloper1ewgq8gtsefakhal687t8hnsw5zl4y8eksup39w"
        }
      }
    },
    "AssignRole": {
      "function_id": 26,
      "description": "MsgAssignRole defines a message to assign a role to an address.",
      "parameters": {
        "address": {
          "type": "string",
          "optional": false,
          "description": "Address to set role to."
        },
        "proposer": {
          "type": "string",
          "optional": false,
          "description": "proposer who propose this message."
        },
        "role": {
          "type": "uint32",
          "optional": false,
          "description": "role identifier."
        }
      }
    },
    "BlacklistPermissions": {
      "function_id": 24,
      "description": "MsgBlacklistPermissions defines a message to blacklist permission of an address.",
      "parameters": {
        "address": {
          "type": "string",
          "optional": false,
          "description": "Address to blacklist permission to."
        },
        "permission": {
          "type": "uint32",
          "optional": false,
          "description": "Permission to be blacklisted."
        },
        "proposer": {
          "type": "string",
          "optional": false,
          "description": "proposer who propose this message."
        }
      }
    },
    "BlacklistRolePermission": {
      "function_id": 29,
      "description": "MsgBlacklistRolePermission defines a message to blacklist permission for a role.",
      "parameters": {
        "permission": {
          "type": "uint32",
          "optional": false,
          "description": "Permission to be blacklisted."
        },
        "proposer": {
          "type": "string",
          "optional": false,
          "description": "proposer who propose this message."
        },
        "role": {
          "type": "uint32",
          "optional": false,
          "description": "role identifier."
        }
      }
    },
    "CancelIdentityRecordsVerifyRequest": {
      "function_id": 16,
      "description": "MsgCancelIdentityRecordsVerifyRequest defines a proposal message to cancel an identity record request.",
      "parameters": {
        "executor": {
          "type": "string",
          "optional": false,
          "description": "the address of requester."
        },
        "verify_request_id": {
          "type": "uint64",
          "optional": false,
          "description": "the id of verification request."
        }
      }
    },
    "ClaimCouncilor": {
      "function_id": 22,
      "description": "MsgClaimCouncilor defines a message to claim councilor when the proposer.",
      "parameters": {
        "address": {
          "type": "string",
          "optional": false,
          "description": "Address to be set as councilor. This address should be proposer address as well."
        },
        "moniker": {
          "type": "string",
          "optional": false,
          "description": "validator's name or nickname."
        }
      }
    },
    "ClaimValidator": {
      "function_id": 32,
      "description": "MsgClaimValidator defines a message for claiming a new validator.",
      "parameters": {
        "moniker": {
          "type": "string",
          "optional": false,
          "description": "validator's name or nickname."
        },
        "pub_key": {
          "type": "string",
          "optional": false,
          "description": "validator bech32 public key"
        },
        "val_key": {
          "type": "val_address",
          "optional": false,
          "description": "validator operator address"
        }
      }
    },
    "CreateRole": {
      "function_id": 25,
      "description": "MsgCreateRole defines a message to create a role.",
      "parameters": {
        "proposer": {
          "type": "string",
          "optional": false,
          "description": "proposer who propose this message."
        },
        "role": {
          "type": "uint32",
          "optional": false,
          "description": "Identifier of this role."
        }
      }
    },
    "CreateSpendingPool": {
      "function_id": 41,
      "description": "MsgCreateSpendingPool represents a message to create a spending pool.",
      "parameters": {
        "beneficiaries": {
          "type": "PermInfo",
          "optional": false,
          "description": ""
        },
        "claim_end": {
          "type": "time",
          "optional": false,
          "description": ""
        },
        "claim_start": {
          "type": "time",
          "optional": false,
          "description": ""
        },
        "expire": {
          "type": "uint64",
          "optional": false,
          "description": ""
        },
        "name": {
          "type": "string",
          "optional": false,
          "description": ""
        },
        "owners": {
          "type": "PermInfo",
          "optional": false,
          "description": ""
        },
        "rate": {
          "type": "decimal",
          "optional": false,
          "description": ""
        },
        "sender": {
          "type": "string",
          "optional": false,
          "description": ""
        },
        "token": {
          "type": "string",
          "optional": false,
          "description": ""
        },
        "vote_enactment": {
          "type": "uint64",
          "optional": false,
          "description": ""
        },
        "vote_period": {
          "type": "uint64",
          "optional": false,
          "description": ""
        },
        "vote_quorum": {
          "type": "uint64",
          "optional": false,
          "description": ""
        }
      }
    },
    "EditIdentityRecord": {
      "function_id": 13,
      "description": "MsgDeleteIdentityRecords defines a method to delete identity records owned by an address.",
      "parameters": {
        "address": {
          "type": "string",
          "optional": false,
          "description": "the address of requester."
        },
        "keys": {
          "type": "array",
          "optional": false,
          "description": "the array string that defines identity record key values to be deleted."
        }
      }
    },
    "HandleIdentityRecordsVerifyRequest": {
      "function_id": 15,
      "description": "MsgHandleIdentityRecordsVerifyRequest defines a proposal message to approve or reject an identity record request.",
      "parameters": {
        "verifier": {
          "type": "string",
          "optional": false,
          "description": "the address of verifier."
        },
        "verify_request_id": {
          "type": "uint64",
          "optional": false,
          "description": "the id of verification request."
        },
        "yes": {
          "type": "bool",
          "optional": true,
          "description": "defines approval or rejecting an identity request (default false)"
        }
      }
    },
    "Pause": {
      "function_id": 36,
      "description": "MsgRefuteSlashingProposal defines a message to refute a validator slash proposal.",
      "parameters": {
        "refutation": {
          "type": "string",
          "optional": false,
          "description": "refutation link of on the proposal"
        },
        "sender": {
          "type": "string",
          "optional": false,
          "description": "proposer of the message."
        },
        "validator": {
          "type": "string",
          "optional": false,
          "description": "bech32 format of validator address. e.g. kiravaloper1ewgq8gtsefakhal687t8hnsw5zl4y8eksup39w"
        }
      }
    },
    "RegisterIdentityRecords": {
      "function_id": 12,
      "description": "MsgRegisterIdentityRecords defines a proposal message to create a identity record.",
      "parameters": {
        "address": {
          "type": "string",
          "optional": false,
          "description": "the address for the identity record."
        },
        "infos": {
          "type": "array",
          "optional": false,
          "description": "key/value array for the mappings of the identity record."
        }
      }
    },
    "RemoveBlacklistRolePermission": {
      "function_id": 31,
      "description": "MsgRemoveBlacklistRolePermission defines a message to remove blacklisted permission for a role.",
      "parameters": {
        "permission": {
          "type": "uint32",
          "optional": false,
          "description": "Permission to be removed from blacklisted listing."
        },
        "proposer": {
          "type": "string",
          "optional": false,
          "description": "proposer who propose this message."
        },
        "role": {
          "type": "uint32",
          "optional": false,
          "description": "role identifier."
        }
      }
    },
    "RemoveRole": {
      "function_id": 27,
      "description": "MsgRemoveRole defines a message to remove a role from an address.",
      "parameters": {
        "address": {
          "type": "string",
          "optional": false,
          "description": "Address to remove role from."
        },
        "proposer": {
          "type": "string",
          "optional": false,
          "description": "proposer who propose this message."
        },
        "role": {
          "type": "uint32",
          "optional": false,
          "description": "role identifier."
        }
      }
    },
    "RemoveWhitelistRolePermission": {
      "function_id": 30,
      "description": "MsgRemoveWhitelistRolePermission defines a message to remove whitelisted permission for a role.",
      "parameters": {
        "permission": {
          "type": "uint32",
          "optional": false,
          "description": "Permission to be removed from whitelisted listing."
        },
        "proposer": {
          "type": "string",
          "optional": false,
          "description": "proposer who propose this message."
        },
        "role": {
          "type": "uint32",
          "optional": false,
          "description": "role identifier."
        }
      }
    },
    "RequestIdentityRecordsVerify": {
      "function_id": 14,
      "description": "MsgRequestIdentityRecordsVerify defines a proposal message to request an identity record verification from a specific verifier.",
      "parameters": {
        "address": {
          "type": "string",
          "optional": false,
          "description": "the address of requester."
        },
        "record_ids": {
          "type": "array<uint64>",
          "optional": false,
          "description": "the id of records to be verified."
        },
        "tip": {
          "type": "coins",
          "optional": false,
          "description": "the amount of coins to be given up-on accepting the request."
        },
        "verifier": {
          "type": "string",
          "optional": false,
          "description": "the address of verifier."
        }
      }
    },
    "SetExecutionFee": {
      "function_id": 21,
      "description": "MsgSetExecutionFee defines a message to set execution fee with specific permission.",
      "parameters": {
        "default_parameters": {
          "type": "bool",
          "optional": false,
          "description": "Default values that the function in question will consume as input parameters before execution"
        },
        "execution_fee": {
          "type": "uint64",
          "optional": false,
          "description": "How much user should pay for executing this specific function"
        },
        "failure_fee": {
          "type": "uint64",
          "optional": false,
          "description": "How much user should pay if function fails to execute"
        },
        "proposer": {
          "type": "address",
          "optional": false,
          "description": "proposer who propose this message."
        },
        "timeout": {
          "type": "uint64",
          "optional": false,
          "description": "After what time function execution should fail"
        },
        "transaction_type": {
          "type": "string",
          "optional": false,
          "description": "Type of the transaction that given permission allows to execute"
        }
      }
    },
    "SetNetworkProperties": {
      "function_id": 20,
      "description": "MsgSetNetworkProperties defines a message to set network properties with specific permission.",
      "parameters": {
        "network_properties": {
          "type": "<Object>",
          "optional": false,
          "description": "network properties to be set.",
          "fields": {
            "enable_foreign_fee_payments": {
              "type": "bool",
              "optional": false,
              "description": "flag to show if foreign fee payment is enabled"
            },
            "max_tx_fee": {
              "type": "uint64",
              "optional": false,
              "description": "maximum transaction fee"
            },
            "min_tx_fee": {
              "type": "uint64",
              "optional": false,
              "description": "minimum transaction fee"
            },
            "proposal_enactment_time": {
              "type": "uint64",
              "optional": false,
              "description": "proposal enactment time"
            },
            "proposal_end_time": {
              "type": "uint64",
              "optional": false,
              "description": "proposal end time"
            },
            "vote_quorum": {
              "type": "uint64",
              "optional": false,
              "description": "vote quorum"
            }
          }
        },
        "proposer": {
          "type": "address",
          "optional": false,
          "description": "proposer who propose this message."
        }
      }
    },
    "SubmitEvidence": {
      "function_id": 3,
      "description": "MsgSubmitEvidence defines a message to submit an evidence",
      "parameters": {
        "evidence": {
          "type": "object",
          "optional": false,
          "description": "evidence object"
        },
        "submitter": {
          "type": "string",
          "optional": false,
          "description": "evidence submitter address"
        }
      }
    },
    "SubmitProposal": {
      "function_id": 10,
      "description": "MsgSubmitProposal defines a proposal message to submit a proposal.",
      "parameters": {
        "content": {
          "type": "object",
          "optional": false,
          "description": "the content of the proposal - different by type of proposal"
        },
        "description": {
          "type": "string",
          "optional": false,
          "description": "the description of the proposal."
        },
        "proposer": {
          "type": "address",
          "optional": false,
          "description": "the proposer of the proposal."
        },
        "title": {
          "type": "string",
          "optional": false,
          "description": "the title of the proposal."
        }
      }
    },
    "Unpause": {
      "function_id": 37,
      "description": "MsgUnpause defines a message to unpause a paused validator.",
      "parameters": {
        "validator_addr": {
          "type": "string",
          "optional": false,
          "description": "bech32 format of validator address. e.g. kiravaloper1ewgq8gtsefakhal687t8hnsw5zl4y8eksup39w"
        }
      }
    },
    "UpsertTokenAlias": {
      "function_id": 33,
      "description": "MsgUpsertTokenAlias represents a message to register token alias.",
      "parameters": {
        "decimals": {
          "type": "uint32",
          "optional": false,
          "description": "Integer number of max decimals."
        },
        "denoms": {
          "type": "array<string>",
          "optional": false,
          "description": "An array of token denoms to be aliased."
        },
        "icon": {
          "type": "string",
          "optional": false,
          "description": "Graphical Symbol (url link to graphics)."
        },
        "name": {
          "type": "string",
          "optional": false,
          "description": "Token Name (e.g. Cosmos, Kira, Bitcoin)."
        },
        "proposer": {
          "type": "string",
          "optional": false,
          "description": "proposer who propose this message."
        },
        "symbol": {
          "type": "string",
          "optional": false,
          "description": "Ticker (eg. ATOM, KEX, BTC)."
        }
      }
    },
    "UpsertTokenRate": {
      "function_id": 34,
      "description": "MsgUpsertTokenRate represents a message to register token rate.",
      "parameters": {
        "denom": {
          "type": "string",
          "optional": false,
          "description": "denomination target."
        },
        "fee_payments": {
          "type": "bool",
          "optional": false,
          "description": "defining if it is enabled or disabled as fee payment method."
        },
        "proposer": {
          "type": "address",
          "optional": false,
          "description": "proposer who propose this message."
        },
        "rate": {
          "type": "float",
          "optional": false,
          "description": "Exchange rate in terms of KEX token. e.g. 0.1, 10.5"
        }
      }
    },
    "VoteProposal": {
      "function_id": 11,
      "description": "MsgVoteProposal defines a proposal message to vote on a submitted proposal.",
      "parameters": {
        "proposal_id": {
          "type": "uint64",
          "optional": false,
          "description": "id of proposal to be voted."
        },
        "value": {
          "type": "enum<VoteOption>",
          "optional": false,
          "description": "vote option: [yes, no, veto, abstain]"
        },
        "voter": {
          "type": "address",
          "optional": false,
          "description": "the address of the voter who vote on the proposal."
        }
      }
    },
    "WhitelistPermissions": {
      "function_id": 23,
      "description": "MsgWhitelistPermissions defines a message to whitelist permission of an address.",
      "parameters": {
        "address": {
          "type": "string",
          "optional": false,
          "description": "Address to whitelist permission to."
        },
        "permission": {
          "type": "uint32",
          "optional": false,
          "description": "Permission to be whitelisted."
        },
        "proposer": {
          "type": "string",
          "optional": false,
          "description": "proposer who propose this message."
        }
      }
    },
    "WhitelistRolePermission": {
      "function_id": 28,
      "description": "MsgWhitelistRolePermission defines a message to whitelist permission for a role.",
      "parameters": {
        "permission": {
          "type": "uint32",
          "optional": false,
          "description": "Permission to be whitelisted."
        },
        "proposer": {
          "type": "string",
          "optional": false,
          "description": "proposer who propose this message."
        },
        "role": {
          "type": "uint32",
          "optional": false,
          "description": "role identifier."
        }
      }
    }
  }
}

```
</details>


[Return to top](#sekai)

### 7. export-minimized-genesis

Get minimized genesis from not standard genesis.

Usage:

```
sekaid export-minimized-genesis [path-to-exported.json] [path-to-standard.json] [flags]
```

| Flags      | Description   | Work  |
| ---------- | ------------- | ----- |
| -h, --help | help for init | ✅ yes |



| Global Flags       |                                                                                    |      |
| ------------------ | ---------------------------------------------------------------------------------- | ---- |
| --home string      | directory for config and data (default "/root/.sekaid")                            | ❌ no |
| --log_format       | The logging format (json\|plain) (default "plain")                                 | ❌ no |
| --log_level string | The logging level (trace\|debug\|info\|warn\|error\|fatal\|panic) (default "info") | ❌ ?  |
| --trace            | print out full stack trace on errors                                               | ❌ ?  |



```
# sekaid export-minimized-genesis -h
Get minimized genesis from not standard genesis.
Example:
$ sekaid export-minimized-genesis exported-genesis.json new-genesis.json

Usage:
  sekaid export-minimized-genesis [path-to-exported.json] [path-to-standard.json] [flags]

Flags:
  -h, --help   help for export-minimized-genesis

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid export-minimized-genesis ~/.sekai/config/genesis.json minimizedGenesis.json
```
<details>
  <summary>Check export-minimized-genesis out</summary>

```
# sekaid export-minimized-genesis ~/.sekai/config/genesis.json minimizedGenesis.json && cat minimizedGenesis.json | jq 
{
  "genesis_time": "2023-05-20T08:48:00.848282749Z",
  "chain_id": "localnet-47",
  "initial_height": "1",
  "consensus_params": {
    "block": {
      "max_bytes": "22020096",
      "max_gas": "-1",
      "time_iota_ms": "1000"
    },
    "evidence": {
      "max_age_num_blocks": "100000",
      "max_age_duration": "172800000000000",
      "max_bytes": "1048576"
    },
    "validator": {
      "pub_key_types": [
        "ed25519"
      ]
    },
    "version": {}
  },
  "app_hash": "",
  "app_state": {
    "auth": {
      "params": {
        "max_memo_characters": "256",
        "tx_sig_limit": "7",
        "tx_size_cost_per_byte": "10",
        "sig_verify_cost_ed25519": "590",
        "sig_verify_cost_secp256k1": "1000"
      },
      "accounts": [
        {
          "@type": "/cosmos.auth.v1beta1.BaseAccount",
          "address": "kira1tcyq0y66mmdpydepexc8lwrmemp9wzgmg2rqhm",
          "pub_key": null,
          "account_number": "0",
          "sequence": "0"
        },
        {
          "@type": "/cosmos.auth.v1beta1.BaseAccount",
          "address": "kira1m2m98e2fahrzvlhdc6rvf645plzt2e626knhn8",
          "pub_key": null,
          "account_number": "0",
          "sequence": "0"
        },
        {
          "@type": "/cosmos.auth.v1beta1.BaseAccount",
          "address": "kira1hcj5m0y3gmr6d80funkn6zj8e4amye3wk5jdmk",
          "pub_key": null,
          "account_number": "0",
          "sequence": "0"
        },
        {
          "@type": "/cosmos.auth.v1beta1.BaseAccount",
          "address": "kira1wh5zkqgg87tes5r6ycp5r56qgyvzglhply5yu4",
          "pub_key": null,
          "account_number": "0",
          "sequence": "0"
        },
        {
          "@type": "/cosmos.auth.v1beta1.BaseAccount",
          "address": "kira16v3lv7qasyn22dad77t4lfg4mqrvpnqr9swylt",
          "pub_key": null,
          "account_number": "0",
          "sequence": "0"
        }
      ]
    },
    "bank": {
      "params": {
        "send_enabled": [],
        "default_send_enabled": true
      },
      "balances": [
        {
          "address": "kira1tcyq0y66mmdpydepexc8lwrmemp9wzgmg2rqhm",
          "coins": [
            {
              "denom": "lol",
              "amount": "1000000"
            },
            {
              "denom": "samolean",
              "amount": "2000000000000000000000000000"
            },
            {
              "denom": "test",
              "amount": "29999780000000000"
            },
            {
              "denom": "ukex",
              "amount": "299998800000000"
            }
          ]
        },
        {
          "address": "kira1wh5zkqgg87tes5r6ycp5r56qgyvzglhply5yu4",
          "coins": [
            {
              "denom": "lol",
              "amount": "1000000"
            },
            {
              "denom": "samolean",
              "amount": "2000000000000000000000000000"
            },
            {
              "denom": "test",
              "amount": "300000000000000"
            },
            {
              "denom": "ukex",
              "amount": "150000000000000"
            }
          ]
        },
        {
          "address": "kira1hcj5m0y3gmr6d80funkn6zj8e4amye3wk5jdmk",
          "coins": [
            {
              "denom": "lol",
              "amount": "1000000"
            },
            {
              "denom": "samolean",
              "amount": "3000000000000000000000000000"
            },
            {
              "denom": "test",
              "amount": "200000000000"
            },
            {
              "denom": "ukex",
              "amount": "1000000000"
            }
          ]
        },
        {
          "address": "kira16v3lv7qasyn22dad77t4lfg4mqrvpnqr9swylt",
          "coins": [
            {
              "denom": "lol",
              "amount": "1000000"
            },
            {
              "denom": "samolean",
              "amount": "2000000000000000000000000000"
            },
            {
              "denom": "test",
              "amount": "300000000000000"
            },
            {
              "denom": "ukex",
              "amount": "150000000000000"
            }
          ]
        },
        {
          "address": "kira1m2m98e2fahrzvlhdc6rvf645plzt2e626knhn8",
          "coins": [
            {
              "denom": "test",
              "amount": "10000000000"
            },
            {
              "denom": "ukex",
              "amount": "100000000"
            }
          ]
        }
      ],
      "supply": [],
      "denom_metadata": []
    },
    "basket": {
      "baskets": [],
      "historical_mints": [],
      "historical_burns": [],
      "historical_swaps": []
    },
    "collectives": {
      "collectives": [],
      "contributers": []
    },
    "custody": {},
    "customevidence": {
      "evidence": []
    },
    "customgov": {
      "starting_proposal_id": "1",
      "next_role_id": "3",
      "roles": [
        {
          "id": 1,
          "sid": "sudo",
          "description": "Sudo role"
        },
        {
          "id": 2,
          "sid": "validator",
          "description": "Validator role"
        }
      ],
      "role_permissions": {
        "1": {
          "blacklist": [],
          "whitelist": [
            1,
            2,
            3,
            6,
            8,
            9,
            12,
            13,
            10,
            11,
            14,
            15,
            18,
            19,
            20,
            21,
            22,
            23,
            31,
            32,
            24,
            25,
            16,
            17,
            4,
            5,
            26,
            27,
            28,
            29,
            30,
            33,
            34,
            35,
            36,
            37,
            38,
            39,
            40,
            41,
            42,
            43,
            44,
            45,
            46,
            47,
            48,
            49,
            50,
            51,
            52,
            53,
            54,
            55,
            56,
            57,
            58,
            59,
            60,
            61,
            62,
            63,
            64,
            65,
            66
          ]
        },
        "2": {
          "blacklist": [],
          "whitelist": [
            2
          ]
        }
      },
      "network_actors": [
        {
          "address": "kira1tcyq0y66mmdpydepexc8lwrmemp9wzgmg2rqhm",
          "roles": [
            "1"
          ],
          "status": "ACTIVE",
          "votes": [
            "VOTE_OPTION_YES",
            "VOTE_OPTION_ABSTAIN",
            "VOTE_OPTION_NO",
            "VOTE_OPTION_NO_WITH_VETO"
          ],
          "permissions": {
            "blacklist": [],
            "whitelist": []
          },
          "skin": "1"
        }
      ],
      "network_properties": {
        "min_tx_fee": "100",
        "max_tx_fee": "1000000",
        "vote_quorum": "33",
        "minimum_proposal_end_time": "360",
        "proposal_enactment_time": "300",
        "min_proposal_end_blocks": "2",
        "min_proposal_enactment_blocks": "1",
        "enable_foreign_fee_payments": true,
        "mischance_rank_decrease_amount": "1",
        "max_mischance": "50",
        "mischance_confidence": "25",
        "inactive_rank_decrease_percent": "0.500000000000000000",
        "min_validators": "1",
        "poor_network_max_bank_send": "1000000",
        "unjail_max_time": "1209600",
        "enable_token_whitelist": false,
        "enable_token_blacklist": true,
        "min_identity_approval_tip": "200",
        "unique_identity_keys": "moniker,username",
        "ubi_hardcap": "6000000",
        "validators_fee_share": "0.500000000000000000",
        "inflation_rate": "0.180000000000000000",
        "inflation_period": "31557600",
        "unstaking_period": "2629800",
        "max_delegators": "100",
        "min_delegation_pushout": "10",
        "slashing_period": "3600",
        "max_jailed_percentage": "0.250000000000000000",
        "max_slashing_percentage": "0.010000000000000000",
        "min_custody_reward": "200",
        "max_custody_buffer_size": "10",
        "max_custody_tx_size": "8192",
        "abstention_rank_decrease_amount": "1",
        "max_abstention": "2",
        "min_collective_bond": "100000",
        "min_collective_bonding_time": "86400",
        "max_collective_outputs": "10",
        "min_collective_claim_period": "14400",
        "validator_recovery_bond": "300000",
        "max_annual_inflation": "0.350000000000000000",
        "max_proposal_title_size": "128",
        "max_proposal_description_size": "1024",
        "max_proposal_poll_option_size": "64",
        "max_proposal_poll_option_count": "128",
        "max_proposal_reference_size": "512",
        "max_proposal_checksum_size": "128",
        "min_dapp_bond": "1000000",
        "max_dapp_bond": "10000000",
        "dapp_liquidation_threshold": "0",
        "dapp_liquidation_period": "0",
        "dapp_bond_duration": "604800",
        "dapp_verifier_bond": "0.001000000000000000",
        "dapp_auto_denounce_time": "60",
        "dapp_mischance_rank_decrease_amount": "1",
        "dapp_max_mischance": "10",
        "dapp_inactive_rank_decrease_percent": "10",
        "dapp_pool_slippage_default": "0.100000000000000000",
        "minting_ft_fee": "100000000000000",
        "minting_nft_fee": "100000000000000"
      },
      "execution_fees": [
        {
          "transaction_type": "claim-validator",
          "execution_fee": "100",
          "failure_fee": "1",
          "timeout": "10",
          "default_parameters": "0"
        },
        {
          "transaction_type": "claim-councilor",
          "execution_fee": "100",
          "failure_fee": "1",
          "timeout": "10",
          "default_parameters": "0"
        },
        {
          "transaction_type": "claim-proposal-type-x",
          "execution_fee": "100",
          "failure_fee": "1",
          "timeout": "10",
          "default_parameters": "0"
        },
        {
          "transaction_type": "vote-proposal-type-x",
          "execution_fee": "100",
          "failure_fee": "1",
          "timeout": "10",
          "default_parameters": "0"
        },
        {
          "transaction_type": "submit-proposal-type-x",
          "execution_fee": "10",
          "failure_fee": "1",
          "timeout": "10",
          "default_parameters": "0"
        },
        {
          "transaction_type": "veto-proposal-type-x",
          "execution_fee": "100",
          "failure_fee": "1",
          "timeout": "10",
          "default_parameters": "0"
        },
        {
          "transaction_type": "upsert-token-alias",
          "execution_fee": "100",
          "failure_fee": "1",
          "timeout": "10",
          "default_parameters": "0"
        },
        {
          "transaction_type": "activate",
          "execution_fee": "100",
          "failure_fee": "1000",
          "timeout": "10",
          "default_parameters": "0"
        },
        {
          "transaction_type": "pause",
          "execution_fee": "100",
          "failure_fee": "100",
          "timeout": "10",
          "default_parameters": "0"
        },
        {
          "transaction_type": "unpause",
          "execution_fee": "100",
          "failure_fee": "100",
          "timeout": "10",
          "default_parameters": "0"
        }
      ],
      "poor_network_messages": {
        "messages": [
          "submit-proposal",
          "set-network-properties",
          "vote-proposal",
          "claim-councilor",
          "whitelist-permissions",
          "blacklist-permissions",
          "create-role",
          "assign-role",
          "remove-role",
          "whitelist-role-permission",
          "blacklist-role-permission",
          "remove-whitelist-role-permission",
          "remove-blacklist-role-permission",
          "claim-validator",
          "activate",
          "pause",
          "unpause",
          "register-identity-records",
          "edit-identity-record",
          "request-identity-records-verify",
          "handle-identity-records-verify-request",
          "cancel-identity-records-verify-request"
        ]
      },
      "proposals": [],
      "votes": [],
      "data_registry": {},
      "identity_records": [
        {
          "id": "1",
          "address": "kira1tcyq0y66mmdpydepexc8lwrmemp9wzgmg2rqhm",
          "key": "moniker",
          "value": "GENESIS VALIDATOR",
          "date": "2023-05-20T08:48:01.608922330Z",
          "verifiers": []
        }
      ],
      "last_identity_record_id": "1",
      "id_records_verify_requests": [],
      "last_id_record_verify_request_id": "0",
      "proposal_durations": {}
    },
    "customslashing": {
      "params": {
        "downtime_inactive_duration": "600s"
      },
      "signing_infos": []
    },
    "customstaking": {
      "validators": [
        {
          "val_key": "kiravaloper1tcyq0y66mmdpydepexc8lwrmemp9wzgmmvlr0h",
          "pub_key": {
            "@type": "/cosmos.crypto.ed25519.PubKey",
            "key": "T6DNclgBnpmdXM4mduAoVwP5T8U2E+fg749Roig9wPI="
          },
          "status": "ACTIVE",
          "rank": "0",
          "streak": "0"
        }
      ]
    },
    "distributor": {
      "fees_treasury": [],
      "fees_collected": [],
      "snap_period": "1000",
      "validator_votes": [],
      "previous_proposer": "",
      "year_start_snapshot": {
        "snapshot_time": "0",
        "snapshot_amount": "0"
      }
    },
    "feeprocessing": {},
    "genutil": {
      "gen_txs": []
    },
    "layer2": {
      "dapps": [],
      "bridge": {
        "helper": null,
        "accounts": [],
        "tokens": [],
        "xams": []
      }
    },
    "multistaking": null,
    "params": null,
    "recovery": {
      "recovery_records": [],
      "recovery_tokens": [],
      "rewards": [],
      "rotations": []
    },
    "spending": {
      "pools": [
        {
          "name": "ValidatorBasicRewardsPool",
          "claim_start": "0",
          "claim_end": "0",
          "claim_expiry": "0",
          "rates": [
            {
              "denom": "ukex",
              "amount": "385.000000000000000000"
            }
          ],
          "vote_quorum": "33",
          "vote_period": "300",
          "vote_enactment": "300",
          "owners": {
            "owner_roles": [
              "2"
            ],
            "owner_accounts": []
          },
          "beneficiaries": {
            "roles": [
              {
                "role": "2",
                "weight": "1"
              }
            ],
            "accounts": []
          },
          "balances": [],
          "dynamic_rate": false,
          "dynamic_rate_period": "0",
          "last_dynamic_rate_calc_time": "0"
        }
      ],
      "claims": []
    },
    "tokens": {
      "aliases": [
        {
          "symbol": "KEX",
          "name": "Kira",
          "icon": "",
          "decimals": 6,
          "denoms": [
            "ukex",
            "mkex"
          ],
          "invalidated": false
        }
      ],
      "rates": [
        {
          "denom": "ukex",
          "fee_rate": "1.000000000000000000",
          "fee_payments": true,
          "stake_cap": "0.500000000000000000",
          "stake_min": "1",
          "stake_token": true,
          "invalidated": false
        },
        {
          "denom": "ubtc",
          "fee_rate": "10.000000000000000000",
          "fee_payments": true,
          "stake_cap": "0.250000000000000000",
          "stake_min": "1",
          "stake_token": true,
          "invalidated": false
        },
        {
          "denom": "xeth",
          "fee_rate": "0.100000000000000000",
          "fee_payments": true,
          "stake_cap": "0.100000000000000000",
          "stake_min": "1",
          "stake_token": false,
          "invalidated": false
        },
        {
          "denom": "frozen",
          "fee_rate": "0.100000000000000000",
          "fee_payments": true,
          "stake_cap": "0.000000000000000000",
          "stake_min": "1",
          "stake_token": false,
          "invalidated": false
        }
      ],
      "tokenBlackWhites": {
        "whitelisted": [
          "ukex"
        ],
        "blacklisted": [
          "frozen"
        ]
      }
    },
    "ubi": {
      "ubi_records": [
        {
          "name": "ValidatorBasicRewardsPoolUBI",
          "distribution_start": "0",
          "distribution_end": "0",
          "distribution_last": "0",
          "amount": "500000",
          "period": "2592000",
          "pool": "ValidatorBasicRewardsPool",
          "dynamic": true
        }
      ]
    },
    "upgrade": {
      "version": "",
      "current_plan": null,
      "next_plan": null
    }
  }
}
```
  
</details>

[Return to top](#sekai)

### 8. gentx

[Return to top](#sekai)

### 9. gentx-claim

Adds validator into the genesis set

Usage:

```
sekaid gentx-claim [key_name] [flags]
```

| Flags                    | Description                                                             | Work  |
| ------------------------ | ----------------------------------------------------------------------- | ----- |
| -h, --help               | help for gentx-claim                                                    | ✅ yes |
| --keyring-backend string | Select keyring's backend (os\|file\|kwallet\|pass\|test) (default "os") | ✅ yes |
| --moniker string         | the Moniker                                                             | ✅ yes |
| --pubkey string          | the public key                                                          | ❌ ?   |



| Global Flags       |                                                                                    |       |
| ------------------ | ---------------------------------------------------------------------------------- | ----- |
| --home string      | directory for config and data (default "/root/.sekaid")                            | ✅ yes |
| --log_format       | The logging format (json\|plain) (default "plain")                                 | ❌ no  |
| --log_level string | The logging level (trace\|debug\|info\|warn\|error\|fatal\|panic) (default "info") | ❌ ?   |
| --trace            | print out full stack trace on errors                                               | ❌ ?   |

```
# sekaid gentx-claim -h
Adds validator into the genesis set

Usage:
  sekaid gentx-claim [key_name] [flags]

Flags:
  -h, --help                     help for gentx-claim
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test) (default "os")
      --moniker string           the Moniker
      --pubkey string            the public key

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
# sekaid gentx-claim validator --keyring-backend=test --moniker="GENESIS VALIDATOR" --home=$SEKAID_HOME
genesis state updated to include validator
```

[Return to top](#sekai)

### 10. help

Help provides help for any command in the application.
Simply type sekaid help [path to command] for full details.

Usage:

```
sekaid help [command] [flags]
```
| Flags      | Description   | Work  |
| ---------- | ------------- | ----- |
| -h, --help | help for init | ✅ yes |



| Global Flags       |                                                                                    |      |
| ------------------ | ---------------------------------------------------------------------------------- | ---- |
| --home string      | directory for config and data (default "/root/.sekaid")                            | ❌ no |
| --log_format       | The logging format (json\|plain) (default "plain")                                 | ❌ no |
| --log_level string | The logging level (trace\|debug\|info\|warn\|error\|fatal\|panic) (default "info") | ❌ ?  |
| --trace            | print out full stack trace on errors                                               | ❌ ?  |

```
# sekaid help -h
Help provides help for any command in the application.
Simply type sekaid help [path to command] for full details.

Usage:
  sekaid help [command] [flags]

Flags:
  -h, --help   help for help

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors

```

```
# sekaid help version
Print the application binary version information

Usage:
  sekaid version [flags]

Flags:
  -h, --help            help for version
      --long            Print long version information
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

[Return to top](#sekai)

### 11. init
Initialize validators's and node's configuration files.

Usage:
```
sekaid init [moniker] [flags]
```
| Flags             | Description                                                     | Work  |
| ----------------- | --------------------------------------------------------------- | ----- |
| --chain-id string | genesis file chain-id, if left blank will be randomly created   | ✅ yes |
| -h, --help        | help for init                                                   | ✅ yes |
| -o --overwrite    | overwrite the genesis.json file                                 | ✅ yes |
| --recover         | provide seed phrase to recover existing key instead of creating | ❌ ?   |



| Global Flags       |                                                                                    |       |
| ------------------ | ---------------------------------------------------------------------------------- | ----- |
| --home string      | directory for config and data (default "/root/.sekaid")                            | ✅ yes |
| --log_format       | The logging format (json\|plain) (default "plain")                                 | ❌ no  |
| --log_level string | The logging level (trace\|debug\|info\|warn\|error\|fatal\|panic) (default "info") | ❌ ?   |
| --trace            | print out full stack trace on errors                                               | ❌ ?   |

```
# sekaid init --help
Initialize validators's and node's configuration files.

Usage:
  sekaid init [moniker] [flags]

Flags:
      --chain-id string   genesis file chain-id, if left blank will be randomly created
  -h, --help              help for init
  -o, --overwrite         overwrite the genesis.json file
      --recover           provide seed phrase to recover existing key instead of creating

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errorsdoc 
```

```
#sekaid init "MONIKER" --overwrite --chain-id=MyBlockchain-03  --home=~/.sekaid-MyBlockchain-03
```
<details>
<summary>Click to expand</summary>

```
#sekaid init "MONIKER" --overwrite --chain-id=MyBlockchain-03  --home=~/.sekaid-MyBlockchain-03
{"app_message":{"auth":{"accounts":[],"params":{"max_memo_characters":"256","sig_verify_cost_ed25519":"590","sig_verify_cost_secp256k1":"1000","tx_sig_limit":"7","tx_size_cost_per_byte":"10"}},"bank":{"balances":[],"denom_metadata":[],"params":{"default_send_enabled":true,"send_enabled":[]},"supply":[]},"basket":{"baskets":[],"historical_burns":[],"historical_mints":[],"historical_swaps":[]},"collectives":{"collectives":[],"contributers":[]},"custody":{},"customevidence":{"evidence":[]},"customgov":{"data_registry":{},"execution_fees":[{"default_parameters":"0","execution_fee":"100","failure_fee":"1","timeout":"10","transaction_type":"claim-validator"},{"default_parameters":"0","execution_fee":"100","failure_fee":"1","timeout":"10","transaction_type":"claim-councilor"},{"default_parameters":"0","execution_fee":"100","failure_fee":"1","timeout":"10","transaction_type":"claim-proposal-type-x"},{"default_parameters":"0","execution_fee":"100","failure_fee":"1","timeout":"10","transaction_type":"vote-proposal-type-x"},{"default_parameters":"0","execution_fee":"10","failure_fee":"1","timeout":"10","transaction_type":"submit-proposal-type-x"},{"default_parameters":"0","execution_fee":"100","failure_fee":"1","timeout":"10","transaction_type":"veto-proposal-type-x"},{"default_parameters":"0","execution_fee":"100","failure_fee":"1","timeout":"10","transaction_type":"upsert-token-alias"},{"default_parameters":"0","execution_fee":"100","failure_fee":"1000","timeout":"10","transaction_type":"activate"},{"default_parameters":"0","execution_fee":"100","failure_fee":"100","timeout":"10","transaction_type":"pause"},{"default_parameters":"0","execution_fee":"100","failure_fee":"100","timeout":"10","transaction_type":"unpause"}],"id_records_verify_requests":[],"identity_records":[],"last_id_record_verify_request_id":"0","last_identity_record_id":"0","network_actors":[],"network_properties":{"abstention_rank_decrease_amount":"1","dapp_auto_denounce_time":"60","dapp_bond_duration":"604800","dapp_inactive_rank_decrease_percent":"10","dapp_liquidation_period":"0","dapp_liquidation_threshold":"0","dapp_max_mischance":"10","dapp_mischance_rank_decrease_amount":"1","dapp_pool_slippage_default":"0.100000000000000000","dapp_verifier_bond":"0.001000000000000000","enable_foreign_fee_payments":true,"enable_token_blacklist":true,"enable_token_whitelist":false,"inactive_rank_decrease_percent":"0.500000000000000000","inflation_period":"31557600","inflation_rate":"0.180000000000000000","max_abstention":"2","max_annual_inflation":"0.350000000000000000","max_collective_outputs":"10","max_custody_buffer_size":"10","max_custody_tx_size":"8192","max_dapp_bond":"10000000","max_delegators":"100","max_jailed_percentage":"0.250000000000000000","max_mischance":"110","max_proposal_checksum_size":"128","max_proposal_description_size":"1024","max_proposal_poll_option_count":"128","max_proposal_poll_option_size":"64","max_proposal_reference_size":"512","max_proposal_title_size":"128","max_slashing_percentage":"0.010000000000000000","max_tx_fee":"1000000","min_collective_bond":"100000","min_collective_bonding_time":"86400","min_collective_claim_period":"14400","min_custody_reward":"200","min_dapp_bond":"1000000","min_delegation_pushout":"10","min_identity_approval_tip":"200","min_proposal_enactment_blocks":"1","min_proposal_end_blocks":"2","min_tx_fee":"100","min_validators":"1","minimum_proposal_end_time":"300","minting_ft_fee":"100000000000000","minting_nft_fee":"100000000000000","mischance_confidence":"10","mischance_rank_decrease_amount":"10","poor_network_max_bank_send":"1000000","proposal_enactment_time":"300","slashing_period":"3600","ubi_hardcap":"6000000","unique_identity_keys":"moniker,username","unjail_max_time":"600","unstaking_period":"2629800","validator_recovery_bond":"300000","validators_fee_share":"0.500000000000000000","vote_quorum":"33"},"next_role_id":"3","poor_network_messages":{"messages":["submit-proposal","set-network-properties","vote-proposal","claim-councilor","whitelist-permissions","blacklist-permissions","create-role","assign-role","remove-role","whitelist-role-permission","blacklist-role-permission","remove-whitelist-role-permission","remove-blacklist-role-permission","claim-validator","activate","pause","unpause","register-identity-records","edit-identity-record","request-identity-records-verify","handle-identity-records-verify-request","cancel-identity-records-verify-request"]},"proposal_durations":{},"proposals":[],"role_permissions":{"1":{"blacklist":[],"whitelist":[1,2,3,6,8,9,12,13,10,11,14,15,18,19,20,21,22,23,31,32,24,25,16,17,4,5,26,27,28,29,30,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66]},"2":{"blacklist":[],"whitelist":[2]}},"roles":[{"description":"Sudo role","id":1,"sid":"sudo"},{"description":"Validator role","id":2,"sid":"validator"}],"starting_proposal_id":"1","votes":[]},"customslashing":{"params":{"downtime_inactive_duration":"600s"},"signing_infos":[]},"customstaking":null,"distributor":{"fees_collected":[],"fees_treasury":[],"previous_proposer":"","snap_period":"1000","validator_votes":[],"year_start_snapshot":{"snapshot_amount":"0","snapshot_time":"0"}},"feeprocessing":{},"genutil":{"gen_txs":[]},"layer2":{"bridge":{"accounts":[],"helper":null,"tokens":[],"xams":[]},"dapps":[]},"multistaking":null,"params":null,"recovery":{"recovery_records":[],"recovery_tokens":[],"rewards":[],"rotations":[]},"spending":{"claims":[],"pools":[{"balances":[],"beneficiaries":{"accounts":[],"roles":[{"role":"2","weight":"1"}]},"claim_end":"0","claim_expiry":"0","claim_start":"0","dynamic_rate":false,"dynamic_rate_period":"0","last_dynamic_rate_calc_time":"0","name":"ValidatorBasicRewardsPool","owners":{"owner_accounts":[],"owner_roles":["2"]},"rates":[{"amount":"385.000000000000000000","denom":"ukex"}],"vote_enactment":"300","vote_period":"300","vote_quorum":"33"}]},"tokens":{"aliases":[{"decimals":6,"denoms":["ukex","mkex"],"icon":"","invalidated":false,"name":"Kira","symbol":"KEX"}],"rates":[{"denom":"ukex","fee_payments":true,"fee_rate":"1.000000000000000000","invalidated":false,"stake_cap":"0.500000000000000000","stake_min":"1","stake_token":true},{"denom":"ubtc","fee_payments":true,"fee_rate":"10.000000000000000000","invalidated":false,"stake_cap":"0.250000000000000000","stake_min":"1","stake_token":true},{"denom":"xeth","fee_payments":true,"fee_rate":"0.100000000000000000","invalidated":false,"stake_cap":"0.100000000000000000","stake_min":"1","stake_token":false},{"denom":"frozen","fee_payments":true,"fee_rate":"0.100000000000000000","invalidated":false,"stake_cap":"0.000000000000000000","stake_min":"1","stake_token":false}],"tokenBlackWhites":{"blacklisted":["frozen"],"whitelisted":["ukex"]}},"ubi":{"ubi_records":[{"amount":"500000","distribution_end":"0","distribution_last":"0","distribution_start":"0","dynamic":true,"name":"ValidatorBasicRewardsPoolUBI","period":"2592000","pool":"ValidatorBasicRewardsPool"}]},"upgrade":{"current_plan":null,"next_plan":null,"version":""}},"chain_id":"MyBlockchain-03","gentxs_dir":"","moniker":"MONIKER","node_id":"09f40bc7966028a2f58fe63130aadf02ff32bf29"}
```
</details>


[Return to top](#sekai)

### 12. keys

Keyring management commands

Usage:
```
sekaid keys [command]
```

Available Commands:

| Subcommand                  | Description                                                                                               |
| --------------------------- | --------------------------------------------------------------------------------------------------------- |
| [`add`](#121-add)           | Add an encrypted private key (either newly generated or recovered), encrypt it, and save to \<name\> file |
| [`delete`](#122-delete)     | Delete the given keys                                                                                     |
| [`export`](#123-export)     | Export private keys                                                                                       |
| [`import`](#124-import)     | Import private keys into the local keybase                                                                |
| [`list`](#125-list)         | List all keys                                                                                             |
| [`migrate`](#126-migrate)   | Migrate keys from the legacy (db-based) Keybase                                                           |
| [`mnemonic`](#127-mnemonic) | Compute the bip39 mnemonic for some input entropy                                                         |
| [`parse`](#128-parse)       | Parse address from hex to bech32 and vice versa                                                           |
| [`show`](#129-show)         | Retrieve key information by name or address                                                               |



| Flags                      | Description                                                                           | Work  |
| -------------------------- | ------------------------------------------------------------------------------------- | ----- |
| `-h, --help`               | Help for keys                                                                         | ✅ yes |
| `--keyring-backend string` | Select keyring's backend (`os\|file\|test`) (default `"os"`)                          | ✅ yes |
| `--keyring-dir string`     | The client Keyring directory; if omitted, the default `'home'` directory will be used | ✅ yes |
| `--output string`          | Output format (`text\|json`) (default `"text"`)                                       | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--home string`       | Directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid keys --help
Keyring management commands. These keys may be in any format supported by the
Tendermint crypto library and can be used by light-clients, full nodes, or any other application
that needs to sign with a private key.

The keyring supports the following backends:

    os          Uses the operating system's default credentials store.
    file        Uses encrypted file-based keystore within the app's configuration directory.
                This keyring will request a password each time it is accessed, which may occur
                multiple times in a single command resulting in repeated password prompts.
    kwallet     Uses KDE Wallet Manager as a credentials management application.
    pass        Uses the pass command line utility to store and retrieve keys.
    test        Stores keys insecurely to disk. It does not prompt for a password to be unlocked
                and it should be use only for testing purposes.

kwallet and pass backends depend on external tools. Refer to their respective documentation for more
information:
    KWallet     https://github.com/KDE/kwallet
    pass        https://www.passwordstore.org/

The pass backend requires GnuPG: https://gnupg.org/

Usage:
  sekaid keys [command]

Available Commands:
  add         Add an encrypted private key (either newly generated or recovered), encrypt it, and save to <name> file
  delete      Delete the given keys
  export      Export private keys
  import      Import private keys into the local keybase
  list        List all keys
  migrate     Migrate keys from the legacy (db-based) Keybase
  mnemonic    Compute the bip39 mnemonic for some input entropy
  parse       Parse address from hex to bech32 and vice versa
  show        Retrieve key information by name or address

Flags:
  -h, --help                     help for keys
      --keyring-backend string   Select keyring's backend (os|file|test) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --output string            Output format (text|json) (default "text")

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors

Use "sekaid keys [command] --help" for more information about a command.
```

[Return to top](#sekai)

#### 12.1 add

Derive a new private key and encrypt to disk.

Usage:
```
sekaid keys add <name> [flags]
```

| Flags                      | Description                                                                          | Work  |
| -------------------------- | ------------------------------------------------------------------------------------ | ----- |
| `--account uint32`         | Account number for HD derivation                                                     | ❌ ?   |
| `--algo string`            | Key signing algorithm to generate keys for (default `"secp256k1"`)                   | ❌ ?   |
| `--coin-type uint32`       | Coin type number for HD derivation (default `118`)                                   | ❌ ?   |
| `--dry-run`                | Perform action, but don't add key to local keystore                                  | ❌ ?   |
| `--hd-path string`         | Manual HD Path derivation (overrides BIP44 config)                                   | ❌ ?   |
| `-h, --help`               | Help for add                                                                         | ✅ yes |
| `--index uint32`           | Address index number for HD derivation                                               | ❌ ?   |
| `-i, --interactive`        | Interactively prompt user for BIP39 passphrase and mnemonic                          | ❌ ?   |
| `--ledger`                 | Store a local reference to a private key on a Ledger device                          | ❌ ?   |
| `--multisig strings`       | List of key names stored in keyring to construct a public legacy multisig key        | ❌ ?   |
| `--multisig-threshold int` | K out of N required signatures. For use in conjunction with --multisig (default `1`) | ❌ ?   |
| `--no-backup`              | Don't print out seed phrase (if others are watching the terminal)                    | ❌ ?   |
| `--nosort`                 | Keys passed to --multisig are taken in the order they're supplied                    | ❌ ?   |
| `--pubkey string`          | Parse a public key in JSON format and saves key info to \<name\> file.               | ❌ ?   |
| `--recover`                | Provide seed phrase to recover existing key instead of creating                      | ✅ yes |



| Global Flags               | Description                                                                            | Work  |
| -------------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--home string`            | The application home directory (default `"/root/.sekaid"`)                             | ✅ yes |
| `--keyring-backend string` | Select keyring's backend (`os\|file\|test`) (default `"os"`)                           | ✅ yes |
| `--keyring-dir string`     | The client Keyring directory; if omitted, the default `'home'` directory will be used  | ✅ yes |
| `--log_format string`      | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`       | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--output string`          | Output format (`text\|json`) (default `"text"`)                                        | ✅ yes |
| `--trace`                  | Print out full stack trace on errors                                                   | ❌ ?   |

```
/# keys add --help
Derive a new private key and encrypt to disk.
Optionally specify a BIP39 mnemonic, a BIP39 passphrase to further secure the mnemonic,
and a bip32 HD path to derive a specific account. The key will be stored under the given name
and encrypted with the given password. The only input that is required is the encryption password.

If run with -i, it will prompt the user for BIP44 path, BIP39 mnemonic, and passphrase.
The flag --recover allows one to recover a key from a seed passphrase.
If run with --dry-run, a key would be generated (or recovered) but not stored to the
local keystore.
Use the --pubkey flag to add arbitrary public keys to the keystore for constructing
multisig transactions.

You can create and store a multisig key by passing the list of key names stored in a keyring
and the minimum number of signatures required through --multisig-threshold. The keys are
sorted by address, unless the flag --nosort is set.
Example:

    keys add mymultisig --multisig "keyname1,keyname2,keyname3" --multisig-threshold 2

Usage:
  sekaid keys add <name> [flags]

Flags:
      --account uint32           Account number for HD derivation
      --algo string              Key signing algorithm to generate keys for (default "secp256k1")
      --coin-type uint32         coin type number for HD derivation (default 118)
      --dry-run                  Perform action, but don't add key to local keystore
      --hd-path string           Manual HD Path derivation (overrides BIP44 config)
  -h, --help                     help for add
      --index uint32             Address index number for HD derivation
  -i, --interactive              Interactively prompt user for BIP39 passphrase and mnemonic
      --ledger                   Store a local reference to a private key on a Ledger device
      --multisig strings         List of key names stored in keyring to construct a public legacy multisig key
      --multisig-threshold int   K out of N required signatures. For use in conjunction with --multisig (default 1)
      --no-backup                Don't print out seed phrase (if others are watching the terminal)
      --nosort                   Keys passed to --multisig are taken in the order they're supplied
      --pubkey string            Parse a public key in JSON format and saves key info to <name> file.
      --recover                  Provide seed phrase to recover existing key instead of creating

Global Flags:
      --home string              The application home directory (default "/root/.sekaid")
      --keyring-backend string   Select keyring's backend (os|file|test) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --log_format string        The logging format (json|plain) (default "plain")
      --log_level string         The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --output string            Output format (text|json) (default "text")
      --trace                    print out full stack trace on errors
```

```
/# sekaid keys add testForDocs --home=/root/.sekai --keyring-backend=test --output=json | jq
{
  "name": "testForDocs",
  "type": "local",
  "address": "kira1xshj7342ar7hk08d6p5ynrwxzux29jzr9pdx40",
  "pubkey": "{\"@type\":\"/cosmos.crypto.secp256k1.PubKey\",\"key\":\"AhPJzI5jcMD1xJ6at1ItLT6d1esQhTN7m8bqnXRVfREZ\"}",
  "mnemonic": "fly dog voice claw pattern found open room extend victory wrap butter vast urban exclude staff private matrix way endorse bone eyebrow already genuine"
}
```

Recover by `--recover`:
```
/# echo "fly dog voice claw pattern found open room extend victory wrap butter vast urban exclude staff private matrix way endorse bone eyebrow already genuine" | sekaid keys add testForDocs --keyring-backend=test --home=/root/.sekai --recover --output=json | jq
{
  "name": "testForDocs",
  "type": "local",
  "address": "kira1xshj7342ar7hk08d6p5ynrwxzux29jzr9pdx40",
  "pubkey": "{\"@type\":\"/cosmos.crypto.secp256k1.PubKey\",\"key\":\"AhPJzI5jcMD1xJ6at1ItLT6d1esQhTN7m8bqnXRVfREZ\"}"
}
```
_If mnemonic will not be provided by pipe, it will be prompted manually_
```
/# sekaid keys add testForDocs2 --keyring-backend=test --home=/root/.sekai --recover --output=json | jq
> Enter your bip39 mnemonic
fly dog voice claw pattern found open room extend victory wrap butter vast urban exclude staff private matrix way endorse bone eyebrow already genuine
{
  "name": "testForDocs2",
  "type": "local",
  "address": "kira1xshj7342ar7hk08d6p5ynrwxzux29jzr9pdx40",
  "pubkey": "{\"@type\":\"/cosmos.crypto.secp256k1.PubKey\",\"key\":\"AhPJzI5jcMD1xJ6at1ItLT6d1esQhTN7m8bqnXRVfREZ\"}"
}
```

Interactive with `-i | --interactive`:
```
/# sekaid keys add testForDocs4 -i --keyring-backend=test --output=json | jq
> Enter your bip39 mnemonic, or hit enter to generate one.

> Enter your bip39 passphrase. This is combined with the mnemonic to derive the seed. Most users should just hit enter to use the default, ""

{
  "name": "testForDocs4",
  "type": "local",
  "address": "kira1h4uu0skupqutvfzw9zyudaj5vczut28pmc6rx0",
  "pubkey": "{\"@type\":\"/cosmos.crypto.secp256k1.PubKey\",\"key\":\"A3xP2BL5RjIT0BZqjEf6j/OqffTX0pgEpsXxdoVIikEI\"}",
  "mnemonic": "spatial forest make border anger armor century bird lava month jeans inch paper alien nation thing interest joy machine cable index wreck spin property"
}
```

#TODO add other usages

🟨  
🟨  
🟨  

[Return to top](#sekai)

#### 12.2 delete

Delete keys from the Keybase backend.

Usage:
```
sekaid keys delete <name> ... [flags]
```

| Flags         | Description                                                                   | Work  |
| ------------- | ----------------------------------------------------------------------------- | ----- |
| `-f, --force` | Remove the key unconditionally without asking for the passphrase. Deprecated. | ❌ no  |
| `-h, --help`  | Help for delete                                                               | ✅ yes |
| `-y, --yes`   | Skip confirmation prompt when deleting offline or ledger key references       | ✅ yes |



| Global Flags               | Description                                                                            | Work  |
| -------------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--home string`            | The application home directory (default `"/root/.sekaid"`)                             | ✅ yes |
| `--keyring-backend string` | Select keyring's backend (`os\|file\|test`) (default `"os"`)                           | ✅ yes |
| `--keyring-dir string`     | The client Keyring directory; if omitted, the default `'home'` directory will be used  | ✅ yes |
| `--log_format string`      | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`       | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--output string`          | Output format (`text\|json`) (default `"text"`)                                        | ❌ no  |
| `--trace`                  | Print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid keys delete --help
Delete keys from the Keybase backend.

Note that removing offline or ledger keys will remove
only the public key references stored locally, i.e.
private keys stored in a ledger device cannot be deleted with the CLI.

Usage:
  sekaid keys delete <name>... [flags]

Flags:
  -f, --force   Remove the key unconditionally without asking for the passphrase. Deprecated.
  -h, --help    help for delete
  -y, --yes     Skip confirmation prompt when deleting offline or ledger key references

Global Flags:
      --home string              The application home directory (default "/root/.sekaid")
      --keyring-backend string   Select keyring's backend (os|file|test) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --log_format string        The logging format (json|plain) (default "plain")
      --log_level string         The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --output string            Output format (text|json) (default "text")
      --trace                    print out full stack trace on errors
```

```
/# sekaid keys delete testForDocs --home=/root/.sekai --keyring-backend=test --yes
Key deleted forever (uh oh!)
```

❌: If key does not exist:
```
/# sekaid keys delete testForDocs --home=/root/.sekai --keyring-backend=test --yes --output=json | jq
Error: testForDocs.info: key not found
```

[Return to top](#sekai)

#### 12.3 export

Export a private key from the local keyring in ASCII-armored encrypted format.

Usage:
```
sekaid keys export <name> [flags]
```

| Flags             | Description                                                                                               | Work  |
| ----------------- | --------------------------------------------------------------------------------------------------------- | ----- |
| `-h, --help`      | Help for export                                                                                           | ✅ yes |
| `--unarmored-hex` | Export unarmored hex privkey. Requires `--unsafe`.                                                        | ✅ yes |
| `--unsafe`        | Enable unsafe operations. This flag must be switched on along with all unsafe operation-specific options. | ✅ yes |



| Global Flags               | Description                                                                            | Work  |
| -------------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--home string`            | The application home directory (default `"/root/.sekaid"`)                             | ✅ yes |
| `--keyring-backend string` | Select keyring's backend (`os\|file\|test`) (default `"os"`)                           | ✅ yes |
| `--keyring-dir string`     | The client Keyring directory; if omitted, the default `'home'` directory will be used  | ✅ yes |
| `--log_format string`      | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`       | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--output string`          | Output format (`text\|json`) (default `"text"`)                                        | ❌ no  |
| `--trace`                  | Print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid keys export --help
Export a private key from the local keyring in ASCII-armored encrypted format.

When both the --unarmored-hex and --unsafe flags are selected, cryptographic
private key material is exported in an INSECURE fashion that is designed to
allow users to import their keys in hot wallets. This feature is for advanced
users only that are confident about how to handle private keys work and are
FULLY AWARE OF THE RISKS. If you are unsure, you may want to do some research
and export your keys in ASCII-armored encrypted format.

Usage:
  sekaid keys export <name> [flags]

Flags:
  -h, --help            help for export
      --unarmored-hex   Export unarmored hex privkey. Requires --unsafe.
      --unsafe          Enable unsafe operations. This flag must be switched on along with all unsafe operation-specific options.

Global Flags:
      --home string              The application home directory (default "/root/.sekaid")
      --keyring-backend string   Select keyring's backend (os|file|test) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --log_format string        The logging format (json|plain) (default "plain")
      --log_level string         The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --output string            Output format (text|json) (default "text")
      --trace                    print out full stack trace on errors
```

```
/# sekaid keys export testForDocs --home=/root/.sekai --keyring-backend=test --output=json
Enter passphrase to encrypt the exported key: <input password>
-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: 03D68CF41D5698BD951C8897917AAA23
type: secp256k1

J4O49A3wRJkioJOAOgHpZq0h0iABo8BoRc79Iy9fKaItridiH/obaiJawUFH0he0
IJ7Yjha4ZVQDoChl44ot9NgN/ulGXKViMPCU2ps=
=RQle
-----END TENDERMINT PRIVATE KEY-----
```

```
/# sekaid keys export testForDocs --home=/root/.sekai --keyring-backend=test --output=json --unarmored-hex --unsafe
WARNING: The private key will be exported as an unarmored hexadecimal string. USE AT YOUR OWN RISK. Continue? [y/N]: y
96176c3b3e94205b662d8cd4e68e536254d813067b4b6e475ac5ebd06a2ef8b9
```
or
```
yes | sekaid keys export testForDocs2 --home=/root/.sekai --keyring-backend=test --output=json --unarmored-hex --unsafe
96176c3b3e94205b662d8cd4e68e536254d813067b4b6e475ac5ebd06a2ef8b9
```

[Return to top](#sekai)

#### 12.4 import

Import a ASCII armored private key into the local keybase.

Usage:
```
sekaid keys import <name> <keyfile> [flags]
```

| Flags        | Description     | Work  |
| ------------ | --------------- | ----- |
| `-h, --help` | Help for import | ✅ yes |



| Global Flags               | Description                                                                            | Work  |
| -------------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--home string`            | The application home directory (default `"/root/.sekaid"`)                             | ✅ yes |
| `--keyring-backend string` | Select keyring's backend (`os\|file\|test`) (default `"os"`)                           | ✅ yes |
| `--keyring-dir string`     | The client Keyring directory; if omitted, the default `'home'` directory will be used  | ✅ yes |
| `--log_format string`      | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`       | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--output string`          | Output format (`text\|json`) (default `"text"`)                                        | ❌ ?   |
| `--trace`                  | Print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid keys import --help
Import a ASCII armored private key into the local keybase.

Usage:
  sekaid keys import <name> <keyfile> [flags]

Flags:
  -h, --help   help for import

Global Flags:
      --home string              The application home directory (default "/root/.sekaid")
      --keyring-backend string   Select keyring's backend (os|file|test) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --log_format string        The logging format (json|plain) (default "plain")
      --log_level string         The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --output string            Output format (text|json) (default "text")
      --trace                    print out full stack trace on errors
```

#TODO

🟨  
🟨  
🟨  

```
/# sekaid keys import <name> <file> --home=/root/.sekai --keyring-backend=test
```

[Return to top](#sekai)

#### 12.5 list

Return a list of all public keys stored by this key manager along with their associated name and address.

Usage:
```
sekaid keys list [flags]
```

| Flags              | Description     | Work  |
| ------------------ | --------------- | ----- |
| `-h, --help`       | Help for list   | ✅ yes |
| `-n, --list-names` | List names only | ✅ yes |



| Global Flags               | Description                                                                            | Work  |
| -------------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--home string`            | The application home directory (default `"/root/.sekaid"`)                             | ✅ yes |
| `--keyring-backend string` | Select keyring's backend (`os\|file\|test`) (default `"os"`)                           | ✅ yes |
| `--keyring-dir string`     | The client Keyring directory; if omitted, the default `'home'` directory will be used  | ✅ yes |
| `--log_format string`      | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`       | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--output string`          | Output format (`text\|json`) (default `"text"`)                                        | ✅ yes |
| `--trace`                  | Print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid keys list --help
Return a list of all public keys stored by this key manager
along with their associated name and address.

Usage:
  sekaid keys list [flags]

Flags:
  -h, --help         help for list
  -n, --list-names   List names only

Global Flags:
      --home string              The application home directory (default "/root/.sekaid")
      --keyring-backend string   Select keyring's backend (os|file|test) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --log_format string        The logging format (json|plain) (default "plain")
      --log_level string         The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --output string            Output format (text|json) (default "text")
      --trace                    print out full stack trace on errors
```

```
/# sekaid keys list --home=/root/.sekai --keyring-backend=test
  type: local
  address: kira1qwqtytwh0p4u6pmh9fzh9t8r3kyyzsaarpmqd0
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A12kfllXSYJJr0iFELFCC7gASzTdPx6QVAFIPHgyRoB9"}'
  mnemonic: ""
- name: test
  type: local
  address: kira1j5a333q9vuazmdtc0r9henhyyw97q8t620x4wn
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AqVtJBdzF5xnHtpDA50RKHz8g9MLUW/KDqynJ5QfqJPK"}'
  mnemonic: ""
. . .
```

```
/# sekaid keys list --home=/root/.sekai --keyring-backend=test --output=json | jq
[
  {
    "name": "00400705-b8ea-41c1-91bb-5cb26e42b7cb",
    "type": "local",
    "address": "kira184lvz68flxvxwpkvavvz7p3dzenvpcten4xaaz",
    "pubkey": "{\"@type\":\"/cosmos.crypto.secp256k1.PubKey\",\"key\":\"A5eVNtogRjILK4TCsXhDicb+FGH9A8JoZuQ6uLQWG/O5\"}"
  },
  . . .
  {
    "name": "validator",
    "type": "local",
    "address": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
    "pubkey": "{\"@type\":\"/cosmos.crypto.secp256k1.PubKey\",\"key\":\"AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo\"}"
  }
]
```

List of names only:
```
/# sekaid keys list --home=/root/.sekai --keyring-backend=test --list-names
signer
test
...
validator
```

[Return to top](#sekai)

#### 12.6 migrate

Migrate key information from the legacy (db-based) Keybase to the new keyring-based Keyring.

Usage:
```
sekaid keys migrate <old_home_dir> [flags]
```

| Flags        | Description                                                              | Work  |
| ------------ | ------------------------------------------------------------------------ | ----- |
| `--dry-run`  | Run migration without actually persisting any changes to the new Keybase | ❌ ?   |
| `-h, --help` | Help for migrate                                                         | ✅ yes |



| Global Flags               | Description                                                                            | Work |
| -------------------------- | -------------------------------------------------------------------------------------- | ---- |
| `--home string`            | The application home directory (default `"/root/.sekaid"`)                             | ❌ ?  |
| `--keyring-backend string` | Select keyring's backend (`os\|file\|test`) (default `"os"`)                           | ❌ ?  |
| `--keyring-dir string`     | The client Keyring directory; if omitted, the default `'home'` directory will be used  | ❌ ?  |
| `--log_format string`      | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?  |
| `--log_level string`       | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?  |
| `--output string`          | Output format (`text\|json`) (default `"text"`)                                        | ❌ ?  |
| `--trace`                  | Print out full stack trace on errors                                                   | ❌ ?  |

#TODO

🟨  
🟨  
🟨  

[Return to top](#sekai)

#### 12.7 mnemonic

Create a bip39 mnemonic

Usage:
```
sekaid keys mnemonic [flags]
```

| Flags              | Description                                                                   | Work  |
| ------------------ | ----------------------------------------------------------------------------- | ----- |
| `-h, --help`       | Help for mnemonic                                                             | ✅ yes |
| `--unsafe-entropy` | Prompt the user to supply their own entropy, instead of relying on the system | ✅ yes |



| Global Flags               | Description                                                                            | Work  |
| -------------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--home string`            | The application home directory (default `"/root/.sekaid"`)                             | ✅ yes |
| `--keyring-backend string` | Select keyring's backend (`os\|file\|test`) (default `"os"`)                           | ✅ yes |
| `--keyring-dir string`     | The client Keyring directory; if omitted, the default `'home'` directory will be used  | ✅ yes |
| `--log_format string`      | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`       | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--output string`          | Output format (`text\|json`) (default `"text"`)                                        | ❌ no  |
| `--trace`                  | Print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid keys mnemonic --help
Create a bip39 mnemonic, sometimes called a seed phrase, by reading from the system entropy. To pass your own entropy, use --unsafe-entropy

Usage:
  sekaid keys mnemonic [flags]

Flags:
  -h, --help             help for mnemonic
      --unsafe-entropy   Prompt the user to supply their own entropy, instead of relying on the system

Global Flags:
      --home string              The application home directory (default "/root/.sekaid")
      --keyring-backend string   Select keyring's backend (os|file|test) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --log_format string        The logging format (json|plain) (default "plain")
      --log_level string         The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --output string            Output format (text|json) (default "text")
      --trace                    print out full stack trace on errors
```

```
/# sekaid keys mnemonic
exit front isolate strike problem truly solve village rack crash radar open spare portion dad box fabric deny custom blame analyst pond merit vicious
```

```
/# sekaid keys mnemonic --unsafe-entropy
> > WARNING: Generate at least 256-bits of entropy and enter the results here:
3Q6yhoMe+zQCtX3Qh/zYr23PiuTZJHP/L7DNojL1Jzk=  // Example command for generating: `dd if=/dev/urandom bs=32 count=1 status=none | base64`
> Input length: 44 [y/N]: y
bicycle choose frame quality symbol super door eye fence glare slow merit danger naive dune violin rare december return measure axis young drastic suggest
```

[Return to top](#sekai)

#### 12.8 parse

Convert and print to stdout key addresses and fingerprints from hexadecimal into bech32 cosmos prefixed format and vice versa.

Usage:
```
sekaid keys parse <hex-or-bech32-address> [flags]
```

| Flags        | Description    | Work  |
| ------------ | -------------- | ----- |
| `-h, --help` | help for parse | ✅ yes |



| Global Flags               | Description                                                                            | Work  |
| -------------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--home string`            | The application home directory (default `"/root/.sekaid"`)                             | ✅ yes |
| `--keyring-backend string` | Select keyring's backend (`os\|file\|test`) (default `"os"`)                           | ✅ yes |
| `--keyring-dir string`     | The client Keyring directory; if omitted, the default `'home'` directory will be used  | ✅ yes |
| `--log_format string`      | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`       | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--output string`          | Output format (`text\|json`) (default `"text"`)                                        | ✅ yes |
| `--trace`                  | Print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid keys parse --help
Convert and print to stdout key addresses and fingerprints from
hexadecimal into bech32 cosmos prefixed format and vice versa.

Usage:
  sekaid keys parse <hex-or-bech32-address> [flags]

Flags:
  -h, --help   help for parse

Global Flags:
      --home string              The application home directory (default "/root/.sekaid")
      --keyring-backend string   Select keyring's backend (os|file|test) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --log_format string        The logging format (json|plain) (default "plain")
      --log_level string         The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --output string            Output format (text|json) (default "text")
      --trace                    print out full stack trace on errors
```

Here we can provide addresses with prefixes: 
- `kira`
- `kirapub`
- `kiravaloper`
- `kiravaloperpub`
- `kiravalcons`
- `kiravalconspub`
```
/# sekaid keys parse kira1j5a333q9vuazmdtc0r9henhyyw97q8t620x4wn
human: kira
bytes: 953B18C405673A2DB57878CB7CCEE4238BE01D7A
```

```
/# sekaid keys parse kira1j5a333q9vuazmdtc0r9henhyyw97q8t620x4wn --output=json | jq
{
  "human": "kira",
  "bytes": "953B18C405673A2DB57878CB7CCEE4238BE01D7A"
}
```

Or we can provide hex address (without prefix)
```
/# sekaid keys parse 953B18C405673A2DB57878CB7CCEE4238BE01D7A --output=json | jq
{
  "formats": [
    "kira1j5a333q9vuazmdtc0r9henhyyw97q8t620x4wn",
    "kirapub1j5a333q9vuazmdtc0r9henhyyw97q8t68c055y",
    "kiravaloper1j5a333q9vuazmdtc0r9henhyyw97q8t6ef6kkl",
    "kiravaloperpub1j5a333q9vuazmdtc0r9henhyyw97q8t657dssf",
    "kiravalcons1j5a333q9vuazmdtc0r9henhyyw97q8t6d6f267",
    "kiravalconspub1j5a333q9vuazmdtc0r9henhyyw97q8t696kc6c"
  ]
}
```

[Return to top](#sekai)

#### 12.9 show

Display keys details

Usage:
```
sekaid keys show [name_or_address [name_or_address...]] [flags]
```

| Flags                      | Description                                                               | Work  |
| -------------------------- | ------------------------------------------------------------------------- | ----- |
| `-a, --address`            | Output the address only (overrides `--output`)                            | ✅ yes |
| `--bech string`            | The Bech32 prefix encoding for a key (`acc\|val\|cons`) (default `"acc"`) | ✅ yes |
| `-d, --device`             | Output the address in a ledger device                                     | ❌ ?   |
| `-h, --help`               | Help for show                                                             | ✅ yes |
| `--multisig-threshold int` | K out of N required signatures (default `1`)                              | ❌ ?   |
| `-p, --pubkey`             | Output the public key only (overrides `--output`)                         | ✅ yes |



| Global Flags               | Description                                                                            | Work  |
| -------------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--home string`            | The application home directory (default `"/root/.sekaid"`)                             | ✅ yes |
| `--keyring-backend string` | Select keyring's backend (`os\|file\|test`) (default `"os"`)                           | ✅ yes |
| `--keyring-dir string`     | The client Keyring directory; if omitted, the default `'home'` directory will be used  | ✅ yes |
| `--log_format string`      | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`       | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--output string`          | Output format (`text\|json`) (default `"text"`)                                        | ✅ yes |
| `--trace`                  | Print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid keys show --help
Display keys details. If multiple names or addresses are provided,
then an ephemeral multisig key will be created under the name "multi"
consisting of all the keys provided by name and multisig threshold.

Usage:
  sekaid keys show [name_or_address [name_or_address...]] [flags]

Flags:
  -a, --address                  Output the address only (overrides --output)
      --bech string              The Bech32 prefix encoding for a key (acc|val|cons) (default "acc")
  -d, --device                   Output the address in a ledger device
  -h, --help                     help for show
      --multisig-threshold int   K out of N required signatures (default 1)
  -p, --pubkey                   Output the public key only (overrides --output)

Global Flags:
      --home string              The application home directory (default "/root/.sekaid")
      --keyring-backend string   Select keyring's backend (os|file|test) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --log_format string        The logging format (json|plain) (default "plain")
      --log_level string         The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --output string            Output format (text|json) (default "text")
      --trace                    print out full stack trace on errors
```

```
/# sekaid keys show validator --home=/root/.sekai --keyring-backend=test --output=json | jq
{
  "name": "validator",
  "type": "local",
  "address": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
  "pubkey": "{\"@type\":\"/cosmos.crypto.secp256k1.PubKey\",\"key\":\"AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo\"}"
}
```

Only address:
```
/# sekaid keys show validator --home=/root/.sekai --keyring-backend=test -a
kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4
```

Only pubkey:
```
/# sekaid keys show validator --home=/root/.sekai --keyring-backend=test -p | jq
{
  "@type": "/cosmos.crypto.secp256k1.PubKey",
  "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
}
```

The Bech32 addresses (based on chosen prefix by `--bech`):
```
/# sekaid keys show validator --home=/root/.sekai --keyring-backend=test --bech acc -a
kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4
```

```
/# sekaid keys show validator --home=/root/.sekai --keyring-backend=test --bech val -a
kiravaloper1vmwdgw426aj9fx33fqusmtg6r65yyucm4ulwne
```

```
/# sekaid keys show validator --home=/root/.sekai --keyring-backend=test --bech cons -a
kiravalcons1vmwdgw426aj9fx33fqusmtg6r65yyucmp0vjlc
```

[Return to top](#sekai)

### 13. new-genesis-from-exported

[Return to top](#sekai)

### 14. query

Querying subcommands

Usage:
```
sekaid query [flags]
sekaid query [command]
```

Available Commands:

| Subcommand                                                   | Description                                                                                                      |
| ------------------------------------------------------------ | ---------------------------------------------------------------------------------------------------------------- |
| [`account`](#141-account)                                    | Query for account by address                                                                                     |
| [`auth`](#142-auth)                                          | Querying commands for the auth module                                                                            |
| [`bank`](#143-bank)                                          | Querying commands for the bank module                                                                            |
| [`basket`](#144-basket)                                      | query commands for the basket module                                                                             |
| [`block`](#145-block)                                        | Get verified data for a the block at given height                                                                |
| [`collectives`](#146-collectives)                            | query commands for the collectives module                                                                        |
| [`custody`](#147-custody)                                    | query commands for the custody module                                                                            |
| [`customevidence`](#148-customevidence)                      | Query for evidence by hash or for all (paginated) submitted evidence                                             |
| [`customgov`](#149-customgov)                                | query commands for the customgov module                                                                          |
| [`customslashing`](#1410-customslashing)                     | Querying commands for the slashing module                                                                        |
| [`customstaking`](#1411-customstaking)                       | Querying commands for the staking module                                                                         |
| [`distributor`](#1412-distributor)                           | query commands for the distributor module                                                                        |
| [`layer2`](#1413-layer2)                                     | query commands for the layer2 module                                                                             |
| [`multistaking`](#1414-multistaking)                         | Querying commands for the multistaking module                                                                    |
| [`params`](#1415-params)                                     | Querying commands for the params module                                                                          |
| [`recovery`](#1416-recovery)                                 | Querying commands for the recovery module                                                                        |
| [`spending`](#1417-spending)                                 | query commands for the tokens module                                                                             |
| [`tendermint-validator-set`](#1418-tendermint-validator-set) | Get the full tendermint validator set at given height                                                            |
| [`tokens`](#1419-tokens)                                     | query commands for the tokens module                                                                             |
| [`tx`](#1420-tx)                                             | Query for a transaction by hash, `"<addr>/<seq>"` combination or comma-separated signatures in a committed block |
| [`txs`](#1421-txs)                                           | Query for paginated transactions that match a set of events                                                      |
| [`ubi`](#1422-ubi)                                           | query commands for the ubi module                                                                                |
| [`upgrade`](#1423-upgrade)                                   | Querying commands for the upgrade module                                                                         |

```
/# sekaid q --help
Querying subcommands

Usage:
  sekaid query [flags]
  sekaid query [command]

Aliases:
  query, q

Available Commands:
  account                  Query for account by address
  auth                     Querying commands for the auth module
  bank                     Querying commands for the bank module
  basket                   query commands for the basket module
  block                    Get verified data for a the block at given height
  collectives              query commands for the collectives module
  custody                  query commands for the custody module
  customevidence           Query for evidence by hash or for all (paginated) submitted evidence
  customgov                query commands for the customgov module
  customslashing           Querying commands for the slashing module
  customstaking            Querying commands for the staking module
  distributor              query commands for the distributor module
  layer2                   query commands for the layer2 module
  multistaking             Querying commands for the multistaking module
  params                   Querying commands for the params module
  recovery                 Querying commands for the recovery module
  spending                 query commands for the tokens module
  tendermint-validator-set Get the full tendermint validator set at given height
  tokens                   query commands for the tokens module
  tx                       Query for a transaction by hash, "<addr>/<seq>" combination or comma-separated signatures in a committed block
  txs                      Query for paginated transactions that match a set of events
  ubi                      query commands for the ubi module
  upgrade                  Querying commands for the upgrade module

Flags:
      --chain-id string   The network chain ID
  -h, --help              help for query

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors

Use "sekaid query [command] --help" for more information about a command.
```

[Return to top](#sekai)

#### 14.1 account

Query for account by address

Usage:
```
sekaid query account [address] [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for account                                                                                 | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--home string`       | Directory for config and data (default `"/root/.sekaid"`)                              | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid query account --help
Query for account by address

Usage:
  sekaid query account [address] [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for account
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid query account kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --output=json | jq
{
  "@type": "/cosmos.auth.v1beta1.BaseAccount",
  "address": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
  "pub_key": {
    "@type": "/cosmos.crypto.secp256k1.PubKey",
    "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
  },
  "account_number": "0",
  "sequence": "85"
}
```

[Return to top](#sekai)

#### 14.2 auth

Querying commands for the auth module

Usage:
```
  sekaid query auth [flags]
  sekaid query auth [command]
```

Available Commands:

| Subcommand                               | Description                              |
| ---------------------------------------- | ---------------------------------------- |
| [`account`](#1421-account)               | Query for account by address             |
| [`accounts`](#1422-accounts)             | Query all the accounts                   |
| [`module-account`](#1423-module-account) | Query module account info by module name |
| [`params`](#1424-params)                 | Query the current auth parameters        |



| Flags        | Description   | Work |
| ------------ | ------------- | ---- |
| `-h, --help` | help for auth |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid query auth account --help
Query for account by address

Usage:
  sekaid query auth account [address] [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for account
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

[Return to top](#sekai)

##### 14.2.1 account

Query for account by address.

Usage:
```
sekaid query auth account [address] [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for account                                                                                 | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--home string`       | Directory for config and data (default `"/root/.sekaid"`)                              | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid query auth account --help
Query for account by address

Usage:
  sekaid query auth account [address] [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for account
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
/# sekaid query auth account kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --output=json | jq
{
  "@type": "/cosmos.auth.v1beta1.BaseAccount",
  "address": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
  "pub_key": {
    "@type": "/cosmos.crypto.secp256k1.PubKey",
    "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
  },
  "account_number": "0",
  "sequence": "85"
}
```

On specific `--height` condition:
```
/# sekaid query auth account kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --height 40980 --output=json | jq
{
  "@type": "/cosmos.auth.v1beta1.BaseAccount",
  "address": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
  "pub_key": {
    "@type": "/cosmos.crypto.secp256k1.PubKey",
    "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
  },
  "account_number": "0",
  "sequence": "70"
}
```

[Return to top](#sekai)

##### 14.2.2 accounts

Query all the accounts

Usage:
```
sekaid query auth accounts [flags]
```

| Flags                 | Description                                                                                         | Work  |
| --------------------- | --------------------------------------------------------------------------------------------------- | ----- |
| `--count-total`       | count total number of records in all-accounts to query for                                          | ✅ yes |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)               | ✅ yes |
| `-h, --help`          | help for accounts                                                                                   | ✅ yes |
| `--limit uint`        | pagination limit of all-accounts to query for (default `100`)                                       | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`)    | ✅ yes |
| `--offset uint`       | pagination offset of all-accounts to query for                                                      | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                     | ✅ yes |
| `--page uint`         | pagination page of all-accounts to query for. This sets offset to a multiple of limit (default `1`) | ✅ yes |
| `--page-key string`   | pagination page-key of all-accounts to query for                                                    | ✅ yes |
| `--reverse`           | results are sorted in descending order                                                              | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--home string`       | Directory for config and data (default `"/root/.sekaid"`)                              | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid query auth accounts --help
Query all the accounts

Usage:
  sekaid query auth accounts [flags]

Flags:
      --count-total       count total number of records in all-accounts to query for
      --height int        Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help              help for accounts
      --limit uint        pagination limit of all-accounts to query for (default 100)
      --node string       <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
      --offset uint       pagination offset of all-accounts to query for
  -o, --output string     Output format (text|json) (default "text")
      --page uint         pagination page of all-accounts to query for. This sets offset to a multiple of limit (default 1)
      --page-key string   pagination page-key of all-accounts to query for
      --reverse           results are sorted in descending order

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid query auth accounts --limit=2 --reverse --output=json | jq
{
  "accounts": [
    {
      "@type": "/cosmos.auth.v1beta1.BaseAccount",
      "address": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
      "pub_key": {
        "@type": "/cosmos.crypto.secp256k1.PubKey",
        "key": "A5mB81789jXij6eUh5QGrRlhXdLheHFL1ix1LtxfMCvJ"
      },
      "account_number": "6",
      "sequence": "5"
    },
    {
      "@type": "/cosmos.auth.v1beta1.ModuleAccount",
      "base_account": {
        "address": "kira17xpfvakm2amg962yls6f84z3kell8c5lqkfw2s",
        "pub_key": null,
        "account_number": "3",
        "sequence": "0"
      },
      "name": "fee_collector",
      "permissions": []
    }
  ],
  "pagination": {
    "next_key": "70V8dQvo1DIIoWZwiWl6lpuJ2s0=",
    "total": "0"
  }
}
```

Other usages:
```
sekaid query auth accounts --limit=2 --offset=4 --reverse --height=80000 --count-total --output=json | jq
```

```
sekaid query auth accounts --limit=2 --page=2 --reverse --height=80000 --count-total --output=json | jq
```

```
sekaid query auth accounts --limit=2 --page-key="<next_key>" --reverse --height=80000 --count-total --output=json | jq
```

**Pay attention**
Can't use together:
- `--page` and `--page-key`
- `--offset` and `--page`
- `--height` for future blocks

[Return to top](#sekai)

##### 14.2.3 module-account

Query module account info by module name

Usage:
```
sekaid query auth module-account [module-name] [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for module-account                                                                          | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--home string`       | Directory for config and data (default `"/root/.sekaid"`)                              | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?       |
```

/# sekaid query auth module-account --help
Query module account info by module name

Usage:
  sekaid query auth module-account [module-name] [flags]

Examples:
sekaid q auth module-account auth

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for module-account
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

#TODO haven't tested (unknown module names)

🟨  
🟨  
🟨  

[Return to top](#sekai)

##### 14.2.4 params

Query the current auth parameters.

Usage:
```
sekaid query auth params [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for params                                                                                  | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--home string`       | Directory for config and data (default `"/root/.sekaid"`)                              | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid query auth params --help
Query the current auth parameters:

$ <appd> query auth params

Usage:
  sekaid query auth params [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for params
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
/# sekaid query auth params --output=json | jq
{
  "max_memo_characters": "256",
  "tx_sig_limit": "7",
  "tx_size_cost_per_byte": "10",
  "sig_verify_cost_ed25519": "590",
  "sig_verify_cost_secp256k1": "1000"
}
```

[Return to top](#sekai)

#### 14.3 bank

Querying commands for the bank module.

Usage:
```
Usage:
  sekaid query bank [flags]
  sekaid query bank [command]
```

Available Commands:

| Subcommand                               | Description                                      |
| ---------------------------------------- | ------------------------------------------------ |
| [`balances`](#1431-balances)             | Query for account balances by address            |
| [`denom-metadata`](#1432-denom-metadata) | Query the client metadata for coin denominations |
| [`total`](#1433-total)                   | Query the total supply of coins of the chain     |

```
/# sekaid query bank --help
Querying commands for the bank module

Usage:
  sekaid query bank [flags]
  sekaid query bank [command]

Available Commands:
  balances       Query for account balances by address
  denom-metadata Query the client metadata for coin denominations
  total          Query the total supply of coins of the chain

Flags:
  -h, --help   help for bank

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors

Use "sekaid query bank [command] --help" for more information about a command.
```

[Return to top](#sekai)

##### 14.3.1 balances

Query the total balance of an account or of a specific denomination.

Usage:
```
sekaid query bank balances [address] [flags]
```

| Flags                 | Description                                                                                         | Work  |
| --------------------- | --------------------------------------------------------------------------------------------------- | ----- |
| `--count-total`       | count total number of records in all balances to query for                                          | ❌ ?   |
| `--denom string`      | The specific balance denomination to query for                                                      | ❌ ?   |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)               | ❌ ?   |
| `-h, --help`          | help for balances                                                                                   | ✅ yes |
| `--limit uint`        | pagination limit of all balances to query for (default `100`)                                       | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`)    | ✅ yes |
| `--offset uint`       | pagination offset of all balances to query for                                                      | ❌ ?   |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                     | ✅ yes |
| `--page uint`         | pagination page of all balances to query for. This sets offset to a multiple of limit (default `1`) | ❌ ?   |
| `--page-key string`   | pagination page-key of all balances to query for                                                    | ❌ ?   |
| `--reverse`           | results are sorted in descending order                                                              | ❌ ?   |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--home string`       | The application home directory (default `"/root/.sekaid"`)                             | ✅ yes |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid query bank balances --help
Query the total balance of an account or of a specific denomination.

Example:
  $ sekaid query bank balances [address]
  $ sekaid query bank balances [address] --denom=[denom]

Usage:
  sekaid query bank balances [address] [flags]

Flags:
      --count-total       count total number of records in all balances to query for
      --denom string      The specific balance denomination to query for
      --height int        Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help              help for balances
      --limit uint        pagination limit of all balances to query for (default 100)
      --node string       <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
      --offset uint       pagination offset of all balances to query for
  -o, --output string     Output format (text|json) (default "text")
      --page uint         pagination page of all balances to query for. This sets offset to a multiple of limit (default 1)
      --page-key string   pagination page-key of all balances to query for
      --reverse           results are sorted in descending order

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
/# sekaid query bank balances kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --output=json | jq
{
  "balances": [
    {
      "denom": "lol",
      "amount": "899900"
    },
    {
      "denom": "samolean",
      "amount": "1999999999999999999999900100"
    },
    {
      "denom": "test",
      "amount": "29999779999900000"
    },
    {
      "denom": "ukex",
      "amount": "299998797027396"
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "0"
  }
}
```

[Return to top](#sekai)

##### 14.3.2 denom-metadata

Query the client metadata for all the registered coin denominations

Usage:
```
sekaid query bank denom-metadata [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--denom string`      | The specific denomination to query client metadata for                                           | ❌ ?   |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ❌ ?   |
| `-h, --help`          | help for denom-metadata                                                                          | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--home string`       | The application home directory (default `"/root/.sekaid"`)                             | ✅ yes |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid query bank denom-metadata --help
Query the client metadata for all the registered coin denominations

Example:
  To query for the client metadata of all coin denominations use:
  $ sekaid query bank denom-metadata

To query for the client metadata of a specific coin denomination use:
  $ sekaid query bank denom-metadata --denom=[denom]

Usage:
  sekaid query bank denom-metadata [flags]

Flags:
      --denom string    The specific denomination to query client metadata for
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for denom-metadata
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

#TODO show real usage \[currently there is no metadata\]

🟨  
🟨  
🟨  

[Return to top](#sekai)

##### 14.3.3 total

Query total supply of coins that are held by accounts in the chain.

Usage:
```
sekaid query bank total [flags]
```

| Flags                 | Description                                                                                         | Work  |
| --------------------- | --------------------------------------------------------------------------------------------------- | ----- |
| `--count-total`       | count total number of records in all balances to query for                                          | ❌ ?   |
| `--denom string`      | The specific balance denomination to query for                                                      | ❌ ?   |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)               | ❌ ?   |
| `-h, --help`          | help for total                                                                                      | ✅ yes |
| `--limit uint`        | pagination limit of all balances to query for (default `100`)                                       | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`)    | ✅ yes |
| `--offset uint`       | pagination offset of all balances to query for                                                      | ❌ ?   |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                     | ✅ yes |
| `--page uint`         | pagination page of all balances to query for. This sets offset to a multiple of limit (default `1`) | ❌ ?   |
| `--page-key string`   | pagination page-key of all balances to query for                                                    | ❌ ?   |
| `--reverse`           | results are sorted in descending order                                                              | ❌ ?   |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--home string`       | The application home directory (default `"/root/.sekaid"`)                             | ✅ yes |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid query bank total --help
Query total supply of coins that are held by accounts in the chain.

Example:
  $ sekaid query bank total

To query for the total supply of a specific coin denomination use:
  $ sekaid query bank total --denom=[denom]

Usage:
  sekaid query bank total [flags]

Flags:
      --count-total       count total number of records in all supply totals to query for
      --denom string      The specific balance denomination to query for
      --height int        Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help              help for total
      --limit uint        pagination limit of all supply totals to query for (default 100)
      --node string       <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
      --offset uint       pagination offset of all supply totals to query for
  -o, --output string     Output format (text|json) (default "text")
      --page uint         pagination page of all supply totals to query for. This sets offset to a multiple of limit (default 1)
      --page-key string   pagination page-key of all supply totals to query for
      --reverse           results are sorted in descending order

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
/# sekaid query bank total --node tcp://10.1.0.2:26657 -o json | jq
{
  "supply": [
    {
      "denom": "lol",
      "amount": "2000000"
    },
    {
      "denom": "samolean",
      "amount": "5000000000000000000000000000"
    },
    {
      "denom": "test",
      "amount": "29999990000000000"
    },
    {
      "denom": "ukex",
      "amount": "300790043210638"
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "0"
  }
}
```

[Return to top](#sekai)

#### 14.4 basket

[Return to top](#sekai)

#### 14.5 block

[Return to top](#sekai)

#### 14.6 collectives

[Return to top](#sekai)

#### 14.7 custody

[Return to top](#sekai)

#### 14.8 customevidence

[Return to top](#sekai)

#### 14.9 customgov

query commands for the customgov module.

Usage:
```
sekaid query customgov [command]
```

Available Commands:

| Subcommand                                                                                            | Description                                                                                       |
| ----------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| [`all-execution-fees`](#1491-all-execution-fees)                                                      | Query all execution fees                                                                          |
| [`all-identity-record-verify-requests`](#1492-all-identity-record-verify-requests)                    | Query all identity records verify requests                                                        |
| [`all-proposal-durations`](#1493-all-proposal-durations)                                              | Query all proposal durations                                                                      |
| [`all-roles`](#1494-all-roles)                                                                        | Query all registered roles                                                                        |
| [`blacklisted-permission-addresses`](#1495-blacklisted-permission-addresses)                          | Query all KIRA addresses by a specific blacklisted permission                                     |
| [`council-registry`](#1496-council-registry)                                                          | Query governance registry.                                                                        |
| [`councilors`](#1497-councilors)                                                                      | Query councilors                                                                                  |
| [`data-registry`](#1498-data-registry)                                                                | Query data registry by specific key                                                               |
| [`data-registry-keys`](#1499-data-registry-keys)                                                      | Query all data registry keys                                                                      |
| [`execution-fee`](#14910-execution-fee)                                                               | Query execution fee by the type of transaction                                                    |
| [`identity-record`](#14911-identity-record)                                                           | Query identity record by id                                                                       |
| [`identity-record-verify-request`](#14912-identity-record-verify-request)                             | Query identity record verify request by id                                                        |
| [`identity-record-verify-requests-by-approver`](#14913-identity-record-verify-requests-by-approver)   | Query identity record verify request by approver                                                  |
| [`identity-record-verify-requests-by-requester`](#14914-identity-record-verify-requests-by-requester) | Query identity records verify requests by requester                                               |
| [`identity-records`](#14915-identity-records)                                                         | Query all identity records                                                                        |
| [`identity-records-by-addr`](#14916-identity-records-by-addr)                                         | Query identity records by address                                                                 |
| [`network-properties`](#14917-network-properties)                                                     | Query network properties                                                                          |
| [`non-councilors`](#14918-non-councilors)                                                             | Query all governance members that are NOT Councilors                                              |
| [`permissions`](#14919-permissions)                                                                   | Query permissions of an address                                                                   |
| [`poll-votes`](#14920-poll-votes)                                                                     | Get poll votes by id                                                                              |
| [`polls`](#14921-polls)                                                                               | Get polls by address                                                                              |
| [`poor-network-messages`](#14922-poor-network-messages)                                               | Query poor network messages                                                                       |
| [`proposal`](#14923-proposal)                                                                         | Query proposal details                                                                            |
| [`proposal-duration`](#14924-proposal-duration)                                                       | Query a proposal duration                                                                         |
| [`proposals`](#14925-proposals)                                                                       | Query proposals with optional filters                                                             |
| [`proposer_voters_count`](#14926-proposer_voters_count)                                               | Query proposer and voters count that can create at least a type of proposal                       |
| [`role`](#14927-role)                                                                                 | Query role by sid or id                                                                           |
| [`roles`](#14928-roles)                                                                               | Query roles assigned to an address                                                                |
| [`vote`](#14929-vote)                                                                                 | Query details of a single vote                                                                    |
| [`voters`](#14930-voters)                                                                             | Query voters of a proposal                                                                        |
| [`votes`](#14931-votes)                                                                               | Query votes on a proposal                                                                         |
| [`whitelisted-permission-addresses](#14932-whitelisted-permission-addresses)                          | Query all KIRA addresses by a specific whitelisted permission                                     |
| [`whitelisted-role-addresses`](#14933-whitelisted-role-addresses)                                     | Query all kira addresses by a specific whitelisted role (address does NOT have to be a Councilor) |



| Flags        | Description        | Work  |
| ------------ | ------------------ | ----- |
| `-h, --help` | help for customgov | ✅ yes |



| Global Flags          | Description                                                                            | Work |
| --------------------- | -------------------------------------------------------------------------------------- | ---- |
| `--chain-id string`   | The network chain ID                                                                   | ❌ ?  |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ❌ ?  |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?  |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?  |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?  |

```
sekaid q customgov --help
query commands for the customgov module

Usage:
  sekaid query customgov [command]

Available Commands:
  all-execution-fees                           Query all execution fees
  all-identity-record-verify-requests          Query all identity records verify requests
  all-proposal-durations                       Query all proposal durations
  all-roles                                    Query all registered roles
  blacklisted-permission-addresses             Query all KIRA addresses by a specific blacklisted permission
  council-registry                             Query governance registry.
  councilors                                   Query councilors
  data-registry                                Query data registry by specific key
  data-registry-keys                           Query all data registry keys
  execution-fee                                Query execution fee by the type of transaction
  identity-record                              Query identity record by id
  identity-record-verify-request               Query identity record verify request by id
  identity-record-verify-requests-by-approver  Query identity record verify request by approver
  identity-record-verify-requests-by-requester Query identity records verify requests by requester
  identity-records                             Query all identity records
  identity-records-by-addr                     Query identity records by address
  network-properties                           Query network properties
  non-councilors                               Query all governance members that are NOT Councilors
  permissions                                  Query permissions of an address
  poll-votes                                   Get poll votes by id
  polls                                        Get polls by address
  poor-network-messages                        Query poor network messages
  proposal                                     Query proposal details
  proposal-duration                            Query a proposal duration
  proposals                                    Query proposals with optional filters
  proposer_voters_count                        Query proposer and voters count that can create at least a type of proposal
  role                                         Query role by sid or id
  roles                                        Query roles assigned to an address
  vote                                         Query details of a single vote
  voters                                       Query voters of a proposal
  votes                                        Query votes on a proposal
  whitelisted-permission-addresses             Query all KIRA addresses by a specific whitelisted permission
  whitelisted-role-addresses                   Query all kira addresses by a specific whitelisted role (address does NOT have to be a Councilor)

Flags:
  -h, --help   help for customgov

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors

Use "sekaid query customgov [command] --help" for more information about a command.
```

[Return to top](#sekai)

##### 14.9.1 all-execution-fees

Query all execution fees.

Usage:
```
sekaid query customgov all-execution-fees [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for all-execution-fees                                                                      | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid q customgov all-execution-fees --help
Query all execution fees

Usage:
  sekaid query customgov all-execution-fees [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for all-execution-fees
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid q customgov all-execution-fees -o json | jq
{
  "fees": [
    {
      "transaction_type": "activate",
      "execution_fee": "100",
      "failure_fee": "1000",
      "timeout": "10",
      "default_parameters": "0"
    },
    {
      "transaction_type": "claim-councilor",
      "execution_fee": "100",
      "failure_fee": "1",
      "timeout": "10",
      "default_parameters": "0"
    },
    {
      "transaction_type": "claim-proposal-type-x",
      "execution_fee": "100",
      "failure_fee": "1",
      "timeout": "10",
      "default_parameters": "0"
    },
    {
      "transaction_type": "claim-validator",
      "execution_fee": "100",
      "failure_fee": "1",
      "timeout": "10",
      "default_parameters": "0"
    },
    {
      "transaction_type": "pause",
      "execution_fee": "100",
      "failure_fee": "100",
      "timeout": "10",
      "default_parameters": "0"
    },
    {
      "transaction_type": "submit-proposal-type-x",
      "execution_fee": "10",
      "failure_fee": "1",
      "timeout": "10",
      "default_parameters": "0"
    },
    {
      "transaction_type": "unpause",
      "execution_fee": "100",
      "failure_fee": "100",
      "timeout": "10",
      "default_parameters": "0"
    },
    {
      "transaction_type": "upsert-token-alias",
      "execution_fee": "100",
      "failure_fee": "1",
      "timeout": "10",
      "default_parameters": "0"
    },
    {
      "transaction_type": "veto-proposal-type-x",
      "execution_fee": "100",
      "failure_fee": "1",
      "timeout": "10",
      "default_parameters": "0"
    },
    {
      "transaction_type": "vote-proposal-type-x",
      "execution_fee": "100",
      "failure_fee": "1",
      "timeout": "10",
      "default_parameters": "0"
    }
  ],
  "tx_types": [
    "remove-blacklist-role-permission",
    "deposit-spending-pool",
    "remove-from-custody-whitelist",
    "submit_evidence",
    "whitelist-permissions",
    "create-role",
    "blacklist-role-permission",
    "create-spending-pool",
    "register-spending-pool-beneficiary",
    "remove-role",
    "claim_rewards",
    "request-identity-records-verify",
    "set-network-properties",
    "blacklist-permissions",
    "whitelist-role-permission",
    "remove-whitelist-role-permission",
    "claim-validator",
    "pause",
    "delegate",
    "submit-proposal",
    "cancel-identity-records-verify-request",
    "assign-role",
    "decline-custody-transaction",
    "undelegate",
    "drop-custody-whitelist",
    "approve-custody-transaction",
    "activate",
    "register_delegator",
    "add-to-custody-custodians",
    "register-identity-records",
    "claim_undelegation",
    "set_compound_info",
    "handle-identity-records-verify-request",
    "set-execution-fee",
    "upsert-token-rate",
    "unpause",
    "claim-spending-pool",
    "multisend",
    "vote-proposal",
    "edit-identity-record",
    "create-custody",
    "add-to-custody-whitelist",
    "upsert_staking_pool",
    "remove-from-custody-custodians",
    "drop-custody-custodians",
    "send",
    "claim-councilor",
    "upsert-token-alias"
  ]
}
```

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.2 all-identity-record-verify-requests

Query all identity records verify requests.

Usage:
```
sekaid query customgov all-identity-record-verify-requests [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--count-total`       | count total number of records in customgov to query for                                          | ✅ yes |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for all-identity-record-verify-requests                                                     | ✅ yes |
| `--limit uint`        | pagination limit of customgov to query for (default `100`)                                       | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `--offset uint`       | pagination offset of customgov to query for                                                      | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |
| `--page uint`         | pagination page of customgov to query for. This sets offset to a multiple of limit (default `1`) | ✅ yes |
| `--page-key string`   | pagination page-key of customgov to query for                                                    | ✅ yes |
| `--reverse`           | results are sorted in descending order                                                           | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid q customgov all-identity-record-verify-requests --help
Query all identity records verify requests.

Example:
$ sekaid query gov all-identity-record-verify-requests

Usage:
  sekaid query customgov all-identity-record-verify-requests [flags]

Flags:
      --count-total       count total number of records in customgov to query for
      --height int        Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help              help for all-identity-record-verify-requests
      --limit uint        pagination limit of customgov to query for (default 100)
      --node string       <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
      --offset uint       pagination offset of customgov to query for
  -o, --output string     Output format (text|json) (default "text")
      --page uint         pagination page of customgov to query for. This sets offset to a multiple of limit (default 1)
      --page-key string   pagination page-key of customgov to query for
      --reverse           results are sorted in descending order

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
/# sekaid q customgov all-identity-record-verify-requests -o json | jq
{
  "verify_records": [
    {
      "id": "3",
      "address": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
      "verifier": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
      "recordIds": [
        "15"
      ],
      "tip": {
        "denom": "ukex",
        "amount": "200"
      },
      "lastRecordEditDate": "2023-05-25T13:18:03.396177800Z"
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "0"
  }
}
```

Other usages:
```
sekaid q customgov all-identity-record-verify-requests --limit=2 --offset=4 --reverse --height=80000 --count-total --output=json | jq
```

```
sekaid q customgov all-identity-record-verify-requests --limit=2 --page=2 --reverse --height=80000 --count-total --output=json | jq
```

```
sekaid q customgov all-identity-record-verify-requests --limit=2 --page-key="<next_key>" --reverse --height=80000 --count-total --output=json | jq
```

**Pay attention**
Can't use together:
- `--page` and `--page-key`
- `--offset` and `--page`
- `--height` for future blocks

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.3 all-proposal-durations

Query all proposal durations.

Usage:
```
sekaid query customgov all-proposal-durations [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for all-proposal-durations                                                                  | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid q customgov all-proposal-durations --help
Query all proposal durations.

Example:
$ sekaid query gov all-proposal-durations

Usage:
  sekaid query customgov all-proposal-durations [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for all-proposal-durations
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid q customgov all-proposal-durations -o json | jq
{
  "proposal_durations": {
    "AssignRoleToAccount": "360",
    "BlacklistAccountPermission": "360",
    "BlacklistRolePermission": "360",
    "CancelSoftwareUpgrade": "360",
    "CreateRoleProposal": "360",
    "RemoveBlacklistedAccountPermission": "360",
    "RemoveBlacklistedRolePermission": "360",
    "RemoveRoleProposal": "360",
    "RemoveWhitelistedAccountPermission": "360",
    "RemoveWhitelistedRolePermission": "360",
    "ResetWholeCouncilorRank": "360",
    "ResetWholeValidatorRank": "360",
    "SetNetworkProperty": "360",
    "SetPoorNetworkMessages": "360",
    "SetProposalDurationsProposal": "360",
    "SlashValidator": "360",
    "SoftwareUpgrade": "360",
    "TokensWhiteBlackChange": "360",
    "UnassignRoleFromAccount": "360",
    "UnjailValidator": "360",
    "UpsertDataRegistry": "360",
    "UpsertTokenAlias": "360",
    "UpsertTokenRates": "360",
    "WhitelistAccountPermission": "360",
    "WhitelistRolePermission": "360"
  }
}
```

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.4 all-roles

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.5 blacklisted-permission-addresses

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.6 council-registry

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.7 councilors

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.8 data-registry

Query data registry by key.

Usage:
```
sekaid query customgov data-registry [flags] ❌ mistake! here we need the one positional argument
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--count-total`       | count total number of records in customgov to query for                                          | ✅ yes |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for data-registry                                                                           | ✅ yes |
| `--limit uint`        | pagination limit of customgov to query for (default `100`)                                       | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `--offset uint`       | pagination offset of customgov to query for                                                      | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |
| `--page uint`         | pagination page of customgov to query for. This sets offset to a multiple of limit (default `1`) | ✅ yes |
| `--page-key string`   | pagination page-key of customgov to query for                                                    | ✅ yes |
| `--reverse`           | results are sorted in descending order                                                           | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid query customgov data-registry --help
Query data registry by key.

Example:
$ sekaid query gov data-registry [key]

Usage:
  sekaid query customgov data-registry [flags]

Flags:
      --count-total       count total number of records in customgov to query for
      --height int        Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help              help for data-registry
      --limit uint        pagination limit of customgov to query for (default 100)
      --node string       <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
      --offset uint       pagination offset of customgov to query for
  -o, --output string     Output format (text|json) (default "text")
      --page uint         pagination page of customgov to query for. This sets offset to a multiple of limit (default 1)
      --page-key string   pagination page-key of customgov to query for
      --reverse           results are sorted in descending order

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
/# sekaid query customgov data-registry icon -o json | jq
{
  "data": {
    "hash": "891bd9d3b2ee0c6eed43a8129b096bebc7e5ae517d0b855b2116a3205211fe21",
    "reference": "https://kira-network.s3-eu-west-1.amazonaws.com/assets/img/tokens/kex.svg",
    "encoding": "picture",
    "size": "1597"
  }
}
```

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.9 data-registry-keys

Query all data registry keys.

Usage:
```
sekaid query customgov data-registry-keys [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--count-total`       | count total number of records in customgov to query for                                          | ✅ yes |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for data-registry-keys                                                                      | ✅ yes |
| `--limit uint`        | pagination limit of customgov to query for (default `100`)                                       | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `--offset uint`       | pagination offset of customgov to query for                                                      | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |
| `--page uint`         | pagination page of customgov to query for. This sets offset to a multiple of limit (default `1`) | ✅ yes |
| `--page-key string`   | pagination page-key of customgov to query for                                                    | ✅ yes |
| `--reverse`           | results are sorted in descending order                                                           | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid query customgov data-registry-keys --help
Query all data registry keys.

Example:
$ sekaid query gov data-registry-keys

Usage:
  sekaid query customgov data-registry-keys [flags]

Flags:
      --count-total       count total number of records in customgov to query for
      --height int        Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help              help for data-registry-keys
      --limit uint        pagination limit of customgov to query for (default 100)
      --node string       <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
      --offset uint       pagination offset of customgov to query for
  -o, --output string     Output format (text|json) (default "text")
      --page uint         pagination page of customgov to query for. This sets offset to a multiple of limit (default 1)
      --page-key string   pagination page-key of customgov to query for
      --reverse           results are sorted in descending order

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid query customgov data-registry-keys --output=json | jq
{
  "keys": [
    "icon"
  ],
  "pagination": {
    "next_key": null,
    "total": "0"
  }
}
```

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.10 execution-fee

Query execution fee by the type of transaction.

Usage:
```
sekaid query customgov execution-fee [transaction_type] [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for execution-fee                                                                           | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid q customgov execution-fee --help
Query execution fee by the type of transaction

Usage:
  sekaid query customgov execution-fee [transaction_type] [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for execution-fee
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid q customgov execution-fee pause -o json | jq
{
  "fee": {
    "transaction_type": "pause",
    "execution_fee": "100",
    "failure_fee": "100",
    "timeout": "10",
    "default_parameters": "0"
  }
}
```

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.11 identity-record

Query identity record by id.

Usage:
```
sekaid query customgov identity-record [id] [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for identity-record                                                                         | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid query customgov identity-record --help
Query identity record by id.

Example:
$ sekaid query gov identity-record 1

Usage:
  sekaid query customgov identity-record [id] [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for identity-record
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
/# sekaid query customgov identity-record 1 -o json | jq
{
  "record": {
    "id": "1",
    "address": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
    "key": "moniker",
    "value": "GENESIS VALIDATOR",
    "date": "2023-05-25T12:56:59.866789362Z",
    "verifiers": []
  }
}
```

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.12 identity-record-verify-request

Query identity record verify request by id.

Usage:
```
sekaid query customgov identity-record-verify-request [id] [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for identity-record-verify-request                                                          | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid query customgov identity-record-verify-request --help
Query identity record verify request by id.

Example:
$ sekaid query gov identity-record-verify-request 1

Usage:
  sekaid query customgov identity-record-verify-request [id] [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for identity-record-verify-request
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
/# sekaid query customgov identity-record-verify-request 3 -o json | jq
{
  "verify_record": {
    "id": "3",
    "address": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
    "verifier": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
    "recordIds": [
      "15"
    ],
    "tip": {
      "denom": "ukex",
      "amount": "200"
    },
    "lastRecordEditDate": "2023-05-25T13:18:03.396177800Z"
  }
}
```

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.13 identity-record-verify-requests-by-approver

Query identity record verify requests by approver.

Usage:
```
sekaid query customgov identity-record-verify-requests-by-approver [addr] [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--count-total`       | count total number of records in customgov to query for                                          | ✅ yes |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for identity-record-verify-requests-by-approver                                             | ✅ yes |
| `--limit uint`        | pagination limit of customgov to query for (default `100`)                                       | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `--offset uint`       | pagination offset of customgov to query for                                                      | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |
| `--page uint`         | pagination page of customgov to query for. This sets offset to a multiple of limit (default `1`) | ✅ yes |
| `--page-key string`   | pagination page-key of customgov to query for                                                    | ✅ yes |
| `--reverse`           | results are sorted in descending order                                                           | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid query customgov identity-record-verify-requests-by-approver --help
Query identity record verify requests by approver.

Example:
$ sekaid query gov identity-record-verify-requests-by-approver [addr]

Usage:
  sekaid query customgov identity-record-verify-requests-by-approver [addr] [flags]

Flags:
      --count-total       count total number of records in customgov to query for
      --height int        Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help              help for identity-record-verify-requests-by-approver
      --limit uint        pagination limit of customgov to query for (default 100)
      --node string       <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
      --offset uint       pagination offset of customgov to query for
  -o, --output string     Output format (text|json) (default "text")
      --page uint         pagination page of customgov to query for. This sets offset to a multiple of limit (default 1)
      --page-key string   pagination page-key of customgov to query for
      --reverse           results are sorted in descending order

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
/# sekaid query customgov identity-record-verify-requests-by-approver kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --output=json | jq
{
  "verify_records": [
    {
      "id": "3",
      "address": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
      "verifier": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
      "recordIds": [
        "15"
      ],
      "tip": {
        "denom": "ukex",
        "amount": "200"
      },
      "lastRecordEditDate": "2023-05-25T13:18:03.396177800Z"
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "0"
  }
}
```

Other usages:
```
sekaid query customgov identity-record-verify-requests-by-approver kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --limit=2 --offset=4 --reverse --height=80000 --count-total --output=json | jq
```

```
sekaid query customgov identity-record-verify-requests-by-approver kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --limit=2 --page=2 --reverse --height=80000 --count-total --output=json | jq
```

```
sekaid query customgov identity-record-verify-requests-by-approver kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --limit=2 --page-key="<next_key>" --reverse --height=80000 --count-total --output=json | jq
```

**Pay attention**
Can't use together:
- `--page` and `--page-key`
- `--offset` and `--page`
- `--height` for future blocks

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.14 identity-record-verify-requests-by-requester

Query identity records verify requests by requester.

Usage:
```
sekaid query customgov identity-record-verify-requests-by-requester [addr] [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--count-total`       | count total number of records in customgov to query for                                          | ✅ yes |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for identity-record-verify-requests-by-requester                                            | ✅ yes |
| `--limit uint`        | pagination limit of customgov to query for (default `100`)                                       | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `--offset uint`       | pagination offset of customgov to query for                                                      | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |
| `--page uint`         | pagination page of customgov to query for. This sets offset to a multiple of limit (default `1`) | ✅ yes |
| `--page-key string`   | pagination page-key of customgov to query for                                                    | ✅ yes |
| `--reverse`           | results are sorted in descending order                                                           | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid query customgov identity-record-verify-requests-by-requester --help
Query identity records verify requests by requester.

Example:
$ sekaid query gov identity-record-verify-requests-by-requester [addr]

Usage:
  sekaid query customgov identity-record-verify-requests-by-requester [addr] [flags]

Flags:
      --count-total       count total number of records in customgov to query for
      --height int        Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help              help for identity-record-verify-requests-by-requester
      --limit uint        pagination limit of customgov to query for (default 100)
      --node string       <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
      --offset uint       pagination offset of customgov to query for
  -o, --output string     Output format (text|json) (default "text")
      --page uint         pagination page of customgov to query for. This sets offset to a multiple of limit (default 1)
      --page-key string   pagination page-key of customgov to query for
      --reverse           results are sorted in descending order

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
/# sekaid query customgov identity-record-verify-requests-by-requester kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt -o json | jq
{
  "verify_records": [
    {
      "id": "3",
      "address": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
      "verifier": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
      "recordIds": [
        "15"
      ],
      "tip": {
        "denom": "ukex",
        "amount": "200"
      },
      "lastRecordEditDate": "2023-05-25T13:18:03.396177800Z"
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "0"
  }
}
```

Other usages:
```
sekaid query customgov identity-record-verify-requests-by-requester kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt --limit=2 --offset=4 --reverse --height=80000 --count-total --output=json | jq
```

```
sekaid query customgov identity-record-verify-requests-by-requester kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt --limit=2 --page=2 --reverse --height=80000 --count-total --output=json | jq
```

```
sekaid query customgov identity-record-verify-requests-by-requester kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt --limit=2 --page-key="<next_key>" --reverse --height=80000 --count-total --output=json | jq
```

**Pay attention**
Can't use together:
- `--page` and `--page-key`
- `--offset` and `--page`
- `--height` for future blocks

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.15 identity-records

Query all identity records.

Usage:
```
sekaid query customgov identity-records [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--count-total`       | count total number of records in customgov to query for                                          | ✅ yes |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for identity-records                                                                        | ✅ yes |
| `--limit uint`        | pagination limit of customgov to query for (default `100`)                                       | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `--offset uint`       | pagination offset of customgov to query for                                                      | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |
| `--page uint`         | pagination page of customgov to query for. This sets offset to a multiple of limit (default `1`) | ✅ yes |
| `--page-key string`   | pagination page-key of customgov to query for                                                    | ✅ yes |
| `--reverse`           | results are sorted in descending order                                                           | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid q customgov identity-records --help
Query all identity records.

Example:
$ sekaid query gov identity-records

Usage:
  sekaid query customgov identity-records [flags]

Flags:
      --count-total       count total number of records in customgov to query for
      --height int        Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help              help for identity-records
      --limit uint        pagination limit of customgov to query for (default 100)
      --node string       <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
      --offset uint       pagination offset of customgov to query for
  -o, --output string     Output format (text|json) (default "text")
      --page uint         pagination page of customgov to query for. This sets offset to a multiple of limit (default 1)
      --page-key string   pagination page-key of customgov to query for
      --reverse           results are sorted in descending order

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid q customgov identity-records -o json | jq
{
  "records": [
    {
      "id": "1",
      "address": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
      "key": "moniker",
      "value": "GENESIS VALIDATOR",
      "date": "2023-05-25T12:56:59.866789362Z",
      "verifiers": []
    },

    // ...

    {
      "id": "17",
      "address": "kira1nqcqufejredawp25ycxqh37983n65pejpfyg9e",
      "key": "moniker",
      "value": "validator4",
      "date": "2023-05-25T13:19:35.602583274Z",
      "verifiers": [
        "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4"
      ]
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "0"
  }
}
```

Other usages:
```
sekaid q customgov identity-records --limit=2 --offset=4 --reverse --height=80000 --count-total --output=json | jq
```

```
sekaid q customgov identity-records --limit=2 --page=2 --reverse --height=80000 --count-total --output=json | jq
```

```
sekaid q customgov identity-records --limit=2 --page-key="<next_key>" --reverse --height=80000 --count-total --output=json | jq
```

**Pay attention**
Can't use together:
- `--page` and `--page-key`
- `--offset` and `--page`
- `--height` for future blocks

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.16 identity-records-by-addr

Query identity records by address.

Usage:
```
sekaid query customgov identity-records-by-addr [addr] [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for identity-records-by-addr                                                                | ✅ yes |
| `--keys string`       | keys required when needs to be filtered                                                          | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |


| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--home string`       | The application home directory (default `"/root/.sekaid"`)                             | ✅ ignored |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid query customgov identity-records-by-addr --help
Query identity records by address.

Example:
$ sekaid query gov identity-records-by-addr [addr]

Usage:
  sekaid query customgov identity-records-by-addr [addr] [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for identity-records-by-addr
      --keys string     keys required when needs to be filtered
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid query customgov identity-records-by-addr kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt -o json | jq
{
  "records": [
    {
      "id": "15",
      "address": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
      "key": "moniker",
      "value": "validator2",
      "date": "2023-05-25T13:18:03.396177800Z",
      "verifiers": []
    }
  ]
}
```

Filter by keys:
```
/# sekaid query customgov identity-records-by-addr kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt --keys=moniker -o json | jq
{
  "records": [
    {
      "id": "15",
      "address": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
      "key": "moniker",
      "value": "validator2",
      "date": "2023-05-25T13:18:03.396177800Z",
      "verifiers": []
    }
  ]
}
```

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.17 network-properties

Query network properties.

Usage:
```
sekaid query customgov network-properties [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for alias                                                                                   | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ ignored |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid query customgov network-properties --help
Query network properties

Usage:
  sekaid query customgov network-properties [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for network-properties
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
/# sekaid query customgov network-properties --output=json | jq
{
  "properties": {
    "min_tx_fee": "99",
    "max_tx_fee": "1000000",
    "vote_quorum": "33",
    "minimum_proposal_end_time": "360",
    "proposal_enactment_time": "300",
    "min_proposal_end_blocks": "2",
    "min_proposal_enactment_blocks": "1",
    "enable_foreign_fee_payments": true,
    "mischance_rank_decrease_amount": "1",
    "max_mischance": "50",
    "mischance_confidence": "25",
    "inactive_rank_decrease_percent": "0.500000000000000000",
    "min_validators": "1",
    "poor_network_max_bank_send": "1000000",
    "unjail_max_time": "1209600",
    "enable_token_whitelist": false,
    "enable_token_blacklist": true,
    "min_identity_approval_tip": "200",
    "unique_identity_keys": "moniker,username",
    "ubi_hardcap": "6000000",
    "validators_fee_share": "0.500000000000000000",
    "inflation_rate": "0.180000000000000000",
    "inflation_period": "31557600",
    "unstaking_period": "2629800",
    "max_delegators": "100",
    "min_delegation_pushout": "10",
    "slashing_period": "3600",
    "max_jailed_percentage": "0.250000000000000000",
    "max_slashing_percentage": "0.010000000000000000",
    "min_custody_reward": "200",
    "max_custody_buffer_size": "10",
    "max_custody_tx_size": "8192",
    "abstention_rank_decrease_amount": "1",
    "max_abstention": "2",
    "min_collective_bond": "100000",
    "min_collective_bonding_time": "86400",
    "max_collective_outputs": "10",
    "min_collective_claim_period": "14400",
    "validator_recovery_bond": "300000",
    "max_annual_inflation": "0.350000000000000000",
    "max_proposal_title_size": "128",
    "max_proposal_description_size": "1024",
    "max_proposal_poll_option_size": "64",
    "max_proposal_poll_option_count": "128",
    "max_proposal_reference_size": "512",
    "max_proposal_checksum_size": "128",
    "min_dapp_bond": "1000000",
    "max_dapp_bond": "10000000",
    "dapp_liquidation_threshold": "0",
    "dapp_liquidation_period": "0",
    "dapp_bond_duration": "604800",
    "dapp_verifier_bond": "0.001000000000000000",
    "dapp_auto_denounce_time": "60",
    "dapp_mischance_rank_decrease_amount": "1",
    "dapp_max_mischance": "10",
    "dapp_inactive_rank_decrease_percent": "10",
    "dapp_pool_slippage_default": "0.100000000000000000",
    "minting_ft_fee": "100000000000000",
    "minting_nft_fee": "100000000000000"
  }
}
```

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.18 non-councilors

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.19 permissions

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.20 poll-votes

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.21 polls

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.22 poor-network-messages

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.23 proposal

Query details for a proposal.

Usage:
```
sekaid query customgov proposal [proposal-id] [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for proposal                                                                                | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ ignored |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid q customgov proposal --help
Query details for a proposal. You can find the
proposal-id by running "sekaid query gov proposals".

Example:
$ sekaid query gov proposal 1

Usage:
  sekaid query customgov proposal [proposal-id] [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for proposal
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid q customgov proposal 13 -o json | jq
{
  "proposal_id": "13",
  "title": "Reducing value of minimal fee for docs",
  "description": "Changing value of minimal fee for testing for docs",
  "content": {
    "@type": "/kira.gov.SetNetworkPropertyProposal",
    "network_property": "MAX_TX_FEE",
    "value": {
      "value": "1000",
      "str_value": ""
    }
  },
  "submit_time": "2023-06-12T09:22:10.251027877Z",
  "voting_end_time": "2023-06-12T09:28:10.251027877Z",
  "enactment_end_time": "2023-06-12T09:33:10.251027877Z",
  "min_voting_end_block_height": "141342",
  "min_enactment_end_block_height": "141376",
  "result": "VOTE_RESULT_PASSED",
  "exec_result": "executed successfully"
}
```

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.24 proposal-duration

Query all all proposal durations.

Usage:
```
sekaid query customgov proposal-duration [proposal_type] [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for proposal-duration                                                                       | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ ignored |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid q customgov proposal-duration --help
Query all all proposal durations.

Example:
$ sekaid query gov proposal-duration SetNetworkProperty

Usage:
  sekaid query customgov proposal-duration [proposal_type] [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for proposal-duration
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid q customgov proposal-duration CreateRoleProposal -o json | jq
{
  "duration": "360"
}
```

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.25 proposals

Query for a all paginated proposals that match optional filters.

Usage:
```
sekaid query customgov proposals [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--count-total`       | count total number of records in customgov to query for                                          | ✅ yes |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for proposals                                                                               | ✅ yes |
| `--limit uint`        | pagination limit of customgov to query for (default `100`)                                       | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `--offset uint`       | pagination offset of customgov to query for                                                      | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |
| `--page uint`         | pagination page of customgov to query for. This sets offset to a multiple of limit (default `1`) | ✅ yes |
| `--page-key string`   | pagination page-key of customgov to query for                                                    | ✅ yes |
| `--reverse`           | results are sorted in descending order                                                           | ✅ yes |
| `--voter string`      | (optional) filter by proposals voted on by voted                                                 | ❌ no  |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ ignored |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid q customgov proposals --help
Query for a all paginated proposals that match optional filters:

Example:
$ sekaid query gov proposals --depositor cosmos1skjwj5whet0lpe65qaq4rpq03hjxlwd9nf39lk
$ sekaid query gov proposals --voter cosmos1skjwj5whet0lpe65qaq4rpq03hjxlwd9nf39lk

Usage:
  sekaid query customgov proposals [flags]

Flags:
      --count-total       count total number of records in customgov to query for
      --height int        Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help              help for proposals
      --limit uint        pagination limit of customgov to query for (default 100)
      --node string       <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
      --offset uint       pagination offset of customgov to query for
  -o, --output string     Output format (text|json) (default "text")
      --page uint         pagination page of customgov to query for. This sets offset to a multiple of limit (default 1)
      --page-key string   pagination page-key of customgov to query for
      --reverse           results are sorted in descending order
      --voter string      (optional) filter by proposals voted on by voted

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid q customgov proposals -o json | jq
{
  "proposals": [
    {
      "proposal_id": "1",
      "title": "Upsert KEX icon URL link",
      "description": "Initial Setup From KIRA Manager",
      "content": {
        "@type": "/kira.tokens.ProposalUpsertTokenAlias",
        "symbol": "KEX",
        "name": "KIRA",
        "icon": "http://kira-network.s3-eu-west-1.amazonaws.com/assets/img/tokens/kex.svg",
        "decimals": 6,
        "denoms": [
          "ukex"
        ],
        "invalidated": false
      },
      "submit_time": "2023-05-25T13:00:51.893680415Z",
      "voting_end_time": "2023-05-25T13:06:51.893680415Z",
      "enactment_end_time": "2023-05-25T13:11:51.893680415Z",
      "min_voting_end_block_height": "25",
      "min_enactment_end_block_height": "58",
      "result": "VOTE_RESULT_PASSED",
      "exec_result": "executed successfully"
    },
    // . . .
    {
      "proposal_id": "28",
      "title": "Test proposal for docs",
      "description": "Test proposal for docs",
      "content": {
        "@type": "/kira.gov.SetProposalDurationsProposal",
        "typeof_proposals": [
          "AssignRoleToAccount",
          "BlacklistAccountPermission"
        ],
        "proposal_durations": [
          "300",
          "310"
        ]
      },
      "submit_time": "2023-06-12T14:41:26.992416500Z",
      "voting_end_time": "2023-06-12T14:47:26.992416500Z",
      "enactment_end_time": "2023-06-12T14:52:26.992416500Z",
      "min_voting_end_block_height": "143199",
      "min_enactment_end_block_height": "143233",
      "result": "VOTE_RESULT_QUORUM_NOT_REACHED",
      "exec_result": ""
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "0"
  }
}
```

Other usages:
```
sekaid q customgov proposals --limit=2 --offset=4 --reverse --height=80000 --count-total --output=json | jq
```

```
sekaid q customgov proposals --limit=2 --page=2 --reverse --height=80000 --count-total --output=json | jq
```

```
sekaid q customgov proposals --limit=2 --page-key="<next_key>" --reverse --height=80000 --count-total --output=json | jq
```

**Pay attention**
Can't use together:
- `--page` and `--page-key`
- `--offset` and `--page`
- `--height` for future blocks

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.26 proposer_voters_count

Query proposer and voters count that can create at least a type of proposal.

Usage:
```
sekaid query customgov proposer_voters_count [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for proposer_voters_count                                                                   | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ ignored |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid q customgov proposer_voters_count --help
Query proposer and voters count that can create at least a type of proposal.

Example:
$ sekaid query gov proposer_voters_count

Usage:
  sekaid query customgov proposer_voters_count [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for proposer_voters_count
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid q customgov proposer_voters_count -o json | jq
{
  "proposers": "4",
  "voters": "2"
}
```

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.27 role

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.28 roles

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.29 vote

Query details for a single vote on a proposal given its identifier.

Usage:
```
sekaid query customgov vote [proposal-id] [voter-addr] [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for vote                                                                                    | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ ignored |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid q customgov vote --help
Query details for a single vote on a proposal given its identifier.

Example:
$ sekaid query gov vote 1 kira1skjwj5whet0lpe65qaq4rpq03hjxlwd9nf39lk

Usage:
  sekaid query customgov vote [proposal-id] [voter-addr] [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for vote
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid query customgov vote 9 kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 -o json | jq
{
  "proposal_id": "9",
  "voter": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
  "option": "VOTE_OPTION_YES",
  "slash": "0.010000000000000000"
}
```

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.30 voters

Query voters for a single proposal by its identifier.

Usage:
```
sekaid query customgov voters [proposal-id] [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for voters                                                                                  | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ ignored |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid q customgov voters --help
Query voters for a single proposal by its identifier.

Example:
$ sekaid query gov voters 1

Usage:
  sekaid query customgov voters [proposal-id] [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for voters
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid query customgov voters 8 -o json | jq
{
  "voters": [
    {
      "address": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
      "roles": [
        "2"
      ],
      "status": "ACTIVE",
      "votes": [
        "VOTE_OPTION_YES",
        "VOTE_OPTION_NO",
        "VOTE_OPTION_ABSTAIN",
        "VOTE_OPTION_NO_WITH_VETO"
      ],
      "permissions": {
        "blacklist": [
          10
        ],
        "whitelist": [
          28,
          12,
          13
        ]
      },
      "skin": "0"
    },
    {
      "address": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
      "roles": [
        "1",
        "2"
      ],
      "status": "ACTIVE",
      "votes": [
        "VOTE_OPTION_YES",
        "VOTE_OPTION_ABSTAIN",
        "VOTE_OPTION_NO",
        "VOTE_OPTION_NO_WITH_VETO"
      ],
      "permissions": {
        "blacklist": [],
        "whitelist": [
          4,
          35,
          14,
          28,
          5,
          36,
          15,
          29
        ]
      },
      "skin": "1"
    }
  ]
}
```

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.31 votes

Query vote details for a single proposal by its identifier.

Usage:
```
sekaid query customgov votes [proposal-id] [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for votes                                                                                   | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ ignored |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid q customgov votes --help
Query vote details for a single proposal by its identifier.

Example:
$ sekaid query gov votes 1

Usage:
  sekaid query customgov votes [proposal-id] [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for votes
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid q customgov votes 8 -o json | jq
{
  "votes": [
    {
      "proposal_id": "8",
      "voter": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
      "option": "VOTE_OPTION_YES",
      "slash": "0.010000000000000000"
    },
    {
      "proposal_id": "8",
      "voter": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
      "option": "VOTE_OPTION_YES",
      "slash": "0.010000000000000000"
    }
  ]
}
```

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.32 whitelisted-permission-addresses

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

##### 14.9.33 whitelisted-role-addresses

[Return to "`query customgov`"](#149-customgov)  
[Return to top](#sekai)

#### 14.10 customslashing

[Return to top](#sekai)

#### 14.11 customstaking

[Return to top](#sekai)

#### 14.12 distributor

[Return to top](#sekai)

#### 14.13 layer2

[Return to top](#sekai)

#### 14.14 multistaking

[Return to top](#sekai)

#### 14.15 params

[Return to top](#sekai)

#### 14.16 recovery

[Return to top](#sekai)

#### 14.17 spending

[Return to top](#sekai)

#### 14.18 tendermint-validator-set

[Return to top](#sekai)

#### 14.19 tokens

Query commands for the tokens module.

Usage:
```
sekaid query tokens [command]
```

Available commands:

| Subcommands                                       | Description                   |
| ------------------------------------------------- | ----------------------------- |
| [`alias`](#14191-alias)                           | Get the token alias by symbol |
| [`aliases-by-denom`](#14192-aliases-by-denom)     | Get token aliases by denom    |
| [`all-aliases`](#14193-all-aliases)               | Get all token aliases         |
| [`all-rates`](#14194-all-rates)                   | Get all token rates           |
| [`rate`](#14195-rate)                             | Get the token rate by denom   |
| [`rates-by-denom`](#14196-rates-by-denom)         | Get token rates by denom      |
| [`token-black-whites`](#14197-token-black-whites) | Get token black whites        |



| Flags        | Description     | Work  |
| ------------ | --------------- | ----- |
| `-h, --help` | help for tokens | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--home string`       | The application home directory (default `"/root/.sekaid"`)                             | ✅ ignored |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid qeury tokens --help
query commands for the tokens module

Usage:
  sekaid query tokens [command]

Available Commands:
  alias              Get the token alias by symbol
  aliases-by-denom   Get token aliases by denom
  all-aliases        Get all token aliases
  all-rates          Get all token rates
  rate               Get the token rate by denom
  rates-by-denom     Get token rates by denom
  token-black-whites Get token black whites

Flags:
  -h, --help   help for tokens

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors

Use "sekaid query tokens [command] --help" for more information about a command.
```

[Return to top](#sekai)

##### 14.19.1 alias

Get the token alias by symbol

Usage:
```
sekaid query tokens alias <symbol> [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for alias                                                                                   | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--home string`       | The application home directory (default `"/root/.sekaid"`)                             | ✅ ignored |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?       |

```
sekaid q tokens alias --help
Get the token alias by symbol

Usage:
  sekaid query tokens alias <symbol> [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for alias
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

<details>
  <summary>Where to find all symbols</summary>

  ```bash
  sekaid query tokens all-aliases --output=json | jq ".data[].symbol"
  "KEX"
  "SAMO"
  "TEST"
  ```
</details>

```
/# sekaid q tokens alias KEX --output=json | jq
{
  "symbol": "KEX",
  "name": "KIRA",
  "icon": "http://kira-network.s3-eu-west-1.amazonaws.com/assets/img/tokens/kex.svg",
  "decimals": 6,
  "denoms": [
    "ukex"
  ],
  "invalidated": false
}
```

[Return to top](#sekai)

##### 14.19.2 aliases-by-denom

Get token aliases by denom.

Usage:
```
sekaid query tokens aliases-by-denom [aliases] [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for alias                                                                                   | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--home string`       | The application home directory (default `"/root/.sekaid"`)                             | ✅ ignored |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid q tokens aliases-by-denom --help
Get token aliases by denom

Usage:
  sekaid query tokens aliases-by-denom [aliases] [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for aliases-by-denom
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid q tokens aliases-by-denom ukex -o json | jq
{
  "data": {
    "ukex": {
      "symbol": "KEX",
      "name": "KIRA",
      "icon": "http://kira-network.s3-eu-west-1.amazonaws.com/assets/img/tokens/kex.svg",
      "decimals": 6,
      "denoms": [
        "ukex"
      ],
      "invalidated": false
    }
  }
}
```

[Return to top](#sekai)

##### 14.19.3 all-aliases

Get all token aliases.

Usage:
```
sekaid query tokens all-aliases [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for alias                                                                                   | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--home string`       | The application home directory (default `"/root/.sekaid"`)                             | ✅ ignored |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid q tokens all-aliases --help
Get all token aliases

Usage:
  sekaid query tokens all-aliases [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for all-aliases
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid q tokens all-aliases --output=json | jq
{
  "data": [
    {
      "symbol": "KEX",
      "name": "KIRA",
      "icon": "http://kira-network.s3-eu-west-1.amazonaws.com/assets/img/tokens/kex.svg",
      "decimals": 6,
      "denoms": [
        "ukex"
      ],
      "invalidated": false
    },
    {
      "symbol": "SAMO",
      "name": "Samolean TestCoin",
      "icon": "http://kira-network.s3-eu-west-1.amazonaws.com/assets/img/tokens/samolean.svg",
      "decimals": 18,
      "denoms": [
        "samolean"
      ],
      "invalidated": false
    },
    {
      "symbol": "TEST",
      "name": "Test TestCoin",
      "icon": "http://kira-network.s3-eu-west-1.amazonaws.com/assets/img/tokens/test.svg",
      "decimals": 8,
      "denoms": [
        "test"
      ],
      "invalidated": false
    }
  ]
}
```

[Return to top](#sekai)

##### 14.19.4 all-rates

Get all token rates.

Usage:
```
sekaid query tokens all-rates [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for alias                                                                                   | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--home string`       | The application home directory (default `"/root/.sekaid"`)                             | ✅ ignored |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid q tokens all-rates --help
Get all token rates

Usage:
  sekaid query tokens all-rates [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for all-rates
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
/# sekaid q tokens all-rates --output=json | jq
{
  "data": [
    {
      "denom": "frozen",
      "fee_rate": "0.100000000000000000",
      "fee_payments": true,
      "stake_cap": "0.000000000000000000",
      "stake_min": "1",
      "stake_token": false,
      "invalidated": false
    },
    {
      "denom": "ubtc",
      "fee_rate": "10.000000000000000000",
      "fee_payments": true,
      "stake_cap": "0.250000000000000000",
      "stake_min": "1",
      "stake_token": true,
      "invalidated": false
    },
    {
      "denom": "ukex",
      "fee_rate": "1.000000000000000000",
      "fee_payments": true,
      "stake_cap": "0.500000000000000000",
      "stake_min": "1",
      "stake_token": true,
      "invalidated": false
    },
    {
      "denom": "xeth",
      "fee_rate": "0.100000000000000000",
      "fee_payments": true,
      "stake_cap": "0.100000000000000000",
      "stake_min": "1",
      "stake_token": false,
      "invalidated": false
    }
  ]
}
```

[Return to top](#sekai)

##### 14.19.5 rate

Get the token rate by denom.

Usage:
```
sekaid query tokens rate <denom> [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for alias                                                                                   | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--home string`       | The application home directory (default `"/root/.sekaid"`)                             | ✅ ignored |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid q tokens rate --help
Get the token rate by denom

Usage:
  sekaid query tokens rate <denom> [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for rate
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
/# sekaid q tokens rate ukex --output=json | jq
{
  "denom": "ukex",
  "fee_rate": "1.000000000000000000",
  "fee_payments": true,
  "stake_cap": "0.500000000000000000",
  "stake_min": "1",
  "stake_token": true,
  "invalidated": false
}
```

[Return to top](#sekai)

##### 14.19.6 rates-by-denom

Get token rates by denom.

Usage:
```
sekaid query tokens rates-by-denom [flags] ❌ mistake! here we need the one positional argument
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for alias                                                                                   | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--home string`       | The application home directory (default `"/root/.sekaid"`)                             | ✅ ignored |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid q tokens rates-by-denom --help
Get token rates by denom

Usage:
  sekaid query tokens rates-by-denom [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for rates-by-denom
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid q tokens rates-by-denom ukex -o json | jq
{
  "data": {
    "ukex": {
      "denom": "ukex",
      "fee_rate": "1.000000000000000000",
      "fee_payments": true,
      "stake_cap": "0.500000000000000000",
      "stake_min": "1",
      "stake_token": true,
      "invalidated": false
    }
  }
}
```

[Return to top](#sekai)

##### 14.19.7 token-black-whites

Get token black whites.

Usage:
```
sekaid query tokens token-black-whites [flags]
```

| Flags                 | Description                                                                                      | Work  |
| --------------------- | ------------------------------------------------------------------------------------------------ | ----- |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)            | ✅ yes |
| `-h, --help`          | help for alias                                                                                   | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`) | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                  | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--home string`       | The application home directory (default `"/root/.sekaid"`)                             | ✅ ignored |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid q tokens token-black-whites --help
Get token black whites

Usage:
  sekaid query tokens token-black-whites [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for token-black-whites
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
/# sekaid q tokens token-black-whites -o json | jq
{
  "data": {
    "whitelisted": [
      "ukex"
    ],
    "blacklisted": [
      "frozen"
    ]
  }
}
```

[Return to top](#sekai)

#### 14.20 tx

🟨 Query for a transaction by hash, "\<addr\>/\<seq\>" combination or comma-separated signatures in a committed block 🟨

Usage:
```
sekaid query tx --type=[hash|acc_seq|signature] [hash|acc_seq|signature] [flags]
```

| Flags                 | Description                                                                                                 | Work  |
| --------------------- | ----------------------------------------------------------------------------------------------------------- | ----- |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)                       | ✅ yes |
| `-h, --help`          | help for tx                                                                                                 | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`)            | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                             | ✅ yes |
| `--type string`       | The type to be used when querying tx, can be one of `"hash"`, `"acc_seq"`, `"signature"` (default `"hash"`) | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--home string`       | The application home directory (default `"/root/.sekaid"`)                             | ✅ ignored |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid q tx --help
Example:
$ sekaid query tx <hash>
$ sekaid query tx --type=acc_seq <addr>/<sequence>
$ sekaid query tx --type=signature <sig1_base64>,<sig2_base64...>

Usage:
  sekaid query tx --type=[hash|acc_seq|signature] [hash|acc_seq|signature] [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for tx
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")
      --type string     The type to be used when querying tx, can be one of "hash", "acc_seq", "signature" (default "hash")

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

By **hash**:
```
sekaid q tx 0C192B8B81789853A107AD6B6283C0AB1784A2B2C32E1F9D6730341F41BDE413 -o json | jq
{
  "height": "110257",
  "txhash": "0C192B8B81789853A107AD6B6283C0AB1784A2B2C32E1F9D6730341F41BDE413",
  "codespace": "",
  "code": 0,
  "data": "0A210A1B2F6B6972612E676F762E4D73675375626D697450726F706F73616C1202080C",
  "raw_log": ". . .",
  "logs": [
    // . . .
  ]
}
```

By **sequence**:
```
sekaid query tx --type=acc_seq kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4/92 -o json | jq
{
  "height": "109602",
  "txhash": "9659E50E56F715EE67C935B2F440A7F02BBA131DDD7697E7510A5F866DAC8BA5",
  "codespace": "",
  "code": 0,
  "data": "0A310A2F2F6B6972612E676F762E4D736743616E63656C4964656E746974795265636F72647356657269667952657175657374",
  "raw_log": ". . .",
  "logs": [
    // . . .
  ]
}
```

By **signature**:
```
sekaid query tx --type=signature G5+araW7aplPdSJiPAvrgkmVMdmE14SNi6QmKgS0t0cEKLDeJ/BotFN2ktlIYl7nDWZ7zNEFYgsLtRKTT1jtfw== -o json | jq
{
  "height": "109602",
  "txhash": "9659E50E56F715EE67C935B2F440A7F02BBA131DDD7697E7510A5F866DAC8BA5",
  "codespace": "",
  "code": 0,
  "data": "0A310A2F2F6B6972612E676F762E4D736743616E63656C4964656E746974795265636F72647356657269667952657175657374",
  "raw_log": ". . .",
  "logs": [
    // . . .
  ]
}
```

[Return to top](#sekai)

#### 14.21 txs

Search for transactions that match the exact given events where results are paginated.

Usage:
```
sekaid query txs [flags]
```

| Flags                 | Description                                                                                                 | Work  |
| --------------------- | ----------------------------------------------------------------------------------------------------------- | ----- |
| `--events string`     | list of transaction events in the form of `{eventType}.{eventAttribute}={value}`                            | ✅ yes |
| `--height int`        | Use a specific height to query state at (this can error if the node is pruning state)                       | ✅ yes |
| `-h, --help`          | help for txs                                                                                                | ✅ yes |
| `--limit int`         | Query number of transactions results per page returned (default `30`)                                       | ✅ yes |
| `--node string`       | \<host\>:\<port\> to Tendermint RPC interface for this chain (default `"tcp://localhost:26657"`)            | ✅ yes |
| `-o, --output string` | Output format (`text\|json`) (default `"text"`)                                                             | ✅ yes |
| `--type string`       | The type to be used when querying tx, can be one of `"hash"`, `"acc_seq"`, `"signature"` (default `"hash"`) | ✅ yes |
| `--page int`          | Query a specific page of paginated results (default `1`)                                                    | ✅ yes |



| Global Flags          | Description                                                                            | Work      |
| --------------------- | -------------------------------------------------------------------------------------- | --------- |
| `--home string`       | The application home directory (default `"/root/.sekaid"`)                             | ✅ ignored |
| `--chain-id string`   | The network chain ID                                                                   | ✅ ignored |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?       |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?       |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?       |

```
/# sekaid q txs --help
Search for transactions that match the exact given events where results are paginated.
Each event takes the form of '{eventType}.{eventAttribute}={value}'. Please refer
to each module's documentation for the full set of events to query for. Each module
documents its respective events under 'xx_events.md'.

Example:
$ sekaid query txs --events 'message.sender=cosmos1...&message.action=withdraw_delegator_reward' --page 1 --limit 30

Usage:
  sekaid query txs [flags]

Flags:
      --events string   list of transaction events in the form of {eventType}.{eventAttribute}={value}
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for txs
      --limit int       Query number of transactions results per page returned (default 30)
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")
      --page int        Query a specific page of paginated results (default 1)

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

> `--events` - is requeired flag!

```
sekaid q txs --events=submit_proposal.proposal_type=UpsertDataRegistry --output=json | jq
{
  "total_count": "4",
  "count": "4",
  "page_number": "1",
  "page_total": "1",
  "limit": "30",
  "txs": [
    {. . .},
    {. . .},
    {. . .},
    {. . .}
  ]
}
```

Other usages:
```
sekaid q txs --events=submit_proposal.proposal_type=UpsertDataRegistry --limit=1 --page=2 --output=json | jq
{
  "total_count": "4",
  "count": "1",
  "page_number": "2",
  "page_total": "4",
  "limit": "1",
  "txs": [
    {. . .}
  ]
}
```

[Return to top](#sekai)

#### 14.22 ubi

[Return to top](#sekai)

#### 14.23 upgrade

[Return to top](#sekai)

### 15. rollback

A state rollback is performed to recover from an incorrect application state transition,
when Tendermint has persisted an incorrect app hash and is thus unable to make
progress. Rollback overwrites a state at height n with the state at height n - 1.
The application also roll back to height n - 1. No blocks are removed, so upon
restarting Tendermint the transactions in block n will be re-executed against the
application.

Usage:
```
sekaid rollback [flags]
```

| Flags        | Description  | Work  |
| ------------ | ------------ | ----- |
| `-h, --help` | help for txs | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--home string`       | The application home directory (default `"/root/.sekaid"`)                             | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?   |


```
# sekaid rollback -h

A state rollback is performed to recover from an incorrect application state transition,
when Tendermint has persisted an incorrect app hash and is thus unable to make
progress. Rollback overwrites a state at height n with the state at height n - 1.
The application also roll back to height n - 1. No blocks are removed, so upon
restarting Tendermint the transactions in block n will be re-executed against the
application.

Usage:
  sekaid rollback [flags]

Flags:
  -h, --help   help for rollback

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

**important, sekaid node has to be stoped to perform this command**

```
# sekaid rollback --home $SEKAID_HOME
Rolled back state to height 31 and hash 89288F2D6B9745B24DE223CA83CC2176CE5EE9B34BD3596D8CCF1F68F76F9E4Broot@669b80378d29
```

[Return to top](#sekai)

### 16. rosetta

[Return to top](#sekai)

### 17. start
Run the full node application with Tendermint in or out of process. By
default, the application will run with Tendermint in process.

Pruning options can be provided via the '--pruning' flag or alternatively with '--pruning-keep-recent', 'pruning-keep-every', and 'pruning-interval' together.

For '--pruning' the options are as follows:

default: the last 100 states are kept in addition to every 500th state; pruning at 10 block intervals
nothing: all historic states will be saved, nothing will be deleted (i.e. archiving node)
everything: all saved states will be deleted, storing only the current and previous state; pruning at 10 block intervals
custom: allow pruning options to be manually specified through 'pruning-keep-recent', 'pruning-keep-every', and 'pruning-interval'

Node halting configurations exist in the form of two flags: '--halt-height' and '--halt-time'. During the ABCI Commit phase, the node will check if the current block height is greater than or equal to
the halt-height or if the current block time is greater than or equal to the halt-time. If so, the node will attempt to gracefully shutdown and the block will not be committed. In addition, the node
will not be able to commit subsequent blocks.

For profiling and benchmarking purposes, CPU profiling can be enabled via the '--cpu-profile' flag which accepts a path for the resulting pprof file.

The node may be started in a 'query only' mode where only the gRPC and JSON HTTP API services are enabled via the 'grpc-only' flag. In this mode, Tendermint is bypassed and can be used when legacy queries are needed after an on-chain upgrade is performed. Note, when enabled, gRPC will also be automatically enabled.

Usage:
```
sekaid start [flags]
```

| Flags                                           | Description                                                                                                                                    | Work |
| ----------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---- |
| --abci string                                   | genesis file chain-id, if left blank will be randomly created                                                                                  | ❌?   |
| --address string                                | Listen address (default "tcp://0.0.0.0:26658")                                                                                                 | ❌?   |
| --consensus.create_empty_blocks                 | set this to false to only produce blocks when there are txs or when the AppHash changes (default true)                                         | ❌?   |
| --consensus.create_empty_blocks_interval string | the possible interval between empty blocks (default "0s")                                                                                      | ❌?   |
| --consensus.double_sign_check_height int        | how many blocks to look back to check existence of the node's consensus votes before joining consensus                                         | ❌?   |
| --cpu-profile string                            | Enable CPU profiling and write to the provided file                                                                                            | ❌?   |
| --db_backend string                             | database backend: goleveldb \| cleveldb \| boltdb \| rocksdb \| badgerdb (default "goleveldb")                                                 | ❌?   |
| --db_dir string                                 | database directory (default "data")                                                                                                            | ❌?   |
| --fast_sync                                     | fast blockchain syncing (default true)                                                                                                         | ❌?   |
| --genesis_hash bytesHex                         | optional SHA-256 hash of the genesis file                                                                                                      | ❌?   |
| --grpc-only                                     | Start the node in gRPC query only mode (no Tendermint process is started)                                                                      | ❌?   |
| --grpc-web.address string                       | The gRPC-Web server address to listen on (default "0.0.0.0:9091")                                                                              | ❌?   |
| --grpc-web.enable                               | Define if the gRPC-Web server should be enabled. (Note: gRPC must also be enabled.) (default true)                                             | ❌?   |
| --grpc.address string                           | the gRPC server address to listen on (default "0.0.0.0:9090")                                                                                  | ❌?   |
| --grpc.enable                                   | Define if the gRPC server should be enabled (default true)                                                                                     | ❌?   |
| --halt-height uint                              | Block height at which to gracefully halt the chain and shutdown the node                                                                       | ❌?   |
| --halt-time uint                                | Minimum block time (in Unix seconds) at which to gracefully halt the chain and shutdown the node                                               | ❌?   |
| -h, --help                                      | help for start                                                                                                                                 | ❌?   |
| --iavl-disable-fastnode                         | Enable fast node for IAVL tree (default true)                                                                                                  | ❌?   |
| --inter-block-cache                             | Enable inter-block caching (default true)                                                                                                      | ❌?   |
| --inv-check-period uint                         | Assert registered invariants every N blocks                                                                                                    | ❌?   |
| --min-retain-blocks uint                        | Minimum block height offset during ABCI commit to prune Tendermint blocks                                                                      | ❌?   |
| --minimum-gas-prices string                     | Minimum gas prices to accept for transactions; Any fee in a tx must meet this minimum (e.g. 0.01photino;0.0001stake)                           | ❌?   |
| --moniker string                                | node name (default "validator.local")                                                                                                          | ❌?   |
| --p2p.external-address string                   | ip:port address to advertise to peers for them to dial                                                                                         | ❌?   |
| --p2p.laddr string                              | node listen address. (0.0.0.0:0 means any interface, any port) (default "tcp://0.0.0.0:26656")                                                 | ❌?   |
| --p2p.persistent_peers string                   | comma-delimited ID@host:port persistent peers                                                                                                  | ❌?   |
| --p2p.pex                                       | enable/disable Peer-Exchange (default true)                                                                                                    | ❌?   |
| --p2p.private_peer_ids string                   | comma-delimited private peer IDs                                                                                                               | ❌?   |
| --p2p.seed_mode                                 | enable/disable seed mode                                                                                                                       | ❌?   |
| --p2p.seeds string                              | comma-delimited ID@host:port seed nodes                                                                                                        | ❌?   |
| --p2p.unconditional_peer_ids string             | comma-delimited IDs of unconditional peers                                                                                                     | ❌?   |
| --p2p.upnp                                      | enable/disable UPNP port forwarding                                                                                                            | ❌?   |
| --priv_validator_laddr string                   | socket address to listen on for connections from external priv_validator process                                                               | ❌?   |
| --proxy_app string                              | proxy app address, or one of: 'kvstore', 'persistent_kvstore', 'counter', 'e2e' or 'noop' for local testing. (default "tcp://127.0.0.1:26658") | ❌?   |
| --pruning string                                | Pruning strategy (default\|nothing\|everything\|custom) (default "default")                                                                    | ❌?   |
| --pruning-interval uint                         | Height interval at which pruned heights are removed from disk (ignored if pruning is not 'custom')                                             | ❌?   |
| --pruning-keep-every uint                       | Offset heights to keep on disk after 'keep-every' (ignored if pruning is not 'custom')                                                         | ❌?   |
| --pruning-keep-recent uint                      | Number of recent heights to keep on disk (ignored if pruning is not 'custom')                                                                  | ❌?   |
| --rpc.grpc_laddr string                         | GRPC listen address (BroadcastTx only). Port required                                                                                          | ❌?   |
| --rpc.laddr string                              | RPC listen address. Port required (default "tcp://127.0.0.1:26657")                                                                            | ✅yes |
| --rpc.pprof_laddr string                        | pprof listen address (https://golang.org/pkg/net/http/pprof)                                                                                   | ❌?   |
| --rpc.unsafe                                    | enabled unsafe rpc methods                                                                                                                     | ❌?   |
| --state-sync.snapshot-interval uint             | State sync snapshot interval                                                                                                                   | ❌?   |
| --state-sync.snapshot-keep-recent uint32        | State sync snapshot to keep (default 2)                                                                                                        | ❌?   |
| --trace-store string                            | Enable KVStore tracing to an output file                                                                                                       | ❌?   |
| --transport string                              | Transport protocol: socket, grpc (default "socket")                                                                                            | ❌?   |
| --unsafe-skip-upgrades ints                     | Skip a set of upgrade heights to continue the old binary                                                                                       | ❌?   |
| --with-tendermint                               | Run abci app embedded in-process with tendermint (default true)                                                                                | ❌?   |
| --x-crisis-skip-assert-invariants               | Skip x/crisis invariants check on startup                                                                                                      | ❌?   |



| Global flags         | Description                                                                            | Work |
| -------------------- | -------------------------------------------------------------------------------------- | ---- |
| `--home`             | directory for config and data (default `"/root/.sekaid"`)                              | ❌?no |
| `--log_format`       | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌?no |
| `--log_level string` | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌?no |
| `--trace`            | Print out full stack trace on errors                                                   | ❌?no |

```
sekaid start --rpc.laddr "tcp://0.0.0.0:26657" 
```
[Return to top](#sekai)

### 18. status
Query remote node for status

#TODO: the descritpion is not accurate. `sekaid status` queries the node where it was launched by default.

🟨  
🟨  
🟨  

Usage:
```
sekaid status [flags]
```

| Flags    | Description                                            | Work  |
| -------- | ------------------------------------------------------ | ----- |
| `--help` | Help for status                                        | ✅ yes |
| `--node` | Node to connect to (default `"tcp://localhost:26657"`) | ✅ yes |



| Global flags         | Description                                                                            | Work  |
| -------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--home`             | Directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format`       | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string` | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`            | Print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid status --help
Query remote node for status

Usage:
  sekaid status [flags]

Flags:
  -h, --help          help for status
  -n, --node string   Node to connect to (default "tcp://localhost:26657")

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
# sekaid status -n=tcp://10.1.0.2:26657 | jq
{
  "NodeInfo": {
    "protocol_version": {
      "p2p": "8",
      "block": "11",
      "app": "0"
    },
    "id": "213267558cc0febde6f1b8149f92fecf16d48b38",
    "listen_addr": "tcp://168.119.228.165:36656",
    "network": "localnet-4",
    "version": "0.34.22",
    "channels": "40202122233038606100",
    "moniker": "KIRA VALIDATOR NODE",
    "other": {
      "tx_index": "on",
      "rpc_address": "tcp://0.0.0.0:26657"
    }
  },
  "SyncInfo": {
    "latest_block_hash": "EF98E1777F8D9477AED36B0BB45238445D3CED2C5E3AA7EF0EAAB984BEED382D",
    "latest_app_hash": "37B64A94D872F0DD081EACE2ECCAD5B397121C7DD6F9E6D3B419299E94D74BD6",
    "latest_block_height": "74627",
    "latest_block_time": "2023-06-04T10:14:45.658686623Z",
    "earliest_block_hash": "26AD6395E9AB39AE545D53F64B2DFF4C664F55448A5A96F9963336386CFE2B3C",
    "earliest_app_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
    "earliest_block_height": "1",
    "earliest_block_time": "2023-05-25T12:56:59.153143882Z",
    "catching_up": false
  },
  "ValidatorInfo": {
    "Address": "AFC8EBD65CE1E7DD38E1E4DD514E9B03A0085E98",
    "PubKey": {
      "type": "tendermint/PubKeyEd25519",
      "value": "4q2MwShsxSt2UBXmEv/rLXevigs6uln9J74uxA/XbRE="
    },
    "VotingPower": "1"
  }
}
```

[Return to top](#sekai)

### 19. tendermint

Tendermint subcommands

Usage:
```
sekaid tendermint [command]
```


Available Commands:

| Subcommand       | Description                                                                        |
| ---------------- | ---------------------------------------------------------------------------------- |
| `reset-state `     | Remove all the data and WAL                                                        |
| `show-address`     | Shows this node's tendermint validator consensus address                           |
| `show-node-id`     | Show this node's ID                                                                |
| `show-validator`   | Show this node's tendermint validator info                                         |
| `unsafe-reset-all` | (unsafe) Remove all the data and WAL, reset this node's validator to genesis state |
| `version`          | Print tendermint libraries' version                                                |



| Flags        | Description | Work  |
| ------------ | ----------- | ----- |
| `-h, --help` | Help for tx | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--home string`       | Directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?   |

```
# sekaid tendermint -h
Tendermint subcommands

Usage:
  sekaid tendermint [command]

Available Commands:
  reset-state      Remove all the data and WAL
  show-address     Shows this node's tendermint validator consensus address
  show-node-id     Show this node's ID
  show-validator   Show this node's tendermint validator info
  unsafe-reset-all (unsafe) Remove all the data and WAL, reset this node's validator to genesis state
  version          Print tendermint libraries' version

Flags:
  -h, --help   help for tendermint

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```


```
# sekaid tendermint reset-state --home $SEKAID_HOME
I[2023-06-13|13:47:39.222] Removed all blockstore.db                    dir=/root/.sekaid-PEPEGENETWORK-1/data/blockstore.db
I[2023-06-13|13:47:39.222] Removed all state.db                         dir=/root/.sekaid-PEPEGENETWORK-1/data/state.db
I[2023-06-13|13:47:39.223] Removed all cs.wal                           dir=/root/.sekaid-PEPEGENETWORK-1/data/cs.wal
I[2023-06-13|13:47:39.223] Removed all evidence.db                      dir=/root/.sekaid-PEPEGENETWORK-1/data/evidence.db
I[2023-06-13|13:47:39.224] Removed tx_index.db                          dir=/root/.sekaid-PEPEGENETWORK-1/data/tx_index.db
```


```
sekaid tendermint show-address --home $SEKAID_HOME
kiravalcons15jeawxkxx9nws6knjkywggmsdn5qa7dhnxg8k
```



```
# sekaid tendermint show-node-id --home $SEKAID_HOME
cc57aecd262219356156532438afea19f83d7960
```


```
# sekaid tendermint show-validator --home $SEKAID_HOME
{"@type":"/cosmos.crypto.ed25519.PubKey","key":"njRqCAsanZhtcntb+qvP/pKTYy77s0CKY5FY4TUk/rI="}
```


```
# sekaid tendermint unsafe-reset-all --home $SEKAID_HOME
I[2023-06-13|13:52:30.009] Removed existing address book                file=/root/.sekaid-PEPEGENETWORK-1/config/addrbook.json
I[2023-06-13|13:52:30.010] Removed all blockchain history               dir=/root/.sekaid-PEPEGENETWORK-1/data
I[2023-06-13|13:52:30.014] Reset private validator file to genesis state keyFile=/root/.sekaid-PEPEGENETWORK-1/config/priv_validator_key.json stateFile=/root/.sekaid-PEPEGENETWORK-1/data/priv_validator_state.json
```


```
# sekaid tendermint version
tendermint: 0.34.22
abci: 0.17.0
blockprotocol: 11
p2pprotocol: 8

```

[Return to top](#sekai)

### 20. testnet

[Return to top](#sekai)

### 21. tx

Transactions subcommands

Usage:
```
sekaid tx [flags]
sekaid tx [command]
```

Available Commands:

| Subcommand                                         | Description                                                     |
| -------------------------------------------------- | --------------------------------------------------------------- |
| [`bank`](#211-bank)                                | Bank transaction subcommands                                    |
| [`basket`](#212-basket)                            | Tokens sub commands                                             |
| [`broadcast`](#213-broadcast)                      | Broadcast transactions generated offline                        |
| [`collectives`](#214-collectives)                  | Collectives sub commands                                        |
| [`custody`](#215-custody)                          | Custody sub commands                                            |
| [`customevidence`](#216-customevidence)            | Evidence transaction subcommands                                |
| [`customgov`](#217-customgov)                      | Custom gov sub commands                                         |
| [`customslashing`](#218-customslashing)            | Slashing transaction subcommands                                |
| [`customstaking`](#219-customstaking)              | Staking module subcommands                                      |
| [`decode`](#2110-decode)                           | Decode a binary encoded transaction string                      |
| [`distributor`](#2111-distributor)                 | Distributor sub commands                                        |
| [`encode`](#2112-encode)                           | Encode transactions generated offline                           |
| [`layer2`](#2113-layer2)                           | Tokens sub commands                                             |
| [`multisign`](#2114-multisign)                     | Generate multisig signatures for transactions generated offline |
| [`multistaking`](#2115-multistaking)               | Multistaking sub commands                                       |
| [`recovery`](#2116-recovery)                       | Recovery transaction subcommands                                |
| [`sign`](#2117-sign)                               | Sign a transaction generated offline                            |
| [`sign-batch`](#2118-sign-batch)                   | Sign transaction batch files                                    |
| [`spending`](#2119-spending)                       | Tokens sub commands                                             |
| [`tokens`](#2120-tokens)                           | Tokens sub commands                                             |
| [`ubi`](#2121-ubi)                                 | UBI sub commands                                                |
| [`upgrade`](#2122-upgrade)                         | Upgrade transaction subcommands                                 |
| [`validate-signatures`](#2123-validate-signatures) | Validate transactions signatures                                |



| Flags               | Description          | Work  |
| ------------------- | -------------------- | ----- |
| `--chain-id string` | The network chain ID | ✅ yes |
| `-h, --help`        | Help for tx          | ✅ yes |



| Global Flags          | Description                                                                            | Work |
| --------------------- | -------------------------------------------------------------------------------------- | ---- |
| `--home string`       | Directory for config and data (default `"/root/.sekaid"`)                              | ❌ ?  |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?  |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?  |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?  |

```
/# sekaid tx --help
Transactions subcommands

Usage:
  sekaid tx [flags]
  sekaid tx [command]

Available Commands:


  bank                Bank transaction subcommands
  basket              Tokens sub commands
  broadcast           Broadcast transactions generated offline
  collectives         Collectives sub commands
  custody             custody sub commands
  customevidence      Evidence transaction subcommands
  customgov           Custom gov sub commands
  customslashing      Slashing transaction subcommands
  customstaking       staking module subcommands
  decode              Decode a binary encoded transaction string
  distributor         distributor sub commands
  encode              Encode transactions generated offline
  layer2              Tokens sub commands
  multisign           Generate multisig signatures for transactions generated offline
  multistaking        multistaking sub commands
  recovery            Recovery transaction subcommands
  sign                Sign a transaction generated offline
  sign-batch          Sign transaction batch files
  spending            Tokens sub commands
  tokens              Tokens sub commands
  ubi                 UBI sub commands
  upgrade             Upgrade transaction subcommands
  validate-signatures validate transactions signatures

Flags:
      --chain-id string   The network chain ID
  -h, --help              help for tx

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors

Use "sekaid tx [command] --help" for more information about a command.
```

[Return to top](#sekai)

#### 21.1 bank

Bank transaction subcommands

Usage:
```
sekaid tx bank [flags]
sekaid tx bank [command]
```

Available Commands:
| Subcommand           | Description                            |
| -------------------- | -------------------------------------- |
| [`send`](#2111-send) | Send funds from one account to another |



| Flags        | Description   | Work  |
| ------------ | ------------- | ----- |
| `-h, --help` | Help for bank | ✅ yes |

```
/# sekaid tx bank --help
Bank transaction subcommands

Usage:
  sekaid tx bank [flags]
  sekaid tx bank [command]

Available Commands:
  send        Send funds from one account to another. Note, the'--from' flag is
ignored as it is implied from [from_key_or_address].

Flags:
  -h, --help   help for bank

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors

Use "sekaid tx bank [command] --help" for more information about a command.
```

[Return to top](#sekai)

##### 21.1.1 send

Send funds from one account to another

Usage:
```
sekaid tx bank send [from_key_or_address] [to_address] [amount] [flags]
```

| Flags                         | Description                                                                                                    | Work      |
| ----------------------------- | -------------------------------------------------------------------------------------------------------------- | --------- |
| `-a, --account-number uint`   | The account number of the signing account (**offline mode only**)                                              | ❌ ?       |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                        | ✅ yes     |
| `--dry-run`                   | Ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                      | ❌ ?       |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                 | ❌ ?       |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                              | ✅ yes     |
| `--from string`               | Name or address of private key with which to sign                                                              | ✅ ignored |
| `--gas string`                | Gas limit to set per-transaction; set to `"auto"` to calculate sufficient gas automatically (default `200000`) | ❌ ?       |
| `--gas-adjustment float`      | Adjustment factor to be multiplied against the estimate returned by the tx simulation ... (default `1`)        | ❌ ?       |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                | ❌ ?       |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)       | ✅ yes     |
| `-h, --help`                  | Help for send                                                                                                  | ✅ yes     |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                            | ✅ yes     |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default `'home'` directory will be used                          | ✅ yes     |
| `--ledger`                    | Use a connected Ledger device                                                                                  | ❌ ?       |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)               | ✅ yes     |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                             | ❌ ?       |
| `--offline`                   | Offline mode (does not allow any online functionality)                                                         | ❌ ?       |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                | ✅ yes     |
| `-s, --sequence uint`         | The sequence number of the signing account (`offline` mode only)                                               | ❌ ?       |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                           | ✅ yes     |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                        | ❌ ?       |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                       | ✅ yes     |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | Directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx bank send --help
Send funds from one account to another. Note, the'--from' flag is
ignored as it is implied from [from_key_or_address].

Usage:
  sekaid tx bank send [from_key_or_address] [to_address] [amount] [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for send
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```


<details>
  <summary>Check before</summary>
  
  validator1:
  ```bash
  sekaid keys show validator -a --home=/root/.sekai --keyring-backend=test
  ```
  `kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4`

  ```bash
  sekaid query bank balances kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --output=json | jq
  ```
  ```json
  {
    "balances": [
      {
        "denom": "lol",
        "amount": "900000"
      },
      {
        "denom": "samolean",
        "amount": "1999999999999999999999900100"
      },
      {
        "denom": "test",
        "amount": "29999779999900000"
      },
      {
        "denom": "ukex",
        "amount": "299998797027495"
      }
    ],
    "pagination": {
      "next_key": null,
      "total": "0"
    }
  }
  ```

  validator2:
  ```bash
  sekaid keys show validator -a --home=/root/.sekai --keyring-backend=test
  ```
  `kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt`

  ```bash
  sekaid query bank balances kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt --output=json | jq
  ```
  ```json
  {
    "balances": [
      {
        "denom": "ukex",
        "amount": "953621"
      }
    ],
    "pagination": {
      "next_key": null,
      "total": "0"
    }
  }
  ```
</details>

```
/# sekaid tx bank send kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt 100lol --keyring-backend=test --home=/root/.sekai --chain-id=localnet-4 --fees 99ukex --output=json --yes | jq
{
  "height": "0",
  "txhash": "CCD7FE3C3E1FB517BB0E0BB81BAE550032BF4A52416F82ABCE66DF470D316C6A",
  "codespace": "",
  "code": 0,
  "data": "",
  "raw_log": "[]",
  "logs": [],
  "info": "",
  "gas_wanted": "0",
  "gas_used": "0",
  "tx": null,
  "timestamp": "",
  "events": []
}
```

<details>
  <summary>Check tx execution</summary>

  ```bash
  sekaid query tx CCD7FE3C3E1FB517BB0E0BB81BAE550032BF4A52416F82ABCE66DF470D316C6A --output=json | jq
  ```

  ```json
  {
    "height": "83663",
    "txhash": "CCD7FE3C3E1FB517BB0E0BB81BAE550032BF4A52416F82ABCE66DF470D316C6A",
    "codespace": "",
    "code": 0,
    "data": "0A1E0A1C2F636F736D6F732E62616E6B2E763162657461312E4D736753656E64",
    "raw_log": "[{\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt\"},{\"key\":\"amount\",\"value\":\"100lol\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\"},{\"key\":\"amount\",\"value\":\"100lol\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.bank.v1beta1.MsgSend\"},{\"key\":\"sender\",\"value\":\"kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\"},{\"key\":\"module\",\"value\":\"bank\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt\"},{\"key\":\"sender\",\"value\":\"kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\"},{\"key\":\"amount\",\"value\":\"100lol\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "coin_received",
            "attributes": [
              {
                "key": "receiver",
                "value": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt"
              },
              {
                "key": "amount",
                "value": "100lol"
              }
            ]
          },
          {
            "type": "coin_spent",
            "attributes": [
              {
                "key": "spender",
                "value": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4"
              },
              {
                "key": "amount",
                "value": "100lol"
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/cosmos.bank.v1beta1.MsgSend"
              },
              {
                "key": "sender",
                "value": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4"
              },
              {
                "key": "module",
                "value": "bank"
              }
            ]
          },
          {
            "type": "transfer",
            "attributes": [
              {
                "key": "recipient",
                "value": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt"
              },
              {
                "key": "sender",
                "value": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4"
              },
              {
                "key": "amount",
                "value": "100lol"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "0",
    "gas_used": "0",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/cosmos.bank.v1beta1.MsgSend",
            "from_address": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
            "to_address": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
            "amount": [
              {
                "denom": "lol",
                "amount": "100"
              }
            ]
          }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
      },
      "auth_info": {
        "signer_infos": [
          {
            "public_key": {
              "@type": "/cosmos.crypto.secp256k1.PubKey",
              "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "83"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "ukex",
              "amount": "99"
            }
          ],
          "gas_limit": "200000",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
        "WK8mHiZNoCwBE0joGrAq/nksi2/xOFTo7Hube03HqapBlm1fbItmR5kE0uzBjSGbm3GLFNJt1zS0O3QyJx1zpQ=="
      ]
    },
    "timestamp": "2023-06-05T12:07:57Z",
    "events": [
      {
        "type": "tx",
        "attributes": [
          {
            "key": "YWNjX3NlcQ==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NC84Mw==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "V0s4bUhpWk5vQ3dCRTBqb0dyQXEvbmtzaTIveE9GVG83SHViZTAzSHFhcEJsbTFmYkl0bVI1a0UwdXpCalNHYm0zR0xGTkp0MXpTME8zUXlKeDF6cFE9PQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "OTl1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "OTl1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "OTl1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "ZmVl",
            "value": "OTl1a2V4",
            "index": true
          },
          {
            "key": "ZmVlX3BheWVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "L2Nvc21vcy5iYW5rLnYxYmV0YTEuTXNnU2VuZA==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwbG9s",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTE3YWVxeHZrbDNnMzdwZmNnd2txeHZrcm42M2pmbGpmdmpyYXZudA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwbG9s",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTE3YWVxeHZrbDNnMzdwZmNnd2txeHZrcm42M2pmbGpmdmpyYXZudA==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwbG9s",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "bW9kdWxl",
            "value": "YmFuaw==",
            "index": true
          }
        ]
      }
    ]
  }
  ```
</details>

<details>
  <summary>Check balances</summary>

  ```bash
  sekaid query bank balances kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --output=json | jq
  ```

  ```json
  {
    "balances": [
      {
        "denom": "lol",
        "amount": "899900"
      },
      {
        "denom": "samolean",
        "amount": "1999999999999999999999900100"
      },
      {
        "denom": "test",
        "amount": "29999779999900000"
      },
      {
        "denom": "ukex",
        "amount": "299998797027396"
      }
    ],
    "pagination": {
      "next_key": null,
      "total": "0"
    }
  }

  /# sekaid query bank balances kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt --output=json | jq
  {
    "balances": [
      {
        "denom": "lol",
        "amount": "100"
      },
      {
        "denom": "ukex",
        "amount": "953621"
      }
    ],
    "pagination": {
      "next_key": null,
      "total": "0"
    }
  }
  ```
</details>

[Return to top](#sekai)

#### 21.2 basket

[Return to top](#sekai)

#### 21.3 broadcast

Broadcast transactions created with the --generate-only flag and signed with the sign command.

Usage:
```
sekaid tx broadcast [file_path] [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `--amino`                     | Generate Amino encoded JSON suitable for submiting to the txs REST endpoint                                                                                 | ❌ ?   |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ❌ ?   |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ❌ ?   |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ✅ yes |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for broadcast                                                                                                                                          | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default `'home'` directory will be used                                                                       | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ❌ ?   |
| `--multisig string`           | Address or key name of the multisig account on behalf of which the transaction shall be signed                                                              | ❌ ?   |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ❌ ?   |
| `--output-document string`    | The document will be written to the given file instead of STDOUT                                                                                            | ❌ ?   |
| `--overwrite`                 | Overwrite existing signatures with a new one. If disabled, new signature will be appended                                                                   | ❌ ?   |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--signature-only`            | Print only the signatures                                                                                                                                   | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ❌ ?   |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | Directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx broadcast --help
Broadcast transactions created with the --generate-only
flag and signed with the sign command. Read a transaction from [file_path] and
broadcast it to a node. If you supply a dash (-) argument in place of an input
filename, the command reads from standard input.

$ <appd> tx broadcast ./mytxn.json

Usage:
  sekaid tx broadcast [file_path] [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for broadcast
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

<details>
  <summary>Pregenerate tx</summary>

  Generate tx (Example of sending tokens via `bank` module):
  ```
  /# sekaid tx bank send kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt 100samolean --home=/root/.sekai --chain-id=localnet-4 --keyring-backend=test --fees=99ukex --generate-only --output=json | jq > tx.json
  ```

  Sign tx:
  ```
  /# sekaid tx sign tx.json --from=kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --home=/root/.sekai --chain-id=localnet-4 --keyring-backend=test --output=json | jq >
tx_signed.json
  ```
</details>

```
/# sekaid tx broadcast tx_signed.json  --home=/root/.sekai --chain-id=localnet-4 --keyring-backend=test --output=json | jq
{
  "height": "0",
  "txhash": "9AE59A7878EE853F6175B77BC83F6A132B243810308A1A17DF35641E703B29A9",
  "codespace": "",
  "code": 0,
  "data": "",
  "raw_log": "[]",
  "logs": [],
  "info": "",
  "gas_wanted": "0",
  "gas_used": "0",
  "tx": null,
  "timestamp": "",
  "events": []
}
```

<details>
  <summary>Check tx execution</summary>

  ```
  /# sekaid q tx 9AE59A7878EE853F6175B77BC83F6A132B243810308A1A17DF35641E703B29A9 -o json | jq
  ```

  ```json
  {
    "height": "91894",
    "txhash": "9AE59A7878EE853F6175B77BC83F6A132B243810308A1A17DF35641E703B29A9",
    "codespace": "",
    "code": 0,
    "data": "0A1E0A1C2F636F736D6F732E62616E6B2E763162657461312E4D736753656E64",
    "raw_log": "[{\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt\"},{\"key\":\"amount\",\"value\":\"100samolean\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\"},{\"key\":\"amount\",\"value\":\"100samolean\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.bank.v1beta1.MsgSend\"},{\"key\":\"sender\",\"value\":\"kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\"},{\"key\":\"module\",\"value\":\"bank\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt\"},{\"key\":\"sender\",\"value\":\"kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\"},{\"key\":\"amount\",\"value\":\"100samolean\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "coin_received",
            "attributes": [
              {
                "key": "receiver",
                "value": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt"
              },
              {
                "key": "amount",
                "value": "100samolean"
              }
            ]
          },
          {
            "type": "coin_spent",
            "attributes": [
              {
                "key": "spender",
                "value": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4"
              },
              {
                "key": "amount",
                "value": "100samolean"
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/cosmos.bank.v1beta1.MsgSend"
              },
              {
                "key": "sender",
                "value": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4"
              },
              {
                "key": "module",
                "value": "bank"
              }
            ]
          },
          {
            "type": "transfer",
            "attributes": [
              {
                "key": "recipient",
                "value": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt"
              },
              {
                "key": "sender",
                "value": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4"
              },
              {
                "key": "amount",
                "value": "100samolean"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "0",
    "gas_used": "0",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/cosmos.bank.v1beta1.MsgSend",
            "from_address": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
            "to_address": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
            "amount": [
              {
                "denom": "samolean",
                "amount": "100"
              }
            ]
          }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
      },
      "auth_info": {
        "signer_infos": [
          {
            "public_key": {
              "@type": "/cosmos.crypto.secp256k1.PubKey",
              "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "84"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "ukex",
              "amount": "99"
            }
          ],
          "gas_limit": "200000",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
        "OukW9MMdqzxI2LUymHKu9DXz1A6B6+RxEZBSn55vw5QSnC7DHi+ryIm2tDcwMPRm84t2uDh55DGyniNyyHoSrg=="
      ]
    },
    "timestamp": "2023-06-06T11:42:47Z",
    "events": [
      {
        "type": "tx",
        "attributes": [
          {
            "key": "YWNjX3NlcQ==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NC84NA==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "T3VrVzlNTWRxenhJMkxVeW1IS3U5RFh6MUE2QjYrUnhFWkJTbjU1dnc1UVNuQzdESGkrcnlJbTJ0RGN3TVBSbTg0dDJ1RGg1NURHeW5pTnl5SG9Tcmc9PQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "OTl1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "OTl1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "OTl1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "ZmVl",
            "value": "OTl1a2V4",
            "index": true
          },
          {
            "key": "ZmVlX3BheWVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "L2Nvc21vcy5iYW5rLnYxYmV0YTEuTXNnU2VuZA==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwc2Ftb2xlYW4=",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTE3YWVxeHZrbDNnMzdwZmNnd2txeHZrcm42M2pmbGpmdmpyYXZudA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwc2Ftb2xlYW4=",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTE3YWVxeHZrbDNnMzdwZmNnd2txeHZrcm42M2pmbGpmdmpyYXZudA==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwc2Ftb2xlYW4=",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "bW9kdWxl",
            "value": "YmFuaw==",
            "index": true
          }
        ]
      }
    ]
  }
  ```
</details>

[Return to top](#sekai)

#### 21.4 collectives

[Return to top](#sekai)

#### 21.5 custody

[Return to top](#sekai)

#### 21.6 customevidence

[Return to top](#sekai)

#### 21.7 customgov

Custom gov sub commands.

Usage:
```
sekaid tx customgov [flags]
sekaid tx customgov [command]
```

Available Commands:

| Subcommands                                                                              | Description                                                                |
| ---------------------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| [`cancel-identity-records-verify-request`](#2171-cancel-identity-records-verify-request) | Submit a transaction to cancel identity records verification request.      |
| [`councilor`](#2172-councilor)                                                           | Councilor subcommands                                                      |
| [`delete-identity-records`](#2173-delete-identity-records)                               | Submit a transaction to delete an identity records.                        |
| [`handle-identity-records-verify-request`](#2174-handle-identity-records-verify-request) | Submit a transaction to approve or reject identity records verify request. |
| [`permission`](#2175-permission)                                                         | Permission management subcommands                                          |
| [`poll`](#2176-poll)                                                                     | Governance poll management subcommands                                     |
| [`proposal`](#2177-proposal)                                                             | Governance proposals management subcommands                                |
| [`register-identity-records`](#2178-register-identity-records)                           | Submit a transaction to create an identity record.                         |
| [`request-identity-record-verify`](#2179-request-identity-record-verify)                 | Submit a transaction to request an identity verify record.                 |
| [`role`](#21710-role)                                                                    | Role management subcommands                                                |
| [`set-execution-fee`](#21711-set-execution-fee)                                          | Submit a transaction to set execution fee                                  |
| [`set-network-properties`](#21712-set-network-properties)                                | Submit a transaction to set network properties                             |

[Return to top](#sekai)

##### 21.7.1 cancel-identity-records-verify-request

Submit a transaction to cancel identity records verification request.

Usage:
```
sekaid tx customgov cancel-identity-records-verify-request [id] [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ✅ yes |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ✅ yes |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ❌ ?   |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for cancel-identity-records-verify-request                                                                                                             | ✅ yes |
| `--infos-file string`         | The infos file for identity request.                                                                                                                        | ✅ yes |
| `--infos-json string`         | The infos json for identity request.                                                                                                                        | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default `'home'` directory will be used                                                                       | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ✅ yes |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ✅ yes |
| `--record-ids string`         | Concatenated identity record ids array. e.g. `1,2`                                                                                                          | ✅ yes |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `--tip string`                | The tip to be given to the verifier.                                                                                                                        | ✅ yes |
| `--verifier string`           | The verifier of the record ids                                                                                                                              | ✅ yes |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx customgov cancel-identity-records-verify-request --help
Submit a transaction to cancel identity records verification request.

Usage:
  sekaid tx customgov cancel-identity-records-verify-request [id] [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for cancel-identity-records-verify-request
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
/# sekaid tx customgov cancel-identity-records-verify-request 5 --from=kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --fees=100ukex --home=/root/.sekai --chain-id=localnet-4 --keyring-backend=test --yes --output=json | jq
{
  "height": "0",
  "txhash": "9659E50E56F715EE67C935B2F440A7F02BBA131DDD7697E7510A5F866DAC8BA5",
  "codespace": "",
  "code": 0,
  "data": "",
  "raw_log": "[]",
  "logs": [],
  "info": "",
  "gas_wanted": "0",
  "gas_used": "0",
  "tx": null,
  "timestamp": "",
  "events": []
}
```

<details>
  <summary>Check tx execution</summary>

  ```
  sekaid q tx 9659E50E56F715EE67C935B2F440A7F02BBA131DDD7697E7510A5F866DAC8BA5 --output=json | jq
  ```

  ```json
  {
    "height": "109602",
    "txhash": "9659E50E56F715EE67C935B2F440A7F02BBA131DDD7697E7510A5F866DAC8BA5",
    "codespace": "",
    "code": 0,
    "data": "0A310A2F2F6B6972612E676F762E4D736743616E63656C4964656E746974795265636F72647356657269667952657175657374",
    "raw_log": "[{\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\"},{\"key\":\"amount\",\"value\":\"200ukex\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"kira18q47tkrtrthf72xyscens2kv58ufpqtyxvxwuv\"},{\"key\":\"amount\",\"value\":\"200ukex\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/kira.gov.MsgCancelIdentityRecordsVerifyRequest\"},{\"key\":\"sender\",\"value\":\"kira18q47tkrtrthf72xyscens2kv58ufpqtyxvxwuv\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\"},{\"key\":\"sender\",\"value\":\"kira18q47tkrtrthf72xyscens2kv58ufpqtyxvxwuv\"},{\"key\":\"amount\",\"value\":\"200ukex\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "coin_received",
            "attributes": [
              {
                "key": "receiver",
                "value": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4"
              },
              {
                "key": "amount",
                "value": "200ukex"
              }
            ]
          },
          {
            "type": "coin_spent",
            "attributes": [
              {
                "key": "spender",
                "value": "kira18q47tkrtrthf72xyscens2kv58ufpqtyxvxwuv"
              },
              {
                "key": "amount",
                "value": "200ukex"
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/kira.gov.MsgCancelIdentityRecordsVerifyRequest"
              },
              {
                "key": "sender",
                "value": "kira18q47tkrtrthf72xyscens2kv58ufpqtyxvxwuv"
              }
            ]
          },
          {
            "type": "transfer",
            "attributes": [
              {
                "key": "recipient",
                "value": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4"
              },
              {
                "key": "sender",
                "value": "kira18q47tkrtrthf72xyscens2kv58ufpqtyxvxwuv"
              },
              {
                "key": "amount",
                "value": "200ukex"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "0",
    "gas_used": "0",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/kira.gov.MsgCancelIdentityRecordsVerifyRequest",
            "executor": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
            "verify_request_id": "5"
          }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
      },
      "auth_info": {
        "signer_infos": [
          {
            "public_key": {
              "@type": "/cosmos.crypto.secp256k1.PubKey",
              "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "92"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "ukex",
              "amount": "100"
            }
          ],
          "gas_limit": "200000",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
        "G5+araW7aplPdSJiPAvrgkmVMdmE14SNi6QmKgS0t0cEKLDeJ/BotFN2ktlIYl7nDWZ7zNEFYgsLtRKTT1jtfw=="
      ]
    },
    "timestamp": "2023-06-08T14:26:31Z",
    "events": [
      {
        "type": "tx",
        "attributes": [
          {
            "key": "YWNjX3NlcQ==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NC85Mg==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "RzUrYXJhVzdhcGxQZFNKaVBBdnJna21WTWRtRTE0U05pNlFtS2dTMHQwY0VLTERlSi9Cb3RGTjJrdGxJWWw3bkRXWjd6TkVGWWdzTHRSS1RUMWp0Znc9PQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "ZmVl",
            "value": "MTAwdWtleA==",
            "index": true
          },
          {
            "key": "ZmVlX3BheWVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "L2tpcmEuZ292Lk1zZ0NhbmNlbElkZW50aXR5UmVjb3Jkc1ZlcmlmeVJlcXVlc3Q=",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTE4cTQ3dGtydHJ0aGY3Mnh5c2NlbnMya3Y1OHVmcHF0eXh2eHd1dg==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTE4cTQ3dGtydHJ0aGY3Mnh5c2NlbnMya3Y1OHVmcHF0eXh2eHd1dg==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTE4cTQ3dGtydHJ0aGY3Mnh5c2NlbnMya3Y1OHVmcHF0eXh2eHd1dg==",
            "index": true
          }
        ]
      }
    ]
  }
  ```
</details>

<details>
  <summary>Check cancelation</summary>

  ```
  sekaid query customgov identity-record-verify-request 5 -o json | jq
  {
    "verify_record": null
  }
  ```
</details>

[Return to top](#sekai)

##### 21.7.2 councilor

[Return to top](#sekai)

##### 21.7.3 delete-identity-records

Submit a transaction to delete an identity records.

Usage:
```
sekaid tx customgov delete-identity-records [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `--approve`                   | The flag to approve or reject (default `true`)                                                                                                              | ✅ yes |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ✅ yes |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ✅ yes |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ✅ yes |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for delete-identity-records                                                                                                                            | ✅ yes |
| `--infos-file string`         | The infos file for identity request.                                                                                                                        | ✅ yes |
| `--infos-json string`         | The infos json for identity request.                                                                                                                        | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default `'home'` directory will be used                                                                       | ✅ yes |
| `--keys string`               | The keys to remove.                                                                                                                                         | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ✅ yes |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ✅ yes |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ✅ yes |


| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx customgov delete-identity-records --help
Submit a transaction to delete an identity records.

Usage:
  sekaid tx customgov delete-identity-records [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for delete-identity-records
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --keys string              The keys to remove.
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
/# sekaid tx customgov delete-identity-records --from=kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --home=/root/.sekai --chain-id=localnet-4 --keyring-backend=test --keys=text --yes --fees=100ukex --output=json | jq
{
  "height": "0",
  "txhash": "DA21A3D291C33308750ACED129E755B884D8B05DFCC94B8F420D63E5454AF9CA",
  "codespace": "",
  "code": 0,
  "data": "",
  "raw_log": "[]",
  "logs": [],
  "info": "",
  "gas_wanted": "0",
  "gas_used": "0",
  "tx": null,
  "timestamp": "",
  "events": []
}
```

<details>
  <summary>Check tx execution</summary>

  ```
  sekaid q tx DA21A3D291C33308750ACED129E755B884D8B05DFCC94B8F420D63E5454AF9CA --output=json | jq
  ```

  ```json
  {
    "height": "109446",
    "txhash": "DA21A3D291C33308750ACED129E755B884D8B05DFCC94B8F420D63E5454AF9CA",
    "codespace": "",
    "code": 0,
    "data": "0A240A222F6B6972612E676F762E4D736744656C6574654964656E746974795265636F726473",
    "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/kira.gov.MsgDeleteIdentityRecords\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/kira.gov.MsgDeleteIdentityRecords"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "0",
    "gas_used": "0",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/kira.gov.MsgDeleteIdentityRecords",
            "address": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
            "keys": [
              "text"
            ]
          }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
      },
      "auth_info": {
        "signer_infos": [
          {
            "public_key": {
              "@type": "/cosmos.crypto.secp256k1.PubKey",
              "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "90"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "ukex",
              "amount": "100"
            }
          ],
          "gas_limit": "200000",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
        "uFhwO7ESlxHgYUG97FYQsMrBxyH67Q8ZGIRS1R8mh+YFO6PQretDorKsuRdkwz03Ua9ea2ahichcLiYdZXWhQQ=="
      ]
    },
    "timestamp": "2023-06-08T13:59:42Z",
    "events": [
      {
        "type": "tx",
        "attributes": [
          {
            "key": "YWNjX3NlcQ==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NC85MA==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "dUZod083RVNseEhnWVVHOTdGWVFzTXJCeHlINjdROFpHSVJTMVI4bWgrWUZPNlBRcmV0RG9yS3N1UmRrd3owM1VhOWVhMmFoaWNoY0xpWWRaWFdoUVE9PQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "ZmVl",
            "value": "MTAwdWtleA==",
            "index": true
          },
          {
            "key": "ZmVlX3BheWVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "L2tpcmEuZ292Lk1zZ0RlbGV0ZUlkZW50aXR5UmVjb3Jkcw==",
            "index": true
          }
        ]
      }
    ]
  }
  ```
</details>

<details>
  <summary>Check removal</summary>

  ```
  /# sekaid query customgov identity-record 19 -o json | jq
  {
    "record": null
  }
  ```
</details>

[Return to top](#sekai)

##### 21.7.4 handle-identity-records-verify-request

Submit a transaction to approve or reject identity records verify request.

Usage:
```
sekaid tx customgov handle-identity-records-verify-request [id] [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `--approve`                   | The flag to approve or reject (default `true`)                                                                                                              | ✅ yes |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ✅ yes |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ✅ yes |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ✅ yes |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for handle-identity-records-verify-request                                                                                                             | ✅ yes |
| `--infos-file string`         | The infos file for identity request.                                                                                                                        | ✅ yes |
| `--infos-json string`         | The infos json for identity request.                                                                                                                        | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default `'home'` directory will be used                                                                       | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ✅ yes |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ✅ yes |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx customgov handle-identity-records-verify-request --help
Submit a transaction to approve or reject identity records verify request.

Usage:
  sekaid tx customgov handle-identity-records-verify-request [id] [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
      --approve                  The flag to approve or reject (default true)
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for handle-identity-records-verify-request
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid tx customgov handle-identity-records-verify-request 4 --home=/root/.sekai  --keyring-backend=test --chain-id=localnet-4 --from=kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt --approve=true --yes --fees=100ukex -o json | jq
{
  "height": "0",
  "txhash": "C2E70D94F60E97251B46702D0258F361969B9DA92511894514872338BC36F497",
  "codespace": "",
  "code": 0,
  "data": "",
  "raw_log": "[]",
  "logs": [],
  "info": "",
  "gas_wanted": "0",
  "gas_used": "0",
  "tx": null,
  "timestamp": "",
  "events": []
}
```

<details>
  <summary>Check tx execution</summary>

  ```
  sekaid q tx C2E70D94F60E97251B46702D0258F361969B9DA92511894514872338BC36F497 -o json | jq
  ```

  ```json
  {
    "height": "109272",
    "txhash": "C2E70D94F60E97251B46702D0258F361969B9DA92511894514872338BC36F497",
    "codespace": "",
    "code": 0,
    "data": "0A310A2F2F6B6972612E676F762E4D736748616E646C654964656E746974795265636F72647356657269667952657175657374",
    "raw_log": "[{\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt\"},{\"key\":\"amount\",\"value\":\"200ukex\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"kira18q47tkrtrthf72xyscens2kv58ufpqtyxvxwuv\"},{\"key\":\"amount\",\"value\":\"200ukex\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/kira.gov.MsgHandleIdentityRecordsVerifyRequest\"},{\"key\":\"sender\",\"value\":\"kira18q47tkrtrthf72xyscens2kv58ufpqtyxvxwuv\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt\"},{\"key\":\"sender\",\"value\":\"kira18q47tkrtrthf72xyscens2kv58ufpqtyxvxwuv\"},{\"key\":\"amount\",\"value\":\"200ukex\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "coin_received",
            "attributes": [
              {
                "key": "receiver",
                "value": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt"
              },
              {
                "key": "amount",
                "value": "200ukex"
              }
            ]
          },
          {
            "type": "coin_spent",
            "attributes": [
              {
                "key": "spender",
                "value": "kira18q47tkrtrthf72xyscens2kv58ufpqtyxvxwuv"
              },
              {
                "key": "amount",
                "value": "200ukex"
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/kira.gov.MsgHandleIdentityRecordsVerifyRequest"
              },
              {
                "key": "sender",
                "value": "kira18q47tkrtrthf72xyscens2kv58ufpqtyxvxwuv"
              }
            ]
          },
          {
            "type": "transfer",
            "attributes": [
              {
                "key": "recipient",
                "value": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt"
              },
              {
                "key": "sender",
                "value": "kira18q47tkrtrthf72xyscens2kv58ufpqtyxvxwuv"
              },
              {
                "key": "amount",
                "value": "200ukex"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "0",
    "gas_used": "0",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/kira.gov.MsgHandleIdentityRecordsVerifyRequest",
            "verifier": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
            "verify_request_id": "4",
            "yes": true
          }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
      },
      "auth_info": {
        "signer_infos": [
          {
            "public_key": {
              "@type": "/cosmos.crypto.secp256k1.PubKey",
              "key": "A5mB81789jXij6eUh5QGrRlhXdLheHFL1ix1LtxfMCvJ"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "5"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "ukex",
              "amount": "100"
            }
          ],
          "gas_limit": "200000",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
        "7HuRJY8w20vSXg5Vu/3p/Yjqm/KRTAvduyOw62vzmlAFxjQY5jME7Xdo1r4fzPPD3wRBJV38exA1z+I3bZNLTQ=="
      ]
    },
    "timestamp": "2023-06-08T13:29:48Z",
    "events": [
      {
        "type": "tx",
        "attributes": [
          {
            "key": "YWNjX3NlcQ==",
            "value": "a2lyYTE3YWVxeHZrbDNnMzdwZmNnd2txeHZrcm42M2pmbGpmdmpyYXZudC81",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "N0h1UkpZOHcyMHZTWGc1VnUvM3AvWWpxbS9LUlRBdmR1eU93NjJ2em1sQUZ4alFZNWpNRTdYZG8xcjRmelBQRDN3UkJKVjM4ZXhBMXorSTNiWk5MVFE9PQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTE3YWVxeHZrbDNnMzdwZmNnd2txeHZrcm42M2pmbGpmdmpyYXZudA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTE3YWVxeHZrbDNnMzdwZmNnd2txeHZrcm42M2pmbGpmdmpyYXZudA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTE3YWVxeHZrbDNnMzdwZmNnd2txeHZrcm42M2pmbGpmdmpyYXZudA==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "ZmVl",
            "value": "MTAwdWtleA==",
            "index": true
          },
          {
            "key": "ZmVlX3BheWVy",
            "value": "a2lyYTE3YWVxeHZrbDNnMzdwZmNnd2txeHZrcm42M2pmbGpmdmpyYXZudA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "L2tpcmEuZ292Lk1zZ0hhbmRsZUlkZW50aXR5UmVjb3Jkc1ZlcmlmeVJlcXVlc3Q=",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTE4cTQ3dGtydHJ0aGY3Mnh5c2NlbnMya3Y1OHVmcHF0eXh2eHd1dg==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTE3YWVxeHZrbDNnMzdwZmNnd2txeHZrcm42M2pmbGpmdmpyYXZudA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTE3YWVxeHZrbDNnMzdwZmNnd2txeHZrcm42M2pmbGpmdmpyYXZudA==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTE4cTQ3dGtydHJ0aGY3Mnh5c2NlbnMya3Y1OHVmcHF0eXh2eHd1dg==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTE4cTQ3dGtydHJ0aGY3Mnh5c2NlbnMya3Y1OHVmcHF0eXh2eHd1dg==",
            "index": true
          }
        ]
      }
    ]
  }
  ```

</details>

<details>
  <summary>Check verification</summary>

  ```
  sekaid query customgov identity-record 19 -o json | jq
  ```

  > After the accepting verification the indentity will have the `verifiers`:

  ```json
  {
    "record": {
      "id": "19",
      "address": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
      "key": "text",
      "value": "Testing item regisrty for docs",
      "date": "2023-06-08T12:25:11.014802004Z",
      "verifiers": [
        "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt"
      ]
    }
  }
  ```
</details>

[Return to top](#sekai)

##### 21.7.5 permission

Permission management subcommands.

Usage:
```
sekaid tx customgov permission [flags]
sekaid tx customgov permission [command]
```

Available Commands:

| Subcommands                                       | Description                                   |
| ------------------------------------------------- | --------------------------------------------- |
| [`blacklist`](#21751-blacklist)                   | Assign permission to a kira account blacklist |
| [`remove-blacklisted`](#21752-remove-blacklisted) | Remove blacklisted permission from an address |
| [`remove-whitelisted`](#21753-remove-whitelisted) | Remove whitelisted permission from an address |
| [`whitelist`](#21754-whitelist)                   | Assign permission to a kira address whitelist |



| Flags        | Description         | Work  |
| ------------ | ------------------- | ----- |
| `-h, --help` | help for permission | ✅ yes |



| Global Flags          | Description                                                                            | Work |
| --------------------- | -------------------------------------------------------------------------------------- | ---- |
| `--chain-id string`   | The network chain ID                                                                   | ❌ ?  |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ❌ ?  |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?  |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?  |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?  |

```
/# sekaid tx customgov permission --help
Permission management subcommands

Usage:
  sekaid tx customgov permission [flags]
  sekaid tx customgov permission [command]

Available Commands:
  blacklist          Assign permission to a kira account blacklist
  remove-blacklisted Remove blacklisted permission from an address
  remove-whitelisted Remove whitelisted permission from an address
  whitelist          Assign permission to a kira address whitelist

Flags:
  -h, --help   help for permission

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors

Use "sekaid tx customgov permission [command] --help" for more information about a command.
```

[Return to top](#sekai)

###### 21.7.5.1 blacklist

[Return to top](#sekai)

###### 21.7.5.2 remove-blacklisted

[Return to top](#sekai)

###### 21.7.5.3 remove-whitelisted

[Return to top](#sekai)

###### 21.7.5.4 whitelist

[Return to top](#sekai)


##### 21.7.6 poll

[Return to top](#sekai)

##### 21.7.7 proposal

Governance proposals management subcommands.

Usage:
```
sekaid tx customgov proposal [flags]
sekaid tx customgov proposal [command]
```

Available Commands:

| Subcommands                                                                         | Description                                             |
| ----------------------------------------------------------------------------------- | ------------------------------------------------------- |
| [`account`](#21771-account)                                                         | Account proposals management subcommands                |
| [`proposal-jail-councilor`](#21772-proposal-jail-councilor)                         | Create a proposal to jail councilors                    |
| [`proposal-reset-whole-councilor-rank`](#21773-proposal-reset-whole-councilor-rank) | Create a proposal to reset whole councilor rank         |
| [`role`](#21774-role)                                                               | Role proposals management subcommands                   |
| [`set-network-property`](#21775-set-network-property)                               | Create a proposal to set a value on a network property. |
| [`set-poor-network-msgs`](#21776-set-poor-network-msgs)                             | Create a proposal to set a value on a network property. |
| [`set-proposal-durations-proposal`](#21777-set-proposal-durations-proposal)         | Create a proposal to set batch proposal durations.      |
| [`upsert-data-registry`](#21778-upsert-data-registry)                               | Create a proposal to upsert a key in the data registry  |
| [`vote`](#21779-vote)                                                               | Vote a proposal.                                        |



| Flags        | Description       | Work  |
| ------------ | ----------------- | ----- |
| `-h, --help` | help for proposal | ✅ yes |



| Global Flags          | Description                                                                            | Work |
| --------------------- | -------------------------------------------------------------------------------------- | ---- |
| `--chain-id string`   | The network chain ID                                                                   | ❌ ?  |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ❌ ?  |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?  |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?  |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?  |

```
sekaid tx customgov proposal --help
Governance proposals management subcommands

Usage:
  sekaid tx customgov proposal [flags]
  sekaid tx customgov proposal [command]

Available Commands:
  account                             Account proposals management subcommands
  proposal-jail-councilor             Create a proposal to jail councilors
  proposal-reset-whole-councilor-rank Create a proposal to reset whole councilor rank
  role                                Role proposals management subcommands
  set-network-property                Create a proposal to set a value on a network property.
  set-poor-network-msgs               Create a proposal to set a value on a network property.
  set-proposal-durations-proposal     Create a proposal to set batch proposal durations.
  upsert-data-registry                Create a proposal to upsert a key in the data registry
  vote                                Vote a proposal.

Flags:
  -h, --help   help for proposal

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors

Use "sekaid tx customgov proposal [command] --help" for more information about a command.
```

[Return to top](#sekai)

###### 21.7.7.1 account

Account proposals management subcommands.

Usage:
```
sekaid tx customgov proposal account [flags]
sekaid tx customgov proposal account [command]
```

Available Commands:

| Subcommand                                                               | Description                                                           |
| ------------------------------------------------------------------------ | --------------------------------------------------------------------- |
| [`assign-role`](#217711-assign-role)                                     | Create a proposal to assign a role to an address.                     |
| [`blacklist-permission`](#217712-blacklist-permission)                   | Create a proposal to blacklist a permission to an address.            |
| [`remove-blacklisted-permission`](#217713-remove-blacklisted-permission) | Create a proposal to remove a blacklisted permission from an address. |
| [`remove-whitelisted-permission`](#217714-remove-whitelisted-permission) | Create a proposal to remove a whitelisted permission from an address. |
| [`unassign-role`](#217715-unassign-role)                                 | Create a proposal to unassign a role from an address.                 |
| [`whitelist-permission`](#217716-whitelist-permission)                   | Create a proposal to whitelist a permission to an address.            |



| Flags         | Description      | Work  |
| ------------- | ---------------- | ----- |
| ` -h, --help` | help for account | ✅ yes |



| Global Flags          | Description                                                                            | Work |
| --------------------- | -------------------------------------------------------------------------------------- | ---- |
| `--chain-id string`   | The network chain ID                                                                   | ❌ ?  |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ❌ ?  |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?  |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?  |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?  |

```
/# sekaid tx customgov proposal account --help
Account proposals management subcommands

Usage:
  sekaid tx customgov proposal account [flags]
  sekaid tx customgov proposal account [command]

Available Commands:
  assign-role                   Create a proposal to assign a role to an address.
  blacklist-permission          Create a proposal to blacklist a permission to an address.
  remove-blacklisted-permission Create a proposal to remove a blacklisted permission from an address.
  remove-whitelisted-permission Create a proposal to remove a whitelisted permission from an address.
  unassign-role                 Create a proposal to unassign a role from an address.
  whitelist-permission          Create a proposal to whitelist a permission to an address.

Flags:
  -h, --help   help for account

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors

Use "sekaid tx customgov proposal account [command] --help" for more information about a command.
```

[Return to top](#sekai)

###### 21.7.7.1.1 assign-role

Create a proposal to assign a role to an address.

Usage:
```
sekaid tx customgov proposal account assign-role role [role_identifier] [flags] ❌ mistake! ... assgin-role [role_sid | role_id] [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `--addr string`               | the address to set permissions                                                                                                                              | ✅ yes |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ❌ ?   |
| `--description string`        | The description of the proposal, it can be a url, some text, etc.                                                                                           | ✅ yes |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ✅ yes |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ✅ yes |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for assign-role                                                                                                                                        | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default 'home' directory will be used                                                                         | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ❌ ?   |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ✅ yes |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `--title string`              | The title of the proposal.                                                                                                                                  | ✅ yes |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx customgov proposal account assign-role --help
Create a proposal to assign a role to an address.

Usage:
  sekaid tx customgov proposal account assign-role role [role_identifier] [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
      --addr string              the address to set permissions
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --description string       The description of the proposal, it can be a url, some text, etc.
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for assign-role
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
      --title string             The title of the proposal.
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

> `--description` and `--title` are required

```
sekaid tx customgov proposal account assign-role 4 --from=kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --addr=kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt --home=/root/.sekai --chain-id=localnet-4 --fees=10ukex --keyring-backend=test --description="Test proposal for docs" --title="Test proposal" --yes -o json | jq
{
  "height": "0",
  "txhash": "99E4837DABC3A3837E6FEFEF6382F3F07F9306068D3CF55EE92AE180F3D091AE",
  "codespace": "",
  "code": 0,
  "data": "",
  "raw_log": "[]",
  "logs": [],
  "info": "",
  "gas_wanted": "0",
  "gas_used": "0",
  "tx": null,
  "timestamp": "",
  "events": []
}
```

<details>
  <summary>Check tx execution</summary>

  ```
  sekaid q tx 99E4837DABC3A3837E6FEFEF6382F3F07F9306068D3CF55EE92AE180F3D091AE -o json | jq
  ```

  ```json
  {
    "height": "141555",
    "txhash": "99E4837DABC3A3837E6FEFEF6382F3F07F9306068D3CF55EE92AE180F3D091AE",
    "codespace": "",
    "code": 0,
    "data": "0A210A1B2F6B6972612E676F762E4D73675375626D697450726F706F73616C1202080E",
    "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/kira.gov.MsgSubmitProposal\"}]},{\"type\":\"submit_proposal\",\"attributes\":[{\"key\":\"proposal_id\",\"value\":\"14\"},{\"key\":\"proposal_type\",\"value\":\"AssignRoleToAccount\"},{\"key\":\"proposal_content\",\"value\":\"proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\\ntitle: Test proposal\\ndescription: Test proposal for docs\\ncontent:\\n  typeurl: /kira.gov.AssignRoleToAccountProposal\\n  value:\\n  - 10\\n  - 20\\n  - 247\\n  - 114\\n  - 3\\n  - 50\\n  - 223\\n  - 138\\n  - 35\\n  - 224\\n  - 167\\n  - 8\\n  - 117\\n  - 128\\n  - 102\\n  - 88\\n  - 115\\n  - 212\\n  - 100\\n  - 159\\n  - 201\\n  - 44\\n  - 18\\n  - 1\\n  - 52\\n  xxx_nounkeyedliteral: {}\\n  xxx_unrecognized: []\\n  xxx_sizecache: 0\\n\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/kira.gov.MsgSubmitProposal"
              }
            ]
          },
          {
            "type": "submit_proposal",
            "attributes": [
              {
                "key": "proposal_id",
                "value": "14"
              },
              {
                "key": "proposal_type",
                "value": "AssignRoleToAccount"
              },
              {
                "key": "proposal_content",
                "value": "proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\ntitle: Test proposal\ndescription: Test proposal for docs\ncontent:\n  typeurl: /kira.gov.AssignRoleToAccountProposal\n  value:\n  - 10\n  - 20\n  - 247\n  - 114\n  - 3\n  - 50\n  - 223\n  - 138\n  - 35\n  - 224\n  - 167\n  - 8\n  - 117\n  - 128\n  - 102\n  - 88\n  - 115\n  - 212\n  - 100\n  - 159\n  - 201\n  - 44\n  - 18\n  - 1\n  - 52\n  xxx_nounkeyedliteral: {}\n  xxx_unrecognized: []\n  xxx_sizecache: 0\n"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "0",
    "gas_used": "0",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/kira.gov.MsgSubmitProposal",
            "proposer": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
            "title": "Test proposal",
            "description": "Test proposal for docs",
            "content": {
              "@type": "/kira.gov.AssignRoleToAccountProposal",
              "address": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
              "role_identifier": "4"
            }
          }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
      },
      "auth_info": {
        "signer_infos": [
          {
            "public_key": {
              "@type": "/cosmos.crypto.secp256k1.PubKey",
              "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "97"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "ukex",
              "amount": "10"
            }
          ],
          "gas_limit": "200000",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
        "nmKNCKdiAhL+Hr5w10UPuiUmkDBQoKfp/5QCIheclaYHWUZ54BpDoYM16xDnAfKaIs0wGu/rc0hG2PpWDvjO/w=="
      ]
    },
    "timestamp": "2023-06-12T09:59:07Z",
    "events": [
      {
        "type": "tx",
        "attributes": [
          {
            "key": "YWNjX3NlcQ==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NC85Nw==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "bm1LTkNLZGlBaEwrSHI1dzEwVVB1aVVta0RCUW9LZnAvNVFDSWhlY2xhWUhXVVo1NEJwRG9ZTTE2eERuQWZLYUlzMHdHdS9yYzBoRzJQcFdEdmpPL3c9PQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTB1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTB1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTB1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "ZmVl",
            "value": "MTB1a2V4",
            "index": true
          },
          {
            "key": "ZmVlX3BheWVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "L2tpcmEuZ292Lk1zZ1N1Ym1pdFByb3Bvc2Fs",
            "index": true
          }
        ]
      },
      {
        "type": "submit_proposal",
        "attributes": [
          {
            "key": "cHJvcG9zYWxfaWQ=",
            "value": "MTQ=",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfdHlwZQ==",
            "value": "QXNzaWduUm9sZVRvQWNjb3VudA==",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfY29udGVudA==",
            "value": "cHJvcG9zZXI6IGtpcmExdm13ZGd3NDI2YWo5ZngzM2ZxdXNtdGc2cjY1eXl1Y214NnJkdDQKdGl0bGU6IFRlc3QgcHJvcG9zYWwKZGVzY3JpcHRpb246IFRlc3QgcHJvcG9zYWwgZm9yIGRvY3MKY29udGVudDoKICB0eXBldXJsOiAva2lyYS5nb3YuQXNzaWduUm9sZVRvQWNjb3VudFByb3Bvc2FsCiAgdmFsdWU6CiAgLSAxMAogIC0gMjAKICAtIDI0NwogIC0gMTE0CiAgLSAzCiAgLSA1MAogIC0gMjIzCiAgLSAxMzgKICAtIDM1CiAgLSAyMjQKICAtIDE2NwogIC0gOAogIC0gMTE3CiAgLSAxMjgKICAtIDEwMgogIC0gODgKICAtIDExNQogIC0gMjEyCiAgLSAxMDAKICAtIDE1OQogIC0gMjAxCiAgLSA0NAogIC0gMTgKICAtIDEKICAtIDUyCiAgeHh4X25vdW5rZXllZGxpdGVyYWw6IHt9CiAgeHh4X3VucmVjb2duaXplZDogW10KICB4eHhfc2l6ZWNhY2hlOiAwCg==",
            "index": true
          }
        ]
      }
    ]
  }
  ```
</details>

<details>
  <summary>Check proposal</summary>

  ```
  sekaid q customgov proposal 14 -o json | jq
  ```

  ```json
  {
    "proposal_id": "14",
    "title": "Test proposal",
    "description": "Test proposal for docs",
    "content": {
      "@type": "/kira.gov.AssignRoleToAccountProposal",
      "address": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
      "role_identifier": "4"
    },
    "submit_time": "2023-06-12T09:59:07.424361608Z",
    "voting_end_time": "2023-06-12T10:05:07.424361608Z",
    "enactment_end_time": "2023-06-12T10:10:07.424361608Z",
    "min_voting_end_block_height": "141557",
    "min_enactment_end_block_height": "141558",
    "result": "VOTE_PENDING",
    "exec_result": ""
  }
  ```
</details>

[Return to top](#sekai)

###### 21.7.7.1.2 blacklist-permission

Create a proposal to blacklist a permission to an address.

Usage:
```
sekaid tx customgov proposal account blacklist-permission [permission_id] [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `--addr string`               | the address to set permissions                                                                                                                              | ✅ yes |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ❌ ?   |
| `--description string`        | The description of the proposal, it can be a url, some text, etc.                                                                                           | ✅ yes |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ✅ yes |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ✅ yes |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for blacklist-permission                                                                                                                               | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default 'home' directory will be used                                                                         | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ❌ ?   |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ✅ yes |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `--title string`              | The title of the proposal.                                                                                                                                  | ✅ yes |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx customgov proposal account blacklist-permission --help
Create a proposal to blacklist a permission to an address.

Usage:
  sekaid tx customgov proposal account blacklist-permission [permission_id] [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
      --addr string              the address to set permissions
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --description string       The description of the proposal, it can be a url, some text, etc.
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for blacklist-permission
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
      --title string             The title of the proposal.
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

> **We can't blacklist whitelisted permission**

```
sekaid tx customgov proposal account blacklist-permission 14 --from=kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --addr=kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt --home=/root/.sekai --chain-id=localnet-4 --fees=10ukex --keyring-backend=test --description="Test proposal for docs" --title="Test proposal" --yes -o json | jq
{
  "height": "0",
  "txhash": "DD748427FEA3D3866007A13887152A5369106380886A47669FFD4932635AC353",
  "codespace": "",
  "code": 0,
  "data": "",
  "raw_log": "[]",
  "logs": [],
  "info": "",
  "gas_wanted": "0",
  "gas_used": "0",
  "tx": null,
  "timestamp": "",
  "events": []
}
```

<details>
  <summary>Check tx execution</summary>

  ```
  sekaid q tx DD748427FEA3D3866007A13887152A5369106380886A47669FFD4932635AC353 -o json | jq
  ```

  ```json
  {
    "height": "141782",
    "txhash": "DD748427FEA3D3866007A13887152A5369106380886A47669FFD4932635AC353",
    "codespace": "",
    "code": 0,
    "data": "0A210A1B2F6B6972612E676F762E4D73675375626D697450726F706F73616C12020810",
    "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/kira.gov.MsgSubmitProposal\"}]},{\"type\":\"submit_proposal\",\"attributes\":[{\"key\":\"proposal_id\",\"value\":\"16\"},{\"key\":\"proposal_type\",\"value\":\"BlacklistAccountPermission\"},{\"key\":\"proposal_content\",\"value\":\"proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\\ntitle: Test proposal\\ndescription: Test proposal for docs\\ncontent:\\n  typeurl: /kira.gov.BlacklistAccountPermissionProposal\\n  value:\\n  - 10\\n  - 20\\n  - 247\\n  - 114\\n  - 3\\n  - 50\\n  - 223\\n  - 138\\n  - 35\\n  - 224\\n  - 167\\n  - 8\\n  - 117\\n  - 128\\n  - 102\\n  - 88\\n  - 115\\n  - 212\\n  - 100\\n  - 159\\n  - 201\\n  - 44\\n  - 16\\n  - 14\\n  xxx_nounkeyedliteral: {}\\n  xxx_unrecognized: []\\n  xxx_sizecache: 0\\n\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/kira.gov.MsgSubmitProposal"
              }
            ]
          },
          {
            "type": "submit_proposal",
            "attributes": [
              {
                "key": "proposal_id",
                "value": "16"
              },
              {
                "key": "proposal_type",
                "value": "BlacklistAccountPermission"
              },
              {
                "key": "proposal_content",
                "value": "proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\ntitle: Test proposal\ndescription: Test proposal for docs\ncontent:\n  typeurl: /kira.gov.BlacklistAccountPermissionProposal\n  value:\n  - 10\n  - 20\n  - 247\n  - 114\n  - 3\n  - 50\n  - 223\n  - 138\n  - 35\n  - 224\n  - 167\n  - 8\n  - 117\n  - 128\n  - 102\n  - 88\n  - 115\n  - 212\n  - 100\n  - 159\n  - 201\n  - 44\n  - 16\n  - 14\n  xxx_nounkeyedliteral: {}\n  xxx_unrecognized: []\n  xxx_sizecache: 0\n"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "0",
    "gas_used": "0",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/kira.gov.MsgSubmitProposal",
            "proposer": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
            "title": "Test proposal",
            "description": "Test proposal for docs",
            "content": {
              "@type": "/kira.gov.BlacklistAccountPermissionProposal",
              "address": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
              "permission": "PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL"
            }
          }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
      },
      "auth_info": {
        "signer_infos": [
          {
            "public_key": {
              "@type": "/cosmos.crypto.secp256k1.PubKey",
              "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "100"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "ukex",
              "amount": "10"
            }
          ],
          "gas_limit": "200000",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
        "NnDiSXreT1opubriO91tTq2DlIo+VWJdRqbE+BnBY5F08FSlmq6ceS/Ltlprc9twZYKkTT+YrwpOoI2Xi7wxmA=="
      ]
    },
    "timestamp": "2023-06-12T10:38:08Z",
    "events": [
      {
        "type": "tx",
        "attributes": [
          {
            "key": "YWNjX3NlcQ==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NC8xMDA=",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "Tm5EaVNYcmVUMW9wdWJyaU85MXRUcTJEbElvK1ZXSmRScWJFK0JuQlk1RjA4RlNsbXE2Y2VTL0x0bHByYzl0d1pZS2tUVCtZcndwT29JMlhpN3d4bUE9PQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTB1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTB1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTB1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "ZmVl",
            "value": "MTB1a2V4",
            "index": true
          },
          {
            "key": "ZmVlX3BheWVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "L2tpcmEuZ292Lk1zZ1N1Ym1pdFByb3Bvc2Fs",
            "index": true
          }
        ]
      },
      {
        "type": "submit_proposal",
        "attributes": [
          {
            "key": "cHJvcG9zYWxfaWQ=",
            "value": "MTY=",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfdHlwZQ==",
            "value": "QmxhY2tsaXN0QWNjb3VudFBlcm1pc3Npb24=",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfY29udGVudA==",
            "value": "cHJvcG9zZXI6IGtpcmExdm13ZGd3NDI2YWo5ZngzM2ZxdXNtdGc2cjY1eXl1Y214NnJkdDQKdGl0bGU6IFRlc3QgcHJvcG9zYWwKZGVzY3JpcHRpb246IFRlc3QgcHJvcG9zYWwgZm9yIGRvY3MKY29udGVudDoKICB0eXBldXJsOiAva2lyYS5nb3YuQmxhY2tsaXN0QWNjb3VudFBlcm1pc3Npb25Qcm9wb3NhbAogIHZhbHVlOgogIC0gMTAKICAtIDIwCiAgLSAyNDcKICAtIDExNAogIC0gMwogIC0gNTAKICAtIDIyMwogIC0gMTM4CiAgLSAzNQogIC0gMjI0CiAgLSAxNjcKICAtIDgKICAtIDExNwogIC0gMTI4CiAgLSAxMDIKICAtIDg4CiAgLSAxMTUKICAtIDIxMgogIC0gMTAwCiAgLSAxNTkKICAtIDIwMQogIC0gNDQKICAtIDE2CiAgLSAxNAogIHh4eF9ub3Vua2V5ZWRsaXRlcmFsOiB7fQogIHh4eF91bnJlY29nbml6ZWQ6IFtdCiAgeHh4X3NpemVjYWNoZTogMAo=",
            "index": true
          }
        ]
      }
    ]
  }
  ```
</details>

<details>
  <summary>Check proposal</summary>

  ```
  sekaid query customgov proposal 16 -o json | jq
  ```

  ```json
  {
    "proposal_id": "16",
    "title": "Test proposal",
    "description": "Test proposal for docs",
    "content": {
      "@type": "/kira.gov.BlacklistAccountPermissionProposal",
      "address": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
      "permission": "PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL"
    },
    "submit_time": "2023-06-12T10:38:08.789572953Z",
    "voting_end_time": "2023-06-12T10:44:08.789572953Z",
    "enactment_end_time": "2023-06-12T10:49:08.789572953Z",
    "min_voting_end_block_height": "141784",
    "min_enactment_end_block_height": "141785",
    "result": "VOTE_PENDING",
    "exec_result": ""
  }
  ```
</details>

[Return to top](#sekai)

###### 21.7.7.1.3 remove-blacklisted-permission

Create a proposal to remove a blacklisted permission from an address.

Usage:
```
sekaid tx customgov proposal account remove-blacklisted-permission [permission_id] [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `--addr string`               | the address to set permissions                                                                                                                              | ✅ yes |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ❌ ?   |
| `--description string`        | The description of the proposal, it can be a url, some text, etc.                                                                                           | ✅ yes |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ✅ yes |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ✅ yes |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for remove-blacklisted-permission                                                                                                                      | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default 'home' directory will be used                                                                         | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ❌ ?   |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ✅ yes |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `--title string`              | The title of the proposal.                                                                                                                                  | ✅ yes |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
sekaid tx customgov proposal account remove-blacklisted-permission 10 --from=kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --addr=kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt --home=/root/.sekai --chain-id=localnet-4 --fees=10ukex --keyring-backend=test --description="Test proposal for docs" --title="Test proposal" --yes -o json | jq
{
  "height": "0",
  "txhash": "29E3FD5033DD532CE7B783F761F9C78B16B2D0F4FE93D0EA64B2B87A581B4F01",
  "codespace": "",
  "code": 0,
  "data": "",
  "raw_log": "[]",
  "logs": [],
  "info": "",
  "gas_wanted": "0",
  "gas_used": "0",
  "tx": null,
  "timestamp": "",
  "events": []
}
```

<details>
  <summary>Check tx execution</summary>

  ```
  sekaid q tx 29E3FD5033DD532CE7B783F761F9C78B16B2D0F4FE93D0EA64B2B87A581B4F01 -o json | jq
  ```

  ```json
  {
    "height": "141933",
    "txhash": "29E3FD5033DD532CE7B783F761F9C78B16B2D0F4FE93D0EA64B2B87A581B4F01",
    "codespace": "",
    "code": 0,
    "data": "0A210A1B2F6B6972612E676F762E4D73675375626D697450726F706F73616C12020811",
    "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/kira.gov.MsgSubmitProposal\"}]},{\"type\":\"submit_proposal\",\"attributes\":[{\"key\":\"proposal_id\",\"value\":\"17\"},{\"key\":\"proposal_type\",\"value\":\"RemoveBlacklistedAccountPermission\"},{\"key\":\"proposal_content\",\"value\":\"proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\\ntitle: Test proposal\\ndescription: Test proposal for docs\\ncontent:\\n  typeurl: /kira.gov.RemoveBlacklistedAccountPermissionProposal\\n  value:\\n  - 10\\n  - 20\\n  - 247\\n  - 114\\n  - 3\\n  - 50\\n  - 223\\n  - 138\\n  - 35\\n  - 224\\n  - 167\\n  - 8\\n  - 117\\n  - 128\\n  - 102\\n  - 88\\n  - 115\\n  - 212\\n  - 100\\n  - 159\\n  - 201\\n  - 44\\n  - 16\\n  - 10\\n  xxx_nounkeyedliteral: {}\\n  xxx_unrecognized: []\\n  xxx_sizecache: 0\\n\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/kira.gov.MsgSubmitProposal"
              }
            ]
          },
          {
            "type": "submit_proposal",
            "attributes": [
              {
                "key": "proposal_id",
                "value": "17"
              },
              {
                "key": "proposal_type",
                "value": "RemoveBlacklistedAccountPermission"
              },
              {
                "key": "proposal_content",
                "value": "proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\ntitle: Test proposal\ndescription: Test proposal for docs\ncontent:\n  typeurl: /kira.gov.RemoveBlacklistedAccountPermissionProposal\n  value:\n  - 10\n  - 20\n  - 247\n  - 114\n  - 3\n  - 50\n  - 223\n  - 138\n  - 35\n  - 224\n  - 167\n  - 8\n  - 117\n  - 128\n  - 102\n  - 88\n  - 115\n  - 212\n  - 100\n  - 159\n  - 201\n  - 44\n  - 16\n  - 10\n  xxx_nounkeyedliteral: {}\n  xxx_unrecognized: []\n  xxx_sizecache: 0\n"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "0",
    "gas_used": "0",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/kira.gov.MsgSubmitProposal",
            "proposer": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
            "title": "Test proposal",
            "description": "Test proposal for docs",
            "content": {
              "@type": "/kira.gov.RemoveBlacklistedAccountPermissionProposal",
              "address": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
              "permission": "PERMISSION_CREATE_UPSERT_DATA_REGISTRY_PROPOSAL"
            }
          }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
      },
      "auth_info": {
        "signer_infos": [
          {
            "public_key": {
              "@type": "/cosmos.crypto.secp256k1.PubKey",
              "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "101"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "ukex",
              "amount": "10"
            }
          ],
          "gas_limit": "200000",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
        "Ot4WugrrQGOdS0e1EikLGJog4Ne+he4NvH5yxW1clyQ5bApg0b7tQL+V6ZPIcpNOF7Jm2K6vwslcy4aInbwqkg=="
      ]
    },
    "timestamp": "2023-06-12T11:04:05Z",
    "events": [
      {
        "type": "tx",
        "attributes": [
          {
            "key": "YWNjX3NlcQ==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NC8xMDE=",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "T3Q0V3VncnJRR09kUzBlMUVpa0xHSm9nNE5lK2hlNE52SDV5eFcxY2x5UTViQXBnMGI3dFFMK1Y2WlBJY3BOT0Y3Sm0ySzZ2d3NsY3k0YUluYndxa2c9PQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTB1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTB1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTB1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "ZmVl",
            "value": "MTB1a2V4",
            "index": true
          },
          {
            "key": "ZmVlX3BheWVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "L2tpcmEuZ292Lk1zZ1N1Ym1pdFByb3Bvc2Fs",
            "index": true
          }
        ]
      },
      {
        "type": "submit_proposal",
        "attributes": [
          {
            "key": "cHJvcG9zYWxfaWQ=",
            "value": "MTc=",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfdHlwZQ==",
            "value": "UmVtb3ZlQmxhY2tsaXN0ZWRBY2NvdW50UGVybWlzc2lvbg==",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfY29udGVudA==",
            "value": "cHJvcG9zZXI6IGtpcmExdm13ZGd3NDI2YWo5ZngzM2ZxdXNtdGc2cjY1eXl1Y214NnJkdDQKdGl0bGU6IFRlc3QgcHJvcG9zYWwKZGVzY3JpcHRpb246IFRlc3QgcHJvcG9zYWwgZm9yIGRvY3MKY29udGVudDoKICB0eXBldXJsOiAva2lyYS5nb3YuUmVtb3ZlQmxhY2tsaXN0ZWRBY2NvdW50UGVybWlzc2lvblByb3Bvc2FsCiAgdmFsdWU6CiAgLSAxMAogIC0gMjAKICAtIDI0NwogIC0gMTE0CiAgLSAzCiAgLSA1MAogIC0gMjIzCiAgLSAxMzgKICAtIDM1CiAgLSAyMjQKICAtIDE2NwogIC0gOAogIC0gMTE3CiAgLSAxMjgKICAtIDEwMgogIC0gODgKICAtIDExNQogIC0gMjEyCiAgLSAxMDAKICAtIDE1OQogIC0gMjAxCiAgLSA0NAogIC0gMTYKICAtIDEwCiAgeHh4X25vdW5rZXllZGxpdGVyYWw6IHt9CiAgeHh4X3VucmVjb2duaXplZDogW10KICB4eHhfc2l6ZWNhY2hlOiAwCg==",
            "index": true
          }
        ]
      }
    ]
  }
  ```
</details>

<details>
  <summary>Check proposal</summary>

  ```
  sekaid q customgov proposal 17 -o json | jq
  ```

  ```json
  {
    "proposal_id": "17",
    "title": "Test proposal",
    "description": "Test proposal for docs",
    "content": {
      "@type": "/kira.gov.RemoveBlacklistedAccountPermissionProposal",
      "address": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
      "permission": "PERMISSION_CREATE_UPSERT_DATA_REGISTRY_PROPOSAL"
    },
    "submit_time": "2023-06-12T11:04:05.818914927Z",
    "voting_end_time": "2023-06-12T11:10:05.818914927Z",
    "enactment_end_time": "2023-06-12T11:15:05.818914927Z",
    "min_voting_end_block_height": "141935",
    "min_enactment_end_block_height": "141936",
    "result": "VOTE_PENDING",
    "exec_result": ""
  }
  ```
</details>

[Return to top](#sekai)

###### 21.7.7.1.4 remove-whitelisted-permission

Create a proposal to remove a whitelisted permission from an address.

Usage:
```
sekaid tx customgov proposal account remove-whitelisted-permission [permission_id] [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `--addr string`               | the address to set permissions                                                                                                                              | ✅ yes |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ❌ ?   |
| `--description string`        | The description of the proposal, it can be a url, some text, etc.                                                                                           | ✅ yes |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ✅ yes |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ✅ yes |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for remove-whitelisted-permission                                                                                                                      | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default 'home' directory will be used                                                                         | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ❌ ?   |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ✅ yes |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `--title string`              | The title of the proposal.                                                                                                                                  | ✅ yes |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx customgov proposal account remove-whitelisted-permission --help
Create a proposal to remove a whitelisted permission from an address.

Usage:
  sekaid tx customgov proposal account remove-whitelisted-permission [permission_id] [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
      --addr string              the address to set permissions
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --description string       The description of the proposal, it can be a url, some text, etc.
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for remove-whitelisted-permission
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
      --title string             The title of the proposal.
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```



```
sekaid tx customgov proposal account remove-whitelisted-permission 12 --from=kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --addr=kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt --home=/root/.sekai --chain-id=localnet-4 --fees=10ukex --keyring-backend=test --description="Test proposal for docs" --title="Test proposal" --yes -o json | jq
{
  "height": "0",
  "txhash": "297D3955DE3C620FA1670BDCFCC567E7EB323D9C300E307E9BED53625378020F",
  "codespace": "",
  "code": 0,
  "data": "",
  "raw_log": "[]",
  "logs": [],
  "info": "",
  "gas_wanted": "0",
  "gas_used": "0",
  "tx": null,
  "timestamp": "",
  "events": []
}
```

<details>
  <summary>Check tx execution</summary>

  ```
  sekaid q tx 297D3955DE3C620FA1670BDCFCC567E7EB323D9C300E307E9BED53625378020F -o json | jq
  ```

  ```json
  {
    "height": "142021",
    "txhash": "297D3955DE3C620FA1670BDCFCC567E7EB323D9C300E307E9BED53625378020F",
    "codespace": "",
    "code": 0,
    "data": "0A210A1B2F6B6972612E676F762E4D73675375626D697450726F706F73616C12020812",
    "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/kira.gov.MsgSubmitProposal\"}]},{\"type\":\"submit_proposal\",\"attributes\":[{\"key\":\"proposal_id\",\"value\":\"18\"},{\"key\":\"proposal_type\",\"value\":\"RemoveWhitelistedAccountPermission\"},{\"key\":\"proposal_content\",\"value\":\"proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\\ntitle: Test proposal\\ndescription: Test proposal for docs\\ncontent:\\n  typeurl: /kira.gov.RemoveWhitelistedAccountPermissionProposal\\n  value:\\n  - 10\\n  - 20\\n  - 247\\n  - 114\\n  - 3\\n  - 50\\n  - 223\\n  - 138\\n  - 35\\n  - 224\\n  - 167\\n  - 8\\n  - 117\\n  - 128\\n  - 102\\n  - 88\\n  - 115\\n  - 212\\n  - 100\\n  - 159\\n  - 201\\n  - 44\\n  - 16\\n  - 12\\n  xxx_nounkeyedliteral: {}\\n  xxx_unrecognized: []\\n  xxx_sizecache: 0\\n\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/kira.gov.MsgSubmitProposal"
              }
            ]
          },
          {
            "type": "submit_proposal",
            "attributes": [
              {
                "key": "proposal_id",
                "value": "18"
              },
              {
                "key": "proposal_type",
                "value": "RemoveWhitelistedAccountPermission"
              },
              {
                "key": "proposal_content",
                "value": "proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\ntitle: Test proposal\ndescription: Test proposal for docs\ncontent:\n  typeurl: /kira.gov.RemoveWhitelistedAccountPermissionProposal\n  value:\n  - 10\n  - 20\n  - 247\n  - 114\n  - 3\n  - 50\n  - 223\n  - 138\n  - 35\n  - 224\n  - 167\n  - 8\n  - 117\n  - 128\n  - 102\n  - 88\n  - 115\n  - 212\n  - 100\n  - 159\n  - 201\n  - 44\n  - 16\n  - 12\n  xxx_nounkeyedliteral: {}\n  xxx_unrecognized: []\n  xxx_sizecache: 0\n"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "0",
    "gas_used": "0",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/kira.gov.MsgSubmitProposal",
            "proposer": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
            "title": "Test proposal",
            "description": "Test proposal for docs",
            "content": {
              "@type": "/kira.gov.RemoveWhitelistedAccountPermissionProposal",
              "address": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
              "permission": "PERMISSION_CREATE_SET_NETWORK_PROPERTY_PROPOSAL"
            }
          }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
      },
      "auth_info": {
        "signer_infos": [
          {
            "public_key": {
              "@type": "/cosmos.crypto.secp256k1.PubKey",
              "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "103"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "ukex",
              "amount": "10"
            }
          ],
          "gas_limit": "200000",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
        "wrJE4bPMqRMJ2AaUEfRGgILw2j2ZtDEFUt5v9WjjyJBqfpQ1sgdw0zOYaJhVLtEnITtj/stl0lwUGkyyqTmeYQ=="
      ]
    },
    "timestamp": "2023-06-12T11:19:14Z",
    "events": [
      {
        "type": "tx",
        "attributes": [
          {
            "key": "YWNjX3NlcQ==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NC8xMDM=",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "d3JKRTRiUE1xUk1KMkFhVUVmUkdnSUx3MmoyWnRERUZVdDV2OVdqanlKQnFmcFExc2dkdzB6T1lhSmhWTHRFbklUdGovc3RsMGx3VUdreXlxVG1lWVE9PQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTB1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTB1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTB1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "ZmVl",
            "value": "MTB1a2V4",
            "index": true
          },
          {
            "key": "ZmVlX3BheWVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "L2tpcmEuZ292Lk1zZ1N1Ym1pdFByb3Bvc2Fs",
            "index": true
          }
        ]
      },
      {
        "type": "submit_proposal",
        "attributes": [
          {
            "key": "cHJvcG9zYWxfaWQ=",
            "value": "MTg=",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfdHlwZQ==",
            "value": "UmVtb3ZlV2hpdGVsaXN0ZWRBY2NvdW50UGVybWlzc2lvbg==",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfY29udGVudA==",
            "value": "cHJvcG9zZXI6IGtpcmExdm13ZGd3NDI2YWo5ZngzM2ZxdXNtdGc2cjY1eXl1Y214NnJkdDQKdGl0bGU6IFRlc3QgcHJvcG9zYWwKZGVzY3JpcHRpb246IFRlc3QgcHJvcG9zYWwgZm9yIGRvY3MKY29udGVudDoKICB0eXBldXJsOiAva2lyYS5nb3YuUmVtb3ZlV2hpdGVsaXN0ZWRBY2NvdW50UGVybWlzc2lvblByb3Bvc2FsCiAgdmFsdWU6CiAgLSAxMAogIC0gMjAKICAtIDI0NwogIC0gMTE0CiAgLSAzCiAgLSA1MAogIC0gMjIzCiAgLSAxMzgKICAtIDM1CiAgLSAyMjQKICAtIDE2NwogIC0gOAogIC0gMTE3CiAgLSAxMjgKICAtIDEwMgogIC0gODgKICAtIDExNQogIC0gMjEyCiAgLSAxMDAKICAtIDE1OQogIC0gMjAxCiAgLSA0NAogIC0gMTYKICAtIDEyCiAgeHh4X25vdW5rZXllZGxpdGVyYWw6IHt9CiAgeHh4X3VucmVjb2duaXplZDogW10KICB4eHhfc2l6ZWNhY2hlOiAwCg==",
            "index": true
          }
        ]
      }
    ]
  }
  ```
</details>

<details>
  <summary>Check proposal</summary>

  ```
  sekaid q customgov proposal 18 -o json | jq
  ```

  ```json
  {
    "proposal_id": "18",
    "title": "Test proposal",
    "description": "Test proposal for docs",
    "content": {
      "@type": "/kira.gov.RemoveWhitelistedAccountPermissionProposal",
      "address": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
      "permission": "PERMISSION_CREATE_SET_NETWORK_PROPERTY_PROPOSAL"
    },
    "submit_time": "2023-06-12T11:19:14.776196895Z",
    "voting_end_time": "2023-06-12T11:25:14.776196895Z",
    "enactment_end_time": "2023-06-12T11:30:14.776196895Z",
    "min_voting_end_block_height": "142023",
    "min_enactment_end_block_height": "142057",
    "result": "VOTE_RESULT_QUORUM_NOT_REACHED",
    "exec_result": ""
  }
  ```
</details>

[Return to top](#sekai)

###### 21.7.7.1.5 unassign-role

Create a proposal to unassign a role from an address.

Usage:
```
sekaid tx customgov proposal account unassign-role role [flags] 🟨 Can use role_id | role_sid
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `--addr string`               | the address to set permissions                                                                                                                              | ✅ yes |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ❌ ?   |
| `--description string`        | The description of the proposal, it can be a url, some text, etc.                                                                                           | ✅ yes |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ✅ yes |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ✅ yes |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for unassign-role                                                                                                                                      | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default 'home' directory will be used                                                                         | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ❌ ?   |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ✅ yes |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `--title string`              | The title of the proposal.                                                                                                                                  | ✅ yes |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx customgov proposal account unassign-role --help
Create a proposal to unassign a role from an address.

Usage:
  sekaid tx customgov proposal account unassign-role role [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
      --addr string              the address to set permissions
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --description string       The description of the proposal, it can be a url, some text, etc.
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for unassign-role
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
      --title string             The title of the proposal.
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid tx customgov proposal account unassign-role validator --from=kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --addr=kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt --home=/root/.sekai --chain-id=localnet-4 --fees=10ukex --keyring-backend=test --description="Test proposal for docs" --title="Test proposal" --yes -o json | jq
{
  "height": "0",
  "txhash": "295F7574ADFD68B5D039F6E3E019B9B24897A443F2861088291194A6EF569D0D",
  "codespace": "",
  "code": 0,
  "data": "",
  "raw_log": "[]",
  "logs": [],
  "info": "",
  "gas_wanted": "0",
  "gas_used": "0",
  "tx": null,
  "timestamp": "",
  "events": []
}
```

<details>
  <summary>Check tx execution</summary>

  ```
  sekaid q tx 295F7574ADFD68B5D039F6E3E019B9B24897A443F2861088291194A6EF569D0D -o json | jq
  ```

  ```json
  {
    "height": "142120",
    "txhash": "295F7574ADFD68B5D039F6E3E019B9B24897A443F2861088291194A6EF569D0D",
    "codespace": "",
    "code": 0,
    "data": "0A210A1B2F6B6972612E676F762E4D73675375626D697450726F706F73616C12020813",
    "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/kira.gov.MsgSubmitProposal\"}]},{\"type\":\"submit_proposal\",\"attributes\":[{\"key\":\"proposal_id\",\"value\":\"19\"},{\"key\":\"proposal_type\",\"value\":\"UnassignRoleFromAccount\"},{\"key\":\"proposal_content\",\"value\":\"proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\\ntitle: Test proposal\\ndescription: Test proposal for docs\\ncontent:\\n  typeurl: /kira.gov.UnassignRoleFromAccountProposal\\n  value:\\n  - 10\\n  - 20\\n  - 247\\n  - 114\\n  - 3\\n  - 50\\n  - 223\\n  - 138\\n  - 35\\n  - 224\\n  - 167\\n  - 8\\n  - 117\\n  - 128\\n  - 102\\n  - 88\\n  - 115\\n  - 212\\n  - 100\\n  - 159\\n  - 201\\n  - 44\\n  - 18\\n  - 9\\n  - 118\\n  - 97\\n  - 108\\n  - 105\\n  - 100\\n  - 97\\n  - 116\\n  - 111\\n  - 114\\n  xxx_nounkeyedliteral: {}\\n  xxx_unrecognized: []\\n  xxx_sizecache: 0\\n\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/kira.gov.MsgSubmitProposal"
              }
            ]
          },
          {
            "type": "submit_proposal",
            "attributes": [
              {
                "key": "proposal_id",
                "value": "19"
              },
              {
                "key": "proposal_type",
                "value": "UnassignRoleFromAccount"
              },
              {
                "key": "proposal_content",
                "value": "proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\ntitle: Test proposal\ndescription: Test proposal for docs\ncontent:\n  typeurl: /kira.gov.UnassignRoleFromAccountProposal\n  value:\n  - 10\n  - 20\n  - 247\n  - 114\n  - 3\n  - 50\n  - 223\n  - 138\n  - 35\n  - 224\n  - 167\n  - 8\n  - 117\n  - 128\n  - 102\n  - 88\n  - 115\n  - 212\n  - 100\n  - 159\n  - 201\n  - 44\n  - 18\n  - 9\n  - 118\n  - 97\n  - 108\n  - 105\n  - 100\n  - 97\n  - 116\n  - 111\n  - 114\n  xxx_nounkeyedliteral: {}\n  xxx_unrecognized: []\n  xxx_sizecache: 0\n"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "0",
    "gas_used": "0",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/kira.gov.MsgSubmitProposal",
            "proposer": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
            "title": "Test proposal",
            "description": "Test proposal for docs",
            "content": {
              "@type": "/kira.gov.UnassignRoleFromAccountProposal",
              "address": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
              "role_identifier": "validator"
            }
          }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
      },
      "auth_info": {
        "signer_infos": [
          {
            "public_key": {
              "@type": "/cosmos.crypto.secp256k1.PubKey",
              "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "105"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "ukex",
              "amount": "10"
            }
          ],
          "gas_limit": "200000",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
        "dqu/YESslTFwwXdRrfpACFPOzdIct5sq3qpE43y07kgppsr5mpJQmstLpp1OHOR5MznU3I9QvplooRS/YUqA8A=="
      ]
    },
    "timestamp": "2023-06-12T11:36:15Z",
    "events": [
      {
        "type": "tx",
        "attributes": [
          {
            "key": "YWNjX3NlcQ==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NC8xMDU=",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "ZHF1L1lFU3NsVEZ3d1hkUnJmcEFDRlBPemRJY3Q1c3EzcXBFNDN5MDdrZ3Bwc3I1bXBKUW1zdExwcDFPSE9SNU16blUzSTlRdnBsb29SUy9ZVXFBOEE9PQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTB1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTB1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTB1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "ZmVl",
            "value": "MTB1a2V4",
            "index": true
          },
          {
            "key": "ZmVlX3BheWVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "L2tpcmEuZ292Lk1zZ1N1Ym1pdFByb3Bvc2Fs",
            "index": true
          }
        ]
      },
      {
        "type": "submit_proposal",
        "attributes": [
          {
            "key": "cHJvcG9zYWxfaWQ=",
            "value": "MTk=",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfdHlwZQ==",
            "value": "VW5hc3NpZ25Sb2xlRnJvbUFjY291bnQ=",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfY29udGVudA==",
            "value": "cHJvcG9zZXI6IGtpcmExdm13ZGd3NDI2YWo5ZngzM2ZxdXNtdGc2cjY1eXl1Y214NnJkdDQKdGl0bGU6IFRlc3QgcHJvcG9zYWwKZGVzY3JpcHRpb246IFRlc3QgcHJvcG9zYWwgZm9yIGRvY3MKY29udGVudDoKICB0eXBldXJsOiAva2lyYS5nb3YuVW5hc3NpZ25Sb2xlRnJvbUFjY291bnRQcm9wb3NhbAogIHZhbHVlOgogIC0gMTAKICAtIDIwCiAgLSAyNDcKICAtIDExNAogIC0gMwogIC0gNTAKICAtIDIyMwogIC0gMTM4CiAgLSAzNQogIC0gMjI0CiAgLSAxNjcKICAtIDgKICAtIDExNwogIC0gMTI4CiAgLSAxMDIKICAtIDg4CiAgLSAxMTUKICAtIDIxMgogIC0gMTAwCiAgLSAxNTkKICAtIDIwMQogIC0gNDQKICAtIDE4CiAgLSA5CiAgLSAxMTgKICAtIDk3CiAgLSAxMDgKICAtIDEwNQogIC0gMTAwCiAgLSA5NwogIC0gMTE2CiAgLSAxMTEKICAtIDExNAogIHh4eF9ub3Vua2V5ZWRsaXRlcmFsOiB7fQogIHh4eF91bnJlY29nbml6ZWQ6IFtdCiAgeHh4X3NpemVjYWNoZTogMAo=",
            "index": true
          }
        ]
      }
    ]
  }
  ```
</details>

<details>
  <summary>Check proposal</summary>

  ```
  sekaid q customgov proposal 19 -o json | jq
  ```

  ```json
  {
    "proposal_id": "19",
    "title": "Test proposal",
    "description": "Test proposal for docs",
    "content": {
      "@type": "/kira.gov.UnassignRoleFromAccountProposal",
      "address": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
      "role_identifier": "validator"
    },
    "submit_time": "2023-06-12T11:36:15.728838159Z",
    "voting_end_time": "2023-06-12T11:42:15.728838159Z",
    "enactment_end_time": "2023-06-12T11:47:15.728838159Z",
    "min_voting_end_block_height": "142122",
    "min_enactment_end_block_height": "142123",
    "result": "VOTE_PENDING",
    "exec_result": ""
  }
  ```
</details>

[Return to top](#sekai)

###### 21.7.7.1.6 whitelist-permission

Create a proposal to whitelist a permission to an address.

Usage:
```
sekaid tx customgov proposal account whitelist-permission [permission_id] [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `--addr string`               | the address to set permissions                                                                                                                              | ✅ yes |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ❌ ?   |
| `--description string`        | The description of the proposal, it can be a url, some text, etc.                                                                                           | ✅ yes |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ✅ yes |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ✅ yes |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for whitelist-permission                                                                                                                               | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default 'home' directory will be used                                                                         | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ❌ ?   |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ✅ yes |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `--title string`              | The title of the proposal.                                                                                                                                  | ✅ yes |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx customgov proposal account whitelist-permission --help
Create a proposal to whitelist a permission to an address.

Usage:
  sekaid tx customgov proposal account whitelist-permission [permission_id] [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
      --addr string              the address to set permissions
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --description string       The description of the proposal, it can be a url, some text, etc.
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for whitelist-permission
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
      --title string             The title of the proposal.
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

> **Can't whitelist blacklisted permission**

```
sekaid tx customgov proposal account whitelist-permission 14 --from=kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --addr=kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt --home=/root/.sekai --chain-id=localnet-4 --fees=10ukex --keyring-backend=test --description="Test proposal for docs" --title="Test proposal" --yes -o json | jq
{
  "height": "0",
  "txhash": "2733898301260774DAA8F962B8D7D680D01BA79C9AF7748EAE9906352BC62102",
  "codespace": "",
  "code": 0,
  "data": "",
  "raw_log": "[]",
  "logs": [],
  "info": "",
  "gas_wanted": "0",
  "gas_used": "0",
  "tx": null,
  "timestamp": "",
  "events": []
}
```

<details>
  <summary>Check tx execution</summary>

  ```
  sekaid q tx 2733898301260774DAA8F962B8D7D680D01BA79C9AF7748EAE9906352BC62102 -o json | jq
  ```

  ```json
  {
    "height": "142200",
    "txhash": "2733898301260774DAA8F962B8D7D680D01BA79C9AF7748EAE9906352BC62102",
    "codespace": "",
    "code": 0,
    "data": "0A210A1B2F6B6972612E676F762E4D73675375626D697450726F706F73616C12020814",
    "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/kira.gov.MsgSubmitProposal\"}]},{\"type\":\"submit_proposal\",\"attributes\":[{\"key\":\"proposal_id\",\"value\":\"20\"},{\"key\":\"proposal_type\",\"value\":\"WhitelistAccountPermission\"},{\"key\":\"proposal_content\",\"value\":\"proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\\ntitle: Test proposal\\ndescription: Test proposal for docs\\ncontent:\\n  typeurl: /kira.gov.WhitelistAccountPermissionProposal\\n  value:\\n  - 10\\n  - 20\\n  - 247\\n  - 114\\n  - 3\\n  - 50\\n  - 223\\n  - 138\\n  - 35\\n  - 224\\n  - 167\\n  - 8\\n  - 117\\n  - 128\\n  - 102\\n  - 88\\n  - 115\\n  - 212\\n  - 100\\n  - 159\\n  - 201\\n  - 44\\n  - 16\\n  - 14\\n  xxx_nounkeyedliteral: {}\\n  xxx_unrecognized: []\\n  xxx_sizecache: 0\\n\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/kira.gov.MsgSubmitProposal"
              }
            ]
          },
          {
            "type": "submit_proposal",
            "attributes": [
              {
                "key": "proposal_id",
                "value": "20"
              },
              {
                "key": "proposal_type",
                "value": "WhitelistAccountPermission"
              },
              {
                "key": "proposal_content",
                "value": "proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\ntitle: Test proposal\ndescription: Test proposal for docs\ncontent:\n  typeurl: /kira.gov.WhitelistAccountPermissionProposal\n  value:\n  - 10\n  - 20\n  - 247\n  - 114\n  - 3\n  - 50\n  - 223\n  - 138\n  - 35\n  - 224\n  - 167\n  - 8\n  - 117\n  - 128\n  - 102\n  - 88\n  - 115\n  - 212\n  - 100\n  - 159\n  - 201\n  - 44\n  - 16\n  - 14\n  xxx_nounkeyedliteral: {}\n  xxx_unrecognized: []\n  xxx_sizecache: 0\n"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "0",
    "gas_used": "0",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/kira.gov.MsgSubmitProposal",
            "proposer": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
            "title": "Test proposal",
            "description": "Test proposal for docs",
            "content": {
              "@type": "/kira.gov.WhitelistAccountPermissionProposal",
              "address": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
              "permission": "PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL"
            }
          }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
      },
      "auth_info": {
        "signer_infos": [
          {
            "public_key": {
              "@type": "/cosmos.crypto.secp256k1.PubKey",
              "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "107"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "ukex",
              "amount": "10"
            }
          ],
          "gas_limit": "200000",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
        "a6/6w6RjWLA7JC4Kkup5WK4rZvl7mggP4Nmxxc/P/04rQMBwBfeu0dno4b8P2luG6DTg8rfXQfoQ5OCEcSTy/A=="
      ]
    },
    "timestamp": "2023-06-12T11:50:00Z",
    "events": [
      {
        "type": "tx",
        "attributes": [
          {
            "key": "YWNjX3NlcQ==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NC8xMDc=",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "YTYvNnc2UmpXTEE3SkM0S2t1cDVXSzRyWnZsN21nZ1A0Tm14eGMvUC8wNHJRTUJ3QmZldTBkbm80YjhQMmx1RzZEVGc4cmZYUWZvUTVPQ0VjU1R5L0E9PQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTB1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTB1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTB1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "ZmVl",
            "value": "MTB1a2V4",
            "index": true
          },
          {
            "key": "ZmVlX3BheWVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "L2tpcmEuZ292Lk1zZ1N1Ym1pdFByb3Bvc2Fs",
            "index": true
          }
        ]
      },
      {
        "type": "submit_proposal",
        "attributes": [
          {
            "key": "cHJvcG9zYWxfaWQ=",
            "value": "MjA=",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfdHlwZQ==",
            "value": "V2hpdGVsaXN0QWNjb3VudFBlcm1pc3Npb24=",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfY29udGVudA==",
            "value": "cHJvcG9zZXI6IGtpcmExdm13ZGd3NDI2YWo5ZngzM2ZxdXNtdGc2cjY1eXl1Y214NnJkdDQKdGl0bGU6IFRlc3QgcHJvcG9zYWwKZGVzY3JpcHRpb246IFRlc3QgcHJvcG9zYWwgZm9yIGRvY3MKY29udGVudDoKICB0eXBldXJsOiAva2lyYS5nb3YuV2hpdGVsaXN0QWNjb3VudFBlcm1pc3Npb25Qcm9wb3NhbAogIHZhbHVlOgogIC0gMTAKICAtIDIwCiAgLSAyNDcKICAtIDExNAogIC0gMwogIC0gNTAKICAtIDIyMwogIC0gMTM4CiAgLSAzNQogIC0gMjI0CiAgLSAxNjcKICAtIDgKICAtIDExNwogIC0gMTI4CiAgLSAxMDIKICAtIDg4CiAgLSAxMTUKICAtIDIxMgogIC0gMTAwCiAgLSAxNTkKICAtIDIwMQogIC0gNDQKICAtIDE2CiAgLSAxNAogIHh4eF9ub3Vua2V5ZWRsaXRlcmFsOiB7fQogIHh4eF91bnJlY29nbml6ZWQ6IFtdCiAgeHh4X3NpemVjYWNoZTogMAo=",
            "index": true
          }
        ]
      }
    ]
  }
  ```
</details>

<details>
  <summary>Check proposal</summary>

  ```
  sekaid q customgov proposal 20 -o json | jq
  ```

  ```json
  {
    "proposal_id": "20",
    "title": "Test proposal",
    "description": "Test proposal for docs",
    "content": {
      "@type": "/kira.gov.WhitelistAccountPermissionProposal",
      "address": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
      "permission": "PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL"
    },
    "submit_time": "2023-06-12T11:50:00.816857099Z",
    "voting_end_time": "2023-06-12T11:56:00.816857099Z",
    "enactment_end_time": "2023-06-12T12:01:00.816857099Z",
    "min_voting_end_block_height": "142202",
    "min_enactment_end_block_height": "142236",
    "result": "VOTE_RESULT_QUORUM_NOT_REACHED",
    "exec_result": ""
  }
  ```
</details>

[Return to top](#sekai)

###### 21.7.7.2 proposal-jail-councilor

Create a proposal to jail councilors.

Usage:
```
sekaid tx customgov proposal proposal-jail-councilor [councilors] [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ❌ ?   |
| `--description string`        | The description of the proposal, it can be a url, some text, etc.                                                                                           | ✅ yes |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ✅ yes |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ✅ yes |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for proposal-jail-councilor                                                                                                                            | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default 'home' directory will be used                                                                         | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ❌ ?   |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ✅ yes |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `--title string`              | The title of the proposal.                                                                                                                                  | ✅ yes |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx customgov proposal proposal-jail-councilor --help
Create a proposal to jail councilors

Usage:
  sekaid tx customgov proposal proposal-jail-councilor [councilors] [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --description string       The description of the proposal, it can be a url, some text, etc.
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for proposal-jail-councilor
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
      --title string             The title of the proposal.
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

#TODO haven't tested

🟨  
🟨  
🟨  

[Return to top](#sekai)

###### 21.7.7.3 proposal-reset-whole-councilor-rank

Create a proposal to reset whole councilor rank.

Usage:
```
sekaid tx customgov proposal proposal-reset-whole-councilor-rank [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ❌ ?   |
| `--description string`        | The description of the proposal, it can be a url, some text, etc.                                                                                           | ✅ yes |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ✅ yes |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ✅ yes |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for proposal-reset-whole-councilor-rank                                                                                                                | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default 'home' directory will be used                                                                         | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ❌ ?   |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ✅ yes |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `--title string`              | The title of the proposal.                                                                                                                                  | ✅ yes |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx customgov proposal proposal-reset-whole-councilor-rank --help
Create a proposal to reset whole councilor rank

Usage:
  sekaid tx customgov proposal proposal-reset-whole-councilor-rank [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --description string       The description of the proposal, it can be a url, some text, etc.
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for proposal-reset-whole-councilor-rank
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
      --title string             The title of the proposal.
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

#TODO haven't tested

🟨  
🟨  
🟨  

[Return to top](#sekai)

###### 21.7.7.4 role

Role proposals management subcommands.

Usage:
```
sekaid tx customgov proposal role [flags]
sekaid tx customgov proposal role [command]
```

Available Commands:

| Subcommands                         | Description                                                               |
| ----------------------------------- | ------------------------------------------------------------------------- |
| [`blacklist-permission`]()          | Raise governance proposal to blacklist a permission for a role.           |
| [`create`]()                        | Raise governance proposal to create a new role.                           |
| [`remove`]()                        | Raise governance proposal to remove a role.                               |
| [`remove-blacklisted-permission`]() | Raise governance proposal to remove a blacklisted permission from a role. |
| [`remove-whitelisted-permission`]() | Raise governance proposal to remove whitelisted permission from a role.   |
| [`whitelist-permission`]()          | Raise governance proposal to whitelist a permission for a role.           |



| Flags        | Description   | Work |
| ------------ | ------------- | ---- |
| `-h, --help` | help for role |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx customgov proposal role --help
Role proposals management subcommands

Usage:
  sekaid tx customgov proposal role [flags]
  sekaid tx customgov proposal role [command]

Available Commands:
  blacklist-permission          Raise governance proposal to blacklist a permission for a role.
  create                        Raise governance proposal to create a new role.
  remove                        Raise governance proposal to remove a role.
  remove-blacklisted-permission Raise governance proposal to remove a blacklisted permission from a role.
  remove-whitelisted-permission Raise governance proposal to remove whitelisted permission from a role.
  whitelist-permission          Raise governance proposal to whitelist a permission for a role.

Flags:
  -h, --help   help for role

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors

Use "sekaid tx customgov proposal role [command] --help" for more information about a command.
```

[Return to top](#sekai)

###### 21.7.7.4.1 blacklist-permission

Raise governance proposal to blacklist a permission for a role.

Usage:
```
sekaid tx customgov proposal role blacklist-permission [role_sid] [role_description] [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `--blacklist int32Slice`      | the blacklist values in format 1,2,3 (default `[]`)                                                                                                         | ❌ ?   |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ❌ ?   |
| `--description string`        | The description of the proposal, it can be a url, some text, etc.                                                                                           | ✅ yes |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ✅ yes |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ✅ yes |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for blacklist-permission                                                                                                                               | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default 'home' directory will be used                                                                         | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ❌ ?   |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ✅ yes |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `--title string`              | The title of the proposal.                                                                                                                                  | ✅ yes |
| `--whitelist int32Slice`      | the whitelist value in format 1,2,3 (default `[]`)                                                                                                          | ❌ ?   |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx customgov proposal role blacklist-permission --help
Raise governance proposal to blacklist a permission for a role.

Usage:
  sekaid tx customgov proposal role blacklist-permission [role_sid] [role_description] [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
      --blacklist int32Slice     the blacklist values in format 1,2,3 (default [])
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --description string       The description of the proposal, it can be a url, some text, etc.
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for blacklist-permission
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
      --title string             The title of the proposal.
      --whitelist int32Slice     the whitelist value in format 1,2,3 (default [])
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid tx customgov proposal role blacklist-permission 3 11 --title="Test proposal" --description="Test proposal" --from=kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --keyring-backend=test --home=/root/.sekai --chain-id=localnet-4 --fees=10ukex --yes --output=json | jq
{
  "height": "0",
  "txhash": "854CC53C8A5860E30E6C9EE4A101E3BAE5F3AF9FD13A023965006ABF6953164B",
  "codespace": "",
  "code": 0,
  "data": "",
  "raw_log": "[]",
  "logs": [],
  "info": "",
  "gas_wanted": "0",
  "gas_used": "0",
  "tx": null,
  "timestamp": "",
  "events": []
}
```

<details>
  <summary>Check tx execution</summary>

  ```
  sekaid q tx 854CC53C8A5860E30E6C9EE4A101E3BAE5F3AF9FD13A023965006ABF6953164B -o json | jq
  ```

  ```json
  {
    "height": "142786",
    "txhash": "854CC53C8A5860E30E6C9EE4A101E3BAE5F3AF9FD13A023965006ABF6953164B",
    "codespace": "",
    "code": 0,
    "data": "0A210A1B2F6B6972612E676F762E4D73675375626D697450726F706F73616C12020815",
    "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/kira.gov.MsgSubmitProposal\"}]},{\"type\":\"submit_proposal\",\"attributes\":[{\"key\":\"proposal_id\",\"value\":\"21\"},{\"key\":\"proposal_type\",\"value\":\"BlacklistRolePermission\"},{\"key\":\"proposal_content\",\"value\":\"proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\\ntitle: Test proposal\\ndescription: Test proposal\\ncontent:\\n  typeurl: /kira.gov.BlacklistRolePermissionProposal\\n  value:\\n  - 10\\n  - 1\\n  - 51\\n  - 16\\n  - 11\\n  xxx_nounkeyedliteral: {}\\n  xxx_unrecognized: []\\n  xxx_sizecache: 0\\n\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/kira.gov.MsgSubmitProposal"
              }
            ]
          },
          {
            "type": "submit_proposal",
            "attributes": [
              {
                "key": "proposal_id",
                "value": "21"
              },
              {
                "key": "proposal_type",
                "value": "BlacklistRolePermission"
              },
              {
                "key": "proposal_content",
                "value": "proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\ntitle: Test proposal\ndescription: Test proposal\ncontent:\n  typeurl: /kira.gov.BlacklistRolePermissionProposal\n  value:\n  - 10\n  - 1\n  - 51\n  - 16\n  - 11\n  xxx_nounkeyedliteral: {}\n  xxx_unrecognized: []\n  xxx_sizecache: 0\n"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "0",
    "gas_used": "0",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/kira.gov.MsgSubmitProposal",
            "proposer": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
            "title": "Test proposal",
            "description": "Test proposal",
            "content": {
              "@type": "/kira.gov.BlacklistRolePermissionProposal",
              "role_identifier": "3",
              "permission": "PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL"
            }
          }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
      },
      "auth_info": {
        "signer_infos": [
          {
            "public_key": {
              "@type": "/cosmos.crypto.secp256k1.PubKey",
              "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "108"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "ukex",
              "amount": "10"
            }
          ],
          "gas_limit": "200000",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
        "x3IfoIjy+d3Qzfl9SjiAgOoaMpY4mLixhc7iIp1jMJE4iM4JNA63pzXB4XfuFN+z1Q3Y5DSrUBkxRHhgdfwxMg=="
      ]
    },
    "timestamp": "2023-06-12T13:30:45Z",
    "events": [
      {
        "type": "tx",
        "attributes": [
          {
            "key": "YWNjX3NlcQ==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NC8xMDg=",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "eDNJZm9JankrZDNRemZsOVNqaUFnT29hTXBZNG1MaXhoYzdpSXAxak1KRTRpTTRKTkE2M3B6WEI0WGZ1Rk4rejFRM1k1RFNyVUJreFJIaGdkZnd4TWc9PQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTB1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTB1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTB1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "ZmVl",
            "value": "MTB1a2V4",
            "index": true
          },
          {
            "key": "ZmVlX3BheWVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "L2tpcmEuZ292Lk1zZ1N1Ym1pdFByb3Bvc2Fs",
            "index": true
          }
        ]
      },
      {
        "type": "submit_proposal",
        "attributes": [
          {
            "key": "cHJvcG9zYWxfaWQ=",
            "value": "MjE=",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfdHlwZQ==",
            "value": "QmxhY2tsaXN0Um9sZVBlcm1pc3Npb24=",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfY29udGVudA==",
            "value": "cHJvcG9zZXI6IGtpcmExdm13ZGd3NDI2YWo5ZngzM2ZxdXNtdGc2cjY1eXl1Y214NnJkdDQKdGl0bGU6IFRlc3QgcHJvcG9zYWwKZGVzY3JpcHRpb246IFRlc3QgcHJvcG9zYWwKY29udGVudDoKICB0eXBldXJsOiAva2lyYS5nb3YuQmxhY2tsaXN0Um9sZVBlcm1pc3Npb25Qcm9wb3NhbAogIHZhbHVlOgogIC0gMTAKICAtIDEKICAtIDUxCiAgLSAxNgogIC0gMTEKICB4eHhfbm91bmtleWVkbGl0ZXJhbDoge30KICB4eHhfdW5yZWNvZ25pemVkOiBbXQogIHh4eF9zaXplY2FjaGU6IDAK",
            "index": true
          }
        ]
      }
    ]
  }
  ```
</details>

<details>
  <summary>Check proposal</summary>

  ```
  sekaid q customgov proposal 21 -o json | jq
  ```

  ```json
  {
    "proposal_id": "21",
    "title": "Test proposal",
    "description": "Test proposal",
    "content": {
      "@type": "/kira.gov.BlacklistRolePermissionProposal",
      "role_identifier": "3",
      "permission": "PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL"
    },
    "submit_time": "2023-06-12T13:30:45.511453097Z",
    "voting_end_time": "2023-06-12T13:36:45.511453097Z",
    "enactment_end_time": "2023-06-12T13:41:45.511453097Z",
    "min_voting_end_block_height": "142788",
    "min_enactment_end_block_height": "142789",
    "result": "VOTE_PENDING",
    "exec_result": ""
  }
  ```
</details>

[Return to top](#sekai)

###### 21.7.7.4.2 create

Raise governance proposal to create a new role.

Usage:
```
sekaid tx customgov proposal role create [role_sid] [role_description] [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `--blacklist int32Slice`      | the blacklist values in format 1,2,3 (default `[]`)                                                                                                         | ✅ yes |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ❌ ?   |
| `--description string`        | The description of the proposal, it can be a url, some text, etc.                                                                                           | ✅ yes |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ✅ yes |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ✅ yes |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for create                                                                                                                                             | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default 'home' directory will be used                                                                         | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ❌ ?   |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ✅ yes |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `--title string`              | The title of the proposal.                                                                                                                                  | ✅ yes |
| `--whitelist int32Slice`      | the whitelist value in format 1,2,3 (default `[]`)                                                                                                          | ✅ yes |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx customgov proposal role create --help
Raise governance proposal to create a new role.

Usage:
  sekaid tx customgov proposal role create [role_sid] [role_description] [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
      --blacklist int32Slice     the blacklist values in format 1,2,3 (default [])
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --description string       The description of the proposal, it can be a url, some text, etc.
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for create
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
      --title string             The title of the proposal.
      --whitelist int32Slice     the whitelist value in format 1,2,3 (default [])
  -y, --yes                      Skip tx broadcasting prompt confirmation
```

```
sekaid tx customgov proposal role create 'bruh3' "'Bruh'(3) role" --title="Test proposal for 'bruh3' role" --description="Test proposal for creating 'bruh3' role" --from=kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --whitelist=20,21 --blacklist=22,23 --keyring-backend=test --chain-id=localnet-4 --home=/root/.sekai --fees=100ukex --broadcast-mode=async --yes --output=json | jq
{
  "height": "0",
  "txhash": "05D7C2595B3AFD249CCB3A9FE2040B1B6AFE2CBAB72465771790EF19EA047F38",
  "codespace": "",
  "code": 0,
  "data": "",
  "raw_log": "",
  "logs": [],
  "info": "",
  "gas_wanted": "0",
  "gas_used": "0",
  "tx": null,
  "timestamp": "",
  "events": []
}
```

<details>
  <summary>Check tx execution</summary>

  ```
  sekaid q tx 05D7C2595B3AFD249CCB3A9FE2040B1B6AFE2CBAB72465771790EF19EA047F38 -o json | jq
  ```

  ```json
  {
    "height": "142853",
    "txhash": "05D7C2595B3AFD249CCB3A9FE2040B1B6AFE2CBAB72465771790EF19EA047F38",
    "codespace": "",
    "code": 0,
    "data": "0A210A1B2F6B6972612E676F762E4D73675375626D697450726F706F73616C12020816",
    "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/kira.gov.MsgSubmitProposal\"}]},{\"type\":\"submit_proposal\",\"attributes\":[{\"key\":\"proposal_id\",\"value\":\"22\"},{\"key\":\"proposal_type\",\"value\":\"CreateRoleProposal\"},{\"key\":\"proposal_content\",\"value\":\"proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\\ntitle: Test proposal for 'bruh3' role\\ndescription: Test proposal for creating 'bruh3' role\\ncontent:\\n  typeurl: /kira.gov.CreateRoleProposal\\n  value:\\n  - 10\\n  - 5\\n  - 98\\n  - 114\\n  - 117\\n  - 104\\n  - 51\\n  - 18\\n  - 14\\n  - 39\\n  - 66\\n  - 114\\n  - 117\\n  - 104\\n  - 39\\n  - 40\\n  - 51\\n  - 41\\n  - 32\\n  - 114\\n  - 111\\n  - 108\\n  - 101\\n  - 26\\n  - 2\\n  - 20\\n  - 21\\n  - 34\\n  - 2\\n  - 22\\n  - 23\\n  xxx_nounkeyedliteral: {}\\n  xxx_unrecognized: []\\n  xxx_sizecache: 0\\n\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/kira.gov.MsgSubmitProposal"
              }
            ]
          },
          {
            "type": "submit_proposal",
            "attributes": [
              {
                "key": "proposal_id",
                "value": "22"
              },
              {
                "key": "proposal_type",
                "value": "CreateRoleProposal"
              },
              {
                "key": "proposal_content",
                "value": "proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\ntitle: Test proposal for 'bruh3' role\ndescription: Test proposal for creating 'bruh3' role\ncontent:\n  typeurl: /kira.gov.CreateRoleProposal\n  value:\n  - 10\n  - 5\n  - 98\n  - 114\n  - 117\n  - 104\n  - 51\n  - 18\n  - 14\n  - 39\n  - 66\n  - 114\n  - 117\n  - 104\n  - 39\n  - 40\n  - 51\n  - 41\n  - 32\n  - 114\n  - 111\n  - 108\n  - 101\n  - 26\n  - 2\n  - 20\n  - 21\n  - 34\n  - 2\n  - 22\n  - 23\n  xxx_nounkeyedliteral: {}\n  xxx_unrecognized: []\n  xxx_sizecache: 0\n"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "0",
    "gas_used": "0",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/kira.gov.MsgSubmitProposal",
            "proposer": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
            "title": "Test proposal for 'bruh3' role",
            "description": "Test proposal for creating 'bruh3' role",
            "content": {
              "@type": "/kira.gov.CreateRoleProposal",
              "role_sid": "bruh3",
              "role_description": "'Bruh'(3) role",
              "whitelisted_permissions": [
                "PERMISSION_CREATE_UNJAIL_VALIDATOR_PROPOSAL",
                "PERMISSION_VOTE_UNJAIL_VALIDATOR_PROPOSAL"
              ],
              "blacklisted_permissions": [
                "PERMISSION_CREATE_CREATE_ROLE_PROPOSAL",
                "PERMISSION_VOTE_CREATE_ROLE_PROPOSAL"
              ]
            }
          }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
      },
      "auth_info": {
        "signer_infos": [
          {
            "public_key": {
              "@type": "/cosmos.crypto.secp256k1.PubKey",
              "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "109"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "ukex",
              "amount": "100"
            }
          ],
          "gas_limit": "200000",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
        "U98UymXe9j7LpZwcbKAo6F4wNZV1X2haFdihaqkgQs4uNCJJNGaW1Kujp/qmvtBVyXBkF/YYWQEKDFtepBg5ZA=="
      ]
    },
    "timestamp": "2023-06-12T13:42:16Z",
    "events": [
      {
        "type": "tx",
        "attributes": [
          {
            "key": "YWNjX3NlcQ==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NC8xMDk=",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "VTk4VXltWGU5ajdMcFp3Y2JLQW82RjR3TlpWMVgyaGFGZGloYXFrZ1FzNHVOQ0pKTkdhVzFLdWpwL3FtdnRCVnlYQmtGL1lZV1FFS0RGdGVwQmc1WkE9PQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "ZmVl",
            "value": "MTAwdWtleA==",
            "index": true
          },
          {
            "key": "ZmVlX3BheWVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "L2tpcmEuZ292Lk1zZ1N1Ym1pdFByb3Bvc2Fs",
            "index": true
          }
        ]
      },
      {
        "type": "submit_proposal",
        "attributes": [
          {
            "key": "cHJvcG9zYWxfaWQ=",
            "value": "MjI=",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfdHlwZQ==",
            "value": "Q3JlYXRlUm9sZVByb3Bvc2Fs",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfY29udGVudA==",
            "value": "cHJvcG9zZXI6IGtpcmExdm13ZGd3NDI2YWo5ZngzM2ZxdXNtdGc2cjY1eXl1Y214NnJkdDQKdGl0bGU6IFRlc3QgcHJvcG9zYWwgZm9yICdicnVoMycgcm9sZQpkZXNjcmlwdGlvbjogVGVzdCBwcm9wb3NhbCBmb3IgY3JlYXRpbmcgJ2JydWgzJyByb2xlCmNvbnRlbnQ6CiAgdHlwZXVybDogL2tpcmEuZ292LkNyZWF0ZVJvbGVQcm9wb3NhbAogIHZhbHVlOgogIC0gMTAKICAtIDUKICAtIDk4CiAgLSAxMTQKICAtIDExNwogIC0gMTA0CiAgLSA1MQogIC0gMTgKICAtIDE0CiAgLSAzOQogIC0gNjYKICAtIDExNAogIC0gMTE3CiAgLSAxMDQKICAtIDM5CiAgLSA0MAogIC0gNTEKICAtIDQxCiAgLSAzMgogIC0gMTE0CiAgLSAxMTEKICAtIDEwOAogIC0gMTAxCiAgLSAyNgogIC0gMgogIC0gMjAKICAtIDIxCiAgLSAzNAogIC0gMgogIC0gMjIKICAtIDIzCiAgeHh4X25vdW5rZXllZGxpdGVyYWw6IHt9CiAgeHh4X3VucmVjb2duaXplZDogW10KICB4eHhfc2l6ZWNhY2hlOiAwCg==",
            "index": true
          }
        ]
      }
    ]
  }
  ```
</details>

<details>
  <summary>Check proposal</summary>

  ```
  sekaid q customgov proposal 22 -o json | jq
  ```

  ```json
  {
    "proposal_id": "22",
    "title": "Test proposal for 'bruh3' role",
    "description": "Test proposal for creating 'bruh3' role",
    "content": {
      "@type": "/kira.gov.CreateRoleProposal",
      "role_sid": "bruh3",
      "role_description": "'Bruh'(3) role",
      "whitelisted_permissions": [
        "PERMISSION_CREATE_UNJAIL_VALIDATOR_PROPOSAL",
        "PERMISSION_VOTE_UNJAIL_VALIDATOR_PROPOSAL"
      ],
      "blacklisted_permissions": [
        "PERMISSION_CREATE_CREATE_ROLE_PROPOSAL",
        "PERMISSION_VOTE_CREATE_ROLE_PROPOSAL"
      ]
    },
    "submit_time": "2023-06-12T13:42:16.485050497Z",
    "voting_end_time": "2023-06-12T13:48:16.485050497Z",
    "enactment_end_time": "2023-06-12T13:53:16.485050497Z",
    "min_voting_end_block_height": "142855",
    "min_enactment_end_block_height": "142856",
    "result": "VOTE_PENDING",
    "exec_result": ""
  }
  ```
</details>

[Return to top](#sekai)

###### 21.7.7.4.3 remove

Raise governance proposal to remove a role.

Usage:
```
sekaid tx customgov proposal role remove [role_sid] [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ❌ ?   |
| `--description string`        | The description of the proposal, it can be a url, some text, etc.                                                                                           | ✅ yes |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ✅ yes |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ✅ yes |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for remove                                                                                                                                             | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default 'home' directory will be used                                                                         | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ❌ ?   |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ✅ yes |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `--title string`              | The title of the proposal.                                                                                                                                  | ✅ yes |
| `--whitelist int32Slice`      | the whitelist value in format 1,2,3 (default `[]`)                                                                                                          | ✅ yes |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx customgov proposal role remove --help
Raise governance proposal to remove a role.

Usage:
  sekaid tx customgov proposal role remove [role_sid] [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --description string       The description of the proposal, it can be a url, some text, etc.
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for remove
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
      --title string             The title of the proposal.
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid tx customgov proposal role remove 'bruh2' --title="Test proposal" --description="Test proposal" --from=kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --keyring-backend=test --chain-id=localnet-4 --home=/root/.sekai --fees=100ukex --broadcast-mode=async --yes --output=json | jq
```

❌ _remove role proposal is not implemented!_

[Return to top](#sekai)

###### 21.7.7.4.4 remove-blacklisted-permission

Raise governance proposal to remove a blacklisted permission from a role.

Usage:
```
sekaid tx customgov proposal role remove-blacklisted-permission [role_sid] [role_description] [flags] ❌ mistake! [role_sid | role_id] [permission_id]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `--blacklist int32Slice`      | the blacklist values in format 1,2,3 (default `[]`)                                                                                                         | ❌ ?   |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ❌ ?   |
| `--description string`        | The description of the proposal, it can be a url, some text, etc.                                                                                           | ✅ yes |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ✅ yes |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ✅ yes |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for remove-blacklisted-permission                                                                                                                      | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default 'home' directory will be used                                                                         | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ❌ ?   |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ✅ yes |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `--title string`              | The title of the proposal.                                                                                                                                  | ✅ yes |
| `--whitelist int32Slice`      | the whitelist value in format 1,2,3 (default `[]`)                                                                                                          | ❌ ?   |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx customgov proposal role remove-blacklisted-permission --help
Raise governance proposal to remove a blacklisted permission from a role.

Usage:
  sekaid tx customgov proposal role remove-blacklisted-permission [role_sid] [role_description] [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
      --blacklist int32Slice     the blacklist values in format 1,2,3 (default [])
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --description string       The description of the proposal, it can be a url, some text, etc.
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for remove-blacklisted-permission
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
      --title string             The title of the proposal.
      --whitelist int32Slice     the whitelist value in format 1,2,3 (default [])
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid tx customgov proposal role remove-blacklisted-permission bruh 10 --title="Test proposal" --description="Test proposal" --from=kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --keyring-backend=test --chain-id=localnet-4 --home=/root/.sekai --fees=100ukex --broadcast-mode=async --yes --output=json | jq
{
  "height": "0",
  "txhash": "0435572323540F92E5F3EC95086FA34DB6C09D1A9AF57B3FAF40BBD6F6AC58CE",
  "codespace": "",
  "code": 0,
  "data": "",
  "raw_log": "",
  "logs": [],
  "info": "",
  "gas_wanted": "0",
  "gas_used": "0",
  "tx": null,
  "timestamp": "",
  "events": []
}
```

<details>
  <summary>Check tx execution</summary>

  ```
  sekaid q tx 0435572323540F92E5F3EC95086FA34DB6C09D1A9AF57B3FAF40BBD6F6AC58CE -o json | jq
  ```

  ```json
  {
    "height": "142956",
    "txhash": "0435572323540F92E5F3EC95086FA34DB6C09D1A9AF57B3FAF40BBD6F6AC58CE",
    "codespace": "",
    "code": 0,
    "data": "0A210A1B2F6B6972612E676F762E4D73675375626D697450726F706F73616C12020817",
    "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/kira.gov.MsgSubmitProposal\"}]},{\"type\":\"submit_proposal\",\"attributes\":[{\"key\":\"proposal_id\",\"value\":\"23\"},{\"key\":\"proposal_type\",\"value\":\"RemoveBlacklistedRolePermission\"},{\"key\":\"proposal_content\",\"value\":\"proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\\ntitle: Test proposal\\ndescription: Test proposal\\ncontent:\\n  typeurl: /kira.gov.RemoveBlacklistedRolePermissionProposal\\n  value:\\n  - 10\\n  - 4\\n  - 98\\n  - 114\\n  - 117\\n  - 104\\n  - 16\\n  - 10\\n  xxx_nounkeyedliteral: {}\\n  xxx_unrecognized: []\\n  xxx_sizecache: 0\\n\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/kira.gov.MsgSubmitProposal"
              }
            ]
          },
          {
            "type": "submit_proposal",
            "attributes": [
              {
                "key": "proposal_id",
                "value": "23"
              },
              {
                "key": "proposal_type",
                "value": "RemoveBlacklistedRolePermission"
              },
              {
                "key": "proposal_content",
                "value": "proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\ntitle: Test proposal\ndescription: Test proposal\ncontent:\n  typeurl: /kira.gov.RemoveBlacklistedRolePermissionProposal\n  value:\n  - 10\n  - 4\n  - 98\n  - 114\n  - 117\n  - 104\n  - 16\n  - 10\n  xxx_nounkeyedliteral: {}\n  xxx_unrecognized: []\n  xxx_sizecache: 0\n"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "0",
    "gas_used": "0",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/kira.gov.MsgSubmitProposal",
            "proposer": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
            "title": "Test proposal",
            "description": "Test proposal",
            "content": {
              "@type": "/kira.gov.RemoveBlacklistedRolePermissionProposal",
              "role_sid": "bruh",
              "permission": "PERMISSION_CREATE_UPSERT_DATA_REGISTRY_PROPOSAL"
            }
          }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
      },
      "auth_info": {
        "signer_infos": [
          {
            "public_key": {
              "@type": "/cosmos.crypto.secp256k1.PubKey",
              "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "113"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "ukex",
              "amount": "100"
            }
          ],
          "gas_limit": "200000",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
        "0UhZaikBOdSDX/4E1Bq/0fyNFCHheRN3QVOyGAdQ+wlxM9vXGnBM0L/JvUYBsAQEYcgxBYsyE8mttkD5nDxzmA=="
      ]
    },
    "timestamp": "2023-06-12T14:00:01Z",
    "events": [
      {
        "type": "tx",
        "attributes": [
          {
            "key": "YWNjX3NlcQ==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NC8xMTM=",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "MFVoWmFpa0JPZFNEWC80RTFCcS8wZnlORkNIaGVSTjNRVk95R0FkUSt3bHhNOXZYR25CTTBML0p2VVlCc0FRRVljZ3hCWXN5RThtdHRrRDVuRHh6bUE9PQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "ZmVl",
            "value": "MTAwdWtleA==",
            "index": true
          },
          {
            "key": "ZmVlX3BheWVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "L2tpcmEuZ292Lk1zZ1N1Ym1pdFByb3Bvc2Fs",
            "index": true
          }
        ]
      },
      {
        "type": "submit_proposal",
        "attributes": [
          {
            "key": "cHJvcG9zYWxfaWQ=",
            "value": "MjM=",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfdHlwZQ==",
            "value": "UmVtb3ZlQmxhY2tsaXN0ZWRSb2xlUGVybWlzc2lvbg==",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfY29udGVudA==",
            "value": "cHJvcG9zZXI6IGtpcmExdm13ZGd3NDI2YWo5ZngzM2ZxdXNtdGc2cjY1eXl1Y214NnJkdDQKdGl0bGU6IFRlc3QgcHJvcG9zYWwKZGVzY3JpcHRpb246IFRlc3QgcHJvcG9zYWwKY29udGVudDoKICB0eXBldXJsOiAva2lyYS5nb3YuUmVtb3ZlQmxhY2tsaXN0ZWRSb2xlUGVybWlzc2lvblByb3Bvc2FsCiAgdmFsdWU6CiAgLSAxMAogIC0gNAogIC0gOTgKICAtIDExNAogIC0gMTE3CiAgLSAxMDQKICAtIDE2CiAgLSAxMAogIHh4eF9ub3Vua2V5ZWRsaXRlcmFsOiB7fQogIHh4eF91bnJlY29nbml6ZWQ6IFtdCiAgeHh4X3NpemVjYWNoZTogMAo=",
            "index": true
          }
        ]
      }
    ]
  }
  ```
</details>

<details>
  <summary>Check proposal</summary>

  ```
  sekaid q customgov proposal 23 -o json | jq
  ```

  ```json
  {
    "proposal_id": "23",
    "title": "Test proposal",
    "description": "Test proposal",
    "content": {
      "@type": "/kira.gov.RemoveBlacklistedRolePermissionProposal",
      "role_sid": "bruh",
      "permission": "PERMISSION_CREATE_UPSERT_DATA_REGISTRY_PROPOSAL"
    },
    "submit_time": "2023-06-12T14:00:01.126696499Z",
    "voting_end_time": "2023-06-12T14:06:01.126696499Z",
    "enactment_end_time": "2023-06-12T14:11:01.126696499Z",
    "min_voting_end_block_height": "142958",
    "min_enactment_end_block_height": "142959",
    "result": "VOTE_PENDING",
    "exec_result": ""
  }
  ```
</details>

[Return to top](#sekai)

###### 21.7.7.4.5 remove-whitelisted-permission

Raise governance proposal to remove whitelisted permission from a role.

Usage:
```
sekaid tx customgov proposal role remove-whitelisted-permission [role_sid] [permission] [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `--blacklist int32Slice`      | the blacklist values in format 1,2,3 (default `[]`)                                                                                                         | ❌ ?   |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ❌ ?   |
| `--description string`        | The description of the proposal, it can be a url, some text, etc.                                                                                           | ✅ yes |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ✅ yes |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ✅ yes |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for remove-whitelisted-permission                                                                                                                      | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default 'home' directory will be used                                                                         | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ❌ ?   |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ✅ yes |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `--title string`              | The title of the proposal.                                                                                                                                  | ✅ yes |
| `--whitelist int32Slice`      | the whitelist value in format 1,2,3 (default `[]`)                                                                                                          | ❌ ?   |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx customgov proposal role remove-whitelisted-permission --help
Raise governance proposal to remove whitelisted permission from a role.

Usage:
  sekaid tx customgov proposal role remove-whitelisted-permission [role_sid] [permission] [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
      --blacklist int32Slice     the blacklist values in format 1,2,3 (default [])
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --description string       The description of the proposal, it can be a url, some text, etc.
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for remove-whitelisted-permission
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
      --title string             The title of the proposal.
      --whitelist int32Slice     the whitelist value in format 1,2,3 (default [])
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid tx customgov proposal role remove-whitelisted-permission 4 28 --title="Test proposal" --description="Test proposal" --from=kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --keyring-backend=test --chain-id=localnet-4 --home=/root/.sekai --fees=100ukex --broadcast-mode=async --yes --output=json | jq
{
  "height": "0",
  "txhash": "445C111D2424B3619FC90554F64E735809F56F9CF21FE689049E48877A2EDAE8",
  "codespace": "",
  "code": 0,
  "data": "",
  "raw_log": "",
  "logs": [],
  "info": "",
  "gas_wanted": "0",
  "gas_used": "0",
  "tx": null,
  "timestamp": "",
  "events": []
}
```

<details>
  <summary>Check tx execution</summary>

  ```
  sekaid q tx 445C111D2424B3619FC90554F64E735809F56F9CF21FE689049E48877A2EDAE8 -o json | jq
  ```

  ```json
  {
    "height": "142994",
    "txhash": "445C111D2424B3619FC90554F64E735809F56F9CF21FE689049E48877A2EDAE8",
    "codespace": "",
    "code": 0,
    "data": "0A210A1B2F6B6972612E676F762E4D73675375626D697450726F706F73616C12020819",
    "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/kira.gov.MsgSubmitProposal\"}]},{\"type\":\"submit_proposal\",\"attributes\":[{\"key\":\"proposal_id\",\"value\":\"25\"},{\"key\":\"proposal_type\",\"value\":\"RemoveWhitelistedRolePermission\"},{\"key\":\"proposal_content\",\"value\":\"proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\\ntitle: Test proposal\\ndescription: Test proposal\\ncontent:\\n  typeurl: /kira.gov.RemoveWhitelistedRolePermissionProposal\\n  value:\\n  - 10\\n  - 1\\n  - 52\\n  - 16\\n  - 28\\n  xxx_nounkeyedliteral: {}\\n  xxx_unrecognized: []\\n  xxx_sizecache: 0\\n\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/kira.gov.MsgSubmitProposal"
              }
            ]
          },
          {
            "type": "submit_proposal",
            "attributes": [
              {
                "key": "proposal_id",
                "value": "25"
              },
              {
                "key": "proposal_type",
                "value": "RemoveWhitelistedRolePermission"
              },
              {
                "key": "proposal_content",
                "value": "proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\ntitle: Test proposal\ndescription: Test proposal\ncontent:\n  typeurl: /kira.gov.RemoveWhitelistedRolePermissionProposal\n  value:\n  - 10\n  - 1\n  - 52\n  - 16\n  - 28\n  xxx_nounkeyedliteral: {}\n  xxx_unrecognized: []\n  xxx_sizecache: 0\n"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "0",
    "gas_used": "0",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/kira.gov.MsgSubmitProposal",
            "proposer": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
            "title": "Test proposal",
            "description": "Test proposal",
            "content": {
              "@type": "/kira.gov.RemoveWhitelistedRolePermissionProposal",
              "role_sid": "4",
              "permission": "PERMISSION_CREATE_SOFTWARE_UPGRADE_PROPOSAL"
            }
          }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
      },
      "auth_info": {
        "signer_infos": [
          {
            "public_key": {
              "@type": "/cosmos.crypto.secp256k1.PubKey",
              "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "115"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "ukex",
              "amount": "100"
            }
          ],
          "gas_limit": "200000",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
        "ZsEjyPOsXBVV4xZ2fIzpAeq1WBaGwH4KgdEIQfpbAftM2TWy4SRWt4yhwQnpcvQH2+tY4px/6I7EOk1DfnbIGg=="
      ]
    },
    "timestamp": "2023-06-12T14:06:33Z",
    "events": [
      {
        "type": "tx",
        "attributes": [
          {
            "key": "YWNjX3NlcQ==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NC8xMTU=",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "WnNFanlQT3NYQlZWNHhaMmZJenBBZXExV0JhR3dINEtnZEVJUWZwYkFmdE0yVFd5NFNSV3Q0eWh3UW5wY3ZRSDIrdFk0cHgvNkk3RU9rMURmbmJJR2c9PQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "ZmVl",
            "value": "MTAwdWtleA==",
            "index": true
          },
          {
            "key": "ZmVlX3BheWVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "L2tpcmEuZ292Lk1zZ1N1Ym1pdFByb3Bvc2Fs",
            "index": true
          }
        ]
      },
      {
        "type": "submit_proposal",
        "attributes": [
          {
            "key": "cHJvcG9zYWxfaWQ=",
            "value": "MjU=",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfdHlwZQ==",
            "value": "UmVtb3ZlV2hpdGVsaXN0ZWRSb2xlUGVybWlzc2lvbg==",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfY29udGVudA==",
            "value": "cHJvcG9zZXI6IGtpcmExdm13ZGd3NDI2YWo5ZngzM2ZxdXNtdGc2cjY1eXl1Y214NnJkdDQKdGl0bGU6IFRlc3QgcHJvcG9zYWwKZGVzY3JpcHRpb246IFRlc3QgcHJvcG9zYWwKY29udGVudDoKICB0eXBldXJsOiAva2lyYS5nb3YuUmVtb3ZlV2hpdGVsaXN0ZWRSb2xlUGVybWlzc2lvblByb3Bvc2FsCiAgdmFsdWU6CiAgLSAxMAogIC0gMQogIC0gNTIKICAtIDE2CiAgLSAyOAogIHh4eF9ub3Vua2V5ZWRsaXRlcmFsOiB7fQogIHh4eF91bnJlY29nbml6ZWQ6IFtdCiAgeHh4X3NpemVjYWNoZTogMAo=",
            "index": true
          }
        ]
      }
    ]
  }
  ```
</details>

<details>
  <summary>Check proposal</summary>

  ```
  sekaid q customgov proposal 25 -o json | jq
  ```

  ```json
  {
    "proposal_id": "25",
    "title": "Test proposal",
    "description": "Test proposal",
    "content": {
      "@type": "/kira.gov.RemoveWhitelistedRolePermissionProposal",
      "role_sid": "4",
      "permission": "PERMISSION_CREATE_SOFTWARE_UPGRADE_PROPOSAL"
    },
    "submit_time": "2023-06-12T14:06:33.096171946Z",
    "voting_end_time": "2023-06-12T14:12:33.096171946Z",
    "enactment_end_time": "2023-06-12T14:17:33.096171946Z",
    "min_voting_end_block_height": "142996",
    "min_enactment_end_block_height": "143030",
    "result": "VOTE_RESULT_QUORUM_NOT_REACHED",
    "exec_result": ""
  }
  ```
</details>

[Return to top](#sekai)

###### 21.7.7.4.6 whitelist-permission

Raise governance proposal to whitelist a permission for a role.

Usage:
```
sekaid tx customgov proposal role whitelist-permission [role_sid] [permission_id] [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `--blacklist int32Slice`      | the blacklist values in format 1,2,3 (default `[]`)                                                                                                         | ❌ ?   |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ❌ ?   |
| `--description string`        | The description of the proposal, it can be a url, some text, etc.                                                                                           | ✅ yes |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ✅ yes |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ✅ yes |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for whitelist-permission                                                                                                                               | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default 'home' directory will be used                                                                         | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ❌ ?   |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ✅ yes |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `--title string`              | The title of the proposal.                                                                                                                                  | ✅ yes |
| `--whitelist int32Slice`      | the whitelist value in format 1,2,3 (default `[]`)                                                                                                          | ❌ ?   |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx customgov proposal role whitelist-permission --help
Raise governance proposal to whitelist a permission for a role.

Usage:
  sekaid tx customgov proposal role whitelist-permission [role_sid] [permission_id] [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
      --blacklist int32Slice     the blacklist values in format 1,2,3 (default [])
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --description string       The description of the proposal, it can be a url, some text, etc.
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for whitelist-permission
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
      --title string             The title of the proposal.
      --whitelist int32Slice     the whitelist value in format 1,2,3 (default [])
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid tx customgov proposal role whitelist-permission 3 11 --title="Test proposal" --description="Test proposal" --from=kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --keyring-backend=test --chain-id=localnet-4 --home=/root/.sekai --fees=100ukex --broadcast-mode=async --yes --output=json | jq
{
  "height": "0",
  "txhash": "383217AAFC0C23FCA0D13D253CEAFE883E1E89B3D512AA985E65D98439624CC3",
  "codespace": "",
  "code": 0,
  "data": "",
  "raw_log": "",
  "logs": [],
  "info": "",
  "gas_wanted": "0",
  "gas_used": "0",
  "tx": null,
  "timestamp": "",
  "events": []
}
```

<details>
  <summary>Check tx execution</summary>

  ```
  sekaid q tx 383217AAFC0C23FCA0D13D253CEAFE883E1E89B3D512AA985E65D98439624CC3 -o json | jq
  ```

  ```json
  {
    "height": "143069",
    "txhash": "383217AAFC0C23FCA0D13D253CEAFE883E1E89B3D512AA985E65D98439624CC3",
    "codespace": "",
    "code": 0,
    "data": "0A210A1B2F6B6972612E676F762E4D73675375626D697450726F706F73616C1202081A",
    "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/kira.gov.MsgSubmitProposal\"}]},{\"type\":\"submit_proposal\",\"attributes\":[{\"key\":\"proposal_id\",\"value\":\"26\"},{\"key\":\"proposal_type\",\"value\":\"WhitelistRolePermission\"},{\"key\":\"proposal_content\",\"value\":\"proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\\ntitle: Test proposal\\ndescription: Test proposal\\ncontent:\\n  typeurl: /kira.gov.WhitelistRolePermissionProposal\\n  value:\\n  - 10\\n  - 1\\n  - 51\\n  - 16\\n  - 11\\n  xxx_nounkeyedliteral: {}\\n  xxx_unrecognized: []\\n  xxx_sizecache: 0\\n\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/kira.gov.MsgSubmitProposal"
              }
            ]
          },
          {
            "type": "submit_proposal",
            "attributes": [
              {
                "key": "proposal_id",
                "value": "26"
              },
              {
                "key": "proposal_type",
                "value": "WhitelistRolePermission"
              },
              {
                "key": "proposal_content",
                "value": "proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\ntitle: Test proposal\ndescription: Test proposal\ncontent:\n  typeurl: /kira.gov.WhitelistRolePermissionProposal\n  value:\n  - 10\n  - 1\n  - 51\n  - 16\n  - 11\n  xxx_nounkeyedliteral: {}\n  xxx_unrecognized: []\n  xxx_sizecache: 0\n"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "0",
    "gas_used": "0",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/kira.gov.MsgSubmitProposal",
            "proposer": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
            "title": "Test proposal",
            "description": "Test proposal",
            "content": {
              "@type": "/kira.gov.WhitelistRolePermissionProposal",
              "role_identifier": "3",
              "permission": "PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL"
            }
          }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
      },
      "auth_info": {
        "signer_infos": [
          {
            "public_key": {
              "@type": "/cosmos.crypto.secp256k1.PubKey",
              "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "116"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "ukex",
              "amount": "100"
            }
          ],
          "gas_limit": "200000",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
        "DoKkT3wrfNzoweZ9Ln7vZdrP5afBxZ68IwP7bqnPsMUoy6eSXIYDJXKIMbfP4OybOduKi4/Ptd6nfZrc3TxcRA=="
      ]
    },
    "timestamp": "2023-06-12T14:19:27Z",
    "events": [
      {
        "type": "tx",
        "attributes": [
          {
            "key": "YWNjX3NlcQ==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NC8xMTY=",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "RG9La1Qzd3JmTnpvd2VaOUxuN3ZaZHJQNWFmQnhaNjhJd1A3YnFuUHNNVW95NmVTWElZREpYS0lNYmZQNE95Yk9kdUtpNC9QdGQ2bmZacmMzVHhjUkE9PQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "ZmVl",
            "value": "MTAwdWtleA==",
            "index": true
          },
          {
            "key": "ZmVlX3BheWVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "L2tpcmEuZ292Lk1zZ1N1Ym1pdFByb3Bvc2Fs",
            "index": true
          }
        ]
      },
      {
        "type": "submit_proposal",
        "attributes": [
          {
            "key": "cHJvcG9zYWxfaWQ=",
            "value": "MjY=",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfdHlwZQ==",
            "value": "V2hpdGVsaXN0Um9sZVBlcm1pc3Npb24=",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfY29udGVudA==",
            "value": "cHJvcG9zZXI6IGtpcmExdm13ZGd3NDI2YWo5ZngzM2ZxdXNtdGc2cjY1eXl1Y214NnJkdDQKdGl0bGU6IFRlc3QgcHJvcG9zYWwKZGVzY3JpcHRpb246IFRlc3QgcHJvcG9zYWwKY29udGVudDoKICB0eXBldXJsOiAva2lyYS5nb3YuV2hpdGVsaXN0Um9sZVBlcm1pc3Npb25Qcm9wb3NhbAogIHZhbHVlOgogIC0gMTAKICAtIDEKICAtIDUxCiAgLSAxNgogIC0gMTEKICB4eHhfbm91bmtleWVkbGl0ZXJhbDoge30KICB4eHhfdW5yZWNvZ25pemVkOiBbXQogIHh4eF9zaXplY2FjaGU6IDAK",
            "index": true
          }
        ]
      }
    ]
  }
  ```
</details>

<details>
  <summary>Check proposal</summary>

  ```
  sekaid q customgov proposal 26 -o json | jq
  ```

  ```json
  {
    "proposal_id": "26",
    "title": "Test proposal",
    "description": "Test proposal",
    "content": {
      "@type": "/kira.gov.WhitelistRolePermissionProposal",
      "role_identifier": "3",
      "permission": "PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL"
    },
    "submit_time": "2023-06-12T14:19:27.042751264Z",
    "voting_end_time": "2023-06-12T14:25:27.042751264Z",
    "enactment_end_time": "2023-06-12T14:30:27.042751264Z",
    "min_voting_end_block_height": "143071",
    "min_enactment_end_block_height": "143072",
    "result": "VOTE_PENDING",
    "exec_result": ""
  }
  ```
</details>

[Return to top](#sekai)

###### 21.7.7.5 set-network-property

🟨 Create a proposal to set a value on a network property. 🟨

Usage:
```
sekaid tx customgov proposal set-network-property <property> <value> [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ❌ ?   |
| `--description string`        | The description of the proposal, it can be a url, some text, etc.                                                                                           | ✅ yes |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ✅ yes |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ✅ yes |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for set-network-property                                                                                                                               | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default 'home' directory will be used                                                                         | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ❌ ?   |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ✅ yes |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `--title string`              | The title of the proposal.                                                                                                                                  | ✅ yes |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx customgov proposal set-network-property --help

                $ %s tx customgov proposal set-network-property MIN_TX_FEE 100 --from=<key_or_address>

                Available properties:
                        MIN_TX_FEE
                        MAX_TX_FEE
                        VOTE_QUORUM
                        PROPOSAL_END_TIME
                        PROPOSAL_ENACTMENT_TIME
                        ENABLE_FOREIGN_TX_FEE_PAYMENTS

Usage:
  sekaid tx customgov proposal set-network-property <property> <value> [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --description string       The description of the proposal, it can be a url, some text, etc.
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for set-network-property
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
      --title string             The title of the proposal.
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid tx customgov proposal set-network-property "MIN_TX_FEE" 10 --title="Reducing value of minimal fee for docs" --description="Changing value of minimal fee for testing for docs" --from=kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --chain-id=localnet-4 --home=/root/.sekai --keyring-backend=test --fees=100ukex --yes --output=json | jq
{
  "height": "0",
  "txhash": "0C192B8B81789853A107AD6B6283C0AB1784A2B2C32E1F9D6730341F41BDE413",
  "codespace": "",
  "code": 0,
  "data": "",
  "raw_log": "[]",
  "logs": [],
  "info": "",
  "gas_wanted": "0",
  "gas_used": "0",
  "tx": null,
  "timestamp": "",
  "events": []
}
```

<details>
  <summary>Check tx execution</summary>

  ```
  sekaid q tx 0C192B8B81789853A107AD6B6283C0AB1784A2B2C32E1F9D6730341F41BDE413 -o json | jq
  ```

  ```json
  {
    "height": "110257",
    "txhash": "0C192B8B81789853A107AD6B6283C0AB1784A2B2C32E1F9D6730341F41BDE413",
    "codespace": "",
    "code": 0,
    "data": "0A210A1B2F6B6972612E676F762E4D73675375626D697450726F706F73616C1202080C",
    "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/kira.gov.MsgSubmitProposal\"}]},{\"type\":\"submit_proposal\",\"attributes\":[{\"key\":\"proposal_id\",\"value\":\"12\"},{\"key\":\"proposal_type\",\"value\":\"SetNetworkProperty\"},{\"key\":\"proposal_content\",\"value\":\"proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\\ntitle: Reducing value of minimal fee for docs\\ndescription: Changing value of minimal fee for testing for docs\\ncontent:\\n  typeurl: /kira.gov.SetNetworkPropertyProposal\\n  value:\\n  - 18\\n  - 2\\n  - 8\\n  - 10\\n  xxx_nounkeyedliteral: {}\\n  xxx_unrecognized: []\\n  xxx_sizecache: 0\\n\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/kira.gov.MsgSubmitProposal"
              }
            ]
          },
          {
            "type": "submit_proposal",
            "attributes": [
              {
                "key": "proposal_id",
                "value": "12"
              },
              {
                "key": "proposal_type",
                "value": "SetNetworkProperty"
              },
              {
                "key": "proposal_content",
                "value": "proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\ntitle: Reducing value of minimal fee for docs\ndescription: Changing value of minimal fee for testing for docs\ncontent:\n  typeurl: /kira.gov.SetNetworkPropertyProposal\n  value:\n  - 18\n  - 2\n  - 8\n  - 10\n  xxx_nounkeyedliteral: {}\n  xxx_unrecognized: []\n  xxx_sizecache: 0\n"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "0",
    "gas_used": "0",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/kira.gov.MsgSubmitProposal",
            "proposer": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
            "title": "Reducing value of minimal fee for docs",
            "description": "Changing value of minimal fee for testing for docs",
            "content": {
              "@type": "/kira.gov.SetNetworkPropertyProposal",
              "network_property": "MIN_TX_FEE",
              "value": {
                "value": "10",
                "str_value": ""
              }
            }
          }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
      },
      "auth_info": {
        "signer_infos": [
          {
            "public_key": {
              "@type": "/cosmos.crypto.secp256k1.PubKey",
              "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "93"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "ukex",
              "amount": "100"
            }
          ],
          "gas_limit": "200000",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
        "UC3fpkO1X1udG0Hmk6PneiOqvRaQ/0ULOCGigKXuwVtIQONk2hVaTsXi9KdNhThD7dM0T+hvhHeUsQ4DAQcdWA=="
      ]
    },
    "timestamp": "2023-06-08T16:19:07Z",
    "events": [
      {
        "type": "tx",
        "attributes": [
          {
            "key": "YWNjX3NlcQ==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NC85Mw==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "VUMzZnBrTzFYMXVkRzBIbWs2UG5laU9xdlJhUS8wVUxPQ0dpZ0tYdXdWdElRT05rMmhWYVRzWGk5S2ROaFRoRDdkTTBUK2h2aEhlVXNRNERBUWNkV0E9PQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "ZmVl",
            "value": "MTAwdWtleA==",
            "index": true
          },
          {
            "key": "ZmVlX3BheWVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "L2tpcmEuZ292Lk1zZ1N1Ym1pdFByb3Bvc2Fs",
            "index": true
          }
        ]
      },
      {
        "type": "submit_proposal",
        "attributes": [
          {
            "key": "cHJvcG9zYWxfaWQ=",
            "value": "MTI=",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfdHlwZQ==",
            "value": "U2V0TmV0d29ya1Byb3BlcnR5",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfY29udGVudA==",
            "value": "cHJvcG9zZXI6IGtpcmExdm13ZGd3NDI2YWo5ZngzM2ZxdXNtdGc2cjY1eXl1Y214NnJkdDQKdGl0bGU6IFJlZHVjaW5nIHZhbHVlIG9mIG1pbmltYWwgZmVlIGZvciBkb2NzCmRlc2NyaXB0aW9uOiBDaGFuZ2luZyB2YWx1ZSBvZiBtaW5pbWFsIGZlZSBmb3IgdGVzdGluZyBmb3IgZG9jcwpjb250ZW50OgogIHR5cGV1cmw6IC9raXJhLmdvdi5TZXROZXR3b3JrUHJvcGVydHlQcm9wb3NhbAogIHZhbHVlOgogIC0gMTgKICAtIDIKICAtIDgKICAtIDEwCiAgeHh4X25vdW5rZXllZGxpdGVyYWw6IHt9CiAgeHh4X3VucmVjb2duaXplZDogW10KICB4eHhfc2l6ZWNhY2hlOiAwCg==",
            "index": true
          }
        ]
      }
    ]
  }
  ```

  > OR find all transactions for
  ```
  sekaid query txs --events=submit_proposal.proposal_type=SetNetworkProperty --output=json | jq
  ```

  __VERY LONG JSON with all txs__
  ```json
  {
    "total_count": "2",
    "count": "2",
    "page_number": "1",
    "page_total": "1",
    "limit": "30",
    "txs": [
      {...},
      {...}
    ]
  }
  ```
</details>

<details>
  <summary>Check the vote</summary>

  ```
  sekaid query customgov proposal 12 -o json | jq
  ```

  ```json
  {
    "proposal_id": "12",
    "title": "Reducing value of minimal fee for docs",
    "description": "Changing value of minimal fee for testing for docs",
    "content": {
      "@type": "/kira.gov.SetNetworkPropertyProposal",
      "network_property": "MIN_TX_FEE",
      "value": {
        "value": "10",
        "str_value": ""
      }
    },
    "submit_time": "2023-06-08T16:19:07.409676597Z",
    "voting_end_time": "2023-06-08T16:25:07.409676597Z",
    "enactment_end_time": "2023-06-08T16:30:07.409676597Z",
    "min_voting_end_block_height": "110259",
    "min_enactment_end_block_height": "110293",
    "result": "VOTE_RESULT_ENACTMENT",
    "exec_result": ""
  }
  ```
</details>

<details>
  <summary>Check network property (after successful vote)</summary>

  List of all network properties:
  ```
  sekaid query customgov network-properties --output=json | jq
  ```

  ```json
  {
    "properties": {
      "min_tx_fee": "10",
      "max_tx_fee": "1000000",
      // . . .
    }
  }
  ```
</details>

[Return to top](#sekai)

###### 21.7.7.6 set-poor-network-msgs

🟨 Create a proposal to set a value on a network property. 🟨

Usage:
```
sekaid tx customgov proposal set-poor-network-msgs <messages> [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ❌ ?   |
| `--description string`        | The description of the proposal, it can be a url, some text, etc.                                                                                           | ✅ yes |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ✅ yes |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ✅ yes |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for set-poor-network-msgs                                                                                                                              | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default 'home' directory will be used                                                                         | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ❌ ?   |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ✅ yes |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `--title string`              | The title of the proposal.                                                                                                                                  | ✅ yes |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx customgov proposal set-poor-network-msgs --help

                $ %s tx customgov proposal set-poor-network-msgs XXXX,YYY --from=<key_or_address>

                All the message types supported could be added here
                        create-role
                        assign-role
                        remove-role
                        ...

Usage:
  sekaid tx customgov proposal set-poor-network-msgs <messages> [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --description string       The description of the proposal, it can be a url, some text, etc.
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for set-poor-network-msgs
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
      --title string             The title of the proposal.
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

#TODO haven't tested

🟨  
🟨  
🟨

[Return to top](#sekai)

###### 21.7.7.7 set-proposal-durations-proposal

Create a proposal to set batch proposal durations.

Usage:
```
sekaid tx customgov proposal set-proposal-durations-proposal [proposal_types] [durations] [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ❌ ?   |
| `--description string`        | The description of the proposal, it can be a url, some text, etc.                                                                                           | ✅ yes |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ✅ yes |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ✅ yes |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for set-proposal-durations-proposal                                                                                                                    | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default 'home' directory will be used                                                                         | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ❌ ?   |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ✅ yes |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `--title string`              | The title of the proposal.                                                                                                                                  | ✅ yes |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx customgov proposal set-proposal-durations-proposal --help
Create a proposal to set batch proposal durations.

Usage:
  sekaid tx customgov proposal set-proposal-durations-proposal [proposal_types] [durations] [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --description string       The description of the proposal, it can be a url, some text, etc.
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for set-proposal-durations-proposal
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
      --title string             The title of the proposal.
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

> You can find current proposal durations here: `sekaid q customgov all-proposal-durations -o json | jq`
>
> You can provide key-value pairs like this: ` AssignRoleToAccount,BlacklistAccountPermission 300,310`

```
sekaid tx customgov proposal set-proposal-durations-proposal AssignRoleToAccount 300 --title="Test proposal for docs" --description="Test proposal for docs" --from=kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --chain-id=localnet-4 --home=/root/.sekai --keyring-backend=test --fees=100ukex --yes --output=json | jq
{
  "height": "0",
  "txhash": "43DA0E8CAF8CDA944D8C6BC2F182E99EBC144187B07F16609CEB7C68535306EB",
  "codespace": "",
  "code": 0,
  "data": "",
  "raw_log": "[]",
  "logs": [],
  "info": "",
  "gas_wanted": "0",
  "gas_used": "0",
  "tx": null,
  "timestamp": "",
  "events": []
}
```

<details>
  <summary>Check tx execution</summary>

  ```
  sekaid q tx 43DA0E8CAF8CDA944D8C6BC2F182E99EBC144187B07F16609CEB7C68535306EB -o json | jq
  ```

  ```json
  {
    "height": "143174",
    "txhash": "43DA0E8CAF8CDA944D8C6BC2F182E99EBC144187B07F16609CEB7C68535306EB",
    "codespace": "",
    "code": 0,
    "data": "0A210A1B2F6B6972612E676F762E4D73675375626D697450726F706F73616C1202081B",
    "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/kira.gov.MsgSubmitProposal\"}]},{\"type\":\"submit_proposal\",\"attributes\":[{\"key\":\"proposal_id\",\"value\":\"27\"},{\"key\":\"proposal_type\",\"value\":\"SetProposalDurationsProposal\"},{\"key\":\"proposal_content\",\"value\":\"proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\\ntitle: Test proposal for docs\\ndescription: Test proposal for docs\\ncontent:\\n  typeurl: /kira.gov.SetProposalDurationsProposal\\n  value:\\n  - 10\\n  - 19\\n  - 65\\n  - 115\\n  - 115\\n  - 105\\n  - 103\\n  - 110\\n  - 82\\n  - 111\\n  - 108\\n  - 101\\n  - 84\\n  - 111\\n  - 65\\n  - 99\\n  - 99\\n  - 111\\n  - 117\\n  - 110\\n  - 116\\n  - 18\\n  - 2\\n  - 172\\n  - 2\\n  xxx_nounkeyedliteral: {}\\n  xxx_unrecognized: []\\n  xxx_sizecache: 0\\n\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/kira.gov.MsgSubmitProposal"
              }
            ]
          },
          {
            "type": "submit_proposal",
            "attributes": [
              {
                "key": "proposal_id",
                "value": "27"
              },
              {
                "key": "proposal_type",
                "value": "SetProposalDurationsProposal"
              },
              {
                "key": "proposal_content",
                "value": "proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\ntitle: Test proposal for docs\ndescription: Test proposal for docs\ncontent:\n  typeurl: /kira.gov.SetProposalDurationsProposal\n  value:\n  - 10\n  - 19\n  - 65\n  - 115\n  - 115\n  - 105\n  - 103\n  - 110\n  - 82\n  - 111\n  - 108\n  - 101\n  - 84\n  - 111\n  - 65\n  - 99\n  - 99\n  - 111\n  - 117\n  - 110\n  - 116\n  - 18\n  - 2\n  - 172\n  - 2\n  xxx_nounkeyedliteral: {}\n  xxx_unrecognized: []\n  xxx_sizecache: 0\n"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "0",
    "gas_used": "0",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/kira.gov.MsgSubmitProposal",
            "proposer": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
            "title": "Test proposal for docs",
            "description": "Test proposal for docs",
            "content": {
              "@type": "/kira.gov.SetProposalDurationsProposal",
              "typeof_proposals": [
                "AssignRoleToAccount"
              ],
              "proposal_durations": [
                "300"
              ]
            }
          }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
      },
      "auth_info": {
        "signer_infos": [
          {
            "public_key": {
              "@type": "/cosmos.crypto.secp256k1.PubKey",
              "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "117"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "ukex",
              "amount": "100"
            }
          ],
          "gas_limit": "200000",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
        "YbR5agxdjERYOsqh7t5tgusY0YeYC7RA1OCz14TXGoQF6I2t8v9P9wQdlR67ZSW4M/x0JI4yQJXy1oX35DFYwg=="
      ]
    },
    "timestamp": "2023-06-12T14:37:29Z",
    "events": [
      {
        "type": "tx",
        "attributes": [
          {
            "key": "YWNjX3NlcQ==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NC8xMTc=",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "WWJSNWFneGRqRVJZT3NxaDd0NXRndXNZMFllWUM3UkExT0N6MTRUWEdvUUY2STJ0OHY5UDl3UWRsUjY3WlNXNE0veDBKSTR5UUpYeTFvWDM1REZZd2c9PQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "ZmVl",
            "value": "MTAwdWtleA==",
            "index": true
          },
          {
            "key": "ZmVlX3BheWVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "L2tpcmEuZ292Lk1zZ1N1Ym1pdFByb3Bvc2Fs",
            "index": true
          }
        ]
      },
      {
        "type": "submit_proposal",
        "attributes": [
          {
            "key": "cHJvcG9zYWxfaWQ=",
            "value": "Mjc=",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfdHlwZQ==",
            "value": "U2V0UHJvcG9zYWxEdXJhdGlvbnNQcm9wb3NhbA==",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfY29udGVudA==",
            "value": "cHJvcG9zZXI6IGtpcmExdm13ZGd3NDI2YWo5ZngzM2ZxdXNtdGc2cjY1eXl1Y214NnJkdDQKdGl0bGU6IFRlc3QgcHJvcG9zYWwgZm9yIGRvY3MKZGVzY3JpcHRpb246IFRlc3QgcHJvcG9zYWwgZm9yIGRvY3MKY29udGVudDoKICB0eXBldXJsOiAva2lyYS5nb3YuU2V0UHJvcG9zYWxEdXJhdGlvbnNQcm9wb3NhbAogIHZhbHVlOgogIC0gMTAKICAtIDE5CiAgLSA2NQogIC0gMTE1CiAgLSAxMTUKICAtIDEwNQogIC0gMTAzCiAgLSAxMTAKICAtIDgyCiAgLSAxMTEKICAtIDEwOAogIC0gMTAxCiAgLSA4NAogIC0gMTExCiAgLSA2NQogIC0gOTkKICAtIDk5CiAgLSAxMTEKICAtIDExNwogIC0gMTEwCiAgLSAxMTYKICAtIDE4CiAgLSAyCiAgLSAxNzIKICAtIDIKICB4eHhfbm91bmtleWVkbGl0ZXJhbDoge30KICB4eHhfdW5yZWNvZ25pemVkOiBbXQogIHh4eF9zaXplY2FjaGU6IDAK",
            "index": true
          }
        ]
      }
    ]
  }
  ```
</details>

<details>
  <summary>Check proposal</summary>

  ```
  sekaid q customgov proposal 27 -o json | jq
  ```

  ```json
  {
    "proposal_id": "27",
    "title": "Test proposal for docs",
    "description": "Test proposal for docs",
    "content": {
      "@type": "/kira.gov.SetProposalDurationsProposal",
      "typeof_proposals": [
        "AssignRoleToAccount"
      ],
      "proposal_durations": [
        "300"
      ]
    },
    "submit_time": "2023-06-12T14:37:29.848780313Z",
    "voting_end_time": "2023-06-12T14:43:29.848780313Z",
    "enactment_end_time": "2023-06-12T14:48:29.848780313Z",
    "min_voting_end_block_height": "143176",
    "min_enactment_end_block_height": "143177",
    "result": "VOTE_PENDING",
    "exec_result": ""
  }
  ```
</details>

[Return to top](#sekai)

###### 21.7.7.8 upsert-data-registry

Create a proposal to upsert a key in the data registry.

Usage:
```
sekaid tx customgov proposal upsert-data-registry [key] [hash] [reference] [encoding] [size] [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ❌ ?   |
| `--description string`        | The description of the proposal, it can be a url, some text, etc.                                                                                           | ✅ yes |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ✅ yes |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ✅ yes |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ❌ ?   |
| `-h, --help`                  | help for upsert-data-registry                                                                                                                               | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default `'home'` directory will be used                                                                       | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ❌ ?   |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ✅ yes |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `--title string`              | The title of the proposal.                                                                                                                                  | ✅ yes |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx customgov proposal upsert-data-registry --help
Create a proposal to upsert a key in the data registry

Usage:
  sekaid tx customgov proposal upsert-data-registry [key] [hash] [reference] [encoding] [size] [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --description string       The description of the proposal, it can be a url, some text, etc.
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for upsert-data-registry
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
      --title string             The title of the proposal.
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid tx customgov proposal upsert-data-registry icon2 891bd9d3b2ee0c6eed43a8129b096bebc7e5ae517d0b855b2116a3205211fe21 https://kira-network.s3-eu-west-1.amazonaws.com/assets/img/tokens/kex.svg picture 1597 --title="Upserting date registry key 'icon'" --description="Assign value '<url>' to key 'icon'" --from=kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --chain-id=localnet-4 --keyring-backend=test --fees=100ukex --yes --log_format=json --broadcast-mode=async --home=/root/.sekai --output=json | jq
{
  "height": "0",
  "txhash": "6B7BCD4B3B2B93D83F477E522EAA5FEA998C6BD68970CCAD746A07E153E43AC4",
  "codespace": "",
  "code": 0,
  "data": "",
  "raw_log": "",
  "logs": [],
  "info": "",
  "gas_wanted": "0",
  "gas_used": "0",
  "tx": null,
  "timestamp": "",
  "events": []
}
```

<details>
  <summary>Check the transaction</summary>

  ```
  /# sekaid query tx 6B7BCD4B3B2B93D83F477E522EAA5FEA998C6BD68970CCAD746A07E153E43AC4 -o json | jq
  ```

  ```json 
  {
    "height": "93606",
    "txhash": "6B7BCD4B3B2B93D83F477E522EAA5FEA998C6BD68970CCAD746A07E153E43AC4",
    "codespace": "",
    "code": 0,
    "data": "0A210A1B2F6B6972612E676F762E4D73675375626D697450726F706F73616C1202080B",
    "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/kira.gov.MsgSubmitProposal\"}]},{\"type\":\"submit_proposal\",\"attributes\":[{\"key\":\"proposal_id\",\"value\":\"11\"},{\"key\":\"proposal_type\",\"value\":\"UpsertDataRegistry\"},{\"key\":\"proposal_content\",\"value\":\"proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\\ntitle: Upserting date registry key 'icon'\\ndescription: Assign value '\\u003curl\\u003e' to key 'icon'\\ncontent:\\n  typeurl: /kira.gov.UpsertDataRegistryProposal\\n  value:\\n  - 10\\n  - 5\\n  - 105\\n  - 99\\n  - 111\\n  - 110\\n  - 50\\n  - 18\\n  - 64\\n  - 56\\n  - 57\\n  - 49\\n  - 98\\n  - 100\\n  - 57\\n  - 100\\n  - 51\\n  - 98\\n  - 50\\n  - 101\\n  - 101\\n  - 48\\n  - 99\\n  - 54\\n  - 101\\n  - 101\\n  - 100\\n  - 52\\n  - 51\\n  - 97\\n  - 56\\n  - 49\\n  - 50\\n  - 57\\n  - 98\\n  - 48\\n  - 57\\n  - 54\\n  - 98\\n  - 101\\n  - 98\\n  - 99\\n  - 55\\n  - 101\\n  - 53\\n  - 97\\n  - 101\\n  - 53\\n  - 49\\n  - 55\\n  - 100\\n  - 48\\n  - 98\\n  - 56\\n  - 53\\n  - 53\\n  - 98\\n  - 50\\n  - 49\\n  - 49\\n  - 54\\n  - 97\\n  - 51\\n  - 50\\n  - 48\\n  - 53\\n  - 50\\n  - 49\\n  - 49\\n  - 102\\n  - 101\\n  - 50\\n  - 49\\n  - 26\\n  - 73\\n  - 104\\n  - 116\\n  - 116\\n  - 112\\n  - 115\\n  - 58\\n  - 47\\n  - 47\\n  - 107\\n  - 105\\n  - 114\\n  - 97\\n  - 45\\n  - 110\\n  - 101\\n  - 116\\n  - 119\\n  - 111\\n  - 114\\n  - 107\\n  - 46\\n  - 115\\n  - 51\\n  - 45\\n  - 101\\n  - 117\\n  - 45\\n  - 119\\n  - 101\\n  - 115\\n  - 116\\n  - 45\\n  - 49\\n  - 46\\n  - 97\\n  - 109\\n  - 97\\n  - 122\\n  - 111\\n  - 110\\n  - 97\\n  - 119\\n  - 115\\n  - 46\\n  - 99\\n  - 111\\n  - 109\\n  - 47\\n  - 97\\n  - 115\\n  - 115\\n  - 101\\n  - 116\\n  - 115\\n  - 47\\n  - 105\\n  - 109\\n  - 103\\n  - 47\\n  - 116\\n  - 111\\n  - 107\\n  - 101\\n  - 110\\n  - 115\\n  - 47\\n  - 107\\n  - 101\\n  - 120\\n  - 46\\n  - 115\\n  - 118\\n  - 103\\n  - 34\\n  - 7\\n  - 112\\n  - 105\\n  - 99\\n  - 116\\n  - 117\\n  - 114\\n  - 101\\n  - 40\\n  - 189\\n  - 12\\n  xxx_nounkeyedliteral: {}\\n  xxx_unrecognized: []\\n  xxx_sizecache: 0\\n\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/kira.gov.MsgSubmitProposal"
              }
            ]
          },
          {
            "type": "submit_proposal",
            "attributes": [
              {
                "key": "proposal_id",
                "value": "11"
              },
              {
                "key": "proposal_type",
                "value": "UpsertDataRegistry"
              },
              {
                "key": "proposal_content",
                "value": "proposer: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\ntitle: Upserting date registry key 'icon'\ndescription: Assign value '<url>' to key 'icon'\ncontent:\n  typeurl: /kira.gov.UpsertDataRegistryProposal\n  value:\n  - 10\n  - 5\n  - 105\n  - 99\n  - 111\n  - 110\n  - 50\n  - 18\n  - 64\n  - 56\n  - 57\n  - 49\n  - 98\n  - 100\n  - 57\n  - 100\n  - 51\n  - 98\n  - 50\n  - 101\n  - 101\n  - 48\n  - 99\n  - 54\n  - 101\n  - 101\n  - 100\n  - 52\n  - 51\n  - 97\n  - 56\n  - 49\n  - 50\n  - 57\n  - 98\n  - 48\n  - 57\n  - 54\n  - 98\n  - 101\n  - 98\n  - 99\n  - 55\n  - 101\n  - 53\n  - 97\n  - 101\n  - 53\n  - 49\n  - 55\n  - 100\n  - 48\n  - 98\n  - 56\n  - 53\n  - 53\n  - 98\n  - 50\n  - 49\n  - 49\n  - 54\n  - 97\n  - 51\n  - 50\n  - 48\n  - 53\n  - 50\n  - 49\n  - 49\n  - 102\n  - 101\n  - 50\n  - 49\n  - 26\n  - 73\n  - 104\n  - 116\n  - 116\n  - 112\n  - 115\n  - 58\n  - 47\n  - 47\n  - 107\n  - 105\n  - 114\n  - 97\n  - 45\n  - 110\n  - 101\n  - 116\n  - 119\n  - 111\n  - 114\n  - 107\n  - 46\n  - 115\n  - 51\n  - 45\n  - 101\n  - 117\n  - 45\n  - 119\n  - 101\n  - 115\n  - 116\n  - 45\n  - 49\n  - 46\n  - 97\n  - 109\n  - 97\n  - 122\n  - 111\n  - 110\n  - 97\n  - 119\n  - 115\n  - 46\n  - 99\n  - 111\n  - 109\n  - 47\n  - 97\n  - 115\n  - 115\n  - 101\n  - 116\n  - 115\n  - 47\n  - 105\n  - 109\n  - 103\n  - 47\n  - 116\n  - 111\n  - 107\n  - 101\n  - 110\n  - 115\n  - 47\n  - 107\n  - 101\n  - 120\n  - 46\n  - 115\n  - 118\n  - 103\n  - 34\n  - 7\n  - 112\n  - 105\n  - 99\n  - 116\n  - 117\n  - 114\n  - 101\n  - 40\n  - 189\n  - 12\n  xxx_nounkeyedliteral: {}\n  xxx_unrecognized: []\n  xxx_sizecache: 0\n"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "0",
    "gas_used": "0",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/kira.gov.MsgSubmitProposal",
            "proposer": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
            "title": "Upserting date registry key 'icon'",
            "description": "Assign value '<url>' to key 'icon'",
            "content": {
              "@type": "/kira.gov.UpsertDataRegistryProposal",
              "key": "icon2",
              "hash": "891bd9d3b2ee0c6eed43a8129b096bebc7e5ae517d0b855b2116a3205211fe21",
              "reference": "https://kira-network.s3-eu-west-1.amazonaws.com/assets/img/tokens/kex.svg",
              "encoding": "picture",
              "size": "1597"
            }
          }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
      },
      "auth_info": {
        "signer_infos": [
          {
            "public_key": {
              "@type": "/cosmos.crypto.secp256k1.PubKey",
              "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "85"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "ukex",
              "amount": "100"
            }
          ],
          "gas_limit": "200000",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
        "nIXu3GR+RioFsgnNfOIlyOxKSNFfqE+EkdCDyF7rzxtXq8X+jcU22Iuh1rrU3IYzzP3fCeIL8kkuuDf2q/0fIg=="
      ]
    },
    "timestamp": "2023-06-06T16:37:02Z",
    "events": [
      {
        "type": "tx",
        "attributes": [
          {
            "key": "YWNjX3NlcQ==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NC84NQ==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "bklYdTNHUitSaW9Gc2duTmZPSWx5T3hLU05GZnFFK0VrZENEeUY3cnp4dFhxOFgramNVMjJJdWgxcnJVM0lZenpQM2ZDZUlMOGtrdXVEZjJxLzBmSWc9PQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "ZmVl",
            "value": "MTAwdWtleA==",
            "index": true
          },
          {
            "key": "ZmVlX3BheWVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "L2tpcmEuZ292Lk1zZ1N1Ym1pdFByb3Bvc2Fs",
            "index": true
          }
        ]
      },
      {
        "type": "submit_proposal",
        "attributes": [
          {
            "key": "cHJvcG9zYWxfaWQ=",
            "value": "MTE=",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfdHlwZQ==",
            "value": "VXBzZXJ0RGF0YVJlZ2lzdHJ5",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfY29udGVudA==",
            "value": "cHJvcG9zZXI6IGtpcmExdm13ZGd3NDI2YWo5ZngzM2ZxdXNtdGc2cjY1eXl1Y214NnJkdDQKdGl0bGU6IFVwc2VydGluZyBkYXRlIHJlZ2lzdHJ5IGtleSAnaWNvbicKZGVzY3JpcHRpb246IEFzc2lnbiB2YWx1ZSAnPHVybD4nIHRvIGtleSAnaWNvbicKY29udGVudDoKICB0eXBldXJsOiAva2lyYS5nb3YuVXBzZXJ0RGF0YVJlZ2lzdHJ5UHJvcG9zYWwKICB2YWx1ZToKICAtIDEwCiAgLSA1CiAgLSAxMDUKICAtIDk5CiAgLSAxMTEKICAtIDExMAogIC0gNTAKICAtIDE4CiAgLSA2NAogIC0gNTYKICAtIDU3CiAgLSA0OQogIC0gOTgKICAtIDEwMAogIC0gNTcKICAtIDEwMAogIC0gNTEKICAtIDk4CiAgLSA1MAogIC0gMTAxCiAgLSAxMDEKICAtIDQ4CiAgLSA5OQogIC0gNTQKICAtIDEwMQogIC0gMTAxCiAgLSAxMDAKICAtIDUyCiAgLSA1MQogIC0gOTcKICAtIDU2CiAgLSA0OQogIC0gNTAKICAtIDU3CiAgLSA5OAogIC0gNDgKICAtIDU3CiAgLSA1NAogIC0gOTgKICAtIDEwMQogIC0gOTgKICAtIDk5CiAgLSA1NQogIC0gMTAxCiAgLSA1MwogIC0gOTcKICAtIDEwMQogIC0gNTMKICAtIDQ5CiAgLSA1NQogIC0gMTAwCiAgLSA0OAogIC0gOTgKICAtIDU2CiAgLSA1MwogIC0gNTMKICAtIDk4CiAgLSA1MAogIC0gNDkKICAtIDQ5CiAgLSA1NAogIC0gOTcKICAtIDUxCiAgLSA1MAogIC0gNDgKICAtIDUzCiAgLSA1MAogIC0gNDkKICAtIDQ5CiAgLSAxMDIKICAtIDEwMQogIC0gNTAKICAtIDQ5CiAgLSAyNgogIC0gNzMKICAtIDEwNAogIC0gMTE2CiAgLSAxMTYKICAtIDExMgogIC0gMTE1CiAgLSA1OAogIC0gNDcKICAtIDQ3CiAgLSAxMDcKICAtIDEwNQogIC0gMTE0CiAgLSA5NwogIC0gNDUKICAtIDExMAogIC0gMTAxCiAgLSAxMTYKICAtIDExOQogIC0gMTExCiAgLSAxMTQKICAtIDEwNwogIC0gNDYKICAtIDExNQogIC0gNTEKICAtIDQ1CiAgLSAxMDEKICAtIDExNwogIC0gNDUKICAtIDExOQogIC0gMTAxCiAgLSAxMTUKICAtIDExNgogIC0gNDUKICAtIDQ5CiAgLSA0NgogIC0gOTcKICAtIDEwOQogIC0gOTcKICAtIDEyMgogIC0gMTExCiAgLSAxMTAKICAtIDk3CiAgLSAxMTkKICAtIDExNQogIC0gNDYKICAtIDk5CiAgLSAxMTEKICAtIDEwOQogIC0gNDcKICAtIDk3CiAgLSAxMTUKICAtIDExNQogIC0gMTAxCiAgLSAxMTYKICAtIDExNQogIC0gNDcKICAtIDEwNQogIC0gMTA5CiAgLSAxMDMKICAtIDQ3CiAgLSAxMTYKICAtIDExMQogIC0gMTA3CiAgLSAxMDEKICAtIDExMAogIC0gMTE1CiAgLSA0NwogIC0gMTA3CiAgLSAxMDEKICAtIDEyMAogIC0gNDYKICAtIDExNQogIC0gMTE4CiAgLSAxMDMKICAtIDM0CiAgLSA3CiAgLSAxMTIKICAtIDEwNQogIC0gOTkKICAtIDExNgogIC0gMTE3CiAgLSAxMTQKICAtIDEwMQogIC0gNDAKICAtIDE4OQogIC0gMTIKICB4eHhfbm91bmtleWVkbGl0ZXJhbDoge30KICB4eHhfdW5yZWNvZ25pemVkOiBbXQogIHh4eF9zaXplY2FjaGU6IDAK",
            "index": true
          }
        ]
      }
    ]
  }
  ```

  > OR find all transactions for 
  ```
  sekaid query txs --events=submit_proposal.proposal_type=UpsertDataRegistry --output=json | jq
  ```

  __VERY LONG JSON with all txs__
  ```json
  {
    "total_count": "4",
    "count": "4",
    "page_number": "1",
    "page_total": "1",
    "limit": "30",
    "txs": [
      {...}, 
      {...}, 
      {...}, 
      {...}
    ]
  }
  ```
</details>

<details>
  <summary>Check the vote</summary>

  ```
  sekaid query customgov proposal 11 -o json | jq
  ```

  ```json
  {
    "proposal_id": "11",
    "title": "Upserting date registry key 'icon'",
    "description": "Assign value '<url>' to key 'icon'",
    "content": {
      "@type": "/kira.gov.UpsertDataRegistryProposal",
      "key": "icon2",
      "hash": "891bd9d3b2ee0c6eed43a8129b096bebc7e5ae517d0b855b2116a3205211fe21",
      "reference": "https://kira-network.s3-eu-west-1.amazonaws.com/assets/img/tokens/kex.svg",
      "encoding": "picture",
      "size": "1597"
    },
    "submit_time": "2023-06-06T16:37:02.959659272Z",
    "voting_end_time": "2023-06-06T16:43:02.959659272Z",
    "enactment_end_time": "2023-06-06T16:48:02.959659272Z",
    "min_voting_end_block_height": "93608",
    "min_enactment_end_block_height": "93609",
    "result": "VOTE_PENDING",
    "exec_result": ""
  }
  ```
</details>

<details>
  <summary>Check data registry (after successful vote)</summary>

  List of all data registries:
  ```
  sekaid query customgov data-registry-keys --output=json | jq
  ```

  ```json
  {
    "keys": [
      "icon" // it's a old passed value!
    ],
    "pagination": {
      "next_key": null,
      "total": "0"
    }
  }
  ```

  Only specific value (_there is an example of old same registry_):
  ```
  /# sekaid query customgov data-registry icon -o json | jq
  ```

  ```json
  {
    "data": {
      "hash": "891bd9d3b2ee0c6eed43a8129b096bebc7e5ae517d0b855b2116a3205211fe21",
      "reference": "https://kira-network.s3-eu-west-1.amazonaws.com/assets/img/tokens/kex.svg",
      "encoding": "picture",
      "size": "1597"
    }
  }
  ```
</details>

[Return to top](#sekai)

###### 21.7.7.9 vote

Vote a proposal.

Usage:
```
sekaid tx customgov proposal vote proposal-id vote-option [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ✅ yes |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ✅ yes |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ✅ yes |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for vote                                                                                                                                               | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default `'home'` directory will be used                                                                       | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ✅ yes |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ✅ yes |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--slash string`              | slash value on the proposal (default `"0.01"`)                                                                                                              | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx customgov proposal vote --help
Vote a proposal.

Usage:
  sekaid tx customgov proposal vote proposal-id vote-option [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for vote
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --slash string             slash value on the proposal (default "0.01")
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

> `vote-option` can be
>
> - 0: Unspecified
> - 1: Yes
> - 2: Abstain
> - 3: No
> - 4: Veto

```
sekaid tx customgov proposal vote 13 1 --from=kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --keyring-backend=test --home=/root/.sekai --chain-id=localnet-4 --fees=100ukex --yes --output=json | jq
{
  "height": "0",
  "txhash": "10C68B927E3A1BDED0EDD8112AB96DB97CC31FA5B91C88566C960D813B64D073",
  "codespace": "",
  "code": 0,
  "data": "",
  "raw_log": "[]",
  "logs": [],
  "info": "",
  "gas_wanted": "0",
  "gas_used": "0",
  "tx": null,
  "timestamp": "",
  "events": []
}
```

<details>
  <summary>Check tx execution</summary>

  ```
  sekaid q tx 10C68B927E3A1BDED0EDD8112AB96DB97CC31FA5B91C88566C960D813B64D073 -o json | jq
  ```

  ```json
  {
    "height": "141342",
    "txhash": "10C68B927E3A1BDED0EDD8112AB96DB97CC31FA5B91C88566C960D813B64D073",
    "codespace": "",
    "code": 0,
    "data": "0A1B0A192F6B6972612E676F762E4D7367566F746550726F706F73616C",
    "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/kira.gov.MsgVoteProposal\"}]},{\"type\":\"proposal_vote\",\"attributes\":[{\"key\":\"proposal_id\",\"value\":\"13\"},{\"key\":\"voter\",\"value\":\"kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\"},{\"key\":\"option\",\"value\":\"VOTE_OPTION_YES\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/kira.gov.MsgVoteProposal"
              }
            ]
          },
          {
            "type": "proposal_vote",
            "attributes": [
              {
                "key": "proposal_id",
                "value": "13"
              },
              {
                "key": "voter",
                "value": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4"
              },
              {
                "key": "option",
                "value": "VOTE_OPTION_YES"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "0",
    "gas_used": "0",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/kira.gov.MsgVoteProposal",
            "proposal_id": "13",
            "voter": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
            "option": "VOTE_OPTION_YES",
            "slash": "0.010000000000000000"
          }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
      },
      "auth_info": {
        "signer_infos": [
          {
            "public_key": {
              "@type": "/cosmos.crypto.secp256k1.PubKey",
              "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "96"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "ukex",
              "amount": "100"
            }
          ],
          "gas_limit": "200000",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
        "pSzlgYtk+hG0zRlF59DHl/rS928uIGAX+w5m2rFCF2dMsSCw9q28te49FLDzr7IDmrgoC3euUB/Qo99N+ZQH7w=="
      ]
    },
    "timestamp": "2023-06-12T09:22:30Z",
    "events": [
      {
        "type": "tx",
        "attributes": [
          {
            "key": "YWNjX3NlcQ==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NC85Ng==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "cFN6bGdZdGsraEcwelJsRjU5REhsL3JTOTI4dUlHQVgrdzVtMnJGQ0YyZE1zU0N3OXEyOHRlNDlGTER6cjdJRG1yZ29DM2V1VUIvUW85OU4rWlFIN3c9PQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "ZmVl",
            "value": "MTAwdWtleA==",
            "index": true
          },
          {
            "key": "ZmVlX3BheWVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "L2tpcmEuZ292Lk1zZ1ZvdGVQcm9wb3NhbA==",
            "index": true
          }
        ]
      },
      {
        "type": "proposal_vote",
        "attributes": [
          {
            "key": "cHJvcG9zYWxfaWQ=",
            "value": "MTM=",
            "index": true
          },
          {
            "key": "dm90ZXI=",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "b3B0aW9u",
            "value": "Vk9URV9PUFRJT05fWUVT",
            "index": true
          }
        ]
      }
    ]
  }
  ```
</details>

<details>
  <summary>Check vote</summary>

  ```
  sekaid q customgov votes 13 -o json | jq
  ```

  ```json
  {
    "votes": [
      {
        "proposal_id": "13",
        "voter": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
        "option": "VOTE_OPTION_YES",
        "slash": "0.010000000000000000"
      }
    ]
  }
  ```

  ```
  sekaid q customgov vote 13 kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 -o json | jq
  ```

  ```json
  {
    "proposal_id": "13",
    "voter": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
    "option": "VOTE_OPTION_YES",
    "slash": "0.010000000000000000"
  }
  ```
</details>

[Return to top](#sekai)

##### 21.7.8 register-identity-records

Submit a transaction to create an identity record.

Usage:
```
sekaid tx customgov register-identity-records [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ✅ yes |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ✅ yes |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ✅ yes |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for register-identity-records                                                                                                                          | ✅ yes |
| `--infos-file string`         | The infos file for identity request.                                                                                                                        | ✅ yes |
| `--infos-json string`         | The infos json for identity request.                                                                                                                        | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default `'home'` directory will be used                                                                       | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ✅ yes |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ✅ yes |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx customgov register-identity-records --help
Submit a transaction to create an identity record.

Usage:
  sekaid tx customgov register-identity-records [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for register-identity-records
      --infos-file string        The infos file for identity request.
      --infos-json string        The infos json for identity request.
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid tx customgov register-identity-records --from=kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --home=/root/.sekai --chain-id=localnet-4 --keyring-backend=test --infos-json='{"text":"Testing item regisrty for docs"}' --yes --fees=99ukex --output=json | jq
{
  "height": "0",
  "txhash": "38C6F2304E9FA5CB80CA4FEDF89E375FDE034E18B299DB068A725397E3D1912D",
  "codespace": "",
  "code": 0,
  "data": "",
  "raw_log": "[]",
  "logs": [],
  "info": "",
  "gas_wanted": "0",
  "gas_used": "0",
  "tx": null,
  "timestamp": "",
  "events": []
}
```

<details>
  <summary>Check transaction execution</summary>

  ```
  sekaid q tx 38C6F2304E9FA5CB80CA4FEDF89E375FDE034E18B299DB068A725397E3D1912D -o json | jq
  ```

  ```json
  {
    "height": "108896",
    "txhash": "38C6F2304E9FA5CB80CA4FEDF89E375FDE034E18B299DB068A725397E3D1912D",
    "codespace": "",
    "code": 0,
    "data": "0A260A242F6B6972612E676F762E4D736752656769737465724964656E746974795265636F726473",
    "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/kira.gov.MsgRegisterIdentityRecords\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/kira.gov.MsgRegisterIdentityRecords"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "0",
    "gas_used": "0",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/kira.gov.MsgRegisterIdentityRecords",
            "address": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
            "infos": [
              {
                "key": "text",
                "info": "Testing item regisrty for docs"
              }
            ]
          }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
      },
      "auth_info": {
        "signer_infos": [
          {
            "public_key": {
              "@type": "/cosmos.crypto.secp256k1.PubKey",
              "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "86"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "ukex",
              "amount": "99"
            }
          ],
          "gas_limit": "200000",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
        "6uSHGVEc3MVhV7THxqKvGeQptF1co9FRiySMAti17zxKW0yxAGdx2cbZC4/SeKoLXWdc9kNaI/lZwHNhQLDkUw=="
      ]
    },
    "timestamp": "2023-06-08T12:25:11Z",
    "events": [
      {
        "type": "tx",
        "attributes": [
          {
            "key": "YWNjX3NlcQ==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NC84Ng==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "NnVTSEdWRWMzTVZoVjdUSHhxS3ZHZVFwdEYxY285RlJpeVNNQXRpMTd6eEtXMHl4QUdkeDJjYlpDNC9TZUtvTFhXZGM5a05hSS9sWndITmhRTERrVXc9PQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "OTl1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "OTl1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "OTl1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "ZmVl",
            "value": "OTl1a2V4",
            "index": true
          },
          {
            "key": "ZmVlX3BheWVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "L2tpcmEuZ292Lk1zZ1JlZ2lzdGVySWRlbnRpdHlSZWNvcmRz",
            "index": true
          }
        ]
      }
    ]
  }
  ```
</details>

<details>
  <summary>Check identity registry</summary>

  List of all identity records:
  ```
  sekaid q customgov identity-records -o json | jq
  ```

  ```json
  {
    "records": [
      {
        "id": "1",
        "address": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
        "key": "moniker",
        "value": "GENESIS VALIDATOR",
        "date": "2023-05-25T12:56:59.866789362Z",
        "verifiers": []
      },
      // . . .
      {
        "id": "19",
        "address": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
        "key": "text",
        "value": "Testing item regisrty for docs",
        "date": "2023-06-08T12:25:11.014802004Z",
        "verifiers": []
      }
    ],
    "pagination": {
      "next_key": null,
      "total": "0"
    }
  }
  ```

  Identity record by ID:
  ```
  sekaid query customgov identity-record 19 -o json | jq
  ```

  ```json
  {
    "record": {
      "id": "19",
      "address": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
      "key": "text",
      "value": "Testing item regisrty for docs",
      "date": "2023-06-08T12:25:11.014802004Z",
      "verifiers": []
    }
  }
  ```

  Identity records by address:
  ```
  sekaid query customgov identity-records-by-addr kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 -o json | jq
  ```

  Sorted by `key` output:
  ```json
{
  "records": [
    // . . .
    {
      "id": "19",
      "address": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
      "key": "text",
      "value": "Testing item regisrty for docs",
      "date": "2023-06-08T12:25:11.014802004Z",
      "verifiers": []
    },
    // . . .
  ]
}
  ```
</details>

[Return to top](#sekai)

##### 21.7.9 request-identity-record-verify

Submit a transaction to request an identity verify record.

Usage:
```
sekaid tx customgov request-identity-record-verify [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ✅ yes |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ✅ yes |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ❌ ?   |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for request-identity-record-verify                                                                                                                     | ✅ yes |
| `--infos-file string`         | The infos file for identity request.                                                                                                                        | ✅ yes |
| `--infos-json string`         | The infos json for identity request.                                                                                                                        | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default `'home'` directory will be used                                                                       | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ✅ yes |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ✅ yes |
| `--record-ids string`         | Concatenated identity record ids array. e.g. `1,2`                                                                                                          | ✅ yes |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `--tip string`                | The tip to be given to the verifier.                                                                                                                        | ✅ yes |
| `--verifier string`           | The verifier of the record ids                                                                                                                              | ✅ yes |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ✅ yes |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx customgov request-identity-record-verify --help
Submit a transaction to request an identity verify record.

Usage:
  sekaid tx customgov request-identity-record-verify [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for request-identity-record-verify
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
      --record-ids string        Concatenated identity record ids array. e.g. 1,2
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
      --tip string               The tip to be given to the verifier.
      --verifier string          The verifier of the record ids
  -y, --yes                      Skip tx broadcasting prompt confirmation
```

> Make sure that `tip` is higher that the minimum value configured by the network
> 
> `min_identity_approval_tip` value 
> 
> from `sekaid query customgov network-properties --output=json | jq`

```
sekaid tx customgov request-identity-record-verify --from=kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --verifier=kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt --keyring-backend=test --chain-id=localnet-4 --home=/root/.sekai --record-ids=19 --fees=99ukex --tip=200ukex --yes -o json | jq
{
  "height": "0",
  "txhash": "5E482557FCDFF9821A5314DABDC2A3F2360265BC30D26AD0100E82227E429CF0",
  "codespace": "",
  "code": 0,
  "data": "",
  "raw_log": "[]",
  "logs": [],
  "info": "",
  "gas_wanted": "0",
  "gas_used": "0",
  "tx": null,
  "timestamp": "",
  "events": []
}
```

<details>
  <summary>Chech tx execution</summary>

  ```
  sekaid q tx 5E482557FCDFF9821A5314DABDC2A3F2360265BC30D26AD0100E82227E429CF0 -o json | jq
  ```

  ```json
  {
    "height": "109062",
    "txhash": "5E482557FCDFF9821A5314DABDC2A3F2360265BC30D26AD0100E82227E429CF0",
    "codespace": "",
    "code": 0,
    "data": "0A2F0A292F6B6972612E676F762E4D7367526571756573744964656E746974795265636F72647356657269667912020804",
    "raw_log": "[{\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"kira18q47tkrtrthf72xyscens2kv58ufpqtyxvxwuv\"},{\"key\":\"amount\",\"value\":\"200ukex\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\"},{\"key\":\"amount\",\"value\":\"200ukex\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/kira.gov.MsgRequestIdentityRecordsVerify\"},{\"key\":\"sender\",\"value\":\"kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"kira18q47tkrtrthf72xyscens2kv58ufpqtyxvxwuv\"},{\"key\":\"sender\",\"value\":\"kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\"},{\"key\":\"amount\",\"value\":\"200ukex\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "coin_received",
            "attributes": [
              {
                "key": "receiver",
                "value": "kira18q47tkrtrthf72xyscens2kv58ufpqtyxvxwuv"
              },
              {
                "key": "amount",
                "value": "200ukex"
              }
            ]
          },
          {
            "type": "coin_spent",
            "attributes": [
              {
                "key": "spender",
                "value": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4"
              },
              {
                "key": "amount",
                "value": "200ukex"
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/kira.gov.MsgRequestIdentityRecordsVerify"
              },
              {
                "key": "sender",
                "value": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4"
              }
            ]
          },
          {
            "type": "transfer",
            "attributes": [
              {
                "key": "recipient",
                "value": "kira18q47tkrtrthf72xyscens2kv58ufpqtyxvxwuv"
              },
              {
                "key": "sender",
                "value": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4"
              },
              {
                "key": "amount",
                "value": "200ukex"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "0",
    "gas_used": "0",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/kira.gov.MsgRequestIdentityRecordsVerify",
            "address": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
            "verifier": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
            "record_ids": [
              "19"
            ],
            "tip": {
              "denom": "ukex",
              "amount": "200"
            }
          }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
      },
      "auth_info": {
        "signer_infos": [
          {
            "public_key": {
              "@type": "/cosmos.crypto.secp256k1.PubKey",
              "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "88"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "ukex",
              "amount": "99"
            }
          ],
          "gas_limit": "200000",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
        "e5TXzbKnTeR9Hlf8980rEoJj7IMi0aXaADkO0nBnFbFXIF7DFyqVy5EZjyP9tVgd9wPdZws7QQLH760gQ+CU0w=="
      ]
    },
    "timestamp": "2023-06-08T12:53:42Z",
    "events": [
      {
        "type": "tx",
        "attributes": [
          {
            "key": "YWNjX3NlcQ==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NC84OA==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "ZTVUWHpiS25UZVI5SGxmODk4MHJFb0pqN0lNaTBhWGFBRGtPMG5CbkZiRlhJRjdERnlxVnk1RVpqeVA5dFZnZDl3UGRad3M3UVFMSDc2MGdRK0NVMHc9PQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "OTl1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "OTl1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "OTl1a2V4",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "ZmVl",
            "value": "OTl1a2V4",
            "index": true
          },
          {
            "key": "ZmVlX3BheWVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "L2tpcmEuZ292Lk1zZ1JlcXVlc3RJZGVudGl0eVJlY29yZHNWZXJpZnk=",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "a2lyYTE4cTQ3dGtydHJ0aGY3Mnh5c2NlbnMya3Y1OHVmcHF0eXh2eHd1dg==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "a2lyYTE4cTQ3dGtydHJ0aGY3Mnh5c2NlbnMya3Y1OHVmcHF0eXh2eHd1dg==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjAwdWtleA==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
            "index": true
          }
        ]
      }
    ]
  }
  ```
</details>

<details>
  <summary>Chech identity record verify requests</summary>

  Check all requests:
  ```
  sekaid q customgov all-identity-record-verify-requests -o json | jq
  ```

  ```json
  {
    "verify_records": [
      // . . .
      {
        "id": "4",
        "address": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
        "verifier": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
        "recordIds": [
          "19"
        ],
        "tip": {
          "denom": "ukex",
          "amount": "200"
        },
        "lastRecordEditDate": "2023-06-08T12:25:11.014802004Z"
      }
    ],
    "pagination": {
      "next_key": null,
      "total": "0"
    }
  }
  ```

  Check request by ID:
  ```
  sekaid query customgov identity-record-verify-request 4 -o json | jq
  ```

  ```json
  {
    "verify_record": {
      "id": "4",
      "address": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
      "verifier": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
      "recordIds": [
        "19"
      ],
      "tip": {
        "denom": "ukex",
        "amount": "200"
      },
      "lastRecordEditDate": "2023-06-08T12:25:11.014802004Z"
    }
  }
  ```

  Check request by requester:
  ```
  sekaid query customgov identity-record-verify-requests-by-requester kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --output=json | jq
  ```

  ```json
  {
    "verify_records": [
      {
        "id": "4",
        "address": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
        "verifier": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
        "recordIds": [
          "19"
        ],
        "tip": {
          "denom": "ukex",
          "amount": "200"
        },
        "lastRecordEditDate": "2023-06-08T12:25:11.014802004Z"
      }
    ],
    "pagination": {
      "next_key": null,
      "total": "0"
    }
  }
  ```

  Check request by approver:
  ```
  sekaid query customgov identity-record-verify-requests-by-approver kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt --output=json | jq
  ```

  ```json
  {
    "verify_records": [
      {
        "id": "4",
        "address": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
        "verifier": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
        "recordIds": [
          "19"
        ],
        "tip": {
          "denom": "ukex",
          "amount": "200"
        },
        "lastRecordEditDate": "2023-06-08T12:25:11.014802004Z"
      }
    ],
    "pagination": {
      "next_key": null,
      "total": "0"
    }
  }
  ```
</details>

[Return to top](#sekai)

##### 21.7.10 role

[Return to top](#sekai)

##### 21.7.11 set-execution-fee

Submit a transaction to set execution fee.

Usage:
```
sekaid tx customgov set-execution-fee [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ✅ yes |
| `--default_parameters uint`   | default parameters                                                                                                                                          | ❌ ?   |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--execution_fee uint`        | execution fee (default `10`)                                                                                                                                | ✅ yes |
| `--failure_fee uint`          | failure fee (default `1`)                                                                                                                                   | ✅ yes |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ✅ yes |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ❌ ?   |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for set-execution-fee                                                                                                                                  | ✅ yes |
| `--infos-file string`         | The infos file for identity request.                                                                                                                        | ✅ yes |
| `--infos-json string`         | The infos json for identity request.                                                                                                                        | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default `'home'` directory will be used                                                                       | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ✅ yes |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ✅ yes |
| `--record-ids string`         | Concatenated identity record ids array. e.g. `1,2`                                                                                                          | ✅ yes |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--timeout uint`              | timeout                                                                                                                                                     | ✅ yes |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `--tip string`                | The tip to be given to the verifier.                                                                                                                        | ✅ yes |
| `--transaction_type string`   | execution type                                                                                                                                              | ✅ yes |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ✅ yes |


| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx customgov set-execution-fee --help
Submit a transaction to set execution fee

Usage:
  sekaid tx customgov set-execution-fee [flags]

Flags:
  -a, --account-number uint       The account number of the signing account (offline mode only)
  -b, --broadcast-mode string     Transaction broadcasting mode (sync|async|block) (default "sync")
      --default_parameters uint   default parameters
      --dry-run                   ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --execution_fee uint        execution fee (default 10)
      --failure_fee uint          failure fee (default 1)
      --fee-account string        Fee account pays fees for the transaction instead of deducting from the signer
      --fees string               Fees to pay along with transaction; eg: 10uatom
      --from string               Name or address of private key with which to sign
      --gas string                gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float      adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string         Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only             Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                      help for set-execution-fee
      --keyring-backend string    Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string        The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                    Use a connected Ledger device
      --node string               <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string               Note to add a description to the transaction (previously --memo)
      --offline                   Offline mode (does not allow any online functionality
  -o, --output string             Output format (text|json) (default "json")
  -s, --sequence uint             The sequence number of the signing account (offline mode only)
      --sign-mode string          Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout uint              timeout
      --timeout-height uint       Set a block timeout height to prevent the tx from being committed past a certain height
      --transaction_type string   execution type
  -y, --yes                       Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
sekaid tx customgov set-execution-fee --execution_fee=100 --failure_fee=100 --transaction_type=create-role --timeout=10 --from=kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --chain-id=localnet-4 --keyring-backend=test --home=/root/.sekai --fees=100ukex --yes --broadcast-mode=async --output=json | jq
{
  "height": "0",
  "txhash": "DDF0EA387745E2733B959E27C9B86BE85CC20AB305D18434674AE51ED5792A7E",
  "codespace": "",
  "code": 0,
  "data": "",
  "raw_log": "",
  "logs": [],
  "info": "",
  "gas_wanted": "0",
  "gas_used": "0",
  "tx": null,
  "timestamp": "",
  "events": []
}
```

<details>
  <summary>Check tx execution</summary>

```
sekaid q tx DDF0EA387745E2733B959E27C9B86BE85CC20AB305D18434674AE51ED5792A7E -o json | jq
```

```json
{
  "height": "144169",
  "txhash": "DDF0EA387745E2733B959E27C9B86BE85CC20AB305D18434674AE51ED5792A7E",
  "codespace": "",
  "code": 0,
  "data": "0A1E0A1C2F6B6972612E676F762E4D7367536574457865637574696F6E466565",
  "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/kira.gov.MsgSetExecutionFee\"}]},{\"type\":\"set_execution_fee\",\"attributes\":[{\"key\":\"proposer\",\"value\":\"kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4\"},{\"key\":\"transaction_type\",\"value\":\"create-role\"},{\"key\":\"execution_fee\",\"value\":\"100\"},{\"key\":\"failure_fee\",\"value\":\"100\"},{\"key\":\"time_out\",\"value\":\"100\"},{\"key\":\"default_parameters\",\"value\":\"0\"}]}]}]",
  "logs": [
    {
      "msg_index": 0,
      "log": "",
      "events": [
        {
          "type": "message",
          "attributes": [
            {
              "key": "action",
              "value": "/kira.gov.MsgSetExecutionFee"
            }
          ]
        },
        {
          "type": "set_execution_fee",
          "attributes": [
            {
              "key": "proposer",
              "value": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4"
            },
            {
              "key": "transaction_type",
              "value": "create-role"
            },
            {
              "key": "execution_fee",
              "value": "100"
            },
            {
              "key": "failure_fee",
              "value": "100"
            },
            {
              "key": "time_out",
              "value": "100"
            },
            {
              "key": "default_parameters",
              "value": "0"
            }
          ]
        }
      ]
    }
  ],
  "info": "",
  "gas_wanted": "0",
  "gas_used": "0",
  "tx": {
    "@type": "/cosmos.tx.v1beta1.Tx",
    "body": {
      "messages": [
        {
          "@type": "/kira.gov.MsgSetExecutionFee",
          "transaction_type": "create-role",
          "execution_fee": "100",
          "failure_fee": "100",
          "timeout": "10",
          "default_parameters": "0",
          "proposer": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4"
        }
      ],
      "memo": "",
      "timeout_height": "0",
      "extension_options": [],
      "non_critical_extension_options": []
    },
    "auth_info": {
      "signer_infos": [
        {
          "public_key": {
            "@type": "/cosmos.crypto.secp256k1.PubKey",
            "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
          },
          "mode_info": {
            "single": {
              "mode": "SIGN_MODE_DIRECT"
            }
          },
          "sequence": "121"
        }
      ],
      "fee": {
        "amount": [
          {
            "denom": "ukex",
            "amount": "100"
          }
        ],
        "gas_limit": "200000",
        "payer": "",
        "granter": ""
      }
    },
    "signatures": [
      "hSfBo7I+9seNSzf0PEDUHLOj+W1Ug0uHxrokuYYXmH5TTqrWYD13Cr//U6eS9qQhVuGjPTg4WlId63D78C0yHQ=="
    ]
  },
  "timestamp": "2023-06-12T17:28:31Z",
  "events": [
    {
      "type": "tx",
      "attributes": [
        {
          "key": "YWNjX3NlcQ==",
          "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NC8xMjE=",
          "index": true
        }
      ]
    },
    {
      "type": "tx",
      "attributes": [
        {
          "key": "c2lnbmF0dXJl",
          "value": "aFNmQm83SSs5c2VOU3pmMFBFRFVITE9qK1cxVWcwdUh4cm9rdVlZWG1INVRUcXJXWUQxM0NyLy9VNmVTOXFRaFZ1R2pQVGc0V2xJZDYzRDc4QzB5SFE9PQ==",
          "index": true
        }
      ]
    },
    {
      "type": "coin_spent",
      "attributes": [
        {
          "key": "c3BlbmRlcg==",
          "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
          "index": true
        },
        {
          "key": "YW1vdW50",
          "value": "MTAwdWtleA==",
          "index": true
        }
      ]
    },
    {
      "type": "coin_received",
      "attributes": [
        {
          "key": "cmVjZWl2ZXI=",
          "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
          "index": true
        },
        {
          "key": "YW1vdW50",
          "value": "MTAwdWtleA==",
          "index": true
        }
      ]
    },
    {
      "type": "transfer",
      "attributes": [
        {
          "key": "cmVjaXBpZW50",
          "value": "a2lyYTE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHFrZncycw==",
          "index": true
        },
        {
          "key": "c2VuZGVy",
          "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
          "index": true
        },
        {
          "key": "YW1vdW50",
          "value": "MTAwdWtleA==",
          "index": true
        }
      ]
    },
    {
      "type": "message",
      "attributes": [
        {
          "key": "c2VuZGVy",
          "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
          "index": true
        }
      ]
    },
    {
      "type": "tx",
      "attributes": [
        {
          "key": "ZmVl",
          "value": "MTAwdWtleA==",
          "index": true
        },
        {
          "key": "ZmVlX3BheWVy",
          "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
          "index": true
        }
      ]
    },
    {
      "type": "message",
      "attributes": [
        {
          "key": "YWN0aW9u",
          "value": "L2tpcmEuZ292Lk1zZ1NldEV4ZWN1dGlvbkZlZQ==",
          "index": true
        }
      ]
    },
    {
      "type": "set_execution_fee",
      "attributes": [
        {
          "key": "cHJvcG9zZXI=",
          "value": "a2lyYTF2bXdkZ3c0MjZhajlmeDMzZnF1c210ZzZyNjV5eXVjbXg2cmR0NA==",
          "index": true
        },
        {
          "key": "dHJhbnNhY3Rpb25fdHlwZQ==",
          "value": "Y3JlYXRlLXJvbGU=",
          "index": true
        },
        {
          "key": "ZXhlY3V0aW9uX2ZlZQ==",
          "value": "MTAw",
          "index": true
        },
        {
          "key": "ZmFpbHVyZV9mZWU=",
          "value": "MTAw",
          "index": true
        },
        {
          "key": "dGltZV9vdXQ=",
          "value": "MTAw",
          "index": true
        },
        {
          "key": "ZGVmYXVsdF9wYXJhbWV0ZXJz",
          "value": "MA==",
          "index": true
        }
      ]
    }
  ]
}
```
</details>

<details>
  <summary>Check changes</summary>

```
sekaid q customgov execution-fee create-role -o json | jq
{
  "fee": {
    "transaction_type": "create-role",
    "execution_fee": "100",
    "failure_fee": "100",
    "timeout": "10",
    "default_parameters": "0"
  }
}
```
</details>

[Return to top](#sekai)

##### 21.7.12 set-network-properties

[Return to top](#sekai)

#### 21.8 customslashing

[Return to top](#sekai)

#### 21.9 customstaking

[Return to top](#sekai)

#### 21.10 decode

[Return to top](#sekai)

#### 21.11 distributor

[Return to top](#sekai)

#### 21.12 encode

[Return to top](#sekai)

#### 21.13 layer2

[Return to top](#sekai)

#### 21.14 multisign

Sign transactions created with the --generate-only flag that require multisig signatures.

Usage:
```
sekaid tx multisign [file] [name] [[signature]...] [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `--amino`                     | Generate Amino-encoded JSON suitable for submitting to the txs REST endpoint                                                                                | ❌ ?   |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ❌ ?   |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ❌ ?   |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ✅ yes |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for multisign                                                                                                                                          | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default `'home'` directory will be used                                                                       | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ❌ ?   |
| `--multisig string`           | Address or key name of the multisig account on behalf of which the transaction shall be signed                                                              | ❌ ?   |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ❌ ?   |
| `--output-document string`    | The document is written to the given file instead of STDOUT                                                                                                 | ❌ ?   |
| `--overwrite`                 | Overwrite existing signatures with a new one. If disabled, new signature will be appended                                                                   | ❌ ?   |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--signature-only`            | Print only the generated signature, then exit                                                                                                               | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ❌ ?   |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | Directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?   |

```
# sekaid tx multisign --help
Sign transactions created with the --generate-only flag that require multisig signatures.

Read one or more signatures from one or more [signature] file, generate a multisig signature compliant to the
multisig key [name], and attach the key name to the transaction read from [file].

Example:
$ sekaid tx multisign transaction.json k1k2k3 k1sig.json k2sig.json k3sig.json

If --signature-only flag is on, output a JSON representation
of only the generated signature.

If the --offline flag is on, the client will not reach out to an external node.
Account number or sequence number lookups are not performed so you must
set these parameters manually.

The current multisig implementation defaults to amino-json sign mode.
The SIGN_MODE_DIRECT sign mode is not supported.'

Usage:
  sekaid tx multisign [file] [name] [[signature]...] [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
      --amino                    Generate Amino-encoded JSON suitable for submitting to the txs REST endpoint
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for multisign
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
      --output-document string   The document is written to the given file instead of STDOUT
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --signature-only           Print only the generated signature, then exit
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

#TODO haven't tested

🟨  
🟨  
🟨  

[Return to top](#sekai)

#### 21.15 multistaking

[Return to top](#sekai)

#### 21.16 recovery

[Return to top](#sekai)

#### 21.17 sign

Sign a transaction created with the '--generate-only' flag.


| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `--amino`                     | Generate Amino encoded JSON suitable for submiting to the txs REST endpoint                                                                                 | ❌ ?   |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ❌ ?   |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ❌ ?   |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ✅ yes |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ✅ yes |
| `-h, --help`                  | help for sign                                                                                                                                               | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ✅ yes |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default `'home'` directory will be used                                                                       | ✅ yes |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ❌ ?   |
| `--multisig string`           | Address or key name of the multisig account on behalf of which the transaction shall be signed                                                              | ❌ ?   |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ✅ yes |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ❌ ?   |
| `--output-document string`    | The document will be written to the given file instead of STDOUT                                                                                            | ❌ ?   |
| `--overwrite`                 | Overwrite existing signatures with a new one. If disabled, new signature will be appended                                                                   | ❌ ?   |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--signature-only`            | Print only the signatures                                                                                                                                   | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ❌ ?   |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | Directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx sign tx --help
Sign a transaction created with the --generate-only flag.
It will read a transaction from [file], sign it, and print its JSON encoding.

If the --signature-only flag is set, it will output the signature parts only.

The --offline flag makes sure that the client will not reach out to full node.
As a result, the account and sequence number queries will not be performed and
it is required to set such parameters manually. Note, invalid values will cause
the transaction to fail.

The --multisig=<multisig_key> flag generates a signature on behalf of a multisig account
key. It implies --signature-only. Full multisig signed transactions may eventually
be generated via the 'multisign' command.

Usage:
  sekaid tx sign [file] [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
      --amino                    Generate Amino encoded JSON suitable for submiting to the txs REST endpoint
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for sign
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --multisig string          Address or key name of the multisig account on behalf of which the transaction shall be signed
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
      --output-document string   The document will be written to the given file instead of STDOUT
      --overwrite                Overwrite existing signatures with a new one. If disabled, new signature will be appended
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --signature-only           Print only the signatures
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

<details>
  <summary>Pregenerate tx</summary>

  ```
  /# sekaid tx bank send kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt 100samolean --home=/root/.sekai --chain-id=localnet-4 --keyring-backend=test --fees=99ukex --generate-only --output=json | jq > tx.json
  ```

  tx.json:
  ```json
  {
    "body": {
      "messages": [
        {
          "@type": "/cosmos.bank.v1beta1.MsgSend",
          "from_address": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
          "to_address": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
          "amount": [
            {
              "denom": "samolean",
              "amount": "100"
            }
          ]
        }
      ],
      "memo": "",
      "timeout_height": "0",
      "extension_options": [],
      "non_critical_extension_options": []
    },
    "auth_info": {
      "signer_infos": [],
      "fee": {
        "amount": [
          {
            "denom": "ukex",
            "amount": "99"
          }
        ],
        "gas_limit": "200000",
        "payer": "",
        "granter": ""
      }
    },
    "signatures": []
  }
  ```

  __Important to use the `--generate-only` flag__
</details>

```
/# sekaid tx sign tx.json --from=kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --home=/root/.sekai --chain-id=localnet-4 --keyring-backend=test --output=json | jq
{
  "body": {
    "messages": [
      {
        "@type": "/cosmos.bank.v1beta1.MsgSend",
        "from_address": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
        "to_address": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
        "amount": [
          {
            "denom": "samolean",
            "amount": "100"
          }
        ]
      }
    ],
    "memo": "",
    "timeout_height": "0",
    "extension_options": [],
    "non_critical_extension_options": []
  },
  "auth_info": {
    "signer_infos": [
      {
        "public_key": {
          "@type": "/cosmos.crypto.secp256k1.PubKey",
          "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
        },
        "mode_info": {
          "single": {
            "mode": "SIGN_MODE_DIRECT"
          }
        },
        "sequence": "84"
      }
    ],
    "fee": {
      "amount": [
        {
          "denom": "ukex",
          "amount": "99"
        }
      ],
      "gas_limit": "200000",
      "payer": "",
      "granter": ""
    }
  },
  "signatures": [
    "OukW9MMdqzxI2LUymHKu9DXz1A6B6+RxEZBSn55vw5QSnC7DHi+ryIm2tDcwMPRm84t2uDh55DGyniNyyHoSrg=="
  ]
}
```

[Return to top](#sekai)

#### 21.18 sign-batch

Sign batch files of transactions generated with `--generate-only`.

Usage:
```
sekaid tx sign-batch [file] [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ❌ ?   |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ❌ ?   |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ❌ ?   |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ❌ ?   |
| `-h, --help`                  | help for sign-batch                                                                                                                                         | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ❌ ?   |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default `'home'` directory will be used                                                                       | ❌ ?   |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ❌ ?   |
| `--multisig string`           | Address or key name of the multisig account on behalf of which the transaction shall be signed                                                              | ❌ ?   |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ❌ ?   |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ❌ ?   |
| `--output-document string`    | The document will be written to the given file instead of STDOUT                                                                                            | ❌ ?   |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--signature-only`            | Print only the generated signature, then exit (default `true`)                                                                                              | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ❌ ?   |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | Directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx sign-batch --help
Sign batch files of transactions generated with --generate-only.
The command processes list of transactions from file (one StdTx each line), generate
signed transactions or signatures and print their JSON encoding, delimited by '\n'.
As the signatures are generated, the command updates the account sequence number accordingly.

If the --signature-only flag is set, it will output the signature parts only.

The --offline flag makes sure that the client will not reach out to full node.
As a result, the account and the sequence number queries will not be performed and
it is required to set such parameters manually. Note, invalid values will cause
the transaction to fail. The sequence will be incremented automatically for each
transaction that is signed.

The --multisig=<multisig_key> flag generates a signature on behalf of a multisig
account key. It implies --signature-only.

Usage:
  sekaid tx sign-batch [file] [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for sign-batch
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --multisig string          Address or key name of the multisig account on behalf of which the transaction shall be signed
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
      --output-document string   The document will be written to the given file instead of STDOUT
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --signature-only           Print only the generated signature, then exit (default true)
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

#TODO haven't tested

🟨  
🟨  
🟨  

[Return to top](#sekai)

#### 21.19 spending

[Return to top](#sekai)

#### 21.20 tokens

[Return to top](#sekai)

#### 21.21 ubi

[Return to top](#sekai)

#### 21.22 upgrade

[Return to top](#sekai)

#### 21.23 validate-signatures

Print the addresses that must sign the transaction, those who have already signed it, and make sure that signatures are in the correct order.

Usage:
```
sekaid tx validate-signatures [file] [flags]
```

| Flags                         | Description                                                                                                                                                 | Work  |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| `-a, --account-number uint`   | The account number of the signing account (offline mode only)                                                                                               | ❌ ?   |
| `-b, --broadcast-mode string` | Transaction broadcasting mode (`sync\|async\|block`) (default `"sync"`)                                                                                     | ❌ ?   |
| `--dry-run`                   | ignore the `--gas` flag and perform a simulation of a transaction, but don't broadcast it                                                                   | ❌ ?   |
| `--fee-account string`        | Fee account pays fees for the transaction instead of deducting from the signer                                                                              | ❌ ?   |
| `--fees string`               | Fees to pay along with transaction; eg: `10uatom`                                                                                                           | ❌ ?   |
| `--from string`               | Name or address of private key with which to sign                                                                                                           | ❌ ?   |
| `--gas string`                | gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default `200000`)                                                | ❌ ?   |
| `--gas-adjustment float`      | adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default `1`) | ❌ ?   |
| `--gas-prices string`         | Gas prices in decimal format to determine the transaction fee (e.g. `0.1uatom`)                                                                             | ❌ ?   |
| `--generate-only`             | Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)                                                    | ❌ ?   |
| `-h, --help`                  | help for validate-signatures                                                                                                                                | ✅ yes |
| `--keyring-backend string`    | Select keyring's backend (`os\|file\|kwallet\|pass\|test\|memory`) (default `"os"`)                                                                         | ❌ ?   |
| `--keyring-dir string`        | The client Keyring directory; if omitted, the default `'home'` directory will be used                                                                       | ❌ ?   |
| `--ledger`                    | Use a connected Ledger device                                                                                                                               | ❌ ?   |
| `--multisig string`           | Address or key name of the multisig account on behalf of which the transaction shall be signed                                                              | ❌ ?   |
| `--node string`               | \<host\>:\<port\> to tendermint rpc interface for this chain (default `"tcp://localhost:26657"`)                                                            | ❌ ?   |
| `--note string`               | Note to add a description to the transaction (previously `--memo`)                                                                                          | ❌ ?   |
| `--offline`                   | Offline mode (does not allow any online functionality                                                                                                       | ❌ ?   |
| `-o, --output string`         | Output format (`text\|json`) (default `"json"`)                                                                                                             | ❌ ?   |
| `--output-document string`    | The document will be written to the given file instead of STDOUT                                                                                            | ❌ ?   |
| `-s, --sequence uint`         | The sequence number of the signing account (offline mode only)                                                                                              | ❌ ?   |
| `--sign-mode string`          | Choose sign mode (`direct\|amino-json`), this is an advanced feature                                                                                        | ❌ ?   |
| `--signature-only`            | Print only the generated signature, then exit (default `true`)                                                                                              | ❌ ?   |
| `--timeout-height uint`       | Set a block timeout height to prevent the tx from being committed past a certain height                                                                     | ❌ ?   |
| `-y, --yes`                   | Skip tx broadcasting prompt confirmation                                                                                                                    | ❌ ?   |



| Global Flags          | Description                                                                            | Work  |
| --------------------- | -------------------------------------------------------------------------------------- | ----- |
| `--chain-id string`   | The network chain ID                                                                   | ✅ yes |
| `--home string`       | Directory for config and data (default `"/root/.sekaid"`)                              | ✅ yes |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ❌ ?   |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ❌ ?   |
| `--trace`             | Print out full stack trace on errors                                                   | ❌ ?   |

```
/# sekaid tx validate-signatures --help
Print the addresses that must sign the transaction, those who have already
signed it, and make sure that signatures are in the correct order.

The command would check whether all required signers have signed the transactions, whether
the signatures were collected in the right order, and if the signature is valid over the
given transaction. If the --offline flag is also set, signature validation over the
transaction will be not be performed as that will require RPC communication with a full node.

Usage:
  sekaid tx validate-signatures [file] [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for validate-signatures
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

<details>
  <summary>Pregenerate signed tx</summary>

  ```
  sekaid tx sign tx1.json --from=kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 --home=/root/.sekai --chain-id=localnet-4 --keyring-backend=test --node=tcp://localhost:26657 --output=json | jq > tx1_signed.json
  ```

  tx1_signed.json:
  ```json
  {
  "body": {
    "messages": [
      {
        "@type": "/cosmos.bank.v1beta1.MsgSend",
        "from_address": "kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4",
        "to_address": "kira17aeqxvkl3g37pfcgwkqxvkrn63jfljfvjravnt",
        "amount": [
          {
            "denom": "samolean",
            "amount": "100"
          }
        ]
      }
    ],
    "memo": "",
    "timeout_height": "0",
    "extension_options": [],
    "non_critical_extension_options": []
  },
  "auth_info": {
    "signer_infos": [
      {
        "public_key": {
          "@type": "/cosmos.crypto.secp256k1.PubKey",
          "key": "AjjA26m3ab7z6Ddrqeons69CREF8q/d815X180ucZLmo"
        },
        "mode_info": {
          "single": {
            "mode": "SIGN_MODE_DIRECT"
          }
        },
        "sequence": "84"
      }
    ],
    "fee": {
      "amount": [],
      "gas_limit": "200000",
      "payer": "",
      "granter": ""
    }
  },
  "signatures": [
    "yqZAxLBrN+vsF3oWozOCItihdziTv1iZkt5aOhS2krZpE2Wk8SaXRMpPaoeBRDiSYd7Nd26u0Zc9uQMYvk1e3g=="
  ]
}
  ```
</details>

```
/# sekaid tx validate-signatures --home=/root/.sekai --chain-id=localnet-4 --keyring-backend=test tx1_signed.json
Signers:
  0: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4

Signatures:
  0: kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4                        [OK]
```

[Return to top](#sekai)

### 22. val-address

Get validator address from account address

Usage: 

```
sekaid val-address [address] [flags]
```

| Flags               | Description                                                                           | Work  |
| ------------------- | ------------------------------------------------------------------------------------- | ----- |
| --height int        | Use a specific height to query state at (this can error if the node is pruning state) | ❌ no  |
| -h, --help          | help for val-address                                                                  | ✅ yes |
| --node string       | : to Tendermint RPC interface for this chain (default "tcp://localhost:26657")        | ❌ no  |
| -o, --output string | Output format (text\|json) (default "text")                                           | ❌ no  |


| Global flags       | Description                                                                        | Work |
| ------------------ | ---------------------------------------------------------------------------------- | ---- |
| --home string      | directory for config and data (default "/root/.sekaid")                            | ❌ no |
| --log_format       | The logging format (json\|plain) (default "plain")                                 | ❌ no |
| --log_level string | The logging level (trace\|debug\|info\|warn\|error\|fatal\|panic) (default "info") | ❌ ?  |
| --trace            | print out full stack trace on errors                                               | ❌ ?  |

```
# sekaid val-address --help
Get validator address from account address

Usage:
  sekaid val-address [address] [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for val-address
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
# sekaid val-address kira1tcyq0y66mmdpydepexc8lwrmemp9wzgmg2rqhm && echo
kiravaloper1tcyq0y66mmdpydepexc8lwrmemp9wzgmmvlr0h
```

[Return to top](#sekai)

### 23. valcons-address

Get validator consensus address from account address

Usage:

```
  sekaid valcons-address [address] [flags]
```

| Flags               | Description                                                                                | Work  |
| ------------------- | ------------------------------------------------------------------------------------------ | ----- |
| --height int        | Use a specific height to query state at (this can error if the node is pruning state)      | ❌ no  |
| -h, --help          | help for val-address                                                                       | ✅ yes |
| --node string       | <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657") | ❌ no  |
| -o, --output string | Output format (text\|json) (default "text")                                                | ❌ no  |



| Global flags       | Description                                                                        | Work |
| ------------------ | ---------------------------------------------------------------------------------- | ---- |
| --home string      | directory for config and data (default "/root/.sekaid")                            | ❌ no |
| --log_format       | The logging format (json\|plain) (default "plain")                                 | ❌ no |
| --log_level string | The logging level (trace\|debug\|info\|warn\|error\|fatal\|panic) (default "info") | ❌ ?  |
| --trace            | print out full stack trace on errors                                               | ❌ ?  |

```
# sekaid valcons-address  --help
Get validator consensus address from account address

Usage:
  sekaid valcons-address [address] [flags]

Flags:
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for valcons-address
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors

```

```
# sekaid valcons-address  kira1vmwdgw426aj9fx33fqusmtg6r65yyucmx6rdt4 && echo
kiravalcons1vmwdgw426aj9fx33fqusmtg6r65yyucmp0vjlc
```


[Return to top](#sekai)

### 24. validate-genesis

Validates the genesis file at the default location or at the location passed as an arg

Usage:

```
sekaid validate-genesis [file] [flags]
```
| Flags  | Description      | Work  |
| ------ | ---------------- | ----- |
| --help | help for version | ✅ yes |



| Global flags       | Description                                                                        | Work  |
| ------------------ | ---------------------------------------------------------------------------------- | ----- |
| --home string      | directory for config and data (default "/root/.sekaid")                            | ✅ yes |
| --log_format       | The logging format (json\|plain) (default "plain")                                 | ❌ no  |
| --log_level string | The logging level (trace\|debug\|info\|warn\|error\|fatal\|panic) (default "info") | ❌ ?   |
| --trace            | print out full stack trace on errors                                               | ❌ ?   |

```
# sekaid validate-genesis --help
validates the genesis file at the default location or at the location passed as an arg

Usage:
  sekaid validate-genesis [file] [flags]

Flags:
  -h, --help   help for validate-genesis

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors

```

```
# sekaid validate-genesis --home $SEKAID_HOME
File at /root/.sekai/config/genesis.json is a valid genesis file
```

[Return to top](#sekai)

### 25. version
Print the application binary version information
Usage:
```
sekaid version [flags]
```
| Flags           | Description                                 | Work  |
| --------------- | ------------------------------------------- | ----- |
| --help          | help for version                            | ✅ yes |
| --long          | Print long version information              | ✅ yes |
| --output string | Output format (text\|json) (default "text") | ❌ ?no |



| Global flags       | Description                                                                        | Work  |
| ------------------ | ---------------------------------------------------------------------------------- | ----- |
| --home string      | directory for config and data (default "/root/.sekaid")                            | ✅ yes |
| --log_format       | The logging format (json\|plain) (default "plain")                                 | ❌ no  |
| --log_level string | The logging level (trace\|debug\|info\|warn\|error\|fatal\|panic) (default "info") | ❌ ?   |
| --trace            | print out full stack trace on errors                                               | ❌ ?   |

```
#sekaid version  --help
Print the application binary version information

Usage:
  sekaid version [flags]

Flags:
  -h, --help            help for version
      --long            Print long version information
  -o, --output string   Output format (text|json) (default "text")

Global Flags:
      --home string         directory for config and data (default "/root/.sekaid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

```
#sekaid version
v0.3.15.1
```

[Return to top](#sekai)
