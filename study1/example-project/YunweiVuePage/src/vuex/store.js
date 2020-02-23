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
    asyncRouters: [],
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
        state.userInfo = userInfo

        sessionStorage.setItem('user', userInfo.user);
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
        state.asyncRouters = asyncRouters
    }
}

// 创建 store 实例
export default new Vuex.Store({
    actions,
    getters,
    state,
    mutations
})