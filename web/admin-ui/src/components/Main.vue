<template>
  <div>
    <!-- key是为了避免缓存 -->
    <router-view :key="$router.path"></router-view>
  </div>
</template>