<template>
  <div class="main-container">
    <div class="header">
      <div>
        <img src="./assets/logo.svg" alt="SimpleGiftApp" width="35" height="35" class="logo">
        <div class="header-name">SimpleGift</div>
      </div>
      <div v-show="currentUserState.data" class="logged-in-user">
        <img src="@/assets/user_icons/user_icon_blue.svg" alt="UserIcon" width="24" height="24" class="user-icon">
        <a @click="logout" class="user">{{ currentUserState.data != null && currentUserState.data.name }}</a>
        <a v-if="currentUserState.data != null && currentUserState.data.admin" @click="selectUser"
           class="admin">Admin</a>
      </div>
    </div>
    <div class="content">
      <router-view></router-view>
    </div>
  </div>

</template>

<script lang="ts">
import {defineComponent, ref} from "vue";
import {logoutUser, useCurrentUserState} from "./state/users";
import {useRouter} from "vue-router";

export interface User {
  id: string,
  name: string,
  email: string,
  admin: boolean,
}

export interface Event {
  id: string,
  name: string,
  ownerUserId: string
}

export interface GiftRequest {
  id: string
  userId: string,
  eventId: string,
  name: string
  isAssigned: boolean
  isAssignedToMe: boolean
}

export default defineComponent({
  name: 'App',
  meta: {
    title: "Simple Gift App",
  },
  setup() {
    const currentUserState = useCurrentUserState()

    const router = useRouter()

    const logoutState = ref({data: "", error: "", loading: false})
    const logout = function () {
      return logoutUser(logoutState).then(() => {
        console.log("logged user out")
        router.push({name: "home"})
      })
    }

    const selectUser = function () {
      console.log("admin moving to selectuser")
      router.push({name: "selectuser"})
    }

    return {
      currentUserState,
      logout,
      selectUser
    }
  }
})
</script>

<style>
body {
  background: transparent linear-gradient(180deg, #F7FBFE 0%, #E6F3FC 100%) 0% 0% no-repeat padding-box;
  color: #2F3237;
  font: normal normal normal 22px/32px Proxima Nova;
  min-height: 100vh;
  margin: 0;
  padding: 0;
}

h1 {
  color: #FFFFFF;
  font: normal normal 600 42px/52px Proxima Nova;
}

h2 {
  color: #2F3237;
  font: normal normal bold 32px/39px Proxima Nova;
}

.main-container {
  margin: 0 auto;
  width: 1119px;
  background: #FFFFFF;
  box-shadow: 0px 0px 20px #0000000D;
  border-radius: 10px;
  max-width: calc(100% - 40px);
}

.header {
  height: 77px;
  width: 100%;
  background: #FBFDFF 0% 0% no-repeat padding-box;
  border-radius: 10px 10px 0px 0px;
  box-shadow: 0px 0px 20px #0000000D;
  line-height: 77px;
  display: flex;
  justify-content: space-between;
}

.logo {
  display: inline-block;
  vertical-align: middle;
  margin-right: 8px;
  margin-left: 25px;
}

.header-name {
  font: normal normal bold 26px/31px Proxima Nova;
  letter-spacing: 0px;
  color: #2F3237;
  display: inline-block;
  vertical-align: middle;
}

.user-icon {
  margin-right: 10px;
  vertical-align: middle;
}

.logged-in-user {
  display: inline-block;
  vertical-align: middle;
}

.user {
  vertical-align: middle;
  margin-right: 22px;
}

.admin {
  vertical-align: middle;
  margin-right: 22px;
}

.content {
  text-align: center;
  width: 100%;
  border-radius: 10px;
}

/* for subpages */
.event-header {
  width: 100%;
  min-height: 115px;
  background: #233D6E 0% 0% no-repeat padding-box;
  display: flex;
  justify-content: space-between;
}

.event-header-title {
  display: inline-block;
  vertical-align: middle;
  margin: 0px;
}

.present-icon-container {
  display: flex;
  align-items: flex-end;
}

.view-event-title {
  margin-left: 91px;
  padding-top: 26px;
}

.event-header-title-back {
  font: normal normal bold 14px/17px Proxima Nova;
  letter-spacing: 0.7px;
  color: #89BF60;
  text-transform: uppercase;
  display: block;
  margin: 0px;
  text-align: left;
  text-decoration: none;
}

.delete-button {
  border: none;
  background: none;
  width: 120px;
  height: 42px;
  font: normal normal bold 14px/17px Proxima Nova;
  letter-spacing: 0.7px;
  text-align: left;
  color: #89BF60;
  text-transform: uppercase;
  cursor: pointer;
  display: inline-block;
  vertical-align: middle;
}

.button {
  background: #89BF60 0% 0% no-repeat padding-box;
  border-radius: 6px;
  width: 177px;
  height: 46px;
  text-align: center;
  font: normal normal bold 16px/19px Proxima Nova;
  letter-spacing: 0px;
  color: #FFFFFF;
  text-transform: uppercase;
  border: none;
  cursor: pointer;
}

.submit-button {
  width: 266px;
  height: 46px;
  background: #89BF60 0% 0% no-repeat padding-box;
  border-width: 0px;
  border-radius: 6px;
  text-align: center;
  font: normal normal bold 16px/19px Proxima Nova;
  letter-spacing: 0px;
  color: #FFFFFF;
  text-transform: uppercase;
  cursor: pointer;
}

.border-only-button {
  border: 2px solid #89BF60;
  border-radius: 6px;
  width: 177px;
  height: 42px;
  text-align: center;
  font: normal normal bold 16px/19px Proxima Nova;
  letter-spacing: 0px;
  color: #89BF60;
  text-transform: uppercase;
  background: none;
  cursor: pointer;
}

/* end for subpages */

#app {
  font-family: 'Proxima Nova', Georgia, sans-serif;
  padding: 60px 0;
  box-sizing: border-box;
  min-height: 100vh;
}
</style>
