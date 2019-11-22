import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementUI from 'element-ui'
import echarts from 'echarts'
import 'element-ui/lib/theme-chalk/index.css'

Vue.config.productionTip = false;

router.beforeEach((to, from, next) => {
  if (!store.state.isLogin && to.path !== '/login' && to.path != '/') {
    ElementUI.Message.error("用户尚未登录");
    next('/login');
  }
  next();
})

Vue.use(ElementUI);
Vue.use(echarts)
Vue.prototype.$echarts = echarts
new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app');
