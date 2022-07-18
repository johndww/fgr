<template>
  <div class="vh-100" style="background-color: #3da2c3;">
    <div class="container py-5 h-100">
      <div class="row d-flex justify-content-center align-items-center h-100">
        <div class="col col-lg-9 col-xl-7">
          <div class="card rounded-3">
            <div class="card-body p-4">

              <div v-show="currentUser().value">
                <nav class="navbar navbar-light bg-light">
                  <div class="container-fluid">
                    <router-link :to="{name: 'selectevent'}" class="navbar-brand" >Events</router-link>

                    <div style="text-align: right">
                      <a @click="logout" >{{ currentUser().value != null && currentUser().value.name }}</a>
                    </div>
                  </div>
                </nav>
              </div>

              <router-view></router-view>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

</template>

<script lang="ts">
import {defineComponent, ref} from "vue";
import {logoutUser, useCurrentUser} from "./state/users";
import {useRouter} from "vue-router";

export interface User {
  id: string,
  name: string,
  email: string,
}

export interface Event {
  id: string,
  name: string,
  ownerUserId: string
}

export interface GiftRequest {
  id: string
  userId: string,
  eventId: string,
  name: string
  isAssigned: boolean
  isAssignedToMe: boolean
}

export default defineComponent({
  name: 'App',
  setup() {
    const currentUser = function () {
      return useCurrentUser();
    }

    const router = useRouter()

    const logoutState = ref({data: "", error: "", loading: false})
    const logout = function () {
      return logoutUser(logoutState).then(() => {
        console.log("logged user out")
        router.push({name: "login"})
      })
    }

    return {
      currentUser,
      logout
    }
  }
})
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
