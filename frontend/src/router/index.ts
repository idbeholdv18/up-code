import { createRouter, createWebHistory } from 'vue-router'
import HomeView from './HomeView.vue'
import RegisterView from './RegisterView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', component: HomeView },
    { path: '/register', component: RegisterView },
  ],
})

export default router
