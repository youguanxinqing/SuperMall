<template>
  <div id="home">
    <nav-bar></nav-bar>
    <home-swiper></home-swiper>
  </div>
</template>

<script>
import NavBar from './childComps/NavBar'
import HomeSwiper from './childComps/HomeSwiper'

export default {
  name: "Home",
  components: {
    NavBar,
    HomeSwiper
  },
  data() {
    return { 
    }
  }
}
</script>

<style>

</style>