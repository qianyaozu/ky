import Vue from 'vue'
import Router from 'vue-router'
import Main from '@/views/Main'
import Login from '@/views/Login'
import Demo from '@/components/Demo'
import WorkPlace from '@/components/WorkPlace'
import FrameSet from '@/components/FrameSet'
import PillarSet from '@/components/PillarSet'
import DipSet from '@/components/DipSet'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/Main',
      name: 'Main',
      redirect: '/Main',
      component: Main,
      children: [
        {
          path: '/Demo',
          name: 'Demo',
          component: Demo,
          hidden: true
        },
        {
          path: '/WorkPlace',
          name: 'WorkPlace',
          component: WorkPlace,
          hidden: true
        },
        {
          path: '/FrameSet',
          name: 'FrameSet',
          component: FrameSet,
          hidden: true
        },
        {
          path: '/PillarSet',
          name: 'PillarSet',
          component: PillarSet,
          hidden: true
        },
        {
          path: '/DipSet',
          name: 'DipSet',
          component: DipSet,
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
