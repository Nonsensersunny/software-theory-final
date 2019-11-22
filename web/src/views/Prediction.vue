<template>
  <div id="prediction">
    <div class="datasetTable">
      <div style="text-align: left;font-size: 18px;margin-bottom: 10px;">
        <span style="font-size: 18px;color: #009FCC">全部预测：</span>
      </div>
      <el-table
        :data="predictionList"
        border
        
        :row-style="{height:'80px'}"
      >
        <el-table-column type="index" align="center" label="序号" width="50px"></el-table-column>
        <el-table-column prop="name" label="预测集名称" align="center"></el-table-column>
        <el-table-column prop="algoname" label="算法名称" align="center"></el-table-column>
        <el-table-column prop="time" label="预测时间" align="center"></el-table-column>
        <el-table-column prop="work" label="操作" align="center">
          <template slot-scope="scope">
            <el-button size="small" type="primary" @click="check(scope.row.predictionid)">查看结果</el-button>
            <!-- <el-button size="small" type="danger" @click="del(scope.row.predictionid)">删除</el-button> -->
          </template>
        </el-table-column>
      </el-table>
    </div>
    <el-dialog title="预测结果" :visible.sync="show" width="700px">
      <div>
        <span>本次预测准确率：</span>
        <el-input
          disabled
          v-model="predictionResult.accuracy"
          style="width:50%;margin-bottom:10px;font-color:red;"
        ></el-input>
      </div>
      <el-table :data="predictionResult.result" height="400px">
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
  </div>
</template>
<script>
export default {
  data() {
    return {
      item: {},
      predictionList:[],
      name:'',
      show:false,
      predictionResult:{
          accuracy:0,
          result:[]
      },
      colTitle:[]
      
    };
  },
  methods: {
    async check(id){
        this.show = true
        let resP = await this.$store.dispatch("GetPredictionsById",id)
        this.predictionResult.accuracy = resP.data.data.prediction.accuracy.toFixed(4) 
        this.predictionResult.result =  JSON.parse(resP.data.data.prediction.result).output
        this.changeData()
    },
    changeData(){
        var datap = this.predictionResult.result;
        var item = {}
        this.colTitle.push("预测结果");
        for (var i = 1; i < datap[0].length;  i++) //表头初始化
        {
            this.colTitle.push("指标 " + i);
        }
        this.predictionResult.result = []
        for (var i = 0;  i < datap.length; i++ )//对每个预测，进行数据格式转换
        {
            item = {}
            item[this.colTitle[0]] = datap[i][datap[i].length - 1]; //结果放在第一列
            for (var j = 0; j < datap[i].length - 1; j++) {
            //将属性存储成对象的格式
                item[this.colTitle[j + 1]] = datap[i][j];
            }
            this.predictionResult.result.push(item);
        }
    },
    async getAlgorithm(aid) {
      let resA = await this.$store.dispatch("getAlgorithms");
      let algo = resA.data.algorithms;
      for (var i = 0; i < algo.length; i++) {
        if (algo[i].id === aid) {
          this.name = algo[i].name;
        //   console.log(this.name)
        }
      }
    //   return this.name;
    },
    async GetPredictionsBytrainId() {
      let resD = await this.$store.dispatch("getDatasets");
      let data = resD.data.datasets;
      this.predictionList = [];
      for (var i = 0; i < data.length; i++) {
        if (data[i].type === "train") {
          let res = await this.$store.dispatch(
            "GetPredictionsBytrainId",
            data[i].id
          );
          let prediction = res.data.data.prediction;
          if (prediction.length != 0) {
            for (var j = 0; j < prediction.length; j++) {
              this.item = {};
              this.item.name = data[i].name;
              this.item.predictionid = prediction[j].id;
              this.item.accuracy = prediction[j].accuracy;
              this.item.time = prediction[j].time;
              for (var k = 0; k < data.length; k++) {
                if (data[k].id === prediction[j].test_id)
                  this.item.testname = data[k].name;
              }
              await this.getAlgorithm(prediction[j].aid);
              this.item.algoname = this.name;
            //   console.log(this.item.algoname)
              this.predictionList.push(this.item);
            }
          }
        }
      }
      this.sentMessage();
    },
    sentMessage(){
        this.$emit("change",this.predictionList)
    }
  },
  created() {
    this.GetPredictionsBytrainId();
  }
};
</script>
<style scoped>
#prediction {
  float: left;
  height: 400px;
  width: 44%;
  border-left: darkcyan solid 1px;
  padding-left: 15px;
  overflow-y: scroll;
}
</style>