protoc -I api/ --go_out=internal/ api/simple.proto
protoc -I api/ --go_out=internal/ api/enum_example.proto
protoc -I api/ --go_out=internal/ api/complex.proto