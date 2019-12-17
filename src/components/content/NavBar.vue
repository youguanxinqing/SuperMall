<template>
  <div id="use-nav-bar">
    <nav-bar>
      <div slot="left">LEFT</div>
      <div slot="center">LEFT</div>
      <div slot="right">LEFT</div>
    </nav-bar>
  </div>
</template>

<script>
import NavBar from 'components/common/navbar/NavBar'

export default {
  components: {
    NavBar
  }
}
</script>

<style>

</style>