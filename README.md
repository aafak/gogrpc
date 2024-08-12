# Go lang grpc examples

go mod init github.com/aafak/gogrpc
go: creating new go.mod: module github.com/aafak/gogrpc

# generate proto
personel\gogrpc>protoc --go_out=./protos --go-grpc_out=./protos --go_opt=Mprotos/user/user.proto=/user --go-grpc_opt=Mprotos/user/user.proto=/user ./protos/user/user.proto


# For ongoing adition of proto
\personel\gogrpc>protoc --go_out=./protos --go-grpc_out=./protos --go_opt=Mprotos/user/user.proto=/user --go-grpc_opt=Mprotos/user/user.proto=/user ./protos/user/user.proto  --go_opt=Mprotos/virtualmachine/virtualmachine.proto=/virtualmachine --go-grpc_opt=Mprotos/virtualmachine/virtualmachine.proto=/virtualmachine ./protos/virtualmachine/virtualmachine.proto

