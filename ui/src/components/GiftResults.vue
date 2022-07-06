<template>
  <div>
    <h1 class="text-center my-3 pb-3">Gifting Results</h1>
    <div class="tabbable tabs-left">
      <ul class="nav nav-tabs">
        <li v-for="user in allUsers" :key="user.name">
          <button
              :class="user.name === this.viewResultsUserName ? (this.currentUserName === user.name ? 'btn btn-warning' : 'btn btn-primary') : 'btn btn-secondary'"
              @click="this.viewResultsUserName = user.name">{{ user.name }}
          </button>
        </li>
      </ul>
    </div>

    <template v-if="this.viewResultsUser && this.currentUserName !== this.viewResultsUserName">
      <div class="tab-content">
        <table class="table mb-4">
          <tbody>
          <tr v-for="gift in unassignedGifts" :key="gift.name">
            <td>{{ gift.name }}</td>
            <td>
              <button
                  @click="$emit('assign-gift', {giftId: gift.id, forUserName: this.viewResultsUserName, byUserName: this.currentUserName})"
                  class="btn btn-primary">
                Claim Gift
              </button>
            </td>
          </tr>

          <tr v-for="gift in assignedGifts" :key="gift.name">
            <td>{{ gift.name }}</td>
            <template v-if="gift.assignedUserName === this.currentUserName">
              <td>
                <button class="btn btn-warning"
                        @click="$emit('release-gift', {giftId: gift.id, forUserName: this.viewResultsUserName})">
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
                     v-if="this.viewResultsUser && this.currentUserName === this.viewResultsUserName"/>

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
    currentUserName: String,
    allUsers: Array as PropType<typeof User[]>
  },
  data() {
    return {
      viewResultsUserName: this.currentUserName,
    }
  },
  watch: {
    currentUserName: function () {
      this.viewResultsUserName = this.currentUserName
    }
  },
  methods: {},
  computed: {
    viewResultsUser: function () {
      return this.allUsers!.find((user: typeof User) => user.name === this.viewResultsUserName)
    },
    unassignedGifts: function () {
      return this.viewResultsUser!.gifts.filter((gift: typeof Gift) => !gift.assignedUserName)
    },

    assignedGifts: function () {
      return this.viewResultsUser!.gifts.filter((gift: typeof Gift) => gift.assignedUserName)
    }
  },
  created() {
  }
})
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
