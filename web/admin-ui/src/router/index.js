import Vue from 'vue'
import VueRouter from 'vue-router'

import Login from '../views/Login.vue'
import Admin from '../views/Admin.vue'
import Index from '../components/Index.vue'
import ArticleAdd from '../components/Article/ArticleAdd.vue'
import ArticleList from '../components/Article/ArticleList.vue'
import UserList from '../components/User/UserList.vue'
import CateList from '../components/Category/CateList.vue'

Vue.use(VueRouter)


// 路由
const routes = [
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/',
    name: 'admin',
    component: Admin,
    // 子路由
    children:[
      {path:'index', component: Index},
      {path:'artadd', component: ArticleAdd},
      {path:'artadd/:id', component: ArticleAdd, props:true},
      {path:'artlist', component: ArticleList},
      {path:'userlist', component: UserList},
      {path:'catelist', component: CateList},
    ]
  },
]

const router = new VueRouter({
  routes
})

// 前置守卫
router.beforeEach((to, from, next)=>{
  const token = window.sessionStorage.getItem('token')
  if(to.path === '/login') return next()
  if(!token && to.path === '/admin'){
    next('/login')
  } else {
    next()
  }
})
export default router
