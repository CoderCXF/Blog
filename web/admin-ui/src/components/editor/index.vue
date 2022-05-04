<template>
  <div>
    <Editor :init="init" v-model="content"></Editor>
  </div>
</template>

<script>
import Editor from '@tinymce/tinymce-vue'
import tinymce from './tinymce.min.js'
import './icons/default/icons.min.js'
import './themes/silver/theme.min.js'
import './langs/zh-Hans.js'

// 注册插件
// 1.预览
import './plugins/preview/plugin.min'
import './plugins/code/plugin.min'
import './plugins/wordcount/plugin.min.js'
import './plugins/paste/plugin.min'
import './plugins/image/plugin.min'
import './plugins/imagetools/plugin.min'

export default {
  data(){
    return{
      init:{
        branding: false,
        language_url: "./langs/zh-Hans.js",
        language: "zh-Hans",
        height: "500px",
        margin:"0",
        padding: "0",
        plugins: "preview code paste wordcount",
        toolbar:['undo redo | formatselect | alignleft aligncenter alignright alignjustify' , 
                'preview paste code | image  imagetools'
        ],
        
      },
      content:this.value,
    }
  },
  props:{
    value:{
      type: String,
      default:'',
    }
  },
  components:{ Editor },
  watch:{
    value(newV){
      this.content = newV
    },
    content(newV){
      this.$emit('input', newV)
    },
  }

}
</script>


<style>
/* 导入外部样式表 */
@import url('./skins/ui/oxide/skin.min.css');
</style>