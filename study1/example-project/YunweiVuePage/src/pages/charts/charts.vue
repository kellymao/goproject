<template>
    <tree :chartData="treeData" />
</template>

<script>
    import tree from './tree';

    
    export default{
        data (){
            return {

              treeData: {
                label: {
                  backgroundColor: '#F4F4F4',
                  borderRadius: [0, 0, 5, 5],
                  formatter: [
                    '{first|综合授信额度}',
                    '{second|(CR20190912000013)\n获批金额：100\n币种：人民币}',
                  ].join('\n'),
                  rich: {
                    first: {
                      backgroundColor: '#078E34',
                      color: '#fff',
                      align: 'center',
                      width: 135,
                      height: 30,
                      borderRadius: [5, 5, 0, 0],
                    },
                    second: {
                      color: '#888',
                      align: 'center',
                      lineHeight: 17,
                    },
                  }
                },
                children: [
                  {
                    label: {
                      formatter: [
                        '{first|渠道额度}',
                      ].join('\n'),
                      rich: {
                        first: {
                          backgroundColor: '#3AC082',
                          color: '#fff',
                          align: 'center',
                          width: 135,
                          height: 65,
                          borderRadius: 5,
                        },
                      }
                    },
                    children: [{
                      label: {
                        formatter: [
                          '{first|保理额度}',
                        ].join('\n'),
                        rich: {
                          first: {
                            backgroundColor: '#3AC082',
                            color: '#fff',
                            align: 'center',
                            width: 135,
                            height: 65,
                            borderRadius: 5,
                          },
                        }
                      },
                      children: [{
                        label: {
                          backgroundColor: '#F4F4F4',
                          borderRadius: [0, 0, 5, 5],
                          formatter: [
                            '{first|反向保理}',
                            '{second|(CR20190912000013)\n获批金额：100\n币种：人民币}',
                          ].join('\n'),
                          rich: {
                            first: {
                              backgroundColor: '#078E34',
                              color: '#fff',
                              align: 'center',
                              width: 135,
                              height: 30,
                              borderRadius: [5, 5, 0, 0],
                            },
                            second: {
                              color: '#888',
                              align: 'center',
                              lineHeight: 17,
                            },
                          }
                        },
                      }]
                    }]
                  },
                  {
                    label: {
                      formatter: [
                        '{first|担保/(乐)集团/其他额度}',
                      ].join('\n'),
                      rich: {
                        first: {
                          backgroundColor: '#3AC082',
                          color: '#fff',
                          align: 'center',
                          width: 135,
                          height: 65,
                          borderRadius: 5,
                        },
                      }
                    },
                    children: [{
                      label: {
                        formatter: [
                          '{first|保理额度}',
                        ].join('\n'),
                        rich: {
                          first: {
                            backgroundColor: '#3AC082',
                            color: '#fff',
                            align: 'center',
                            width: 135,
                            height: 65,
                            borderRadius: 5,
                          },
                        }
                      },
                      children: [{
                        label: {
                          backgroundColor: '#F4F4F4',
                          borderRadius: [0, 0, 5, 5],
                          formatter: [
                            '{first|正向保理}',
                            '{second|(CR20190912000013)\n获批金额：100\n币种：人民币}',
                          ].join('\n'),
                          rich: {
                            first: {
                              backgroundColor: '#B8D87E',
                              color: '#fff',
                              align: 'center',
                              width: 135,
                              height: 30,
                              borderRadius: [5, 5, 0, 0],
                            },
                            second: {
                              color: '#888',
                              align: 'center',
                              lineHeight: 17,
                            },
                          }
                        },
                      }]
                    },
                      {
                        label: {
                          formatter: [
                            '{first|租赁额度}',
                          ].join('\n'),
                          rich: {
                            first: {
                              backgroundColor: '#3AC082',
                              color: '#fff',
                              align: 'center',
                              width: 135,
                              height: 65,
                              borderRadius: 5,
                            },
                          }
                        },
                        children: [
                          {
                            label: {
                              backgroundColor: '#F4F4F4',
                              borderRadius: [0, 0, 5, 5],
                              formatter: [
                                '{first|车辆租赁}',
                                '{second|(CR20190912000013)\n获批金额：100\n币种：人民币}',
                              ].join('\n'),
                              rich: {
                                first: {
                                  backgroundColor: '#FF6C6A',
                                  color: '#fff',
                                  align: 'center',
                                  width: 135,
                                  height: 30,
                                  borderRadius: [5, 5, 0, 0],
                                },
                                second: {
                                  color: '#888',
                                  align: 'center',
                                  lineHeight: 17,
                                },
                              }
                            },
                          },
                        ]
                      }]
                  }
                ]
              }





            }
        },
      components: {
        tree,
      }

    }
</script>

<style scoped>

</style>
