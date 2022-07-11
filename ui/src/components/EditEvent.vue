<template>
  <div>
    <h1 class="text-center my-3 pb-3">Edit {{ this.event.name }}</h1>
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

<script lang="ts">

import {defineComponent} from "vue";
import {
  getCurrentUser,
  getEvent,
  getEventUsers,
  getMyEvents,
  getSessionUserId,
  persistUpdateEvent
} from "../state/store";

export default defineComponent({
  name: "EditEvent",
  components: {},
  props: {
  },
  data() {
    return {
      event: getEvent(this.$route.params.id)!,
      eventUsers: getEventUsers(this.$route.params.id),
      eventName: "",
      updateUserList: [{email: ""}],
    }
  },
  methods: {
    updateEvent: function (e: any) {
      e.preventDefault()

      if (!this.eventName) {
        alert("Event name cannot be empty")
        return
      }

      let uniqueEmails = new Set()
      this.updateUserList.forEach(user => {
        // could alert them on accidentally including themselves, but easier to just silently remove that email
        if (user.email != "" && user.email != getCurrentUser()!.email) {
          uniqueEmails.add(user.email)
        }
      })

      persistUpdateEvent(this.event.id, this.eventName, Array.from(uniqueEmails) as string[])
      this.$router.push({name: 'viewevent', params: {id: this.event.id}})
    },

    deleteUser(userIdx: number) {
      this.updateUserList.splice(userIdx, 1)
    },

    addNewMember() {
      this.updateUserList.push({email: ""})
    }
  },
  computed: {
    currentUser: function () {
      return getCurrentUser()
    },
    recentUsers: function () {
      return getMyEvents()
          .flatMap(event => getEventUsers(event.id))
          .filter((v,i,a)=>a.findIndex(v2=>(v2.id===v.id))===i) // remove dup users from multiple events (from stackoverflow)
          .filter(user => {
            if (user.id === getSessionUserId().value) {
              // dont show the logged in user
              return false
            }

            if (this.updateUserList.find(updateUser => user.email === updateUser.email)) {
              // user is already in the update list, no need to show as a recent user
              return false
            }

            return true
          })
    }
  },

  created() {
    this.updateUserList = []
    this.eventUsers.forEach(user => {
      if (user.id !== getSessionUserId().value) {
        this.updateUserList.push({email: user.email})
      }
    })
    this.eventName = this.event.name
  }
})
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->

<style scoped>

.edit {
  color: #42b983;
  font-size: 12px;
}
</style>
