<template>
  <div>
    <h1 class="text-center my-3 pb-3">Admin: Login to another user</h1>

    <div>
      <form @submit="login">
        <input type="text" v-model="userId">
        <button type="submit">Login</button>
      </form>
    </div>
  </div>
</template>

<script lang="ts">

import {ref} from "vue";
import {devLogin, useCurrentUserState,} from "../state/users";
import {useRouter} from "vue-router";

export default {
  components: {},
  setup() {

    const userId = ref("")

    const router = useRouter()

    const login = function(e: any) {
      e.preventDefault()

      if (!userId) {
        alert("Please enter a userid")
        return
      }

      const promise = devLogin(userId.value)

      promise.then(() => {
        useCurrentUserState().value.fetch().then(
            () => {
              console.log("redirecting to selectevent")
              router.push({ name: 'selectevent'})
            }
        )
      })
    }

    return {
      userId,
      login,
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
