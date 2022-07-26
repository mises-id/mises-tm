<template>
  <div class="card">
    <div class="title-box flex">
      <div>
        <p class="title">{{ title }}</p>
        <p v-if="secondTitle" class="second-title">{{ secondTitle }}</p>
      </div>
      <div><slot name="rightTool"></slot></div>
    </div>
    <div class="card-content">
      <template v-if="loading">
        <Skeleton active />
        <Skeleton active />
      </template>
      <slot name="content" v-if="!loading"></slot>
    </div>
  </div>
</template>

<script lang="ts">
import {Skeleton} from 'ant-design-vue'
import { defineComponent, useSlots } from 'vue'
export default defineComponent({
  name: 'MisesCard',
  components:{
    Skeleton
  },
  props: {
    title: {
      type: String,
      default: ''
    },
    secondTitle: {
      type: String,
      default: ''
    },
    loading:{
      type: Boolean,
      default: false
    }
  },
})
</script>

<style scoped lang="scss">
.card {
  padding: 50px 30px;
  background: white;
  border: 1px solid #eeeeee;
  box-shadow: 0px 10px 30px 0px rgba(0, 0, 0, 0.05);
  border-radius: 10px;
}
.title-box {
  justify-content: space-between;
}
@media (max-width: 768px) {
  .title-box {
    flex-direction: column;
    justify-content: flex-start;
    align-items: flex-start;
  }

  .card {
    padding: 25px 15px;
    max-width: 92vw;
    box-sizing: border-box;
    margin:0 auto;
  }
}
.title {
  margin: 0;
  font-size: 22px;
  font-family: 'hlt-75';
  color: #16161d;
  line-height: 1;
}
.second-title {
  font-size: 14px;
  font-family: 'hlt-55';
  font-weight: normal;
  color: #666666;
  margin-top: 12px;
  margin-bottom: 0;
  line-height: 1;
}
.card-content{
  margin-top: 26px;
}
</style>