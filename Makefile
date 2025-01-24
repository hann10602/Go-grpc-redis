proto/gen:
	protoc --go_out=. \
		--go_opt=paths=source_relative \
		--go_grpc_out=. \
		--go_grpc_opt=paths=source_relative \
		$(shell find . -name '*.proto')