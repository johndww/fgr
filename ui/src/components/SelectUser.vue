<template>
  <div>
    <h1 class="text-center my-3 pb-3">Admin: Login to another user</h1>
    <div v-if="createUserState.error">{{createUserState.error}}</div>

    <div v-if="allUsersState.data.length > 0">
      <ul>
        <li v-for="user in allUsersState.data" :key="user.name"><button @click="selectUser(user.id)">{{ user.name }}</button></li>
      </ul>
      <form @submit="onCreateUser">
        <input name="name" v-model="createUserName"> <button type="submit">Join Registry</button>
      </form>
    </div>
    <LoadingOrError :loading="allUsersState.loading" :error="allUsersState.error"></LoadingOrError>
  </div>
</template>

<script lang="ts">

import {Ref, ref} from "vue";
import {createUser, CreateUserState, login, useAllUsers, useCurrentUserState,} from "../state/users";
import {useRouter} from "vue-router";
import LoadingOrError from "./LoadingOrError.vue";

export default {
  components: {LoadingOrError},
  setup() {
    const allUsersState = useAllUsers()
    allUsersState.value.fetch()

    const createUserName = ref("")

    const router = useRouter()

    const createUserState: Ref<CreateUserState> = ref({
      userId: '',
      error: '',
      loading: false,
    })

    const onCreateUser = function(e: any) {
      e.preventDefault()

      if (!createUserName) {
        alert("Please enter a username")
        return
      }

      //TODO remember where you came from before login
      const promise = createUser(createUserName.value, createUserState)

      promise.then(() => {
        let loginPromise = login(createUserState.value.userId)
        loginPromise.then(() => {
          router.push({ name: 'selectevent'})
        }).catch(err => {
          console.log("unable to login user: " + err)
          return Promise.reject(err)
        })
      })
    }

    const selectUser = function(id: string) {
      login(id).then(() => {
        router.push({ name: 'selectevent'})
      })
      .catch(err => {
        console.log("unable to login user: " + err)
        return Promise.reject(err)
      })
    }

    return {
      allUsersState,
      createUserName,
      createUserState,
      onCreateUser,
      selectUser
    }
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
