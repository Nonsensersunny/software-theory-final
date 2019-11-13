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
          <el-table-column type="index" align="center" label="序号" width="50px" ></el-table-column>
          <el-table-column prop="test" label="算法名称" align="center"></el-table-column>
          <el-table-column prop="algorithm" label="算法名称" align="center" ></el-table-column>
          <el-table-column prop="time" label="预测时间" align="center"></el-table-column>
          <el-table-column prop="work" label="操作" align="center" >
            <template slot-scope="scope">
              <el-button style="color:#409EFF" @click="check(scope.row)">查看结果</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <el-dialog title="预测结果" :visible.sync="data_table" width="700px"> 
        <div>
          <span>本次预测准确率：</span>
          <el-input disabled v-model="show_prediction.accuracy" style="width:50%;margin-bottom:10px;font-color:red;"></el-input>
        </div>
        <el-table :data="show_prediction.output" height="400px">
          <el-table-column
            align="center"
            v-for="(item, index) in colTitle"
            :key="index"
            :label="item"
            :prop="item"
            :height="10"
            :width="100"
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
      tableData: [],
      show_prediction:{},//
      colTitle:['预测结果','指标1'],//标题
      data: {//传进来的数据
        id:5,
        test: "乳腺癌细胞核数据集.csv",
        algorithm: "朴素贝叶斯算法",
        time: "2019-10-30 下午06：28",
        accuracy: 90.42553191489363,
        output:[[14.05, 27.15, 91.38, 600.4, 0.09929, 0.1126, 0.04462, 0.04304, 0.1537, 0.06171, 0.3645, 1.492, 2.888, 29.84, 0.007256, 0.02678, 0.02071, 0.01626, 0.0208, 0.005304, 15.3, 33.17, 100.2, 706.7, 0.1241, 0.2264, 0.1326, 0.1048, 0.225, 0.08321, 0.0], [11.2, 29.37, 70.67, 386.0, 0.07449, 0.03558, 0.0, 0.0, 0.106, 0.05502, 0.3141, 3.896, 2.041, 22.81, 0.007594, 0.008878, 0.0, 0.0, 0.01989, 0.001773, 11.92, 38.3, 75.19, 439.6, 0.09267, 0.05494, 0.0, 0.0, 0.1566, 0.05905, 0.0], [15.22, 30.62, 103.4, 716.9, 0.1048, 0.2087, 0.255, 0.09429, 0.2128, 0.07152, 0.2602, 1.205, 2.362, 22.65, 0.004625, 0.04844, 0.07359, 0.01608, 0.02137, 0.006142, 17.52, 42.79, 128.7, 915.0, 0.1417, 0.7917, 1.17, 0.2356, 0.4089, 0.1409, 1.0], [20.92, 25.09, 143.0, 1347.0, 0.1099, 0.2236, 0.3174, 0.1474, 0.2149, 0.06879, 0.9622, 1.026, 8.758, 118.8, 0.006399, 0.0431, 0.07845, 0.02624, 0.02057, 0.006213, 24.29, 29.41, 179.1, 1819.0, 0.1407, 0.4186, 0.6599, 0.2542, 0.2929, 0.09873, 1.0], [21.56, 22.39, 142.0, 1479.0, 0.111, 0.1159, 0.2439, 0.1389, 0.1726, 0.05623, 1.176, 1.256, 7.673, 158.7, 0.0103, 0.02891, 0.05198, 0.02454, 0.01114, 0.004239, 25.45, 26.4, 166.1, 2027.0, 0.141, 0.2113, 0.4107, 0.2216, 0.206, 0.07115, 1.0], [20.13, 28.25, 131.2, 1261.0, 0.0978, 0.1034, 0.144, 0.09791, 0.1752, 0.05533, 0.7655, 2.463, 5.203, 99.04, 0.005769, 0.02423, 0.0395, 0.01678, 0.01898, 0.002498, 23.69, 38.25, 155.0, 1731.0, 0.1166, 0.1922, 0.3215, 0.1628, 0.2572, 0.06637, 1.0], [16.6, 28.08, 108.3, 858.1, 0.08455, 0.1023, 0.09251, 0.05302, 0.159, 0.05648, 0.4564, 1.075, 3.425, 48.55, 0.005903, 0.03731, 0.0473, 0.01557, 0.01318, 0.003892, 18.98, 34.12, 126.7, 1124.0, 0.1139, 0.3094, 0.3403, 0.1418, 0.2218, 0.0782, 1.0], [20.6, 29.33, 140.1, 1265.0, 0.1178, 0.277, 0.3514, 0.152, 0.2397, 0.07016, 0.726, 1.595, 5.772, 86.22, 0.006522, 0.06158, 0.07117, 0.01664, 0.02324, 0.006185, 25.74, 39.42, 184.6, 1821.0, 0.165, 0.8681, 0.9387, 0.265, 0.4087, 0.124, 1.0], [7.76, 24.54, 47.92, 181.0, 0.05263, 0.04362, 0.0, 0.0, 0.1587, 0.05884, 0.3857, 1.428, 2.548, 19.15, 0.007189, 0.00466, 0.0, 0.0, 0.02676, 0.002783, 9.456, 30.37, 59.16, 268.6, 0.08996, 0.06444, 0.0, 0.0, 0.2871, 0.07039, 0.0]],
      }
    };
  },
  created(){
    this.init()
  },
  methods: {
    init(){
      this.changeForm();
      // this.tableData.push(this.data)
      // this.$axios.get(this.$axios.defaults.baseURL+'/prediction')
      // .then(response=>{
      //   console.log(response)
      //   var count = response.data.data.prediction.length;
      //   var item = {};
      //   for(var i=0;i<count;i++)
      //   {//挨个对后端传过来的数据格式化
      //     //.......处理之后赋值给item this.changeForm()
      //     this.tableData.push(item)
      //   }

      // })
    },
    changeForm(){
      this.data.accuracy = this.data.accuracy.toFixed(4)+'%';//准确率格式化
      var data = this.data.output;
      var item = {}
      var i=0;
      this.colTitle = []
      this.data.output = []
      this.colTitle.push('预测结果')
      for(i =1; i<data[0].length;i++)//表头初始化
      {
        this.colTitle.push('指标 '+i)
      }
      for(i=0;i<data.length;i++)//对每个预测，进行数据格式转换
      {
        item = {}
        item[this.colTitle[0]] = data[i][data[i].length-1]//结果放在第一列
        for(var j=0;j<data[i].length-1;j++)
        { //将属性存储成对象的格式
          item[this.colTitle[j+1]] = data[i][j]
        }
        this.data.output.push(item)
      }
      this.tableData.push(this.data)
      console.log(this.colTitle)
      console.log(this.tableData)
    },
    check(params) {
      console.log(params);
      this.show_prediction.accuracy = params.accuracy;
      this.show_prediction.output = params.output;
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
  width: 1050px;
  /* display: inline; */
}
</style>
