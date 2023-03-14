# 5GOTN-backend

The 5G challenge competition backend

## 1. Architecture

Just look at the dictionary :)

The core part is the cmd package, in this package, all the services will provide the rpc server (use the grpc

Out proto file should be put in the pb dictionary.

```md
.
├── api
│   └── route.go
├── app
│   └── main.go
├── client
├── cmd
│   └── core
│       └── main.py
├── dal
├── docker-compose.yml
├── dockerfile
├── pb
│   └── demo.pb
└── README.md
```

## 2. Run this project

Install the grpc requirements for golang and python first [doc link](https://grpc.io/docs/languages/go/quickstart/)

Generate the grpc code for python server and the golang client.

```shell
make rpc

cd cmd/core/pb
make rpc
```

```shell
python main.cpp

go run ./app
```

```shell
docker-compose up -d
```e

Start the services

Not implemented yet :)
