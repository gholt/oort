# ortkv

### A fast persistent clustered key/value store

"It's not the future" ~ wreese

### Installation

You need godep.

* go get -u github.com/pandemicsyn/ort
* cd $GOPATH/src/github.com/pandemicsyn/ort
* make build or make run or make packages

### Ort Daemons/Binaries/Backends

* ortd - The redis speaking storage daemon
* ort-bench - The redis speaking simple benchmark utility
* mapstore - The map based kv storage
* valuestore - gholt's sstable/bitcask like persistent value store


### Other temporary (i.e. testing) stuff

* cfs - A test fuse file system you don't want to use for anything.
* apid - The frontend api server that's the relay between fuse and ortd

# Building packages

Requires fpm.

# Testing out a POC using cfs -> apid -> redis (or ort if you want)

### Run a backend

1. You can either fire up ort or a redis-server instance 

### Run apid the api server
1. cd apid/; go run main.go

### Build and mount an instance of cfs
1. cd cfs; go build .
2. mkdir /mnt/test
3. ./cfs -debug=true /mnt/test

# Testing out ort with the gholt valuestore and replication:

1. make ring
2. Fire up instance0: `VALUESTORE_PATH=/tmp/store0 VALUESTORE_PATHTOC=/tmp/store0 VALUESTORE_OUTPUSHREPLICATIONWORKERS=2 ORT_LOCALID=0 ORT_STORETYPE=ortstore ORT_RINGFILE=/tmp/ort.ring ORT_LISTENADDR=127.0.0.1:6379 go run ortd/main.go`
3. Fire up instance1: `VALUESTORE_PATH=/tmp/store1 VALUESTORE_PATHTOC=/tmp/store1 VALUESTORE_OUTPUSHREPLICATIONWORKERS=2 ORT_LOCALID=1 ORT_STORETYPE=ortstore ORT_RINGFILE=/tmp/ort.ring ORT_LISTENADDR=127.0.0.1:6479 go run ortd/main.go`
4. Fire up instance2: `VALUESTORE_PATH=/tmp/store2 VALUESTORE_PATHTOC=/tmp/store2 VALUESTORE_OUTPUSHREPLICATIONWORKERS=2 ORT_LOCALID=2 ORT_STORETYPE=ortstore ORT_RINGFILE=/tmp/ort.ring ORT_LISTENADDR=127.0.0.1:6579 go run ortd/main.go`


