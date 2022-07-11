<template>
  <div>
    <h1 class="text-center my-3 pb-3">Who are you?</h1>

    <ul>
      <li v-for="user in allUsers" :key="user.name"><button @click="selectUser(user.id)">{{ user.name }}</button></li>
    </ul>
    <form @submit="onCreateUser">
      <input name="name" v-model="createdUserName"> <button type="submit">Join Registry</button>
    </form>
  </div>
</template>

<script lang="ts">

import {defineComponent} from "vue";
import {createUser, getAllUsers, login} from "../state/store";

export default defineComponent({
  name: "SelectUser",
  props: {
  },
  data() {
    return {
      createdUserName: ''
    }
  },
  methods: {
    onCreateUser(e: any) {
      e.preventDefault()

      if (!this.createdUserName) {
        alert("Please enter a username")
        return
      }

      //TODO prevent dup names
      //TODO remember where you came from before login
      createUser(this.createdUserName)
      this.$router.push({ name: 'selectevent'})
    },

    selectUser(id: string) {
      login(id)
      this.$router.push({ name: 'selectevent'})
    }
  },
  computed: {
    allUsers: function () {
      return getAllUsers()
    }
  }
})
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
