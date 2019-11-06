<template>
  <div id="app">
    <div class="head">
      <div class="block">
    <el-carousel height="150px">
      <el-carousel-item v-for="item in 4" :key="item">
        <h1> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;疾 病 分 析 与 预 测 平 台 </h1>
      </el-carousel-item>
    </el-carousel>
  </div>
  
 </div>
    <div id="nav">
      <el-col :span="18">
      
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
  height:150px;
  text-align: center;
  padding-top :30px;
  font-size: 28px;
  color: #fff;
  font-family: '微软雅黑';

  width: 100%;
}
.el-carousel__item h3 {
    color: #475669;
    font-size: 14px;
    opacity: 0.75;
    line-height: 150px;
    margin: 0;
  }

  .el-carousel__item:nth-child(2n) {
     background-color: #99a9bf;
  }
  
  .el-carousel__item:nth-child(2n+1) {
     background-color: #d3dce6;
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
