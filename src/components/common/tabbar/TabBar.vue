<template>
  <div id="tab-bar">
    <slot></slot>
  </div>
</template>

<script>
export default {
  name: "TabBar"
}
</script>

<style>

#tab-bar {
  width: 100%;
  height: 65px;
  background: #e9e9e9;

  position: fixed;
  left: 0px;
  bottom: 0px;
}

</style>