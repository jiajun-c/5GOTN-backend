FROM golang:alpine

COPY . /otn
WORKDIR /otn
RUN go env -w GO111MODULE=on&&go env -w GOPROXY=https://goproxy.cn,direct
CMD go run ./app