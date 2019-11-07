<template>
  <div class="algoControl">
    <center>
      <div class="datasetTable">
        <div style="text-align: left;font-size: 18px;margin-bottom: 10px;">
          <span style="font-size: 18px;color: #009FCC">全部预测：</span>
        </div>
        <el-table
          :data="tableData"
          border
          style="font-size:18px;border-color: #009FCC"
          :row-style="{height:'80px'}"
        >
          <el-table-column prop="id" align="center" label="序号" width="100px"></el-table-column>
          <el-table-column prop="dataset" label="算法名称" align="center" width="280px"></el-table-column>
          <el-table-column prop="algo" label="算法名称" align="center" width="280px"></el-table-column>
          <el-table-column prop="time" label="预测时间" align="center" width="280px"></el-table-column>
          <el-table-column prop="work" label="操作" align="center" width="280px">
            <template slot-scope="scope">
              <el-button style="color:#409EFF" @click="check(scope)">查看结果</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <el-dialog title="预测结果" :visible.sync="data_table" width="600px">
        <div>
          <span>本次预测准确率：</span>
          <el-input disabled v-model="data.percent" style="width:50%;margin-bottom:10px;font-color:red;"></el-input>
        </div>
        <el-table :data="data.data_yuce">
          <el-table-column
            align="center"
            v-for="(item, index) in colTitle"
            :key="index"
            :label="item"
            :prop="item"
            :height="10"
          ></el-table-column>
        </el-table>
      </el-dialog>
    </center>
  </div>
</template>

<script>
export default {
  name: "datasetControl",
  data() {
    return {
      data_table: false,
      tableData: [
        {
          id: 5,
          dataset: "乳腺癌细胞核数据集.csv",
          algo: "SVM向量机算法",
          time: "2019-10-30 下午06：30"
        },
        {
          id: 7,
          dataset: "乳腺癌细胞核数据集.csv",
          algo: "朴素贝叶斯算法",
          time: "2019-10-30 下午06：28"
        }
      ],
      colTitle:['预测结果','指标1'],//标题
      data: {
        percent: 0,
        data_yuce: [{预测结果:'1',指标1:'30'}]
      }
    };
  },
  methods: {
    check(params) {
      console.log(params);
      this.data_table = true;
    }
  }
};
</script>
<style scoped>
* {
  font-size: 16px;
}
.upload {
}
.algoControl {
  /* float: left; */
  font-family: "微软雅黑";
  font-size: 16px;
  display: inline-block;
  margin-left: 30px;
  margin-top: 50px;
}
.el-table--border::after,
.el-table--group::after,
.el-table::before {
  background-color: #009fcc;
}
.datasetTable {
  /* display: inline; */
}
</style>
