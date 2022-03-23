<template>
	<div class="container">
		<div class="sp-box sp-shadow">
			<h2>VALIDATORS</h2>
			<table class="SpTable">
				<thead>
					<tr>
						<th><strong>Status</strong></th>
						<th><strong>moniker</strong></th>
						<th><strong>OperatorAddress</strong></th>
						<th><strong>VotingPower</strong></th>
					</tr>
				</thead>
				<tbody>
					<tr class="sp-blockdisplayline" v-for="validator in validators" :key="validator.address">
						<td>
							<SpStatusLED :status="validator.status" style="margin-left: 45%" />
						</td>
						<td>
							{{ validator.moniker }}
						</td>
						<td>
							{{ validator.address }}
						</td>
						<td>
							{{ validator.voting_power }}
						</td>
					</tr>
				</tbody>
			</table>
		</div>
	</div>
</template>

<script>
import axios from 'axios'
export default {
	name: 'Validators',
	data() {
		return {
			validators: [],
			pages: 1
		}
	},
	computed: {
		depsLoaded() {
			return this._depsLoaded
		},
		page() {
			return parseInt(this.$route.params.page) || 1
		}
	},
	watch: {
		page: async function (newPage) {}
	},
	async mounted() {
		const page = await axios.get(this.$store.getters['common/env/apiCosmos'] + '/cosmos/staking/v1beta1/validators')
		for (let validator_meta of page.data.validators) {
			const validator = {
				moniker: validator_meta.description.moniker,
				address: validator_meta.operator_address,
				status: !validator_meta.jailed,
				voting_power: validator_meta.tokens / 1000000,
				proposer_priority: validator_meta.proposer_priority
			}
			this.validators.push(validator)
		}
	}
}
</script>
