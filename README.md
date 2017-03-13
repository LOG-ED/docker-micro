## Setup

### Requirements

1. Download the protol buffers compiler from https://github.com/google/protobuf/releases   

2. Run the following command to install the Go protocol buffers plugin:   
   `$ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}` 

3. If changes are made to the Protocol Buffer file use the following command to regenerate:  
   `$ protoc -I$GOPATH/src --go_out=plugins=micro:$GOPATH/src \ $GOPATH/src/github.com/log-ed/docker-micro/proto/task.proto`


