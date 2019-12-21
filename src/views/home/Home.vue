<template>
  <div id="home">
    <nav-bar></nav-bar>
    <home-swiper :banners="banners"></home-swiper>
    <home-recommand :recommands="recommands"></home-recommand>
    <home-feature></home-feature>
  </div>
</template>

<script>
import { getHomeMultiData } from 'network/home'

import NavBar from './childComps/NavBar'
import HomeSwiper from './childComps/HomeSwiper'
import HomeRecommand from './childComps/HomeRecommand'
import HomeFeature from './childComps/HomeFeature'


export default {
  name: "Home",
  components: {
    NavBar,
    HomeSwiper,
    HomeRecommand,
    HomeFeature
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