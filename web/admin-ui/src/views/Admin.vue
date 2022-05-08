<template>
    <el-container>
      <el-aside  width="250px">
        <!-- 侧边栏组件 -->
        <Aside></Aside>
      </el-aside>
      <el-container>
        <el-header>
          <!-- 头部组件 -->
          <Header></Header>
        </el-header>
        <el-main>
          <!-- 内容组件 -->
          <Main></Main>
        </el-main>
      </el-container>
    </el-container>
</template>

<script>
import Aside from "../components/Aside.vue"
import Header from "../components/Header.vue"
import Main from "../components/Main.vue"

export default {
  components: {
    Aside,
    Header,
    Main
  }
}
</script>

<style scoped>
.el-container, .el-main{
  /* width: 250px; */
  height: 100%;
}
.el-aside, .el-header{
  height: 100%;
  background-color: #282c34;
  padding: 0;
}

</style>