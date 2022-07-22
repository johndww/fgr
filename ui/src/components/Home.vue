<template>
<div>
  <p>This is a simple app that lets friends and families organize gifting for holidays like Christmas, without
  duplicating gifts to anyone or spoiling the gift surprise.</p>

  <p>To get started, login:</p>
  <button type="button" @click="toLogin()" class="btn btn-primary btn-block mb-4">Login</button>
</div>
</template>

<script lang="ts">

import {defineComponent} from "vue";
import {useRouter} from "vue-router";

export default defineComponent({
  name: "Home",
  setup() {
    const router = useRouter()

    const selectUser = function() {
      router.push({
        name: "selectuser"
      })
    }

    const toLogin = function () {
      router.push({
        name: 'login',
      })
    }

    return {
      selectUser,
      toLogin
    }
  }
})
</script>

<style scoped>

</style>