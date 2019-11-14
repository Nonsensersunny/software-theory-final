import axios from "axios";

import {ErrorCode, RespError} from "@/axios/type.js"
import {server} from '@/axios/config.js'

export const GuestHttp = {
  client: axios.create({
    baseURL: 'http://${server.host}:${server.port}'
  }),
  async uoload_dataset(dataset){
    let data = new FormData();
    data.set('测试连接',dataset.uid)
    data.set('',dataset.path)
    let res = await this.client.post('/dataset',dataset)
  },
  async register(user) {
    let data = new FormData();
    data.set('mail', user.mail);
    data.set('password', user.password);
    let res = await this.client.put('/user', user);
    return res;
  },
  async signin(user) {
    let data = new FormData();
    data.set('mail', user.mail);
    data.set('password', user.password);
    let res = await this.client.post('/user', user);
    return res;
  },
  async signup(user) {
    let data = new FormData();
    data.set('username', user.username);
    data.set('password', user.password);
    let res = await this.client.post('/register', user);
    return res;
  },
  async getDataset(dataset){
    let resp = await this.client.get('/dataset');
    return resp.data.data['data']

  }
}

export const UserHttp = {
  client: axios.create({
    baseURL: 'http://127.0.0.1:8080',
  }),

}

function errorHandler(e) {
  if (!e.response) {
    console.error(e);
    let error = new RespError(0, ErrorCode.Undefined, '');
    return Promise.reject(error);
  }
  let resp = e.response;
  if (resp.status >= 500) {
    let error = new RespError(resp.status, ErrorCode.ServerError, '');
    return Promise.reject(error);
  }
  let errorBody = resp.data.error;
  if (!errorBody || !errorBody.code) {
    let error = new RespError(resp.status, ErrorCode.Undefined, resp.data);
    return Promise.reject(error);
  }
  let error = new RespError(resp.status, errorBody.code, resp.msg);
  return Promise.reject(error);
}

// GuestHttp.client.interceptors.response.use(resp => resp, errorHandler);
// UserHttp.client.interceptors.response.use(resp => resp, errorHandler);
