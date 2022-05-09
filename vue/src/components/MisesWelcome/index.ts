import { App as Application } from 'vue'
import MisesWelcome from './MisesWelcome.vue'

import { registerComponent } from './../../utils/plugins/index'

export const Plugin = {
  install(vue: Application): void {
    registerComponent(vue, MisesWelcome)
  },
}

export default MisesWelcome
