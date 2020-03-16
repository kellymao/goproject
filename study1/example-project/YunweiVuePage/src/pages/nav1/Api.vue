<template>
    <div>

      <Form ref="formInline" :model="formInline" :rules="ruleInline" inline>
        <FormItem prop="authorityName">
          <Input type="text" v-model="formInline.path" placeholder="API路径">
            <!--<Icon type="ios-person-outline" slot="prepend"></Icon>-->
          </Input>
        </FormItem>
        <FormItem prop="authorityId">
          <Input type="text" v-model="formInline.group" placeholder="API所属组">
            <!--<Icon type="ios-person-outline" slot="prepend"></Icon>-->
          </Input>
        </FormItem>
        <FormItem>
          <Button type="primary" @click="handleSubmit('formInline')">查询</Button>
        </FormItem>
      </Form>
      <Table ref="tablerole" border :columns="columns12" :data="data6" @on-selection-change="showselect">
        <template slot-scope="{ row }" slot="name">
          <strong>{{ row.path }}</strong>
        </template>
        <template slot-scope="{ row, index }" slot="action">
          <Button type="primary" size="small" style="margin-right: 5px" @click="show(row)">编辑</Button>

<!--          <Poptip confirm title="您确认删除该api吗？" @on-ok="remove(row)" placement="left-start" transfer>-->
            <Button type="error" size="small" @click="remove(row)">删除</Button>
<!--          </Poptip>-->

        </template>
      </Table>

      <pageinfo v-if="showPage" :total="total"  :page="page" :pagesize="pageSize" v-on:queryData="getTableData"></pageinfo>




    </div>
</template>

<script>
  import {getapidata, roleadd} from '@/api/api'
  import common from '@/common/util'
  import DataTree from '@/component/tree/DataTree.vue'
  import pageinfo from '@/component/page/pageinfo.vue'

    export default {
        name: "Api",
        data(){
          return {
            // 查询框及规则
            formInline: {
              path: '',
              group: ''
            },
            ruleInline: {
              // name: [
              //   { required: true, message: 'Please fill in the user name', trigger: 'blur' }
              // ],
              // passwd: [
              //   { required: true, message: 'Please fill in the password.', trigger: 'blur' },
              //   //{ type: 'string', min: 6, message: 'The password length cannot be less than 6 bits', trigger: 'blur' }
              // ]
            },
            searchInfo:{},

            columns12: [
              {
                type: 'selection',
                width: 60,
                align: 'center'
              },
              {
                title: '路径',
                slot: 'name'
              },
              {
                title: '资源描述',
                key: 'description'
              },
              {
                title: '所属组',
                key: 'group'
              },
              {
                title: '操作',
                slot: 'action',
                width: 150,
                align: 'center'
              }
            ],
            data6: [
              {
                authorityId: '888',
                CreatedAt: '2019-09-08 16:18:31',
                authorityName: '普通用户'
              },
              {
                authorityId: '888',
                CreatedAt: '2019-09-08 16:18:31',
                authorityName: '普通用户'
              }

            ],
            showPage: true,
            total:0,
            page: 1,
            pageSize: 10,




          }


        },
      methods:{
        showselect(row){
          console.log(row);
        } ,
        handleSubmit(name) {
          this.$refs[name].validate((valid) => {
            if (valid) {
              //this.$Message.success('Success!');

              this.searchInfo =  this.formInline;
              this.getTableData(this.page,this.pageSize,this.searchInfo);
            } else {
              this.$Message.error('条件输入错误!');
            }
          })


        },
        show(row){
          this.$Modal.info({
            title: 'API信息',
            content: `PATH：${row.path}<br>描述：${row.description}<br>属组：${row.group}`
          })

        },
        remove(row){

          let self = this;
          this.$Modal.confirm({
            //title: 'API信息',
            content: `是否删除此记录`,
            onOk: function(){
              this.$Loading.start();
              let para = {id: row.ID}
              // removeUser(para).then((res)=> {
              //   self.mockTableData();
              // });

              this.$Message.success('删除API成功!' + para.id);

            }
          })

        },
        getTableData(page = this.page, pageSize = this.pageSize, searchInfo = {} ){


          //alert(JSON.stringify(searchInfo));
          //alert(page + '-' + pageSize);
          // getroledata({ page, pageSize, ...searchInfo }).then(res=>{
          //
          //
          //   if (res.success) {
          //
          //     const tmp_data  =res.data.list;
          //
          //     for (let i=0;i<res.data.list.length;i++){
          //       tmp_data[i].CreatedAt = common.formatDate.utcformat(tmp_data[i].CreatedAt, "yyyy-MM-dd HH:mm:ss")
          //     }
          //
          //     this.data6 =tmp_data;
          //     this.total = res.data.total;
          //     this.page = res.data.page;
          //     this.pageSize = res.data.pageSize;
          //   }
          // }).catch(err=>{
          //   alert(err);
          // });

            getapidata({page,pageSize, ...searchInfo}).then(res =>{

               if(res.success){

                 this.data6 = res.data.list;
                 this.total = res.data.total;
                 this.page = res.data.page;
                 this.pageSize = res.data.pageSize ;

                 console.log(this.data6);


               }


            }).catch(err=>{
              alert(err);
            })

        },







      },
      mounted() {

        // 页面加载时刷新表格
        this.getTableData()
      },
      components: {
        pageinfo,
        DataTree
      }
    }
</script>

<style scoped>

</style>
