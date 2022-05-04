<template>
  <div class="header">
    <el-button type="danger" @click="loginOut">退出</el-button>
  </div>
</template>

<script>
export default {
  data(){
    return{}
  },
  methods:{
    loginOut(){
      window.sessionStorage.clear('token')
      this.$router.push('/login')
    }
  }
}
</script>

<style scoped>
.header{
  width: 100%;
  height: 50px;
  display: flex;
  justify-content: flex-end;
  align-items: center;

}
.header span{
  font-size: 18px;
}
</style>