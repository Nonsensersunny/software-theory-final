import axios from 'axios'
import Vue from 'vue'
axios.defaults.baseURL = 'http://106.13.90.235:10000' // 配置你的接口请求地址
// axios.defaults.headers.common['Authorization'] = AUTH_TOKEN //配置token,看情况使用
axios.defaults.headers.post['Content-Type'] =
  'application/x-www-form-urlencoded' // 配置请求头信息。

Vue.prototype.$axios = axios

// 请求拦截器
axios.interceptors.request.use(
  function (config) {
    return config
  },
  function (error) {
    return Promise.reject(error)
  }
)
// 响应拦截器
axios.interceptors.response.use(
  function (response) {
    return response
  },
  function (error) {
    return Promise.reject(error)
  }
)

// 封装axios的post请求
export function fetch (url, params) {
  return new Promise((resolve, reject) => {
    axios
      .post(url, params)
      .then(response => {
        resolve(response.data)
      })
      .catch(error => {
        reject(error)
      })
  })
}

// 封装axios的post请求
export function fetchfile (url, params) {
  return new Promise((resolve, reject) => {
    axios
      .post(url, params, {
        responseType: 'blob'
      })
      .then(response => {
        resolve(response.data)
      })
      .catch(error => {
        reject(error)
      })
  })
}

export function getFiles (url) {
  return new Promise((resolve, reject) => {
    axios({
      method: 'get',
      url,
      responseType: 'arraybuffer'
    })
      .then(data => {
        resolve(data.data)
      })
      .catch(error => {
        reject(error.toString())
      })
  })
}

export default {
  JH_news (url, params) {
    return fetch(url, params)
  },
  Hosturl () {
    return axios.defaults.baseURL + '/'
  },
  getFile (url) {
    return getFiles(url)
  },
  JH_File (url, params) {
    return fetchfile(url, params)
  }
}