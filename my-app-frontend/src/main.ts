import { createApp } from 'vue'
import App from './App.vue'
import router from './router/index'
import store from './store/index'
import ElementUI from 'element-plus'
import 'element-plus/dist/index.css'

createApp(App).use(router).use(store).use(ElementUI).mount('#app')
