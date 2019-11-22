#!/usr/bin/env python
# coding: utf-8

from sklearn.ensemble import RandomForestClassifier
from sklearn.model_selection import GridSearchCV, cross_val_score, train_test_split
from sklearn.metrics import make_scorer, accuracy_score
import pandas as pd
import warnings
import sys
import uuid
import json
warnings.filterwarnings("ignore")

def main(train_path, test_path, output_path):
    res = {}
    data = pd.read_csv(train_path, header=None)
    col=data.columns.size
    features_remain = data.columns[0:col-1]
    # print(features_remain)
    result = data.columns[(col-1):col]
    # 抽取30%的数据作为测试集，其余作为训练集
    train, test = train_test_split(data, test_size = 0.3)# in this our main data is splitted into train and test
    # 抽取特征选择的数值作为训练和测试数据
    X_train = train[features_remain]
    y_train=train[result]
    X_test= test[features_remain]
    y_test =test[result]

    #选择分类器的类型，我没试过其他的哦，因为在这个案例中，有人做过试验发现随机森林模型是最好的，所以选了它。呜呜，我下次试试其他的
    clf = RandomForestClassifier()

    #可以通过定义树的各种参数，限制树的大小，防止出现过拟合现象哦，也可以通过剪枝来限制，但sklearn中的决策树分类器目前不支持剪枝
    parameters = {'n_estimators': [4, 6, 9], 
                'max_features': ['log2', 'sqrt','auto'], 
                'criterion': ['entropy', 'gini'],        #分类标准用熵，基尼系数
                'max_depth': [2, 3, 5, 10], 
                'min_samples_split': [2, 3, 5],
                'min_samples_leaf': [1,5,8]
                }

    #以下是用于比较参数好坏的评分，使用'make_scorer'将'accuracy_score'转换为评分函数
    acc_scorer = make_scorer(accuracy_score)

    #自动调参，GridSearchCV，它存在的意义就是自动调参，只要把参数输进去，就能给出最优化的结果和参数
    #GridSearchCV用于系统地遍历多种参数组合，通过交叉验证确定最佳效果参数。
    grid_obj = GridSearchCV(clf,parameters,scoring=acc_scorer)
    grid_obj = grid_obj.fit(X_train,y_train)

    #将clf设置为参数的最佳组合
    clf = grid_obj.best_estimator_

    #将最佳算法运用于数据中
    clf.fit(X_train,y_train)

    predictions = clf.predict(X_test)
    print(accuracy_score(y_test,predictions))
    #如果把上面多执行几次，会发现这里的执行的结果都会有所不同，取决于用了parameter的哪些参数

    #预测
    testset = pd.read_csv(test_path, header=None)
    #t = list(data.columns[0:8])
    #t_x = test[t]
    #test = testset.columns[0:10]
    #model.predict(ch)
    #model.fit(testset)
    print(testset.head(None))
    clf.predict(testset)

    with open(output_path, 'w') as f:
        f.write(json.dumps(res))
    print('{}@{}'.format(output_path, accu))

if __name__ == "__main__":
    train_path = sys.argv[1]
    test_path = sys.argv[2]
    output_path = 'data/{}/prediction/{}.json'.format(sys.argv[3], uuid.uuid1())
    main(train_path, test_path, output_path)
