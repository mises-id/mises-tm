<template>
  <div class="content content-top">
    <Card title="Validators Leaderboard">
      <template #content>
        <a-table
          :dataSource="validators"
          :columns="columns"
          :loading="loading"
          @change="handleChange"
          :bordered="false"
          rowKey="address"
          :customRow="handleClickRow"
          :pagination="false"

          :scroll="{x:1000}"
          :rowClassName="(_record, index) => (index % 2 === 1 ? 'table-striped' : undefined)"
        ></a-table>
      </template>
    </Card>
  </div>
</template>

<script>
import axios from 'axios'
import Card from '../components/MisesCard'
import { h } from 'vue'
import { getBondedMISCount, getCalcVotingPowerRate, uptime_estimated, valConsAddress } from '../utils/plugins'
import BigNumber from 'bignumber.js'
import {Table} from 'ant-design-vue'
export default {
  components: {
    Card,
    ATable: Table
  },
  name: 'Validators',
  data() {
    return {
      validators: [],
      pages: 1,
      loading: false,
      columns: [
        {
          title: 'Rank',
          dataIndex: 'index',
          customRender: ({ index }) => {
            return `#${index + 1}`
          }
        },
        {
          title: 'Moniker',
          dataIndex: 'moniker'
        },
        {
          title: 'Voting power',
          dataIndex: 'voting_power',
					customRender: ({ text }) => {
						return `${text}%`
					}
        },
        {
          title: 'First Block',
          dataIndex: 'first_block'
        },
        {
          title: 'Last Block',
          dataIndex: 'last_block'
        },
        {
          title: 'Bonded MIS',
          dataIndex: 'bondedMIS'
        },
        {
          title: 'Uptime',
          dataIndex: 'uptime'
        },
        {
          title: 'Active',
          dataIndex: 'status',
          customRender: ({ text }) => {
            return h(
              'span',
              {
                style: {
                  color: text ? '#50C68D' : '#FF6060'
                }
              },
              text ? 'Yes' : 'No'
            )
          }
        }
      ]
    }
  },
  computed: {
    page() {
      return parseInt(this.$route.params.page) || 1
    }
  },
  async created() {
    this.loading = true
    const {
      data: { info: validators }
    } = await axios.get(this.$store.getters['common/env/apiCosmos'] + '/cosmos/slashing/v1beta1/signing_infos')
    const page = await axios.get(this.$store.getters['common/env/apiCosmos'] + '/cosmos/staking/v1beta1/validators')
		const calcRate = getCalcVotingPowerRate(page.data.validators)
    for (let validator_meta of page.data.validators) {
			const voting_power_rate = calcRate(validator_meta.operator_address)
      const info = validators?.find(({ address }) => address === valConsAddress(validator_meta)) || {}
      const uptime = uptime_estimated(info);

      const bondedMIS = getBondedMISCount(page.data.validators, validator_meta.operator_address)
			// console.log(info)
      const validator = {
        moniker: validator_meta.description.moniker,
        address: validator_meta.operator_address,
        status: !validator_meta.jailed,
        voting_power: new BigNumber(voting_power_rate).times(100).toFixed(2),
        proposer_priority: validator_meta.proposer_priority,
        first_block: info?.start_height,
				last_block: new BigNumber(info?.start_height).plus(info?.index_offset).plus(info?.missed_blocks_counter).toString(),
        uptime:`${uptime&&new BigNumber(uptime).times(100).toFixed(2)}%`,
        bondedMIS:`${bondedMIS}MIS`,
      }
      this.validators.push(validator)
    }
		this.validators.sort((a,b)=>a.voting_power - b.voting_power>0?-1:1)
    this.loading = false
  },
  methods: {
    handleClickRow(record) {
      return {
        onClick: () => {
          this.$router.push(`/validators/${record.address}`)
        }
      }
    },
  }
}
</script>
