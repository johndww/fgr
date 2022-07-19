<template>
  <div>
    <GoogleLogin :callback="gCallback" />
  </div>
</template>

<script>

import {loginGoogle} from "../state/users";
import {useRouter} from "vue-router";

export default {
  name: "Login",
  setup() {
    const router = useRouter()
    const gCallback = function (resp) {
      loginGoogle(resp.credential)
          .then(() => {
            router.push({
              name: "selectevent"
            })
          })
    }

    return {
      gCallback
    }
  }
}
</script>

<style scoped>

</style>