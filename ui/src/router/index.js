import { createRouter, createWebHistory } from 'vue-router'
// import Home from '../views/Home.vue'

const routes = [
  {path: '/',name: 'Home', component:() => import('../views/Home.vue')},
  {path: '/login',name:'Login', component: () => import('../views/auth/Login.vue')},
  {path: '/register',name:'Resister', component: () => import('../views/auth/Register.vue')},
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
