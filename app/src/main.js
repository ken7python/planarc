import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './route/index.js'
import { vLongpress } from '@nanogiants/vue3-longpress'

import '@event-calendar/core/index.css';

const app = createApp(App)
app.directive('use-longpress', vLongpress)
app.use(router).mount('#app')
//
// createApp(App).use(router).mount('#app')
