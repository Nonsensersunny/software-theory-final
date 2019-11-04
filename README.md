# software-theory-final

### 1. 数据字典

- User

  | 列名     | 类型    | 说明       |
  | :------- | ------- | ---------- |
  | id       | varchar | 用户存储ID |
  | mail     | varchar | 用户邮箱   |
  | password | varchar | 用户密码   |

- Algorithm

  | 列名        | 类型    | 说明         |
  | ----------- | ------- | ------------ |
  | id          | varchar | 算法存储ID   |
  | path        | varchar | 算法存储路径 |
  | description | varchar | 算法描述     |

- Dataset

  | 列名        | 类型    | 说明           |
  | ----------- | ------- | -------------- |
  | id          | varchar | 数据集存储ID   |
  | uid         | varchar | 用户ID（外键） |
  | path        | varchar | 数据集存储路径 |
  | description | varchar | 数据集描述     |
  | time        | time    | 数据集上传时间 |

- Prediction

  | 列名     | 类型    | 说明             |
  | -------- | ------- | ---------------- |
  | id       | varchar | 预测结果存储ID   |
  | aid      | varchar | 算法ID（外键）   |
  | did      | varchar | 数据集ID（外键） |
  | accuracy | float   | 预测结果准确度   |
  | path     | varchar | 测试集存储路径   |
  | result   | varchar | 预测结果         |
  | time     | time    | 预测操作时间     |

### 2. 后端REST规范

- 后端接口定义与测试均上传至 [SwaggerHub](https://app.swaggerhub.com/apis/B755/software-theory_final/1.0.0#/User)
