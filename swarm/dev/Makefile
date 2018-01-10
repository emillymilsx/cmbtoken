.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/CoinMarketBrasil/cmbtoken/cmd/swarm
	go build -o bin/geth     github.com/CoinMarketBrasil/cmbtoken/cmd/geth
	go build -o bin/bootnode github.com/CoinMarketBrasil/cmbtoken/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/CoinMarketBrasil/cmbtoken/swarm/...
