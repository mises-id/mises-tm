import { createRouter, createWebHistory } from 'vue-router'
import Portfolio from '../views/Portfolio.vue'
import Data from '../views/Data.vue'
import Relayers from '../views/Relayers.vue'
import LatestBlocks from '../views/LatestBlocks.vue'
import Blocks from '../views/Blocks.vue'
import Block from '../views/Block.vue'
import Validators from '../views/Validators.vue'

const routerHistory = createWebHistory()
const routes = [
	{ path: '/', component: Portfolio },
	{ path: '/portfolio', component: Portfolio },
	{ path: '/data', component: Data },
	{ path: '/relayers', component: Relayers },
	{ path: '/latestblocks', component: LatestBlocks },
	{ path: '/blocks/:page?', component: Blocks },
	{ path: '/block/:block', component: Block },
	{ path: '/validators/:page?', component: Validators },
	{ path: '/address/:address', component: Data },
	{ path: '/tx/:txhash', component: Block },
]

const router = createRouter({
	history: routerHistory,
	routes
})

export default router
