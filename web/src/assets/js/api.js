import axios from "axios";

import {ErrorCode, RespError, User, Dataset, Prediction} from "@/assets/js/type.js"
import {server} from '@/assets/js/config.js';
import da from "element-ui/src/locale/lang/da";

export const GuestHttp = {
  client: axios.create({
    baseURL: `http://${server.host}:${server.port}`,
    withCredentials: true
  }),
  async register(user) {
    let res = await this.client.put('user/', user);
    return res.data;
  },
  async signin(user) {
    let res = await this.client.post('user/', user);
    return res.data;
  },
}

export const UserHttp = {
  client: axios.create({
    baseURL: `http://${server.host}:${server.port}`,
    withCredentials: true,
  }),
  async logout() {
    let res = await this.client.delete('user/');
    return res;
  },
  async createDataset(dataset) {
    let data = new FormData();
    data.append("name", dataset.name);
    data.append("type", dataset.type);
    data.append("upload", dataset.upload);
    data.append("description", dataset.description);
    let resp = await this.client.put("dataset/", data);
    return resp.data;
  },
  async getDatasets() {
    let resp = await this.client.get("dataset/");
    return resp.data;
  },
  async getAlgorithms() {
    let resp = await this.client.get("algorithm/");
    return resp.data;
  },
  async createPrediction(train, test, algorithm) {
    // let data = new FormData();
    // data.append('train', train);
    // data.append('test', test);
    // data.append('algorithm', algorithm)
    let data = {
      train: train,
      test: test,
      algorithm: algorithm
    }
    let resp = await this.client.post('prediction/', data);
    return resp.data;
  }
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

GuestHttp.client.interceptors.response.use(resp => resp, errorHandler);
UserHttp.client.interceptors.response.use(resp => resp, errorHandler);
