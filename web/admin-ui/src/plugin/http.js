import Vue from 'vue'
import axios from 'axios'

let URL = 'http://localhost:3000/api/v1'
axios.defaults.baseURL = URL
Vue.prototype.$http = axios

axios.interceptors.request.use(config=>{
  config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
  return config
})

export {URL}