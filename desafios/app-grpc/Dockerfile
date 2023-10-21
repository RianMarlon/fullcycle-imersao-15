FROM golang:1.21

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLED=1

RUN apt-get update && \
  apt-get install protobuf-compiler -y && \
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 && \
  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

CMD ["go", "run", "main.go"]
