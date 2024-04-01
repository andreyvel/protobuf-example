# dynamic-protobuf

## Inslall protoc
https://github.com/protocolbuffers/protobuf/releases

wget https://github.com/protocolbuffers/protobuf/releases/download/v26.1/protoc-26.1-linux-x86_64.zip
sudo unzip -o protoc-26.1-linux-x86_64.zip -d /usr/local bin/protoc
sudo unzip -o protoc-26.1-linux-x86_64.zip -d /usr/local 'include/*'
protoc --version

## Configure project
go mod init protobuf-dynamic-go
go get github.com/stretchr/testify/assert
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

## Generate impl
protoc --proto_path=./schemas --go_out=. --go_opt=Mmsg-v1.proto=impl.v1 msg-v1.proto
protoc --proto_path=./schemas --go_out=. --go_opt=Mmsg-v2.proto=impl.v2 msg-v2.proto
