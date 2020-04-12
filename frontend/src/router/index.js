import Vue from 'vue'
import Router from 'vue-router'
const home = () => import('@/components/home')

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: home
    }
  ]
})
