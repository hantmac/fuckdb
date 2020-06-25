<template>
  <div id="app">
    <router-view v-if="isRouterAlive"></router-view>
  </div>
</template>

<script>
export default {
  name: 'App',
  provide(){
    return {
      reload: this.reload
    }
  },
  data () {
    return {
      isRouterAlive:true
    }
  },
  methods: {
      reload() {
        this.isRouterAlive = false
        this.$nextTick(function () {
          this.isRouterAlive = true;
        })
      }
  }
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
