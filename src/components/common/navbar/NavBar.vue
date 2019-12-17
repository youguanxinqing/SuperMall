<template>
  <div id="nav-bar">
    <div class="right"><slot name="right"></slot></div>
    <div class="center"><slot name="center"></slot></div>
    <div class="left"><slot name="left"></slot></div>
  </div>
</template>

<script>
export default {
  name: "NavBar"
}

</script>

<style>
#nav-bar {
  display: flex;
  height: 60px;
}

.left, .right {
  width: 60px;
  background-color: red;
}

.center {
  flex: 1;
  background-color: green;
}

</style>