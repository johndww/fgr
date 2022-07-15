<template>
  <div class="tab-content">
    <table class="table mb-4">
      <tbody>
      <tr v-for="gift in myGifts" :key="gift.name">
        <td>{{ gift.name }}</td>
        <td>
          <button class="btn btn-danger" @click="deleteGift(gift.id)">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-trash"
                 viewBox="0 0 16 16">
              <path
                  d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z"/>
              <path fill-rule="evenodd"
                    d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118zM2.5 3V2h11v1h-11z"/>
            </svg>
          </button>
        </td>
      </tr>
      </tbody>
    </table>
    <form @submit="addGift">
      <div class="form-group row">
        <div class="col-xs-3">
          <input type="text" class="form-control" v-model="giftToAdd"/>
        </div>
        <div class="col-xs-4">
          <button type="submit" class="btn btn-primary">Add Gift</button>
        </div>
      </div>
    </form>

  </div>
</template>

<script lang="ts">

import {useCurrentUserId} from "../state/users";
import {GetGiftRequestsState} from "../state/events";
import {computed, defineComponent, ref} from "vue";

export default defineComponent({
  props: ['eventId', 'giftRequestState'],
  emits: ['add-gift', 'delete-gift'],
  setup(props, {emit}) {
    const currentUserId = useCurrentUserId()

    const myGifts = computed(() => {
      const giftRequestState: GetGiftRequestsState = props.giftRequestState
      return giftRequestState.data.filter(request => request.userId == currentUserId.value)
    })

    const giftToAdd = ref("")
    const addGift = function(e: any) {
      e.preventDefault()

      if (!giftToAdd.value) {
        alert("Please enter a gift name")
        return
      }

      //TODO prevent dup gifts

      emit('add-gift', giftToAdd.value)

      giftToAdd.value = ''
    }

    const deleteGift = function (giftIdToDelete: string) {
      if (!confirm("Are you sure you want to remove this item? Others may already have already gotten this for you, but will be notified of this change")) {
        return
      }

      emit('delete-gift', giftIdToDelete)
    }

    return {
      giftToAdd,
      myGifts,
      addGift,
      deleteGift
    }
  }
})
</script>
