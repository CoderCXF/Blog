import Vue from 'vue'
import App from './App.vue'
import router from './router'

import './plugin/element'
import './plugin/http'
import './assets/css/style.css'

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
