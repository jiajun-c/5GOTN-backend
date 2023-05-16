rpc:
	 protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	pb/core.proto

# 安装python的环境
install:
	python -m pip install grpcio
	python -m pip install grpcio-tools