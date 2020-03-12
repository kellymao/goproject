<template>
    <Form ref="fromData" :model="formData"  :label-width="80" >
        <!--此处rowItem 是一个列表-->
        <Row v-for="(rowItem, rowIndex) in data" :key="'rowItem'+rowIndex">
            <!-- item 是一个对象-->
            <i-col v-for="(item, index) in rowItem" :span="item.span" :key="'col'+index">

                <FormItem :key="'formItem'+index"
                          :label="item.label==null?'':item.label"
                          v-show="!item.hidden"
                          :rules="item.rules==null?null:item.rules"
                          :prop="item.name">


                  <template v-if="item.type==='text'">
                    <i-input
                            type="text"
                            v-model="formData[item.name]"
                            :placeholder="item.placeholder==null?'':item.placeholder"
                            :disabled="item.disabled==null?false:item.disabled"
                            :autosize="item.type==='editor'?item.textarea:''" >
                            <!--
                            :icon="item.icon"
                            :autosize="item.textarea"
                            :type="item.type==null?'text':item.type"
                            -->

                    </i-input>
                  </template>
                  <template v-else-if="item.type==='select'">
                    <i-select v-model="formData[item.name]" placeholder="请选择">

<!--                      item.select_options = [-->
<!--                      {title:"普通组",authorityId:"888"},-->
<!--                      {title:"普通1组",authorityId:"999"}-->
<!--                      ]-->

                      <i-option v-for="i in item.select_options" :value="i.authorityId" >{{ i.title }}</i-option>

                    </i-select>

                  </template>

                </FormItem>

            </i-col>
        </Row>

    </Form>

</template>

<script>
    export default {
        name: "FormDynamic",
        props: {
            data: {},
        },
        data(){

            return {

                formData: {},
            }
        },
        methods: {
            // 提交表单
            submit () {
                alert('提交表单数据');
                alert(JSON.stringify(this.formData));


            },
            getData () {
                return this.fromData;
            },
            /**
             * 设置表单数据
             * @param fromData
             */
            setFormData (formdata) {



                //this.formData = formdata;
                // 此处需要动态设置值，静态设置值表单不能自动刷新

                for (let name in formdata) {
                    let val = formdata[name];
                    this.$set(this.formData, name, val);
                }

            },
        },

    }
</script>

<style scoped>

</style>
