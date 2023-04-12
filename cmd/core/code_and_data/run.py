import CorrelationAnalysis
import predict
import train
import xmlrpc.client
if __name__ == '__main__':

    # Set up RPC client
    server = xmlrpc.client.ServerProxy("http://localhost:8000")

    data = server.get_data()
    all_account, level0, level1, level2, level3 = Train(data)
    # Send results back to server
    server.send_results(all_account, level0, level1, level2, level3)

    state = CorrelationAnalysis(data)
    server.send_results(state)
    input_data = server.get_data()
    result_list = predict(input_data)
    server.send_results(result_list)