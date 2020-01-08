<template>
  <div id="home-tab-control">
    <tab-control :fields="['流行', '新款', '精选']"></tab-control>
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

#home-tab-control {
  margin-bottom: 5px;
}

</style>
