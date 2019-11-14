<template>
	<div id="chart" :style="{width:'100%',height:'500px'}">
		<!-- <chart :options="myOption" ></chart> -->
	</div>
</template>
<script type="text/javascript">
export default{
	data(){
		return{
			series:[],
			dataset:[],
			algorithm:[]
		}
	},
	mounted(){
		this.init()
	},
	methods:{
		changeForm(response){
			//只要name一样就放在一个系列里，根据{'算法名':0.955} 这个格式存进series的data里
			//并把数据集name存进series的name里
			//对x轴的类别 记得改一下，要统一。
				var item = {
					id:'',
					name:'',
					type:'scatter',
					symbolSize: function (data) {
			            return (data[1]-0.5)*100;
			        },
			        data:[]
				}
				var count = response.data.data.prediction.length;
				for(var i=0;i<count;i++)
				{
					item.name = ""
					item.data = []
					var data = response.data.data.prediction[i]//第i条预测数据
					this.changeData(data)
					//通过name去找所在的series//并把算法名algorithm_name和准确率放到data里
					if(this.series.length != 0)
					{
						for(var j=0;j<this.series.length;j++)
						{
							if(data.name == this.series[j].name)
							{//已经存在该数据集的系列，只要放入series的data就行
								this.series[j].data.push([data.algorithm_name,data.accuracy])
								break; 
							}
							if(j == this.series.length-1)
							{//说明该数据集不存在，那么新建一个系列
								item.name = data.name;
								item.data.push([data.algorithm_name,data.accuracy])
								this.series.push(item)
							}
						}
					}
					else{//第一次添加
						item.name = data.name;
						item.data.push([data.algorithm_name,data.accuracy])
						this.series.push(item)
					}
				}

		},
		changeData(params)//把数据集算法的name属性加到预测结果里。
		{
			for(var i=0;i<this.dataset.length;i++)
			{
				if(params.test == this.dataset[i].id)
					params.name = this.dataset[i].name
				if(params.algorithm == this.algorithm[i].name)
					params.algorithm_name = this.algorithm[i].name
			}
		},
		init(){
			this.$axios.get(this.$axios.defaults.baseURL+'/algorithm')
			.then(response=>{
				var len=response.data.data.algorithms.length;
		        for(var i=0;i<len;i++)
		        {
		          this.algorithm.push(response.data.data.algorithms[i])
		        }
			});
			this.$axios.get(this.$axios.defaults.baseURL+'/dataset')
			.then(response=>{
				var len=response.data.data.datasets.length;
		        for(var i=0;i<len;i++)
		        {
		          this.dataset.push(response.data.data.datasets[i])
		        }
			});
			this.$axios.get(this.$axios.defaults.baseURL+'/prediction')
			.then(response=>{
				console.log(response)
				this.changeForm(response);
			})	;	
			
			var myChart = this.$echarts.init(document.getElementById('chart'));
			myChart.setOption({//横坐标是方法 纵坐标是准确率 要有多个类别（样本类别）
		      axisPointer:{
		      	show:true,
		      	type:'line',
		      	lineStyle:{
		      		type:'dashed'
		      	}
		      },
		      legend:{
		      	show:true,
		      	top:'45px',
		      },
		      title:{
		      	show:true,
		      	text:'不同数据集在不同算法下的准确率'
		      },
		      grid:{
		      	top:'100px'
		      },
		      dataZoom: [
                       {
                            type: 'inside',
                            xAxisIndex: [0],
                            start:0 ,
                            end: 100
                        },
                        {
                          type:'inside',
                          yAxisIndex:[0],
                          start:71,
                          end:100
                        }
              ],
              tooltip:{
              	show:true,
              	trigger:'item',
              	triggerOn:'mousemove'
              },
		      xAxis: {
			       type: 'category',
			       data: ['SVM','LOGISTIC','DESIDETREE','BAYES'],
			       axisLine:{
			       	lineStyle:{
			       		color:'#61a0a8'
			       	}
			       },
			       axisTick:{
			       	inside:true,
			       	lineStyle:{
			       		color:'#61a0a8'
			       	}
			       }
		       },
		       yAxis: {
		           type: 'value',
		           min:'dataMin',
		           axisLine:{
		           	symbol:['none','arrow'],
		           	lineStyle:{
			       		color:'#61a0a8'
			       	}
		           },
		           axisTick:{
		           	show:false
		           }
		       },
		       series:this.series,
		       // series: [
		       // {
		       //      name: '乳腺癌细胞核',
		       //      type: 'scatter',
		       //      // itemStyle: itemStyle,
		       //      // data: [[0.94,0.91,0.86,0.74],[0.84,0.75,0.94,0.92]]
		       //      data:[['SVM',0.96],['LOGISTIC',0.75],['DESIDETREE',0.88],['BAYES',0.89]],
		       //      symbolSize: function (data) {
			      //       return (data[1]-0.5)*100;
			      //   }
		       //  },{
		       //  	name:'心脏病',
		       //  	type:'scatter',
		       //  	data:[['SVM',0.94],['LOGISTIC',0.85],['DESIDETREE',0.87],['BAYES',0.84]],
		       //      symbolSize: function (data) {
			      //       return (data[1]-0.5)*100;
			      //   }
		       //  }
		        // ],
		        color:['#c23531', '#61a0a8', '#d48265', '#91c7ae','#749f83',  '#ca8622', '#bda29a','#6e7074', '#546570', '#c4ccd3']
		     });
		},
		resize() {
          this.__resizeHanlder = debounce(() => {
              if (this.chart) {
                  this.chart.resize()
              }
          }, 100);
          this.__resizeHanlder();
      },
	}
}
</script>
<style type="text/css">
#chart{
	margin-left: 30px;
	margin-top: 30px;
}
</style>