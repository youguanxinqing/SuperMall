<template>
  <div id="app">
    <router-view></router-view>
    <tab-bar>
      <tab-bar-item path="/home" textColor="#d4237a">
        <img slot="item-active" src="@/assets/img/tabbar/index-active.svg" alt="">
        <!-- <img src="" alt=""> -->
        <img slot="item-deactive" src="@/assets/img/tabbar/index.svg">
        <div slot="item-text">首页</div>
      </tab-bar-item>
      <tab-bar-item path="/category" textColor="#d4237a">
        <img slot="item-active" src="@/assets/img/tabbar/category-active.svg" alt="">
        <img slot="item-deactive" src="@/assets/img/tabbar/category.svg" alt="">
        <div slot="item-text">分类</div>
      </tab-bar-item>
      <tab-bar-item path="/cart" textColor="#d4237a">
        <img slot="item-active" src="@/assets/img/tabbar/cart-active.svg" alt="">
        <img slot="item-deactive" src="@/assets/img/tabbar/cart.svg" alt="">
        <div slot="item-text">购物车</div>
      </tab-bar-item>
      <tab-bar-item path="/profile" textColor="#d4237a">
        <img slot="item-active" src="@/assets/img/tabbar/profile-active.svg" alt="">
        <img slot="item-deactive" src="@/assets/img/tabbar/profile.svg" alt="">
        <div slot="item-text">我的</div>
      </tab-bar-item>
    </tab-bar>
  </div>
</template>

<script>

import TabBar from 'components/common/tabbar/TabBar'
import TabBarItem from 'components/common/tabbar/TabBarItem'

export default {
  name: 'app',
  components: {
    TabBar,
    TabBarItem
  }
}
</script>

<style>
  @import "assets/css/base.css";
</style>
