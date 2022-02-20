import {createApp} from 'vue'
import App from './App.vue'
import {router} from './plugins/routes';
import {i18n} from './plugins/translations'
import {pinia} from "./plugins/pinia";

const app = createApp(App)
app.use(router)
app.use(i18n)
app.use(pinia)

app.mount('#app')
