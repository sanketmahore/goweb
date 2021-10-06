FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /code

COPY . .

WORKDIR /code/admin_microservice

RUN go build -o admin_microservice .

CMD ["./admin_microservice"]