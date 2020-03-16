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
    <Tree ref="tree" :showCheckbox="showCheckbox" :data="treeData"  size="large"
          @on-select-change="onSelectChange" @on-check-change="onCheckChange"></Tree>
<!--    <spin v-if="loading" size="large" class="tree-spin"></spin>-->
  </div>
</template>

<script>
  import {getallapis,getroleapis,saveroleapis} from "@/api/api";


  export default {
        name: "DataApiTree",
      props: {
        dataUrl: {},
        checkedUrl: {},
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
        async queryData (roleid) {

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

          let checked_ids =[];
          this.roleid = roleid;
          await  getroleapis(roleid).then(res=>{   // tmp_roleid 会实时变化，roleid不能实时变

            if (res.success){
              //console.log(res.data.menu);
              checked_ids = res.data.menu;


            }


          }).catch(err=>{
            alert(err);
          });



          await getallapis().then(res=>{


            if (res.success) {

              //alert(JSON.stringify(res.data.menu));
              //menutree_convert_checked(res.data.menu);

              /*

              js 中数组赋值给新数组，原来的数组的值也跟着变了 ？？？？？？
              因此赋值的是引用。所以原数组值也变了。折腾了半天。

               this.treeData =res.data.menu;

               */

              //this.treeData =res.data.menu;
              const datalist = res.data.list;

              const treedataobj = {};
              const treedata = [];


              res.data.list.map(item => {
                const k=item.group + '组';

                if (k in treedataobj) {
                  treedataobj[k].push(item);
                } else {
                  treedataobj[k] = [item];
                }

              });

              console.log(checked_ids);

              for (let group in treedataobj) {
                var temp_data = {title: group, children: []};
                for (let index in treedataobj[group]) {
                  let item = treedataobj[group][index];


                  if (checked_ids && checked_ids.indexOf(item.ID.toString()  ) !== -1){  // 元素在列表中存在 并且列表不为[]
                    temp_data.children.push({ID: item.ID, title: item.description,checked:true});

                  }else{
                    temp_data.children.push({ID: item.ID, title: item.description});

                  }
                }
                ;

                treedata.push(temp_data);
              }
              ;


              this.treeData = treedata;


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

        submit(){

          console.log(this.treeData);
          console.log(this.$refs.tree.getCheckedNodes());

          const checkArr = this.$refs.tree.getCheckedNodes();

          const check_ids = [];
          checkArr&&checkArr.map(item=>{
            if(item.ID){
              check_ids.push(item.ID);
            }
          });

          saveroleapis({roleid:this.roleid,check_ids:check_ids}).then(res=>{

            if (res.success){

              this.$Message.info({
                content: '保存成功！',
                duration: 2
              });

            } else {
              alert("保存失败！");}
            }
          ).catch(err=>{
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
    }
</script>

<style scoped>

</style>
