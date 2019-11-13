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
    result = data.columns[(col-1):col]

    # 抽取30%的数据作为测试集，其余作为训练集
    train, test = train_test_split(data, test_size = 0.3)# in this our main data is splitted into train and test
    # 抽取特征选择的数值作为训练和测试数据
    train_X = train[features_remain]
    train_y=train[result]
    test_X= test[features_remain]
    test_y =test[result]

    #使用SGDC进行分类，并查看其分类准确率
    from sklearn.linear_model import SGDClassifier
    from sklearn.metrics import accuracy_score
    sgdc = SGDClassifier()
    sgdc.fit(train_X, train_y)
    sgdc_y_predict = sgdc.predict(test_X)
    accu = accuracy_score(sgdc_y_predict, test_y)
    #预测
    testset = pd.read_csv(filenamepre, header=None)
    #print(testset.head(None))
    sgdc_predict = sgdc.predict(testset)
    testset = testset.values.tolist()
    sgdc_predict =list(sgdc.predict(testset))

    for i in range(len(sgdc_predict)):
        testset[i].append(float(sgdc_predict[i]))

    # 返回结果
    res['output'] = testset
    res['accuracy'] = accu
    with open(output_path, 'w') as f:
        f.write(json.dumps(res))
    print(f'{output_path}@{accu}')


if __name__ == "__main__":
    train_path = sys.argv[1]
    test_path = sys.argv[2]
    output_path = f'data/{sys.argv[3]}/prediction/{uuid.uuid1()}.json'
    main(train_path, test_path, output_path)

