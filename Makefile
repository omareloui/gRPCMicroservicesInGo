all:

generate-proto:
	@echo "Generating the protobuffers..."
	@mkdir -p proto/go
	@protoc -I ./proto \
		--go_out=./proto/go/ \
		--go_opt=paths=source_relative \
		--go-grpc_out=./proto/go \
		--go-grpc_opt=paths=source_relative \
		./proto/**/*.proto
