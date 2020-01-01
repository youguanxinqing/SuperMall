<template>
  <div id="home-tab-control">
    <tab-control :fields="['慧慧', '庭英', '我是']"></tab-control>
  </div>
</template>

<script>
import TabControl from 'components/content/tabcontrol/TabControl'

export default {
  name: 'HomeTabControl',
  data() {
    return {

    }
  },
  components: {
    TabControl
  }
}
</script>

<style>
</style>
