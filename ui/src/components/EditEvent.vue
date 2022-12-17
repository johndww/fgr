<template>
  <div class="event-header">
    <div class="view-event-title">
      <router-link class="event-header-title-back" :to="{name: 'selectevent'}">&lt All Events</router-link>
      <h1 class="event-header-title">Event Details</h1>
    </div>
    <div class="present-icon-container">
      <img src="../assets/present_icon.svg" alt="SimpleGiftApp" width="124" height="115">
    </div>
  </div>
  <div class="edit-event-contents">

    <div>
      <LoadingOrError :loading="this.getEventUsersState.loading"
                      :error="this.getEventUsersState.error"></LoadingOrError>
    </div>

    <form @submit="updateEvent">
    <div class="event-name">
      <span class="event-name-input-title">Event Name</span>
      <input type="text" class="event-name-input" v-model="eventName" />
    </div>

    <h2 class="members">Members</h2>

    <div class="member-list">


      <div class="member">
        <div class="member-logo-name">
          <img class="member-icon" src="../assets/user_icons/user_icon_purple.svg" alt="UserIcon" width="24"
               height="24">
          <span class="member-name">{{ currentUserState.data.email}}</span>
        </div>

<!--        <span class="member-email">{{ currentUserState.data.email }}</span>-->

        <div class="owner-or-trash">
          <span class="member-owner">Owner</span>
        </div>
      </div>

      <div class="member" v-for="(user, userIdx) in updateUserList" :key="userIdx">
        <div class="member-logo-name">
          <img class="member-icon" :src="userIcon(userIdx)" alt="UserIcon" width="24"
               height="24">
          <span class="member-name">{{ user.email }}</span>
        </div>

<!--        <span class="member-email">{{ user.email }}</span>-->

        <div class="owner-or-trash">
          <img class="member-trash" src="../assets/trash.svg" alt="Trash" width="18" height="16" @click="deleteUser(userIdx)">
        </div>
      </div>
    </div>

      <h2 class="members">Add New Member</h2>

      <datalist id="recentUserList">-->
        <option v-for="recentUser in recentUsers">{{ recentUser.email }}</option>
      </datalist>

      <div class="new-member-container">
<!--        <div class="new-member-name">-->
<!--          <span class="input-title">Name</span>-->
<!--          <input class="new-member-input" v-model="newUser.name">-->
<!--        </div>-->

        <div class="new-member-email">
          <span class="input-title">Email</span>
          <input class="new-member-input" list="recentUserList"  v-model="newUser.email">
        </div>

        <div class="new-member-button-container">
          <span class="input-title">&nbsp;</span>
          <button class="border-only-button" type="button" @click="addNewMember">Add Member</button>
        </div>
      </div>

      <div class="delete-save-event-footer">
        <div class="delete-event">
          <img src="../assets/trash.svg" alt="User" width="18" height="16" class="trash-icon">
          <button class="delete-button" type="button">Delete Event</button>
        </div>
        <button class="button" type="submit">Save Changes</button>
      </div>
    </form>
  </div>

  <!--  <LoadingOrError :loading="getEventUsersState.loading" :error="getEventUsersState.error"></LoadingOrError>-->
  <!--  <template v-if="currentUserState.data && event">-->
  <!--    <div>-->
  <!--      <h1 class="text-center my-3 pb-3">Edit {{ event.name }}</h1>-->
  <!--       <form @submit="updateEvent">-->
  <!--        <div class="tab-content">-->
  <!--          <div class="text-start">-->
  <!--          <input type="text" class="form-control" v-model="eventName" placeholder="Enter event name..">-->
  <!--            <p></p>-->
  <!--          <h3>Members:</h3>-->
  <!--          </div>-->
  <!--          <table class="table mb-4">-->
  <!--            <tbody>-->
  <!--            <datalist id="recentUserList">-->
  <!--              <option v-for="recentUser in recentUsers">{{recentUser.email}}</option>-->
  <!--            </datalist>-->
  <!--            <tr>-->
  <!--              <td>-->
  <!--                <input type="text" class="form-control" disabled v-model="currentUserState.data.email">-->
  <!--              </td>-->
  <!--              <td></td>-->
  <!--            </tr>-->
  <!--            <tr v-for="(user, userIdx) in updateUserList" :key="userIdx">-->
  <!--              <td>-->
  <!--                <input type="text" class="form-control" v-model="user.email" list="recentUserList">-->
  <!--              </td>-->

  <!--              <td><button type="button" class="btn btn-danger" @click="deleteUser(userIdx)">-->
  <!--              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-trash"-->
  <!--                   viewBox="0 0 16 16">-->
  <!--                <path-->
  <!--                    d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z"/>-->
  <!--                <path fill-rule="evenodd"-->
  <!--                      d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118zM2.5 3V2h11v1h-11z"/>-->
  <!--              </svg>-->
  <!--            </button></td>-->
  <!--            </tr>-->
  <!--            <tr>-->
  <!--             <td><button type="button" class="btn btn-success mt-5 mb-5" @click="addNewMember">Add Member</button></td>-->
  <!--            </tr>-->
  <!--            <tr><td><button type="submit">Update Event</button></td></tr>-->
  <!--            </tbody>-->
  <!--          </table>-->
  <!--        </div>-->
  <!--       </form>-->
  <!--    </div>-->
  <!--  </template>-->
</template>

<script lang="ts">

import {getUsersForEvent, useCurrentUserState, useCurrentUserId} from "../state/users";
import {useRoute, useRouter} from "vue-router";
import {useMyEventsState, persistUpdateEvent, MembersToAdd} from "../state/events";
import {computed, Ref, ref, watch, watchEffect} from "vue";
import {User} from "../App.vue";
import {State} from "../state/store";
import LoadingOrError from "./LoadingOrError.vue";

export default {
  components: {LoadingOrError},
  setup() {
    const myEventsState = useMyEventsState()
    myEventsState.value.fetch()

    const route = useRoute()
    let eventId = route.params.id as string;

    const event = computed(() => myEventsState.value.data.find((event) => event.id === eventId))

    const getEventUsersState = ref({
      data: <User[]>[],
      error: "",
      loading: false
    })
    getUsersForEvent(eventId, getEventUsersState)

    const eventName = ref("")
    watchEffect(() => {
      eventName.value = event.value?.name || ""
    })

    const userIcon = function (index: number) {
      const icons = [
        "user_icon_blue.svg",
        "user_icon_pink.svg",
        // "user_icon_purple.svg",
        "user_icon_yellow.svg",
      ]

      const iconIdx = index % icons.length
      return new URL(`/src/assets/user_icons/${icons[iconIdx]}`, import.meta.url).href
    }

    const updateUserList: Ref<{email: string }[]> = ref([])

    const currentUserState = useCurrentUserState()

    const router = useRouter()

    const updateEvent = function (e: any) {
      console.log("updateEvent called")
      e.preventDefault()

      if (!eventName.value) {
        alert("Event name cannot be empty")
        return
      }

      let uniqueEmails = new Set()
      updateUserList.value.forEach(user => {
        // could alert them on accidentally including themselves, but easier to just silently remove that email
        if (user.email != "" && user.email != currentUserState.value.data!.email) {
          uniqueEmails.add(user.email)
        }
      })

      uniqueEmails.add(currentUserState.value.data!.email)

      const state = ref({
        data: "",
        loading: false,
        error: ""
      })

      const promise = persistUpdateEvent(eventId, eventName.value, Array.from(uniqueEmails) as string[], state)
      promise.then(() => {
        router.push({name: 'viewevent', params: {id: eventId}})
      })
    }

    const deleteUser = function (userIdx: number) {
      updateUserList.value.splice(userIdx, 1)
    }

    const newUser = ref({email: ""})

    const addNewMember = function () {
      if (newUser.value.email === "") {
        alert("Must provide an email")
        return
      }

      const userAlreadyExists = updateUserList.value.find((existingUser) => {
        return existingUser.email === newUser.value.email
      })

      if (userAlreadyExists) {
        alert("Duplicate email found: " + userAlreadyExists.email)
        return
      }

      updateUserList.value.push({email: newUser.value.email})
    }

    const allEventsUsers: Ref<User[]> = ref([])

    watch(myEventsState, () => {
      const states: Ref<State<User[]>>[] = []

      Promise.all(myEventsState.value.data
          .map(event => {
            const getEventUsersState = ref({data: <User[]>[], error: "", loading: false})
            states.push(getEventUsersState)
            return getUsersForEvent(event.id, getEventUsersState)
          }))
          .then(() => {
            let filteredUsers = states
                .flatMap(state => state.value.data)
                .filter((v, i, a) => a.findIndex(v2 => (v2.id === v.id)) === i); // remove dup users from multiple events (from stackoverflow)
            allEventsUsers.value = [...filteredUsers]
          })
    }, {deep: true})

    const recentUsers = computed(() => {
      return allEventsUsers.value.filter(user => {
        if (user.id === currentUserState.value.data!.id) {
          // dont show the logged in user
          return false
        }

        if (updateUserList.value.find(updateUser => user.email === updateUser.email)) {
          // user is already in the update list, no need to show as a recent user
          return false
        }

        return true
      })
    })

    const currentUserId = useCurrentUserId()
    watchEffect(() => {
      getEventUsersState.value.data.forEach(user => {
        if (user.id !== currentUserId.value) {
          updateUserList.value.push({email: user.email})
        }
      })
    })

    return {
      getEventUsersState,
      userIcon,
      event,
      eventName,
      newUser,
      updateUserList,
      currentUserState,
      updateEvent,
      deleteUser,
      addNewMember,
      recentUsers
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->

<style scoped>

.edit-event-contents {
  padding-bottom: 62px;
  margin-left: 91px;
  margin-right: 91px;
}

.event-name {
  margin-top: 39px;
}

.event-name-input-title {
  text-align: left;
  font: normal normal 600 12px/22px Proxima Nova;
  letter-spacing: 0.3px;
  color: #97989B;
  display: block;
  text-transform: uppercase;
  margin-bottom: 3px;
}

.event-name-input {
  font: normal normal 600 16px/22px Proxima Nova;
  letter-spacing: 0px;
  color: #2F3237;
  width: 450px;
  height: 42px;
  border: 1px solid #E9E9E9;
  border-radius: 6px;
  display: block;
}

.members {
  margin-top: 39px;
  margin-bottom: 13px;
  font: normal normal bold 32px/39px Proxima Nova;
  letter-spacing: 0px;
  color: #2F3237;
  text-align: left;
}

.member-list {
  display: flex;
  flex-direction: column;
  row-gap: 20px;
  margin-bottom: 29px;
}

.member {
  background: #FAF9F8 0% 0% no-repeat padding-box;
  border: 1px solid #70707040;
  border-radius: 10px;
  height: 79px;
  display: inline-grid;
  grid-template-columns: auto auto;
  align-items: center;
}

.member-logo-name {
  display: flex;
  align-items: center;
}

.member-icon {
  margin-right: 10px;
  margin-left: 20px;
}

.owner-or-trash {
  margin-right: 20px;
  text-align: right;
}

.member-name {
  text-align: left;
}

.member-email {
  text-align: left;
  font: normal normal normal 16px/22px Proxima Nova;
  letter-spacing: 0px;
  color: #2F3237;
}

.member-owner {
  text-align: right;
  font: normal normal bold 12px/22px Proxima Nova;
  letter-spacing: 0.6px;
  color: #2F3237;
  text-transform: uppercase;
}

.new-member-container {
  display: flex;
  justify-content: space-between;
  height: 105px;
  align-items: center;
  column-gap: 25px;
  background: #FFFFFF 0% 0% no-repeat padding-box;
  border: 1px solid #70707040;
  border-radius: 10px;
  padding-left: 32px;
  padding-right: 32px;
}

/*.new-member-name {*/
/*  display: flex;*/
/*  flex-direction: column;*/
/*}*/

.input-title {
  text-align: left;
  font: normal normal 600 12px/22px Proxima Nova;
  letter-spacing: 0.3px;
  color: #97989B;
  text-transform: uppercase;
}

.new-member-input {
  border: 1px solid #E9E9E9;
  border-radius: 6px;
  width: 525px;
  height: 42px;
}

.new-member-email {
  display: flex;
  flex-direction: column;
}

.new-member-button-container {
  display: flex;
  flex-direction: column;
}

.delete-save-event-footer {
  display: flex;
  justify-content: space-between;
  margin-top: 30px;
}

.member-trash {
  cursor: pointer;
}

</style>
