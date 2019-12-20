<template>
  <div id="swiper-item">
    <a :href="swiperItemLink">
      <img :src="swiperItemImage" :style="styleImg" alt="">
    </a>
  </div>
</template>

<script>
export default {
  name: "SwiperItem",
  props: {
    swiperItemLink: String,
    swiperItemImage: String,
  },
  data() {
    return {
      itemWidth: String,
      itemHeight: String
    }
  },
  computed: {
    styleImg() {
      return {
        width: this.itemWidth,
        height: this.itemHeight
      }
    }
  },
  created() {
    // 增加一个组件, dotNum + 1
    this.$parent.$emit("add-dot-num")
    // ...
    this.itemWidth = this.$parent.swiperWidth
    this.itemHeight = this.$parent.swiperHeight
  }
}
</script>

<style>

</style>