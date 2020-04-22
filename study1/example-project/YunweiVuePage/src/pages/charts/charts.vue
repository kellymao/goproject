<template>
    <!-- <tree :chartData="treeData" /> -->
    <el-section id="chart" class="chart"  ref="companyStyle">
    <p style="margin: 6px; auto;text-align: center;">
        <img style="height:80px;" :src="host_comment"/>
    </p>
    <div id="main_orgStructure" style="width:100%; height:500px;">
        有内容
    </div>
    </el-section>
</template>

<script>
    import tree from './tree';
    import echarts from 'echarts'
    import host_comment from '@/assets/host-comment.png'

    
    export default {
        data() {
            // return {
            //
            //   treeData: {
            //     label: {
            //       backgroundColor: '#F4F4F4',
            //       borderRadius: [0, 0, 5, 5],
            //       formatter: [
            //         '{first|综合授信额度}',
            //         '{second|(CR20190912000013)\n获批金额：100\n币种：人民币}',
            //       ].join('\n'),
            //       rich: {
            //         first: {
            //           backgroundColor: '#078E34',
            //           color: '#fff',
            //           align: 'center',
            //           width: 135,
            //           height: 30,
            //           borderRadius: [5, 5, 0, 0],
            //         },
            //         second: {
            //           color: '#888',
            //           align: 'center',
            //           lineHeight: 17,
            //         },
            //       }
            //     },
            //     children: [
            //       {
            //         label: {
            //           formatter: [
            //             '{first|渠道额度}',
            //           ].join('\n'),
            //           rich: {
            //             first: {
            //               backgroundColor: '#3AC082',
            //               color: '#fff',
            //               align: 'center',
            //               width: 135,
            //               height: 65,
            //               borderRadius: 5,
            //             },
            //           }
            //         },
            //         children: [{
            //           label: {
            //             formatter: [
            //               '{first|保理额度}',
            //             ].join('\n'),
            //             rich: {
            //               first: {
            //                 backgroundColor: '#3AC082',
            //                 color: '#fff',
            //                 align: 'center',
            //                 width: 135,
            //                 height: 65,
            //                 borderRadius: 5,
            //               },
            //             }
            //           },
            //           children: [{
            //             label: {
            //               backgroundColor: '#F4F4F4',
            //               borderRadius: [0, 0, 5, 5],
            //               formatter: [
            //                 '{first|反向保理}',
            //                 '{second|(CR20190912000013)\n获批金额：100\n币种：人民币}',
            //               ].join('\n'),
            //               rich: {
            //                 first: {
            //                   backgroundColor: '#078E34',
            //                   color: '#fff',
            //                   align: 'center',
            //                   width: 135,
            //                   height: 30,
            //                   borderRadius: [5, 5, 0, 0],
            //                 },
            //                 second: {
            //                   color: '#888',
            //                   align: 'center',
            //                   lineHeight: 17,
            //                 },
            //               }
            //             },
            //           }]
            //         }]
            //       },
            //       {
            //         label: {
            //           formatter: [
            //             '{first|担保/(乐)集团/其他额度}',
            //           ].join('\n'),
            //           rich: {
            //             first: {
            //               backgroundColor: '#3AC082',
            //               color: '#fff',
            //               align: 'center',
            //               width: 135,
            //               height: 65,
            //               borderRadius: 5,
            //             },
            //           }
            //         },
            //         children: [{
            //           label: {
            //             formatter: [
            //               '{first|保理额度}',
            //             ].join('\n'),
            //             rich: {
            //               first: {
            //                 backgroundColor: '#3AC082',
            //                 color: '#fff',
            //                 align: 'center',
            //                 width: 135,
            //                 height: 65,
            //                 borderRadius: 5,
            //               },
            //             }
            //           },
            //           children: [{
            //             label: {
            //               backgroundColor: '#F4F4F4',
            //               borderRadius: [0, 0, 5, 5],
            //               formatter: [
            //                 '{first|正向保理}',
            //                 '{second|(CR20190912000013)\n获批金额：100\n币种：人民币}',
            //               ].join('\n'),
            //               rich: {
            //                 first: {
            //                   backgroundColor: '#B8D87E',
            //                   color: '#fff',
            //                   align: 'center',
            //                   width: 135,
            //                   height: 30,
            //                   borderRadius: [5, 5, 0, 0],
            //                 },
            //                 second: {
            //                   color: '#888',
            //                   align: 'center',
            //                   lineHeight: 17,
            //                 },
            //               }
            //             },
            //           }]
            //         },
            //           {
            //             label: {
            //               formatter: [
            //                 '{first|租赁额度}',
            //               ].join('\n'),
            //               rich: {
            //                 first: {
            //                   backgroundColor: '#3AC082',
            //                   color: '#fff',
            //                   align: 'center',
            //                   width: 135,
            //                   height: 65,
            //                   borderRadius: 5,
            //                 },
            //               }
            //             },
            //             children: [
            //               {
            //                 label: {
            //                   backgroundColor: '#F4F4F4',
            //                   borderRadius: [0, 0, 5, 5],
            //                   formatter: [
            //                     '{first|车辆租赁}',
            //                     '{second|(CR20190912000013)\n获批金额：100\n币种：人民币}',
            //                   ].join('\n'),
            //                   rich: {
            //                     first: {
            //                       backgroundColor: '#FF6C6A',
            //                       color: '#fff',
            //                       align: 'center',
            //                       width: 135,
            //                       height: 30,
            //                       borderRadius: [5, 5, 0, 0],
            //                     },
            //                     second: {
            //                       color: '#888',
            //                       align: 'center',
            //                       lineHeight: 17,
            //                     },
            //                   }
            //                 },
            //               },
            //             ]
            //           }]
            //       }
            //     ]
            //   }
            //
            //
            //
            //
            //
            // }
            return {
                chartorg: null,
                host_comment: host_comment,
            }

        },

        mounted: function () {

            this.chartorg = echarts.init(document.getElementById('main_orgStructure'));

            var option = {
                title: {
                    text: 'lover3 集群结构图',
                    subtext: '写VIP:192.168.0.247 \n \n 读VIP:192.168.0.153 \n \n 主控机:192.168.1.28-MMM'
                    ,
                    x: '123px;',
                    textAlign: 'center',
                    padding: 1,
                    //itemGap:10,
                    subtextStyle: {//副标题文本样式{"color": "#aaa"}
                        color: 'blue',
                        //fontFamily: 'Arial, Verdana, sans...',
                        fontSize: 15,
                        fontStyle: 'normal',
                        fontWeight: 'normal',
                    }
                },


                tooltip: {
                    show: false,
                    trigger: 'item',
                    formatter: "{b}: {c}"
                },
                toolbox: {
                    show: true,
                    feature: {
                        mark: {show: true},
                        dataView: {show: true, readOnly: false},
                        restore: {show: true},
                        saveAsImage: {show: true}
                    }
                },
                calculable: false,
                // legend: {
                //     x: 60,
                //     y: 30,
                //     itemWidth: 39,
                //     itemHeight: 30,
                //     orient: 'Vertical',
                //     icon: 'square',
                //     borderColor: 'rgba(178,34,34,0.8)',
                //     borderWidth: 0,
                //     textStyle: {
                //         color: 'blue',
                //     },
                //
                //
                //     data:[{
                //           name:'树图',
                //         textStyle: {
                //             color: 'black',
                //             fontSize: 16,
                //             fontWeight:  'bolder'
                //         },
                //         icon:  'image://static/static/M.png',
                //
                //
                //
                //
                //     }]
                // },

                series: [
                    {
                        name: '树图',
                        type: 'tree',
                        orient: 'vertical',  // vertical horizontal
                        rootLocation: {x: '50%', y: '15%'}, // 根节点位置  {x: 'center',y: 10}
                        nodePadding: 20,
                        layerPadding: 40,
                        symbol: 'rectangle',
                        borderColor: 'black',
                        //edgeShape:'curve',
                        itemStyle: {
                            normal: {
                                color: '#fff',//节点背景色
                                label: {
                                    show: true,
                                    position: 'bottom',
                                    rotate: -15,
                                    textStyle: {
                                        color: 'black',
                                        fontSize: 13,
                                        fontWeight: 'bolder'
                                    }
                                },
                                lineStyle: {
                                    color: '#000',
                                    width: 1,
                                    type: 'polyline', // 'curve'|'broken'|'solid'|'dotted'|'dashed'
                                    curveness: 0.5
                                }
                            },
                            emphasis: {
                                label: {
                                    show: false
                                }
                            }
                        },
                        data: [
                            {
                                name: '192.168.1.61:3306',
                                value: 6,
                                symbolSize: [76, 60],
                                symbol: 'image://static/static/M.png',
                                itemStyle: {
                                    normal: {
                                        borderWidth: 2,
                                        borderColor: 'black'
                                    },
                                    label: {
                                        show: true,
                                        position: 'left',
                                        rotate: -15,
                                        textStyle: {
                                            color: 'black',
                                            fontSize: 12,
                                            fontWeight: 'bolder'
                                        }
                                    },
                                    lineStyle: {
                                        opacity: 0
                                    },
                                    opacity: 0
                                },
                                label: {
                                    show: true,
                                    position: 'left',
                                    rotate: -15,
                                    textStyle: {
                                        color: 'black',
                                        fontSize: 12,
                                        fontWeight: 'bolder'
                                    }
                                },

                                children: [
                                    {
                                        name: '192.168.1.60:3306',
                                        value: 6,
                                        symbolSize: [76, 60],
                                        symbol: 'image://static/static/M.png',
                                        itemStyle: {
                                            normal: {
                                                borderWidth: 2,
                                                borderColor: 'black'
                                            }
                                        },
                                        lineStyle: {
                                            opacity: 0,
                                            color: '#0000ff',
                                            borderColor: '#0000ff'
                                        },
                                    },
                                    {
                                        name: '192.168.0.109:3306',
                                        value: 6,
                                        symbolSize: [76, 60],
                                        symbol: 'image://static/static/s-r.png',
                                        itemStyle: {
                                            normal: {
                                                borderWidth: 2,
                                                borderColor: 'black'
                                            }
                                        },
                                        lineStyle: {
                                            opacity: 0,
                                            color: '#0000ff',
                                            borderColor: '#0000ff'
                                        },
                                    },
                                    {
                                        name: '192.168.1.74:3306',
                                        value: 6,
                                        symbolSize: [76, 60],
                                        symbol: 'image://static/static/s-r.png',
                                        itemStyle: {
                                            normal: {
                                                borderWidth: 2,
                                                borderColor: 'black'
                                            }
                                        },
                                        lineStyle: {
                                            opacity: 0,
                                            color: '#0000ff',
                                            borderColor: '#0000ff'
                                        },
                                    },
                                    {
                                        name: '192.168.0.159:3306',
                                        value: 6,
                                        symbolSize: [76, 60],
                                        symbol: 'image://static/static/s.png',
                                        itemStyle: {
                                            normal: {
                                                borderWidth: 2,
                                                borderColor: 'black'
                                            }
                                        },
                                        lineStyle: {
                                            opacity: 0,
                                            color: '#0000ff',
                                            borderColor: '#0000ff'
                                        },
                                        children: [
                                            {
                                                name: '192.168.18.23:3306',
                                                value: 6,
                                                symbolSize: [76, 60],
                                                symbol: 'image://static/static/bg.png',
                                                itemStyle: {
                                                    normal: {
                                                        borderWidth: 2,
                                                        borderColor: 'black'
                                                    }
                                                },
                                                lineStyle: {
                                                    opacity: 0,
                                                    color: '#0000ff',
                                                    borderColor: '#0000ff'
                                                },
                                            },
                                            {
                                                name: '192.168.18.24:3306',
                                                value: 6,
                                                symbolSize: [76, 60],
                                                symbol: 'image://static/static/bg.png',
                                                itemStyle: {
                                                    normal: {
                                                        borderWidth: 2,
                                                        borderColor: 'black'
                                                    }
                                                },
                                                lineStyle: {
                                                    opacity: 0,
                                                    color: '#0000ff',
                                                    borderColor: '#0000ff'
                                                },
                                            },
                                        ]

                                    },
                                    {
                                        name: '192.168.0.65:3306',
                                        value: 6,
                                        symbolSize: [76, 60],
                                        symbol: 'image://static/static/s.png',
                                        itemStyle: {
                                            normal: {
                                                borderWidth: 2,
                                                borderColor: 'black'
                                            }
                                        },
                                        lineStyle: {
                                            opacity: 0,
                                            color: '#0000ff',
                                            borderColor: '#0000ff'
                                        },
                                    },
                                    {
                                        name: '192.168.0.161:3306',
                                        value: 6,
                                        symbolSize: [76, 60],
                                        symbol: 'image://static/static/s.png',
                                        itemStyle: {
                                            normal: {
                                                borderWidth: 2,
                                                borderColor: 'black'
                                            }
                                        },
                                        lineStyle: {
                                            opacity: 0,
                                            color: '#0000ff',
                                            borderColor: '#0000ff'
                                        },
                                    },
                                    {
                                        name: '192.168.0.196:3306',
                                        value: 6,
                                        symbolSize: [76, 60],
                                        symbol: 'image://static/static/s.png',
                                        itemStyle: {
                                            normal: {
                                                borderWidth: 2,
                                                borderColor: 'black'
                                            }
                                        },
                                        lineStyle: {
                                            opacity: 0,
                                            color: '#0000ff',
                                            borderColor: '#0000ff'
                                        },
                                    },
                                ]
                            }
                        ]
                    }
                ]
            };
            // 为echarts对象加载数据
            this.chartorg.setOption(option);


        }
        ,

        components: {
            tree,
        },
        methods: {

            autosize() {
                window.addEventListener('resize', function () {
                    this.chartorg.resize();
                })
            }
        },
        computed: {
            newName() {

                let inputStyWidth
                let me = this;
                this.$nextTick(function () {
                    inputStyWidth = me.$refs.companyStyle.$el.clientWidth + 'px';
                });
                return inputStyWidth;
            }
        },
        watch: {
            newName(val) {
                alert(val);

                this.autosize();
            }
        }
    }

</script>

<style scoped>
    .chart {
        width: 100%;

    }

</style>
