<template>
    <div>
        <slot name="treeTop"></slot>
        <div>
            <Button  size="small" shape="circle" type="ghost" icon="android-create"
                    @click="$router.push(treeOptions.manageUrl);">管理
            </Button>
        </div>

        <div class="data-tree-panel">
            <Tree ref="tree" :showCheckbox="showCheckbox" :data="treeData" :class="{'tree-loading': loading}" size="large"
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
            onSelectChange(){

            },
            onCheckChange(){

            }

        },
        created:{

        }
    }
</script>

<style scoped>

</style>