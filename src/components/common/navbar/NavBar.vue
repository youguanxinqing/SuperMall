<template>
  <div id="nav-bar">
    <div class="left" :style="styleBackgroundColor"><slot name="left"></slot></div>
    <div class="center" :style="centerStyle"><slot name="center"></slot></div>
    <div class="right" :style="styleBackgroundColor"><slot name="right"></slot></div>
  </div>
</template>

<script>
export default {
  name: "NavBar",
  props: {
    textColor: {
      type: String,
      default: '#fff'
    },
    backgroundColor: {
      type: String,
      default: '#d4237a'
    }
  },
  computed: {
    centerStyle() {
      return {
        backgroundColor: this.backgroundColor,
        color: this.textColor
      }
    },
    styleBackgroundColor() {
      return {
        backgroundColor: this.backgroundColor
      }
    }
  }
}

</script>

<style>
#nav-bar {
  display: flex;
  height: 60px;
}

.left, .right {
  width: 60px;
  /* background-color: red; */
}

.center {
  flex: 1;
  /* background-color: green; */
  text-align: center;
  line-height: 60px;
  letter-spacing: 1.5px;
}

</style>