.PHONY: proto
proto:
	find . -name "*.proto" | xargs -L 1 protoc --go_out=plugins=grpc,paths=source_relative:. -I.

.PHONY: breaking_change
breaking_change:
	buf check breaking --against .git#branch=main
	
.PHONY: lint
lint:
	buf check lint
	
.PHONY: test
test:
	go test ./... -v 
	
.PHONY: generate
generate:
	go generate ./...