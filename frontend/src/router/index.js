import Vue from 'vue'
import Router from 'vue-router'
const home = () => import('@/components/home')
const dump = () => import('@/components/dump')

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: home
    },{
      path: '/dump',
      name: 'dump',
      component: dump
    },{
      path:'*',
      redirect:'/'
    }
  ]
})
