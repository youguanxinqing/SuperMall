<template>
  <div id="good-item">
    <img :src="'http://127.0.0.1:8081'+goodItem.image" alt="" srcset="">
    <div class="item-info">
      <p class="item-title">{{ goodItem.title }}</p>
      <span class="item-price" >{{ goodItem.price.toFixed(2) }}</span>
      <span class="item-fav" >{{ goodItem.fav }}</span>
    </div>
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
  width: 45%;
  height: 210px;
  font-size: 13px;
  margin: 5px 0 5px;
  padding-bottom: 40px;
  text-align: center;
}

img {
  display: inline-block;
  margin-bottom: 5px;
  width: 95%;
  height: 100%;
  border-radius: 3%;
}

.item-info p {
  display: inline-block;
  width: 95%;
  height: 13px;
  white-space: nowrap;
  overflow:hidden;
  text-overflow: ellipsis;
  color: #888;
  margin: 0 0 1px;
}

.item-info span{
  margin-left: 5px; 
}

.item-price {
  color: var(--color-high-text)
}

.item-fav {
  padding-left: 15px;
  position: relative;
}
.item-fav::before {
  content: "";
  position: absolute;
  left: 0px;
  width: 13px;
  height: 12px;
  background: url("~assets/img/common/coll-active.svg") 0 0/13px 12px;
}

</style>
