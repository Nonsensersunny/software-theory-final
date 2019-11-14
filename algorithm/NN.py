#!/usr/bin/env python2
# -*- coding: utf-8 -*-
import numpy as np
import pandas as pd
import json
import sys
import uuid
 
def initialize_parameters(n_x, n_h, n_y):
    W1 = np.random.randn(n_h,n_x) * 0.01
    b1 = np.zeros((n_h,1))
    W2 = np.random.randn(n_y,n_h) * 0.01
    b2 = np.zeros((n_y,1))
  
    assert (W1.shape == (n_h, n_x))
    assert (b1.shape == (n_h, 1))
    assert (W2.shape == (n_y, n_h))
    assert (b2.shape == (n_y, 1))
    
    parameters = {"W1": W1,
                  "b1": b1,
                  "W2": W2,
                  
                  "b2": b2}
    
    return parameters
 
def forward_propagation(X, parameters):
   
    W1 = parameters["W1"]
    b1 = parameters["b1"]
    W2 = parameters["W2"]
    b2 = parameters["b2"]
 
    # Implement Forward Propagation to calculate A2 (probabilities)
    Z1 = np.dot(W1,X)+b1
    A1 = np.tanh(Z1)
    Z2 = np.dot(W2,A1)+b2
    A2 = sigmoid(Z2)
    #print(W2.shape)
    #print(A2.shape)
    
    cache = {"Z1": Z1,
             "A1": A1,
             "Z2": Z2,
             "A2": A2}
    
    return A2, cache
 
def compute_cost(A2, Y, parameters):
    #计算损失时，用的是转换之后的标签
    #m = Y.shape[1] # number of example
 
    # Compute the cross-entropy cost
    logprobs = -np.multiply(np.log(A2),Y)
    cost = np.sum(logprobs)/A2.shape[1]
    #print(logprobs.shape)
    #print(cost.shape)
    return cost
 
def backward_propagation(parameters, cache, X, Y):
    #反向传播时，用的也是转换之后的标签Y->label_train
    m = X.shape[1]
    
    W1 = parameters["W1"]
    W2 = parameters["W2"]
        
    A1 = cache["A1"]
    A2 = cache["A2"]
    
    # Backward propagation: calculate dW1, db1, dW2, db2. 
    dZ2 = A2-Y
    dW2 = np.dot(dZ2,A1.T)/m
    db2 = np.sum(dZ2,axis=1,keepdims=True)/m
    dZ1 = np.dot(W2.T,dZ2)*(1 - np.power(A1, 2))
    dW1 =  np.dot(dZ1,X.T)/m
    db1 = np.sum(dZ1,axis=1,keepdims=True)/m
    
    grads = {"dW1": dW1,
             "db1": db1,
             "dW2": dW2,
             "db2": db2}
    
    return grads
 
# GRADED FUNCTION: update_parameters
 
def update_parameters(parameters, grads, learning_rate = 1.2):
   
    W1 = parameters["W1"]
    b1 = parameters["b1"]
    W2 = parameters["W2"]
    b2 = parameters["b2"]
    
    dW1 = grads["dW1"]
    db1 = grads["db1"]
    dW2 = grads["dW2"]
    db2 = grads["db2"]
 
    # Update rule for each parameter
    W1 = W1-learning_rate*dW1
    b1 = b1-learning_rate*db1
    W2 = W2-learning_rate*dW2
    b2 = b2-learning_rate*db2
   
    parameters = {"W1": W1,
                  "b1": b1,
                  "W2": W2,
                  
                  "b2": b2}
    
    return parameters
 
#softmax函数定义
def softmax(x):
    exp_scores = np.exp(x)
    probs = exp_scores / np.sum(exp_scores, axis=1, keepdims=True)
    return probs
def sigmoid(x):
    s=1/(1+np.exp(-x))
    return s

def main(train_path, test_path, output_path):
 
    data = pd.read_csv(train_path,header = None)

    train = np.array(data)

    X = train[:,0:-1]
    Y = train[:,-1]
    Y = Y.reshape((Y.shape[0],1))
    
    X = X.T
    Y = Y.T
    n_x = X.shape[0]
    n_h = 32
    n_y = 2
    
    parameters = initialize_parameters(n_x, n_h, n_y)
    

    num_iterations = 5000
    nuber = Y.shape[1]
    label_train=np.zeros((2,nuber))

    for k in range(nuber):
        if Y[0][k] == 0:
            label_train[0][k] = 1
        if Y[0][k] == 1:
            label_train[1][k] = 1
        
    for i in range(0, num_iterations):
        learning_rate=0.00095
        A2, cache = forward_propagation(X, parameters)
            
        pre_model = softmax(A2)
        cost = compute_cost(pre_model, label_train, parameters)
    
        grads = backward_propagation(parameters, cache, X, label_train)
    
        parameters = update_parameters(parameters, grads, learning_rate)  
            
    data = pd.read_csv(train_path,header = None)
    test = np.array(data)

    X_test = test[:,0:len(test[0])-1]
    Y_label = test[:,len(test[0])-1]
    Y_label = Y_label.reshape((Y_label.shape[0],1))

    X_test = X_test.T
    Y_label = Y_label.T

    count = 0
    a2,xxx=forward_propagation(X_test,parameters)

    answer_list = []

    for j in range(Y_label.shape[1]):
        #np.argmax返回最大值的索引
        predict_value=np.argmax(a2[:,j])
        answer_list.append(np.argmax(a2[:,j]))
        #在val上计算准确度
        if predict_value==int(Y_label[:,j]):
            count=count+1
            
    accu = float(count)/Y_label.shape[1]
    data = pd.read_csv(test_path,header = None)
    test = np.array(data)
    #print test.shape

    X_test = test[:,0:len(test[0])]
    #Y_label = test[:,-1]
    #Y_label = Y_label.reshape((Y_label.shape[0],1))
    X_test = X_test.T
    #Y_label = Y_label.T
    #print Y_label

    count = 0
    a2,xxx=forward_propagation(X_test,parameters)

    answer_list = []
    for j in range(len(X_test[0])):
        #np.argmax返回最大值的索引
        predict_value=np.argmax(a2[:,j])
        answer_list.append(float(predict_value))
        #在val上计算准确度
        
    res = {}
    res["output"] = answer_list
    res["accuracy"] = accu
    with open(output_path, "w") as f:
        f.write(json.dumps(res))
    print(f'{answer_list}@{accu}')
    
if __name__ == "__main__":
    train_path = sys.argv[1]
    test_path = sys.argv[2]
    output_path = f'data/{sys.argv[3]}/prediction/{uuid.uuid1()}.json'
    main(train_path, test_path, output_path)
