import App from './App.vue'
import SelectEvent from './components/SelectEvent.vue'
import {createRouter, createWebHistory} from 'vue-router';
import {createApp} from 'vue'
import SelectUser from "./components/SelectUser.vue";
import Login from "./components/Login.vue";
import ViewEvent from "./components/ViewEvent.vue";
import EditEvent from "./components/EditEvent.vue";
import Home from "./components/Home.vue";
import {haveCheckedIfLoggedIn, isLoggedIn, useCurrentUserId, useCurrentUserState} from "./state/users";
import vue3GoogleLogin from 'vue3-google-login'
import {setupCsrfInterceptor} from "./state/store";

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

router.beforeEach(async (to, from, next) => {
    if (!isLoggedIn() && !haveCheckedIfLoggedIn()) {
        console.log("checking if the user already has a session")

        const isLoggedIn = await useCurrentUserState().value.fetch().then(() => {
            console.log("detected valid user session and is logged in")
            return Promise.resolve(true)
        }).catch(() => {
            console.log("user does not have a valid session")
            return Promise.resolve(false)
        })

        if (isLoggedIn) {
            if (to.name === "home") {
                console.log("user was attempting to go home, but is already logged in. redirecting to select event")
                next({name: 'selectevent'})
            } else {
                next()
            }
            return
        }
    }

    if (!isLoggedIn() && !(to.name == "login" || to.name === "selectuser" || to.name === "home")) {
        console.log("not logged in or trying to login, redirecting back to home")
        next({name: 'home'})
    } else {
        next()
    }
})

setupCsrfInterceptor()

const app = createApp(App)

app.use(router)
app.use(vue3GoogleLogin, {
    clientId: '186100627326-iqnh1vj4bbbse1i1qh24p1br61c9hgjh.apps.googleusercontent.com'
})

app.mount('#app')