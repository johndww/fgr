<template>
  <div class="view-event-contents">
    <div class="event-header">
      <div class="view-event-title">
        <router-link class="event-header-title-back" :to="{name: 'selectevent'}">&lt All Events</router-link>
        <h1 class="event-header-title">{{ this.event.name }}</h1>
      </div>
      <div class="present-icon-container">
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
            <img src="@/assets/user_icons/user_icon_empty.svg" alt="User" width="28" height="29" class="member-icon">
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

    <SelfGiftResults @add-gift="(giftToAdd, description) => addGift(giftToAdd, description)"
                     @delete-gift="(giftId) => deleteGift(giftId)"
                     :eventId="this.event.id" :giftRequestState="giftRequestState"
                     v-if="this.viewResultsUser && currentUserId === this.viewResultsUserId"/>

    <div class="items" v-if="this.viewResultsUser && currentUserId !== this.viewResultsUserId">

      <div class="item-container" v-for="gift in unassignedGifts" :key="gift.name">
        <div class="item-claim">
          <div class="item-name">{{ gift.name }}</div>
          <button class="button" @click="assignGift(gift.id, currentUserId)" :disabled="this.claimGiftState.loading">
            <span v-if="this.claimGiftState.loading && this.claimGiftState.giftId === gift.id">Claiming...</span>
            <span v-if="!this.claimGiftState.loading || this.claimGiftState.giftId !== gift.id">Claim Gift</span>
          </button>
        </div>
        <div class="gift-description">{{ gift.description }}</div>
      </div>

      <div class="item-container" v-for="gift in assignedGifts" :key="gift.name">
        <div class="item-release">
          <div class="item-name">{{ gift.name }}</div>
          <template v-if="gift.isAssignedToMe">
            <button class="release-gift-button" @click="releaseGift(gift.id)" :disabled="this.releaseGiftState.loading">
              <span v-if="this.releaseGiftState.loading && this.releaseGiftState.giftId === gift.id">Releasing...</span>
              <span v-if="!this.releaseGiftState.loading || this.releaseGiftState.giftId !== gift.id">Release Gift</span>
            </button>
          </template>
        </div>
        <div class="gift-description">{{ gift.description }}</div>
      </div>

    </div>

    <div class="delete-edit-event-footer" v-if="this.event.ownerUserId === currentUserId">
      <div class="delete-event">
        <img src="../assets/trash.svg" alt="User" width="18" height="16" class="trash-icon">
        <button class="delete-button" @click="deleteEvent(event.id)">Delete Event</button>
      </div>
      <router-link class="edit" :to="{name: 'editevent', params: { id: event.id }}"><button class="button">Edit Event</button></router-link>
    </div>

    <div>
      <div>
        <!-- footer -->
        <ins class="adsbygoogle"
             style="display:block"
             data-ad-client="ca-pub-3785547202124831"
             data-ad-slot="4497174860"
             data-ad-format="auto"
             data-full-width-responsive="true"></ins>
      </div>
    </div>
  </div>

</template>

<script lang="ts">

import {computed, ref} from "vue";
import {User} from "../App.vue";
import {getUsersForEvent, useCurrentUserId} from "../state/users";
import {useRoute, useRouter} from "vue-router";
import {
  getGiftRequests,
  persistClaimGift, persistDeleteEvent,
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

    const router = useRouter()

    const deleteEventState = ref({data: "", loading: false, error: ""})
    const deleteEvent = function(eventId: string) {
      if (confirm("Are you sure you want to delete this event?")) {
        persistDeleteEvent(eventId, deleteEventState).finally(() => {
          console.log("event deleted. redirecting to selectevent")
          router.push({ name: 'selectevent'})
        })
      }
    }

    //TODO errors from these should get collated and displayed. need to rethink loading or error
    const claimGiftState = ref({data: "", loading: false, giftId: "", error: ""})
    const claimGift = function(giftId: string, byUserId: string) {
      persistClaimGift(eventId, giftId, byUserId, claimGiftState).finally(() => getGiftRequests(eventId, giftRequestState))
    }

    const addGiftState = ref({data: "", loading: false, error: ""})
    const addGift = function(giftName: string, description: string) {
      persistGiftRequest(giftName, description, eventId, addGiftState).finally(() => getGiftRequests(eventId, giftRequestState))
    }

    const deleteGiftState = ref({data: "", loading: false, error: ""})
    const deleteGift = function(giftId: string) {
      persistDeleteGiftRequest(eventId, giftId, deleteGiftState).finally(() => getGiftRequests(eventId, giftRequestState))
    }

    const releaseGiftState = ref({data: "", loading: false, giftId: "", error: ""})
    const releaseGift = function(giftId: string) {
      persistReleaseGift(eventId, giftId, releaseGiftState).finally(() => getGiftRequests(eventId, giftRequestState))
    }

    const userIcon = function (index: number) {
      const icons = [
        "user_icon_blue.svg",
        "user_icon_pink.svg",
        "user_icon_purple.svg",
        "user_icon_yellow.svg",
      ]

      const iconIdx = index % icons.length
      return new URL(`/src/assets/user_icons/${icons[iconIdx]}`, import.meta.url).href
    }

    return {
      viewResultsUserId,
      event,
      getEventUsersState,
      currentUserId,
      assignGift: claimGift,
      deleteEvent,
      addGift,
      deleteGift,
      releaseGift,
      viewResultsUser,
      unassignedGifts,
      assignedGifts,
      giftRequestState,
      userIcon,
      releaseGiftState,
      claimGiftState
    }
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->

<style scoped>

.view-event-contents {
  padding-bottom: 62px;
}

.event-members {
  display: flex;
  flex-wrap: wrap;
  column-gap: 9px;
  margin-top: 34px;
  margin-left: 91px;
  margin-right: 91px;
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

.item-container{
  min-height: 79px;
  padding-top: 15px;
  background: #FFFFFF 0% 0% no-repeat padding-box;
  border: 1px solid #70707040;
  border-radius: 10px;
}

.gift-description {
  text-align: left;
  margin-left: 20px;
  font: normal normal bold 12px/22px Proxima Nova;
  white-space: pre-wrap;
}

.item-claim {
  height: 79px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-left: 20px;
  padding-right: 20px;
}

.item-release {
  height: 79px;
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
  margin: 20px 91px 10px;
}

.trash-icon {
  display: inline-block;
  vertical-align: middle;
}

</style>
