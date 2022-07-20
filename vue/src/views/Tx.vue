<template>
	<div class="content content-top">
    <Card title="Transactions Detail" :loading="loading">
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
import {getBlock, getTx, getTxs} from '../api/serve'
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
		try {
			const res = await getTx({hash:this.$route.params.txhash})
			this.block = [{
				title:'Transaction Hash',
				value: res.hash
			},{
				title:'Status',
				value: res.transactions,
				render:({value})=>{
					const isSuccess = res.tx_result?.code===0
					return h('div',{
						style:{
							display:'flex',
							alignItems:'center'
						}
					},[
						h('img',{
							class:'iconfont',
							src:isSuccess?'/images/index/success@2x.png':'icon-error',
						}),
						h('span',{
							style:{
								marginLeft:'5px',
								color:isSuccess ? '#50C68D' : '#FF6060'
							}
						},isSuccess ? 'Success' : 'Fail')
					])
				}
			},{
				title:'Block',
				value: res.height
			},{
				title:'Timestamp',
				value: `${formatTime(res.block_time)}(${dayjs(res.block_time).format('MMM-DD-YYYY HH:mm:ss A [ +UTC]')})`
			},{
				title:'Form',
				value: res.tx_msg?.from_address ?? '-'
			},{
				title:'To',
				value: res.tx_msg?.to_address ?? '-'
			},{
				title:'Value',
				value: `${res.value_mis} MIS`,
			},{
				title:'Gas Fee',
				value: `${res.gas_used} MIS`,
			}] as dataItem[]
		} finally{
			this.loading = false;
		}
	}
}
</script>
<style lang="scss">
.iconfont{
	width: 14px;
	height: 14px;
}
</style>