import {createApp} from 'vue/dist/vue.esm-bundler'
import App from './App.vue'
import {router} from './plugins/routes';
import {i18n} from './plugins/translations'

const app = createApp(App)
app.use(router)
app.use(i18n)

app.mount('#app')
