<template>
	<SpSidebar
		v-on:sidebar-open="sidebarOpen = true"
		v-on:sidebar-close="sidebarOpen = false"
	>
		<template v-slot:default>
			<SpLinkIcon link="/" text="Dashboard" icon="Dashboard" />
			<SpLinkIcon link="/types" text="Mises Social Data" icon="Form" />
			<SpLinkIcon link="/latestblocks" text="Latest Blocks" icon="Transactions" />
			<div class="sp-dash"></div>
			<SpLinkIcon
				href="https://github.com/mises-id/doc"
				target="_blank"
				text="Documentation"
				icon="Docs"
			/>
		</template>
		<template v-slot:footer>
			<SpStatusAPI :showText="sidebarOpen" />
			<SpStatusRPC :showText="sidebarOpen" />
			<SpStatusWS :showText="sidebarOpen" />
			<div class="sp-text">Build: v0.3.8</div>
		</template>
	</SpSidebar>
</template>

<script>
export default {
	name: 'Sidebar',
	data() {
		return {
			sidebarOpen: true
		}
	},
	computed: {
		hasWallet() {
			return this.$store.hasModule(['common', 'wallet'])
		}
	}
}
</script>
