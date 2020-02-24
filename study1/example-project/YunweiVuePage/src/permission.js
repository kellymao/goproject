import router from './router'
import  store  from '@/vuex/store'

let asyncRouterFlag = false;

router.beforeEach((to, from, next) => {
    if (to.path == '/login') {
        sessionStorage.removeItem('user');
        sessionStorage.removeItem('token');
    }

    // asyncRouterFlag++ // 浏览器每次打开一个新页面，就从0 开始重新计数。 在同一页面点不同的菜单是直接+1，不重新计数
    // console.log('asyncRouterFlag is ' + asyncRouterFlag);



    console.log(router);
    //let token = JSON.parse(sessionStorage.getItem('token'));
    let token = sessionStorage.getItem('token');

    if (token) {
        if (!asyncRouterFlag){ // 浏览器每次打开一个新页面时，只执行一次
            asyncRouterFlag = true ;
            store.dispatch('SetAsyncRouter')
            const asyncRouters = store.getters['asyncRouters']
            console.log("-----------");
            console.log(asyncRouters);
            console.log("-----------");
            router.options.routes = router.options.routes.concat(asyncRouters);
            router.addRoutes(asyncRouters); // 不能实时更新

            console.log(router.options);



        }
    }



    if (!token && to.path != '/login') { // token 不存在并且访问的非登录页面
        next({ path: '/login' ,
               query: {
                    redirect: to.path
                }
            })
    } else {
        next()
    }
})
