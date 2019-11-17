import Vue from 'vue'
import App from './App'
import router from './router'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import axios from 'axios';
import echarts from 'echarts'

Vue.use(ElementUI);
Vue.config.productionTip = false;
Vue.prototype.$axios = axios;

Vue.use(echarts)
Vue.prototype.$echarts = echarts;

new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
