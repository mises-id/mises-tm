import starportLibrary from '@starport/vue'
import { createApp } from 'vue'

import App from './App.vue'
import router from './router'
import store from './store'
import 'ant-design-vue/dist/antd.css';

const app = createApp(App)
app.use(store).use(router).use(starportLibrary).mount('#app')