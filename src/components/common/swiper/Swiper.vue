<template>
  <!-- 可视区域 -->
  <div id="swiper" :style="customStyle">
    
    <!-- 轮播区域 -->
    <div class="swiper-imgs" 
      :style="autoMove">
      <slot></slot>
    </div>
    
    <!-- 圆点 -->
    <div class="swiper-indexes" 
      :style="autoMoveCenter" 
    >
      <ul v-for="(item, index) in Array(dotNum)" :key="index">
        <li :style="isActive(index)"></li>
      </ul>
    </div>

  </div>
</template>

<script>
export default {
  name: "Swiper",
  props: {
    swiperWidth: {
      type: String,
      default: "100%"
    },
    swiperHeight: {
      type: String,
      default: "80px"
    },
    dotActiveColor: {
      type: String,
      default: "orange"
    },
    dotDeactiveColor: {
      type: String,
      default: "gray"
    }
  },
  data() {
    return {
      curIndex: 0, /* 当前图片索引 */
      dotNum: 0, /* 圆点个数 */
    }
  },
  computed: {
    // 用户自定义样式：可视区域的大小
    customStyle() {
      return {
        width: this.swiperWidth,
        height: this.swiperHeight
      }
    },
    isActive() {
      return (index) => this.curIndex === index ? {
        backgroundColor: this.dotActiveColor
      } : {
        backgroundColor: this.dotDeactiveColor
      }
    },
    // 轮播
    autoMove() {
      return {
        left: String(this.curIndex * -360) + 'px'
      }
    },
    // 计算索引点自动居中
    autoMoveCenter() {
      let wholeWidth = this.swiperWidth.match(/\d+/)
      if (wholeWidth === null) {
        throw "does not know width"
      } else {
        wholeWidth = Number(wholeWidth)
      }
      // 计算偏移距离
      let leftDistance = (wholeWidth - 20 * this.dotNum) / 2
      // 计算便宜百分比
      return {
        left: String(leftDistance / wholeWidth * 100) + '%'
      }
    }
  },
  methods: {
  },
  mounted() {
    setInterval(() => {
      ++this.curIndex
      if (this.curIndex > this.dotNum - 1) {
        this.curIndex = 0
      }
    }, 3000)
  },
  created() {
    this.$on("add-dot-num", () => {
      let slotnum = Object.keys(this.$slots)
      // 限制圆点数超过 swiper-item 个数
      if (slotnum !== 0 && this.$slots.default.length === this.dotNum) {
        return
      }
      ++this.dotNum
    })
  },
}
</script>

<style>
#swiper {
  position: relative;
  width: 100%;
  height: 100px;
  overflow: hidden;
}

.swiper-imgs {
  position: absolute;
  display: flex;
  justify-content: center;
}

.swiper-indexes {
  position: absolute;
  display: flex;
  bottom: 10px;
}

.swiper-indexes li {
  list-style: none;
  width: 10px;
  height: 10px;
  background-color: gray;
  border-radius: 50%;
  margin: 0 5px;
}

</style>