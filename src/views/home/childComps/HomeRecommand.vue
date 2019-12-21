<template>
  <div id="home-recommand">
    <a v-for="(item, index) in recommands" :key="index" 
      :href="item.link"
    ><img :src="item.image" alt=""><div>{{ item.title }}</div></a>
  </div>
</template>

<script>
export default {
  name: 'HomeRecommand',
  props: {
    recommands: {
      type: Array,
      default: [],
    }
  }
}
</script>

<style>

#home-recommand {
  display: flex;
  text-align: center;
  /* margin: 8px 0; */
  padding-bottom: 17px;
  border-bottom: 13px solid #eee8e8;
}

#home-recommand a {
  flex: 1;
  font-size: 13px;
  letter-spacing: 2px;
}

#home-recommand img {
  width: 70px;
  height: 70px;
  margin: 10px 0 5px;
}

</style>