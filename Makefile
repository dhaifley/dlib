NOVENDOR = $(shell go list ./... | grep -v vendor | grep -v node_modules)
NOVENDOR_LINTER = $(shell go list ./... | grep -v vendor | grep -v ptypes | grep -v node_modules)

all: build

fix:
	go fix $(NOVENDOR)
.PHONY: fix

vet:
	go vet $(NOVENDOR)
.PHONY: vet

lint:
	printf "%s\n" "$(NOVENDOR)" | xargs -I {} sh -c 'golint -set_exit_status {}'
.PHONY: lint

test:
	go test -v -cover $(NOVENDOR)
.PHONY: test

cleanproto:
	rm -rf ptypes/*.pb.go
.PHONY: cleanprotobuf

proto: cleanproto
	protoc ptypes/*.proto --go_out=plugins=grpc,import_path=ptypes:.
.PHONY: protobuf

clean:
	rm -rf ./bin
.PHONY: clean

build: clean proto fix vet lint test
	mkdir bin
	GOOS=linux GOARCH=386 go build -v -o ./bin/dlib.a .
	GOOS=windows GOARCH=amd64 go build -v -o ./bin/dlib.dll .
.PHONY: build
