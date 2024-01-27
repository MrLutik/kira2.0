#  Operating System Configuration After Node Initialization:

## Configuration File

The configuration file should be placed within the `~/.config/km2/` directory and named `config.toml`.



## Firewalld Configuration

Ensure that the firewalld is set up with an active `validator` zone, featuring the following parameters:

Open system ports `22/tcp` and `53/udp`.
Open the ports `RpcPort`, `GrpcPort`, `P2PPort`, `PrometheusPort`, and `InterxPort` with the corresponding values specified in the `config.toml` file.



## Docker Setup

For Docker, make sure there are two existing containers, namely sekaid and interx, along with a volume named kira_volume.

The `sekaid` container should have the following ports forwarded: `RpcPort`, `P2PPort`, and `PrometheusPort`, with values obtained from the `config.toml` file.

The `interx` container should have the `InterxPort` forwarded with the value specified in the config.toml file.

Both containers must have a mounted volume using the `VolumeName` value from the `config.toml` file.
	
	
	
	
	