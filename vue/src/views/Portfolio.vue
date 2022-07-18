<template>
  <div>
    <div class="homepage">
      <p class="welcome">The Mises Blockchain Explorer</p>
      <div class="search">
        <input type="text" placeholder="Search by Txn Hash/Block/Address/User Name" />
        <div class="search-btn">
          <img src="/images/index/search@2x.png" alt="" />
        </div>
      </div>
      <div class="mises-stats">
        <div v-for="(item, index) in state.misesStats" :key="index" class="flex stats-item">
          <div class="flex stats-item-icon-box">
            <img :src="`/images/index/${item.icon}@2x.png`" alt="" class="stats-item-icon" />
          </div>
          <div class="stats-item-value">
            <span class="title">{{ item.title }}</span>
            <p class="value">{{ item.value }}</p>
          </div>
        </div>
      </div>
    </div>
    <div class="block-container">
      <div class="block-box">
        <p class="title">MIS Coin Top Holders</p>
        <div class="block-list">
          <div v-for="(item, index) in state.holders" :key="index" class="block-item flex">
            <div>
              <p>mises-w1</p>
              <p class="second-value">misemisesvaloper1zt9wqeewewqeeee</p>
            </div>
            <div>
              <p>Quantity</p>
              <p>1109250 MIS</p>
            </div>
            <div>
              <p>Percentage</p>
              <p>25.14%</p>
            </div>
          </div>
        </div>
        <div class="block-btn">View all Holders</div>
      </div>
      <div class="block-box">
        <p class="title">MIS Coin Top Holders</p>
        <div class="block-list">
          <div v-for="(item, index) in state.blocks" :key="index" class="block-item flex">
            <div>
              <p>Height: 1544940</p>
              <p class="second-value">2022.7.1 15:23:36</p>
            </div>
            <div>
              <p>Miner: Mises-West</p>
              <p>12 Txns in 12 sec</p>
            </div>
            <div>
              <p>Block Reward</p>
              <p>0.0032 MIS</p>
            </div>
          </div>
        </div>
        <div class="block-btn">View all Blocks</div>
      </div>
    </div>
  </div>
</template>

<script>
import MisesWelcome from '../components/MisesWelcome'
import { SpAssets, SpTokenTransferList } from '@starport/vue'
import { computed, reactive } from 'vue'
import { useStore } from 'vuex'
export default {
  name: 'Portfolio',
  components: { SpAssets, SpTokenTransferList, MisesWelcome },
  setup() {
    const state = reactive({
      misesStats: [
        {
          icon: 'icon_1',
          title: 'Total MIS Supply',
          value: '105M'
        },
        {
          icon: 'icon_2',
          title: 'Med Gas Price',
          value: '0.005 MIS'
        },
        {
          icon: 'icon_3',
          title: 'Latest Block',
          value: '105M'
        },
        {
          icon: 'icon_4',
          title: 'Social Relationships',
          value: '70,098'
        }
      ],
      holders: [{}, {}, {}, {}, {}],
      blocks: [{}, {}, {}, {}, {}]
    })
    // store
    // let $s = useStore()
    // // computed
    // let address = computed(() => $s.getters['common/wallet/address'])
    return {
      state
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
      font-weight: 'hlt-55';
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
    }
  }
}
.block-container {
  width: 1200px;
  margin: 110px auto 20px;
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 30px;
  .block-box {
    background: #ffffff;
    border: 1px solid #eeeeee;
    box-shadow: 0px 10px 30px 0px rgba(0, 0, 0, 0.05);
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
      cursor: pointer;
      height: 34px;
      background: #5d61ff;
      border-radius: 5px;
      text-align: center;
      line-height: 37px;
      color: white;
      font-family: 'hlt-55';
      margin:15px;
    }
    .block-item {
      border-top: 1px solid #eee;
      padding: 19px 0;
      font-family: 'hlt-55';
      font-size: 14px;
      color: #16161d;
      justify-content: space-between;
      p {
        margin: 0;
        &:last-child {
          margin-top: 9px;
        }
      }
      .second-value {
        color: #999999;
        width: 120px;
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