# -*- coding: utf-8 -*-
# @Time : 2023/3/18 14:27
# @Author : Zhangmiaosong
from concurrent import futures
import grpc
import Data_pb2 as data_pb2
import Data_pb2_grpc as data_pb2_grpc

class DataService(data_pb2_grpc.DataServiceServicer):
    def SendData(self, request, context):
        print("Send data successfully!")
        print(request)
        return request

def serve():
    port = '8000'
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    data_pb2_grpc.add_DataServiceServicer_to_server(DataService(), server)
    server.add_insecure_port('[::]:'+port)
    server.start()
    print("server started, listening on port 8000")
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
