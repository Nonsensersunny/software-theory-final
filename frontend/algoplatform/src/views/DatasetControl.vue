<template>
  <div class="datasetControl">
    <!-- <img alt="Vue logo" src="../assets/logo.png"> -->
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <div class="upload">
      <el-button size="small" type="primary" @click="dialogVisible = true">上传数据集</el-button>
    </div>
    <div class="datasetTable"></div>

    <el-dialog title="上传数据集" :visible.sync="dialogVisible" width="30%">
      <el-upload
        accept=".csv"
        class="upload-demo"
        action="/"
     
        :show-file-list="false"
        :http-request="uploadSuccess"
      >
        <el-button class="btn_upload" size="small" type="primary">上传</el-button>
        <div slot="tip" class="el-upload__tip">只能上传csv/excel文件</div>
      </el-upload>
      <el-form :model="form">
        <el-form-item label="数据集名称" :label-width="formLabelWidth">
          <el-input v-model="form.fileName" disabled></el-input>
        </el-form-item>
        <el-form-item label="类型" :label-width="formLabelWidth">
          <el-select v-model="form.fileType" placeholder="请选择" style="width:453.13px;">
            <el-option
              v-for="item in options"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="描述" :label-width="formLabelWidth">
          <el-input type="textarea" v-model="form.fileInfo" :rows="3" placeholder="请输入对数据集的描述"></el-input>
        </el-form-item>
      </el-form>

      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="uploadDataset">确 定</el-button>
      </div>
    </el-dialog>
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
        fileName: "", //数据集名字
        fileType: "", //训练集还是预测集
        fileInfo: "", //描述
        fileTime: "", //上传时间
        fileList: "" //数据
      },
      options: [
        { value: "预测集", label: "预测集" },
        { value: "训练集", label: "训练集" }
      ]
    };
  },
  methods: {
    // checkFile(params) {
    //   // if(params.type === "application/vnd.ms-excel")
    //   // console.log(params)
    //   // console.log(params.type)
    //   // return true
    // },
    uploadSuccess(params) {//上传到前端，缓存 ?
      this.form.fileName = params.file.name;
      
      // console.log(params);
    },
    uploadDataset() {
      //提交表单，上传到数据库里
      var myDate = new Date();
      this.form.fileTime = myDate.toLocaleString();
      if(this.form.fileType ==="预测集")
        this.$message.success("预测集 "+this.form.fileName + " 上传成功!");
      else if(this.form.fileType === "训练集")
        this.$message.success("训练集 "+this.form.fileName + " 上传成功!");
      
      console.log(this.form)
      this.dialogVisible = false;
    }
  }
  // components: {
  //   HelloWorld
  // }
};
</script>
<style scoped>
.upload {
  margin: 10px;
}
/* .btn_upload{
  float: left;
} */
</style>
