create:

	protoc --go_out=gen --go_opt=paths=source_relative \
        --go-grpc_out=gen --go-grpc_opt=paths=source_relative \
        proto/inmemoryservice.proto

clean:
	rm gen/proto/*.go