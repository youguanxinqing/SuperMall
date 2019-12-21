<template>
  <div id="home-swiper">
    <Swiper 
      :swiper-width="styleSwiperSize.width" 
      :swiper-height="styleSwiperSize.height"
    ><SwiperItem 
        v-for="(item, index) in banners" 
        :key="index" 
        :swiper-item-link="item.link" 
        :swiper-item-image="item.image"
      ></SwiperItem>
    </Swiper>
  </div>
</template>

<script>

import { Swiper, SwiperItem } from 'components/common/swiper'

export default {
  name: "HomeSwiper",
  props: {
    banners: {
      type: Array,
      default: []
    }
  },
  components: {
    Swiper,
    SwiperItem
  },
  data() {
    return {
      // 轮播图图片尺寸设置
      styleSwiperSize: {
        width: "360px",
        height: "160px"
      }
    }
  },
}
</script>

<style>

</style>