<template>
  <div class="container">
    <div class="logo">
      <span>blog</span>
    </div>
    <el-menu class="navmenu"
      default-active="index" 
      background-color="#282c34"
      text-color="#fff"
      active-text-color="#ff0000"
      router
      @select="goToPage">
      <!-- 仪表盘 -->
      <el-menu-item index="index">
        <i class="el-icon-s-data"></i>
        <span>仪表盘</span>
      </el-menu-item>

      <!-- 导航栏一：文章管理 -->
      <el-submenu index="idnexArt">
          <template slot="title">
            <i class="el-icon-document"></i>
            <span>文章管理</span>
        </template>
        <el-menu-item index="startart">
          <i class="el-icon-document"></i>
          <span slot="title">开始写文章</span>
        </el-menu-item>
        <!-- 打算删除这个界面 -->
        <el-menu-item index="artadd">
          <i class="el-icon-document"></i>
          <span slot="title">写文章</span>
        </el-menu-item>
        <el-menu-item index="drafts">
          <i class="el-icon-document"></i>
          <span slot="title">草稿箱</span>
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
        <el-menu-item index="taglist">
          <i class="el-icon-document"></i>
          <span slot="title">标签管理</span>
        </el-menu-item>
      <!-- 用户中心 -->
      <el-menu-item index="userlist">
        <i class="el-icon-s-check"></i>
        <span slot="title">用户管理</span>
      </el-menu-item>
      <!-- 个人中心 -->
      <el-menu-item index="myinfo">
        <i class="el-icon-user"></i>
        <span slot="title">个人中心</span>
      </el-menu-item>
    </el-menu>
    </div>
</template>

<script>

export default {
  methods:{
    goToPage(index){
      // 选中事件，进行路由
      this.$router.push('/'+index).catch(err=>err)
    }
  }

}
</script>
<style scoped>
.container{
  height: 100%;
}
.logo{
  width: 100%;
  height: 60px;
  background-color:#002140;
  color: red;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 17px;
}
</style>