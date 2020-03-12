<template>
  <div style="margin: 10px;overflow: hidden; height: 30px;">
    <div :style="'text-align: '+this.pageAlign">
      <Page :total="total" :page-size="pagesize" placement="top" :current="page"
            @on-change="changePage" @on-page-size-change="changePageSize" :size="tableSize" prev-text="上一页" next-text="下一页"
            show-total show-sizer show-elevator>
      </Page>
      <!--
      total : 总记录条数
      page-size: 每页的条数
      current： 当前的页码
      size: 可选择small 或默认

      on-change: 页码改变的回调，返回改变后的页码
      on-page-size-change	切换每页条数时的回调，返回切换后的每页条数

      -->
    </div>
  </div>
</template>

<script>
    export default {
        name: "pageinfo",
      props: {
        total:{
          type: Number,
          default: 0
        },
        page:{
          type: Number,
          default: 1
        },
        pagesize:{
          type: Number,
          default: 10
        } ,
      },
        data () {
        return {
          pageAlign:'center',
          tableSize:"small",
          pageParam:{
            pageSize:10,
            page:0

          },
        };
      },
        methods:{
          changePage (page) {
            this.pageParam.page = page;
            this.$parent.$data.page = page;
            //this.$emit("getTableData",this.pageParam.page,this.pageParam.pageSize);
            this.$emit("queryData");
          },
          changePageSize (pageSize) {
            this.pageParam.pageSize = pageSize;

            this.$parent.$data.page = 1;
            this.$parent.$data.pageSize = pageSize;
            //alert(pageSize);
            this.$emit("queryData");
          },



      }
    }
</script>

<style scoped>

</style>
