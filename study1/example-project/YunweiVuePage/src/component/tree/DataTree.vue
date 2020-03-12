<style scoped>
  .data-tree-panel {
    width: 100%;
    position: relative;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
  }
  .tree-spin {
    position: absolute;
    left: 40%;
    top: 50%;
  }
</style>

<style>
  .data-tree-panel .tree-loading {
    opacity: 0.5
  }
  .data-tree-panel .ivu-tree .ivu-tree-title {
    font-size: 13px;
  }
  .data-tree-panel .ivu-tree .ivu-tree-arrow {
    font-size: 14px;
  }
</style>

<template>
  <div class="data-tree-panel">
    <Tree ref="tree" :showCheckbox="showCheckbox" :data="treeData" :class="{'tree-loading': loading}" size="large"
          @on-select-change="onSelectChange" @on-check-change="onCheckChange"></Tree>
    <spin v-if="loading" size="large" class="tree-spin"></spin>
  </div>
</template>

<script>
  import {prase_tree,menutree_convert_checked} from "@/common/parse_datatree";
  import {get_role_menutree,save_role_menutree} from "@/api/api";


  export default {
    props: {
      dataUrl: {},
      param: {},
      showCheckbox: {
        type: Boolean,
        default: false
      },
      constructTree: {
        type: Boolean,
        default: true
      },
      lastStep: {
        type: Boolean,
        default: false
      },
      lazy: {
        type: Boolean,
        default: false
      },
      roleid:''
    },
    data () {
      return {
        total: 0,
        loading: false,
        treeData: [],
        mData: [],
        tmp_roleid:''
      };
    },
    methods: {
      /**
       * 获取数据
       */
      queryData (url=this.dataUrl,roleid=this.roleid) {

        // const data3 = [
        //   {
        //     id: 2,
        //     title: '系统管理',
        //     expand: true,
        //     selected: true,
        //     children: [
        //       {
        //         id:7,
        //         title: '表格',
        //         expand: false,
        //         children: [
        //           {
        //             title: 'leaf 1-1-1',
        //             disabled: true
        //           },
        //           {
        //             title: 'leaf 1-1-2'
        //           }
        //         ]
        //       },
        //       {
        //         id:8,
        //         title: '表单',
        //         expand: false,
        //         children: [
        //           {
        //             title: 'leaf 1-2-1',
        //             checked: true
        //           },
        //           {
        //             title: 'leaf 1-2-1'
        //           }
        //         ]
        //       }
        //       ,
        //       {
        //         id:9,
        //         title: '用户管理',
        //         expand: false,
        //       }
        //       ,
        //       {
        //         id:10,
        //         title: '角色管理',
        //         expand: false,
        //       }
        //     ]
        //   }
        // ];

        get_role_menutree({url,roleid}).then(res=>{


          if (res.success) {

            //alert(JSON.stringify(res.data.menu));
            //menutree_convert_checked(res.data.menu);

            /*

            js 中数组赋值给新数组，原来的数组的值也跟着变了 ？？？？？？
            因此赋值的是引用。所以原数组值也变了。折腾了半天。

             this.treeData =res.data.menu;

             */
            //const rel = res.data.menu.slice(0);

            this.treeData =menutree_convert_checked(res.data.menu);

          }
        }).catch(err=>{
          alert(err);
        });


        //this.treeData = data3;

      },
      onSelectChange (slected) {
          alert(1);
          console.log(slected);
          alert(JSON.stringify(slected));
      },
      onCheckChange (slected) {

        //console.log(slected);
        //console.log(this.$refs.tree.getCheckedNodes()); // 打勾的时候触发
        //console.log(this.$refs.tree.getSelectedNodes());

        console.log(this.treeData);
      },
      /**
       * 选中结果包含父节点
       * @param parentId
       * @param data
       */
      getData () {
        return this.mData;
      },
      submit(){
        //const data = this.$refs.tree.getCheckedNodes();

        // copydata = $.extend(true, [], this.treeData)  js 深拷贝

        function copy (obj) {   // 自定义二维数组深拷贝 https://www.cnblogs.com/lvonve/p/11334628.html

          if (!obj){   // 排除 null
            return
          }

          var newobj = obj.constructor === Array ? [] : {};
          if(typeof obj !== 'object'){
            return;
          }
          for(var i in obj){
            newobj[i] = typeof obj[i] === 'object' ? copy(obj[i]) : obj[i];
          }
          return newobj
        }

        const copyobj = copy(this.treeData);
        let rel = prase_tree(copyobj,this.tmp_roleid); // 避免修改原数据

        console.log(rel);


        save_role_menutree(rel).then(res=>{


          if (res.success) {

            //alert(JSON.stringify(res.data.menu));
            //menutree_convert_checked(res.data.menu);

            /*

            js 中数组赋值给新数组，原来的数组的值也跟着变了 ？？？？？？
            因此赋值的是引用。所以原数组值也变了。折腾了半天。

             this.treeData =res.data.menu;

             */
            //const rel = res.data.menu.slice(0);

            this.$Message.info({
              content: '保存成功！',
              duration: 10
            });

          }else{
            alert("保存失败！")
          }
        }).catch(err=>{
          alert(err);
        });



      }
    },
    created () {

      //this.queryData();
      // if (!this.lazy) {
      //   this.queryData();
      // }
    },
    watch: {
      roleid(newvalue,oldvalue) {
        this.tmp_roleid = newvalue;
      }
    },
    components: {}
  };
</script>


