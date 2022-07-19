import {User} from "../App.vue";
import axios from "axios";
import {Ref, ref} from "vue";
import {SharedState, State} from "../state/store";

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

    return axios.get("http://localhost/users")
        .then(resp => {
            allUsersState.value.data = resp.data.users
            allUsersState.value.error = ""
        })
        .catch(err => {
            allUsersState.value.error = "Unable to fetch users"
            console.log("error fetching users: " + err)
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

    return axios.post("http://localhost/users/create", {
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
        })
        .finally(() => createUserState.value.loading = false)
}

export function login(userId: string): Promise<any> {
    return axios.post('http://localhost/login/user/' + userId, {}, {
        withCredentials: true
    })
        .then(() => {
            return fetchCurrentUser()
        })
        .catch(err => {
            console.log("unable to login: " + err)
        })
}

export function loginGoogle(credential: string): Promise<any> {
    return axios.post('http://localhost/login/google', {
        token: credential
    }, {
        withCredentials: true
    })
        .then(() => {
            return fetchCurrentUser()
        })
        .catch(err => {
            console.log("unable to login with google: " + err)
        })
}

export function useCurrentUserId() {
    return currentUserId
}

export function isLoggedIn(): boolean {
    return currentUserId.value != ""
}

const currentUserId = ref("")

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

    return axios.get("http://localhost/users/me", {
        withCredentials: true,
    })
        .then(resp => {
            currentUserState.value.data = resp.data
            currentUserId.value = resp.data.id
            currentUserState.value.error = ""
        })
        .catch(err => {
            console.log("unable to fetch current user: " + err)
            currentUserState.value.error = "Unable to fetch current user"
        })
        .finally(() => currentUserState.value.loading = false)
}

export interface GetUsersForEventState extends State<User[]> {}

export function getUsersForEvent(eventId: string, state: Ref<GetUsersForEventState>): Promise<any> {
    console.log("fetching event users")
    state.value.loading = true

    return axios.get("http://localhost/users/event/" + eventId, {withCredentials: true})
        .then(resp => {
            state.value.data = resp.data.users
            state.value.error = ""
        })
        .catch(err => {
            state.value.error = "Unable to fetch users for event"
            console.log("error fetching users for event: " + err)
        })
        .finally(() => state.value.loading = false)
}

export interface LogoutState extends State<string> {}

export function logoutUser(state: Ref<LogoutState>): Promise<any> {
    state.value.loading = true

    return axios.post('http://localhost/logout', {}, {
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
        })
        .finally(() => state.value.loading = false)
}