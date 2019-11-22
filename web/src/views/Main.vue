<template>
<div id="main">
  <dataset @finish="finish" />
  <prediction ref="predictions" @change="receiveMessage" />
  <div style="clear:both"></div>
  <chart ref="chart" :message="message" />
  </div>
</template>

<script>
import dataset from '../views/Dataset.vue'
import chart from '../views/Chart'
import prediction from '../views/Prediction'
export default {
 components:{dataset,chart,prediction},
 data(){
     return {
         message:[]
     }
 },
 methods:{
     finish(){
         this.$refs.predictions.GetPredictionsBytrainId()
     },
     receiveMessage(params){//把chart信息传到chart表中
         this.$refs.chart.init(params)
     }
 },
}
</script>

<style scoped>

</style>