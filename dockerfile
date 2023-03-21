FROM golang:alpine

LABEL maintainer="chengjiajun20@gmail.com"
COPY . /otn
WORKDIR /otn
RUN go env -w GO111MODULE=on&&go env -w GOPROXY=https://goproxy.cn,direct
CMD go run ./app