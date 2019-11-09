1. protoc -I api/proto --go_out=plugins=grpc:pkg/api api/proto/adder.proto
# name method, path to our protofile, method and plugins grpc, output dir, name of file

# in output file, intersting us methods AddServer and AddClient

2. go build -v ./cmd/server

3. ./server

4. evans api/proto/adder.proto
5. call Add
