import { createRouter, createWebHistory } from 'vue-router'
import Index from '@/views/Index.vue'
import Types from '@/views/Types.vue'
import Relayers from '@/views/Relayers.vue'
import LatestBlocks from '@/views/LatestBlocks.vue'
import Blocks from '@/views/Blocks.vue'
import Block from '@/views/Block.vue'
import Validators from '@/views/Validators.vue'

const routerHistory = createWebHistory()
const routes = [
	{
		path: '/',
		component: Index
	},
	{ path: '/types', component: Types },
	{ path: '/relayers', component: Relayers },
	{ path: '/latestblocks', component: LatestBlocks },
	{ path: '/blocks/:page?', component: Blocks },
	{ path: '/block/:block', component: Block },
	{ path: '/validators/:page?', component: Validators },
]

const router = createRouter({
	history: routerHistory,
	routes
})

export default router
