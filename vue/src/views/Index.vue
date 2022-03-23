<template>
	<div>
	
			
		<div class="container">
		
			<MisesWelcome />
			<SpTokenSend :address="address" />
			<SpTransferList :address="address" />
		</div>
		
	</div>

</template>

<script>
import MisesWelcome from '../components/MisesWelcome'
export default {
	components: {
		MisesWelcome
	},
	name: 'Index',
	computed: {
		address() {
			return this.$store.getters['common/wallet/address']
		}
	}
}
</script>
