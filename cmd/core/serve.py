import grpc
import core_pb2_grpc
import core_pb2
import CorrelationAnalysis
import predict
import Train
from concurrent import futures
import grpc

class CoreServicer(core_pb2_grpc.coreServicer):
    # 实现Train方法
    def Train(self, request, context):
        # 进行训练，并返回结果
        data = request.file_content
        all_account, level0, level1, level2, level3 = Train(data)
        response = core_pb2.TrainResponse(all_account=all_account, level0=level0, level1=level1, level2=level2, level3=level3)
        return response

    # 实现CorrelationAnalysis方法
    def CorrelationAnalysis(self, request, context):
        # 进行相关性分析，并返回结果
        data = request.file_content
        state = CorrelationAnalysis(data)
        response = core_pb2.CorrelationAnalysisResponse(success=state, model_name="model", model_content=b"content", train_name="train", train_content=b"content", loc_name="loc", loc_content=b"content", locToIndexOflocs_name="locToIndexOflocs", locToIndexOflocs_content=b"content")
        return response

    # 实现Alert方法
    def Alert(self, request, context):
        # 进行预测，并返回结果
        data_request = request.datarequest
        input_data = []
        for data in data_request:
            input_data.append(data.clocationinfo)
        result_list = predict(input_data)
        data_outputs = []
        for result in result_list:
            data_outputs.append(core_pb2.Dataoutput(clocationinfo=result))
        response = core_pb2.AlertResponse(dataresponse=data_outputs, predict_resutlt="result")
        return response

def serve():
    port = '8080'
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    core_pb2_grpc.add_coreServicer_to_server(CoreServicer(), server)
    server.add_insecure_port('[::]:'+port)
    server.start()
    print("server started, listening on port 8080")
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
