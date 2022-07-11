<template>
  <div>
    <h1 class="text-center my-3 pb-3">{{ this.event.name }} <a class="edit" @click="editEvent" v-if="this.event.ownerUserId === currentUserId.value">(edit)</a></h1>
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
                  @click="assignGift(gift.id, currentUserId.value)"
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
                        @click="releaseGift(gift.id)">
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
                     :eventId="this.event.id"
                     v-if="this.viewResultsUser && currentUserId.value === this.viewResultsUserId"/>

  </div>
</template>

<script lang="ts">

import SelfGiftResults from "./SelfGiftResults.vue";
import {defineComponent} from "vue";
import {
  getEvent,
  getEventUsers,
  getGiftRequests,
  getSessionUserId,
  persistAssignGift,
  persistDeleteGiftRequest,
  persistGiftRequest,
  persistReleaseGift
} from "../state/store";
import {GiftRequest, User} from "../App.vue";

export default defineComponent({
  name: "ViewEvent",
  components: {SelfGiftResults},
  props: {
  },
  data() {
    return {
      viewResultsUserId: getSessionUserId().value,
      event: getEvent(this.$route.query.id)!,
      eventUsers: getEventUsers(this.$route.query.id),
    }
  },
  methods: {
    editEvent() {
      this.$router.push({path: '/editevent', query: {id: this.event.id}})
    },

    assignGift(giftId: string, byUserId: string) {
      persistAssignGift(giftId, byUserId)
    },

    addGift(giftName: string) {
      persistGiftRequest(giftName, this.event.id)
    },

    deleteGift(giftId: string) {
      persistDeleteGiftRequest(giftId)
    },

    releaseGift(giftId: string) {
      persistReleaseGift(giftId)
    }
  },
  computed: {
    currentUserId: function () {
      return getSessionUserId()
    },

    viewResultsUser: function (): User | null {
      return this.eventUsers.find(user => user.id === this.viewResultsUserId) || null
    },

    unassignedGifts: function (): GiftRequest[] | null {
      return getGiftRequests(this.event.id).filter(request => request.userId == this.viewResultsUserId && !request.assignedUserId)
    },

    assignedGifts: function (): GiftRequest[] | null {
      return getGiftRequests(this.event.id).filter(request => request.userId == this.viewResultsUserId  && request.assignedUserId)
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
