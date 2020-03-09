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
  import {prase_tree} from "@/common/parse_datatree";

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
      }
    },
    data () {
      return {
        total: 0,
        loading: false,
        treeData: [],
        mData: []
      };
    },
    methods: {
      /**
       * 获取数据
       */
      queryData (param) {

        const data3 = [
          {
            id: 2,
            title: '系统管理',
            expand: true,
            selected: true,
            children: [
              {
                id:7,
                title: '表格',
                expand: false,
                children: [
                  {
                    title: 'leaf 1-1-1',
                    disabled: true
                  },
                  {
                    title: 'leaf 1-1-2'
                  }
                ]
              },
              {
                id:8,
                title: '表单',
                expand: false,
                children: [
                  {
                    title: 'leaf 1-2-1',
                    checked: true
                  },
                  {
                    title: 'leaf 1-2-1'
                  }
                ]
              }
              ,
              {
                id:9,
                title: '用户管理',
                expand: false,
              }
              ,
              {
                id:10,
                title: '角色管理',
                expand: false,
              }
            ]
          }
        ];
        this.treeData = data3;

      },
      onSelectChange (slected) {
          alert(1);
          console.log(slected);
          alert(JSON.stringify(slected));
      },
      onCheckChange (slected) {
        alert(2);
        //console.log(slected);
        console.log(this.$refs.tree.getCheckedNodes()); // 打勾的时候触发

        console.log(this.$refs.tree.getSelectedNodes())
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
        const data = this.$refs.tree.getCheckedNodes();
        let rel = prase_tree(data);

        console.log(rel);
        alert(rel);


      }
    },
    created () {
      if (!this.lazy) {
        this.queryData();
      }
    },
    components: {}
  };
</script>


