setup-project:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install golang.org/x/perf/cmd/benchstat@latest
	#export PATH="$PATH:$(go env GOPATH)/bin"

mod-tidy:
	go mod tidy

mod-vendor: mod-tidy
	go mod vendor

dev:
	nodemon

start-services:
	docker build -t apiary . --progress=plain
	docker-compose down
	docker-compose up -d

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

build: mod-vendor compile-proto-go
	rm -rf bin
	go build -o bin/main .
	#CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -tags=containers -o bin/main .

benchmark-cache:
	#go test -bench=. -benchmem -count=10 -benchtime=4s ./server -timeout 30m | tee current_bench.txt
	go test -bench=. -benchmem -count=10 -benchtime=4s ./server -timeout 30m | tee new_bench.txt
	benchstat current_bench.txt new_bench.txt

compile-proto:
	find ./pb -name "*.pb.go" -exec rm {} +
	buf generate

tests:
	go test -v `go list ./... | grep -v ./pb` -race -coverprofile=coverage.out; go tool cover -html=coverage.out

pull-submodules:
	git submodule update --remote --merge --recursive
	make compile-proto
