<template>

  <div class="vh-100" style="background-color: #3da2c3;">
    <div class="container py-5 h-100">
      <div class="row d-flex justify-content-center align-items-center h-100">
        <div class="col col-lg-9 col-xl-7">
          <div class="card rounded-3">
            <div class="card-body p-4">

              <div v-show="currentUserId" style="text-align: right">
                {{ currentUserName }}
              </div>
              <SelectUser @user-selection="userSelection" @user-creation="createUser" v-show="!currentUserId"
                          :allUsers="allUsers"/>
              <SelectEvent @event-selection="eventSelection" @event-creation="createEvent" v-if="currentUserId && !selectedEventId"
                          :allEvents="allEvents"/>
              <GiftResults v-if="currentUserId && selectedEventId" :currentUserId="currentUserId" :eventUsers="eventUsers"
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
import SelectEvent from './components/SelectEvent.vue';
import GiftResults from "./components/GiftResults.vue";
import {defineComponent} from "vue";

declare interface User {
  id: string,
  name: string,
  gifts: Gift[]
  eventIds: string[] // joiner, not persisted like this todo
}

declare interface Event {
  id: string,
  name: string,
  ownerUserId: string
}

declare interface Gift {
  id: string
  name: string
  assignedUserId?: string
}

export { User, Event, Gift }

export default defineComponent({
  name: 'App',
  props: {},
  data() {
    return {
      currentUserId: "",
      selectedEventId: "",
      allUsers: [] as User[],
      allEvents: [] as Event[],
    }
  },
  methods: {
    userSelection(userId: string) {
      this.currentUserId = userId
    },

    eventSelection(eventId: string) {
      this.selectedEventId = eventId
    },

    createUser(userName: string) {
      this.currentUserId = Math.floor(Math.random() * 1000).toString()
      this.allUsers = [...this.allUsers, {id: this.currentUserId, name: userName, gifts: [], eventIds: []}]
    },

    createEvent(eventName: string) {
      this.selectedEventId = Math.floor(Math.random() * 1000).toString()
      this.allEvents = [...this.allEvents, {id: "", name: eventName, ownerUserId: this.currentUserId}]
    },

    assignGift(details: { forUserId: string; giftId: string; byUserId: string; }) {
      const user = this.allUsers.find((user) => user.id === details.forUserId) || null
      if (user == null) {
        console.log("unable to assign gift. cannot find user: " + details.forUserId)
        return
      }

      const gift = user.gifts.find((gift: { id: string; }) => gift.id === details.giftId) || null
      if (gift == null) {
        console.log("unable to assign gift. cannot find gift: " + details.giftId)
        return
      }
      gift.assignedUserId = details.byUserId
    },

    addGift(giftName: string) {
      const user = this.allUsers.find((user) => user.id === this.currentUserId) || null
      if (user == null) {
        console.log("unable to add gift. cannot find user: " + this.currentUserId)
        return
      }

      user.gifts = [...user.gifts, {id: Math.floor(Math.random() * 1000).toString(), name: giftName, assignedUserId: ""}]
    },

    deleteGift(giftId: string) {
      const user = this.allUsers.find((user) => user.id === this.currentUserId) || null
      if (user == null) {
        console.log("unable to delete gift. cannot find user: " + this.currentUserId)
        return
      }

      user.gifts = user.gifts.filter((gift) => gift.id !== giftId)
    },

    releaseGift(details: { forUserId: string; giftId: string; }) {
      const user = this.allUsers.find((user) => user.id === details.forUserId) || null
      if (user == null) {
        console.log("unable to release gift. cannot find user: " + details.forUserId)
        return
      }

      const gift = user.gifts.find((gift) => gift.id === details.giftId) || null
      if (gift == null) {
        console.log("unable to release gift. cannot find gift: " + details.giftId)
        return
      }

      gift.assignedUserId = undefined
    }
  },
  computed: {
    eventUsers: function (): User[] {
      return this.allUsers.filter((user) => user.eventIds.includes(this.selectedEventId))
    },
    currentUserName: function (): User[] {
      return this.allUsers.filter((user) => user.id === this.currentUserId)
    },
  },
  components: {
    GiftResults,
    SelectUser,
    SelectEvent
  },
  created() {
    this.allEvents = [
      {
        id: "1",
        name: "2020 Wright's Christmas",
        ownerUserId: "1",
      },
      {
        id: "2",
        name: "2021 Tapa's Thanksgiving",
        ownerUserId: "2",
      },
    ]

    this.allUsers = [
      {
        id: "1",
        name: "John",
        gifts: [
          {
            id: "1",
            name: "Xbox",
            assignedUserId: "2"
          },
          {
            id: "2",
            name: "Ps6",
          },
          {
            id: "3",
            name: "Gamecube"
          }
        ],
        eventIds: ["1","2"]
      },
      {
        name: "Haritha",
        id: "2",
        gifts: [
          {
            id: "4",
            name: "Shoes"
          },
          {
            id: "5",
            name: "Ski goggles",
            assignedUserId: "1"
          },
          {
            id: "6",
            name: "Hat"
          }
        ],
        eventIds: ["1","2"]
      },
      {
        id: "3",
        name: "Sue",
        gifts: [
          {
            id: "7",
            name: "Skis"
          }
        ],
        eventIds: ["1"]
      },
      {
        id: "4",
        name: "Bruce",
        gifts: [
          {
            id: "8",
            name: "wine"
          },
          {
            id: "9",
            name: "beer",
            assignedUserId: "3"
          },
        ],
        eventIds: ["1"]
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
