import Vue from 'vue';
import App from './App.vue';
import router from './router';
import layer from 'vue-layer';
import 'vue-layer/lib/vue-layer.css';
import ElementUI from 'element-ui';
import VueI18n from 'vue-i18n';
import Mint from 'mint-ui';
import { messages } from './components/common/i18n';
import 'element-ui/lib/theme-chalk/index.css'; // 默认主题
// import './assets/css/theme-green/index.css'; // 浅绿色主题
import './assets/css/icon.css';
import './components/common/directives';
import './components/common/filter';
import 'babel-polyfill';
import 'mint-ui/lib/style.css';
import './icons'; // icon
import './utils/error-log'; // error log
import store from './store';

Vue.config.productionTip = false;
Vue.prototype.$layer = layer(Vue);
Vue.use(VueI18n);
Vue.use(ElementUI, {
    size: 'small'
});
Vue.use(Mint);
const i18n = new VueI18n({
    locale: 'zh',
    messages
});

//使用钩子函数对路由进行权限跳转
router.beforeEach((to, from, next) => {
    document.title = `${to.meta.title} | 彩票`;
    next();
});

new Vue({
    router,
    i18n,
    store,
    render: h => h(App)
}).$mount('#app');
