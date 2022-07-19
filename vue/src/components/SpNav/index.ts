import { App as Application } from 'vue'
import SpNav from './SpNav.vue'

import { registerComponent } from './../../utils/plugins/index'

export const Plugin = {
  install(vue: Application): void {
    registerComponent(vue, SpNav)
  },
}

export default SpNav
