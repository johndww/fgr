<template>
  <h1>Welcome</h1>
  <p class="welcome-message"><b>SimpleGift</b> is a simple app that lets friends and families organize gifting for holidays like Christmas, without duplicating gifts to anyone or spoiling the gift surprise.</p>
  <div>
    <GoogleLogin :callback="gCallback" />
  </div>
  <img src="../assets/welcome.svg" alt="SimpleGiftApp" width="570" height="235" class="welcome-img">
</template>

<script lang="ts">

import {defineComponent} from "vue";
import {useRouter} from "vue-router";
import {loginGoogle} from "../state/users";

export default defineComponent({
  name: "Home",
  setup() {
    const router = useRouter()

    const selectUser = function() {
      router.push({
        name: "selectuser"
      })
    }

    const gCallback = function (resp: any) {
      loginGoogle(resp.credential)
          .then(() => {
            router.push({
              name: "selectevent"
            })
          })
    }

    return {
      selectUser,
      gCallback
    }
  }
})
</script>

<style scoped>

.welcome-message {
  margin-top: 19px;
  margin-left: auto;
  margin-right: auto;
  width: 642px;
  text-align: center;
  max-width: calc(100% - 40px);
}

.welcome-img {
  margin-top: 89px;
  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translate(-50%);

}
</style>