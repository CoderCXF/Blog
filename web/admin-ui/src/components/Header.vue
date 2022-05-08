<template>
  <div class="container">
    <!-- <span><i class="el-icon-link"></i></span> -->
    <el-tooltip class="item" effect="dark" content="回到首页" :hide-after="500" placement="top-start">
      <div class="el-icon-link iconitem" @click="onLinkToIndex"></div>
    </el-tooltip>
    <el-tooltip class="item" effect="dark" content="后台布局设置" :hide-after="500" placement="top-start" >
      <div class="el-icon-setting iconitem"></div>
    </el-tooltip>
    <el-tooltip class="item" effect="dark" content="评论信息" :hide-after="500" placement="top-start" >
      <div class="el-icon-bell iconitem"></div>
    </el-tooltip>
    <el-dropdown class="item" @command="handleCommand">
      <span><img src="../assets/avatar.jpg" alt=""></span>
      <el-dropdown-menu slot="dropdown">
        <el-dropdown-item command="a">个人中心</el-dropdown-item>
        <el-dropdown-item command="b">退出登录</el-dropdown-item>
      </el-dropdown-menu>
    </el-dropdown>
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
    },
    handleCommand(command){
      if(command === "a"){
        this.$router.push('/myinfo').catch(error => error)
      } else if(command === 'b') {
        this.loginOut()
      }
    },
    onLinkToIndex(){
      this.$router.push('/index').catch(error => error)
    }
  }
}
</script>

<style scoped>
.container{
  width: 100%;
  height: 60px;
  background-color: #fff;
  display: flex;
  justify-content: flex-end;
  align-items: center;
  box-shadow: 10px 1px 15px 1px  gray;
}

.container img{
  width: 25px;
  height: 25px;
  border-radius: 18px;
}

.container .item{
  padding: 0 20px;
}

.container .iconitem:hover{
  color:blue;
}

.container .item .iconitem{
  font-size: 16px;
  color: black;
}
</style>