<template>

<div>

  <Table  border :columns="columns12" :data="data6" @on-selection-change="showselect">

    <slot name="buttons">

    </slot>

  </Table>

  <div v-if="showPage" style="margin: 10px;overflow: hidden">
    <div :style="'text-align: '+this.pageAlign">
      <Page :total="total" :page-size="pageParam.pageSize" placement="top" :current="pageParam.page"
            @on-change="changePage" @on-page-size-change="changePageSize" :size="tableSize"
            show-total show-sizer show-elevator>
      </Page>
    </div>
  </div>

</div>
</template>

<script>
    export default {
      name: "DataTable",
      props: {
        columns12: [],
        data6:[],
      },
      data () {
        return {
          showPage:false,
          pageAlign:'center',
          total:10,
          tableSize:10,
          pageParam:{
            pageSize:10,
            page:0

          },
          selections:""
        };
      },
      methods:{
        changePage(){},
        changePageSize(){},



      }
    }
</script>

<style scoped>

</style>
