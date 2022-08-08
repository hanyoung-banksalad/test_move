# imageproxy

[![Actions Status](https://github.com/banksalad/imageproxy/workflows/ci/badge.svg)](https://github.com/banksalad/imageproxy/actions) ![Golang Badge](https://badgen.net/badge/Language/Go/cyan) ![GRPC Badge](https://badgen.net/badge/Use/gRPC/blue)

> image proxy service

## Docs
<!-- TODO: Update to the actual document url -->
- [TechSpec](https://docs.google.com/document/d/1dUBkovA0tv8K6XYX5qATWO1Pm2cbGH0ajxmucJvTANA/edit?usp=sharing)
- [Go in Banksalad](https://docs.google.com/document/d/1mPUGKlfA6pFLMUuUCHv54ejnUDrrldJ5z06AbvinRQA)

## Getting Started

### Start a Server
```sh
$ git config --global url."https://${GITHUB_ACCESS_TOKEN}@github.com/".insteadOf "https://github.com/"  # insert your github access token
$ make init
$ make run

# Use Docker
$ docker build --build-arg GH_ACCESS_TOKEN=${GITHUB_ACCESS_TOKEN} --tag imageproxy .  # insert your github access token
$ docker run --rm -p 18081:18081 -p 18082:18082 imageproxy

# Use Onebox
$ make deploy-to-local-k8s
```

### Test & Lint
```sh
$ make test

$ make lint
```

## APIs
<!-- TODO: Update to actual urls -->
- [imageproxy.proto](https://github.com/banksalad/idl/blob/master/protos/apis/v1/imageproxy/imageproxy.proto)
- [imageproxy.swagger.json](https://github.com/banksalad/idl/blob/master/gen/swagger/apis/v1/imageproxy/imageproxy.swagger.json)

## Directory Structure
```
.
├── client.go # dependency service 들에 대한 client
├── cmd       # server를 실행시키기 위한 main.go
│   └── ...
├── config    # 설정 파일
│   └── ...
└── server
    ├── grpc_server.go  # main gRPC server
    ├── http_server.go  # HTTP <-> gRPC 변환해주는 grpc-gateway layer
    └── handler         # gRPC server handlers
```
