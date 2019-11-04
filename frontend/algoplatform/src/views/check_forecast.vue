<template>
  <div class="algoControl">
   <div class="datasetTable" >
      <el-table :data="tableData" border style="font-size:18px;" :row-style="{height:'80px'}"  >
        <el-table-column  prop="id" align="center"  label="序号" width="100px"></el-table-column>
        <el-table-column prop="dataset" label="算法名称"  align="center"  width="320px"></el-table-column>
        <el-table-column prop="algo" label="算法名称"  align="center"  width="320px"></el-table-column>
        <el-table-column prop="time" label="预测时间"  align="center"  width="280px"></el-table-column>
        <el-table-column prop="work" label="操作"  align="center"  width="280px">
          <template slot-scope="scope">
            <el-button style="color:#409EFF"  @click="check(scope.row)" >查看预测结果</el-button>
          </template>
        </el-table-column>
        
      </el-table>
    </div>   
  </div>
</template>

<script>
export default {
  name: "datasetControl",
  data() {
    return {
      tableData: [
        {
            id:5,
          dataset:'乳腺癌细胞核数据集.csv',
          algo:"SVM向量机算法",
          time: "2019-10-30 下午06：30"
        },{
            id:7,
          dataset:'乳腺癌细胞核数据集.csv',
          algo:"朴素贝叶斯算法",
          time: "2019-10-30 下午06：28"
        }
      ]
    };
  },
  methods: {
    check(params){
        console.log(params)
    },
   
  }
};
</script>
<style scoped>
*{
font-size: 18px;
}
.upload {
  margin: 50px;
}
.algoControl {
  /* float: left; */
  font-family:'微软雅黑';
  font-size: 18px;
  display: inline-block;
}

.datasetTable {
  /* display: inline; */
  margin-left: 50px;
}


</style>
