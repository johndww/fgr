import {Event} from "../App.vue";
import {ref, Ref} from "vue";
import axios from "axios";
import {SharedState, State} from "@/state/store";
import {GiftRequest} from "../App.vue";

export interface MyEventsState extends SharedState<Event[]> {}

const myEventsState: Ref<MyEventsState> = ref({
    data: <Event[]>[],
    loading: false,
    error: "",
    fetch(): Promise<any> {
        return fetchMyEvents()
    }
})

export function useMyEventsState(): Ref<MyEventsState> {
    return myEventsState
}

function fetchMyEvents(): Promise<any> {
    console.log("fetching my events")
    myEventsState.value.loading = true

    return axios.get("http://localhost/events", {withCredentials: true})
        .then(resp => {
            myEventsState.value.data = resp.data.events
            myEventsState.value.error = ""
        })
        .catch(err => {
            myEventsState.value.error = "Unable to fetch my events"
            console.log("error fetching my events: " + err)
        })
        .finally(() => myEventsState.value.loading = false)
}

export interface CreateEventState {
    eventId: string,
    error: string,
    loading: boolean
}

export function createEvent(eventName: string, state: Ref<CreateEventState>) {
    state.value.loading = true

    return axios.post("http://localhost/events/create", {
        name: eventName
    }, {
        withCredentials: true
    })
        .then(resp => {
            state.value.eventId = resp.data.eventId
            state.value.error = ""
        })
        .catch(err => {
            console.log("unable to create event: " + err)
            state.value.error = "Unable to create event"
        })
        .finally(() => state.value.loading = false)
}

export interface GetGiftRequestsState extends State<GiftRequest[]> {}

export function getGiftRequests(eventId: string, state: Ref<GetGiftRequestsState>): Promise<any> {
    console.log("fetching gift requests for event")
    state.value.loading = true

    return axios.get("http://localhost/events/" + eventId + "/gift-requests", {withCredentials: true})
        .then(resp => {
            state.value.data = resp.data.gifts
            state.value.error = ""
        })
        .catch(err => {
            state.value.error = "Unable to fetch gift requests"
            console.log("error fetching gift requests: " + err)
        })
        .finally(() => state.value.loading = false)
}

export interface PersistGiftRequestState extends State<string>{}

export function persistGiftRequest(giftName: string, eventId: string, state: Ref<PersistGiftRequestState>) {
    state.value.loading = true

    return axios.post("http://localhost/events/" + eventId + "/gift-requests/create", {
        name: giftName
    }, {
        withCredentials: true
    })
        .then(resp => {
            state.value.data = resp.data.requestId
            state.value.error = ""
        })
        .catch(err => {
            console.log("unable to persist gift request: " + err)
            state.value.error = "Unable to persist gift request"
        })
        .finally(() => state.value.loading = false)
}

export interface DeleteGiftRequestState extends State<string>{}

export function persistDeleteGiftRequest(eventId: string, giftId: string, state: Ref<DeleteGiftRequestState>) {
    state.value.loading = true

    return axios.delete("http://localhost/events/" + eventId + "/gift-requests/" + giftId + "/delete", {
        withCredentials: true
    })
        .then(() => {
            state.value.error = ""
        })
        .catch(err => {
            console.log("unable to delete gift request: " + err)
            state.value.error = "Unable to delete gift request"
        })
        .finally(() => state.value.loading = false)
}

export interface ReleaseGiftRequestState extends State<string>{}

export function persistReleaseGift(eventId: string, giftId: string, state: Ref<ReleaseGiftRequestState>) {
    state.value.loading = true

    return axios.post("http://localhost/events/" + eventId + "/gift-requests/" + giftId + "/release", {}, {
        withCredentials: true
    })
        .then(() => {
            state.value.error = ""
        })
        .catch(err => {
            console.log("unable to release gift request: " + err)
            state.value.error = "Unable to release gift request"
        })
        .finally(() => state.value.loading = false)
}

export interface ClaimGiftRequestState extends State<string>{}

export function persistClaimGift(eventId: string, giftId: string, byUserId: string, state: Ref<ClaimGiftRequestState>) {
    state.value.loading = true

    return axios.post("http://localhost/events/" + eventId + "/gift-requests/" + giftId + "/claim", {}, {
        withCredentials: true
    })
        .then(() => {
            state.value.error = ""
        })
        .catch(err => {
            console.log("unable to claim gift request: " + err)
            state.value.error = "Unable to claim gift request"
        })
        .finally(() => state.value.loading = false)
}

export interface UpdateEventState extends State<string>{}

export function persistUpdateEvent(eventId: string, name: string, memberEmails: string[], state: Ref<UpdateEventState>): Promise<any> {
    state.value.loading = true

    return axios.post("http://localhost/events/" + eventId + "/update", {
        name: name,
        emails: memberEmails
    }, {
        withCredentials: true
    })
        .then(() => {
            state.value.error = ""
        })
        .catch(err => {
            console.log("unable to update event: " + err)
            state.value.error = "Unable to update event"
        })
        .finally(() => state.value.loading = false)
}
