import axios from "axios";

export const backendBaseUrl = "https://simplegift.app"
export const v1api = "/api/v1"
export const backendV1Url = backendBaseUrl + v1api

export interface SharedState<T> {
    data: T
    loading: boolean
    error: string
    fetch(): Promise<any>
}

export interface State<T> {
    data: T
    loading: boolean
    error: string
}

let csrfToken = ""

export function setupCsrfInterceptor() {
    axios.interceptors.response.use(undefined, err => {
        const error = err.response;
        if (error.status===401 && error.config &&
            !error.config.__isRetryRequest) {

            console.log("detected 401, fetching new CSRF token and then retrying request")
            return fetchCSRFToken().then(_ => {
                error.config.__isRetryRequest = true;
                error.config.headers['X-CSRF-TOKEN']= csrfToken
                console.log("acquired new CSRF token. retrying request")
                return axios(error.config);
            });
        }
        return Promise.reject(err)
    });
}

export function fetchCSRFToken(): Promise<any> {
    return axios.get(backendV1Url + "/csrf", {
        withCredentials: true
    }).then(resp => {
        csrfToken = resp.data.token
        axios.defaults.headers.common['X-CSRF-TOKEN'] = csrfToken
        console.log("successfully fetched new CSRF token")
    }).catch((err) => {
        console.log("unable to fetch CSRF token")
        return Promise.reject(err)
    })
}