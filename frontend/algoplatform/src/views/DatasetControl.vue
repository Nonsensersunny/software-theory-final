<template>
  <div class="datasetControl">
    <!-- <img alt="Vue logo" src="../assets/logo.png"> -->
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <center>
      <div class="upload">
        <el-button size="small" type="primary" @click="dialogVisible = true">上传数据集</el-button>
      </div>
      <div class="datasetTable">
        <el-table
          :data="tableData"
          border
          style="font-size:18px;border-color:#009FCC;width: 90%;"
          :row-style="{height:'80px'}"
        >
          <el-table-column type="index" align="center" label="序号" ></el-table-column>
          <el-table-column prop="name" label="数据集名称" align="center" ></el-table-column>
          <el-table-column prop="type" label="类型" align="center" ></el-table-column>
          <el-table-column prop="time" label="上传时间" align="center"></el-table-column>
          <el-table-column prop="work" label="操作" align="center" >
            <template slot-scope="scope">
              <el-button style="color:#409EFF" @click="check(scope.row)">查看</el-button>
              <el-button disabled style="color:#409EFF" @click="dele(scope.row)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>

      <el-dialog title="上传数据集" :visible.sync="dialogVisible" width="600px">
        <el-form style="width:590px;" :model='form' size="small">
          <el-form-item label="文件" :label-width="formLabelWidth">
            <el-input type="file" id="uploadFile" name="uploadFile" v-model="form.path" style="width:80%;margin-bottom:10px"></el-input>
          </el-form-item>
            <el-form-item label="名称" :label-width="formLabelWidth">
              <el-input v-model="form.name" style="width:80%;margin-bottom:10px"></el-input>
            </el-form-item>
            <!-- <el-form-item label="路径" :label-width="formLabelWidth">
              <el-input v-model="form.filePath" disabled style="width:80%;margin-bottom:10px"></el-input>
            </el-form-item> -->
            <el-form-item label="类型" :label-width="formLabelWidth">
              <el-select
                v-model="form.type"
                placeholder="请选择"
                style="width:80%;margin-bottom:10px"
              >
                <el-option
                  v-for="item in options"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                ></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="描述" :label-width="formLabelWidth">
              <el-input
                type="textarea"
                v-model="form.description"
                :rows="3"
                style="width:80%;margin-bottom:10px"
                placeholder="请输入对数据集的描述"
              ></el-input>
            </el-form-item>
           <el-form-item>
              <el-button @click="dialogVisible = false">取 消</el-button>
            <el-button type="primary" @click="uploadDataset">确 定</el-button>
           </el-form-item>
        </el-form>
       
      </el-dialog>
      <el-dialog :title="show_dataset.name" :visible.sync="data_table" width="600px">
        <el-form :ref="show_dataset">
          <el-form-item label="名称" :label-width="formLabelWidth">
            <el-input v-model="show_dataset.name" disabled style="width:80%;margin-bottom:10px"></el-input>
          </el-form-item>
          <el-form-item label="路径" :label-width="formLabelWidth">
            <el-input v-model="show_dataset.path" disabled style="width:80%;margin-bottom:10px"></el-input>
          </el-form-item>
          <el-form-item label="类型" :label-width="formLabelWidth">
            <el-input v-model="show_dataset.type" disabled style="width:80%;margin-bottom:10px"></el-input>
          </el-form-item>
          <el-form-item label="上传时间" :label-width="formLabelWidth">
            <el-input v-model="show_dataset.time" disabled style="width:80%;margin-bottom:10px"></el-input>
          </el-form-item>
          <el-form-item label="描述" :label-width="formLabelWidth">
            <el-input
              type="textarea"
              v-model="show_dataset.description"
              disabled
              :rows="3"
              style="width:80%;margin-bottom:10px"
            ></el-input>
          </el-form-item>
          <!-- </el-form-item> -->
        </el-form>
      </el-dialog>
    </center>
  </div>
</template>

<script>
// import api from '@/axios/api'
export default {
  // name: "DatasetControl",
  props: {},
  data() {
    return {
      fileList: [],
      dialogVisible: false,
      data_table: false,
      formLabelWidth: "90px",
      form: {
        name: "", //数据集名字
        path: "",//路径
        type: "", //训练集还是预测集
        description: "", //描述
        time: "" //上传时间
      },
      show_dataset: {
        name: "", //数据集名字
        path: "",
        type: "", //训练集还是预测集
        description: "", //描述
        time: "" //上传时间
      },
      options: [
        { value: 'test', label: "预测集" },
        { value: "train", label: "训练集" }
      ],
      tableData: [],
      fileObj:{},
    };
  },
  created() {
    this.init()
  },
  methods: {
    init(){

      this.$axios.get(this.$axios.defaults.baseURL+'/dataset')
      .then(response=>{
        console.log("dateset get")
        console.log(response)
        var len=response.data.data.datasets.length;
        this.tableData = [];
        var item = {};
        for(var i=0;i<len;i++)
        {
          item = response.data.data.datasets[i];
          if(item.type ==="train")
            item.type = "训练集"
          else
            item.type ="预测集"
          this.tableData.push(item)
        }
      })
    },
    check(params) {
      this.show_dataset.name = params.name;
      this.show_dataset.path = params.path;
      this.show_dataset.description = params.description;
      this.show_dataset.type = params.type;
      this.show_dataset.time = params.time;
      this.data_table = true;
      
    },
    dele(params) {
      var name = params.name;
      var type = params.type;
      let betsData = this.tableData;
      betsData.forEach(function(item, idx) {
        if (item.name == params.name) betsData.splice(idx, 1);
      });
      this.$message.success(params.type + "  " + params.name + "  删除成功!");
    },
    
    uploadDataset() {
      //提交表单，上传到数据库里
      var myDate = new Date();
      this.form.time = myDate.toLocaleString();
      this.fileObj = document.getElementById("uploadFile").files[0];
      // console.log(this.fileObj)
      //请求后端
      var formdata = new FormData();
      formdata.append('name',this.form.name)
      formdata.append('upload',this.fileObj)
      formdata.append('description',this.form.description)
      formdata.append('time',this.form.time)
      formdata.append('name',this.form.name)
      formdata.append('type',this.form.type)
      if(this.form.name != ''&&this.form.path !='' && this.form.description != ''&&this.form.type != '' && this.form.time != '')
      {
        this.$axios.put(this.$axios.defaults.baseURL+'/dataset',formdata)
        .then(response=>{
          console.log("dataset put")
          console.log(response)
          if (this.form.type === "test")
            this.$message.success("预测集 " + this.form.name + " 上传成功!");
          else if (this.form.type === "train")
            this.$message.success("训练集 " + this.form.name + " 上传成功!");
          this.dialogVisible = false;
          this.tableData.push(this.form);
          });
      }
      else{
        this.$message.error("请填写完整信息！")
      }
    }
  }
  // components: {
  //   HelloWorld
  // }
};
</script>
<style scoped>
* {
  font-size: 18px;
}
.upload {
  margin: 50px;
}

.el-table--border::after,
.el-table--group::after,
.el-table::before {
  background-color: #009fcc;
}

.datasetControl {
  /* float: left; */
  font-family: "微软雅黑";
  font-size: 18px;
  display: inline-block;
}

.datasetTable {
  width: 1050px;
  /* display: inline; */
  margin: 10px;
}
</style>
