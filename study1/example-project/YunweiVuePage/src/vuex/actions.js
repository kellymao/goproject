//test


import { login } from '@/api/api'
import router from '@/router/index'

import Home from '@/pages/Home.vue'
import echarts from '@/pages/charts/echarts.vue'


export const increment = ({commit}) => {
    commit('INCREMENT')
}
export const decrement = ({commit}) => {
    commit('DECREMENT')
}



export const  LoginIn = ({commit},loginInfo) => {

    login(loginInfo).then(res=>{


        console.log(res);
        commit('setUserInfo', res.data.user);
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
                console.log('yyyyyyyy');
                router.push({ path: redirect, replace: true });

            }else{
                console.log('xxxxxxxxx');
                router.push({ path: '/table' });

            }
        }



    }); // res.success 如果为false 直接 拦截了。不会往下走


}

export const  SetAsyncRouter = ({commit}) => {


    const baseRouter =  [
        {
            path: '/',
                component: Home,
            name: '导航四',
            iconCls: 'stats-bars',
            children: [
            { path: '/echarts', component: echarts, name: 'echarts' }
        ]
        }
    ];

    // const asyncRouterRes =  asyncMenu()
    // const asyncRouter = asyncRouterRes.data.menus

    // asyncRouterHandle(baseRouter)
    commit('setAsyncRouter', baseRouter)



}