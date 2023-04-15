import grpc
import core_pb2
import core_pb2_grpc

def run():

    with grpc.insecure_channel('[::1]:8000') as channel:
        client = core_pb2_grpc.coreStub(channel)

        # # 测试 Analyse RPC
        # print("Testing Analyse RPC...")
        # response1 = client.Analyse(core_pb2.AnalyseRequest(file_path='t_alarmlogcur.csv'))
        # print("Response: all_account=%d, level0=%d, level1=%d, level2=%d, level3=%d" % (
        #     response1.all_account, response1.level0, response1.level1, response1.level2, response1.level3))
        #
        # # 测试 Train RPC
        # print("Testing Train RPC...")
        # response2 = client.Train(core_pb2.TrainRequest(file_path='t_alarmlogcur.csv'))
        # print("Response: success=%s" % response2.success)

        # 测试 Alert RPC
        print("Testing Alert RPC...")
        request = [("1077-15-赣州寻乌-南环-OTM2:1LN4-λ2-水南方向[45]:L_PORT1/OCH-1",),
                   ("1077-11-赣州定南-南环-OTM2:8TN1-水南[49]:背板口-L_PORT7/ODU0-1",), ]
        datarequest = []
        for item in request:
            datainput = core_pb2.Datainput(clocationinfo=item[0])
            datarequest.append(datainput)
        response3 = client.Alert(core_pb2.AlertRequest(datarequest=datarequest))

        print("Response:")
        for dataoutput in response3.dataresponse:
            print("- clocationinfo=%s" % dataoutput.clocationinfo)



if __name__ == '__main__':
    run()
