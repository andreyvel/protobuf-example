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

protoc --descriptor_set_out=./schemas/user.desc ./schemas/user.proto

protoc < ./schemas/user.desc --decode=google.protobuf.FileDescriptorSet google/protobuf/descriptor.proto

````
file {
    name: "schemas/user.proto"
    package: "api.v1"
    dependency: "google/protobuf/timestamp.proto"
    message_type {
        name: "User"
        field {
            name: "id"
            number: 1
            label: LABEL_OPTIONAL
            type: TYPE_INT32
            json_name: "id"
        }
        field {
            name: "name"
            number: 2
            label: LABEL_OPTIONAL
            type: TYPE_STRING
            json_name: "name"
        }
        field {
            name: "email"
            number: 3
            label: LABEL_OPTIONAL
            type: TYPE_STRING
            json_name: "email"
        }
        field {
            name: "last_updated"
            number: 5
            label: LABEL_OPTIONAL
            type: TYPE_MESSAGE
            type_name: ".google.protobuf.Timestamp"
            json_name: "lastUpdated"
        }
    }
    syntax: "proto3"
}
````