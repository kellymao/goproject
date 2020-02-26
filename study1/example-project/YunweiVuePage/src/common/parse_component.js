const _import = require('./_import') //获取组件的方法


export const prase_route = (asyncRouter) => {
  asyncRouter.map(item => {
    if (item.component) {

      item.component = _import(item.component)
    } else {
      delete item['component']
    }
    if (item.children) {
      prase_route(item.children)
    }
  })
}

/*
js中的列表的 filter方法和map方法 :https://blog.51cto.com/11871779/2126561

*/
