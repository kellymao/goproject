<template>
    <!--
    <Table border :columns="columns5" :data="data5"></Table>
    -->
<div>
    <Row>
        <!--左边树-->
        <i-col  :xs="8" :sm="4" :md="4" :lg="3"  >

            <CrudTree ref="crudTree" :treeOptions="treeOptions" @on-data-loaded="onTreeDataLoaded" @on-select="onTreeChange">
                <template slot="treeTop">

                </template>
            </CrudTree>
        </i-col>

        <!--右边表格-->
        <i-col  :xs="16" :sm="20" :md="20" :lg="21"  >
          <Table :context="self" :data="tableData" :columns="tableColumns" stripe border></Table>

        </i-col>


    </Row>



<!--  <Modal :title="modal_data.title"-->
<!--         v-model="modal_data.visible"-->
<!--         class-name="vertical-center-modal"-->
<!--         class="popup-edit-modal"-->
<!--         width=580-->
<!--         @on-ok="ok">-->
<!--    <Form ref="modlerole" :model="modal_data.roledata" :rules="rulemodlerole" :label-width="80" >-->
<!--      <FormItem prop="id" label="ID:" >-->
<!--        <Input type="text" v-model="modal_data.roledata.ID" placeholder="id"  disabled-->
<!--        >-->
<!--          &lt;!&ndash;<Icon type="ios-person-outline" slot="prepend"></Icon>&ndash;&gt;-->
<!--        </Input>-->
<!--      </FormItem>-->

<!--      <FormItem prop="authorityName1" label="角色名:" >-->
<!--        <Input type="text" v-model="modal_data.roledata.authorityName" placeholder="角色名" class="roledata">-->
<!--          &lt;!&ndash;<Icon type="ios-person-outline" slot="prepend"></Icon>&ndash;&gt;-->
<!--        </Input>-->
<!--      </FormItem>-->
<!--      <FormItem prop="authorityId1" label="角色ID:" >-->
<!--        <Input type="text" v-model="modal_data.roledata.authorityId" placeholder="角色id" class="roledata">-->
<!--          &lt;!&ndash;<Icon type="ios-person-outline" slot="prepend"></Icon>&ndash;&gt;-->
<!--        </Input>-->
<!--      </FormItem>-->
<!--      &lt;!&ndash;-->
<!--      <FormItem>-->
<!--        <Button type="primary" @click="saveSubmit('modlerole')">保存</Button>-->
<!--      </FormItem>-->
<!--      &ndash;&gt;-->
<!--    </Form>-->

<!--    <div slot="footer">-->
<!--      <Button type="primary" @click="ok">确定</Button>-->
<!--    </div>-->
<!--  </Modal>-->

  <Modal :title="modal_data.title"
         v-model="modal_data.visible"
         class-name="vertical-center-modal"
         class="popup-edit-modal"
         width=580
         @on-ok="ok">
  <formDynamic ref="dynamicForm" :data="options.dynamic" :label-width="options.labelWidth">
  </formDynamic>

    <div slot="footer">
      <Button type="primary" @click="ok">确定</Button>
    </div>
  </Modal>
</div>
</template>
<script>
    import CrudTree from '@/component/tree/CrudTree.vue';
    import FormDynamic from '@/component/form/FormDynamic';

    import {getroledata, getuserdata} from '@/api/api'


    /**
     * 树参数
     */
    const treeOptions = {
        dataUrl: '/sys/role/list',
        categoryField: 'roleId',
        manageUrl: '/role',
        showToolbar: true,
        treedata:[],
    };

    const editOptions = {
      labelWidth:80,
      dynamic: [
        [
          {name: 'ID',  type: 'text',hidden: false, label: 'ID',disabled: true,span: 12},
          {name: 'username', type: 'text', span: 12, label: '用户名', rules: {required: true}},
          {name: 'nickname', type: 'text', span: 12, label: '别名', rules: {required: true}},
          // {name: 'authorityId', type: 'text', span: 12, label: '所属组id', rules: {required: true}},
          // {name: 'authorityname', type: 'text', textarea: {minRows: 2, maxRows: 3}, span: 24, label: '所属组',rules: {required: true}},

          {name: 'authorityId', type: 'select', select_options:'',span: 24, label: '所属组',rules: {required: true}}
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


    export default {
        data () {

            return {
                treeOptions,
              self: this,
              table_origindata:[],
              tableData: [],
              tableColumns: [
                {
                  type: 'selection',
                  width: 60,
                  align: 'center'
                },
                {
                  title: '名字',
                  key: 'username',
                  align: 'center',
                  sortable: true

                },
                {
                  title: '所属组',
                  align: 'center',
                  key: 'authorityname',
                  width: 100
                },
                {
                  title: '别名',
                  align: 'center',
                  key: 'nickname'
                },
                {
                  title: '创建时间',
                  align: 'center',
                  key: 'created'
                },
                // {
                //   title: '性别',
                //   key: 'sex',
                //   align: 'center',
                //   render: function(row, column, index){
                //     return row.sex == 0 ? '男' : '女';
                //   }
                // },
                {
                  title: '操作',
                  key: 'action',
                  width: 150,
                  align: 'center',

                  render: (h, params) => {
                    return h('div', [
                      h('Button', {
                        props: {
                          type: 'warning',
                          size: 'small'
                        },
                        style: {
                          marginRight: '5px'
                        },
                        on: {
                          click: () => {
                            this.show(params.row);
                          }
                        }
                      }, '编辑'),
                      h('Button', {
                        props: {
                          type: 'success',
                          size: 'small'
                        },
                        on: {
                          click: () => {
                            this.remove(params.row)
                          }
                        }
                      }, '删除')
                    ]);
                  }
                }
              ],
              options:editOptions,
              modal_data:{
                  title:"编辑用户",
                  visible:false,
              }
            }
        },

        components: {
            CrudTree,FormDynamic
        },
        methods:{

            onTreeDataLoaded(){

            },
            onTreeChange(id){
              //alert(id);
              this.tableData =  this.table_origindata.filter(item=>{
                return item.authorityId === id;
              });
            },
           querytreedata(){
            getroledata({}).then(res=>{


              if (res.success) {


                const tmp_data  =[];

                for (let i=0;i<res.data.list.length;i++){
                  let obj = {};
                  obj.title = res.data.list[i].authorityName;
                  obj.authorityId = res.data.list[i].authorityId;

                  tmp_data[i] = obj;
                }

                this.treeOptions.treedata  = tmp_data;

                for( let i=0; i<this.options.dynamic[0].length; i++){

                  if (this.options.dynamic[0][i].name === 'authorityId'){
                    this.options.dynamic[0][i].select_options = tmp_data;
                  }
                }

                console.log(this.treeOptions.treedata);

                console.log(this.options);
                alert(JSON.stringify(this.options));

              }
            }).catch(err=>{
              alert(err);
            });

          },
          querytabledata(){

            getuserdata({}).then(res=>{


              if (res.success) {


                const tmp_data  =[];
                console.log(res.data.list);

                for (let i=0;i<res.data.list.length;i++){
                  let obj = {};
                  obj.ID = res.data.list[i].ID
                  obj.username = res.data.list[i].userName;
                  obj.authorityname = res.data.list[i].authority.authorityName;
                  obj.authorityId= res.data.list[i].authority.authorityId;
                  obj.nickname = res.data.list[i].nickName;

                  obj.created = res.data.list[i].CreatedAt;

                  tmp_data[i] = obj;
                }

                this.tableData  = this.table_origindata = tmp_data;


              }
            }).catch(err=>{
              alert(err);
            });

          },

          // 点击编辑俺就触发
          show(row){
              console.log(row);

            this.$refs['dynamicForm'].setFormData(row);

            this.modal_data.visible = true;
          },

          // 编辑后点击确定后触发
          ok(){
            this.$refs['dynamicForm'].submit();
            this.modal_data.visible = false;

          }

        },
        created() {
          this.querytreedata();
          this.querytabledata();
        },
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
