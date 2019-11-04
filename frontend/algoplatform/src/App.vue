<template>
  <div id="app">
    <div class="head">疾 病 分 析 与 预 测 平 台</div>
    <div id="nav">
      <el-col :span="18">
        <h5>导航</h5>
        <el-menu
          default-active="1"
          class="el-menu-vertical-demo"
          @open="handleOpen"
          @close="handleClose"
          background-color="#DDDDDD"
          text-color="#000"
          active-text-color="#ffd04b"
        >
          <el-menu-item index="1">
            <i class="el-icon-document"></i>
            <router-link to="/dataset">数据集管理</router-link>
            <!-- <span slot="title">数据集管理</span> -->
          </el-menu-item>
          <el-menu-item index="2">
            <i class="el-icon-document"></i>
            <router-link to="/algorithm">算法管理</router-link>
          </el-menu-item>
          <el-submenu index="3">
            <template slot="title">
              <i class="el-icon-location"></i>
              <span font-size="18px">预测进程</span>
            </template>
            <el-menu-item-group>
              <el-menu-item index="3-1">
                <i class="el-icon-document"></i>
                <router-link to="/forecast/begin">开始预测</router-link>
              </el-menu-item>
              <el-menu-item index="3-2">
                <i class="el-icon-document"></i>
                <router-link to="/forecast/check">预测结果</router-link>
              </el-menu-item>
            </el-menu-item-group>
          </el-submenu>
          <!-- <el-menu-item index="3" >
            <i class="el-icon-notebook-1"></i>
              <router-link to='/'>预测结果</router-link>

          </el-menu-item>-->
          <el-menu-item index="4" disabled>
            <i class="el-icon-notebook-2"></i>
            <span slot="title">可视化</span>
          </el-menu-item>
        </el-menu>
      </el-col>
      <!-- <router-link to="/">Home</router-link> | -->
      <!-- <router-link to="/about">About</router-link> -->
    </div>

    <router-view />
  </div>
</template>
<script>
export default {
  methods: {
    handleOpen(key, keyPath) {
      console.log(key, keyPath);
    },
    handleClose(key, keyPath) {
      console.log(key, keyPath);
    }
  }
};
</script>
<style lang="stylus">
span{
  font-size :18px;
}
.head {
  height:100px;
  text-align: center;
  padding-top :30px;
  font-size: 28px;
  color: #fff;
  font-family: '微软雅黑';
  background-color: #888888;
  width: 100%;
}

#nav {
  width: 200px;
  height: 800px;
  background-color: #DDDDDD;
  float: left;
}

.el-menu {
  width: 200px;
}

.el-menu-item span, .el-menu-item a {
  font-size: 18px;
  color: #000;
  text-decoration: none;
}

.el-menu-item :focus, .el-menu-item :hover {
  color: #ffd04b;
}
</style>
