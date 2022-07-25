<template>
  <div>
    <div class="homepage">
      <p class="welcome">The Mises Blockchain Explorer</p>
      <div class="search">
        <input type="text" placeholder="Search by Txn Hash/Block Height/Address" @keyup="getKeyUp" v-model="search" />
        <div class="search-btn" @click="getText">
          <img src="/images/index/search@2x.png" alt="" />
        </div>
      </div>
      <div class="mises-stats">
        <div class="flex stats-item">
          <div class="flex stats-item-icon-box">
            <img :src="`/images/index/icon_1@2x.png`" alt="" class="stats-item-icon" />
          </div>
          <div class="stats-item-value">
            <span class="title">Total MIS Supply</span>
            <p class="value">{{ supply }} MIS</p>
          </div>
        </div>
        <div class="flex stats-item">
          <div class="flex stats-item-icon-box">
            <img :src="`/images/index/icon_2@2x.png`" alt="" class="stats-item-icon" />
          </div>
          <div class="stats-item-value">
            <span class="title">Med Gas Price</span>
            <p class="value">{{ gasPrice }}</p>
          </div>
        </div>
        <div class="flex stats-item">
          <div class="flex stats-item-icon-box">
            <img :src="`/images/index/icon_1@2x.png`" alt="" class="stats-item-icon" />
          </div>
          <div class="stats-item-value">
            <span class="title">Latest Block</span>
            <p class="value">{{ blockHeight }}</p>
          </div>
        </div>
        <div class="flex stats-item">
          <div class="flex stats-item-icon-box">
            <img :src="`/images/index/icon_1@2x.png`" alt="" class="stats-item-icon" />
          </div>
          <div class="stats-item-value">
            <span class="title">Social Relationships</span>
            <p class="value">{{ relationships }}</p>
          </div>
        </div>
      </div>
    </div>
    <div class="block-container">
      <div class="block-box">
        <p class="title">MIS Coin Top Holders</p>
        <div class="block-list">
          <div v-for="(item, index) in holderList" :key="index" class="block-item flex" @click="holder(item)">
            <div>
              <p class="name">{{ item.name }}</p>
              <p class="second-value">{{ item.misesid }}</p>
            </div>
            <div>
              <p>Quantity</p>
              <p>{{ item.quantity }} MIS</p>
            </div>
            <div>
              <p>Percentage</p>
              <p>{{ item.percentage }}%</p>
            </div>
          </div>
          <div v-if="!holderList.length">
            <Skeleton active />
            <Skeleton active />
            <Skeleton active />
          </div>
        </div>
        <router-link class="block-btn" to="/holders">View all Holders</router-link>
      </div>
      <div class="block-box">
        <p class="title">Latest Blocks</p>
        <div class="block-list">
          <div v-for="item in blocks" :key="item.height" class="block-item flex" @click="block(item)">
            <div>
              <p>Height: {{ item.height }}</p>
              <p class="second-value">{{ formatTS(item.timestamp) }}</p>
            </div>
            <div>
              <p>Miner: {{ item.validdator }}</p>
              <p>{{ item.transactions || 0 }} Txns in {{ item.timeAgo }}</p>
            </div>
            <div>
              <p>Block Reward</p>
              <p>{{ item.block_reward || 0 }} MIS</p>
            </div>
          </div>

          <div v-if="!blocks.length">
            <Skeleton active />
            <Skeleton active />
            <Skeleton active />
          </div>
        </div>
        <router-link class="block-btn" to="/blocks">View all Blocks</router-link>
      </div>
    </div>
  </div>
</template>

<script  lang="ts">
import { computed } from 'vue'
import { useStore } from 'vuex'
import { getIndexPageStats, getBlocks, getTopHolder } from '../api/serve'
import { Block } from '../utils/interfaces'
import dayjs from 'dayjs'
import BigNumber from 'bignumber.js'
import { formatTime, shortenAddress } from '../utils/plugins/index'
import { Skeleton } from 'ant-design-vue'
export default {
  components: {
    Skeleton
  },
  name: 'Portfolio',
  data() {
    return {
      supply: 'N/A',
      gasPrice: 'N/A',
      blockHeight: 'N/A',
      relationships: 'N/A',
      initBlocks: [],
      holders: [],
      search: ''
    }
  },
  setup() {
    const $s = useStore()
    const supply = computed(() => {
      const data = $s.getters['cosmos.bank.v1beta1/getTotalSupply']()
      const supply = data.supply?.[0]
      const totalSupply = supply?.amount ? new BigNumber(supply.amount).div(1000000).toFixed(4).toString() : 0
      return totalSupply
    })
    return {
      supply
    }
  },

  computed: {
    blocks(): Array<Block> {
      try {
        let $s = useStore()
        const blocks = $s.getters['common/blocks/getBlocks'](5).map((val) => {
          return {
            ...val,
            timeAgo: formatTime(val.timestamp)
          }
        })
        if (blocks[0]?.height) this.blockHeight = blocks[0].height
        return [...blocks, ...this.initBlocks].slice(0, 5)
      } catch (error) {
        return []
      }
    },
    holderList() {
      if (this.holders.length && this.supply) {
        return this.holders.map((val) => {
          return {
            ...val,
            percentage: new BigNumber(val.quantity).div(this.supply).times(100).integerValue().toString() || '0'
          }
        })
      }
      console.log('xxx')
      return this.holders
    }
  },
  created() {
    this.getStatsData()
    this.getBlockList()
    this.getHolderList()
  },
  methods: {
    formatTS(timestamp: string): string {
      const momentTime = dayjs(timestamp)
      return momentTime.format('YYYY.MM.DD HH:mm:ss')
    },
    // update home page data
    getStatsData() {
      getIndexPageStats().then((res) => {
        this.relationships = res.social_relationships
        this.gasPrice = `${res.gas_price}MIS`
      })
    },
    getBlockList() {
      getBlocks({
        page_num: 1,
        page_size: 5
      }).then((res) => {
        this.initBlocks = res.data.map((val) => {
          return {
            details: val.block,
            timestamp: val.block.header.time,
            height: val.height,
            block_reward: val.block_reward || 0,
            transactions: val.transactions,
            validdator: val.validdator?.moniker ?? '',
            timeAgo: formatTime(val.block.header.time)
          }
        })
        this.blockHeight = this.initBlocks[0].height
      })
    },
    getHolderList() {
      getTopHolder({
        list_num: 5
      }).then((res) => {
        this.holders = res.map((val) => {
          return {
            ...val,
            name: val.pubinfo?.name ?? shortenAddress(val.misesid),
            quantity: new BigNumber(val.user_ext.quantity).integerValue().toString(),
            relationships: 0
          }
        })
      })
    },
    getKeyUp(val: KeyboardEvent) {
      if (val.code.toLocaleLowerCase() === 'enter') {
        this.getText()
      }
    },
    getText() {
      const hashLength = 64
      const addressLength = 44
      const blockHeight = (number: string) => !new BigNumber(number).isNaN()
      if (this.search.length === hashLength) {
        this.$router.push(`/tx/${this.search}`)
        return
      }
      if (this.search.length === addressLength) {
        this.$router.push(`/holders/${this.search}`)
        return
      }
      if (blockHeight(this.search)) {
        this.$router.push(`/block/${this.search}`)
      }
    },
    holder(item) {
      this.$router.push(`/holders/${item.misesid}`)
    },
    block(item) {
      this.$router.push(`/block/${item.height}`)
    }
  }
}
</script>

<style scoped lang="scss">
.row {
  display: flex;
  flex-wrap: wrap;
}
.col {
  flex-grow: 1;
  padding: 20px;
}
.bg {
  width: 100%;
}
.homepage {
  position: relative;
  background-image: url('/images/index/bg@2x.png');
  background-repeat: no-repeat;
  background-position: center;
  background-color: #16161d;
  height: 250px;
  background-size: 100vw 13vw;
  .welcome {
    padding-top: 60px;
    text-align: center;
    color: white;
    margin: 0;
    font-size: 22px;
    font-family: 'hlt-75';
  }
  .search {
    margin: 22px auto 0;
    display: flex;
    justify-content: center;
    input {
      height: 44px;
      background: #ffffff;
      border-radius: 5px;
      width: 596px;
      outline: none;
      border: none;
      padding-left: 16px;
      font-size: 16px;
      font-family: 'hlt-55';
      position: relative;
      right: -6px;
      color: #16161d;
      &::placeholder {
        color: #ccc;
      }
    }
    .search-btn {
      border-radius: 5px;
      background: #5d61ff;
      width: 50px;
      height: 44px;
      display: flex;
      align-items: center;
      justify-content: center;
      cursor: pointer;
      img {
        width: 16px;
        height: 16px;
      }
    }
  }
}
.mises-stats {
  width: 1200px;
  height: 120px;
  background: #ffffff;
  border: 1px solid #eeeeee;
  box-shadow: 0px 10px 30px 0px rgba(0, 0, 0, 0.05);
  border-radius: 10px;
  margin: 0 auto;
  position: relative;
  top: 64px;
  display: flex;
  justify-content: space-between;
}
.stats-item {
  .stats-item-icon-box {
    padding: 10px 0 10px 30px;
  }
  &:not(:first-child) {
    .stats-item-icon-box {
      border-left: 1px solid #eeeeee;
    }
  }
  &:last-child {
    margin-right: 90px;
  }
  &-icon {
    display: block;
    width: 40px;
  }
  &-value {
    margin-left: 14px;
    .title {
      font-size: 14px;
      color: #999;
      font-family: 'hlt-55';
    }
    .value {
      color: #16161d;
      margin: 15px 0 0;
      font-family: 'hlt-55';
      line-height: 1;
    }
  }
}
@media (max-width: 768px) {
  $fullWidth: 92vw;
  body {
    .block-container {
      width: $fullWidth;
      margin: 235px auto 0;
      padding-bottom: 50px;
      display: grid;
      gap: 10px;
      .block-box {
        .block-item {
          gap: 10px;
          .second-value {
            width: 30vw;
          }
          .name {
            width: 30vw;
          }
          div {
            &:last-child {
              margin-right: 0;
            }
          }
        }
      }
    }
    .mises-stats {
      flex-direction: column;
      height: auto;
      width: $fullWidth;
      margin: 0 auto;
      top: 30px;
      .stats-item{
        width:100%;
        &:not(:last-child){
          border-bottom: 1px solid #eee;
        }
        .stats-item-icon-box{
          padding-left: 10px;
        }
        .value{
          margin-top:5px;
        }
      }
    }
    .homepage {
      height: 200px;
      .search {
        margin: 22px auto 0;
        display: flex;
        justify-content: center;
        input {
          width: 80vw;
          height: 32px;
          padding-left: 8px;
          font-size: 12px;
        }
        .search-btn {
          width: 40px;
          height: 32px;
          img {
            width: 12px;
            height: 12px;
          }
        }
      }
    }
  }
}
@media (min-width: 768px) {
  .block-container {
    width: 1200px;
    margin: 110px auto 20px;
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 30px;
  }
}
.block-container {
  .block-box {
    background: #ffffff;

    box-shadow: 0px 10px 30px 0px rgba(0, 0, 0, 0.05);
    border: 1px solid #eeeeee;
    border-radius: 10px;
    .title {
      font-size: 16px;
      color: #16161d;
      font-family: 'hlt-75';
      line-height: 50px;
      padding-left: 16px;
      border-bottom: 1px solid #eeeeee;
      margin: 0;
    }
    .block-list {
      padding: 0 15px;
      border-bottom: 1px solid #eee;
    }
    .block-btn {
      display: block;
      cursor: pointer;
      height: 34px;
      background: #5d61ff;
      border-radius: 5px;
      text-align: center;
      line-height: 37px;
      color: white;
      font-family: 'hlt-55';
      margin: 15px;
    }
    .block-item {
      padding: 19px 0;
      font-family: 'hlt-55';
      font-size: 14px;
      color: #16161d;
      justify-content: space-between;
      &:not(:first-child) {
        border-top: 1px solid #eeeeee;
      }
      cursor: pointer;
      p {
        margin: 0;
        &:last-child {
          margin-top: 9px;
        }
      }
      .second-value {
        color: #999999;
        width: 180px;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
      }
      .name {
        width: 180px;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
      }
      div {
        &:last-child {
          margin-right: 80px;
        }
      }
    }
  }
}
</style>