import { env, blocks, wallet } from '@starport/vuex'
import mises_generated from './mises-store/generated'
export default function init(store) {
	for (const moduleInit of Object.values(mises_generated)) {
		moduleInit(store)
	}
	blocks(store)
	env(store)
	wallet(store)
}
