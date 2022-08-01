<template>
  <div class="content content-top">
    <Card title="Block Detail" :loading="loading">
      <template #content>
        <MisesList :data="block" />
        <a-empty v-if="error" :style="{
					margin:'120px auto'
				}">
          <template #description>
            <span class="error-tips">Error getting block data</span>
          </template>
          <!-- <a-button type="primary" @click="init">REFRESH</a-button> -->
        </a-empty>
      </template>
    </Card>
    <!-- <SpBlockDisplayFull :block="block" /> -->
  </div>
</template>

<script lang="ts">
import Card from '../components/MisesCard'
import MisesList from '../components/MisesList'
import { getBlock } from '../api/serve'
import { dataItem } from '../components/MisesList/MisesList.vue'
import { formatTime, shortenAddress } from '../utils/plugins'
import dayjs from 'dayjs'
import { h } from 'vue'
import { message, Empty } from 'ant-design-vue'
export default {
  components: {
    Card,
    MisesList,
    AEmpty: Empty,
  },
  data() {
    return {
      block: [] as dataItem[],
      loading: false,
      error: false
    }
  },
  created() {
		this.init()
	},
  methods: {
    async init() {
      this.loading = true
      try {
        const res = await getBlock({ height: this.$route.params.block })
        this.block = [
          {
            title: 'Block',
            value: res.height
          },
          {
            title: 'Transaction Hash',
            value: res.block_id.hash
          },
          {
            title: 'Timestamp',
            value: `${formatTime(res.block.header.time)}(${dayjs(res.block.header.time).format('MMM-DD-YYYY HH:mm:ss A [ +UTC]')})`
          },
          {
            title: 'Transactions',
            value: res.transactions,
            render: ({ value }) => {
              return h(
                'span',
                {
                  class: 'active',
                  onClick: () => {
                    if (value > 0) {
                      this.$router.push(`/blockTx/${this.$route.params.block}`)
                    } else {
                      message.info('No transactions in this block')
                    }
                  }
                },
                `${value} Txns in this block`
              )
            }
          },
          {
            title: 'Validated',
            value: res.validdator ? `${res.validdator.moniker}(${shortenAddress(res.validdator.address)})` : '-'
          },
          {
            title: 'Gas Used',
            value: res.gas_used
          },
          {
            title: 'Gas Limit',
            value: res.gas_limit
          }
        ] as dataItem[]
      } catch (error) {
        this.error = true
      } finally {
        this.loading = false
      }
    }
  }
}
</script>
<style lang="scss">
.error-tips{
	font-size: 16px;
	font-family: 'hlt-55';
}
</style>