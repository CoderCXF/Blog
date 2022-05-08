<template>
  <mavon-editor v-model="artInfo.content" :ishljs="true" ref=md @imgAdd="$imgAdd" @imgDel="$imgDel" class="editor"/>
</template>

<script>
export default {
  data(){
    return {
      artInfo:{
        content:'',
      }
    }
  },
  methods:{
    $imgAdd(pos, $file){
      var formdata = new FormData();
      formdata.append('File', $file);
      this.img_file[pos] = $file;
      this.$http({
        url: '/upload',
        method: 'post',
        data: formdata,
        headers: { 'Content-Type': 'multipart/form-data' },
      }).then((res) => {
        console.log(res)
        let _res = res.data;
        // 第二步.将返回的url替换到文本原位置![...](0) -> ![...](url)
        this.$refs.md.$img2Url(pos, _res.url);
      })
    },
    $imgDel(pos, $file){
      // console.log(pos, $file)
    }
  }
}
</script>

<style scoped>
.editor{
  /* margin-top: 50px; */
  width: 100%;
  height: 100%;
}
</style>