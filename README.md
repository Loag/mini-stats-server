# mini-stats


## usage

``` bash
# build rpc
cd protobuf && protoc --go_out=../rpc --go_opt=paths=source_relative --go-grpc_out=../rpc --go-grpc_opt=paths=source_relative *.proto
```

### endpoints

``` 
  /ingest 
```