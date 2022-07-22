<template>
  <div class="content content-top">
    <Card title="Blocks">
      <template #content>
        <a-table
          :dataSource="blocks"
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
import { getBlocks } from '../api/serve'
import Card from '../components/MisesCard'
import { Table } from 'ant-design-vue'
import { formatTime } from '../utils/plugins'
export default {
  components: {
    Card,
    ATable: Table
  },
  name: 'Blocks',
  data() {
    return {
      blocks: [],
      pages: 1,
      page_size: 10,
      total: 0,
      columns: [
        {
          title: 'Block Height',
          dataIndex: 'height',
          width: '15%'
        },
        {
          title: 'Hash',
          dataIndex: 'hash',
          ellipsis: true
        },
        {
          title: 'Age',
          dataIndex: 'timeAgo',
          width: '20%'
        },
        {
          title: 'Txns',
          dataIndex: 'transactions',
          customRender: ({ text }) => {
            return `${text} Txns in this block`
          },
          width: '15%'
        },
        {
          title: 'Block Reward',
          dataIndex: 'block_reward',
          customRender: ({ text }) => {
            return `${text} MIS`
          },
          width: '15%'
        }
      ],
      loading: false
    }
  },
  created() {
    this.getBlockList()
  },
  methods: {
    getBlockList() {
      this.loading = true
      getBlocks({
        page_num: this.pages,
        page_size: this.page_size
      })
        .then((res) => {
          this.blocks = res.data.map((val) => {
            return {
              details: val.block,
              timestamp: val.block.header.time,
              height: val.height,
              block_reward: val.block_reward || 0,
              transactions: val.transactions,
              validdator: val.validdator?.moniker ?? '',
              hash: val.block_id.hash,
              timeAgo: formatTime(val.block.header.time)
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
      this.getBlockList()
    },
    handleClickRow(record) {
      return {
				onClick: () => {
					this.$router.push(`/block/${record.height}`)
				}
      }
    }
  }
}
</script>
