<template>
  <div class="select-event-content">
    <div class="event-header">
      <div class="add-events">
        <h1 class="event-header-title">Events</h1>
        <span class="add-event-circle plus"></span>
      </div>
      <div class="present-icon-container">
        <img src="../assets/present_icon.svg" alt="SimpleGiftApp" width="124" height="115">
      </div>
    </div>
    <div>
      <LoadingOrError :loading="myEventsState.loading" :error="myEventsState.error"></LoadingOrError>
    </div>
    <div class="pick-event-container">
      <div class="title-and-date">
        <h2>Pick an event:</h2>
        <div class="date-container">
<!--          <span class="year">Year</span>-->
<!--          <select>-->
<!--            <option>2022</option>-->
<!--          </select>-->
        </div>
      </div>

      <div class="events-container" v-if="!myEventsState.loading && !myEventsState.error">
        <div class="event-container" v-for="event in myEventsState.data" v-on:click="selectEvent(event.id)">
          <div>
            <span class="event-name">{{ event.name }}</span>
            <span class="event-author">by {{ event.ownerName }}</span>
          </div>
          <div class="event-member-count">
            <img src="@/assets/user_icons/user_icon_empty.svg" alt="UserIcon" width="24" height="24" class="user-icon">
            <span class="member-count">{{ event.membershipCount }} Members</span>
          </div>
        </div>
      </div>

    </div>

    <div class="create-event">
      <form @submit="onCreateEvent">
        <input name="name" class="create-event-name" placeholder="Event Name..." v-model="eventNameToCreate">
        <button type="submit" class="submit-button">Create new event</button>
      </form>
    </div>
  </div>

</template>

<script lang="ts">

import {createEvent} from "../state/events";
import {ref} from "vue";
import {useMyEventsState} from "../state/events";
import {useRouter} from "vue-router";
import LoadingOrError from "./LoadingOrError.vue";

export default {
  components: {LoadingOrError},
  setup() {
    const eventNameToCreate = ref("")

    const myEventsState = useMyEventsState()
    myEventsState.value.fetch()

    const router = useRouter()

    const createEventState = ref({
      eventId: '',
      error: '',
      loading: false,
    })

    const onCreateEvent = function (e: any) {
      e.preventDefault()

      if (!eventNameToCreate.value) {
        alert("Please enter an event name")
        return
      }

      //TODO prevent dup names

      const promise = createEvent(eventNameToCreate.value, createEventState)
      promise.then(() => {
        router.push({name: "viewevent", params: {id: createEventState.value.eventId}})
      })
    }

    const selectEvent = function (eventId: string) {
      router.push({name: "viewevent", params: {id: eventId}})
    }

    return {
      eventNameToCreate,
      myEventsState,
      onCreateEvent,
      selectEvent
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.select-event-content {
  padding-bottom: 62px;
}

.add-events {
  margin-left: 91px;
  display: flex;
  align-items: center;
}

.add-event-circle {
  border: 3px solid #89BF60;
  width:31px;
  height:31px;
  border-radius:100%;
  position:relative;
  margin-left:20px;
  display:inline-block;
  vertical-align:middle;
}

.add-event-circle.plus:before,
.add-event-circle.plus:after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: #89BF60;
}

.add-event-circle.plus:before{
  width: 3px;
  margin: 8px auto;
}

.add-event-circle.plus:after{
  margin: auto 8px;
  height: 3px;
}

.pick-event-container {
  margin: 38px 91px 30px;
}

.title-and-date {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.year {
  font: normal normal 600 16px/22px Proxima Nova;
  letter-spacing: 0;
  color: #97989B;
  text-transform: uppercase;
  margin-right: 10px;
}

select {
  color: #2F3237;
  border-width: 1px;
  border-radius: 6px;
  border-color: #E9E9E9;
  font: normal normal 600 16px/22px Proxima Nova;
}

.events-container {
  display: flex;
  flex-direction: column;
  row-gap: 20px;
  margin-bottom: 20px;
}

.event-container {
  padding-left: 20px;
  padding-right: 20px;
  background: #FFFFFF 0% 0% no-repeat padding-box;
  border: 1px solid #70707040;
  border-radius: 10px;
  height:79px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  cursor: pointer;
}

.event-name {
  font: normal normal bold 20px/24px Proxima Nova;
  letter-spacing: 0px;
  color: #3A5892;
  display: block;
  text-align: left;
}

.event-author {
  font: normal normal normal 12px/22px Proxima Nova;
  letter-spacing: 0.3px;
  color: #97989B;
  display: block;
  text-align: left;
}

.member-count {
  font: normal normal normal 16px/22px Proxima Nova;
  letter-spacing: 0px;
  color: #97989B;
}

.create-event-name {
  margin-right: 38px;
  border: 1px solid #E9E9E9;
  border-radius: 6px;
  width: 325px;
  height: 42px;
  font: normal normal 600 16px/22px Proxima Nova;
  letter-spacing: 0px;
  color: #2F3237;
  padding-left: 13px;
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

</style>
