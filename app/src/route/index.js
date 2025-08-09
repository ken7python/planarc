import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Auth from '../views/Auth.vue'
import Signup from "../views/Signup.vue";
import Login from '../views/Login.vue'

import { CONST } from '@/logic/const.js'
import Daily from "@/views/Daily.vue";
import Flag from "@/views/Flag.vue";

const routes = [
    { path: CONST.homeLink, component: Home},
    { path: '/auth', component: Auth },
    { path: '/signup', component: Signup },
    { path: '/login', component: Login },
    { path: CONST.dailyLink, component: Daily },
    { path: CONST.flagLink, component: Flag },

]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router