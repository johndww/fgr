<template>
  <div>
    <h1>Gifting Results</h1>
    <div class="tabbable tabs-left">
      <ul class="nav nav-tabs">
        <li>
          <button>My List</button>
        </li>
        <li>
          <button>Sue</button>
        </li>
        <li>
          <button>Bruce</button>
        </li>
        <li>
          <button>Peter</button>
        </li>
        <li>
          <button>Chelsea</button>
        </li>
        <li>
          <button>Haritha</button>
        </li>
      </ul>
    </div>
    <div class="tab-content">

    </div>
  </div>
</template>

<script lang="ts">
import {Options, Vue} from 'vue-class-component';

@Options({
  props: {}
})
export default class GiftResults extends Vue {
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
