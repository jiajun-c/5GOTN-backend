import pickle

import grpc
from concurrent import futures

import pandas as pd

import core_pb2
import core_pb2_grpc
import analyse
import train
def predict(input_data):
    # 从文件中加载模型
    print("开始预测")
    with open('model.pkl', 'rb') as f:
        model = pickle.load(f)

    # 从文件中读取变量
    with open("locs.pkl", "rb") as f:
        locs = pickle.load(f)
    with open("locToIndexOflocs.pkl", "rb") as f:
        locToIndexOflocs = pickle.load(f)
    with open("train.pkl", "rb") as f:
        train = pickle.load(f)

    loc_index = [0] * len(input_data)
    for i in range(len(input_data)):
        loc_index[i] = locToIndexOflocs[input_data[i]]
    Seqs = [loc_index]
    p = model.predict(train, Seqs, 5, 8)
    result = ([locs[i] for i in itemset] for itemset in p)
    result_list = list(result)
    count = 0

    for i in range(len(result_list)):
        for j in range(len(result_list[i])):
            count += 1
    return result_list
file_path = "temp.csv"
class CoreServicer(core_pb2_grpc.coreServicer):
    # 实现 Analyse RPC
    def Analyse(self, request, context):
        # 处理请求数据
        print("Testing Analyse RPC...")
        new_data = []
        data = request.data
        for item in data:
            item_data = [
                item.clogid,
                item.calarmcode,
                item.cneid,
                item.calarmlevel,
                item.coccurutctime,
                item.clocationinfo,
                item.clineport
            ]
            new_data.append(item_data)
        # 写入到temp.csv文件中
        df = pd.DataFrame(new_data, columns=['logid', 'alarmcode', 'neid', 'alarmlevel', 'occurtime', 'locationinfo', 'lineport'])
        #print(df)
        df.to_csv(file_path, index=False, encoding="gbk")
        # 分析数据
        all_account,level0,level1,level2,level3 = analyse.Analyse(file_path)

        # 返回响应
        return core_pb2.AnalyseResponse(
            all_account=all_account,
            level0=level0,
            level1=level1,
            level2=level2,
            level3=level3,
        )
    def Train(self, request, context):
        print("Testing Train RPC...")
        if request.active == True:
            state = train.Train(file_path)
            return core_pb2.TrainResponse(success=state)
    def Alert(self, request, context):
        print("Testing Alert RPC...")
        datarequest = request.datarequest
        # print(datarequest)
        # 进行预测
        result_list = []
        input_data = []
        for data in datarequest:
            print(data)
            input_data.append(data.clocationinfo)
        print(input_data)
        list = predict(input_data)
        for sublist in list:
            for item in sublist:
                if item not in result_list:
                    result_list.extend(sublist)
        print(result_list)
        dataoutput_list = []
        for result in result_list:
            dataoutput = core_pb2.Dataoutput(clocationinfo=result)
            dataoutput_list.append(dataoutput)
        print("一共产生了" + str(len(result_list)) + "条关联告警")
        # 返回响应
        return core_pb2.AlertResponse(dataresponse=dataoutput_list)


def serve():
    port = '8000'
    # 创建服务器
    # server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10), options=[
        ('grpc.max_send_message_length', 1000 * 1024 * 1024),
        ('grpc.max_receive_message_length', 1000 * 1024 * 1024),
    ])
    # 注册服务
    core_pb2_grpc.add_coreServicer_to_server(CoreServicer(), server)
    # 监听端口
    server.add_insecure_port('[::]:'+port)
    # 启动服务器
    server.start()
    print("Server started. Listening on port 8000.")
    # 阻塞线程
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
