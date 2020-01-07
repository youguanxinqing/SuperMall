<template>
  <div id="good-item">
    <img :src="'http://127.0.0.1:8081'+goodItem.image" alt="" srcset="">
    <span class="item-title">{{ goodItem.title }}</span>
    <span class="item-fav" >{{ goodItem.fav }}</span>
    <span class="item-org-price" >{{ goodItem.orgPrice }}</span>
    <span class="item-price" >{{ goodItem.price }}</span>
  </div>
</template>

<script>
export default {
  name: 'GoodItem',
  props: {
    goodItem: {
      type: Object,
      default: {
        title: "",
        fav: 0,
        orgPrice: 0,
        price: 0,
      }
    }
  },
  data() {
    return {

    }
  },
}
</script>

<style>
#good-item {
  flex: 1;
  width: 47%;
  font-size: 12px;
  margin: 5px 0 15px;
  justify-content: space-around;
}

img {
  margin-bottom: 5px;
  width: 95%;
  height: 95%;
}

.item-title {
  display: block;
  width: 98%;
  height: 12px;
  white-space: nowrap;
  overflow:hidden;
  text-overflow: ellipsis;
  color: #888;
  margin: 0px;
}
</style>
