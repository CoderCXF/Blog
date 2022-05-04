<template>
    <el-container class="container">
      <el-aside width="200px" class="elAside">
        <Aside></Aside>
      </el-aside>
      <el-container>
        <el-header class="elHeader">
          <Header></Header>
        </el-header>
        <el-main>
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


.container{
  height: 100%;
  background-color: #fff;
}

.elAside, .elHeader{
  background-color: #282c34;
}


</style>