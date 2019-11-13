<template>
  <div class="algoControl">
  
    <center>
      <div class="upload">
        <!-- <el-button size="small" type="primary" @click="dialogVisible = true">上传算法</el-button> -->
        <el-button size="small" type="primary" @click="NotAllow">上传算法</el-button>
      </div>
      <div class="datasetTable">
        <el-table
          :data="tableData"
          border
          style="font-size:18px;border-color:#009FCC"
          :row-style="{height:'80px'}"
        >
          <el-table-column type="index" align="center" label="序号" ></el-table-column>
          <el-table-column prop="name" label="算法名称" align="center" ></el-table-column>
          <el-table-column prop="description" label="描述" align="center"></el-table-column>
          <el-table-column prop="time" label="上传时间" align="center" ></el-table-column>
       
        </el-table>
      </div>

      <el-dialog title="上传数据集" :visible.sync="dialogVisible" width="600px">
        <el-form style="width:590px;" size="small">
          <el-form-item label="上传文件"
              :label-width="formLabelWidth">
             <el-input type="file" id="uploadFile" name="uploadFile" v-model="form.path" style="width:80%;margin-bottom:10px"></el-input>
          </el-form-item>
            <el-form-item
              label="算法名称"
              :label-width="formLabelWidth" >
              <el-input v-model="form.name" style="width:80%;margin-bottom:10px"></el-input>
            </el-form-item>
            
            <el-form-item
              label="算法描述"
              :label-width="formLabelWidth">
              <el-input
                type="textarea"
                v-model="form.description"
                :rows="3" style="width:80%;margin-bottom:10px"
                placeholder="请输入对算法的描述，如算法的特点、适用情况，"
              ></el-input>
            </el-form-item>
          <el-form-item>
            <el-button @click="dialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="uploadDataset">确 定</el-button>
          </el-form-item>
        </el-form>
      
      </el-dialog>
     
    </center>
  </div>
</template>

<script>
export default {
  name: "datasetControl",
  data() {
    return {
      fileList: [],
      dialogVisible: false,
      data_table:false,
      formLabelWidth: "90px",
      form: {
        name: "", //数据集名字
        description: "", //描述
        time: '',//上传时间
        path:'', 
      },
      fileObj:{},
    
      tableData: []
    };
  },
  created(){
    this.init();
  },
  methods: {
    init(params) {//这里能获取所有算法信息
      this.$axios.get(this.$axios.defaults.baseURL+'/algorithm')
      .then(response=>{
        console.log(response)
        var data={};
        this.tableData = []
        var count = response.data.data.algorithms.length;
        for(var i=0;i<count;i++)
        {
          var item = response.data.data.algorithms[i];
          data = {};
          data.name = item.name;
          data.description = item.description;
          var datetime = new Date('2019/11/5');
          // console.log(datetime)
          data.time = datetime.toLocaleString();
          // console.log(data.time)
          this.tableData.push(data)
        }
      })
    },
    NotAllow(){
      this.$message.error("抱歉，您没有权限上传算法文件！")
    },
    dele(params) {
      var name = params.name;
      let betsData = this.tableData;
      betsData.forEach(function(item, idx) {
        if (item.name == params.name) betsData.splice(idx, 1);
      });
      this.$message.success("算法文件  " + params.name + "  删除成功!");
    },
    uploadDataset(){
      var myDate = new Date();
      this.form.time = myDate.toLocaleString();
      this.fileObj = document.getElementById("uploadFile").files[0];
      console.log(this.fileObj)
      var item={};
      item.upload = this.fileObj;
      item.name = this.form.name;
      item.description = this.form.description;
      if(this.form.name !='' && this.form.description !='' && this.form.path != '')
      {
          this.$axios.put(this.$axios.defaults.baseURL+'/algorithm',JSON.stringify(item))
        .then(response=>{
          console.log(response)
          if(response.data.data.status == 'success')
          {
            this.tableData.push(this.form)
            this.$message.success("算法文件 " + this.form.name + " 上传成功!");
          }
          else
          {
            this.$message.error("上传失败" +response.data.error);
          }
          // this,tableData.push(response.data)
        })
        this.dialogVisible = false;
      }
      else
      {
        this.$message.error('请填写完整信息！')
      }
  }
  }
};
</script>
<style scoped>
* {
  font-size: 18px;
}
.upload {
  margin: 50px auto;
}
.el-table--border::after,
.el-table--group::after,
.el-table::before {
  background-color: #009fcc;
}

.algoControl {
  /* float: left; */
  font-family: "微软雅黑";
  font-size: 18px;
  display: inline-block;
  
}
.datasetTable::-webkit-scrollbar{
  display: none;
}

.datasetTable {
  /* display: inline; */
  width: 1050px;
  margin-left: 10px;
  max-height: 650px;
  overflow-y: scroll;
}
</style>
