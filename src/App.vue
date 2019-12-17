<template>
  <div id="app">
    <nav-bar></nav-bar>
    <router-view></router-view>
    <tab-bar></tab-bar>
  </div>
</template>

<script>
import TabBar from 'components/content/TabBar'
import NavBar from 'components/content/NavBar'


export default {
  name: 'app',
  components: {
    TabBar,
    NavBar
  }
}
</script>

<style>
  @import "assets/css/base.css";
</style>
