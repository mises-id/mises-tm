<script lang="ts">
import { defineComponent, h, PropType } from 'vue'
export default defineComponent({
  name:"MisesValue",
  props: {
    value: {
      default: h('div')
    }
  },
  render() {
    return this.value
  }
})
</script>