<template>
<div>
  <p>This is a simple app that lets friends and families organize gifting for holidays like Christmas, without
  duplicating gifts to anyone or spoiling the gift surprise.</p>

  <p>To get started, login:</p>
  <button type="button" @click="toHome()" class="btn btn-primary btn-block mb-4">Login</button>
</div>
</template>

<script lang="ts">

import {defineComponent} from "vue";

export default defineComponent({
  name: "Home",
  methods: {
    toHome() {
      this.$router.push({
        path: '/login',
      })
    }
  }
})
</script>

<style scoped>

</style>