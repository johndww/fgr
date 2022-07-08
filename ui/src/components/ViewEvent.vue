<template>
  <div>
    <h1 class="text-center my-3 pb-3">{{ this.event.name }} <a class="edit" onclick="alert('todo')" v-if="this.event.ownerUserId === currentUserId.value">(edit)</a></h1>
    <div class="tabbable tabs-left">
      <ul class="nav nav-tabs">
        <li v-for="user in eventUsers" :key="user.id">
          <button
              :class="user.id === this.viewResultsUserId ? (this.currentUserId.value === user.id ? 'btn btn-warning' : 'btn btn-primary') : 'btn btn-secondary'"
              @click="this.viewResultsUserId = user.id">{{ user.name }}
          </button>
        </li>
      </ul>
    </div>

    <template v-if="this.viewResultsUser && currentUserId.value !== this.viewResultsUserId">
      <div class="tab-content">
        <table class="table mb-4">
          <tbody>
          <tr v-for="gift in unassignedGifts" :key="gift.name">
            <td>{{ gift.name }}</td>
            <td>
              <button
                  @click="assignGift({giftId: gift.id, forUserId: this.viewResultsUserId, byUserId: currentUserId.value})"
                  class="btn btn-primary">
                Claim Gift
              </button>
            </td>
          </tr>

          <tr v-for="gift in assignedGifts" :key="gift.name">
            <td>{{ gift.name }}</td>
            <template v-if="gift.assignedUserId === currentUserId.value">
              <td>
                <button class="btn btn-warning"
                        @click="releaseGift({giftId: gift.id, forUserId: this.viewResultsUserId})">
                  Release Gift
                </button>
              </td>
            </template>
          </tr>
          </tbody>
        </table>

      </div>
    </template>

    <SelfGiftResults @add-gift="(giftToAdd) => addGift(giftToAdd)"
                     @delete-gift="(giftId) => deleteGift(giftId)"
                     :viewResultsUser="this.viewResultsUser"
                     v-if="this.viewResultsUser && currentUserId.value === this.viewResultsUserId"/>

  </div>
</template>

<script lang="ts">

import SelfGiftResults from "./SelfGiftResults.vue";
import {defineComponent} from "vue";
import {getEvent, getEventUsers, getSessionUserId} from "../state/store";
import {Gift, User} from "@/App.vue";

export default defineComponent({
  name: "ViewEvent",
  components: {SelfGiftResults},
  props: {
  },
  data() {
    return {
      viewResultsUserId: getSessionUserId().value,
      event: getEvent(this.$route.query.id),
      eventUsers: getEventUsers(this.$route.query.id),
    }
  },
  methods: {
    assignGift(details: { forUserId: string; giftId: string; byUserId: string }) {
      const user = this.eventUsers.find((user) => user.id === details.forUserId) || null
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
      const user = this.eventUsers.find((user) => user.id === this.currentUserId.value) || null
      if (user == null) {
        console.log("unable to add gift. cannot find user: " + this.currentUserId)
        return
      }

      user.gifts = [...user.gifts, {id: Math.floor(Math.random() * 1000).toString(), name: giftName, assignedUserId: ""}]
    },

    deleteGift(giftId: string) {
      const user = this.eventUsers.find((user) => user.id === this.currentUserId.value) || null
      if (user == null) {
        console.log("unable to delete gift. cannot find user: " + this.currentUserId)
        return
      }

      user.gifts = user.gifts.filter((gift) => gift.id !== giftId)
    },

    releaseGift(details: { forUserId: string; giftId: string; }) {
      const user = this.eventUsers.find((user) => user.id === details.forUserId) || null
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
    currentUserId: function () {
      return getSessionUserId()
    },
    viewResultsUser: function (): User | null {
      return this.eventUsers.find((user) => user.id === this.viewResultsUserId) || null
    },

    unassignedGifts: function (): Gift[] | null {
      return this.viewResultsUser?.gifts.filter((gift) => !gift.assignedUserId) || null
    },

    assignedGifts: function ():  Gift[] | null {
      return this.viewResultsUser?.gifts.filter((gift) => gift.assignedUserId) || null
    },
  },
  created() {
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
