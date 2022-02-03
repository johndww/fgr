<template>
  <div>
    <h1 class="text-center my-3 pb-3">Who are you?</h1>

    <ul>
      <li v-for="user in allUsers" :key="user.name"><button @click="$emit('user-selection', user.name)">{{ user.name }}</button></li>
    </ul>
    <form @submit="onCreateUser">
      <input name="name" v-model="createdUser"> <button type="submit">Join Registry</button>
    </form>
  </div>
</template>

<script lang="ts">

export default {
  name: "SelectUser",
  props: {
    allUsers: Array
  },
  data() {
    return {
      createdUser: ''
    }
  },
  methods: {
    onCreateUser(e) {
      e.preventDefault()

      if (!this.createdUser) {
        alert("Please enter a username")
        return
      }

      //TODO prevent dup names

      this.$emit('user-creation', this.createdUser)
    }
  }
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
