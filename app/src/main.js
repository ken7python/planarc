import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './route/index.js'
import { vLongpress } from '@nanogiants/vue3-longpress'

import '@event-calendar/core/index.css';

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'


const app = createApp(App)
app.directive('use-longpress', vLongpress)
app.use(router).mount('#app')
app.use(ElementPlus)
//
// createApp(App).use(router).mount('#app')
