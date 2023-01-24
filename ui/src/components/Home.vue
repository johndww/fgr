<template>
  <div class="home-content">
    <h1 class="welcome-header">Welcome</h1>
    <p class="welcome-message"><b>SimpleGift</b> is a simple app that lets friends and families organize gifting for
      holidays like Christmas, without duplicating gifts to anyone or spoiling the gift surprise.</p>
    <p>
      To demo the app, <a @click="loginDemo">Click Here</a>
    </p>
    <div>
      <GoogleLogin :callback="gCallback"/>
    </div>
    <img src="../assets/welcome.svg" alt="SimpleGiftApp" width="570" height="235" class="welcome-img">
  </div>
</template>

<script lang="ts">

import {defineComponent, ref} from "vue";
import {useRouter} from "vue-router";
import {loginDemoUser, loginGoogle} from "../state/users";

export default defineComponent({
  name: "Home",
  setup() {
    const router = useRouter()

    const selectUser = function () {
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

    const loginDemo = function () {
      loginDemoUser()
          .then(() => {
            router.push({
              name: "selectevent"
            })
          })
    }

    return {
      selectUser,
      gCallback,
      loginDemo
    }
  }
})
</script>

<style scoped>

.welcome-header {
  font: normal normal bold 46px/56px Proxima Nova;
  color: #233D6E;
}

.welcome-message {
  margin-top: 19px;
  margin-left: auto;
  margin-right: auto;
  width: 642px;
  text-align: center;
  max-width: calc(100% - 40px);
}

.home-content {
  position: relative;
  padding-bottom: 300px;
}

.welcome-img {
  margin-top: 89px;
  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translate(-50%);
}
</style>