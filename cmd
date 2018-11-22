 protoc -I /usr/local/include -I ./  --go_out=plugins=grpc:./example  ./example/example.proto


packr && go build && protoc -I /usr/local/include -I  ./  --plugin=protoc-gen-zap=protoc-gen-zap  --zap_out=:./example  ./example/example.proto
