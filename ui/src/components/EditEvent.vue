<template>
  <LoadingOrError :loading="getEventUsersState.loading" :error="getEventUsersState.error"></LoadingOrError>
  <template v-if="currentUser && event">
    <div>
      <h1 class="text-center my-3 pb-3">Edit {{ event.name }}</h1>
       <form @submit="updateEvent">
        <div class="tab-content">
          <div class="text-start">
          <input type="text" class="form-control" v-model="eventName" placeholder="Enter event name..">
            <p></p>
          <h3>Members:</h3>
          </div>
          <table class="table mb-4">
            <tbody>
            <datalist id="recentUserList">
              <option v-for="recentUser in recentUsers">{{recentUser.email}}</option>
            </datalist>
            <tr>
              <td>
                <input type="text" class="form-control" disabled v-model="currentUser.email">
              </td>
              <td></td>
            </tr>
            <tr v-for="(user, userIdx) in updateUserList" :key="userIdx">
              <td>
                <input type="text" class="form-control" v-model="user.email" list="recentUserList">
              </td>

              <td><button type="button" class="btn btn-danger" @click="deleteUser(userIdx)">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-trash"
                   viewBox="0 0 16 16">
                <path
                    d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z"/>
                <path fill-rule="evenodd"
                      d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118zM2.5 3V2h11v1h-11z"/>
              </svg>
            </button></td>
            </tr>
            <tr>
             <td><button type="button" class="btn btn-success mt-5 mb-5" @click="addNewMember">Add Member</button></td>
            </tr>
            <tr><td><button type="submit">Update Event</button></td></tr>
            </tbody>
          </table>
        </div>
       </form>
    </div>
  </template>
</template>

<script lang="ts">

import {getUsersForEvent, useCurrentUser, useCurrentUserId} from "../state/users";
import {useRoute, useRouter} from "vue-router";
import {useMyEventsState, persistUpdateEvent} from "../state/events";
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

    const updateUserList: Ref<{ email: string }[]> = ref([])

    const currentUser: Ref<User | null> = useCurrentUser()

    const router = useRouter()

    const updateEvent = function(e: any) {
      e.preventDefault()

      if (!eventName.value) {
        alert("Event name cannot be empty")
        return
      }

      let uniqueEmails = new Set()
      updateUserList.value.forEach(user => {
        // could alert them on accidentally including themselves, but easier to just silently remove that email
        if (user.email != "" && user.email != currentUser.value!.email) {
          uniqueEmails.add(user.email)
        }
      })

      uniqueEmails.add(currentUser.value!.email)

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

    const deleteUser = function(userIdx: number) {
      updateUserList.value.splice(userIdx, 1)
    }

    const addNewMember = function() {
      updateUserList.value.push({email: ""})
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
        if (user.id === currentUser.value!.id) {
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
      event,
      eventName,
      updateUserList,
      currentUser,
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

.edit {
  color: #42b983;
  font-size: 12px;
}
</style>
