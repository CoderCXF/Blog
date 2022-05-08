<template>
  <div class="container">
    <!-- key是为了避免缓存 -->
    <router-view :key="$router.path"></router-view>
  </div>
</template>


<style scoped>
.container{
  height: 100%;
}
</style>