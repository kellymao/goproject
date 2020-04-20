<template>
  <div :class="className" :style="{height:height,width:width}" />
</template>

<script>
import echarts from "echarts";
require("echarts/theme/macarons");
//import { debounce } from "@/utils";

export default {
  props: {
    className: {
      type: String,
      default: "chart"
    },
    width: {
      type: String,
      default: "100%"
    },
    height: {
      type: String,
      default: "500px"
    },
    chartData: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      chart: null,
    };
  },

  mounted() {
    this.initChart();
    //是否需要自适应-加了防抖函数
    //this.__resizeHandler = debounce(() => {
    //  if (this.chart) {
    //    this.chart.resize();
    //  }
    //}, 100);
    //window.addEventListener("resize", this.__resizeHandler);

    // 监听侧边栏的变化以实现自适应缩放
    const sidebarElm = document.getElementsByClassName("sidebar-container")[0];
    //sidebarElm.addEventListener("transitionend", this.sidebarResizeHandler);
  },
  beforeDestroy() {
    if (!this.chart) {
      return;
    }
    //window.removeEventListener("resize", this.__resizeHandler);
    this.chart.dispose();
    this.chart = null;

    const sidebarElm = document.getElementsByClassName("sidebar-container")[0];
    //sidebarElm.removeEventListener("transitionend", this.sidebarResizeHandler);
  },
  methods: {
    initChart() {
      this.chart = echarts.init(this.$el, "macarons");


      this.chart.setOption({
        backgroundColor: 'rgba(221,221,221,0.16)',
        tooltip: {
          trigger: 'item',
          formatter: '{b}'
        },
        legend: {
          show:false,
          top: '2%',
          left: '3%',
          orient: 'vertical',
          data: [{
            name: 'tree1',
            icon: 'rectangle'
          },
            {
              name: 'tree2',
              icon: 'rectangle'
            }],

        },
        series: [{
          type: 'tree',
          name: 'tree1',
          data: data1,
          left:'310',
          right: '50%',
          symbol: 'image://static/images/opacity-img.png',
          initialTreeDepth: 10,
          orient: 'RL',
          label: {
            normal: {
              position: 'left',
              verticalAlign: 'middle',
              backgroundColor: '#fff',
              distance:0,
              borderRadius:10,
              fontFamily:'Arial,NotoSansSC, PingFang SC, Microsoft Yahei, sans-serif',
              formatter:function(params){
                var paramIndex=params.name.indexOf("专业");
                var fristParam=params.name.substring(0,paramIndex);
                var lastParam=params.name.substring(paramIndex);
                if(params.value ==2   ){
                  return '{box1|'+fristParam+'}'+'\n'+'{box11|'+lastParam+'}'
                }else if(params.value ==3){
                  return '{box2|'+fristParam+'}'+'\n'+'{box22|'+lastParam+'}'
                }else if(params.value ==4 ){
                  return '{box3|'+fristParam+'}'+'\n'+'{box33|'+lastParam+'}'
                }else if(params.value ==1 ){
                  return '{box|'+fristParam+'}'+'\n'+'{boxx|'+lastParam+'}'
                }
              },
              rich: {
                box: {
                  height: 30,
                  backgroundColor: 'rgba(255,0,0,0.3)',
                  width:240,
                  color: '#455A74',
                  padding: [0,10],
                  align: 'left',
                  borderColor:'none',
                  borderRadius:[10,10,0,0],
                  position: 'center',
                },
                boxx: {
                  height: 30,
                  backgroundColor: 'rgba(255,0,0,0.3)',
                  width:240,
                  color: '#455A74',
                  padding: [0,10],
                  align: 'left',
                  borderColor:'none',
                  borderRadius:[0,0,10,10],
                  position: 'center',
                },
                box1: {
                  height: 30,
                  width:240,
                  color: '#455A74',
                  padding: [0,10],
                  align: 'left',
                  backgroundColor: 'rgba(244,155,25,0.3)',
                  borderColor:'none',
                  borderRadius:[10,10,0,0]
                },
                box11: {
                  height: 30,
                  width:240,
                  color: '#455A74',
                  padding: [0,10],
                  align: 'left',
                  backgroundColor: 'rgba(244,155,25,0.3)',
                  borderColor:'none',
                  borderRadius:[0,0,10,10]
                },
                box2: {
                  height: 30,
                  width:240,
                  color: '#455A74',
                  padding: [0,10],
                  align: 'left',
                  backgroundColor: 'rgba(213,176,0,0.3)',
                  borderColor:'none',
                  borderRadius:[10,10,0,0]

                },
                box22: {
                  height: 30,
                  width:240,
                  color: '#455A74',
                  padding: [0,10],
                  align: 'left',
                  backgroundColor: 'rgba(213,176,0,0.3)',
                  borderColor:'none',
                  borderRadius:[0,0,10,10]

                },
                box3: {
                  height: 30,
                  width:240,
                  color: '#455A74',
                  padding: [0,10],
                  align: 'left',
                  backgroundColor: 'rgba(69,185,255,0.3)',
                  borderColor:'none',
                  borderRadius:[10,10,0,0]

                },
                box33: {
                  height: 30,
                  width:240,
                  color: '#455A74',
                  padding: [0,10],
                  align: 'left',
                  backgroundColor: 'rgba(69,185,255,0.3)',
                  borderColor:'none',
                  borderRadius:[0,0,10,10]


                }
              }
            }
          },
          expandAndCollapse: true,
          animationDuration: 550,
          animationDurationUpdate: 750
        },{
          type: 'tree',
          name: 'tree2',
          data: data2,
          left: '50%',
          right:'310',
          symbol: 'image://static/images/opacity-img.png',
          initialTreeDepth: 10,
          label: {
            normal: {
              position: 'right',
              verticalAlign: 'middle',
              height: 60,
              backgroundColor: '#fff',
              distance:0,
              borderRadius:10,
              fontFamily:'NotoSansSC, PingFang SC, Microsoft Yahei, sans-serif',
              formatter:function(params){
                var paramIndex=params.name.indexOf("专业");
                var fristParam=params.name.substring(0,paramIndex);
                var lastParam=params.name.substring(paramIndex);
                if(params.value ==2   ){
                  return '{box1|'+fristParam+'}'+'\n'+'{box11|'+lastParam+'}'
                }else if(params.value ==3){
                  return '{box2|'+fristParam+'}'+'\n'+'{box22|'+lastParam+'}'
                }else if(params.value ==4 ){
                  return '{box3|'+fristParam+'}'+'\n'+'{box33|'+lastParam+'}'
                }else if(params.value ==1 ){
                  return '{box|'+fristParam+'}'+'\n'+'{boxx|'+lastParam+'}'
                }
              },
              rich: {
                box: {
                  position: [10,10],
                  height: 30,
                  backgroundColor: 'rgba(255,0,0,0.3)',
                  width:240,
                  color: '#455A74',
                  padding: [0,10],
                  align: 'left',
                  borderColor:'none',
                  borderRadius:[10,10,0,0],
                },
                boxx: {
                  position: [10,10],
                  height: 30,
                  backgroundColor: 'rgba(255,0,0,0.3)',
                  width:240,
                  color: '#455A74',
                  padding: [0,10],
                  align: 'left',
                  borderColor:'none',
                  borderRadius:[0,0,10,10]
                },
                box1: {
                  height: 30,
                  width:240,
                  color: '#455A74',
                  padding: [0,10],
                  align: 'left',
                  backgroundColor: 'rgba(244,155,25,0.3)',
                  borderColor:'none',
                  borderRadius:[10,10,0,0]
                },
                box11: {
                  height: 30,
                  width:240,
                  color: '#455A74',
                  padding: [0,10],
                  align: 'left',
                  backgroundColor: 'rgba(244,155,25,0.3)',
                  borderColor:'none',
                  borderRadius:[0,0,10,10]
                },
                box2: {
                  height: 30,
                  width:240,
                  color: '#455A74',
                  padding: [0,10],
                  align: 'left',
                  backgroundColor: 'rgba(213,176,0,0.3)',
                  borderColor:'none',
                  borderRadius:[10,10,0,0]

                },
                box22: {
                  height: 30,
                  width:240,
                  color: '#455A74',
                  padding: [0,10],
                  align: 'left',
                  backgroundColor: 'rgba(213,176,0,0.3)',
                  borderColor:'none',
                  borderRadius:[0,0,10,10]

                },
                box3: {
                  height: 30,
                  width:240,
                  color: '#455A74',
                  padding: [0,10],
                  align: 'left',
                  backgroundColor: 'rgba(69,185,255,0.3)',
                  borderColor:'none',
                  borderRadius:[10,10,0,0]

                },
                box33: {
                  height: 30,
                  width:240,
                  color: '#455A74',
                  padding: [0,10],
                  align: 'left',
                  backgroundColor: 'rgba(69,185,255,0.3)',
                  borderColor:'none',
                  borderRadius:[0,0,10,10]


                }
              }
            }
          },
          expandAndCollapse: true,
          animationDuration: 550,
          animationDurationUpdate: 750
        }]

      });
    },

    sidebarResizeHandler(e) {
      if (e.propertyName === "width") {
      //  this.__resizeHandler();
      }
    }
  }
};
</script>
