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
  display: flex;

  width: 100%;
  height: 65px;
  background: #efefef;

  position: fixed;
  left: 0px;
  bottom: 0px;

  box-shadow: 1px 0px 1px rgba(100, 100, 100, 0.1);
}

</style>