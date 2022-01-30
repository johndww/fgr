<template>
  <SelectUser msg="Welcome to Your Vue.js + TypeScript App"/>
  <GiftResults></GiftResults>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import SelectUser from './components/SelectUser.vue';
import GiftResults from "./components/GiftResults.vue";

@Options({
  components: {
    GiftResults,
    SelectUser,
  },
})
export default class App extends Vue {}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
