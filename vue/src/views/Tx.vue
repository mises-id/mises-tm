<template>
  <div class="content content-top">
    <Card title="Transactions Detail" :loading="loading">
      <template #content>
        <MisesList :data="block" />
      </template>
    </Card>
    <MisesDialog v-model:visible="dialogVisable" title="More Info" :width="1500">
      <MisesList :data="fields" v-if="fields.length" />
    </MisesDialog>
    <!-- <SpBlockDisplayFull :block="block" /> -->
  </div>
</template>

<script lang="ts">
import Card from '../components/MisesCard'
import MisesList from '../components/MisesList'
import { getBlock, getTx, getTxs } from '../api/serve'
import { dataItem } from '../components/MisesList/MisesList.vue'
import MisesDialog from '../components/MisesDialog'
import { formatTime, shortenAddress } from '../utils/plugins'
import dayjs from 'dayjs'
import { h } from 'vue'
import { message } from 'ant-design-vue'
export default {
  components: {
    Card,
    MisesList,
    MisesDialog
  },
  data() {
    return {
      block: [] as dataItem[],
      loading: false,
      dialogVisable: false,
      fields: []
    }
  },
  async created() {
    this.loading = true
    try {
      const res = await getTx({ hash: this.$route.params.txhash })
      this.block = [
        {
          title: 'Transaction Hash',
          value: res.hash
        },
        {
          title: 'Status',
          value: res.transactions,
          render: ({ value }) => {
            const isSuccess = res.tx_result?.code === 0
            return h(
              'div',
              {
                style: {
                  display: 'flex',
                  alignItems: 'center'
                }
              },
              [
                h('img', {
                  class: 'iconfont',
                  src: isSuccess ? '/images/index/success@2x.png' : 'icon-error'
                }),
                h(
                  'span',
                  {
                    style: {
                      marginLeft: '5px',
                      color: isSuccess ? '#50C68D' : '#FF6060'
                    }
                  },
                  isSuccess ? 'Success' : 'Fail'
                )
              ]
            )
          }
        },
        {
          title: 'Block',
          value: res.height
        },
        {
          title: 'Timestamp',
          value: `${formatTime(res.block_time)}(${dayjs(res.block_time).format('MMM-DD-YYYY HH:mm:ss A [ +UTC]')})`
        },
        {
          title: 'From',
          value: res.tx_msg?.from_address ?? '-'
        },
        {
          title: 'To',
          value: res.tx_msg?.to_address ?? '-'
        },
        {
          title: 'Value',
          value: `${res.value_mis} MIS`
        },
        {
          title: 'Gas Fee',
          value: `${res.gas_fee} MIS`
        },
        {
          title: 'More Info',
          value: res.tx_result.tx,
          render: ({ value }) => {
            return h(
              'span',
              {
                class: 'active',
                onClick: () => {
                  this.fields = Object.keys(value).map((val) => {
                    return {
                      title: val,
                      value: value[val]
                    }
                  })
                  this.dialogVisable = true
                }
              },
              'View All'
            )
          }
        }
      ] as dataItem[]
    } finally {
      this.loading = false
    }
  }
}
</script>
<style lang="scss">
.iconfont {
  width: 14px;
  height: 14px;
}
</style>