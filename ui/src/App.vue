<template>

  <div class="vh-100" style="background-color: #3da2c3;">
    <div class="container py-5 h-100">
      <div class="row d-flex justify-content-center align-items-center h-100">
        <div class="col col-lg-9 col-xl-7">
          <div class="card rounded-3">
            <div class="card-body p-4">

              <div v-show="currentUserName" style="text-align: right">
                {{ currentUserName }}
              </div>
              <SelectUser @user-selection="userSelection" @user-creation="createUser" v-show="!currentUserName"
                          :allUsers="allUsers"/>
              <GiftResults v-show="currentUserName" :currentUserName="currentUserName" :allUsers="allUsers"
                           @assign-gift="assignGift" @add-gift="addGift" @delete-gift="deleteGift"
                           @release-gift="releaseGift"/>

            </div>
          </div>
        </div>
      </div>
    </div>
  </div>


</template>

<script lang="ts">
import SelectUser from './components/SelectUser.vue';
import GiftResults from "./components/GiftResults.vue";
import { defineComponent } from "vue";

declare interface User {
  name: string,
  gifts: Gift[]
}

declare interface Gift {
  id: string
  name: string
  assignedUserName?: string
}

export default defineComponent({
  name: 'App',
  props: {},
  data() {
    return {
      currentUserName: "",
      allUsers: [] as User[],
    }
  },
  methods: {
    userSelection(userName: string) {
      this.currentUserName = userName
    },

    createUser(userName: string) {
      this.currentUserName = userName
      this.allUsers = [...this.allUsers, {name: userName, gifts: []}]
    },
    assignGift(details: { forUserName: string; giftId: string; byUserName: string; }) {
      const user = this.allUsers.find((user) => user.name === details.forUserName) || null
      if (user == null) {
        console.log("unable to assign gift. cannot find user: " + details.forUserName)
        return
      }

      const gift = user.gifts.find((gift: { id: string; }) => gift.id === details.giftId) || null
      if (gift == null) {
        console.log("unable to assign gift. cannot find gift: " + details.giftId)
        return
      }
      gift.assignedUserName = details.byUserName
    },
    addGift(giftName: string) {
      const user = this.allUsers.find((user) => user.name === this.currentUserName) || null
      if (user == null) {
        console.log("unable to add gift. cannot find user: " + this.currentUserName)
        return
      }

      user.gifts = [...user.gifts, {id: Math.floor(Math.random() * 1000).toString(), name: giftName, assignedUserName: ""}]
    },
    deleteGift(giftId: string) {
      const user = this.allUsers.find((user) => user.name === this.currentUserName) || null
      if (user == null) {
        console.log("unable to delete gift. cannot find user: " + this.currentUserName)
        return
      }

      user.gifts = user.gifts.filter((gift) => gift.id !== giftId)
    },
    releaseGift(details: { forUserName: string; giftId: string; }) {
      const user = this.allUsers.find((user) => user.name === details.forUserName) || null
      if (user == null) {
        console.log("unable to release gift. cannot find user: " + details.forUserName)
        return
      }

      const gift = user.gifts.find((gift) => gift.id === details.giftId) || null
      if (gift == null) {
        console.log("unable to release gift. cannot find gift: " + details.giftId)
        return
      }

      gift.assignedUserName = undefined
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
            id: "1",
            name: "Xbox",
            assignedUserName: "Haritha"
          },
          {
            id: "2",
            name: "Ps6",
          },
          {
            id: "3",
            name: "Gamecube"
          }
        ]
      },
      {
        name: "Haritha",
        gifts: [
          {
            id: "4",
            name: "Shoes"
          },
          {
            id: "5",
            name: "Ski goggles",
            assignedUserName: "John"
          },
          {
            id: "6",
            name: "Hat"
          }
        ]
      },
      {
        name: "Sue",
        gifts: [
          {
            id: "7",
            name: "Skis"
          }
        ]
      },
      {
        name: "Bruce",
        gifts: [
          {
            id: "8",
            name: "wine"
          },
          {
            id: "9",
            name: "beer",
            assignedUserName: "Sue"
          },
        ]
      }
    ]
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
  margin-top: 60px;
}
</style>
