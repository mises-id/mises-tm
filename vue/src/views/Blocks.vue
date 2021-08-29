<template>
	<div>
		<div class="container">
			<SpBlocksTable :blocks="blocks"/>
		</div>
	</div>
</template>

<script>
export default {
	name: 'Blocks',
	computed: {
		blocks() {
			if (this._depsLoaded) {
				return []
			} else {
				return []
			}
		},
		depsLoaded() {
			return this._depsLoaded
		}
	},
	beforeCreate() {
		const module = ['common', 'blocks'];
		for (let i = 1; i <= module.length; i++) {
			let submod = module.slice(0, i)
			if (!this.$store.hasModule(submod)) {
				console.log('Module `common.blocks` has not been registered!')
				this._depsLoaded = false
				break
			}
		}
	}
}
</script>
