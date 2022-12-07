<template>
  <div class="view-event-contents">
    <div class="event-header">
      <div class="view-event-title">
        <h3 class="event-header-title-back">&lt All Events</h3>
        <h1 class="event-header-title">Wrights 2022 Christmas</h1>
      </div>
      <div>
        <img src="../assets/present_icon.svg" alt="SimpleGiftApp" width="124" height="115">
      </div>
    </div>

    <div>
      <LoadingOrError :loading="this.getEventUsersState.loading"
                      :error="this.getEventUsersState.error"></LoadingOrError>
    </div>

    <div class="event-members">
      <div v-for="(user, index) in getEventUsersState.data" :key="user.id">

        <!-- selected user-->
        <div v-if="user.id === this.viewResultsUserId" class="member-selected-container">
          <div class="member-selected">
            <img src="../assets/user_icon_empty.svg" alt="User" width="28" height="29" class="member-icon">
            <span class="member-name-selected">{{ user.name }}</span>
          </div>
          <div class="user-selected-arrow"></div>
        </div>


        <!-- non-selected user-->
        <div v-if="user.id !== this.viewResultsUserId" class="member" @click="this.viewResultsUserId = user.id">
          <img :src="userIcon(index)" alt="User" width="28" height="29" class="member-icon">
          <span class="member-name">{{ user.name }}</span>
        </div>
      </div>
    </div>

    <SelfGiftResults @add-gift="(giftToAdd) => addGift(giftToAdd)"
                     @delete-gift="(giftId) => deleteGift(giftId)"
                     :eventId="this.event.id" :giftRequestState="giftRequestState"
                     v-if="this.viewResultsUser && currentUserId === this.viewResultsUserId"/>

    <div class="items" v-if="this.viewResultsUser && currentUserId !== this.viewResultsUserId">

      <div class="item-claim" v-for="gift in unassignedGifts" :key="gift.name">
        <div class="item-name">{{ gift.name }}</div>
        <button class="claim-gift-button" @click="assignGift(gift.id, currentUserId)">Claim Gift</button>
      </div>

      <div class="item-release" v-for="gift in assignedGifts" :key="gift.name">
        <div class="item-name">{{ gift.name }}</div>
        <template v-if="gift.isAssignedToMe">
          <button class="release-gift-button" @click="releaseGift(gift.id)">Release Gift</button>
        </template>
      </div>

    </div>

    <div class="delete-edit-event-footer">
      <div class="delete-event">
        <img src="../assets/trash.svg" alt="User" width="18" height="16" class="trash-icon">
        <button class="delete-event-button">Delete Event</button>
      </div>
      <router-link class="edit" :to="{name: 'editevent', params: { id: event.id }}" v-if="this.event.ownerUserId === currentUserId"><button class="edit-event-button">Edit Event</button></router-link>
    </div>

  </div>

</template>

<script lang="ts">

import {computed, ref} from "vue";
import {User} from "../App.vue";
import {getUsersForEvent, useCurrentUserId} from "../state/users";
import {useRoute} from "vue-router";
import {
  getGiftRequests,
  persistClaimGift,
  persistDeleteGiftRequest,
  persistGiftRequest, persistReleaseGift,
  useMyEventsState
} from "../state/events";
import LoadingOrError from "./LoadingOrError.vue";
import {GiftRequest} from "../App.vue";
import SelfGiftResults from "./SelfGiftResults.vue";

export default {
  components: {LoadingOrError, SelfGiftResults},
  setup() {
    const currentUserId = useCurrentUserId()
    const viewResultsUserId = ref(currentUserId.value)

    // only really need current event. could just fetch that. should do this eventually for perf
    const myEventsState = useMyEventsState()
    myEventsState.value.fetch()

    const route = useRoute()

    let eventId = route.params.id as string;
    const event = computed(() => myEventsState.value.data.find((event) => event.id === eventId))

    const getEventUsersState = ref({
      data: <User[]>[],
      error: "",
      loading: false
    })
    getUsersForEvent(eventId, getEventUsersState)

    const viewResultsUser = computed(() => getEventUsersState.value.data.find(user => user.id === viewResultsUserId.value))

    const giftRequestState = ref({
      data: <GiftRequest[]>[],
      loading: false,
      error: ""
    })
    getGiftRequests(eventId, giftRequestState)

    const unassignedGifts = computed(() => {
      return giftRequestState.value.data.filter(request => request.userId == viewResultsUserId.value && !request.isAssigned)
    })

    const assignedGifts = computed(() => {
      return giftRequestState.value.data.filter(request => request.userId == viewResultsUserId.value && request.isAssigned)
    })

    //TODO errors from these should get collated and displayed. need to rethink loading or error
    const claimGiftState = ref({data: "", loading: false, error: ""})
    const claimGift = function(giftId: string, byUserId: string) {
      persistClaimGift(eventId, giftId, byUserId, claimGiftState).finally(() => getGiftRequests(eventId, giftRequestState))
    }

    const addGiftState = ref({data: "", loading: false, error: ""})
    const addGift = function(giftName: string) {
      persistGiftRequest(giftName, eventId, addGiftState).finally(() => getGiftRequests(eventId, giftRequestState))
    }

    const deleteGiftState = ref({data: "", loading: false, error: ""})
    const deleteGift = function(giftId: string) {
      persistDeleteGiftRequest(eventId, giftId, deleteGiftState).finally(() => getGiftRequests(eventId, giftRequestState))
    }

    const releaseGiftState = ref({data: "", loading: false, error: ""})
    const releaseGift = function(giftId: string) {
      persistReleaseGift(eventId, giftId, releaseGiftState).finally(() => getGiftRequests(eventId, giftRequestState))
    }

    const userIcon = function(index: number) {
      //TODO why is src needed..
      const icons = [
          "../src/assets/user_icon_blue.svg",
          "../src/assets/user_icon_pink.svg",
          "../src/assets/user_icon_purple.svg",
          "../src/assets/user_icon_yellow.svg",
      ]

      const iconIdx = index % icons.length
      return icons[iconIdx]
    }

    return {
      viewResultsUserId,
      event,
      getEventUsersState,
      currentUserId,
      assignGift: claimGift,
      addGift,
      deleteGift,
      releaseGift,
      viewResultsUser,
      unassignedGifts,
      assignedGifts,
      giftRequestState,
      userIcon
    }
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->

<style scoped>

.view-event-contents {
  padding-bottom: 62px;
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
}

.event-header-title {
  font: normal normal 600 42px/52px Proxima Nova;
  letter-spacing: 0px;
  color: #FFFFFF;
  display: block;
  margin: 0px;
  text-align: left;
}

.event-members {
  display: flex;
  column-gap: 9px;
  width: 100%;
  margin-top: 34px;
  margin-left: 91px;
}

.member {
  height: 48px;
  padding-left: 11px;
  min-width: 152px;
  padding-right: 11px;
  border: 1px solid #E9E9E9;
  border-radius: 6px;
  display: flex;
  align-items: center;
  cursor: pointer;
}

.member-selected-container {
  position: relative
}

.member-selected {
  height: 48px;
  padding-left: 11px;
  padding-right: 11px;
  min-width: 152px;
  border-radius: 6px;
  display: flex;
  align-items: center;
  background: #89ACF1 0% 0% no-repeat padding-box;
}

.member-icon {
  margin-right: 11.5px;
  text-align: left;
}

.member-name {
  text-align: left;
  font: normal normal 600 19px/26px Proxima Nova;
  letter-spacing: 0px;
  color: #2F3237;
}

.member-name-selected {
  text-align: left;
  font: normal normal 600 19px/26px Proxima Nova;
  letter-spacing: 0px;
  color: #FFFFFF;
}

.user-selected-arrow {
  position: absolute;
  top: 37px;
  left: 0;
  right: 0;
  margin-left: auto;
  margin-right: auto;
  width: 13px;
  height: 13px;
  transform: matrix(0.71, 0.71, -0.71, 0.71, 0, 0);
  background: #89ACF1 0% 0% no-repeat padding-box;
}

.items {
  display: flex;
  flex-direction: column;
  row-gap: 15px;
  margin: 20px 91px 30px;
}

.item-claim {
  height: 79px;
  background: #FFFFFF 0% 0% no-repeat padding-box;
  border: 1px solid #70707040;
  border-radius: 10px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-left: 20px;
  padding-right: 20px;
}

.item-release {
  height: 79px;
  background: #FFFFFF 0% 0% no-repeat padding-box;
  border: 1px solid #89BF60;
  border-radius: 10px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-left: 20px;
  padding-right: 20px;
}

.item-name {
  font: normal normal bold 20px/24px Proxima Nova;
  letter-spacing: 0px;
  color: #2F3237;
}

.claim-gift-button {
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

.release-gift-button {
  width: 177px;
  height: 46px;
  background:none;
  border: none;
  cursor: pointer;
  font: normal normal bold 16px/19px Proxima Nova;
  letter-spacing: 0px;
  color: #89BF60;
  text-transform: uppercase;
  text-align: center;
}

.delete-edit-event-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin: 20px 91px 0px;
}

.delete-event-button {
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

.trash-icon {
  display: inline-block;
  vertical-align: middle;
}

.edit-event-button {
  width: 177px;
  height: 42px;
  border: none;
  background: #89BF60 0% 0% no-repeat padding-box;
  border-radius: 6px;
  text-align: center;
  font: normal normal bold 16px/19px Proxima Nova;
  letter-spacing: 0px;
  color: #FFFFFF;
  text-transform: uppercase;
  cursor: pointer;
}

</style>
