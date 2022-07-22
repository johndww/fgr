import App from './App.vue'
import SelectEvent from './components/SelectEvent.vue'
import {createRouter, createWebHistory} from 'vue-router';
import {createApp} from 'vue'
import SelectUser from "./components/SelectUser.vue";
import Login from "./components/Login.vue";
import ViewEvent from "./components/ViewEvent.vue";
import EditEvent from "./components/EditEvent.vue";
import Home from "./components/Home.vue";
import vue3GoogleLogin from 'vue3-google-login'
import {setupCsrfInterceptor} from "./state/store";
import {setupRouterSecurity} from "./router_security";

const routes = [
    {
        name: 'home',
        path: '/',
        component: Home
    },
    {
        name: 'selectevent',
        path: '/event',
        component: SelectEvent,
    },
    {
        name: 'viewevent',
        path: '/event/:id',
        component: ViewEvent,
    },
    {
        name: 'editevent',
        path: '/event/:id/edit',
        component: EditEvent
    },
    {
        name: 'login',
        path: '/login',
        component: Login
    },
    {
        name: 'selectuser',
        path: '/selectuser',
        component: SelectUser
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

setupRouterSecurity(router)
setupCsrfInterceptor()

const app = createApp(App)

app.use(router)
app.use(vue3GoogleLogin, {
    clientId: '186100627326-iqnh1vj4bbbse1i1qh24p1br61c9hgjh.apps.googleusercontent.com'
})

app.mount('#app')