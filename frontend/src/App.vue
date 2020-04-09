<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
* {
  margin: 0;
  padding: 0;
}

html, body, #app{
  height: 100%;
}
.el-tabs__content{
  display: none
}
</style>
