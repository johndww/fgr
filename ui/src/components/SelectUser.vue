<template>
  <div>
    <h1>Who are you?</h1>
    <ul>
      <li><button>John</button></li>
      <li><button>John2</button></li>
      <li><button>John3</button></li>
      <li><button>John4</button></li>
    </ul>
    <input name="name"> <button>Join Registry</button>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';

@Options({
  props: {
    msg: String
  }
})
export default class SelectUser extends Vue {
  msg!: string
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
