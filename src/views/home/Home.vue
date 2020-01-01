<template>
  <div id="home">
    <nav-bar></nav-bar>
    <home-swiper :banners="banners"></home-swiper>
    <home-recommand :recommands="recommands"></home-recommand>
    <home-feature></home-feature>
    <home-tab-control class="position-tab"></home-tab-control>
    
    <!-- 伪造数据 -->
    <div v-for="(x, i) in Array(fakeNum)" :key="i">{{ i }}</div>
  </div>
</template>

<script>
import { getHomeMultiData } from 'network/home'

import NavBar from './childComps/NavBar'
import HomeSwiper from './childComps/HomeSwiper'
import HomeRecommand from './childComps/HomeRecommand'
import HomeFeature from './childComps/HomeFeature'
import HomeTabControl from './childComps/HomeTabControl'


export default {
  name: "Home",
  components: {
    NavBar,
    HomeSwiper,
    HomeRecommand,
    HomeFeature,
    HomeTabControl
  },
  data() {
    return { 
      banners: [],
      recommands: [],
      
      fakeNum: 100,  /* 伪造数据 */
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

/* 临时用法 */
.position-tab {
  position: sticky;
  top: 0px;
}

</style>