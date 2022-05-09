<template>
	<div class="sp-component sp-welcome">
		<div class="sp-box sp-shadow sp-welcome__main">
			<div class="sp-welcome__header">
				<h2>Welcome to Mises</h2>
			</div>
			<p>Mises is a decentralized social platform</p>
			<SpBlockHeight />
			<div class="sp-welcome__btns">
				<SpButton href="https://github.com/mises-id/mises-tm" target="_blank" type="primary">Documentation</SpButton>
				<SpButton href="https://github.com/mises-id" target="_blank" type="secondary" class="sp-welcome__btns__github">Github</SpButton>
			</div>
		</div>
	</div>
</template>

<script>
import SpBlockHeight from '../SpBlockHeight'
import { defineComponent } from 'vue'
export default defineComponent({
	components: {
		SpBlockHeight
	},
	name: 'MisesWelcome'
})
</script>
