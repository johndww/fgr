<template>
  <div>
    <h1 class="text-center my-3 pb-3">Select Event</h1>

    <div v-if="!myEventsState.loading && !myEventsState.error">
    <ul>
      <li v-for="event in myEventsState.data" :key="event.id"><button @click="selectEvent(event.id)">{{ event.name }}</button></li>
    </ul>
    <form @submit="onCreateEvent">
      <input name="name" v-model="eventNameToCreate"> <button type="submit">Create Event</button>
    </form>
    </div>
    <LoadingOrError :loading="myEventsState.loading" :error="myEventsState.error"></LoadingOrError>
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
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
