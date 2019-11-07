<template>
    <div id="loginDiv">
      <div class="title"><span>疾病分析与预测平台</span></div>
    <el-form :model="ruleForm"  ref="ruleForm" label-width="100px" class="demo-ruleForm login-container">
      <el-form-item label="账号" prop="name">
        <el-input v-model="ruleForm.mail"></el-input>
      </el-form-item>
       <el-form-item label="密码" prop="pass">
        <el-input type="password" v-model="ruleForm.password" auto-complete="off"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="register">注册</el-button>
        <el-button type="primary" @click="login">登录</el-button>
      </el-form-item>
    </el-form>
    </div>
</template>
<script>
  import {GuestHttp} from '@/axios/api.js'
import Cookies from 'js-cookie'

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
    register:function(){
      console.log(this.ruleForm)
      var data = new FormData();
      data.append("mail",this.ruleForm.mail);
      data.append("password",this.ruleForm.password);
      console.log(data.get("mail"));
      this.$axios.put(this.$axios.defaults.baseURL+'/user/',JSON.stringify(this.ruleForm))
      .then(response=>{
        console.log(response)
        var str = "账号： "+this.ruleForm.mail+" 注册成功！请登录！" ;
        this.$message({
          message:str,
          type:'success'
        })
      })
    },
    login:function(){
     this.$axios.post(this.$axios.defaults.baseURL+'/user/',JSON.stringify(this.ruleForm))
      .then(response=>{
        console.log(response)
        this.$router.push('/main/')
        // console.log(this.$cookies.set(this.ruleForm.mail))
        // console.log(Cookies.get())
      });
      // this.$axios.get(this.$axios.defaults.baseURL+'/user/',JSON.stringify(this.ruleForm))
      // .then(response=>{
      //   console.log(response)
      // })
    }
  }
};
</script>
<style type="text/css">
  #loginDiv{
    width: 500px;
    height: 250px;
    margin:100px auto;
    padding: 50px;
    text-align: center;
    border-radius: 15px;
    background-color: #BBFFEE;
  }
  .title{
    height: 50px;
    border-bottom-style: solid ;
    border-color: #fff;
    margin-bottom: 30px;
    border-width: 2px;
  }
  .title span{
    font-size: 26px;

  }
  .login-container{
    width: 80%;

  }
</style>