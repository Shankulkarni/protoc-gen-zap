# protoc-gen-zap

Generates code for zap core.ObjectMarshaler from proto messages for faster marshaling and redact sensitive information in logs.

For more info on zap core.ObjectMarshaler, please refer: https://github.com/uber-go/zap/blob/master/zapcore/marshaler.go
### Example
* Open ./example/example.proto
* Tag the desired field of message with tag zap.redact
    ```
    string first_name = 2 [(zap.redact) = true];
    ```
 * Run following cmd to build zap-gen and generate its .zap file
   ```
   go build && protoc -I  ./  --plugin=protoc-gen-zap=protoc-gen-zap  --zap_out=:./example  ./example/example.proto 
     ````
 * Example.pb.zap.go file is generated in example/example directory
