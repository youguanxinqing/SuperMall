<template>
  <div id="tab-control">
    <div
      v-for="(field, index) in fields" 
      @click="itemClick(index)"
      :key="index"
    ><span
      :class="{active: isActive(index)}"
    >{{ field }}</span></div>
  </div>
</template>

<script>
export default {
  name: 'TabControl',
  props: {
    fields: {
        tyepe: Array,
        default: []
    }
  },
  data() {
    return {
      curIndex: 0
    }
  },
  computed: {
    isActive() {
      return (index) => {
        return this.curIndex === index ? true : false
      }
    }
  },
  methods: {
    itemClick(index) {
      this.curIndex = index
    }
  },
}
</script>

<style>
#tab-control {
  display: flex;
  text-align: center; 
  font-size: 15px;
  letter-spacing: 1px;
  height: 40px;
  line-height: 40px;
  background-color: #fff;
}

#tab-control div {
  flex: 1;
}

#tab-control div span {
  padding: 0 5px 5px;
}

.active {
  color: var(--color-high-text);
  border-bottom: 2px solid var(--color-high-text);
}


</style>
