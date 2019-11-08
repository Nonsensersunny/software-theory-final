<template>
	<div id="chart" :style="{width:'100%',height:'500px'}">
		<!-- <chart :options="myOption" ></chart> -->
	</div>
</template>
<script type="text/javascript">
export default{
	data(){
		return{
			
		}
	},
	mounted(){
		this.init()
	},
	methods:{
		init(){
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
		       series: [{
		            name: '乳腺癌细胞核',
		            type: 'scatter',
		            // itemStyle: itemStyle,
		            // data: [[0.94,0.91,0.86,0.74],[0.84,0.75,0.94,0.92]]
		            data:[['SVM',0.96],['LOGISTIC',0.75],['DESIDETREE',0.88],['BAYES',0.89]],
		            symbolSize: function (data) {
			            return (data[1]-0.6)*100;
			        }
		        },{
		        	name:'心脏病',
		        	type:'scatter',
		        	data:[['SVM',0.94],['LOGISTIC',0.85],['DESIDETREE',0.87],['BAYES',0.84]],
		            symbolSize: function (data) {
			            return (data[1]-0.5)*100;
			        }
		        }],
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