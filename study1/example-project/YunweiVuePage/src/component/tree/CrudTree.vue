<template>
    <div>
        <slot name="treeTop"></slot>
        <div>
            <Button  size="small" shape="circle" type="info" icon="android-create" style="margin-left: 15px;"
                    @click="$router.push(treeOptions.manageUrl);">管理
            </Button>
        </div>

        <div class="data-tree-panel">
            <Tree ref="tree" :showCheckbox="showCheckbox" :data="treeOptions.treedata" :class="{'tree-loading': loading}" size="large"
                  @on-select-change="onSelectChange" @on-check-change="onCheckChange"></Tree>
            <spin v-if="loading" size="large" class="tree-spin"></spin>
        </div>

    </div>
</template>

<script>
    export default {
        name: 'CrudTree',
        data(){
            return {
                total: 0,
                loading: false,
                treeData: [],
                mData: []
            };
        },
        props:{
            treeOptions:{},
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
        methods: {
            onSelectChange(row){
              this.$emit('on-select',row[0].authorityId)
            },
            onCheckChange(){
              alert(2);


            },


        },
        /*
        created:{

        }
        */
    }
</script>

<style scoped>
  .data-tree-panel {
    width: 100%;
    position: relative;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
    font-size: 15px;
  }
  .tree-spin {
    position: absolute;
    left: 40%;
    top: 50%;
  }
  .data-tree-panel .tree-loading {
    opacity: 0.5
  }
  .data-tree-panel .ivu-tree .ivu-tree-title {
    font-size: 15px;
  }
  .data-tree-panel .ivu-tree .ivu-tree-arrow {
    font-size: 15px;
  }
</style>
