import 'element-plus/es/components/message/style/css'
import 'element-plus/es/components/loading/style/css'
import 'element-plus/es/components/notification/style/css'
import 'element-plus/es/components/message-box/style/css'
import './style/element_visiable.scss'

import '@/styles/index.scss';
import '@/styles/common.scss';
import '@/assets/iconfont/iconfont.css';
import '@/assets/iconfont/iconfont.js';
import '@/styles/style.css';
import { createApp } from 'vue'
// 引入gin-vue-admin前端初始化相关内容
import './core/gin-vue-admin'
// 引入封装的router
import router from '@/router/index'
import '@/permission'
import run from '@/core/gin-vue-admin.js'
import auth from '@/directive/auth'
import { store } from '@/pinia'
import App from './App.vue'
import { initDom } from './utils/positionToCode'
import SvgIcon from './components/svg-icon/svg-icon.vue';
import Components from '@/components';
import Fit2CloudPlus from 'fit2cloud-ui-plus';
import i18n from '@/lang/index';
initDom()
    /**
     * @description 导入加载进度条，防止首屏加载时间过长，用户等待
     *
     * */
import Nprogress from 'nprogress'
import 'nprogress/nprogress.css'
import ElementPlus from 'element-plus';
Nprogress.configure({ showSpinner: false, ease: 'ease', speed: 500 })
Nprogress.start()
/**
 * 无需在这块结束，会在路由中间件中结束此块内容
 * */
const app = createApp(App)
// app.config.productionTip = false

app.component('SvgIcon', SvgIcon);
app.use(ElementPlus)
app.use(Fit2CloudPlus, { locale: i18n.global.messages.value[localStorage.getItem('lang') || 'zh'] });

app
    .use(run)
    .use(store)
    .use(auth)
    .use(router)
    .use(i18n)
    .use(Components)
    .mount('#app')

export default app