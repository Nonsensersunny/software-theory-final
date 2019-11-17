import Vue from 'vue'
import Vuex from 'vuex'
import {GuestHttp, UserHttp} from "@/assets/js/api";

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    isLogin: false,
    userInfo: {},
  },
  getters: {

  },
  mutations: {

  },
  actions: {
    async login(context, playload) {
      context.state.playload = GuestHttp.signin(playload);
      context.state.isLogin = true;
    },
    async register(context, playload) {
      return await GuestHttp.register(playload);
    },
    async logout(context) {
      return await UserHttp.logout();
    },
    async createDataset(context, dataset) {
      return await UserHttp.createDataset(dataset);
    }
  }
})
