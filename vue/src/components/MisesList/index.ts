/*
 * @Author: lmk
 * @Date: 2022-04-20 16:39:08
 * @LastEditTime: 2022-04-20 16:59:46
 * @LastEditors: lmk
 * @Description: 
 */
import { App as Application } from 'vue'
import MisesList from './MisesList.vue'

import { registerComponent } from './../../utils/plugins/index'

export const Plugin = {
  install(vue: Application): void {
    registerComponent(vue, MisesList)
  },
}

export default MisesList
