#!/usr/bin/env python
# coding: utf-8

import pandas as pd
from sklearn.model_selection import train_test_split
import warnings
warnings.filterwarnings("ignore")
import json
import sys
import uuid


def main(train_path, test_path, output_path):
    # 定义路径参数
    res = {}
    filename = train_path
    filenamepre = test_path
    # 读取数据集
    data = pd.read_csv(filename, header=None, encoding='gbk')
    col=data.columns.size
    features_remain = data.columns[0:col-1]
    #print(features_remain)
    result = data.columns[(col-1):col]

    # 抽取30%的数据作为测试集，其余作为训练集
    train, test = train_test_split(data, test_size = 0.3)# in this our main data is splitted into train and test
    # 抽取特征选择的数值作为训练和测试数据
    train_X = train[features_remain]
    train_y=train[result]
    test_X= test[features_remain]
    test_y =test[result]

    #使用逻辑回归进行分类，并查看其正确率
    from sklearn.linear_model import LogisticRegression
    from sklearn.metrics import accuracy_score
    lr = LogisticRegression()
    lr.fit(train_X, train_y)
    lr_y_predict = lr.predict(test_X)
    accu=accuracy_score(lr_y_predict, test_y)
    #预测
    testset = pd.read_csv(test_path, header=None)
    lr_predict = lr.predict(testset)
    testset = testset.values.tolist()
    lr_predict =list(lr.predict(testset))

    for i in range(len(lr_predict)):
        testset[i].append(float(lr_predict[i]))
    # 返回结果
    res['output'] = testset
    res['accuracy'] = accu
    with open(output_path, 'w') as f:
        f.write(json.dumps(res))
    print('{}@{}'.format(output_path, accu))


if __name__ == "__main__":
    train_path = sys.argv[1]
    test_path = sys.argv[2]
    output_path = 'data/{}/prediction/{}.json'.format(sys.argv[3], uuid.uuid1())
    main(train_path, test_path, output_path)






