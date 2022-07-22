<template>
  <div class="content content-top">
    <Card title="Transactions" :secondTitle="secondTitle">
      <template #content>
        <a-table
          :dataSource="txs"
          :columns="columns"
          :loading="loading"
          @change="handleChange"
          :bordered="false"
					rowKey="height"
          :customRow="handleClickRow"
          :pagination="{
            total: total,
            current: pages,
            pageSize: page_size,
						showQuickJumper:true
          }"
          :scroll="{x:1000}"
          :rowClassName="(_record, index) => (index % 2 === 1 ? 'table-striped' : undefined)"
        ></a-table>
      </template>
    </Card>
  </div>
</template>

<script lang="ts">
import { getTxs } from '../api/serve'
import Card from '../components/MisesCard'
import { Table } from 'ant-design-vue'
import { formatTime, shortenAddress } from '../utils/plugins'
import { h, HtmlHTMLAttributes } from 'vue'
export default {
  components: {
    Card,
    ATable: Table
  },
  name: 'Txs',
  data() {
    return {
      txs: [],
      pages: 1,
      page_size: 10,
      total: 0,
      secondTitle:'',
      columns: [
        {
          title: 'Txn Hash',
          dataIndex: 'hash',
          width: '25%',
          ellipsis: true,
        },
        {
          title: 'Block',
          dataIndex: 'height',
          width:'8%',
          customRender: ({text}) => {
            return h('span',{
              class: 'active',
              onClick: (e:Event) => {
                e.stopPropagation()
                this.$router.push(`/block/${text}`)
              }
            },text)
          }
        },
        {
          title: 'Time',
          dataIndex: 'block_time',
          width: '15%',
        },
        {
          title: 'From',
          dataIndex: 'from_address',
        },
        {
          title: 'To',
          dataIndex: 'to_address',
        },
        {
          title: 'Value',
          dataIndex: 'value_mis',
          width:'10%'
        },
        {
          title: 'Gas Fee',
          dataIndex: 'gas_fee',
          width:'10%'
        },
      ],
      loading: false
    }
  },
  created() {
    if(this.$route.params.block) {
      this.secondTitle = `For Block ${this.$route.params.block}`
    }
    this.getTxsList()
  },
  methods: {
    getTxsList() {
      this.loading = true
      getTxs({
        page_num: this.pages,
        page_size: this.page_size,
        block_height: this.$route.params.block
      })
        .then((res) => {
          this.txs = res.data.map((val) => {
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
          this.loading = false
        })
        .catch(() => {
          this.loading = false
        })
    },
    handleChange(val) {
      this.page_size = val.pageSize
      this.pages = val.current
      this.getTxsList()
    },
    handleClickRow(record) {
      return {
				onClick: () => {
					this.$router.push(`/tx/${record.hash}`)
				}
      }
    }
  }
}
</script>
