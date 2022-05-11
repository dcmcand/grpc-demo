# A toy gRPC server for learning.

To build the proto files run `make build-proto`. This requires protoc and go-grpc plugin to be installed.

To build the server, run `make build-server` and to run the server run `make run-server`. This requires docker to be installed.
The server will start and listen on port 9000.

To run the client, run `make run-client`. This requires Go 1.16 or higher to be installed.

To edit the message, edit the `msg` variable at the top of the `client/client.go` file.

To edit the number of repeats, edit the `repeats` variable at the top of the `client/client.go` file.