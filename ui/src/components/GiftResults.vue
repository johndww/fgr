<template>
  <div>
    <h1 class="text-center my-3 pb-3">Gifting Results</h1>
    <div class="tabbable tabs-left">
      <ul class="nav nav-tabs">
        <li v-for="user in eventUsers" :key="user.id">
          <button
              :class="user.id === this.viewResultsUserId ? (this.currentUserId === user.id ? 'btn btn-warning' : 'btn btn-primary') : 'btn btn-secondary'"
              @click="this.viewResultsUserId = user.id">{{ user.name }}
          </button>
        </li>
      </ul>
    </div>

    <template v-if="this.viewResultsUser && this.currentUserId !== this.viewResultsUserId">
      <div class="tab-content">
        <table class="table mb-4">
          <tbody>
          <tr v-for="gift in unassignedGifts" :key="gift.name">
            <td>{{ gift.name }}</td>
            <td>
              <button
                  @click="$emit('assign-gift', {giftId: gift.id, forUserId: this.viewResultsUserId, byUserId: this.currentUserId})"
                  class="btn btn-primary">
                Claim Gift
              </button>
            </td>
          </tr>

          <tr v-for="gift in assignedGifts" :key="gift.name">
            <td>{{ gift.name }}</td>
            <template v-if="gift.assignedUserId === this.currentUserId">
              <td>
                <button class="btn btn-warning"
                        @click="$emit('release-gift', {giftId: gift.id, forUserId: this.viewResultsUserId})">
                  Release Gift
                </button>
              </td>
            </template>
          </tr>
          </tbody>
        </table>

      </div>
    </template>

    <SelfGiftResults @add-gift="(giftToAdd) => $emit('add-gift', giftToAdd)"
                     @delete-gift="(giftId) => $emit('delete-gift', giftId)"
                     :viewResultsUser="this.viewResultsUser"
                     v-if="this.viewResultsUser && this.currentUserId === this.viewResultsUserId"/>

  </div>
</template>

<script lang="ts">

import SelfGiftResults from "./SelfGiftResults.vue";
import User from "./App.vue";
import Gift from "./App.vue";
import {defineComponent, PropType} from "vue";


export default defineComponent({
  name: "GiftResults",
  components: {SelfGiftResults},
  props: {
    currentUserId: String,
    eventUsers: Array as PropType<typeof User[]>
  },
  data() {
    return {
      viewResultsUserId: this.currentUserId,
    }
  },
  watch: {
    currentUserId: function () {
      this.viewResultsUserId = this.currentUserId
    }
  },
  methods: {},
  computed: {
    viewResultsUser: function () {
      return this.eventUsers!.find((user: typeof User) => user.id === this.viewResultsUserId)
    },

    unassignedGifts: function () {
      return this.viewResultsUser!.gifts.filter((gift: typeof Gift) => !gift.assignedUserId)
    },

    assignedGifts: function () {
      return this.viewResultsUser!.gifts.filter((gift: typeof Gift) => gift.assignedUserId)
    }
  },
  created() {
  }
})
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
