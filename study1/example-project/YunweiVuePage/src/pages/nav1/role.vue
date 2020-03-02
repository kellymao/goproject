<template>
<div>
  <Form ref="formInline" :model="formInline" :rules="ruleInline" inline>
    <FormItem prop="rolename">
      <Input type="text" v-model="formInline.rolename" placeholder="角色名">
        <!--<Icon type="ios-person-outline" slot="prepend"></Icon>-->
      </Input>
    </FormItem>
    <FormItem prop="roleid">
      <Input type="text" v-model="formInline.roleid" placeholder="角色id">
        <!--<Icon type="ios-person-outline" slot="prepend"></Icon>-->
      </Input>
    </FormItem>
    <FormItem>
      <Button type="primary" @click="handleSubmit('formInline')">查询</Button>
    </FormItem>
  </Form>

  <Table border :columns="columns12" :data="data6">
    <template slot-scope="{ row }" slot="name">
      <strong>{{ row.authorityId }}</strong>
    </template>
    <template slot-scope="{ row, index }" slot="action">
      <Button type="primary" size="small" style="margin-right: 5px" @click="show(index)">编辑</Button>
      <Button type="error" size="small" @click="remove(index)">删除</Button>
    </template>
  </Table>

</div>
</template>
<script>

  import { getroledata } from '@/api/api'


  export default {
    name:'role',
    data () {
      return {
        searchInfo:{},
        columns12: [
          {
            title: '角色id',
            slot: 'name'
          },
          {
            title: '创建时间',
            key: 'created_at'
          },
          {
            title: '角色名',
            key: 'authorityName'
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
            created_at: '2019-09-08 16:18:31',
            authorityName: '普通用户'
          },
          {
            authorityId: '888',
            created_at: '2019-09-08 16:18:31',
            authorityName: '普通用户'
          }

        ],
        formInline: {
          rolename: '',
          roleid: ''
        },
        ruleInline: {
          rolename: [
            { required: true, message: 'Please fill in the user name', trigger: 'blur' }
          ],
          roleid: [
            { required: true, message: 'Please fill in the password.', trigger: 'blur' },
            //{ type: 'string', min: 6, message: 'The password length cannot be less than 6 bits', trigger: 'blur' }
          ]
        }
      }
    },
    methods: {
      show (index) {
        this.$Modal.info({
          title: 'User Info',
          content: `Name：${this.data6[index].name}<br>Age：${this.data6[index].age}<br>Address：${this.data6[index].address}`
        })
      },
      remove (index) {
        this.data6.splice(index, 1);
      },

      handleSubmit(name) {
        this.$refs[name].validate((valid) => {
          if (valid) {
            //this.$Message.success('Success!');

            this.searchInfo =  this.formInline;
          } else {
            this.$Message.error('条件输入错误!');
          }
        })


      },


      getTableData(page = this.page, pageSize = this.pageSize ){


        getroledata({ page, pageSize, ...this.searchInfo }).then(res=>{


          if (res.success) {

            this.data6 =res.data.data;
            this.total = res.data.total;
            this.page = res.data.page;
            this.pageSize = res.data.pageSize;
          }
        }).catch(err=>{
          alert(err);
        });



      }



    },
    mounted() {
      this.getTableData()
    }
  }
</script>
