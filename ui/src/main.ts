import App from './App.vue'
import SelectEvent from './components/SelectEvent.vue'
import {createRouter, createWebHashHistory} from 'vue-router';
import {createApp} from 'vue'
import SelectUser from "./components/SelectUser.vue";
import ViewEvent from "./components/ViewEvent.vue";
import EditEvent from "./components/EditEvent.vue";
import Home from "./components/Home.vue";
import {isLoggedIn} from "./state/users";

const routes = [
    {
        name: 'home',
        path: '/',
        component: Home
    },
    {
        path: '/event',
        name: 'selectevent',
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
        component: SelectUser
    },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

router.beforeEach((to, from, next) => {
    if (!isLoggedIn() && !(to.name == "login" || to.name == "home")) {
        next({name: 'login'})
    } else {
        next()
    }
})

const app = createApp(App)

app.use(router)

app.mount('#app')