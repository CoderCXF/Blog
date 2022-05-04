<template>
  <div class="elAside">
    <div class="logo">
          <span>my blog</span>
    </div>
    <el-menu class="hidden-xs-only"
      default-active="1" 
      background-color="#282c34"
      text-color="#fff"
      active-text-color="#ffd04b"
      router
      @select="goToPage">
      <!-- 仪表盘 -->
      <el-menu-item index="index">
        <i class="el-icon-s-data"></i>
        <span>仪表盘</span>
      </el-menu-item>

      <!-- 导航栏一：文章管理 -->
      <el-submenu index="2">
          <template slot="title">
            <i class="el-icon-document"></i>
            <span>文章管理</span>
        </template>
        <el-menu-item index="artadd">
          <i class="el-icon-document"></i>
          <span slot="title">写文章</span>
        </el-menu-item>
        <el-menu-item index="artlist">
          <i class="el-icon-document"></i>
          <span slot="title">文章列表</span>
        </el-menu-item>
      </el-submenu>
      <el-menu-item index="catelist">
        <i class="el-icon-document"></i>
        <span slot="title">分类列表</span>
      </el-menu-item>
      <el-menu-item index="userlist">
        <i class="el-icon-user"></i>
        <span slot="title">用户列表</span>
      </el-menu-item>
    </el-menu>
    </div>
</template>

<script>

export default {
  data(){
    return{}
  },
  methods:{
    goToPage(index){
      this.$router.push('/'+index).catch(err=>err)
    }
  }

}
</script>
<style scoped>

.logo{
  height: 60px;
  background-color:antiquewhite;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 17px;
}
</style>