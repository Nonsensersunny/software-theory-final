<template>
    <div id="login">
        <div id="title">疾 病 分 析 与 预 测 平 台</div>
        <el-form :model="user">
            <el-form-item label="账号" required><el-input v-model="user.mail" /> </el-form-item>
            <el-form-item label="密码"  required><el-input v-model="user.password" type="password" /> </el-form-item>
        </el-form>
        <div>
            <el-button @click="login">登录</el-button>
            <el-button @click="register" >注册</el-button>
        </div>
    </div>
</template>

<script>
    export default {
        name: "Login",
        data() {
            return {
                user: {}
            }
        },
        methods: {
            async login() {
                try {
                    let resp = await this.$store.dispatch("login", this.user);
                    if (resp.code !== 0) {
                        this.$message.error("Login Failed");
                    } else {
                        this.$message.success("Login Succeeded");
                        this.$store.commit("changeLoginStatus", true);
                        this.$router.push("Main")
                    }
                } catch (e) {
                    console.log(e)
                }
            },
            async register() {
                if(this.user.mail === "123456")
                    this.$message.error("账号已注册")
                else{
                    try {
                        let resp = this.$store.dispatch("register", this.user);
                    } catch (e) {
                        console.log(e)
                    }
                }
                
            }
        }
    }
</script>

<style scoped>
#login{
    width:30%;
    margin: 0 auto;
    border:steelblue 1px solid;
    padding: 50px 50px 50px 50px ;
    border-radius: 15px;
}
#title{
    height: 70px;
    font-family: 'Franklin Gothic Medium', 'Arial Narrow', Arial, sans-serif;
    border-bottom:1px solid steelblue;
}
</style>