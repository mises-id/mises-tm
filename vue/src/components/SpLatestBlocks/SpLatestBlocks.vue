<template>
  <div class="sp-component sp-latestblocks">
    <div class="sp-box sp-shadow sp-latestblocks__main" ref="blockList">
      <h2>LATEST BLOCKS</h2>
      <div v-if="blocks.length >= 10" class="sp-latestblocks__main__list">
        <SpBlockDisplaySmall
          v-for="block in blocks"
          :id="'block-' + block.height"
          :key="block.hash"
          :block="block"
          tsFormat="MMM D YYYY, HH:mm:ss"
        >
        </SpBlockDisplaySmall>
      </div>
      <div v-else class="sp-latestblocks__main__empty">
        <p>Generating blocks</p>
      </div>
      <SpButton type="primary" link="/blocks">View all blocks</SpButton>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent } from 'vue'
import { useStore } from 'vuex'
import SpBlockDisplaySmall from '../SpBlockDisplaySmall'
import {SpButton} from '@starport/vue'

export default defineComponent({
  name: 'SpLatestBlocks',
  components: {
    SpBlockDisplaySmall,
    SpButton,
  },
  setup() {
    let $s = useStore()
    let blocks = computed<Array<unknown> >(() => {
      return $s.getters['common/blocks/getBlocks'](10)
    })
    return {
      //computed
      blocks
    }
  },
  watch: {
    blocks: function (): void {
      let initialPos = window.scrollY
      this.$nextTick(() => {
        window.scrollTo(0, initialPos)
      })
    },
  },
})
</script>
