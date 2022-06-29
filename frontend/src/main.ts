import { createApp } from 'vue'
import App from './App.vue'
import { pinia } from '@/store'
import { SayHello } from '../wailsjs/go/backend/Config'
import router from "@/router";

let res = await SayHello()
console.log(res)
console.log(`%cmain set res success`, `color:red;font-size:16px;background:transparent`)

const app = createApp(App)

app.use(router)
app.use(pinia)
app.mount('#app')
