<template>
  <div id="home">
    <nav-bar></nav-bar>
    <home-swiper :banners="banners"></home-swiper>
    <home-recommand :recommands="recommands"></home-recommand>
    <home-feature></home-feature>
    <home-tab-control class="position-tab"></home-tab-control>
    <good-list :goods="goods['pop'].list"/>
    
    <!-- 伪造数据 -->
    <div v-for="(x, i) in Array(fakeNum)" :key="i">{{ i }}</div>
  </div>
</template>

<script>
import { 
  getHomeMultiData,
  getHomeGoods,
} from 'network/home'
import GoodList from 'components/content/goods/GoodList'

import NavBar from './childComps/NavBar'
import HomeSwiper from './childComps/HomeSwiper'
import HomeRecommand from './childComps/HomeRecommand'
import HomeFeature from './childComps/HomeFeature'
import HomeTabControl from './childComps/HomeTabControl'

export default {
  name: "Home",
  components: {
    NavBar,
    GoodList,
    HomeSwiper,
    HomeRecommand,
    HomeFeature,
    HomeTabControl
  },
  data() {
    return { 
      banners: [],
      recommands: [],
      goods: {
        pop: {page: 0, list: []},
        news: {page: 0, list: []},
        sell: {page: 0, list: []},
      },
      
      fakeNum: 100,  /* 伪造数据 */
    }
  },
  methods: {
    // 轮播图相关数据
    getHomeMultiData() {
      getHomeMultiData()
      .then(resp => {
        this.banners = resp.data.banners
        this.recommands = resp.data.recommands
      })
      .catch(err => {
        console.log(err)
      })
    },
    // 商品相关数据
    getHomeGoods(type) {
      let page = this.goods[type].page + 1
      getHomeGoods(type, page)
      .then(resp => {
        this.goods[type].list.push(...resp.data[type].list)
        this.goods[type].page = page
      })
      .catch(err => {
        console.log(err)
      })
    }
  },
  created() {
    // prepare data
    this.getHomeMultiData()
    for (let goodType in this.goods) {
      this.getHomeGoods(goodType, 0)
    }
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