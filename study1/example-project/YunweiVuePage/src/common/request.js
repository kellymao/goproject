import axios from 'axios'; // 引入axios

import store from '@/vuex/store'
import { Message } from 'iview'



const service = axios.create({
    //baseURL: process.env.VUE_APP_BASE_API,
    baseURL: '/api',
    timeout: 5000
})

//http request 拦截器
service.interceptors.request.use(
    config => {
        const token = store.getters['token']
        config.data = JSON.stringify(config.data);
        config.headers = {
            'Content-Type': 'application/json',
            'x-token': token
        }
        return config;
    },
    error => {

        Message.error(error);
        return Promise.reject(error);
    }
);


//http response 拦截器
service.interceptors.response.use(
    response => {

        if (response.data.success) {
            return response.data
        } else {
            /*Message({
                showClose: true,
                message: response.data.msg,
                type: 'error'
            })*/
            Message.error(response.data.msg);
            if (response.data.data && response.data.data.reload) {
                store.commit('LoginOut')
            }


            return Promise.reject(response.data.msg)
        }
    },
    error => {
        Message.error(error);
        return Promise.reject(error)
    }
)

export default service

/*

https://www.jianshu.com/p/86122178002a 封装axios
 */