import math
import csv
import random
import sys
import json
import uuid

class Bayes(object):
    def __init__(self, trainData):
        self.trainData = trainData
        self.model_para = {}
    def tarin_bayesModel(self):# 将训练集按照类别进行提取
        separated_class = self.separateByClass()
        # vectors是列表，包含的是每个类别对应的向量集
        for classValue, vectors in separated_class.items():
            # 将每一个类别的均值和方差保存在对应的键值对中
            self.model_para[classValue] = self.summarize(vectors)
        return self.model_para
    def mean(self, numbers):# 计算均值
        return sum(numbers) / float(len(numbers))
    def stdev(self, numbers):# 计算方差，注意是分母是n-1
        avg = self.mean(numbers)
        variance = sum([pow(x - avg, 2) for x in numbers]) / float(len(numbers) - 1)
        return math.sqrt(variance)
    # 对每一类样本的每个特征计算均值和方差，结果保存在列表中，依次为第一维特征、第二维特征等...的均值和方差
    def summarize(self, vectors):
        # zip利用 * 号操作符，可以将不同元组或者列表压缩为为列表集合。用来提取每类样本下的每一维的特征集合
        summaries = [(self.mean(attribute), self.stdev(attribute)) for attribute in zip(*vectors)]
        # 将代表类别的最后一个数据删掉，只保留均值和方差
        del summaries[-1]
        return summaries
    # 将训练集按照类别进行提取，以字典形式存放，Key为类别，value为列表，列表中包含的是每个类别对应的向量集
    def separateByClass(self):
        # 字典用于存放分类后的向量集合
        separated_class = {}
        for i in range(len(self.trainData)):
            vector = self.trainData[i]
            # vector[-1]为每组数据的类别
            if (vector[-1] not in separated_class):
                separated_class[vector[-1]] = []
            separated_class[vector[-1]].append(vector)
        return separated_class

    def calProbabilityDensity(self, x, mean, stdev):
        exponent = math.exp(-(math.pow(x - mean, 2) / (2 * math.pow(stdev, 2))))
        return (1 / (math.sqrt(2 * math.pi) * stdev)) * exponent

    # 计算待分类数据的联合概率
    def calClassProbabilities(self, inputVector):
        probabilities = {}
        for classValue, classSummaries in self.model_para.items():
            probabilities[classValue] = 1
            for i in range(len(classSummaries)):
                mean, stdev = classSummaries[i]
                x = inputVector[i]
                probabilities[classValue] *= self.calProbabilityDensity(x, mean, stdev)
        prediction = max(probabilities, key=probabilities.get)
        return prediction
    def calClassProbabilitiespre(self, inputVector):
        probabilities = {}
        for classValue, classSummaries in self.model_para.items():
            probabilities[classValue] = 1
            for i in range(len(classSummaries)):
                mean, stdev = classSummaries[i]
                x = inputVector[i]
                probabilities[classValue] *= self.calProbabilityDensity(x, mean, stdev)
        prediction = max(probabilities, key=probabilities.get)
        return prediction

result = {}

# 准备数据
def loadCsv(filename):
    lines = csv.reader(open(filename, "r"))
    dataset = list(lines)
    for i in range(len(dataset)):
        dataset[i] = [float(x) for x in dataset[i]]
    return dataset
# 将原始数据集划分为训练集和测试集，splitRatio为划分比例。
def splitDataset(dataset, splitRatio):
    trainSize = int(len(dataset) * splitRatio)
    trainSet = []
    copy = list(dataset)
    while len(trainSet) < trainSize:
        index = random.randrange(len(copy))
        trainSet.append(copy.pop(index))
    return [trainSet, copy]
# 计算分类准确率
def calAccuracy(testData, bayes):
    correct_nums = 0
    error_sum = []
    j = 0
    for i in range(len(testData)):
        if testData[i][-1] == bayes.calClassProbabilities(testData[i]):
            correct_nums += 1
        else:
#             print(testData[i])
            j += 1
    return [correct_nums, error_sum]

def main(train_path, test_path, output_path):
    """
    读取数据集进行训练和测试
    """
    filename = train_path
    # 训练集和测试集的划分比例
    splitRatio = 0.67
    dataset = loadCsv(filename)
    trainData, testData = splitDataset(dataset, splitRatio)

    bayes = Bayes(trainData)
    model = bayes.tarin_bayesModel()
    correct_nums, error_info = calAccuracy(testData, bayes)
    accu = correct_nums / len(testData) * 100.0
    # print("测试集分类准确率 %f%%" % (correct_nums / len(testData) * 100.0))
    # 读取预测集
    filenamepre = test_path
    datasetpre = loadCsv(filenamepre)
    bayespre = Bayes(datasetpre)
    for i in range(len(datasetpre)):
        datasetpre[i].append(bayes.calClassProbabilitiespre(datasetpre[i]))
    result['output'] = datasetpre
    result['accuracy'] = accu
    with open(output_path, 'w') as f:
        f.write(json.dumps(result))
    print(f'{output_path}@{accu}')
if __name__ == "__main__":
    train_path = sys.argv[1]
    test_path = sys.argv[2]
    output_path = f'data/{sys.argv[3]}/prediction/{uuid.uuid1()}.json'
    main(train_path, test_path, output_path)
