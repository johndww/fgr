import {User} from "../App.vue";
import axios from "axios";
import {Ref, ref} from "vue";
import {fetchCSRFToken, SharedState, State} from "../state/store";

export interface AllUsersState extends SharedState<User[]> {}

const allUsersState: Ref<AllUsersState> = ref({
    data: [],
    loading: false,
    error: "",
    fetch(): Promise<any> {
        return fetchAllUsers()
    }
})

export function useAllUsers(): Ref<AllUsersState> {
    return allUsersState
}

function fetchAllUsers(): Promise<any> {
    console.log("fetching users")
    allUsersState.value.loading = true

    return axios.get("http://localhost/api/v1/users", {
        withCredentials: true,
    })
        .then(resp => {
            allUsersState.value.data = resp.data.users
            allUsersState.value.error = ""
        })
        .catch(err => {
            allUsersState.value.error = "Unable to fetch users"
            console.log("error fetching users: " + err)
            return Promise.reject(err)
        })
        .finally(() => allUsersState.value.loading = false)
}

export interface CreateUserState {
    userId: string,
    error: string,
    loading: boolean
}

export function createUser(userName: string, createUserState: Ref<CreateUserState>): Promise<any> {
    createUserState.value.loading = true

    return axios.post("http://localhost/api/v1/users/create", {
        name: userName
    }, {
        withCredentials: true
    })
        .then(resp => {
            createUserState.value.userId = resp.data.userId
            createUserState.value.error = ""
        })
        .catch(err => {
            console.log("unable to create user: " + err)
            createUserState.value.error = "Unable to create user"
            return Promise.reject(err)
        })
        .finally(() => createUserState.value.loading = false)
}

export function login(userId: string): Promise<any> {
    return axios.post('http://localhost/api/v1/login/user/' + userId, {}, {
        withCredentials: true
    })
        .then(() => {
            return fetchCSRFToken()
                .then(() => {
                    return fetchCurrentUser().catch((err) => {
                        console.log("unable to fetch user after logging in: " + err.message)
                        return Promise.reject(err)
                    })
                }).catch((err) => {
                    console.log("unable to fetch CSRF token immediately after logging in. will try again: " + err.message)
                    return Promise.resolve()
                })
        })
        .catch(err => {
            console.log("unable to login: " + err.message)
            return Promise.reject(err)
        })
}

export function loginGoogle(credential: string): Promise<any> {
    return axios.post('http://localhost/api/v1/login/google', {
        token: credential
    }, {
        withCredentials: true
    })
        .then(() => {
            return fetchCSRFToken()
                .then(() => {
                    return fetchCurrentUser().catch((err) => {
                        console.log("unable to fetch user after logging in: " + err.message)
                        return Promise.reject(err)
                    })
                }).catch((err) => {
                    console.log("unable to fetch CSRF token immediately after logging in. will try again: " + err.message)
                    return Promise.resolve()
                })
        })
        .catch(err => {
            console.log("unable to login with google: " + err)
            return Promise.reject(err)
        })
}

export function useCurrentUserId() {
    return currentUserId
}

export function isLoggedIn(): boolean {
    return !!currentUserId.value;
}

export function haveCheckedIfLoggedIn(): boolean {
    return currentUserId.value === null || currentUserId.value == ""
}

const currentUserId: Ref<string | undefined | null> = ref(undefined)

export interface CurrentUserState extends SharedState<User | null> {}

const currentUserState: Ref<CurrentUserState> = ref({
    data: null,
    loading: false,
    error: "",
    fetch(): Promise<any> {
        return fetchCurrentUser()
    }
})

export function useCurrentUserState(): Ref<CurrentUserState> {
    return currentUserState
}

function fetchCurrentUser(): Promise<any> {
    console.log("fetching current user")
    currentUserState.value.loading = true

    return axios.get("http://localhost/api/v1/users/me", {
        withCredentials: true,
    })
        .then(resp => {
            currentUserState.value.data = resp.data
            currentUserId.value = resp.data.id
            currentUserState.value.error = ""
        })
        .catch(err => {
            // this API is used to check if the user is logged in or not. if we get a 403, it's not really an error, no need pollute the logs
            if (err.response.status != 403) {
                console.log("unable to fetch current user: " + err)
            }
            currentUserState.value.error = "Unable to fetch current user"
            currentUserId.value = null // undefined -> null, shows that we have attempted evaluation
            return Promise.reject(err)
        })
        .finally(() => currentUserState.value.loading = false)
}

export interface GetUsersForEventState extends State<User[]> {}

export function getUsersForEvent(eventId: string, state: Ref<GetUsersForEventState>): Promise<any> {
    console.log("fetching event users")
    state.value.loading = true

    return axios.get("http://localhost/api/v1/users/event/" + eventId, {withCredentials: true})
        .then(resp => {
            state.value.data = resp.data.users
            state.value.error = ""
        })
        .catch(err => {
            state.value.error = "Unable to fetch users for event"
            console.log("error fetching users for event: " + err)
            return Promise.reject(err)
        })
        .finally(() => state.value.loading = false)
}

export interface LogoutState extends State<string> {}

export function logoutUser(state: Ref<LogoutState>): Promise<any> {
    state.value.loading = true

    return axios.post('http://localhost/api/v1/logout', {}, {
        withCredentials: true
    })
        .then(() => {
            state.value.error = ""
            currentUserState.value.data = null
            currentUserId.value = ""
        })
        .catch(err => {
            state.value.error = "Unable to logout"
            console.log("unable to login: " + err)
            return Promise.reject(err)
        })
        .finally(() => state.value.loading = false)
}