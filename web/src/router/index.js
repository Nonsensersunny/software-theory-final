import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'login',
    component:  () => import('../views/Login.vue')
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  },
  {
    path: "/login",
    name: "login",
    component: () => import('../views/Login.vue')
  },
  {
    path: "/main",
    name: "main",
    component: () => import('../views/Main.vue')
  },
  {
    path: "/dataset",
    name: "dataset",
    component: () => import('../views/Dataset.vue')
  },
  {
    path: "/prediction",
    name: "prediction",
    component: () => import('../views/Prediction.vue')
  },
  {
    path: "/chart",
    name: "chart",
    component: () => import('../views/Chart.vue')
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
