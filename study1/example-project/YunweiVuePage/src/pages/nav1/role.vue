<template>
<div>
  <Form ref="formInline" :model="formInline" :rules="ruleInline" inline>
    <FormItem prop="authorityName">
      <Input type="text" v-model="formInline.authorityName" placeholder="角色名">
        <!--<Icon type="ios-person-outline" slot="prepend"></Icon>-->
      </Input>
    </FormItem>
    <FormItem prop="authorityId">
      <Input type="text" v-model="formInline.authorityId" placeholder="角色id">
        <!--<Icon type="ios-person-outline" slot="prepend"></Icon>-->
      </Input>
    </FormItem>
    <FormItem>
      <Button type="primary" @click="handleSubmit('formInline')">查询</Button>
    </FormItem>
  </Form>


  <div>
    <div class="tooltip">
      <!--增删改按钮-->
        <i-button type="success" shape="circle"  icon="ios-add" @click="tableAddData()">添 加
        </i-button>
        <i-button type="success" shape="circle"  icon="ios-add" @click="tableEditData()"
                  :disabled="isSingle">编 辑
        </i-button>
        <Poptip confirm title="您确认删除选中的内容吗？" @on-ok="tableDelData()" placement="bottom-start" transfer>
          <i-button type="success" shape="circle" icon="ios-trash"
                    :disabled="isMultiple">删 除
          </i-button>
        </Poptip>
      <!--增删改按钮结束-->

      <i-button type='success' shape='circle' icon='ios-add' @click='roleResEdit()' :disabled='isSingle'>角色资源
      </i-button>
    </div>

    <!--角色资源弹出框-->
    <Modal title='角色资源' v-model='editVisible'
           class-name='vertical-center-modal'
           width='400'
           :loading='editLoading'
           @on-ok='roleResEditOk'>

      <DataTree
                ref='resTree'
                style='height:400px;overflow: auto'
                dataUrl='/sys/roleRes/getRoleRes'
                :show-checkbox = 'showcheckbox'
                :lazy='lazy'>

      </DataTree>
    </Modal>

    <div>
  <Table ref="tablerole" border :columns="columns12" :data="data6" @on-selection-change="showselect">
    <template slot-scope="{ row }" slot="name">
      <strong>{{ row.authorityId }}</strong>
    </template>
    <template slot-scope="{ row, index }" slot="action">
      <Button type="primary" size="small" style="margin-right: 5px" @click="show(row)">编辑</Button>

      <Poptip confirm title="您确认删除该角色吗？" @on-ok="remove(row)" placement="left-start" transfer>
      <Button type="error" size="small">删除</Button>
      </Poptip>

    </template>
  </Table>
    </div>

  </div>

  <!-- 编辑弹出框 -->
  <Modal :title="modal_data.title"
         v-model="modal_data.visible"
         class-name="vertical-center-modal"
         class="popup-edit-modal"
         width=580
         @on-ok="ok">
    <Form ref="modlerole" :model="modal_data.roledata" :rules="rulemodlerole" :label-width="80" >
      <FormItem prop="id" label="ID:" >
        <Input type="text" v-model="modal_data.roledata.ID" placeholder="id"  disabled
        >
          <!--<Icon type="ios-person-outline" slot="prepend"></Icon>-->
        </Input>
      </FormItem>

      <FormItem prop="authorityName1" label="角色名:" >
        <Input type="text" v-model="modal_data.roledata.authorityName" placeholder="角色名" class="roledata">
          <!--<Icon type="ios-person-outline" slot="prepend"></Icon>-->
        </Input>
      </FormItem>
      <FormItem prop="authorityId1" label="角色ID:" >
        <Input type="text" v-model="modal_data.roledata.authorityId" placeholder="角色id" class="roledata">
          <!--<Icon type="ios-person-outline" slot="prepend"></Icon>-->
        </Input>
      </FormItem>
      <!--
      <FormItem>
        <Button type="primary" @click="saveSubmit('modlerole')">保存</Button>
      </FormItem>
      -->
    </Form>

    <div slot="footer">
      <Button type="primary" @click="ok">确定</Button>
    </div>
  </Modal>
</div>
</template>
<script>

  import { getroledata,roleadd,roledel } from '@/api/api'
  import DataTree from '@/component/tree/DataTree.vue'


  export default {
    name:'role',
    data () {
      return {
        page: 0,
        pageSize: 0,
        searchInfo:{},
        selections:[],
        columns12: [
          {
            type: 'selection',
            width: 60,
            align: 'center'
          },
          {
            title: '角色id',
            slot: 'name'
          },
          {
            title: '创建时间',
            key: 'CreatedAt'
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
            CreatedAt: '2019-09-08 16:18:31',
            authorityName: '普通用户'
          },
          {
            authorityId: '888',
            CreatedAt: '2019-09-08 16:18:31',
            authorityName: '普通用户'
          }

        ],

        // 查询框及规则
        formInline: {
          authorityName: '',
          authorityId: ''
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

        // 表格后面的编辑弹出框及规则
        rulemodlerole:{


           authorityName1: [
             { required: true, message: 'Please fill in the authorityName.', trigger: 'blur' }
           ],
           authorityId1: [
             { required: true, message: 'Please fill in the authorityId.', trigger: 'blur' },
             //{ type: 'string', min: 6, message: 'The password length cannot be less than 6 bits', trigger: 'blur' }
           ]
        },

        modal_data:{
          visible:false ,
          title: '',
          roledata:{},
        },

        // 角色资源模态框
        editVisible: false,
        editLoading: true,
        showcheckbox:true,
        lazy:false,

      }
    },
    computed:{
      isSingle() {
        return this.selections.length !== 1
      },
      isMultiple(){
        return this.selections.length === 0
      }
    },
    methods: {
      ok(){

      },

      // 每行后面的编辑按钮
      show (row) {
        // this.$Modal.info({
        //   title: 'User Info',
        //   content: `Name：${this.data6[index].name}<br>Age：${this.data6[index].age}<br>Address：${this.data6[index].address}`
        // })
        this.modal_data.title = '修改角色';
        this.modal_data.visible = true;
        this.modal_data.roledata = row;

      },
      // 每行后面的删除按钮
      remove (row) {
        alert(JSON.stringify(row.ID));
      },

      // 查询条件提交按钮
      handleSubmit(name) {
        this.$refs[name].validate((valid) => {
          if (valid) {
            //this.$Message.success('Success!');

            this.searchInfo =  this.formInline;
            this.getTableData(0,0,this.searchInfo);
          } else {
            this.$Message.error('条件输入错误!');
          }
        })


      },


      // 获取表格的数据
      getTableData(page = this.page, pageSize = this.pageSize, searchInfo = {} ){


        //alert(JSON.stringify(searchInfo));
        getroledata({ page, pageSize, ...searchInfo }).then(res=>{


          if (res.success) {

            this.data6 =res.data.list;
            this.total = res.data.total;
            this.page = res.data.page;
            this.pageSize = res.data.pageSize;
          }
        }).catch(err=>{
          alert(err);
        });



      },

      // 表格多选框，选择有变动时触发
      showselect(selection){
        this.selections = selection;
        // alert(selection.length);
        // alert(JSON.stringify(selection));
      },

      // 表格上面的新加和编辑调用此方法
      showEdit(post_api,name,obj){

        const editOptions = {
          labelWidth:80,
          dynamic: [
            [
              {name: 'ID',  hidden: false, label: 'ID',disabled: true,span: 12},
              {name: 'authorityId', type: 'text', span: 12, label: '角色id', rules: {required: true}},
              {name: 'authorityName', type: 'editor', textarea: {minRows: 2, maxRows: 3}, span: 24, label: '角色名',rules: {required: true}},
              // {
              //   name: 'status',
              //   openText: '启用',
              //   closeText: '禁用',
              //   type: 'switch',
              //   span: 24,
              //   label: '角色状态',
              //   value: 1,
              //   trueValue: 1,
              //   falseValue: 0,
              //   rules: {required: true, type: 'number'}
              // }
            ]
          ],
        };

        alert(this.$router.currentRoute.path + '/edit');
        this.$router.push({
          path: this.$router.currentRoute.path + '/edit',
          query: {options: editOptions, action: post_api, data: obj}
        });
      },

      // 表格上面的新加按钮
      tableAddData(){

        this.showEdit(roleadd,'添加',{})
      },

      // 表格上面的编辑按钮
      tableEditData(){
        this.showEdit(roleadd,'编辑',this.selections[0])
      },

      // 表格上面的删除按钮
      tableDelData(){},


      roleResEdit(){
        this.editVisible = true;
      },
      roleResEditOk(){
        this.$refs.resTree.submit();



      },




    },
    mounted() {

      // 页面加载时刷新表格
      this.getTableData()
    },
    components: {DataTree}
  }
</script>

<style scoped>

  .tooltip{
    margin-bottom: 6px ;
  }

  .popup-edit-modal .ivu-modal-wrap {
    z-index: 1051;
    border-radius: 10px;
    -webkit-box-shadow: 4px 4px 5px 0 rgba(0, 0, 0, 0.2);
    box-shadow: 4px 4px 5px 0 rgba(0, 0, 0, 0.2);
  }

  .popup-edit-modal .ivu-modal-body {
    max-height: calc(100vh - 200px);
    overflow: auto;
  }

  .popup-edit-modal .ivu-modal-mask {
    z-index: 1050;
  }

  .popup-edit-modal .ivu-modal {
    top: 0;
  }



</style>
