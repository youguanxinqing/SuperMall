<template>
  <div id="tab-bar-item" @click="clickItem">
    <div v-if="isActive" class="img-box">
      <slot name="item-active"></slot>
    </div>
    <div v-else class="img-box">
      <slot name="item-deactive"></slot>
    </div>
    <div :style="styleTextColor">
      <slot name="item-text"></slot>
    </div>
  </div>
</template>

<script>
export default {
  name: "TabBarItem",
  props: {
    textColor: {
      type: String,
      default: "red"
    },
    path: String
  },
  computed: {
    isActive() {
      return this.$route.path.indexOf(this.path) !== -1
    },
    styleTextColor() {
      return this.isActive === true ? {color: this.textColor} : {}
    }
  },
  methods: {
    clickItem() {
      this.$router.replace(this.path)
    }
  }
}
</script>

<style>

#tab-bar-item {
  flex: 1;
  text-align: center;
}

.img-box img {
  width: 44%;
}

</style>