import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Auth from '../views/Auth.vue'
import Signup from "../views/Signup.vue";
import Login from '../views/Login.vue'

const routes = [
    { path: '/', component: Home},
    { path: '/auth', component: Auth },
    { path: '/signup', component: Signup },
    { path: '/login', component: Login },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router