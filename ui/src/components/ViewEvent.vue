<template>
  <LoadingOrError :loading="this.getEventUsersState.loading" :error="this.getEventUsersState.error"></LoadingOrError>
  <div v-if="this.event">
    <h1 class="text-center my-3 pb-3">{{ this.event.name }} <router-link class="edit" :to="{name: 'editevent', params: { id: event.id }}" v-if="this.event.ownerUserId === currentUserId">(edit)</router-link></h1>
    <div class="tabbable tabs-left">
      <ul class="nav nav-tabs">
        <li v-for="user in getEventUsersState.data" :key="user.id">
          <button
              :class="user.id === this.viewResultsUserId ? (this.currentUserId === user.id ? 'btn btn-warning' : 'btn btn-primary') : 'btn btn-secondary'"
              @click="this.viewResultsUserId = user.id">{{ user.name }}
          </button>
        </li>
      </ul>
    </div>

    <template v-if="this.viewResultsUser && currentUserId !== this.viewResultsUserId">
      <div class="tab-content">
        <table class="table mb-4">
          <tbody>
          <tr v-for="gift in unassignedGifts" :key="gift.name">
            <td>{{ gift.name }}</td>
            <td>
              <button
                  @click="assignGift(gift.id, currentUserId)"
                  class="btn btn-primary">
                Claim Gift
              </button>
            </td>
          </tr>

          <tr v-for="gift in assignedGifts" :key="gift.name">
            <td>{{ gift.name }}</td>
            <template v-if="gift.isAssignedToMe">
              <td>
                <button class="btn btn-warning"
                        @click="releaseGift(gift.id)">
                  Release Gift
                </button>
              </td>
            </template>
          </tr>
          </tbody>
        </table>

      </div>
    </template>

    <SelfGiftResults @add-gift="(giftToAdd) => addGift(giftToAdd)"
                     @delete-gift="(giftId) => deleteGift(giftId)"
                     :eventId="this.event.id" :giftRequestState="giftRequestState"
                     v-if="this.viewResultsUser && currentUserId === this.viewResultsUserId"/>

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
      giftRequestState
    }
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->

<style scoped>

.edit {
  color: #42b983;
  font-size: 12px;
}
</style>
