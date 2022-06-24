<template>
  <div v-if="moduleAvailable" class="container crud--position">
    <div class="row">
      <div class="col-6">
        <SpTypography modifier="highlight" size="md" style="font-weight: 700">
          {{ title }}
        </SpTypography>
      </div>
    </div>

    <SpCrudReadonlyList
      :store-name="storeName"
      :item-name="moduleNameNormalized"
      :command-name="commandName"
      :command-key="commandKey"
      :creator="address"
      :query="query"
    />

  </div>
</template>

<script lang="ts">
import { computed, defineComponent, reactive, ref, toRefs } from 'vue'
import { useStore } from 'vuex'
import SpCrudReadonlyList from '../SpCrudReadonlyList'
import {SpTypography, SpButton} from '@starport/vue'

export interface State {
  visibleModal: string
  activeItem: any
  moduleAvailable: boolean
}

export let initialState: State = {
  visibleModal: '',
  activeItem: {},
  moduleAvailable: false
}

export default defineComponent({
  name: 'SpCrudReadonly',

  components: {
    SpTypography,
    SpButton,
    SpCrudReadonlyList,
  },

  props: {
    displayName: {
      type: String,
      required: true
    },

    storeName: {
      type: String,
      required: true
    },

    itemName: {
      type: String,
      required: true
    },

    address: {
      type: String,
      required: true
    },
  },

  setup(props) {
    // store
    let $s = useStore()

    // composables
    let address = computed(() => {
        return props.address
      }
      
    )
    let title = computed(() => {
        return props.displayName
      }
      
    )

    // state
    let state: State = reactive(initialState)

    // computed
    let moduleNameNormalized = computed(() =>
      props.itemName.replace(/^\w/, (c) => c.toUpperCase())
    )

    state.moduleAvailable = $s.hasModule(props.storeName)

    let commandName = computed(() => {
      if (props.itemName == "UserInfo") {
        return "/RestQueryQueryUser"
      } else if (props.itemName == "Tx"){
        return "/ServiceGetTxsEvent"
      } else {
        return "/RestQueryQueryUserRelation"
      }
      
    })
    let commandKey = computed(() =>{
      if (props.itemName == "UserInfo") {
        return "QueryUser"
      } else if (props.itemName == "Tx"){
        return "GetTxsEvent"
      } else {
        return "QueryUserRelation"
      }
      
    })
    let query = computed(() =>{
      if (props.address.length == 0) {
        return {}
      }
      if (props.displayName == "User Info") {
        return {"mises_id": "did:mises:" + props.address}
      } else if (props.displayName == "Transactions Recv"){
        return {"events": "transfer.recipient='" + props.address + "'", "order_by": "ORDER_BY_DESC"}
      }  else if (props.displayName == "Transactions Send"){
        return {"events": "transfer.sender='" + props.address + "'", "order_by": "ORDER_BY_DESC"}
      } else {
        return {"mises_id": "did:mises:" + props.address}
      }
      
    })
    return {
      ...toRefs(state),
      title,
      address,
      moduleNameNormalized,
      commandName,
      commandKey,
      query
    }
  }
})
</script>

<style scoped lang="scss">
.crud {
  &--position {
    margin-top: 48px;
  }
}

.item-title {
  font-size: 13px;
  line-height: 153.8%;
  color: rgba(0, 0, 0, 0.667);
}

.item-value {
  font-size: 16px;
  line-height: 150%;
  color: #000000;
}

.dropdown-option {
  padding: 1rem 1.4rem;
}

.sp-label {
  display: block;
  text-align: left;
  width: 100%;
  margin: 0 4px;
  font-family: Inter;
  font-style: normal;
  font-weight: normal;
  font-size: 13px;
  line-height: 153.8%;
  /* identical to box height, or 20px */

  /* light/muted */

  color: rgba(0, 0, 0, 0.667);
}
.sp-input {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 12px 16px;
  width: 100%;
  height: 48px;
  background: rgba(0, 0, 0, 0.03);
  border-radius: 10px;
  margin: 4px 0px;
}
</style>
