build-proto:
	protoc --go-grpc_out=. --go_out=. echo.proto

build-server:
	docker build -t grpc-test-server .
	
run-server:
	docker run --rm -p 9000:9000 grpc-test-server