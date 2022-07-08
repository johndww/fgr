import {ref} from "vue";

let sessionUserId = ref("")

export function getSessionUserId() {
    return sessionUserId
}

export function getCurrentUser() {
    return allUsers.find((user) => user.id === sessionUserId.value)
}

export function login(userId: string) {
    sessionUserId.value = userId
}

export function createUser(userName: string) {
    let userId = Math.floor(Math.random() * 1000).toString();
    allUsers = [...allUsers, {id: userId, name: userName, gifts: [], eventIds: []}]
    login(userId)
}

export function createEvent(eventName: string) {
    let eventId = Math.floor(Math.random() * 1000).toString()
    allEvents = [...allEvents, {id: eventId, name: eventName, ownerUserId: sessionUserId.value}]
    return eventId
}

export function isLoggedIn(): boolean {
    return sessionUserId.value != ""
}

export function getAllUsers() {
    return allUsers
}

export function getEvent(id: any) {
    return allEvents.find((event) => event.id === id)
}

export function getMyEvents() {
    return allEvents.filter((event) => getCurrentUser()?.eventIds.includes(event.id))
}

export function getEventUsers(eventId: any) {
    return allUsers.filter((user) => user.eventIds.includes(eventId))
}

let allEvents = [
    {
        id: "1",
        name: "2020 Wright's Christmas",
        ownerUserId: "1",
    },
    {
        id: "2",
        name: "2021 Tapa's Thanksgiving",
        ownerUserId: "2",
    },
]

let allUsers = [
    {
        id: "1",
        name: "John",
        gifts: [
            {
                id: "1",
                name: "Xbox",
                assignedUserId: "2"
            },
            {
                id: "2",
                name: "Ps6",
            },
            {
                id: "3",
                name: "Gamecube"
            }
        ],
        eventIds: ["1","2"]
    },
    {
        name: "Haritha",
        id: "2",
        gifts: [
            {
                id: "4",
                name: "Shoes"
            },
            {
                id: "5",
                name: "Ski goggles",
                assignedUserId: "1"
            },
            {
                id: "6",
                name: "Hat"
            }
        ],
        eventIds: ["1","2"]
    },
    {
        id: "3",
        name: "Sue",
        gifts: [
            {
                id: "7",
                name: "Skis"
            }
        ],
        eventIds: ["1"]
    },
    {
        id: "4",
        name: "Bruce",
        gifts: [
            {
                id: "8",
                name: "wine"
            },
            {
                id: "9",
                name: "beer",
                assignedUserId: "3"
            },
        ],
        eventIds: ["1"]
    }
]