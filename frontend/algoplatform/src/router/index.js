import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: '登录',
      component: () => import(/* webpackChunkName: "about" */ '../views/login.vue')
  
      // component: DatasetControl
    },{
      path: '/main',
      name: '主界面',
      component: () => import(/* webpackChunkName: "about" */ '../views/main.vue')
  
      // component: DatasetControl
    },
    {
      path: '/dataset',
      name: '数据集管理',
      component: () => import(/* webpackChunkName: "about" */ '../views/DatasetControl.vue')
  
      // component: DatasetControl
    },
    {
      path: '/dataset',
      name: '数据集管理',
      component: () => import(/* webpackChunkName: "about" */ '../views/DatasetControl.vue')
  
      // component: DatasetControl
    },
    {
      path: '/algorithm',
      name: '算法管理',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ '../views/AlogrithmControl.vue')
    },
    {
      path: '/forecast',
      name: '预测进程',
      component: () => import(/* webpackChunkName: "about" */ '../views/begin_forecast.vue')
    },
    {
      path: '/forecast/begin',
      name: '开始预测',
      component: () => import(/* webpackChunkName: "about" */ '../views/begin_forecast.vue')
    }, 
    {
      path: '/forecast/check',
      name: '预测结果',
      component: () => import(/* webpackChunkName: "about" */ '../views/check_forecast.vue')
    },
  ]
})
