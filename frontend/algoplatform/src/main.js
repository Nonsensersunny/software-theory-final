// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
Vue.use(ElementUI);
Vue.config.productionTip = false
import axios from 'axios';
Vue.prototype.$axios = axios;
import VueCookie from 'vue-cookie'
Vue.use(VueCookie)
import echarts from 'echarts'
Vue.use(echarts)
Vue.prototype.$echarts = echarts
// Vue.use(Cookies)
axios.defaults.timeout = 50000;
// axios.defaults.headers.put['Access-Control-Allow-Origin']=''
// axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded;charset=utf-8';
// axios.defaults.headers.get['Content-Type'] = 'application/x-www-form-urlencoded;charset=utf-8';
// axios.defaults.headers.put['Content-Type'] = 'application/x-www-form-urlencoded;charset=utf-8';
// axios.defaults.baseURL = "http://106.13.90.235:10000";
axios.defaults.baseURL = "http://47.94.233.125:10000";

axios.defaults.withCredentials = true;
// Vue.http.interceptors.push(function(request,next){
//   request.credentials = true;
// })
// Vue.http.options.xhr = { withCredentials: true }
/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
