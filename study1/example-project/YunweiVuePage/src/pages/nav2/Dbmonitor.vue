<template>
  <iframe
    id="iframeContain"
    name="iframeContain"
    seamless
    scrolling="yes"
    :src="iframeURL"
  >
  </iframe>


</template>

<script>
    export default {
        name: "Dbmonitor",
      data () {
        return {
          message: '有一种遇见，不曾邀约，却心有灵犀；有一种目光，不远不近，却一直在守望；山水相逢，相约陌上，你明媚的笑脸，是我心中倾城的暖，你在，我在，最美好的懂得也在，便是用一朵花开的明媚，见证了整个春天。',
          formItem: {
            textarea: '<h1>许你一场春暖花开</h1></br>'+
              '不曾邀约，却心有灵犀；</br>'+
              '有一种目光，不远不近，却一直在守望;</br>'+
              '山水相逢，相约陌上，</br>'+
              '你明媚的笑脸，是我心中倾城的暖，</br>'+
              '你在，我在，</br>'+
              '最美好的懂得也在，</br>'+
              '便是用一朵花开的明媚，</br>'+
              '见证了整个春天。'
          },
          //iframeURL:"http://172.18.8.178:8000/"
          iframeURL:"http://www.baidu.com/"
        }
      }
    }
</script>

<style scoped>
  .text-copy{
    text-align: center;
    font-size: 14px;
    letter-spacing: 2px;
    margin: 20px;
  }



  /** iframe样式 */
  #iframeContain{
    width: 100%;
    height: 800px;
  }
</style>
