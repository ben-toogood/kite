proto:
	find . -name "*.proto" | xargs -L 1 protoc --go_out=plugins=grpc,paths=source_relative:. -I.