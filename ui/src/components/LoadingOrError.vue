<template>
  <div>
    <div v-if="errorc">
      {{ errorc }}
    </div>
    <div v-if="loadingc">
      Loading...
    </div>
  </div>
</template>

<script lang="ts">
import {computed} from "vue";

export default {
  props: ['error', 'loading'],
  setup(props: any) {
    const errorc = computed(() => props.error)
    const loadingc = computed(() => props.loading)
    return {
      errorc,
      loadingc
    }
  }
}
</script>

<style scoped>

</style>