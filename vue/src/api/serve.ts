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
export function getTopHolder(params){
  return request({
    url: '/api/v1/address/top',
    params
  })
}