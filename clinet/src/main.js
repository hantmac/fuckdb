// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import VueHighlightJS from 'vue-highlight.js';
import go from 'highlight.js/lib/languages/go';
import 'highlight.js/styles/atom-one-dark.css'


import router from './router'

Vue.config.productionTip = false

Vue.use(ElementUI);
// Vue.use(Highlight)
Vue.use(VueHighlightJS,{
  // Register only languages that you want
  languages: {
    go
  }
})



/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
