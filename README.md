# Installation Process 

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

```sudo bash -c 'echo "{\"iptables\": false}" > /etc/docker/daemon.json && systemctl restart docker && systemctl restart firewalld'```

## Disabling sudo request for docker 
```sudo usermod -aG docker $USER```




## Installing KM2

```git clone -b development https://github.com/MrLutik/kira2.0.git```

```cd kira2.0```

```go build -o ./kira2 ./cmd/kira2/main.go```

```sudo mv ./kira2  /usr/local/go/bin/kira2```

### Check if Kira2 was installed 

```kira2 --help```




# Node initialization process 

## To initialize new network 

```export GITHUB_TOKEN=ghp_zQIII30pN114wyJv5rpNyfqVxjXpws3UfjYu``` 

for now it is required to have github access token and set it as variable name 
for example




``` kira2 init new --log-level info ```











sdf
