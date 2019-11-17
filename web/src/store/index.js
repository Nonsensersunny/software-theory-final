import Vue from 'vue'
import Vuex from 'vuex'
import {UserHttp, GuestHttp} from "@/assets/js/api";

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    isLogin: false,
    userInfo: {},
  },
  getters: {
    loginStatus: (state) => { return state.isLogin; }
  },
  mutations: {
    changeLoginStatus(state, status) {
      state.isLogin = status;
    }
  },
  actions: {
    async login(context, playload) {
      return await GuestHttp.signin(playload);
    },
    async register(context, playload) {
      return await GuestHttp.register(playload);
    },
    async logout(context) {
      return await UserHttp.logout();
    },
    async createDataset(context, dataset) {
      return await UserHttp.createDataset(dataset);
    },
    async getDatasets(context) {
      return await UserHttp.getDatasets();
    },
    async getAlgorithms(context) {
      return await UserHttp.getAlgorithms();
    },
    async createPrediction(context, {trainId, testId, algorithmId}) {
      return await UserHttp.createPrediction(trainId, testId, algorithmId);
    }
  }
})