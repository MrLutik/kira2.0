#  Operating System Configuration After Node Initialization:

## Configuration File

The configuration file should be placed within the `~/.config/km2/` directory and named `config.toml`.



## Firewalld Configuration

Delete `docker` zone.

Ensure that the firewalld is set up with an active `validator` zone, featuring the following parameters:

Open system ports `22/tcp` and `53/udp`.
Open the ports `RpcPort`, `GrpcPort`, `P2PPort`, `PrometheusPort`, and `InterxPort` with the corresponding values specified in the `config.toml` file.

Then, the `docker0` network interface and the `kira_network` docker's network (as network interface) must be added to the validator zone.
 
<details>
<summary>How to get name of the interface from the docker's network</summary>
To obtain the network interface name from Docker's <code>kira_network</code>, execute <code>docker network ls</code> to list all Docker networks. Identify the ID of <code>kira_network</code> and prepend <code>br-</code> to this ID. Verify the accuracy of the network interface name by using the <code>ip a</code> command and locating the interface in the output.

</details>

<br>

Additionally, the subnet of the `kira_network` interface should be included in the `validator` zone's rich rule: `rule family="ipv4" source address="subnet_ip" accept`. To get docker's network subnet execute: `docker network inspect kira_network`

<strong>Note:</strong> all docker related information can be obtained through `docker sdk`


## Docker Setup

For Docker, make sure there are two existing containers, namely sekaid and interx, along with a volume named kira_volume.

The `sekaid` container should have the following ports forwarded: `RpcPort`, `P2PPort`, and `PrometheusPort`, with values obtained from the `config.toml` file.

The `interx` container should have the `InterxPort` forwarded with the value specified in the config.toml file.

Both containers must have a mounted volume using the `VolumeName` value from the `config.toml` file.
	
	
	
	
	