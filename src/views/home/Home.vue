<template>
  <div id="home">
    <nav-bar></nav-bar>
    <div v-for="(item, index) in banners" :key="index">
      <img style = "width: 100%; height: 200px" :src="item.link" alt="">
    </div>
    <!-- <img :src="url" alt=""> -->
    <!-- <h1> 首页 </h1> -->
  </div>
</template>

<script>
import NavBar from './childComps/NavBar'
import { getHomeMultiData } from 'network/home'

export default {
  name: "Home",
  components: {
    NavBar
  },
  data() {
    return { 
      url: "",
      banners: [],
      recommands: [] 
    }
  },
  created() {
    getHomeMultiData()
    .then(resp => {
      this.banners = resp.data.banner
      this.recommands = resp.data.recommand
    })
    .catch(err => {
      console.log(err)
    })
  }
}
</script>

<style>

</style>