### Generate Golang file form *.proto file

```
docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc:0.4.0 --go_out=plugins=grpc:. -I. <filename>.proto
```

#### Reference

https://developers.google.com/protocol-buffers/docs/gotutorial

https://github.com/grpc/grpc-go#compiling-error-undefined-grpcsupportpackageisversion

https://github.com/grpc/grpc-go/tree/master/examples/helloworld
