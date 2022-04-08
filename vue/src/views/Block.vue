<template>
	<div>
		<div class="container">
			<SpBlockDisplayFull :block="block" />
		</div>
	</div>
</template>

<script>
import SpBlockDisplayFull from '../components/SpBlockDisplayFull'
import axios from 'axios'
import { fromBase64, toHex } from '@cosmjs/encoding'
import { sha256 } from '@cosmjs/crypto'
function formatTx({
	txHash = '',
	messages = [],
	memo = '',
	signer_infos = [],
	fee = {},
	gas_used = null,
	gas_wanted = null,
	height = null,
	code = 0,
	log = null
}) {
	return {
		txHash,
		body: {
			messages,
			memo
		},
		auth_info: {
			signer_infos,
			fee
		},
		meta: {
			gas_used,
			gas_wanted,
			height,
			code,
			log
		}
	}
}
async function getTx(apiCosmos, apiTendermint, encodedTx) {
	const txHash = sha256(fromBase64(encodedTx))
	try {
		const rpcRes = await axios.get(apiTendermint + '/tx?hash=0x' + toHex(txHash))
		const apiRes = await axios.get(apiCosmos + '/cosmos/tx/v1beta1/txs/' + toHex(txHash))
		return { rpcRes, apiRes, txHash: toHex(txHash).toUpperCase() }
	} catch (e) {
		throw 'Error fetching TX data'
	}
}
async function decodeTx(apiCosmos, apiTendermint, encodedTx) {
	const fullTx = await getTx(apiCosmos, apiTendermint, encodedTx)
	const { data } = fullTx.rpcRes
	const { height, tx_result } = data.result
	const { code, log, gas_used, gas_wanted } = tx_result
	const { body, auth_info } = fullTx.apiRes.data.tx
	const { messages, memo } = body

	return formatTx({
		txHash: fullTx.txHash,
		messages,
		memo,
		signer_infos: auth_info.signer_infos,
		fee: auth_info.fee,
		gas_used,
		gas_wanted,
		height,
		code,
		log
	})
}
export default {
	components: {
		SpBlockDisplayFull
	},
	data() {
		return {
			block: {height:0,timestamp:"",hash:"",details:"",txDecoded:[]}
		}
	},
	async created() {
		const blockDetails = await axios.get(this.$store.getters['common/env/apiTendermint'] + '/block?height=' + this.$route.params.block)
		const txDecoded = blockDetails.data.result.block.data.txs.map(async (x) => {
			const dec = await decodeTx(this.$store.getters['common/env/apiCosmos'], this.$store.getters['common/env/apiTendermint'], x)
			return dec
		})
		const txs = await Promise.all(txDecoded)
		this.block = {
			height: blockDetails.data.result.block.header.height,
			timestamp: blockDetails.data.result.block.header.time,
			hash: blockDetails.data.result.block_id.hash,
			details: blockDetails.data.result.block,
			txDecoded: [...txs]
		}
	}
}
</script>
