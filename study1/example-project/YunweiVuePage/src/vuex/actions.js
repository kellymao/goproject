//test


import { login,getmenu } from '@/api/api'
import router from '@/router/index'

import Home from '@/pages/Home.vue'
import echarts from '@/pages/charts/echarts.vue'
import {prase_route} from "../common/parse_component";


export const increment = ({commit}) => {
    commit('INCREMENT')
}
export const decrement = ({commit}) => {
    commit('DECREMENT')
}



export const  LoginIn = ({commit},loginInfo) => {

    login(loginInfo).then(res=>{


        commit('setUserInfo', res.data.user); // "user":{"ID":1,"CreatedAt":"2019-09-13T17:23:46+08:00","UpdatedAt":"2019-10-21T11:16:03+08:00","DeletedAt":null,"uuid":"ce0d6685-c15f-4126-a5b4-890bc9d2356d","userName":"admin","nickName":"超级管理员","headerImg":"http://www.henrongyi.top/avatar/lufu.jpg","authority":{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"authorityId":"","authorityName":""}}}
        commit('setToken', res.data.token);
        commit('setExpiresAt', res.data.expiresAt);


        if (res.success) {

            //
            // const baseRouter = [
            //     {
            //         path: '/',
            //             component: Home,
            //         name: '导航四',
            //         iconCls: 'stats-bars',
            //         children: [
            //         { path: '/echarts', component: echarts, name: 'echarts' }
            //     ]
            //     }
            // ]
            //
            //
            // router.addRoutes(baseRouter);
            //
            // console.log(router.options);

            const redirect = router.history.current.query.redirect;
            if( redirect){
                router.push({ path: redirect, replace: true });

            }else{
                router.push({ path: '/table' });

            }
        }



    }); // res.success 如果为false 直接 拦截了。不会往下走


}

export const  SetAsyncRouter = ({commit}) => {


    getmenu().then(res=>{





        if (res.success) {

          const baseRouter = res.data.menu ;  // 处理一个菜单列表

          baseRouter.push({
            path: '*',
            name: 'other',
            hidden: true,
            redirect: '/404'


          })

          prase_route(baseRouter);

          router.options.routes = router.options.routes.concat(baseRouter); // 静态路由和动态路由合并后，重新赋值。
          router.addRoutes(baseRouter); // 不能实时更新

          console.log(router.options.routes);
          commit('setAsyncRouter', router.options.routes)
        }



    });

    // const baseRouter =  [
    //     {
    //         path: '/',
    //             component: Home,
    //         name: '导航四',
    //         iconCls: 'stats-bars',
    //         children: [
    //         { path: '/echarts', component: echarts, name: 'echarts' }
    //     ]
    //     }
    // ];

    // const asyncRouterRes =  asyncMenu()
    // const asyncRouter = asyncRouterRes.data.menus

    // asyncRouterHandle(baseRouter)
    // commit('setAsyncRouter', baseRouter)



}
