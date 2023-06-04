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

[Return to top](#sekai)

### 2. collect-gentxs

[Return to top](#sekai)

### 3. config

[Return to top](#sekai)

### 4. debug

[Return to top](#sekai)

#### 4.1 addr

[Return to top](#sekai)

#### 4.2 pubkey

[Return to top](#sekai)

#### 4.3 raw-bytes

[Return to top](#sekai)

### 5. export

[Return to top](#sekai)

### 6. export-metadata

[Return to top](#sekai)

### 7. export-minimized-genesis

[Return to top](#sekai)

### 8. gentx

[Return to top](#sekai)

### 9. gentx-claim

[Return to top](#sekai)

### 10. help

[Return to top](#sekai)

### 11. init

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

| Flags        | Description       | Work  |
| ------------ | ----------------- | ----- |
| `-h, --help` | Help for mnemonic | ✅ yes |



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

### 14. query

### 15. rollback

### 16. rosetta

### 17. start

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

| Flags    | Description                                            | Work |
| -------- | ------------------------------------------------------ | ---- |
| `--help` | Help for status                                        | yes  |
| `--node` | Node to connect to (default `"tcp://localhost:26657"`) | yes  |



| Global flags         | Description                                                                            | Work |
| -------------------- | -------------------------------------------------------------------------------------- | ---- |
| `--home`             | Directory for config and data (default `"/root/.sekaid"`)                              | yes  |
| `--log_format`       | The logging format (`json\|plain`) (default `"plain"`)                                 | ?    |
| `--log_level string` | The logging level (`trace\|debug\|info\|warn\|error\|fatal\|panic`) (default `"info"`) | ?    |
| `--trace`            | Print out full stack trace on errors                                                   | ?    |

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

[Return to top](#sekai)

### 20. testnet

[Return to top](#sekai)

### 21. tx

[Return to top](#sekai)

### 22. val-address

[Return to top](#sekai)

### 23. valcons-address

[Return to top](#sekai)

### 24. validate-genesis

[Return to top](#sekai)

### 25. version

[Return to top](#sekai)
