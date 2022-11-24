import {haveCheckedIfLoggedIn, isLoggedIn, useCurrentUserState} from "./state/users";
import {RouteLocationNormalized, Router} from "vue-router";

const defaultLoggedInLocation = 'selectevent';
const defaultNotLoggedInLocation = 'home';

const isDev = import.meta.env.DEV

function adminRedirect(to: RouteLocationNormalized, from: RouteLocationNormalized): string | null {
    if ((to.name == "selectuser" && !useCurrentUserState().value.data?.admin) || (to.name === "devlogin" && !isDev)) {
        console.log("user is not an admin, invalid page location")
        return defaultLoggedInLocation
    }
    return null
}

export function setupRouterSecurity(router: Router) {
    router.beforeEach((to, from, next) => {
        // i'm not a fan of these chains, but not entirely sure how to simplify without another library
        // goal here is to have ordered redirect checks, and bail the first time a check requires a redirect
        initLoginStateAndRedirect(to, from)
            .then(redirectName => {
                if (redirectName) {
                    console.log("redirecting to: " + redirectName)
                    next({name: redirectName})
                    return
                }

                let nextRedirect = notLoggedInRedirect(to, from)
                if (nextRedirect) {
                    console.log("redirecting to: " + nextRedirect)
                    next({name: nextRedirect})
                    return
                }

                nextRedirect = adminRedirect(to, from)
                if (nextRedirect) {
                    console.log("redirecting to: " + nextRedirect)
                    next({name: nextRedirect})
                    return
                }

                next()
            })
    })
}

function notLoggedInRedirect(to: RouteLocationNormalized, from: RouteLocationNormalized): string | null {
    if (!isLoggedIn() && !(to.name == "login" || to.name === "selectuser" || to.name === "home" || to.name === "devlogin")) {
        console.log("not logged in or trying to login")
        return defaultNotLoggedInLocation
    }
    return null
}

function initLoginStateAndRedirect(to: RouteLocationNormalized, from: RouteLocationNormalized): Promise<string | null> {
    if (!isLoggedIn() && !haveCheckedIfLoggedIn()) {
        console.log("checking if the user already has a session")

        return useCurrentUserState().value.fetch()
            .then(() => {
                console.log("detected valid user session and is logged in")
                return true
            }).catch(() => {
                console.log("user does not have a valid session")
                return false
            }).then((isLoggedIn) => {
                if (isLoggedIn) {
                    if (to.name === "home") {
                        console.log("user was attempting to go home, but is already logged in")
                        return defaultLoggedInLocation
                    }
                }
                return null
            })
    }
    return Promise.resolve(null)
}