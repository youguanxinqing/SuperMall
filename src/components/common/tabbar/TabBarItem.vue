<template>
  <div id="tab-bar-item">
    <div v-if="isActive" class="img-box">
      <slot name="item-active"></slot>
    </div>
    <div v-else class="img-box">
      <slot name="item-deactive"></slot>
    </div>
    <div>
      <slot name="item-text"></slot>
    </div>
  </div>
</template>

<script>
export default {
  name: "TabBarItem",
  computed: {
    isActive() {
      return true
    }
  }
}
</script>

<style>

.img-box {
  
}

</style>