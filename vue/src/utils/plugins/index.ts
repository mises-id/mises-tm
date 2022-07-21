import { App as Application, Component } from 'vue'

export const registerComponent = (instance: Application, component: Component): void => {
  if (component) {
    instance.component(component.name || '', component)
  }
}

export const registerComponentProgrammatic = (instance: Application, property: string, component: Component): void => {
  if (!instance.config.globalProperties.$test) instance.config.globalProperties.$test = {}
  instance.config.globalProperties.$test[property] = component
}
export function shortenAddress(
  address = '',
  prefix = 8,
) {
  if (address.length < 44) {
    return address;
  }

  return `${address.slice(0, prefix)}...${address.slice(
    -4,
  )}`;
}
import relativeTime from 'dayjs/plugin/relativeTime';
import dayjs from 'dayjs'
dayjs.extend(relativeTime)
export function formatTime(time){
  if(!time) return ''
  const timestamp = dayjs(time).valueOf()
  const now = dayjs().valueOf()
  const diff = now - timestamp
  if(diff < 60000){
    return Math.floor(diff / 1000) + ' sec ago'
  }
  if(diff < 3600000){
    return `${Math.floor(diff/60000)} min ago`
  }
  if(diff < 86400000){
    return `${Math.floor(diff/3600000)} hours ago`
  }
  return dayjs(time).format('YYYY.MM.DD HH:mm:ss')
}

import { useStore } from 'vuex'
import { computed } from 'vue'
import BigNumber from 'bignumber.js'
export function useSupply(){
  const $s = useStore()
  const supply = computed(() => {
    const data = $s.getters['cosmos.bank.v1beta1/getTotalSupply']()
    const supply = data.supply?.[0]
    const totalSupply = supply?.amount ? new BigNumber(supply.amount).div(1000000).integerValue().toString() : 0
    return totalSupply
  })
  return supply
}
export const getCalcVotingPowerRate = (TerraValidators: any[]) => {
  TerraValidators = TerraValidators.map(val=>{
    return {
      ...val,
      voting_power:new BigNumber(val.tokens).div(1000000).toString()
    }
  })
  const total = BigNumber.sum(
    ...TerraValidators.map(val => val.voting_power)
  ).toNumber()
  return (address: any) => {
    const validator = TerraValidators.find(
      ({ operator_address }) => operator_address === address
    )

    if (!validator) return
    const { voting_power } = validator
    return voting_power ? Number(validator.voting_power) / total : undefined
  }
}
import axios from 'axios'
export const useCalcVotingPower = () => {
  const $s = useStore()
  console.log($s)
  const init = async ()=>{
    const page = await axios.get($s.getters['common/env/apiCosmos'] + '/cosmos/staking/v1beta1/validators')
    return getCalcVotingPowerRate(page.data.validators)
  }
  return init()
}
import jscrypto from 'jscrypto'
import {Buffer} from 'Buffer'
export const rawAddress = (key) => {
  if(!key) return []
  var pubkeyData = Buffer.from(key, 'base64');
  return sha256(pubkeyData).slice(0, 20)
}
function sha256(data) {
  return jscrypto.SHA256.hash(new jscrypto.Word32Array(data)).toUint8Array();
}
import bech32 from 'bech32'
export const valConsAddress = (val) => {
  const data = rawAddress(val?.consensus_pubkey.key) as Uint8Array;
  return val?.consensus_pubkey ? bech32.encode('misesvalcons', bech32.toWords(data)) : ''
}
export function uptime_estimated(info): number {
  if (!info) {
    return 0
  }
  const { missed_blocks_counter } = info
  if (missed_blocks_counter) {
    return 1 - Number(missed_blocks_counter) / 10000
  }
  return 1
}
export const getBondedMISCount = (
  TerraValidators: any[],
  address: string
) => {

  TerraValidators = TerraValidators.map(val=>{
    return {
      ...val,
      voting_power:new BigNumber(val.tokens).div(1000000).toString()
    }
  })
  const validator = TerraValidators.find(
    ({ operator_address }) => operator_address === address
  )

  if (!validator) return
  const { voting_power } = validator
  return voting_power ? Number(validator.voting_power) : undefined
}