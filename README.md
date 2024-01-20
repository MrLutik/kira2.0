# Installation Process 

## System Requirements

To run this application, your system must meet the following minimum requirements:

- **CPU**: 2 cores
- **RAM**: 4 GB
- **Operating System**: Ubuntu Server 20.04 or higher

## Upgrading system
```sudo apt update```

```sudo apt upgrade```



## Installing dependencies

```sudo apt install git docker.io firewalld```



## Installing golang

```wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz ```

```sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz```

```echo 'export PATH=$PATH:/usr/local/go/bin' | sudo tee -a /etc/profile```




## Disabling ip tables for docker 

To disable IP tables for Docker, execute the following command. This command updates the Docker daemon configuration and restarts the Docker and Firewalld services:

```sudo bash -c 'echo "{\"iptables\": false}" > /etc/docker/daemon.json && systemctl restart docker && systemctl restart firewalld'```

**Note:** This action modifies Docker's default networking behavior and might affect how containers access network resources. Ensure this change aligns with your network security policies and requirements.

## Disabling Sudo Requirement for Docker Commands

```sudo usermod -aG docker $USER```

```newgrp docker```

Or, log out and log back into your user session to apply these changes.




## Installing Kira2

```git clone -b development https://github.com/MrLutik/kira2.0.git```

```cd kira2.0```

```go build -o ./kira2 ./cmd/kira2/main.go```

```sudo mv ./kira2  /usr/local/go/bin/kira2```

### Check if Kira2 was installed 

```kira2 --help```




# Node initialization process 

## To initialize new network 
For initial setup, it is required to have a GitHub access token. Set this token as an environment variable with the following command:

```export GITHUB_TOKEN=ghp_zQIII30pN114wyJv5rpNyfqVxjXpws3UfjYu``` 

**Note:** The token provided above is an example. [Follow this guide to obtain your personal access token.](https://docs.github.com/en/enterprise-server@3.9/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens)


After setting up the environment variable, initialize the network with:

```kira2 init new --log-level info```












