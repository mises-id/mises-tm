import { createRouter, createWebHistory } from 'vue-router'
import Portfolio from '../views/Portfolio.vue'
import Blocks from '../views/Blocks.vue'
import Holders from '../views/Holders.vue'
import Block from '../views/Block.vue'
import Address from '../views/Address.vue'
import Validators from '../views/Validators.vue'
import Validator from '../views/Validator.vue'
import TXs from '../views/TXs.vue'
import Tx from '../views/Tx.vue'

const routerHistory = createWebHistory()
const routes = [
	{ path: '/', component: Portfolio },
	{ path: '/portfolio', component: Portfolio },

	{ path: '/blocks', component: Blocks },
	{ path: '/block/:block', component: Block },
	{ path: '/blockTx/:block', component: TXs },

	{ path: '/holders', component: Holders },
	{ path: '/holders/:misesid', component: Address },

	{ path: '/tx', component: TXs },
	{ path: '/tx/:txhash', component: Tx },

	{ path: '/validators', component: Validators },
	{ path: '/validators/:misesid', component: Validator },

]

const router = createRouter({
	history: routerHistory,
	routes
})

export default router
