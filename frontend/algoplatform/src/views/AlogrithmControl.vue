<template>
  <div class="algoControl">
    <!-- <img alt="Vue logo" src="../assets/logo.png"> -->
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <center>
      <div class="upload">
        <el-button size="small" type="primary" @click="dialogVisible = true">上传算法</el-button>
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
          <el-table-column prop="info" label="描述" align="center"></el-table-column>
          <el-table-column prop="uploadtime" label="上传时间" align="center" ></el-table-column>
          <el-table-column prop="work" label="操作" align="center" >
            <template slot-scope="scope">
              <el-button style="color:#409EFF" size="small" @click="down(scope.row)">下载</el-button>
              <el-button style="color:#409EFF" size="small" @click="dele(scope.row)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>

      <el-dialog title="上传数据集" :visible.sync="dialogVisible" width="30%">
        <!-- <el-upload
          accept=".py"
          class="upload-demo"
          action="/"
          :show-file-list="false"
          :on-success="upload1"
          :before-upload="beforeupload"
          :http-request="uploadSuccess"
        >
          <el-button class="btn_upload" type="primary" style="float:left">选择文件</el-button>
          <div
            slot="tip"
            style="display:inline-block;margin-bottom:20px;margin-left:10px;"
            class="el-upload__tip"
          >只能上传python文件</div>
        </el-upload>-->
        <div style="width:590px;">
          <!-- <span>选择文件</span> -->
          <el-input type="file" v-model="form.src" style="width:85%;margin-bottom:10px"></el-input>
        </div>
        <el-form style="width:590px;">
          <el-form-item :model="form">
            <el-form-item
              label="算法名称"
              :label-width="formLabelWidth"
              style="width:100%;margin-bottom:10px"
            >
              <el-input v-model="form.fileName"></el-input>
            </el-form-item>
            <el-form-item
              label="算法文件"
              :label-width="formLabelWidth"
              style="width:100%;margin-bottom:10px"
            >
              <el-input v-model="form.src" disabled></el-input>
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
            </el-form-item>-->
            <el-form-item
              label="算法描述"
              :label-width="formLabelWidth"
              style="width:100%;margin-bottom:10px"
            >
              <el-input
                type="textarea"
                v-model="form.fileInfo"
                :rows="3"
                placeholder="请输入对算法的描述，如算法的特点、适用情况，"
              ></el-input>
            </el-form-item>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="dialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="uploadDataset">确 定</el-button>
        </div>
      </el-dialog>
      <el-dialog :title="show_algo.fileName" :visible.sync="data_table" width="600px">
        <el-form :ref="show_algo">
          <el-form-item label="名称" :label-width="formLabelWidth">
            <el-input v-model="show_algo.fileName" disabled style="width:80%;margin-bottom:10px"></el-input>
          </el-form-item>
          <el-form-item label="路径" :label-width="formLabelWidth">
            <el-input v-model="show_algo.filePath" disabled style="width:80%;margin-bottom:10px"></el-input>
          </el-form-item>
          <!-- <el-form-item label="类型" :label-width="formLabelWidth">
            <el-input v-model="show_dataset.fileType" disabled style="width:80%;margin-bottom:10px"></el-input>
          </el-form-item>-->
          <el-form-item label="上传时间" :label-width="formLabelWidth">
            <el-input v-model="show_algo.fileTime" disabled style="width:80%;margin-bottom:10px"></el-input>
          </el-form-item>
          <el-form-item label="描述" :label-width="formLabelWidth">
            <el-input
              type="textarea"
              v-model="show_algo.fileInfo"
              disabled
              :rows="3"
              style="width:80%;margin-bottom:10px"
            ></el-input>
          </el-form-item>
          <!-- </el-form> -->
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
      formLabelWidth: "120px",
      form: {
        src: "",
        fileName: "", //数据集名字
        fileInfo: "", //描述
        fileTime: "", //上传时间
      },
      show_algo: {
        fileName: "1", //数据集名字
        filePath: "",
        fileInfo: "", //描述
        fileTime: "" //上传时间
      },

    
      tableData: [
        {
          src: "svm.py",
          name: "SVM向量机算法",
          info: "适用于.....",
          uploadtime: "2019-10-30 下午06：30"
        },
        {
          src: "bayes.py",
          name: "朴素贝叶斯算法",
          info: "适用于.....",
          uploadtime: "2019-10-30 下午06：28"
        }
      ]
    };
  },
  methods: {
    
    down(params) {//测试能否执行Python文件
      
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
      this.form.fileTime = myDate.toLocaleString();
      
      this.$message.success("算法文件 " + this.form.fileName + " 上传成功!");
      this.dialogVisible = false;
      this.tableData.push({
        name: this.form.fileName,
        type: this.form.fileType,
        uploadtime: this.form.fileTime,
        info: this.form.fileInfo
      // this.tableData
      })
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

.datasetTable {
  /* display: inline; */
  width: 1050px;
  margin-left: 10px;
}
</style>
