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
        data_reply = data_pb2.DataReply()
        data_list = []
        # 假设这里要返回一个包含5个Data对象的数组
        for i in range(1,6):
            data = data_pb2.DataRequest()
            # 填充data对象的属性值
            data.clogid = i
            data.cneid = i
            data.clineport = str(i)  # 必须转换为字符串类型
            data.clocationinfo = f"location_{i}"
            data.coccurutctime = f"2022-01-01 00:0{i}:00"
            data.calarmcode = i
            data.calarmlevel = i
            data_list.append(data)
        data_reply.data.extend(data_list)
        return data_reply

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