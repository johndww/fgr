import {reactive, ref} from "vue";

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
    allUsers = [...allUsers, {id: userId, name: userName, email: userName + "@email.com"}]
    login(userId)
}

export function createEvent(eventName: string) {
    let newEventId = Math.floor(Math.random() * 1000).toString()
    allEvents = [...allEvents, {id: newEventId, name: eventName, ownerUserId: sessionUserId.value}]

    memberships.push({
        id: Math.floor(Math.random() * 1000).toString(),
        eventId: newEventId,
        userId: sessionUserId.value
    })
    return newEventId
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
    return memberships
        .filter(membership => membership.userId == getSessionUserId().value)
        .map(membership => allEvents.find(event => event.id == membership.eventId)!)
}

export function getEventUsers(eventId: any) {
    return memberships
        .filter(membership => membership.eventId == eventId)
        .map(membership => allUsers.find(user => user.id == membership.userId)!)
}

export function getGiftRequests(eventId: string) {
    return giftRequests.filter(request => request.eventId == eventId)
}

export function persistGiftRequest(giftName: string, eventId: string) {
    giftRequests.push(
        {
            id: Math.floor(Math.random() * 1000).toString(),
            userId: sessionUserId.value,
            eventId: eventId,
            name: giftName,
            assignedUserId: ""
        })
}

export function persistDeleteGiftRequest(giftId: string) {
    const idx = giftRequests.findIndex(request => request.id = giftId)
    giftRequests.splice(idx, 1)
}

export function persistReleaseGift(giftId: string) {
    let giftRequest = giftRequests.find(request => request.id == giftId)!;
    giftRequest.assignedUserId = ""
}

export function persistAssignGift(giftId: string, byUserId: string) {
    const request = giftRequests.find(request => request.id == giftId)!
    request.assignedUserId = byUserId
}

export function persistUpdateEvent(id: string, name: string, memberEmails: string[]) {
    console.log("memberEmails: " + JSON.stringify(memberEmails))
    //TODO  honestly, logic is super difficult in mem. easier in a db with joiner tables

    memberEmails.push(getCurrentUser()!.email)

    let event = getEvent(id)!
    event.name = name

    // add new members
    const userEmailsToCreate: string[] = []
    const userIdsToAddMemberships: string[] = []
    const memberUserIds = memberEmails
        .filter(email => {
            const existingUser = allUsers.find(user => user.email == email)
            console.log("for email: " + email + " existingUser: " + JSON.stringify(existingUser))
            if (!existingUser) {
                userEmailsToCreate.push(email)
                return false
            }

            const existingMembership = memberships.find(membership => membership.userId == existingUser.id && membership.eventId == id)
            console.log("for email: " + email + " existingMembership: " + JSON.stringify(existingMembership))
            if (!existingMembership) {
                userIdsToAddMemberships.push(existingUser.id)
            }
            return true
        })
        .map(email => {
        return allUsers.find(user => user.email == email)!.id
    })

    console.log("memberEmailsToCreate:" + JSON.stringify(userEmailsToCreate))
    console.log("userIdsToAddMemberships:" + JSON.stringify(userIdsToAddMemberships))

    userEmailsToCreate.forEach(memberEmail => {
        const newUserId = Math.floor(Math.random() * 1000).toString();
        allUsers.push({
            id: newUserId,
            name: memberEmail,
            email: memberEmail,
        })

        memberUserIds.push(newUserId)

        memberships.push({
            id: Math.floor(Math.random() * 1000).toString(),
            eventId: id,
            userId: newUserId,
        })
    })

    userIdsToAddMemberships.forEach(userId => {
        memberships.push({
            id: Math.floor(Math.random() * 1000).toString(),
            eventId: id,
            userId: userId,
        })
    })

    console.log("updatedMemberships: " + JSON.stringify(memberships))

    // delete old memberships
    const membershipsToDelete: string[] = []
    memberships.forEach(membership => {
        if (membership.eventId == id && !memberUserIds.includes(membership.userId)) {
            membershipsToDelete.push(membership.id)

            // remove gift requests for this event for the user that's no longer a member
            giftRequests
                .filter(request => request.userId == membership.userId && request.eventId == id)
                .forEach(requestToDelete => {
                    console.log("gift request deleted: " + JSON.stringify(requestToDelete))
                    const idx = giftRequests.findIndex(request => request.id == requestToDelete.id)
                    giftRequests.splice(idx, 1)
                })

            // un-assign gift requests for this event that the user had claimed
            giftRequests
                .filter(request => request.assignedUserId == membership.userId && request.eventId == id)
                .forEach(request => request.assignedUserId = "")
        }
    })

    console.log("membershipsToDelete: " + JSON.stringify(membershipsToDelete))

    membershipsToDelete.forEach(membershipId => {
        const idx = memberships.findIndex(membership => membership.id == membershipId)
        memberships.splice(idx, 1)
    })
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

let allUsers = reactive([
    {
        id: "1",
        name: "John",
        email: "john.d.wright@gmail.com",
    },
    {
        id: "2",
        name: "Haritha",
        email: "haritha.tapa@gmail.com",

    },
    {
        id: "3",
        name: "Sue",
        email: "paubsue@gmail.com",
    },
    {
        id: "4",
        name: "Bruce",
        email: "bruce.d.wright@gmail.com",
    }
])

let giftRequests = reactive([
    {
        id: "111",
        userId: "1",
        eventId: "1",
        name: "Xbox",
        assignedUserId: "2"
    },
    {
        id: "112",
        userId: "1",
        eventId: "1",
        name: "Ps6",
        assignedUserId: ""
    },
    {
        id: "113",
        userId: "1",
        eventId: "1",
        name: "Gamecube",
        assignedUserId: ""
    },
    {
        id: "121",
        userId: "1",
        eventId: "2",
        name: "Kite",
        assignedUserId: ""
    },
    {
        id: "211",
        userId: "2",
        eventId: "1",
        name: "Shoes",
        assignedUserId: ""
    },
    {
        id: "212",
        userId: "2",
        eventId: "1",
        name: "Ski goggles",
        assignedUserId: "1"
    },
    {
        id: "213",
        userId: "2",
        eventId: "1",
        name: "Hat",
        assignedUserId: ""
    },
    {
        id: "221",
        userId: "2",
        eventId: "2",
        name: "Purse",
        assignedUserId: ""
    },
    {
        id: "222",
        userId: "2",
        eventId: "2",
        name: "Skis",
        assignedUserId: ""
    },
    {
        id: "7",
        userId: "3",
        eventId: "1",
        name: "Skis",
        assignedUserId: ""
    },
    {
        id: "8",
        userId: "4",
        eventId: "1",
        name: "wine",
        assignedUserId: ""
    },
    {
        id: "9",
        userId: "4",
        eventId: "1",
        name: "beer",
        assignedUserId: "3"
    }
])

let memberships = reactive([
    {
        id: "1",
        eventId: "1",
        userId: "1",
    },
    {
        id: "2",
        eventId: "1",
        userId: "2",
    },
    {
        id: "3",
        eventId: "1",
        userId: "3",
    },
    {
        id: "4",
        eventId: "1",
        userId: "4",
    },
    {
        id: "5",
        eventId: "2",
        userId: "1",
    },
    {
        id: "6",
        eventId: "2",
        userId: "2",
    },
])