<template>
    <div id="login">
        <el-form :model="user">
            <el-form-item label="Mail" required><el-input v-model="user.mail" /> </el-form-item>
            <el-form-item label="Password" required><el-input v-model="user.password" /> </el-form-item>
        </el-form>
        <div>
            <el-button @click="login">Login</el-button>
            <el-button @click="register">Register</el-button>
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
                        this.$router.push("dataset")
                    }
                } catch (e) {
                    console.log(e)
                }
            },
            async register() {
                try {
                    let resp = this.$store.dispatch("register", this.user);
                } catch (e) {
                    console.log(e)
                }
            }
        }
    }
</script>

<style scoped>

</style>