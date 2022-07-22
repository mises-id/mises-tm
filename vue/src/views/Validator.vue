<template>
  <div class="content content-top">
    <Card title="Validator Detail" :loading="loading">
      <template #content>
        <MisesList :data="block" title="Overview" />
        <div class="">
          <p class="list-title">Transactions</p>
          <a-table
            :dataSource="txns"
            :columns="columns"
            :loading="listloading"
            @change="handleChange"
            :bordered="false"
            rowKey="height"
            :customRow="handleClickRow"
            :pagination="{
              total: total,
              current: pages,
              pageSize: page_size,
              showQuickJumper: true
            }"

          :scroll="{x:1000}"
            :rowClassName="(_record, index) => (index % 2 === 1 ? 'table-striped' : undefined)"
          />
        </div>
      </template>
    </Card>
  </div>
</template>

<script lang="ts">
import Card from '../components/MisesCard'
import MisesList from '../components/MisesList'
import { getAddress, getTxs } from '../api/serve'
import { dataItem } from '../components/MisesList/MisesList.vue'
import { formatTime, getCalcVotingPowerRate, shortenAddress, uptime_estimated, valConsAddress } from '../utils/plugins'
import {Table} from 'ant-design-vue'
import { h } from '@vue/runtime-core'
import axios from 'axios'
import BigNumber from 'bignumber.js'
export default {
  components: {
    Card,
    MisesList,
    ATable: Table
  },
  data() {
    return {
      block: [] as dataItem[],
      txns: [],
      loading: false,
      pages: 1,
      page_size: 10,
      total: 0,
      columns: [
        {
          title: 'Txn Hash',
          dataIndex: 'hash',
          width: '25%',
          ellipsis: true
        },
        {
          title: 'Block',
          dataIndex: 'height',
          width: '8%'
        },
        {
          title: 'Time',
          dataIndex: 'block_time',
          width: '15%'
        },
        {
          title: 'From',
          dataIndex: 'from_address',
          // customRender:({text}, record)=> {
          //   const isMe = text === this.$route.params.misesid
          //   return h('div',{},[
          //     h('span',{},isMe ? this.username : shortenAddress(text)),
          //     h('span',{class:isMe ? 'out': 'in'},isMe ? 'Out' : 'In')
          //   ])
          // }
        },
        {
          title: 'To',
          dataIndex: 'to_address'
        },
        {
          title: 'Value',
          dataIndex: 'value_mis',
          width: '10%'
        },
        {
          title: 'Gas Fee',
          dataIndex: 'gas_fee',
          width: '10%'
        }
      ]
    }
  },
  async created() {
    this.loading = true
    try {
      const page = await axios.get(this.$store.getters['common/env/apiCosmos'] + '/cosmos/staking/v1beta1/validators')
      const calcRate = getCalcVotingPowerRate(page.data.validators)
      const data = await axios.get(this.$store.getters['common/env/apiCosmos'] + `/cosmos/staking/v1beta1/validators/${this.$route.params.misesid}`)
      const res = data.data.validator;
    const {
      data: { info: validators }
    } = await axios.get(this.$store.getters['common/env/apiCosmos'] + '/cosmos/slashing/v1beta1/signing_infos')
      const {rate } = res.commission.commission_rates
      const voting_power_rate = calcRate(res.operator_address)
      const info = validators?.find(({ address }) => address === valConsAddress(res)) || {}
      const uptime = uptime_estimated(info);
      this.block = [
        {
          title: 'Moliker',
          value: res.description.moniker,
        },
        {
          title: 'Address',
          value: res.operator_address
        },
        {
          title: 'Bonded MIS',
          value: `${new BigNumber(res.tokens).div(1000000).toFixed(6).toString()}MIS`
        },
        {
          title: 'Commission',
          value: `${new BigNumber(rate).times(100)}%`
        },
        {
          title: 'Voting Power',
          value: `${voting_power_rate&&new BigNumber(voting_power_rate).times(100).toFixed(2)}%`
        },
        {
          title: 'Uptime',
          value: `${uptime&&new BigNumber(uptime).times(100).toFixed(2)}%`
        },
        {
          title: 'Active',
          value: res.jailed,
          render({value}){
            return h('span',{
              style:{
                color:!value?'#50C68D':'#FF6060'
              }
            },!value ? 'Yes' : 'No')
          }
        }
      ] as dataItem[]
    } finally {
      this.loading = false
    }
    this.getTxsList()
  },
  methods: {
    getTxsList() {
      this.listloading = true
      getTxs({
        page_num: this.pages,
        page_size: this.page_size,
        validator_address: this.$route.params.misesid
      })
        .then((res) => {
          this.txns = res.data.map((val) => {
            return {
              block_time: formatTime(val.block_time),
              from_address: shortenAddress(val.tx_msg?.from_address) ?? '-',
              to_address: shortenAddress(val.tx_msg?.to_address) ?? '-',
              value_mis:`${val.value_mis} MIS`,
              gas_fee:`${val.gas_fee} MIS`,
              hash: val.hash,
              height: val.height
            }
          })
          this.total = res.pagination.total_records
          this.listloading = false
        })
        .catch(() => {
          this.listloading = false
        })
    },
    handleClickRow(record) {
      return {
        onClick: () => {
          this.$router.push(`/tx/${record.hash}`)
        }
      }
    },
    handleChange(val) {
      this.page_size = val.pageSize
      this.pages = val.current
      this.getTxsList()
    }
  }
}
</script>
<style lang="scss">
.list-title {
  font-size: 18px;
  color: #666666;
  font-family: 'hlt-75';
  margin-top: 40px;
  margin-bottom: 25px;
}
.out{
  color: #FF9600;
  background: #FFEACC;
  display: inline-block;
  margin-left: 5px;
  border-radius: 3px;
  width: 30px;
  height: 18px;
  text-align: center;
  font-size: 13px;
  font-family: 'hlt-45';
  line-height: 22px;
}
.in{
  color: #50C68D;
  background: #DCF4E8;
  display: inline-block;
  margin-left: 5px;
  width: 30px;
  height: 18px;
  text-align: center;
  font-size: 13px;
  font-family: 'hlt-45';
  line-height: 22px;
}
</style>