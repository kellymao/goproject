import router from './router'
import  store  from '@/vuex/store'

import Home from '@/pages/Home.vue'
import echarts from '@/pages/charts/echarts.vue'


let asyncRouterFlag = false;

router.beforeEach((to, from, next) => {
    if (to.path == '/login') {
        sessionStorage.removeItem('user');
        sessionStorage.removeItem('token');
        next()
        return
    }

    // asyncRouterFlag++ // 浏览器每次打开一个新页面，就从0 开始重新计数。 在同一页面点不同的菜单是直接+1，不重新计数
    // console.log('asyncRouterFlag is ' + asyncRouterFlag);



    console.log(router);
    //let token = JSON.parse(sessionStorage.getItem('token'));
    let token = sessionStorage.getItem('token');

    if (token) {
        //if (!asyncRouterFlag){ // 浏览器每次打开一个新页面时，只执行一次

        alert(from.path + ' ' + to.path + ' ' + sessionStorage.getItem('flag'));
        if (!sessionStorage.getItem('flag')){
            asyncRouterFlag = true ;
            sessionStorage.setItem('flag',"1");
            store.dispatch('SetAsyncRouter');
            // const aaa = store.getters['asyncRouters'];


            const baseRouter =  [
                {
                    path: '/',
                    component: Home,
                    name: '导航四',
                    iconCls: 'stats-bars',
                    children: [
                        { path: '/table', component: echarts, name: 'echarts' }
                    ]
                }
            ];




        }

        next()
        return
    }



    if (!token && to.path != '/login') { // token 不存在并且访问的非登录页面
        next({ path: '/login' ,
               query: {
                    redirect: to.path
                }
            })
    } else {

        alert('why');
        next()
    }
})
