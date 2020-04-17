protos:
	cd models/protos && \
	protoc \
    --plugin=protoc-gen-go=${GOPATH}/bin/protoc-gen-go \
    --plugin=protoc-gen-micro=${GOPATH}/bin/protoc-gen-micro \
    --proto_path=${GOPATH}/src:. \
    --micro_out=../  --go_out=../ \
    *.proto
	${GOPATH}/bin/protoc-go-inject-tag  -input=models/prods.pb.go


