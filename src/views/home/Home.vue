<template>
  <div id="home">
    <nav-bar></nav-bar>
    <home-swiper></home-swiper>
    <home-recommand></home-recommand>
  </div>
</template>

<script>
import NavBar from './childComps/NavBar'
import HomeSwiper from './childComps/HomeSwiper'
import HomeRecommand from './childComps/HomeRecommand'

export default {
  name: "Home",
  components: {
    NavBar,
    HomeSwiper,
    HomeRecommand
  },
  data() {
    return { 
    }
  }
}
</script>

<style>

</style>