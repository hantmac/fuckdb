// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import { Dropdown, DropdownMenu, DropdownItem, Form, FormItem, Select, Option, Button, Row, Col, Input, Radio, RadioGroup, Tabs, TabPane, ButtonGroup, Message} from 'element-ui'
import VueHighlightJS from 'vue-highlight.js';
import go from 'highlight.js/lib/languages/go';
import sql from 'highlight.js/lib/languages/sql';
import 'highlight.js/styles/atom-one-dark.css'


Vue.use(Dropdown);
Vue.use(DropdownMenu);
Vue.use(DropdownItem);
Vue.use(Form);
Vue.use(FormItem);
Vue.use(Select);
Vue.use(Option);
Vue.use(Button);
Vue.use(Row);
Vue.use(Col);
Vue.use(Input)
Vue.use(Radio);
Vue.use(RadioGroup);
Vue.use(Tabs);
Vue.use(TabPane);
Vue.use(ButtonGroup);
Vue.prototype.$message = Message;

import router from './router'

Vue.config.productionTip = false

// Vue.use(Highlight)
Vue.use(VueHighlightJS,{
  // Register only languages that you want
  languages: {
    go,
    sql
  }
})



/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
