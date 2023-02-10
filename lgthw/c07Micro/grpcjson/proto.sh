# grpc service
protoc --go-grpc_out=./keyvalue keyvalue/keyvalue.proto

# protobuf
protoc --go_out=keyvalue  keyvalue/keyvalue.proto