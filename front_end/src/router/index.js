import { createRouter, createWebHistory } from 'vue-router'
import LoginPage from '../authorization/LoginPage.vue'
import SignUp from '../authorization/SignUp.vue'
import MainPage from '../main/MainPage.vue'
import HomeVie from '../views/HomeView.vue'

const routes = [
  {
    path: '/login',
    name: 'login',
    component: LoginPage
  },
  {
    path: '/signup',
    name: 'SignUp',
    component: SignUp
  },
  {
    path: "/",
    name: "main",
    component: MainPage
  },
  {
    path:"/test",
    name: "test",
    component: HomeVie
  }
  
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
