FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /code

COPY . .

WORKDIR /code/grpc_server_workdir

RUN go build -o grpc_server .

CMD ["./grpc_server"]