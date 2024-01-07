import { createApp } from 'vue'
import App from './App.vue'
import Antd from 'ant-design-vue'
import qs from 'qs'

import router from '@/router/index'
// import './test/mock'

import './assets/main.css'
import 'ant-design-vue/dist/reset.css'

// context define
const myEnum = {
    0: '日用品',
    1: '食品',
    2: '文件',
    3: '衣物',
    4: '数码产品',
    5: '药品',
    6: '生鲜',
    7: '易碎品',
    8: '其他',
}

export default myEnum;

const app = createApp(App)
app.use(Antd).use(router).use(qs)
app.mount('#app')
