import Vue from 'vue'
import Router from 'vue-router'
import Main from '@/views/Main'
import Login from '@/views/Login'
import SystemConfig from '@/components/SystemConfig'
import WorkPlace from '@/components/WorkPlace'
import FrameSet from '@/components/FrameSet'
import PillarSet from '@/components/PillarSet'
import DipSet from '@/components/DipSet'
import FrameReal from '@/components/FrameReal'
import PillarReal from '@/components/PillarReal'
import PillarHis from '@/components/PillarHis'
import FrameHis from '@/components/FrameHis'
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
          path: '/SystemConfig',
          name: 'SystemConfig',
          component: SystemConfig,
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
        },
        {
          path: '/FrameReal',
          name: 'FrameReal',
          component: FrameReal,
          hidden: true
        },
        {
          path: '/PillarReal',
          name: 'PillarReal',
          component: PillarReal,
          hidden: true
        },
        {
          path: '/PillarHis',
          name: 'PillarHis',
          component: PillarHis,
          hidden: true
        },
        {
          path: '/FrameHis',
          name: 'FrameHis',
          component: FrameHis,
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
