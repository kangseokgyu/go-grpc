# go-grpc

Go와 gRPC를 사용해보자.

<br>

## Go 설치

https://go.dev/doc/install

<br>

## Protocol buffer 컴파일러 설치

```
$ brew install protobuf
```

https://grpc.io/docs/protoc-installation/

<br>

## Go 플러그인 설치

```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

### 플러그인 path 설정

```
$ export PATH="$PATH:$(go env GOPATH)/bin"
```
