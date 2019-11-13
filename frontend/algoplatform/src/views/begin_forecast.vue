<template>
  <div class="forecast">
    <el-steps :active="active" finish-status="success" align-center>
      <el-step title="步骤 1" icon="el-icon-edit" description="选择预测集"></el-step>
      <el-step title="步骤 2" icon="el-icon-edit" description="选择训练集"></el-step>
      <el-step title="步骤 3" icon="el-icon-edit" description="选择本次分析算法"></el-step>
      <el-step title="分析" icon="el-icon-upload" description="结果请在“预测结果”中查看"></el-step>
    </el-steps>
    <center>
      <el-button type="primary" size="small" @click="next">继续</el-button>
      <el-button v-if="active ==1 || active ==2" size="small" @click="front">回退</el-button>
      <template>
        <transition name="el-zoom-in-top">
          <div v-if="active ==0" class="group">
            <span>预测集：</span>
            <el-select v-model="this_forecast.test" placeholder="请选择">
              <el-option
                v-for="item in yuce_options"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              ></el-option>
            </el-select>
          </div>
        </transition>
        <transition name="el-zoom-in-top">
          <div v-if="active ==1" class="group">
            <span>训练集：</span>
            <el-select v-model="this_forecast.train" placeholder="请选择">
              <el-option
                v-for="item in xunlian_options"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              ></el-option>
            </el-select>
          </div>
        </transition>
        <transition name="el-zoom-in-top">
          <div v-if="active ==2" class="group">
            <span>算法：</span>
            <el-select v-model="this_forecast.algorithm" placeholder="请选择">
              <el-option
                v-for="item in algo_options"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              ></el-option>
            </el-select>
          </div>
        </transition>
      </template>
    </center>
  </div>
</template>
<script>
export default {
  data() {
    return {
      active: 0,
      index: [0, 1, 2, 3],
      yuce_options: [],
      xunlian_options: [],
      algo_options: [],
      this_forecast: {
        algorithm:'',
        train:'',
        test:'' 
      }
    };
  },
  created(){
    this.init()
  },
  methods: {
    init(){
      this.$axios.get(this.$axios.defaults.baseURL+'/dataset')
      .then(response=>{
        // console.log(response)
        var len=response.data.data.datasets.length;
        var item = {};
        this.xunlian_options = [];
        this.yuce_options = [];
        for(var i=0;i<len;i++)
        {
          item = response.data.data.datasets[i];
          if(item.type ==="train")
            this.xunlian_options.push(item)
          else
            this.yuce_options.push(item)
        }
      });
      this.$axios.get(this.$axios.defaults.baseURL+'/algorithm')
      .then(response=>{
        // console.log(response)
        var data={};
        var count = response.data.data.algorithms.length;
        for(var i=0;i<count;i++)
        {
          var item = response.data.data.algorithms[i];
          this.algo_options.push(item)
        }
      })
    },
    front() {
      this.active--;
    },
    next() {
      if (this.active == 2) this.active++;
      if (this.active++ > 3) {
        this.beginForecast();
      }
    },
    beginForecast() {
      console.log(this.this_forecast)
      // this.$axios.get(this.$axios.defaults.baseURL+'/prediction')
      // .then(response=>{
      //   console.log(response)
      // })
      this.$axios.post(this.$axios.defaults.baseURL+'/prediction',JSON.stringify(this.this_forecast))
      .then(response=>{
        console.log("prediction post")
        console.log(response)
        this.$emit('change')
      // this.$router.push('/forecast/check')
    })
    }
  }
};
</script>
<style type="text/css">
.forecast {
  margin-top: 100px;
  margin-left: 100px;
}
.group {
  margin: 50px;
}
</style>