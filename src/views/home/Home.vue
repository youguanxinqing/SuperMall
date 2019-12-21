<template>
  <div id="home">
    <nav-bar></nav-bar>
    <home-swiper :banners="banners"></home-swiper>
    <home-recommand :recommands="recommands"></home-recommand>
  </div>
</template>

<script>
import NavBar from './childComps/NavBar'
import HomeSwiper from './childComps/HomeSwiper'
import HomeRecommand from './childComps/HomeRecommand'
import { getHomeMultiData } from 'network/home'

export default {
  name: "Home",
  components: {
    NavBar,
    HomeSwiper,
    HomeRecommand
  },
  data() {
    return { 
      banners: [],
      recommands: []
    }
  },
  created() {
    getHomeMultiData()
    .then(resp => {
      this.banners = resp.data.banners
      this.recommands = resp.data.recommands
    })
    .catch(err => {
      console.log(err)
    })
  }
}
</script>

<style>

</style>