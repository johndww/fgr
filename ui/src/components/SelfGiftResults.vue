<template>

  <div class="items">
    <div class="item-container" v-for="gift in myGifts" :key="gift.name">
      <div class="item-delete">
        <div class="item-name">{{ gift.name }}</div>
        <button class="delete-gift-button" @click="deleteGift(gift.id)">Delete Gift</button>
      </div>
      <div class="gift-description">{{gift.description}}</div>
    </div>
  </div>

  <form @submit="addGift">
    <div class="create-gift">
      <input name="name" type="text" class="create-gift-name" placeholder="Gift Name..." v-model="giftToAdd"/>
      <textarea name="description" class="create-gift-description" v-model="description"
                placeholder="Description... (optional)"></textarea>
      <button type="submit" class="submit-button">Add Gift</button>
    </div>
  </form>


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
    const description = ref("")
    const addGift = function(e: any) {
      e.preventDefault()

      if (!giftToAdd.value) {
        alert("Please enter a gift name")
        return
      }

      //TODO prevent dup gifts

      emit('add-gift', giftToAdd.value, description.value)

      giftToAdd.value = ''
      description.value = ''
    }

    const deleteGift = function (giftIdToDelete: string) {
      if (!confirm("Are you sure you want to remove this item? Others may already have already gotten this for you, but will be notified of this change")) {
        return
      }

      emit('delete-gift', giftIdToDelete)
    }

    return {
      giftToAdd,
      description,
      myGifts,
      addGift,
      deleteGift
    }
  }
})
</script>


<style scoped>

.items {
  display: flex;
  flex-direction: column;
  row-gap: 15px;
  margin: 20px 91px 30px;
}

@media (max-width: 600px) {
  .items {
    margin-left: 10px;
  }
}

.item-container{
  min-height: 79px;
  padding-top: 15px;
  background: #FFFFFF 0% 0% no-repeat padding-box;
  border: 1px solid #70707040;
  border-radius: 10px;
}

.item-delete {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-left: 20px;
  padding-right: 20px;
}

.gift-description {
  text-align: left;
  margin-left: 20px;
  font: normal normal bold 12px/22px Proxima Nova;
  white-space: pre-wrap;
}

.delete-gift-button {
  background: #a32323 0% 0% no-repeat padding-box;
  border-radius: 6px;
  width: 177px;
  height: 46px;
  text-align: center;
  font: normal normal bold 16px/19px Proxima Nova;
  letter-spacing: 0px;
  color: #FFFFFF;
  text-transform: uppercase;
  border: none;
  cursor: pointer;
}

.create-gift {
  display: flex;
  flex-direction: column;
  margin: 20px 91px 30px;
  row-gap: 15px;
}

@media (max-width: 600px) {
  .create-gift {
    margin-left: 10px;
  }
}

.create-gift-name {
  margin-right: 38px;
  border: 1px solid #E9E9E9;
  border-radius: 6px;
  width: 725px;
  height: 42px;
  font: normal normal 600 16px/22px Proxima Nova;
  letter-spacing: 0px;
  color: #2F3237;
  padding-left: 13px;
}

.create-gift-description {
  margin-right: 38px;
  border: 1px solid #E9E9E9;
  border-radius: 6px;
  width: 725px;
  height: 100px;
  font: normal normal 600 16px/22px Proxima Nova;
  letter-spacing: 0px;
  color: #2F3237;
  padding-left: 13px;
  resize: none;
  white-space: pre-wrap;
}

</style>