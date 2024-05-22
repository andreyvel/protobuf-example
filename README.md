# dynamic-protobuf

## Install protoc
GitHub releases https://github.com/protocolbuffers/protobuf/releases

```shell
wget https://github.com/protocolbuffers/protobuf/releases/download/v26.1/protoc-26.1-linux-x86_64.zip
sudo unzip -o protoc-26.1-linux-x86_64.zip -d /usr/local bin/protoc
sudo unzip -o protoc-26.1-linux-x86_64.zip -d /usr/local 'include/*'
protoc --version
```

## Configure New project

```shell
go mod init protobuf-dynamic-go
go get github.com/stretchr/testify/assert
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

## Generate GO code

```shell
protoc --proto_path=./schemas --go_out=. --go_opt=Mmsg-v1.proto=impl.v1 msg-v1.proto
protoc --proto_path=./schemas --go_out=. --go_opt=Mmsg-v2.proto=impl.v2 msg-v2.proto

protoc --proto_path=./schemas --go_out=. --go_opt=Mmsg-pack.proto=impl.v1 msg-pack.proto
```

## Generate descriptor

```shell
protoc --descriptor_set_out=./schemas/user.desc ./schemas/user.proto
```

## Show descriptor

```shell
protoc < ./schemas/user.desc --decode=google.protobuf.FileDescriptorSet google/protobuf/descriptor.proto
```

```shell
file {
  name: "schemas/user.proto"
  package: "api.v1"
  dependency: "google/protobuf/timestamp.proto"
  message_type {
    name: "PhoneNumber"
    field {
      name: "number"
      number: 1
      label: LABEL_OPTIONAL
      type: TYPE_STRING
      json_name: "number"
    }
    field {
      name: "type"
      number: 2
      label: LABEL_OPTIONAL
      type: TYPE_ENUM
      type_name: ".api.v1.PhoneType"
      json_name: "type"
    }
  }
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
      name: "emails"
      number: 3
      label: LABEL_REPEATED
      type: TYPE_STRING
      json_name: "emails"
    }
    field {
      name: "phones"
      number: 4
      label: LABEL_OPTIONAL
      type: TYPE_MESSAGE
      type_name: ".api.v1.PhoneNumber"
      json_name: "phones"
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
  enum_type {
    name: "PhoneType"
    value {
      name: "MOBILE"
      number: 0
    }
    value {
      name: "HOME"
      number: 1
    }
    value {
      name: "WORK"
      number: 2
    }
  }
  syntax: "proto3"
}
```