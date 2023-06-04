# Sekai
`ver.: v0.3.16`
               
* [Sekai](#sekai)
  * [Context](#context)
    * [1. add-genesis-account](#1-add-genesis-account)
    * [2. collect-gentxs](#2-collect-gentxs)
    * [3. config](#3-config)
    * [4. debug](#4-debug)
      * [4.1 addr](#41-addr)
      * [4.2 pubkey](#42-pubkey)
      * [4.3 raw-bytes](#43-raw-bytes)
    * [5. export](#5-export)
    * [6. export-metadata](#6-export-metadata)
    * [7. export-minimized-genesis](#7-export-minimized-genesis)
    * [8. gentx](#8-gentx)
    * [9. gentx-claim](#9-gentx-claim)
    * [10. help](#10-help)
    * [11. init](#11-init)
    * [12. keys](#12-keys)
      * [12.1 add](#121-add)
      * [12.2 delete](#122-delete)
      * [12.3 export](#123-export)
      * [12.4 import](#124-import)
      * [12.5 list](#125-list)
      * [12.6 migrate](#126-migrate)
      * [12.7 mnemonic](#127-mnemonic)
      * [12.8 parse](#128-parse)
      * [12.9 show](#129-show)
    * [13. new-genesis-from-exported](#13-new-genesis-from-exported)
    * [14. query](#14-query)
    * [15. rollback](#15-rollback)
    * [16. rosetta](#16-rosetta)
    * [17. start](#17-start)
    * [18. status](#18-status)
    * [19. tendermint](#19-tendermint)
    * [20. testnet](#20-testnet)
    * [21. tx](#21-tx)
    * [22. val-address](#22-val-address)
    * [23. valcons-address](#23-valcons-address)
    * [24. validate-genesis](#24-validate-genesis)
    * [25. version](#25-version)

## Context

### 1. add-genesis-account

### 2. collect-gentxs

### 3. config

### 4. debug

#### 4.1 addr

#### 4.2 pubkey

#### 4.3 raw-bytes

### 5. export

### 6. export-metadata

### 7. export-minimized-genesis

### 8. gentx

### 9. gentx-claim

### 10. help

### 11. init

### 12. keys

Keyring management commands

Usage:
```
sekaid keys [command]
```

Available Commands:

| Subcommand | Description                                                                                               |
| ---------- | --------------------------------------------------------------------------------------------------------- |
| `add`      | Add an encrypted private key (either newly generated or recovered), encrypt it, and save to \<name\> file |
| `delete`   | Delete the given keys                                                                                     |
| `export`   | Export private keys                                                                                       |
| `import`   | Import private keys into the local keybase                                                                |
| `list`     | List all keys                                                                                             |
| `migrate`  | Migrate keys from the legacy (db-based) Keybase                                                           |
| `mnemonic` | Compute the bip39 mnemonic for some input entropy                                                         |
| `parse`    | Parse address from hex to bech32 and vice versa                                                           |
| `show`     | Retrieve key information by name or address                                                               |

| Flags                      | Description                                                                           | Work |
| -------------------------- | ------------------------------------------------------------------------------------- | ---- |
| `-h, --help`               | Help for keys                                                                         | yes  |
| `--keyring-backend string` | Select keyring's backend (`os\|file\|test`) (default `"os"`)                          | yes  |
| `--keyring-dir string`     | The client Keyring directory; if omitted, the default `'home'` directory will be used | yes  |
| `--output string`          | Output format (`text\|json`) (default `"text"`)                                       | yes  |

| Global Flags          | Description                                                                            | Work |
| --------------------- | -------------------------------------------------------------------------------------- | ---- |
| `--home string`       | Directory for config and data (default `"/root/.sekaid"`)                              | yes  |
| `--log_format string` | The logging format (`json\|plain`) (default `"plain"`)                                 | ?no  |
| `--log_level string`  | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ?no  |
| `--trace`             | Print out full stack trace on errors                                                   | ?no  |

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

#### 12.1 add

#### 12.2 delete

#### 12.3 export

#### 12.4 import

#### 12.5 list

#### 12.6 migrate

#### 12.7 mnemonic

#### 12.8 parse

#### 12.9 show

### 13. new-genesis-from-exported

### 14. query

### 15. rollback

### 16. rosetta

### 17. start

### 18. status
Query remote node for status

#TODO: the descritpion is not accurate. `sekaid status` queries the node where it was launched by default.

Usage:
```
sekaid status [flags]
```

| Flags    | Description                                            | Work |
| -------- | ------------------------------------------------------ | ---- |
| `--help` | Help for status                                        | yes  |
| `--node` | Node to connect to (default `"tcp://localhost:26657"`) | yes  |

| Global flags         | Description                                                                            | Work |
| -------------------- | -------------------------------------------------------------------------------------- | ---- |
| `--home`             | directory for config and data (default `"/root/.sekaid"`)                              | yes  |
| `--log_format`       | The logging format (`json\|plain`) (default `"plain"`)                                 | ?no  |
| `--log_level string` | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ?no  |
| `--trace`            | Print out full stack trace on errors                                                   | ?no  |

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

### 19. tendermint

### 20. testnet

### 21. tx

### 22. val-address

### 23. valcons-address

### 24. validate-genesis

### 25. version
