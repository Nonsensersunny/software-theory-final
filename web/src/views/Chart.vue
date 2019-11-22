<template>
  <div id="chart" :style="{width:'100%',height:'500px'}"></div>
</template>
<script>
export default {
  props: {
    message: {
      type: Array,
      default: []
    }
  },
  data() {
    return {
      series: [],
      // dataset:[],
      // algorithm:[],
      xAxis: []
    };
  },
  mounted() {
    // this.init(message)
  },
  methods: {
    changeForm(response) {
      //只要name一样就放在一个系列里，根据{'算法名':0.955} 这个格式存进series的data里
      //并把数据集name存进series的name里
      //对x轴的类别 记得改一下，要统一。

      var count = response.length;
      this.series = [];
      for (var i = 0; i < count; i++) {
        var data = response[i]; //第i条预测数据
        if (this.xAxis.length !== 0) {
          var k = 0;
          for (k = 0; k < this.xAxis.length; k++) {
            if (this.xAxis[k] === data.algoname) break;
          }
          if (k >= this.xAxis.length) {
            this.xAxis.push(data.algoname);
          }
        } else {
          this.xAxis.push(data.algoname);
        }
      }
      for (var j = 0; j < count; j++) {
        if (this.series.length === 0) {
          // console.log(response[j])
          var item = {
            // id:'',
            name: "",
            type: "scatter",
            symbolSize: function(data) {
              return (data[1] - 0.5)*100;
            },
            data: []
          };
          // item.name = " "
          item.name = response[j].name.toString();
          // item.data = []
          item.data.push([response[j].algoname, response[j].accuracy]);
          this.series.push(item);
          // console.log(this.series)
        } else {
          for (var m = 0; m < this.series.length; m++) {
            if (response[j].name === this.series[m].name) {
              this.series[m].data.push([
                response[j].algoname,
                response[j].accuracy
              ]);
              break;
            }
            if (this.series.length - 1 === m) {
              var item = {
                name: "",
                type: "scatter",
                symbolSize: function(data) {
                  return (data[1] - 0.5)*100;
                },
                data: []
              };
              console.log(response[j].name);
              item.name = response[j].name.toString();
              item.data.push([response[j].algoname, response[j].accuracy]);
              this.series.push(item);
            }
          }
        }
      }
      // console.log(this.series)
    },
 
    init(params) {
      for (var ii = 0; ii < params.length; ii++) {
        if (params[ii].accuracy > 10) params[ii].accuracy /= 100;
        params[ii].accuracy = params[ii].accuracy.toFixed(4);
      }
      this.changeForm(params);
      var myChart = this.$echarts.init(document.getElementById("chart"));
      myChart.setOption({
        //横坐标是方法 纵坐标是准确率 要有多个类别（样本类别）
        axisPointer: {
          show: true,
          type: "line",
          lineStyle: {
            type: "dashed"
          }
        },
        legend: {
          show: true,
          top: "45px"
        },
        title: {
          show: true,
          text: "不同数据集在不同算法下的准确率"
        },
        grid: {
          top: "100px"
        },
        dataZoom: [
          {
            type: "inside",
            xAxisIndex: [0],
            start: 0,
            end: 100
          },
          {
            type: "inside",
            yAxisIndex: [0],
            start: 71,
            end: 100
          }
        ],
        tooltip: {
          show: true,
          trigger: "item",
          triggerOn: "mousemove"
        },
        xAxis: {
          type: "category",
          data: this.xAxis,
          axisLine: {
            lineStyle: {
              color: "#61a0a8"
            }
          },
          axisTick: {
            inside: true,
            lineStyle: {
              color: "#61a0a8"
            }
          }
        },
        yAxis: {
          type: "value",
          min: "dataMin",
          axisLine: {
            symbol: ["none", "arrow"],
            lineStyle: {
              color: "#61a0a8"
            }
          },
          axisTick: {
            show: false
          }
        },
        series: this.series,
        //    series: [
        //    {
        //         name: '乳腺癌细胞核',
        //         type: 'scatter',
        //         // itemStyle: itemStyle,
        //         // data: [[0.94,0.91,0.86,0.74],[0.84,0.75,0.94,0.92]]
        //         data:[['SVM',0.96],['LOGISTIC',0.75],['DESIDETREE',0.88],['BAYES',0.89]],
        //         symbolSize: function (data) {
        //             return (data[1]-0.5)*100;
        //         }
        //     },{
        //     	name:'心脏病',
        //     	type:'scatter',
        //     	data:[['SVM',0.94],['LOGISTIC',0.85],['DESIDETREE',0.87],['BAYES',0.84]],
        //         symbolSize: function (data) {
        //             return (data[1]-0.5)*100;
        //         }
        //     }
        //     ],
        color: [
          "#c23531",
          "#61a0a8",
          "#d48265",
          "#91c7ae",
          "#749f83",
          "#ca8622",
          "#bda29a",
          "#6e7074",
          "#546570",
          "#c4ccd3"
        ]
      });
    }
    // 	resize() {
    //       this.__resizeHanlder = debounce(() => {
    //           if (this.chart) {
    //               this.chart.resize()
    //           }
    //       }, 100);
    //       this.__resizeHanlder();
    //   },
  }
};
</script>
<style scoped>
#chart {
  width: 70%;
  margin: 15px auto;
}
</style>