# OS state after init:

## Config file
config file has to be located inside `~/.config/km2/` folder with `config.toml` file name.

## Firewald 

firewalld must be configured so that it has active `validator` zone with next parameters:
- Opened system ports `22/tcp`, `53/udp`. 
- Opened `RpcPort`, `GrpcPort`, `P2PPort`, `PrometheusPort`, `InterxPort` with values from `config.toml` file. 


## Docker
 Docker must have 2 existing containers `sekaid`, `interx` and volume with `kira_volume`. 
 
- sekaid container has to have forwarded ports: `RpcPort`, `P2PPort`, `PrometheusPort` with values from `config.toml` file. 
- interx container has to have forwarded  `InterxPort` port with value from `config.toml` file. 
- both containers must have mounted volume with `VolumeName` value from `config.toml` file.