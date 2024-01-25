# KIRA2

- [KIRA2](#kira2)
	- [Context](#context)
		- [1. Completion](#1-completion)
		- [2. Firewall](#2-firewall)
		- [3. Help](#3-help)
		- [4. Init](#4-init)
		- [5. Maintenance](#5-maintenance)
		- [6. Monitoring](#6-monitoring)
		- [7. Start](#7-start)
		- [8. Stop](#8-stop)


## Context

### 1. Completion
??


### 2. Firewall
Command configure firewall

Usage
```
Usage:
  kira2 firewall [flags]
  kira2 firewall [command]
```


Available Commands:

| Subcommand              | Description                     |
| ----------------------- | ------------------------------- |
| [`blacklist`](#121-add) | subcommand for blacklisting ips |
| [`delete`](#122-delete) | subcommand for port closing     |
| [`export`](#123-export) | subcommand for port opening     |
| [`import`](#124-import) | subcommand for whitelisting ips |


| Flags          | Description                                                                                     | Work  |
| -------------- | ----------------------------------------------------------------------------------------------- | ----- |
| `--blacklist`  | Set this flag to block all ports except ssh                                                     | ✅ yes |
| `--close-port` | Set this flag to restore default setting for firewall (what km2 is set after node installation) | ✅ yes |
| `-h, --help`   | help for firewall                                                                               | ✅ yes |
| `--open-ports` | Set this flag to open all km2 default ports                                                     | ✅ yes |



### 3. Help
### 4. Init
### 5. Maintenance
### 6. Monitoring
### 7. Start
### 8. Stop