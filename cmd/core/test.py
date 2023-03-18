import grpc
import Data_pb2 as pb2
import Data_pb2_grpc as pb2_grpc


def run():
    with grpc.insecure_channel('localhost:8000') as channel:
        stub = pb2_grpc.DataServiceStub(channel)
        data = pb2.Data(
            clogid=1,
            cneid=2,
            clinport="port1",
            clocationinfo="location",
            coccurutctime="2022-01-01 00:00:00",
            calarmcode=100,
            calarmlevel=1,
        )
        response = stub.SendData(data)
        print("Client received: \n" + str(response))


if __name__ == '__main__':
    run()
