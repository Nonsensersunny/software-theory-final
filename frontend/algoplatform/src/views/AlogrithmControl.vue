<template>
  <div class="algoControl">
    <!-- <img alt="Vue logo" src="../assets/logo.png"> -->
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <center>
    <div class="upload">
      <el-button size="small" type="primary" @click="dialogVisible = true">上传算法</el-button>
    </div>
    <div class="datasetTable" >
      <el-table :data="tableData" border style="font-size:18px;border-color:#009FCC" :row-style="{height:'80px'}"  >
        <el-table-column type="index"  align="center"  label="序号" width="100px"></el-table-column>
        <el-table-column prop="name" label="算法名称"  align="center"  width="320px"></el-table-column>
        <el-table-column prop="info" label="描述"  align="center"  width="300px"></el-table-column>
        <el-table-column prop="uploadtime" label="上传时间"  align="center"  width="280px"></el-table-column>
        <el-table-column prop="work" label="操作"  align="center"  width="280px">
          <template slot-scope="scope">
            <el-button style="color:#409EFF"  @click="check(scope.row)" >查看</el-button>
            <el-button style="color:#409EFF"  @click="delete(scope.row)" >删除</el-button>
          </template>
        </el-table-column>
        
      </el-table>
    </div>

    <el-dialog title="上传数据集" :visible.sync="dialogVisible" width="30%">
      <el-upload
        accept=".py"
        class="upload-demo"
        action="/"
       
        :show-file-list="false"
        :on-success="upload1"
        :before-upload="beforeupload"
        :http-request="uploadSuccess"
      >
        <el-button class="btn_upload"  type="primary" style="float:left">选择文件</el-button>
        <div slot="tip" style="display:inline-block;margin-bottom:20px;margin-left:10px;" class="el-upload__tip">只能上传python文件</div>
      </el-upload>
      <el-form style="width:590px;">
              <el-form-item :model="form">
        <el-form-item label="算法文件" :label-width="formLabelWidth" style="width:100%;margin-bottom:10px">
          <el-input v-model="form.src" disabled></el-input>
        </el-form-item>
        <el-form-item label="算法名称" :label-width="formLabelWidth" style="width:100%;margin-bottom:10px">
          <el-input v-model="form.fileName" ></el-input>
        </el-form-item>
        <!-- <el-form-item label="类型" :label-width="formLabelWidth">
          <el-select v-model="form.fileType" placeholder="请选择" style="width:100%;margin-bottom:10px">
            <el-option
              v-for="item in options"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            ></el-option>
          </el-select>
        </el-form-item> -->
        <el-form-item label="算法描述" :label-width="formLabelWidth" style="width:100%;margin-bottom:10px">
          <el-input type="textarea" v-model="form.fileInfo" :rows="3" placeholder="请输入对算法的描述，如算法的特点、适用情况，"></el-input>
        </el-form-item>
      </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="uploadDataset">确 定</el-button>
      </div>
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
      formLabelWidth: "120px",
      form: {
        src:'',
        fileName: "", //数据集名字
        fileType: "", //训练集还是预测集
        fileInfo: "", //描述
        fileTime: "", //上传时间
        fileList: "" //数据
      },
     
      // options: [
      //   { value: "预测集", label: "预测集" },
      //   { value: "训练集", label: "训练集" }
      // ],
      tableData: [
        {
          src:'svm.py',
          name: "SVM向量机算法",
          info: "适用于.....",
          uploadtime: "2019-10-30 下午06：30"
        },{
          src:'bayes.py',
          name: "朴素贝叶斯算法",
          info: "适用于.....",
          uploadtime: "2019-10-30 下午06：28"
        }
      ]
    };
  },
  methods: {
    beforeupload(file) {
      //    let fd = new FormData();
      // fd.append('file',file);//传文件
      //   //  fd.append('srid',this.aqForm.srid);//传其他参数
      // axios.post('/api/up/file',fd).then(function(res){
      //         alert('成功');
      // })
    },
    check(params){},
    delete(params){},
    // checkFile(params) {
    //   // if(params.type === "application/vnd.ms-excel")
    //   // console.log(params)
    //   // console.log(params.type)
    //   // return true
    // },
    upload1(response, file) {
      // console.log(response)
      // console.log(file)
    },
    uploadSuccess(params) {
      //上传到前端，缓存 ?
      const formDate = new FormData();
      formDate.append("file", params.file);
      // console.log(formDate)
      this.form.src = params.file.name;
      var str = "";
      // str = params.
      // console.log(this.pps);
    },
    uploadDataset() {
      //提交表单，上传到数据库里
      var myDate = new Date();
      this.form.fileTime = myDate.toLocaleString();
      if (this.form.fileType === "预测集")
        this.$message.success("预测集 " + this.form.src + " 上传成功!");
      else if (this.form.fileType === "训练集")
        this.$message.success("训练集 " + this.form.src + " 上传成功!");

      // console.log(this.form)
      this.dialogVisible = false;
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
   margin: 50px  auto ;
}
.el-table--border::after, .el-table--group::after, .el-table::before{
  background-color:#009FCC;
}

.algoControl {
  /* float: left; */
  font-family:'微软雅黑';
  font-size: 18px;
  display: inline-block;
}

.datasetTable {
  /* display: inline; */
  margin-left: 150px;
}


</style>
