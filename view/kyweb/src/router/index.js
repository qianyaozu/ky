import Vue from 'vue'
import Router from 'vue-router'
import Main from '@/views/Main'
import Login from '@/views/Login'
import Demo from '@/components/Demo'
import Script from '@/components/Script'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/Main',
      name: 'Main',
      redirect: '/Demo',
      component: Main,
      children: [
        {
          path: '/Demo',
          name: 'Demo',
          component: Demo,
          hidden: true
        },
        {
          path: '/Script',
          name: 'Script',
          component: Script,
          hidden: true
        }
      ]
    }, {
      path: '/Login',
      name: 'Login',
      component: Login,
      hidden: true
    }
  ]
})
