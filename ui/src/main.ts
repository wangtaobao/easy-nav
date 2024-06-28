import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import 'nprogress/nprogress.css'
import router from './router'
// import { DatePicker } from 'ant-design-vue';

const app = createApp(App)

app.use(router)

app.mount('#app')