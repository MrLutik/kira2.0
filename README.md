# Installation Process 

## System Requirements

To run this application, your system must meet the following minimum requirements:

- **CPU**: 2 cores
- **RAM**: 4 GB
- **Operating System**: Ubuntu based system 18.04 or later 

**Note:** Due to the interactions between firewalld and manipulation with Docker's iptables settings, it is highly recommended to run kira2 in an isolated environment. This precaution helps to avoid potential conflicts and ensures more stable operation.

## Upgrading system
`sudo apt update`

`sudo apt upgrade`



## Installing dependencies

`sudo apt install git docker.io firewalld`



## Installing golang

`wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz `

`sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz`

`echo 'export PATH=$PATH:/usr/local/go/bin' | sudo tee -a /etc/profile`




## Disabling ip tables for docker 

To disable IP tables for Docker, execute the following command. This command updates the Docker daemon configuration and restarts the Docker and Firewalld services:

`sudo bash -c 'echo "{\"iptables\": false}" > /etc/docker/daemon.json && systemctl restart docker && systemctl restart firewalld'`

**Note:** This action modifies Docker's default networking behavior and might affect how containers access network resources. Ensure this change aligns with your network security policies and requirements.

## Disabling Sudo Requirement for Docker Commands

`sudo usermod -aG docker $USER`

`newgrp docker`

Or, log out and log back into your user session to apply these changes.




## Installing Kira2

`git clone -b development https://github.com/MrLutik/kira2.0.git`

`cd kira2.0`

`go build -o ./kira2 ./cmd/kira2/main.go`

`sudo mv ./kira2  /usr/local/go/bin/kira2`

### Check if Kira2 was installed 

`kira2 --help`




# Node initialization process 

## Initializing new network 
For initial setup, it is required to have a GitHub access token (browse only). Set this token as an environment variable with the following command:

`export GITHUB_TOKEN=ghp_zQIII30pN114wyJv5rpNyfqVxjXpws3UfjYu`

**Note:** The token provided above is an example. [Follow this guide to obtain your personal access token.](https://docs.github.com/en/enterprise-server@3.9/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens)


After setting up the environment variable, initialize the network with:

`kira2 init new --log-level info`

<details>

  <summary>Initialize kira2 with specific sekai and interx versions</summary>

By default kira2 initialize node with latest version of sekai and interx but you can choose specific versions with <strong>--interx-version</strong> and <strong>--sekai-version</strong> flags.


For example: <pre>kira2 init new --interx-version v0.4.43 --sekai-version v0.3.41</pre>

<strong>Note:</strong> some versions of sekai and interx are not compatible with each other.
</details>


## Join to existing network
For joining, it is required to have Github access token (browse only).

`export GITHUB_TOKEN=ghp_zQIII30pN114wyJv5rpNyfqVxjXpws3UfjYu`

**Note:** The token provided above is an example. [Follow this guide to obtain your personal access token.](https://docs.github.com/en/enterprise-server@3.9/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens)

Once the environment variable is set, you can initialize your node with the following command:

`kira2 init join --ip <ip_of_the_node_to_join>`

The **join** command has additional flags that are essential for making a proper connection, if the node you are connecting to uses a custom configuration.

<details> 
<summary>join flags</summary>

- --p2p-port string         Sekaid P2P port of the validator (default "26656")
- --rpc-port string         Sekaid RPC port of the validator (default "26657")

- -h, --help                    help for join
- --interx-port string      Interx port of the validator (default "11000")
- --interx-version string   Set this flag to choose what interx version will be initialized (default "latest")
- --ip string               IP address of the validator to join
- --p2p-port string         Sekaid P2P port of the validator (default "26656")
- --recover                 If true recover keys and mnemonic from master mnemonic, otherwise generate random one
- --rpc-port string         Sekaid RPC port of the validator (default "26657")
- --sekai-version string    Set this flag to choose what sekai version will be initialized (default "latest")
</details>
<br>

For example, to connect to a node operated by kira1, you need to specify  the **--rpc-port**, **--p2p-port** ports individually and also use **--sekai-version** and **--interx-version** flags. 

### Joining to node with kira1

To join a node configured by **kira1** with version **v0.11.27**, execute the following command:

`kira2 init join --ip <ip_of_the_node_to_join>  --log-level info  --interx-version v0.4.35  --sekai-version v0.3.17 --p2p-port 36656 --rpc-port 36657`

## Key recovering

Both the **kira2 init** new and **kira2 init join** commands include a recovery option, which can be activated using the `--recover` flag. For instance: 

``kira2 init join --ip <ip_of_the_node_to_join>  --log-level info  --interx-version v0.4.35  --sekai-version v0.3.17 --p2p-port 36656 --rpc-port 36657 --recover``

**Note:** When initializing a node with `--recover` option, **kira2** will prompt you for your BIP39 24-word ***master mnemonic***. The key delivery algorithm used in **kira2** is the same as in **kira1**, allowing for seamless migration from a **kira1** node to **kira2**.








