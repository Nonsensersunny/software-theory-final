<template>
  <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm login-container">
  <el-form-item label="账号" prop="name">
    <el-input v-model="ruleForm.mail"></el-input>
  </el-form-item>
   <el-form-item label="密码" prop="pass">
    <el-input type="password" v-model="ruleForm.password" auto-complete="off"></el-input>
  </el-form-item>
  <el-form-item>
    <el-button type="primary" @click="init">提交</el-button>
  </el-form-item>
  </el-form>
</template>

<script>
  import {GuestHttp} from '@/axios/api.js'
export default {
  data() {
    return {
      ruleForm: {
        mail: '',
        password: ''
      },
      rules: {
        name: [
          { required: true, message: '请输入登录账号', trigger: 'blur' }
        ],
        pass: [
          { required: true, message: '请输入登录密码', trigger: 'blur' }
        ]
      }
    }
  },

  created(){
    // this.init()
  },
  methods:{
    init:function(){
      console.log(this.ruleForm)
      var data = new FormData();
      // let 
      data.append("mail",this.ruleForm.mail);
      data.append("password",this.ruleForm.password);
      // console.log(data.get("mail"));
      // console.log(this.$axios.defaults.baseURL)
      this.$axios.put(this.$axios.defaults.baseURL+'/user/',data)
      .then(response=>{
        console.log(response)
      })
      // var str= this.GuestHttp.register(this.ruleForm)
      // console.log(str)
    }
  }
};
</script>