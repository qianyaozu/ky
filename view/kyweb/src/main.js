// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import elementui from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import './assets/css/index.css'
import App from './App'
import router from './router'
import store from './vuex/store'
import Mock from './mockjs/index'
Vue.config.productionTip = false
Vue.use(elementui)

Mock.bootstrap()
router.beforeEach((to, from, next) => {
  next()
})
/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  template: '<App/>',
  components: { App }
})
