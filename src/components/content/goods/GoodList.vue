<template>
  <div id="good-list">
    <good-item 
      v-for="(good, index) in goods" 
      :key="index"
      :good-item="good"
    />
  </div>
</template>

<script>

import GoodItem from './GoodItem'

export default {
  name: 'GoodList',
  props: {
    goods: {
      type: Array,
      default: []
    }
  },
  components: {
    GoodItem
  },
  data() {
    return {

    }
  },
}
</script>

<style>

#good-list {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
}

</style>
