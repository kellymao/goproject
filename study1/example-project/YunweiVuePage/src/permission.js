import router from './router'
import  store  from '@/vuex/store'

let asyncRouterFlag = false;

router.beforeEach((to, from, next) => {
    if (to.path == '/login') {
        localStorage.removeItem('user');
        localStorage.removeItem('token');
        asyncRouterFlag = false;
    }

    // asyncRouterFlag++ // 浏览器每次打开一个新页面或者刷新页面，就从0 开始重新计数。 在同一页面里面点不同的菜单是直接+1，不重新计数
    // console.log('asyncRouterFlag is ' + asyncRouterFlag);


  //let token = JSON.parse(sessionStorage.getItem('token'));
    let token = localStorage.getItem('token');

    if (token) {
        if (!asyncRouterFlag){ // 浏览器每次打开一个新页面时，只执行一次
            asyncRouterFlag = true ;
            store.dispatch('SetAsyncRouter')
            //
            // router.options.routes = router.options.routes.concat(asyncRouters);
            // router.addRoutes(asyncRouters); // 不能实时更新
            //
            // console.log(router.options);
            alert("从后端获取route");
            next({...to, replace: true})




        }else{
            alert("从router.options.routes获取route");
            next()
        }

        return
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
