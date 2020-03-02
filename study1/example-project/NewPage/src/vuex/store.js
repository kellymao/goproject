import Vue from 'vue'
import Vuex from 'vuex'
import * as actions from './actions'
import * as getters from './getters'
import router from '@/router/index'


Vue.use(Vuex)

// 应用初始状态
const state = {

    userInfo: {
        uuid: "",
        nickName: "",
        headerImg: "",
        authority: "",
    },
    token: "",
    expiresAt: "",
    count: 10,
    asyncRouters: "",
    flag: false,
}

// 定义所需的 mutations
const mutations = {
    INCREMENT(state) {
        state.count++
    },
    DECREMENT(state) {
        state.count--
    },

    setUserInfo(state, userInfo) {
        // 这里的 `state` 对象是模块的局部状态
        state.userInfo = userInfo // userInfo:{"ID":1,"CreatedAt":"2019-09-13T17:23:46+08:00","UpdatedAt":"2019-10-21T11:16:03+08:00","DeletedAt":null,"uuid":"ce0d6685-c15f-4126-a5b4-890bc9d2356d","userName":"admin","nickName":"超级管理员","headerImg":"http://www.henrongyi.top/avatar/lufu.jpg","authority":{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"authorityId":"","authorityName":""}}}

        sessionStorage.setItem('user', userInfo.userName);
    },
    setToken(state, token) {
        // 这里的 `state` 对象是模块的局部状态
        state.token = token
        sessionStorage.setItem('token', token);

    },
    setExpiresAt(state, expiresAt) {
        // 这里的 `state` 对象是模块的局部状态
        state.expiresAt = expiresAt
    },



    LoginOut(state) {
        state.userInfo = {};
        state.token = "";
        state.expiresAt = "";
        router.push({ name: 'login', replace: true })
    },

    setAsyncRouter(state, asyncRouters) {
        state.asyncRouters = asyncRouters;

        console.log("检查");
        console.log(asyncRouters);
    },

    setflag(state){
        state.flag = true;
    }
}

// 创建 store 实例
export default new Vuex.Store({
    actions,
    getters,
    state,
    mutations
})