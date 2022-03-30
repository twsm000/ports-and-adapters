FROM golang:1.18-alpine3.15

EXPOSE 9000

RUN apk update \
    && apk add --no-cache \
    mysql-client \
    build-base

RUN mkdir /app
WORKDIR /app

# Commands from <HOST> to <CONTAINER> workdir
# copy go.mod to container workdir
COPY go.mod .
# copy go.sum to container workdir
COPY go.sum .
RUN go mod download
# copy all remaining files/folders to container workdir
COPY . .
COPY ./grpc_entrypoint.sh /usr/local/bin/grpc_entrypoint.sh
RUN /bin/chmod +x /usr/local/bin/grpc_entrypoint.sh

# Building the binary
RUN go build cmd/main.go
RUN mv main /usr/local/bin/main

CMD ["main"]
ENTRYPOINT ["grpc_entrypoint.sh"]
