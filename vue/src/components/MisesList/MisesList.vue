<template>
  <div>
    <p v-if="title" class="list-title">{{title}}</p>
    <div class="data-item" v-for="(item, index) in data" :key="index">
      <div class="title">{{item.title}}</div>
      <MisesValue  v-if="item.render" :value="item.render?.(item)"></MisesValue>
      <div class="value" v-else>{{item.value}}</div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue'
import MisesValue from './render.vue'
export interface dataItem {
  title: string
  value: string
  render?: (h: any) => any
}
export default defineComponent({
  components:{
    MisesValue,
  },
  props: {
    data: {
      type: Array as PropType<dataItem[]>,
      default: ''
    },
    title:{
      type:String,
      default:''
    }
  },
})
</script>

<style scoped lang="scss">
.data-item {
  display: flex;
  background-color: #fcfcfc;
  color: #16161D;
  height: 50px;
  line-height: 50px;
  font-size: 16px;
  &:nth-child(2n) {
    background-color: #f8f8f8;
  }
  .title{
    width: 40%;
    font-family: 'hlt-75';
    box-sizing: border-box;
    padding-left: 30px;
  }
  .value{
    width: 60%;
    font-family: 'hlt-45';
  }
}
.list-title{
	font-size: 18px;
	color: #666666;
	font-family: 'hlt-75';
  margin-top: 25px;
}
</style>