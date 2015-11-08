Item Price Service (IPS)
===
A microservice to provide listing price recommendation given an item.

Requirement
-----------
- Go (>=1.4)

Setup Go
------------
Example:

    wget https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz
    tar -xzf go1.5.1.linux-amd64.tar.gz
    export GOROOT=$HOME/go
    export PATH=$PATH:$GOROOT/bin
    export GOPATH=$HOME/work

Checkout
-------------
    mkdir -p $HOME/work/src/github.com/jiang0131/ && cd $HOME/work/src/github.com/jiang0131
    git clone https://github.com/jiang0131/ips.git

Build
-------------
    go get
    go build

Run
-------------
    ./ips

Endpoints
-------------
* GET /item-price-service/

  Query parameters:
  - item [string]
  - city [string]