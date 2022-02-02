<template>
  <div>
    <h1>Gifting Results</h1>
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
        <ul class="list-group">
          <li class="list-group-item list-group-item-light" v-for="gift in unassignedGifts" :key="gift.name">
            {{ gift.name }}
            <button
                @click="$emit('assign-gift', {giftId: gift.id, forUserName: this.viewResultsUserName, byUserName: this.currentUserName})"
                class="btn btn-primary">
              Claim Gift
            </button>
          </li>

          <li class="list-group-item list-group-item-dark" v-for="gift in assignedGifts" :key="gift.name">
            {{ gift.name }}
            <template v-if="gift.assignedUserName === this.currentUserName">
              <button class="btn btn-warning"
                      @click="$emit('release-gift', {giftId: gift.id, forUserName: this.viewResultsUserName})">
                Release Gift
              </button>
            </template>
          </li>
        </ul>
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

export default {
  name: "GiftResults",
  components: {SelfGiftResults},
  props: {
    currentUserName: String,
    allUsers: Array
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
      return this.allUsers.find((user) => user.name === this.viewResultsUserName)
    },
    unassignedGifts: function () {
      return this.viewResultsUser.gifts.filter((gift) => !gift.assignedUserName)
    },

    assignedGifts: function () {
      return this.viewResultsUser.gifts.filter((gift) => gift.assignedUserName)
    }
  },
  created() {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
