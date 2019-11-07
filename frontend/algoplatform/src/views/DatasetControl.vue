<template>
  <div class="datasetControl">
    <!-- <img alt="Vue logo" src="../assets/logo.png"> -->
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <center>
    <div class="upload">
      <el-button size="small" type="primary" @click="dialogVisible = true">上传数据集</el-button>
    </div>
    <div class="datasetTable" >
      <el-table :data="tableData" border style="font-size:18px;border-color:#009FCC" :row-style="{height:'80px'}"  >
        <el-table-column type="index"  align="center"  label="序号" width="100px"></el-table-column>
        <el-table-column prop="name" label="数据集名称"  align="center"  width="320px"></el-table-column>
        <el-table-column prop="type" label="类型"  align="center"  width="180px"></el-table-column>
        <el-table-column prop="uploadtime" label="上传时间"  align="center"  width="280px"></el-table-column>
        <el-table-column prop="work" label="操作"  align="center"  width="280px">
          <template slot-scope="scope">
            <el-button style="color:#409EFF"  @click="check(scope.row)" >查看</el-button>
            <el-button style="color:#409EFF"  @click="dele(scope.row)" >删除</el-button>
          </template>
        </el-table-column>
        
      </el-table>
    </div>

    <el-dialog title="上传数据集" :visible.sync="dialogVisible" width="600px">
      <div style="width:590px;">
      <!-- <span>选择文件</span> -->
      <el-input type="file"
      v-model="form.filePath"  style="width:85%;margin-bottom:10px"></el-input>
    </div>
      <el-form style="width:590px;">
              <el-form-item :model="form">
        <el-form-item label="名称" :label-width="formLabelWidth" >
          <el-input v-model="form.fileName"  style="width:80%;margin-bottom:10px"></el-input>
        </el-form-item>
        <el-form-item label="路径" :label-width="formLabelWidth" >
          <el-input v-model="form.filePath" disabled style="width:80%;margin-bottom:10px"></el-input>
        </el-form-item>
        <el-form-item label="类型" :label-width="formLabelWidth">
          <el-select v-model="form.fileType" placeholder="请选择" style="width:80%;margin-bottom:10px">
            <el-option
              v-for="item in options"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="描述" :label-width="formLabelWidth" >
          <el-input type="textarea" v-model="form.fileInfo" :rows="3" style="width:80%;margin-bottom:10px"
           placeholder="请输入对数据集的描述"></el-input>
        </el-form-item>
      </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="uploadDataset">确 定</el-button>
      </div>
    </el-dialog>
    <el-dialog :title="show_dataset.fileName" :visible.sync="data_table" width="600px">
      <el-form :ref="show_dataset">
        <el-form-item label="名称" :label-width="formLabelWidth" >
          <el-input v-model="show_dataset.fileName" disabled style="width:80%;margin-bottom:10px"></el-input>
        </el-form-item>
        <el-form-item label="路径" :label-width="formLabelWidth" >
          <el-input v-model="show_dataset.filePath" disabled style="width:80%;margin-bottom:10px"></el-input>
        </el-form-item>
        <el-form-item label="类型" :label-width="formLabelWidth">
          <el-input v-model="show_dataset.fileType" disabled style="width:80%;margin-bottom:10px"></el-input>
        </el-form-item>
        <el-form-item label="上传时间" :label-width="formLabelWidth">
          <el-input v-model="show_dataset.fileTime" disabled style="width:80%;margin-bottom:10px"></el-input>
        </el-form-item>
        <el-form-item label="描述" :label-width="formLabelWidth" >
          <el-input type="textarea" v-model="show_dataset.fileInfo" disabled :rows="3" style="width:80%;margin-bottom:10px"></el-input>
        </el-form-item>
      </el-form-item>
      </el-form>
    </el-dialog>
  </center>
  </div>
</template>

<script>
  // import api from '@/axios/api'
export default {
  // name: "DatasetControl",
  props:{


  },
  data() {
    return {
      
      fileList: [],
      dialogVisible: false,
      data_table:false,
      formLabelWidth: "90px",
      form: {
        fileName: "", //数据集名字
        filePath:"",
        fileType: "", //训练集还是预测集
        fileInfo: "", //描述
        fileTime: "", //上传时间
      },
      show_dataset:{
        fileName: "1", //数据集名字
        filePath:"",
        fileType: "", //训练集还是预测集
        fileInfo: "", //描述
        fileTime: "", //上传时间
      },
      options: [
        { value: "预测集", label: "预测集" },
        { value: "训练集", label: "训练集" }
      ],
      tableData: [
        {
          name: "乳腺癌细胞核特征测试集.csv",
          type: "测试集",
          path:'',
          uploadtime: "2019-10-30 下午06：30",
          info:''
        },{
          name: "乳腺癌细胞核特征测试集.csv",
          type: "训练集",
          path:'',
          uploadtime: "2019-10-30 下午06：28",
          info:''
        }
      ]
    };
  },
  created(){
    },
  methods: {
   
    check(params){
      this.show_dataset.fileName = params.name
      this.show_dataset.filePath = params.path
      this.show_dataset.fileInfo = params.info
      this.show_dataset.fileType = params.type
      this.show_dataset.fileTime = params.uploadtime
      this.data_table = true
    },
    dele(params){
      var name = params.name
      var type = params.type
      let betsData = this.tableData
      betsData.forEach(function(item,idx){
        if(item.name == params.name)
          betsData.splice(idx,1)
      })
      this.$message.success(params.type+"  " + params.name+"  删除成功!");

    },
    uploadSuccess(params) {
      //上传到前端，缓存 ?
      const formDate = new FormData();
      formDate.append("file", params.file);
      // console.log(formDate)
      this.form.fileName = params.file.name;
      var str = "";
      // str = params.
      // console.log(this.pps);
    },
    uploadDataset() {
      //提交表单，上传到数据库里
      var myDate = new Date();
      this.form.fileTime = myDate.toLocaleString();
      if (this.form.fileType === "预测集")
        this.$message.success("预测集 " + this.form.fileName + " 上传成功!");
      else if (this.form.fileType === "训练集")
        this.$message.success("训练集 " + this.form.fileName + " 上传成功!");
      this.dialogVisible = false;
      this.tableData.push({name:this.form.fileName,type:this.form.fileType,uploadtime:this.form.fileTime,info:this.form.fileInfo})
    }
  }
  // components: {
  //   HelloWorld
  // }
};
</script>
<style scoped>
*{
font-size: 18px;
}
.upload {
  margin: 50px ;
}

.el-table--border::after, .el-table--group::after, .el-table::before{
  background-color:#009FCC;
}

.datasetControl {
  /* float: left; */
  font-family:'微软雅黑';
  font-size: 18px;
  display: inline-block;
}

.datasetTable {
  /* display: inline; */
  margin: 10px;
}


</style>
