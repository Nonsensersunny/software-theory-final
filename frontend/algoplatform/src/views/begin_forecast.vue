<template>
	<div class="forecast">
		<el-steps :active="active" finish-status="success" align-center>
		  <el-step title="步骤 1" icon="el-icon-edit" description="选择预测集"></el-step>
		  <el-step title="步骤 2" icon="el-icon-edit" description="选择训练集"></el-step>
		  <el-step title="步骤 3" icon="el-icon-edit" description="选择本次分析算法"></el-step>
		  <el-step title="分析" icon="el-icon-upload" description="结果请在“预测结果”中查看"></el-step>
		</el-steps>
		<center><el-button type="primary" style="" @click="next">继续</el-button>
		<el-button  v-if="active ==1 || active ==2"  @click="front">回退</el-button>
		<template >
			<transition name="el-zoom-in-top">
			<div v-if="active ==0" class="group">
				<span>预测集：</span>
				<el-select v-model="forecast_dataset" placeholder="请选择">
					<el-option
					v-for="item in options"
				      :key="item.value"
				      :label="item.label"
				      :value="item.value">
				      	
				      </el-option>
				</el-select>
			</div>
		</transition>
		<transition name="el-zoom-in-top">
			<div v-if="active ==1" class="group"><span>训练集：</span>
				<el-select v-model="xunlian_dataset" placeholder="请选择">
					<el-option
					v-for="item in xunlian_options"
				      :key="item.value"
				      :label="item.label"
				      :value="item.value">
				      	
				      </el-option>
				</el-select>
			</div>
			</transition>
			<transition name="el-zoom-in-top">
			<div v-if="active ==2" class="group"><span>算法：</span>
				<el-select v-model="algo" placeholder="请选择">
					<el-option
					v-for="item in algo_options"
				      :key="item.value"
				      :label="item.label"
				      :value="item.value">
				      	
				      </el-option>
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
        index:[0,1,2,3],
        forecast_dataset:'',
        xunlian_dataset:'',
        algo:'',
        options:[{label:'1号病人细胞核特征.csv',value:'1号病人细胞核特征.csv'},
        		{label:'2号病人细胞核特征.csv',value:'2号病人细胞核特征.csv'},
        		{label:'3号病人细胞核特征.csv',value:'3号病人细胞核特征.csv'}],
        xunlian_options:[{label:'乳腺癌细胞核特征训练集.csv',value:'乳腺癌细胞核特征训练集.csv'},
        		{label:'糖尿病血液指标训练集.csv',value:'糖尿病血液指标训练集.csv'},
        		{label:'心脏病心脏指标训练集.csv',value:'心脏病心脏指标训练集.csv'}],
        algo_options:[{label:'朴素贝叶斯算法',value:'朴素贝叶斯算法'},
        		{label:'支持向量机SVM算法',value:'支持向量机SVM算法'},
        		{label:'Logistic逻辑回归',value:'Logistic逻辑回归'}],
      };
    },

    methods: {
    	front(){
    		this.active--;
    	},
      next() {
      	
      	if(this.active == 2) this.active ++;
        if (this.active++ > 3) {
        	this.beginForecast();
        }
      },
      beginForecast(){
      	this.$emit('change','check_forecast')
      	// this.$router.push('/forecast/check')
      }
    }
  }
</script>
<style type="text/css">
	.forecast{
		margin-top: 100px ;
		margin-left: 100px;
	}
	.group{
		margin:50px ;
	}
</style>