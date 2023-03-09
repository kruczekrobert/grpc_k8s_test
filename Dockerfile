FROM golang
COPY . /go/src/grpc_k8s
WORKDIR /go/src/grpc_k8s
RUN go get .
RUN go build .
EXPOSE 3000 8080