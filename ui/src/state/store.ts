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

// export function persistUpdateEvent(id: string, name: string, memberEmails: string[]) {
    // console.log("memberEmails: " + JSON.stringify(memberEmails))
    // //TODO  honestly, logic is super difficult in mem. easier in a db with joiner tables
    //
    // memberEmails.push(useCurrentUser().value!.email)
    //
    // let event = getEvent(id)!
    // event.name = name
    //
    // // add new members
    // const userEmailsToCreate: string[] = []
    // const userIdsToAddMemberships: string[] = []
    // const memberUserIds = memberEmails
    //     .filter(email => {
    //         const existingUser = allUsersContainer.users.find(user => user.email == email)
    //         console.log("for email: " + email + " existingUser: " + JSON.stringify(existingUser))
    //         if (!existingUser) {
    //             userEmailsToCreate.push(email)
    //             return false
    //         }
    //
    //         const existingMembership = memberships.find(membership => membership.userId == existingUser.id && membership.eventId == id)
    //         console.log("for email: " + email + " existingMembership: " + JSON.stringify(existingMembership))
    //         if (!existingMembership) {
    //             userIdsToAddMemberships.push(existingUser.id)
    //         }
    //         return true
    //     })
    //     .map(email => {
    //     return allUsersContainer.users.find(user => user.email == email)!.id
    // })
    //
    // console.log("memberEmailsToCreate:" + JSON.stringify(userEmailsToCreate))
    // console.log("userIdsToAddMemberships:" + JSON.stringify(userIdsToAddMemberships))
    //
    // userEmailsToCreate.forEach(memberEmail => {
    //     const newUserId = Math.floor(Math.random() * 1000).toString();
    //     allUsersContainer.users.push({
    //         id: newUserId,
    //         name: memberEmail,
    //         email: memberEmail,
    //     })
    //
    //     memberUserIds.push(newUserId)
    //
    //     memberships.push({
    //         id: Math.floor(Math.random() * 1000).toString(),
    //         eventId: id,
    //         userId: newUserId,
    //     })
    // })
    //
    // userIdsToAddMemberships.forEach(userId => {
    //     memberships.push({
    //         id: Math.floor(Math.random() * 1000).toString(),
    //         eventId: id,
    //         userId: userId,
    //     })
    // })
    //
    // console.log("updatedMemberships: " + JSON.stringify(memberships))
    //
    // // delete old memberships
    // const membershipsToDelete: string[] = []
    // memberships.forEach(membership => {
    //     if (membership.eventId == id && !memberUserIds.includes(membership.userId)) {
    //         membershipsToDelete.push(membership.id)
    //
    //         // remove gift requests for this event for the user that's no longer a member
    //         giftRequests
    //             .filter(request => request.userId == membership.userId && request.eventId == id)
    //             .forEach(requestToDelete => {
    //                 console.log("gift request deleted: " + JSON.stringify(requestToDelete))
    //                 const idx = giftRequests.findIndex(request => request.id == requestToDelete.id)
    //                 giftRequests.splice(idx, 1)
    //             })
    //
    //         // un-assign gift requests for this event that the user had claimed
    //         giftRequests
    //             .filter(request => request.assignedUserId == membership.userId && request.eventId == id)
    //             .forEach(request => request.assignedUserId = "")
    //     }
    // })
    //
    // console.log("membershipsToDelete: " + JSON.stringify(membershipsToDelete))
    //
    // membershipsToDelete.forEach(membershipId => {
    //     const idx = memberships.findIndex(membership => membership.id == membershipId)
    //     memberships.splice(idx, 1)
    // })
// }
