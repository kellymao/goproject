<style scoped>
    .edit-row {
        padding-bottom: 45px;
    }
</style>

<template>
    <div class="main-view main-view-full" style="padding-top: 20px">
        <row style="height: 80px">
            <Button type="text" icon='ios-arrow-back' @click='back'>返 回</Button>
        </row>
        <row class="edit-row" type="flex" justify="center">
            <i-col span="22">

                <formDynamic ref="dynamicForm" :data="options.dynamic" :label-width="options.labelWidth">
                </formDynamic>

                <div :style="{marginLeft: options.labelWidth+'px'}">
                    <i-button type='primary' icon='folder' @click='ok' :loading="loading">
                        保 存
                    </i-button>
                </div>

            </i-col>
        </row>

    </div>
</template>

<script>
    import FormDynamic from '@/component/form/FormDynamic';

    export default {
        props: {},
        data () {
            return {
                options: {},
                action: {},
                data: {},
                loading: false
            };
        },
        computed: {},
        methods: {
            ok () {
                this.$refs.dynamicForm.submit();
            },
            back () {
                this.$router.go(-1);
            }
        },
        mounted () {
            alert('active');
            this.options = this.$route.query.options;
            this.action = this.$route.query.action;
            this.data = this.$route.query.data;
            console.log(this.$route);
            console.log(this.options);
            console.log(this.data);

            // 在created()钩子函数执行的时候DOM 其实并未进行任何渲染，而此时进行DOM操作并无作用，而在created()里使用this.$nextTick()可以等待dom生成以后再来获取dom对象
            this.$nextTick(() => {

                if (JSON.stringify(this.data) !== '{}') {
                    alert(" this.data  is not null");
                    this.$refs['dynamicForm'].setFormData(this.data);

                }
            });
        },
        components: { FormDynamic}
    };
</script>
