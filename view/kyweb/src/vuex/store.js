import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)

const state = {
  Token: localStorage.getItem('Token'),
  UserName: localStorage.getItem('UserName'),
  MineName: localStorage.getItem('MineName'),
  CurrentPath: localStorage.getItem('CurrentPath')
}

export default new Vuex.Store({
  state
})
export function SetStore (key, value) {
  if (key === 'UserName') {
    state.UserName = value
  } else if (key === 'MineName') {
    state.MineName = value
  } else if (key === 'Token') {
    state.Token = value
  } else if (key === 'CurrentPath') {
    state.CurrentPath = value
  }

  localStorage.setItem(key, value)
}
export function GetStore (key) {
  if (key === 'UserName') {
    return state.UserName
  } else if (key === 'MineName') {
    return state.MineName
  } else if (key === 'Token') {
    return state.Token
  } else if (key === 'CurrentPath') {
    return state.CurrentPath
  }
}
