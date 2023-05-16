import grpc
import pandas as pd

import core_pb2
import core_pb2_grpc
def read_data(file_path):
    '''读取初始表数据，先转成csv格式'''
    df = pd.read_csv( \
        file_path, usecols=[0, 6, 8, 13, 16, 21, 22], \
        dtype={ \
            'clogid': int, \
            'calarmcode': int, \
            'cneid': int, \
            'calarmlevel': int, \
            'coccuructime': object, \
            'clocationinfo': object, \
            'clinport': object \
            }, \
        low_memory=False, encoding="gbk")
    data = df.values
    return data
def run():

    with grpc.insecure_channel('[::1]:8000') as channel:
        client = core_pb2_grpc.coreStub(channel)

        # # 测试 Analyse RPC
        data = read_data("t_alarmlogcur.csv")
        request1 = core_pb2.AnalyseRequest()
        for row in data:
            data_struct = core_pb2.DataStruct(
                clogid=int(row[0]),
                calarmcode=int(row[1]),
                cneid=int(row[2]),
                calarmlevel=int(row[3]),
                coccurutctime=str(row[4]),
                clocationinfo=str(row[5]),
                clineport=str(row[6])
            )
            request1.data.append(data_struct)
        #print(request1)
        response1 = client.Analyse(request1)
        print("Response: all_account=%d, level0=%d, level1=%d, level2=%d, level3=%d" % (
            response1.all_account, response1.level0, response1.level1, response1.level2, response1.level3))
        # #
        # # # 测试 Train RPC
        # print("Testing Train RPC...")
        # response2 = client.Train(core_pb2.TrainRequest(active=1))
        # print("Response: success=%s" % response2.success)

        # 测试 Alert RPC
        # print("Testing Alert RPC...")
        # request = [("1077-15-赣州寻乌-南环-OTM2:1LN4-λ2-水南方向[45]:L_PORT1/OCH-1",),
        #            ("1077-11-赣州定南-南环-OTM2:8TN1-水南[49]:背板口-L_PORT7/ODU0-1",), ]
        # datarequest = []
        # for item in request:
        #     datainput = core_pb2.Datainput(clocationinfo=item[0])
        #     datarequest.append(datainput)
        # response3 = client.Alert(core_pb2.AlertRequest(datarequest=datarequest))
        #
        # print("Response:")
        # for dataoutput in response3.dataresponse:
        #     print("- clocationinfo=%s" % dataoutput.clocationinfo)
        print("Testing Alert RPC...")
        # request = [('1075-2-赣州信丰大塘-OADM1:IOSC_F1K[04]:OSC_LINE_W',),
        #            ('1075-2-赣州信丰大塘-OADM1:IOSC_F1K[04]:OSC_LINE_W',),
        #            ('1075-6-赣州信丰小江-OADM1:OCP_F1K[03]:INOUT-1/OCH-1',), ]
        # datarequest = []
        # for item in request:
        #     datainput = core_pb2.Datainput(clocationinfo=item[0])
        #     datarequest.append(datainput)
        try:
            request = ['1075-2-赣州信丰大塘-OADM1:IOSC_F1K[04]:OSC_LINE_W',
                       '1075-2-赣州信丰大塘-OADM1:IOSC_F1K[04]:OSC_LINE_W',
                       '1075-6-赣州信丰小江-OADM1:OCP_F1K[03]:INOUT-1/OCH-1']

            datarequest = [core_pb2.Datainput(clocationinfo=item) for item in request]
            response3 = client.Alert(core_pb2.AlertRequest(datarequest=datarequest))

            print("Response:")
            for dataoutput in response3.dataresponse:
                print("- clocationinfo=%s" % dataoutput.clocationinfo)
        except Exception as e:
            print("Error")



if __name__ == '__main__':
    run()
