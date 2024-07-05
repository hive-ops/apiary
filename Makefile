mod-tidy:
	go mod tidy

compile-proto-go:
	#go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	#go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	#export PATH="$PATH:$(go env GOPATH)/bin"
	find ./pb -name "*.pb.go" -exec rm {} +
	protoc \
		--go_out=pb \
		--go_opt=paths=source_relative \
		--go-grpc_out=pb \
		--go-grpc_opt=paths=source_relative \
		proto/*.proto
