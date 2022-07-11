<template>
  <div>
    <h1 class="text-center my-3 pb-3">Select Event</h1>

    <ul>
      <li v-for="event in myEvents" :key="event.id"><button @click="selectedEvent(event.id)">{{ event.name }}</button></li>
    </ul>
    <form @submit="onCreateEvent">
      <input name="name" v-model="eventNameToCreate"> <button type="submit">Create Event</button>
    </form>
  </div>
</template>

<script lang="ts">

import {defineComponent} from "vue";
import {createEvent, getMyEvents} from "../state/store";

export default defineComponent({
  name: "SelectEvent",
  props: {
  },
  data() {
    return {
      eventNameToCreate: ''
    }
  },
  computed: {
    myEvents: function () {
      return getMyEvents()
    }
  },
  methods: {
    onCreateEvent(e: any) {
      e.preventDefault()

      if (!this.eventNameToCreate) {
        alert("Please enter an event name")
        return
      }

      //TODO prevent dup names

      const eventId = createEvent(this.eventNameToCreate)
      this.$router.push({name: "viewevent", params: { id: eventId }})
    },

    selectedEvent(eventId: string) {
      this.$router.push({name: "viewevent", params: { id: eventId }})
    }
  }
})
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
