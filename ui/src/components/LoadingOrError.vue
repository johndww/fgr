<template>
  <div>
    <div v-if="error">
      {{ error }}
    </div>
    <div v-if="loading">
      Loading...
    </div>
  </div>
</template>

<script lang="ts">
import {computed} from "vue";

export default {
  props: ['error', 'loading'],
  setup(props: any) {
    const error = computed(() => props.error)
    const loading = computed(() => props.loading)
    return {
      error,
      loading
    }
  }
}
</script>

<style scoped>

</style>