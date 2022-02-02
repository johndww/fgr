<template>
  <div v-show="currentUserName" style="text-align: right">
    {{ currentUserName }}
  </div>
  <SelectUser @user-selection="userSelection" @user-creation="createUser" v-show="!currentUserName" :allUsers="allUsers"/>
  <GiftResults v-show="currentUserName" :currentUserName="currentUserName" :allUsers="allUsers"/>
</template>

<script lang="ts">
import SelectUser from './components/SelectUser.vue';
import GiftResults from "./components/GiftResults.vue";

export default {
  name: 'App',
  props: {
  },
  data() {
    return {
      currentUserName: "",
      allUsers: []
    }
  },
  methods: {
    userSelection(userName: String) {
      this.currentUserName = userName
    },

    createUser(userName: String) {
      this.currentUserName = userName
      this.allUsers = [...this.allUsers, {name: userName}]
    }
  },
  components: {
    GiftResults,
    SelectUser
  },
  created() {
    this.allUsers = [
      {
        name: "John",
        gifts: [
          {
            name: "Xbox"
          },
          {
            name: "Ps6"
          },
          {
            name: "Gamecube"
          }
        ]
      },
      {
        name: "Haritha",
        gifts: [
          {
            name: "hGift1"
          },
          {
            name: "hGift2"
          },
          {
            name: "hGift3"
          }
        ]
      },
      {
        name: "Sue",
        gifts: [
          {
            name: "test"
          }
        ]
      },
      {
        name: "Bruce",
        gifts: [
          {
            name: "wine"
          },
          {
            name: "beer"
          },
        ]
      }
    ]
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
