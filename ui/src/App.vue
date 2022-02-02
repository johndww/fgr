<template>
  <div v-show="currentUserName" style="text-align: right">
    {{ currentUserName }}
  </div>
  <SelectUser @user-selection="userSelection" @user-creation="createUser" v-show="!currentUserName"
              :allUsers="allUsers"/>
  <GiftResults v-show="currentUserName" :currentUserName="currentUserName" :allUsers="allUsers"
               @assign-gift="assignGift" @add-gift="addGift" @delete-gift="deleteGift" @release-gift="releaseGift"/>
</template>

<script lang="ts">
import SelectUser from './components/SelectUser.vue';
import GiftResults from "./components/GiftResults.vue";

export default {
  name: 'App',
  props: {},
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
      this.allUsers = [...this.allUsers, {name: userName, gifts: []}]
    },
    assignGift(details) {
      const user = this.allUsers.find((user) => user.name === details.forUserName)
      const gift = user.gifts.find((gift) => gift.id === details.giftId)
      gift.assignedUserName = details.byUserName
    },
    addGift(giftName: String) {
      const user = this.allUsers.find((user) => user.name === this.currentUserName)
      user.gifts = [...user.gifts, {id: Math.floor(Math.random() * 1000), name: giftName}]
    },
    deleteGift(giftId) {
      const user = this.allUsers.find((user) => user.name === this.currentUserName)
      user.gifts = user.gifts.filter((gift) => gift.id !== giftId)
    },
    releaseGift(details) {
      const user = this.allUsers.find((user) => user.name === details.forUserName)
      const gift = user.gifts.find((gift) => gift.id === details.giftId)
      gift.assignedUserName = ''
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
            id: 1,
            name: "Xbox",
            assignedUserName: "Haritha"
          },
          {
            id: 2,
            name: "Ps6"
          },
          {
            id: 3,
            name: "Gamecube"
          }
        ]
      },
      {
        name: "Haritha",
        gifts: [
          {
            id: 4,
            name: "Shoes"
          },
          {
            id: 5,
            name: "Ski goggles",
            assignedUserName: "John"
          },
          {
            id: 6,
            name: "Hat"
          }
        ]
      },
      {
        name: "Sue",
        gifts: [
          {
            id: 7,
            name: "Skis"
          }
        ]
      },
      {
        name: "Bruce",
        gifts: [
          {
            id: 8,
            name: "wine"
          },
          {
            id: 9,
            name: "beer",
            assignedUserName: "Sue"
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
