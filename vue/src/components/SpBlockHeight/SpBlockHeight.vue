<template>
  <div class="SpBlockHeight" >
    Block Height: <strong>{{ blockHeight }}</strong>
  </div>
</template>
<script lang="ts">
import { defineComponent } from 'vue'
import { useStore } from 'vuex'
import { Block } from '../../utils/interfaces'

export default defineComponent({
  name: 'SpBlockHeight',
  computed: {
    blocks: function (): Array<Block> {
      let $s = useStore()
      return $s.getters['common/blocks/getBlocks'](10)
    },
    blockHeight: function (): number | 'N/A' {
      if (this.blocks.length > 0) {
        return this.blocks[0].height
      } else {
        return 'N/A'
      }
    }
  }
})
</script>
