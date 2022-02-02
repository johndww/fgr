<template>
  <div>
    <h1>Gifting Results</h1>
    <div class="tabbable tabs-left">
      <ul class="nav nav-tabs">
        <li v-for="user in allUsers" :key="user.name">
          <button :class="user.name === this.viewResultsUserName ? 'btn btn-primary' : 'btn btn-secondary'"
                  @click="this.viewResultsUserName = user.name">{{ user.name }}
          </button>
        </li>
      </ul>
    </div>
    <template v-if="this.viewResultsUser">
      <div class="tab-content" v-show="this.viewResultsUser">
        <ul class="list-group">
          <li class="list-group-item list-group-item-light" v-for="gift in this.viewResultsUser.gifts" :key="gift.name">
            {{ gift.name }}
          </li>
        </ul>
      </div>
    </template>
  </div>
</template>

<script lang="ts">

export default {
  name: "GiftResults",
  props: {
    currentUserName: String,
    allUsers: Array
  },
  data() {
    return {
      viewResultsUserName: this.currentUserName
    }
  },
  watch: {
    currentUserName: function() {
        this.viewResultsUserName = this.currentUserName
    }
  },
  methods: {
  },
  computed: {
    viewResultsUser: function() {
      return this.allUsers.find((user) => user.name === this.viewResultsUserName)
    }
  },
  created() {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
