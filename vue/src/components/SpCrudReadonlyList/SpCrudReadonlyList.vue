<template>
  <div>

    <div v-if="items">
      <table border="none" class="table" v-if="(items || []).length">
        <thead>
          <tr>
            <th v-for="field in state.itemHeader" :key="field"><span class="table-content">{{field}}</span></th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="(item, index) in items" :key="index">
            <td  v-for="(value,key) in item" :key="key">
              <span class="table-content view" v-if="key==='More Info'" @click="triggerModel(value)">
                View
              </span>
              <span class="table-content" :class="key" v-else>
                {{value}}
              </span>
            </td>
          </tr>
        </tbody>
      </table>
      <!-- <div
        v-for="item in items"
        :key="item.id"
        style="
          display: flex;
          justify-content: space-between;
          gap: 14px;
          margin-bottom: 3rem;
        "
      >
        <div style="width: 100%">
          <div v-for="field in itemFields" :key="field.name">
            <div class="item-title capitalize-first-letter">
              {{ field.name }}
            </div>
            <div class="item-value">
              {{ item[field.name] }}

            </div>
            <SpSpacer size="xsm" />
          </div>
        </div>
      
      </div> -->
      <div v-if="(items || []).length === 0">
        <SpSpacer size="md" />
        <SpTypography size="md" class="empty">No items</SpTypography>
      </div>
    </div>
    <MisesDialog v-model:visible="state.dialogVisable" title="More Info">
      <div v-if="state.selectedItem">
        <div v-for="field in itemFields" :key="field.name">
          
        <div class="item-title capitalize-first-letter">
            {{ field.name }}
          </div>
          <div class="item-value">
            {{ state.selectedItem[field.name] }}
          </div>
        </div>
      </div>
    </MisesDialog>
  </div>
</template>

<script lang="ts">
// import { sentenceCase } from 'sentence-case'
import { computed, defineComponent, onMounted, reactive } from 'vue'
import { useStore } from 'vuex'

import {SpTypography, SpModal, SpButton, SpSpacer} from '@starport/vue'
import { precisionRound } from '../../utils/helpers'
import MisesDialog from '../MisesDialog'
export default defineComponent({
  name: 'SpCrudReadonlyList',

  components: {
    SpSpacer,
    SpTypography,
    SpButton,
    SpModal,
    MisesDialog
  },

  props: {
    storeName: {
      type: String,
      required: true
    },

    itemName: {
      type: String,
      required: true
    },

    query: {
      type: Object,
      required: true
    },

    commandName: {
      type: String,
      required: true
    },

    commandKey: {
      type: String,
      required: true
    },
    
    creator: {
      type: String,
      required: true
    },
  },

  setup(props) {
    // store
    let $s = useStore()

    // computed
    let itemFields = computed(() =>{
      const fields = $s.getters[props.storeName + '/getTypeStructure'](props.itemName)
      //get table keys from store ['pub_info','signatures',"body","auth_info"]
      const getFields = fields.filter(field => {
        if(['pub_info','signatures',"body","auth_info"].includes(field.name)) {
          return field
        }
      })
      return getFields;
    })

    let formatCoins = (data: any) => data.map(formatCoin).join(',')
    let formatCoin = ({ amount, denom }: any) => `${amount}${denom}`

    let defaultMessage = (type: string) => {
      const result = type.split('.')
      return sentenceCase(result[result.length - 1].replace('Msg', ''))
    }


    let summary = (msg: any): string  =>{
      const creator = props.creator
      console.log(creator)
      try {
        const msgType = msg["@type"]
        const data = msg

        switch (msgType) {
          case '/cosmos.bank.v1beta1.MsgSend': {
            const { amount, to_address, from_address } = data
            if (creator === to_address) {
              return `Recv ${formatCoins(amount)} from ${from_address}`
            }
            return `Send ${formatCoins(amount)} to ${to_address}`
          }

          case '/cosmos.staking.v1beta1.MsgDelegate': {
            const { amount, validator_address } = data
            return `Delegate ${formatCoin(amount)} to ${validator_address}`
          }

          case '/cosmos.staking.v1beta1.MsgBeginRedelegate': {
            const { amount, validator_dst_address, validator_src_address } = data
            return `Redelegate ${formatCoin(
              amount
            )} from ${validator_src_address} to ${validator_dst_address}`
          }

          case '/cosmos.staking.v1beta1.MsgUndelegate': {
            const { amount, validator_address } = data
            return `Undelegate ${formatCoin(amount)} from ${validator_address}`
          }

          case '/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward': {
            const { validator_address } = data
            return `Withdraw rewards from ${validator_address}`
          }

          default:
            return defaultMessage(msgType)
        }
      } catch {
        return ''
      }
    }

    // let itemFields = computed(() =>
    //   [
    //     {
    //       name: 'title',
    //     },
    //     {
    //       name: 'description'
    //     },
    //     {
    //       name: 'by'
    //     }
    //   ]
    // )
    const state = reactive({
      itemHeader:[],
      dialogVisable: false,
      selectedItem:null
    })
    let reloadSocial = async () => {
        $s.dispatch(`${props.storeName}${props.commandName}`, {
          options: { subscribe: false },
          params: {},
          query: props.query
        })
    }
    let items = computed(() => {
        if (Object.keys(props.query).length > 0) {
          
          // console.log("props.query", props.query)
          const itemData = $s.state[props.storeName][props.commandKey]
          let queryKey 
          Object.keys(itemData).forEach((key) => {
            // console.log("key", key)
            if (key.indexOf(Object.values(props.query)[0]) >= 0) {
              queryKey = key
            }
          })
          // console.log("queryKey", queryKey)
          if (!queryKey) {
            reloadSocial()
          } else {
            // console.log(itemData,queryKey,'queryKeyqueryKeyqueryKeyqueryKeyqueryKey')
            if (queryKey && itemData[queryKey]) {
              // console.log([itemData[queryKey]],'============')
              if(itemData[queryKey]['pub_info']){
                const key = Object.keys(itemData[queryKey]['pub_info']);
                state.itemHeader = key;
                return [itemData[queryKey]['pub_info']]
              }
              if (itemData[queryKey]["pagination"]) {
                const tx_responses = itemData[queryKey]["tx_responses"];
                const data = tx_responses.reduce((result, val)=>{
                  const {messages=[{}]} = val.tx.body;
                  return result.concat(messages.map(msg=>{
                    return {
                      hash: val.txhash,
                      blockHeight:val.height,
                      amount:msg.amount ? `${precisionRound((Array.isArray(msg.amount)?msg.amount[0]:msg.amount).amount / 1000000, 6)}MIS` : '',

                      summary: summary(msg),
                      gas_used:val.gas_used,
                      status: val.code===0 ? 'success' : 'failed',
                      [`More Info`]:val.tx,
                    }
                  }))
                } , []);
                if(data.length){
                  const key = Object.keys(data[0])
                  state.itemHeader = key;
                  return data
                }
                return data;
              }
              return [itemData[queryKey]]
            }
          }

        }
      return []
    })
    const triggerModel = (item)=>{
      state.dialogVisable = !state.dialogVisable
      if(state.dialogVisable&&item){
        state.selectedItem = item
      }
    }
    return {
      itemFields,
      items,
      state,
      triggerModel
    }
  }
})
</script>

<style scoped lang="scss">
.page-background {
  background: white;
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

.dropdown-option:not(:last-child) {
  padding: 0 0 16px 0;
}

.sp-label {
  display: block;
  text-align: left;
  width: 100%;
  margin: 0 4px;
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

.capitalize-first-letter:first-letter {
  text-transform: uppercase;
}

.empty {
  font-size: 16px;
  color: rgba(0, 0, 0, 0.667);
}
.table-content{
  display: block;
  padding: 13px 10px;
  font-size: 13px;
  text-align: center;
  word-break: break-all;
}
$borderColor:#ebeef5;
.table{
  margin-top: 20px;
  border-collapse: collapse;
  width:100%;
  border:1px solid $borderColor ;
  margin-bottom:20px;
  th{
    border-collapse: collapse;
    border-right:1px solid $borderColor ;
    border-bottom:1px solid $borderColor ;
    background-color:#f5f7fa ; 
    font-size:14px;
    font-weight:normal;
    text-align:center;
    white-space: nowrap;
  }
  td{
    border-collapse: collapse;
    border-right:1px solid $borderColor ;
    border-bottom:1px solid $borderColor ;
    font-size:12px;
    font-weight:normal;
    text-align:center;
    word-break: break-all;
  }
  .status,.amount,info{
    white-space: nowrap;
  }
  .view{
    color: #1989fa;
    cursor: pointer;
  }
}
</style>