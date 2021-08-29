import { env, blocks, wallet, transfers, relayers } from '@starport/vuex'
import generated from './generated'
import mises_generated from './mises-store/generated'
export default function init(store) {
	for (const moduleInit of Object.values(generated)) {
		moduleInit(store)
	}
	for (const moduleInit of Object.values(mises_generated)) {
		moduleInit(store)
	}
	transfers(store)
	blocks(store)
	env(store)
	wallet(store)
	relayers(store)
}
