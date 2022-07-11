import App from './App.vue'
import SelectEvent from './components/SelectEvent.vue'
import {createRouter, createWebHashHistory} from 'vue-router';
import {createApp} from 'vue'
import SelectUser from "./components/SelectUser.vue";
import ViewEvent from "./components/ViewEvent.vue";
import {isLoggedIn} from "./state/store";
import EditEvent from "./components/EditEvent.vue";
import Home from "./components/Home.vue";

const routes = [
    { name: 'Home', path: '/', component: Home },
    { path: '/selectevent', component: SelectEvent },
    { name: 'Login', path: '/login', component: SelectUser },
    { path: '/event', component: ViewEvent },
    { path: '/editevent', component: EditEvent },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

router.beforeEach((to, from, next) => {
    if (!isLoggedIn() && !(to.name == "Login" || to.name == "Home")) {
        next({name: 'Login'})
    } else {
        next()
    }
})

const app = createApp(App)

app.use(router)

app.mount('#app')