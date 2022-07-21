import request from "../utils/request";

export function getIndexPageStats(){
  return request({
    url: '/api/v1/info',
  })
}
export function getBlocks(params){
  return request({
    url: '/api/v1/block/page',
    params
  })
}
export function getTxs(params){
  return request({
    url: 'api/v1/tx/page',
    params
  })
}
export function getTopHolder(params?:any){
  return request({
    url: '/api/v1/address/top',
    params
  })
}
export function getHolders(params?:any){
  return request({
    url: 'api/v1/address/page',
    params
  })
}
export function getBlock(params){
  return request({
    url: `/api/v1/block/${params.height}`
  })
}
export function getTx(params){
  return request({
    url: `api/v1/tx/${params.hash}`
  })
}
export function getAddress(params){
  return request({
    url: `/api/v1/address/${params.misesid}`,
  })
}