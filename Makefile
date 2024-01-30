# Set arguments for grpcui
ifeq (grpcui,$(firstword $(MAKECMDGOALS)))
  PORT := $(word 2, $(MAKECMDGOALS))
endif


all:
	@echo "TODO"

generate-proto:
	@echo "Generating the protobuffers..."
	@mkdir -p proto/go
	@protoc -I ./proto \
		--go_out=./proto/go/ \
		--go_opt=paths=source_relative \
		--go-grpc_out=./proto/go \
		--go-grpc_opt=paths=source_relative \
		./proto/**/*.proto


grpcui:
	grpcui --plaintext 127.0.0.1:$(PORT)
