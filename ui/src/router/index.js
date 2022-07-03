import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'

const routes = [
  { path: '/',name: 'Home',component: Home },
  { path: '/login',name: 'Login',component: () => import('../views/Login.vue')},
  { path: '/register',name:'Resister',component: () => import('../views/Resister.vue')},
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router