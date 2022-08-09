<template>
  <div class="content content-top">
    <Card title="Mises Coin Holders">
      <template #content>
        <a-table
          :dataSource="holderList"
          :columns="columns"
          :loading="loading"
          @change="handleChange"
          :bordered="false"
					rowKey="misesid"
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
      <template #rightTool>
        <div class="search">
          <input type="text" placeholder="Search by Mises ID/User Name/Validator Address" @keyup="getKeyUp" v-model="keywords"/>
          <div class="search-btn" @click="getText">
            <img src="/images/index/search@2x.png" alt="" />
          </div>
        </div>
      </template>
    </Card>
  </div>
</template>

<script lang="ts">
import BigNumber from 'bignumber.js'
import dayjs from 'dayjs'
import { getHolders } from '../api/serve'
import Card from '../components/MisesCard'
import { Table } from 'ant-design-vue'
import { formatTime, shortenAddress, useSupply } from '../utils/plugins'
export default {
  components: {
    Card,
    ATable: Table
  },
  name: 'Holders',
  data() {
    return {
      holders: [],
      pages: 1,
      page_size: 10,
      total: 0,
      keywords:'',
      columns: [
        {
          title: 'Address',
          dataIndex: 'misesid',
        },
        {
          title: 'Name Tag',
          dataIndex: 'name',
        },
        {
          title: 'Quantity',
          dataIndex: 'quantity',
          customRender: ({ text }) => {
            return `${text} MIS`
          },
        },
        {
          title: 'Percentage',
          dataIndex: 'percentage',
          customRender: ({ text }) => {
            return `${text} %`
          },
        }
      ],
      loading: false
    }
  },
  setup() {
    const supply = useSupply()
    return {
      supply
    }
  },
  created() {
    this.getHolderList()
  },
  computed:{
    holderList(){
      if(this.holders.length&&this.supply){
        return this.holders.map(val=>{
          return {
            ...val,
            percentage: new BigNumber(val.quantity).div(this.supply).times(100).toFixed(4).toString()
          }
        })
      }
      return this.holders
    }
  },
  methods: {
    getHolderList() {
      this.loading = true
      getHolders({
        page_num: this.pages,
        page_size: this.page_size,
        keywords: this.keywords,
      }).then((res) => {
          this.holders = res.data.filter(val=>val).map((val) => {
            return {
              misesid: val.misesid,
              name: val.name_tag,  
              quantity: new BigNumber(val.user_ext.quantity).integerValue().toString(),
              percentage: 0,
            }
          })
          this.total = res.pagination.total_records
          this.loading = false
        })
        .catch(() => {
          this.loading = false
        })
    },
    handleClickRow(record) {
      return {
				onClick: () => {
					this.$router.push(`/holders/${record.misesid}`)
				}
      }
    },
    handleChange(val) {
      this.page_size = val.pageSize
      this.pages = val.current
      this.getHolderList()
    },
    getKeyUp(val:KeyboardEvent){
      if(val.code.toLocaleLowerCase() === 'enter'){
        this.getText()
      }
    },
    getText(){
      this.getHolderList()
    }
  }
}
</script>
<style lang="scss" scoped>
@media (max-width: 768px) {
  .content .search {
    margin: 12px auto 0;
    display: flex;
    justify-content: center;
    input {
      width: 73vw;
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
      border: 1px solid #EEEEEE;
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
      cursor: pointer;
      img {
        width: 16px;
        height: 16px;
      }
    }
  }
</style>