<template>
    <!--
    <Table border :columns="columns5" :data="data5"></Table>
    -->

    <Row>
        <!--左边树-->
        <i-col  :xs="8" :sm="4" :md="4" :lg="3"  >

            <CrudTree ref="crudTree" :treeOptions="treeOptions" @on-data-loaded="onTreeDataLoaded" @on-select="onTreeChange">
                <template slot="treeTop">
                    <h1>准备开始搞！！！</h1>
                </template>
            </CrudTree>
        </i-col>

        <!--左边表格-->
        <i-col  :xs="16" :sm="20" :md="20" :lg="21"  >

        </i-col>
    </Row>
</template>
<script>
    import CrudTree from '@/component/tree/CrudTree.vue';

    /**
     * 树参数
     */
    const treeOptions = {
        dataUrl: '/sys/role/list',
        categoryField: 'roleId',
        manageUrl: '/role',
        showToolbar: true
    };


    export default {
        data () {

            return {
                treeOptions
            }
        },

        components: {
            CrudTree
        },
        methods:{

            onTreeDataLoaded(){

            },
            onTreeChange(){

            }

        }
            /*
            return {
                columns5: [
                    {
                        title: '日期',
                        key: 'date',
                        sortable: true
                    },
                    {
                        title: '姓名',
                        key: 'name'
                    },
                    {
                        title: '年龄',
                        key: 'age',
                        sortable: true
                    },
                    {
                        title: '地址',
                        key: 'address'
                    }
                ],
                data5: [
                    {
                        name: '王小明',
                        age: 18,
                        address: '北京市朝阳区芍药居',
                        date: '2016-10-03'
                    },
                    {
                        name: '张小刚',
                        age: 25,
                        address: '北京市海淀区西二旗',
                        date: '2016-10-01'
                    },
                    {
                        name: '李小红',
                        age: 30,
                        address: '上海市浦东新区世纪大道',
                        date: '2016-10-02'
                    },
                    {
                        name: '周小伟',
                        age: 26,
                        address: '深圳市南山区深南大道',
                        date: '2016-10-04'
                    }
                ]
            }
        }
        */
    }


</script>
