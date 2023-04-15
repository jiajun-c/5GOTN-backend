# -*- coding: utf-8 -*-
# @Time : 2023/3/25 21:21
# @Author : Zhangmiaosong
import operator
import pickle
from datetime import datetime

import numpy as np
from numpy import array

import DataProcessing as dp
def Predict(input_data):
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
    # Seqs = [[1657, 278, 253, 1664]]
    # print(Seqs)
    p = model.predict(train, Seqs, 5, 15)
    result = ([locs[i] for i in itemset] for itemset in p)
    result_list = list(result)
    count = 0

    for i in range(len(result_list)):
        for j in range(len(result_list[i])):
            print(result_list[i][j])
            count += 1
    print("一共产生了" + str(count) + "条关联告警")
    return result_list
if __name__ == '__main__':
    '''
    接收的传入参数为告警的位置信息，输出得到的结果
    '''
    # 从文件中加载模型
    with open('model.pkl', 'rb') as f:
        model = pickle.load(f)

    # 从文件中读取变量
    with open("locs.pkl", "rb") as f:
        locs = pickle.load(f)
    with open("locToIndexOflocs.pkl", "rb") as f:
        locToIndexOflocs = pickle.load(f)
    with open("train.pkl", "rb") as f:
        train = pickle.load(f)


    print("请输入告警的位置信息")
    input_data = input().replace("\"", "").split(",")
    input_data = [x.strip() for x in input_data]
    print(input_data)
    loc_index=[0] * len(input_data)
    for i in range(len(input_data)):
        loc_index[i] = locToIndexOflocs[input_data[i]]
    Seqs = [loc_index]
        #Seqs = [[1657, 278, 253, 1664]]
        #print(Seqs)
    p = model.predict(train,Seqs,5,15)
    result = ([locs[i] for i in itemset] for itemset in p)
    result_list = list(result)
    count = 0

    for i in range(len(result_list)):
        for j in range(len(result_list[i])):
            print(result_list[i][j])
            count+=1
    print("一共产生了" + str(count) + "条关联告警")
