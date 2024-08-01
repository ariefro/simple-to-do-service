.PHONY: proto
proto:
	rm -f api/protogen/*.go
	protoc --proto_path=api/proto --go_out=api/protogen --go_opt=paths=source_relative \
    --go-grpc_out=api/protogen --go-grpc_opt=paths=source_relative \
    api/proto/**/*.proto
