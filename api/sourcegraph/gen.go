package sourcegraph

//go:generate sh -c "env GOBIN=`pwd`/../../vendor/.bin go install ../../vendor/github.com/gogo/protobuf/protoc-gen-gogo"
//go:generate env PATH=../../vendor/.bin:$PATH protoc -I../../vendor -I../../vendor/github.com/gengo/grpc-gateway/third_party/googleapis -I../../../../.. -I. --gogo_out=Mgoogle/api/annotations.proto=github.com/gengo/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. sourcegraph.proto

//go:generate sh -c "env GOBIN=`pwd`/../../vendor/.bin go install ../../vendor/github.com/gengo/grpc-gateway/protoc-gen-grpc-gateway"
//go:generate env PATH=../../vendor/.bin:$PATH protoc -I../../vendor -I../../vendor/github.com/gengo/grpc-gateway/third_party/googleapis -I../../../../.. -I. --grpc-gateway_out=logtostderr=true:. sourcegraph.proto
