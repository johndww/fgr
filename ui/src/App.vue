<template>
  <div class="main-container">
    <div class="header">
      <div><img src="./assets/logo.svg" alt="SimpleGiftApp" width="35" height="35" class="logo">
      <div class="header-name">SimpleGift</div>
      </div>
      <div v-show="currentUserState.data" class="logged-in-user">
        <img src="./assets/user_icon_blue.svg" alt="UserIcon" width="24px" height="24px" class="user-icon">
        <a @click="logout" class="user">{{ currentUserState.data != null && currentUserState.data.name }}</a>
        <a v-if="currentUserState.data != null && currentUserState.data.admin" @click="selectUser">Admin</a>
      </div>
<!--      <div v-show="currentUserState.data" style="text-align: right">-->
<!--        <a @click="logout">{{ currentUserState.data != null && currentUserState.data.name }}</a><br/>-->
<!--        <a v-if="currentUserState.data != null && currentUserState.data.admin" @click="selectUser">Admin</a>-->
<!--      </div>-->
    </div>
    <div class="content">
      <router-view></router-view>
    </div>
  </div>


  <!--  <div class="vh-100" style="background-color: #3da2c3;">-->
<!--    <div class="container py-5 h-100">-->
<!--      <div class="row d-flex justify-content-center align-items-center h-100">-->
<!--        <div class="col col-lg-9 col-xl-7">-->
<!--          <div class="card rounded-3">-->
<!--            <div class="card-body p-4">-->

<!--              <div v-show="currentUserState.data">-->
<!--                <nav class="navbar navbar-light bg-light">-->
<!--                  <div class="container-fluid">-->
<!--                    <router-link :to="{name: 'selectevent'}" class="navbar-brand" >Events</router-link>-->

<!--                    <div style="text-align: right">-->
<!--                      <a @click="logout" >{{ currentUserState.data != null && currentUserState.data.name }}</a><br />-->
<!--                      <a v-if="currentUserState.data != null && currentUserState.data.admin" @click="selectUser">Admin</a>-->
<!--                    </div>-->
<!--                  </div>-->
<!--                </nav>-->
<!--              </div>-->

<!--              <router-view></router-view>-->
<!--            </div>-->
<!--          </div>-->
<!--        </div>-->
<!--      </div>-->
<!--    </div>-->
<!--  </div>-->

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

.main-container {
  margin: 0 auto;
  width: 1119px;
  background: #FFFFFF;
  box-shadow: 0px 0px 20px #0000000D;
  border-radius: 10px;
  max-width: calc(100% - 40px);
}

.header {
  /* Layout Properties */
  height: 77px;
  width: 100%;
  /* UI Properties */
  background: #FBFDFF 0% 0% no-repeat padding-box;
  border-radius: 10px 10px 0px 0px;
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

.content {
  text-align: center;
  padding-top: 71px;
  width: 100%;
  box-shadow: 0px 0px 20px #0000000D;
  border-radius: 10px;
  position: relative;
  padding-bottom: 300px;
}

h1 {
  font: normal normal bold 46px/56px Proxima Nova;
  color: #233D6E;
}

#app {
  font-family: 'Proxima Nova', Georgia, sans-serif;
  padding: 60px 0;
  box-sizing: border-box;
  min-height: 100vh;
  /*-webkit-font-smoothing: antialiased;*/
  /*-moz-osx-font-smoothing: grayscale;*/
  /*text-align: center;*/
  /*color: #2c3e50;*/
}
</style>
