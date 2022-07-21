<template>
	<div class="content content-top">
    <Card title="Block Detail" :loading="loading">
			<template #content>
				<MisesList :data="block"/>
			</template>
		</Card>
		<!-- <SpBlockDisplayFull :block="block" /> -->
	</div>
</template>

<script lang="ts">
import Card from '../components/MisesCard'
import MisesList from '../components/MisesList'
import {getBlock} from '../api/serve'
import { dataItem } from '../components/MisesList/MisesList.vue'
import { formatTime, shortenAddress } from '../utils/plugins'
import dayjs from 'dayjs'
import { h } from 'vue'
import { message } from 'ant-design-vue'
export default {
	components: {
		Card,MisesList
	},
	data() {
		return {
			block: [] as dataItem[],
			loading:false
		}
	},
	async created() {
		this.loading = true;
		// const blockDetails = await axios.get(this.$store.getters['common/env/apiTendermint'] + '/block?height=' + this.$route.params.block)
		// const txDecoded = blockDetails.data.result.block.data.txs.map(async (x) => {
		// 	const txhash = toHex(sha256(fromBase64(x)))
		// 	const dec = await decodeTx(this.$store.getters['common/env/apiCosmos'], this.$store.getters['common/env/apiTendermint'], txhash)
		// 	return dec
		// })
		// const txs = await Promise.all(txDecoded)
		// this.block = {
		// 	height: blockDetails.data.result.block.header.height,
		// 	timestamp: blockDetails.data.result.block.header.time,
		// 	hash: blockDetails.data.result.block_id.hash,
		// 	details: blockDetails.data.result.block,
		// 	txDecoded: [...txs]
		// }
		try {
			const res = await getBlock({height:this.$route.params.block})
			this.block = [{
				title:'Block',
				value: res.height
			},{
				title:'Timestamp',
				value: `${formatTime(res.block.header.time)}(${dayjs(res.block.header.time).format('MMM-DD-YYYY HH:mm:ss A [ +UTC]')})`
			},{
				title:'Transactions',
				value: res.transactions,
				render:({value})=>{
					return h('span',{
						class:'active',
						onClick: ()=>{
							if(value>0){
								this.$router.push(`/blockTx/${this.$route.params.block}`)
							}else{
								message.info('No transactions in this block')
							}
						}
					},`${value} Txns in this block`)
				}
			},{
				title:'Validated',
				value: res.validdator ? `${res.validdator.moniker}(${shortenAddress(res.validdator.address)})`: '-'
			},{
				title:'Gas Used',
				value: res.gas_used
			},{
				title:'Gas Limit',
				value: res.gas_limit
			}] as dataItem[]
		} finally{
			this.loading = false;
		}
	}
}
</script>
<style lang="scss">
.active{
	color:#5D61FF;
	cursor: pointer;
}
</style>