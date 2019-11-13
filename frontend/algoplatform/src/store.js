import Vue from 'vue'
import Vuex from 'vuex'
import {GuestHttp, UserHttp} from "@/axios/api";
// import {langs, i18n} from "./i18n";

Vue.use(Vuex)


export default new Vuex.Store({
  state: {
    dataset_url:'',
    isLogin: false,
    contents: [],
    rankeduser: [],
    profile: {},
    // lang: "en",
    // langs: langs,
    avatar: '',
    cats: []
  },
  getters: {
    rankedUsers(state) { return state.rankeduser },
    profile(state) { return state.profile },
    // selectedLang(state) { return state.lang },
    // langs(state) { return state.langs },
    categories(state) { return state.cats },
    wsAddr(state) { return UserHttp.getWSAddr() }
  },
  mutations: {
    selectLang(state, lang) {
      // state.lang = lang
      // i18n.locale = lang
    }
  },
  actions: {
    async uoload_dataset(context,fileinfo){
      return await GuestHttp.upload_dataset(fileinfo);
    },
    async checkLoginStatus(context) {
      let resp = await UserHttp.checkLoginStatus()
      if (resp) {
        context.state.profile = resp
        context.state.isLogin = true
      } else {
        context.state.profile = {}
        context.state.isLogin = false
      }
    },
    async userLogin(context, playload) {
      context.state.profile = await GuestHttp.signin(playload);
      context.state.isLogin = true
    },
    async userRegister(context, playload) {
      return await GuestHttp.signup(playload);
    },
    async userLogout(context) {
      await UserHttp.userLogout(context.state.profile.username)
      context.state.isLogin = false
      context.state.profile = {}
    },
    async getProfileByName(context, username) {
      return await GuestHttp.getProfileByName(username);
    },
    async getProfileById(context, id) {
      return await GuestHttp.getProfileById(id);
    },
    async getRankedUsers(context) {
      context.state.rankeduser = await UserHttp.getRankedUsers()
    },
    async getContentById(context, id) {
      return await GuestHttp.getContentById(id);
    },
    async createContent(context, {title, author, category, content}) {
      return await UserHttp.createContent(title, author, category, content);
    },
    async getRankedContent(context) {
      return await GuestHttp.getRankedContent()
    },
    async getContentByCat(context, cat) {
      return await GuestHttp.getContentByCat(cat)
    },
    async getTopContent(context) {
      return await UserHttp.getTopContent()
    },
    async updateUserProfile(context, user) {
      context.state.profile = await UserHttp.updateUserProfile(user)
    },
    async getCategories(context) {
      context.state.cats = await GuestHttp.getCategories()
    },
    async createComment(context, comment) {
      return await UserHttp.createComment(comment)
    },
    async userGetContentById(context, {cid, uid}) {
      return await UserHttp.userGetContentById(cid, uid)
    },
    async createFavorite(context, {uid, cid}) {
      return await UserHttp.createFavorite(uid, cid)
    },
    async deleteFavorite(context, {uid, cid}) {
      return await UserHttp.deleteFavorite(uid, cid)
    },
    async createLink(context, link) {
      return await UserHttp.createLink(link)
    },
    async deleteLink(context, id) {
      return await UserHttp.deleteLink(id)
    },
    async getLinksByUserId(context, id) {
      return await UserHttp.getLinksByUserId(id)
    },
    // async createCategory(context, category) {
    //   return await AdminHttp.createCategory(category)
    // },
    // async toggleKeyContent(context, id) {
    //   return await AdminHttp.toggleKeyContent(id)
    // },
    async getCommentsByUid(context, id) {
      return await UserHttp.getCommentsByUid(id)
    },
    async delCommentById(context, id) {
      return await UserHttp.delCommentById(id)
    },
    async getAuthMail(context, user) {
      return await GuestHttp.getAuthMail(user)
    }
  }
})