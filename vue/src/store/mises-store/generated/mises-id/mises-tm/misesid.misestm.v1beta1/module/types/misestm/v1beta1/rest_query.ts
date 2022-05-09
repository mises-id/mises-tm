/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import { Timestamp } from '../../google/protobuf/timestamp'
import * as Long from 'long'
import { DidRegistry } from '../../misestm/v1beta1/DidRegistry'
import { PublicUserInfo, PrivateUserInfo } from '../../misestm/v1beta1/UserInfo'
import { PageRequest, PageResponse } from '../../cosmos/base/query/v1beta1/pagination'
import { PublicAppInfo } from '../../misestm/v1beta1/AppInfo'
import { TxResponse } from '../../cosmos/base/abci/v1beta1/abci'
import { Coin } from '../../cosmos/base/v1beta1/coin'
import { Duration } from '../../google/protobuf/duration'

export const protobufPackage = 'misesid.misestm.v1beta1'

export interface RestQueryDidRequest {
  mises_id: string
}

export interface RestQueryDidResponse {
  didRegistry: DidRegistry | undefined
}

export interface RestQueryUserRequest {
  mises_uid: string
}

export interface RestQueryUserResponse {
  pub_info: PublicUserInfo | undefined
  pri_info: PrivateUserInfo | undefined
  version: number
}

export interface RestQueryUserRelationRequest {
  mises_uid: string
  filter: string
  pagination: PageRequest | undefined
}

export interface MisesID {
  mises_id: string
  rel_type: string
}

export interface RestQueryUserRelationResponse {
  mises_list: MisesID[]
  pagination: PageResponse | undefined
}

export interface RestQueryAppRequest {
  mises_appid: string
}

export interface RestQueryAppResponse {
  pub_info: PublicAppInfo | undefined
  version: number
}

export interface RestQueryTxRequest {
  txhash: string
}

export interface RestTxResponse {
  /** tx_response is the queried TxResponses. */
  tx_response: TxResponse | undefined
  code: number
}

export interface RestQueryAppFeeGrantRequest {
  mises_appid: string
  mises_uid: string
}

export interface AppFeeGrant {
  spend_limit: Coin | undefined
  period: Duration | undefined
  expiration: Date | undefined
}

export interface RestQueryAppFeeGrantResponse {
  grant: AppFeeGrant | undefined
}

const baseRestQueryDidRequest: object = { mises_id: '' }

export const RestQueryDidRequest = {
  encode(message: RestQueryDidRequest, writer: Writer = Writer.create()): Writer {
    if (message.mises_id !== '') {
      writer.uint32(10).string(message.mises_id)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): RestQueryDidRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseRestQueryDidRequest } as RestQueryDidRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.mises_id = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): RestQueryDidRequest {
    const message = { ...baseRestQueryDidRequest } as RestQueryDidRequest
    if (object.mises_id !== undefined && object.mises_id !== null) {
      message.mises_id = String(object.mises_id)
    } else {
      message.mises_id = ''
    }
    return message
  },

  toJSON(message: RestQueryDidRequest): unknown {
    const obj: any = {}
    message.mises_id !== undefined && (obj.mises_id = message.mises_id)
    return obj
  },

  fromPartial(object: DeepPartial<RestQueryDidRequest>): RestQueryDidRequest {
    const message = { ...baseRestQueryDidRequest } as RestQueryDidRequest
    if (object.mises_id !== undefined && object.mises_id !== null) {
      message.mises_id = object.mises_id
    } else {
      message.mises_id = ''
    }
    return message
  }
}

const baseRestQueryDidResponse: object = {}

export const RestQueryDidResponse = {
  encode(message: RestQueryDidResponse, writer: Writer = Writer.create()): Writer {
    if (message.didRegistry !== undefined) {
      DidRegistry.encode(message.didRegistry, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): RestQueryDidResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseRestQueryDidResponse } as RestQueryDidResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.didRegistry = DidRegistry.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): RestQueryDidResponse {
    const message = { ...baseRestQueryDidResponse } as RestQueryDidResponse
    if (object.didRegistry !== undefined && object.didRegistry !== null) {
      message.didRegistry = DidRegistry.fromJSON(object.didRegistry)
    } else {
      message.didRegistry = undefined
    }
    return message
  },

  toJSON(message: RestQueryDidResponse): unknown {
    const obj: any = {}
    message.didRegistry !== undefined && (obj.didRegistry = message.didRegistry ? DidRegistry.toJSON(message.didRegistry) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<RestQueryDidResponse>): RestQueryDidResponse {
    const message = { ...baseRestQueryDidResponse } as RestQueryDidResponse
    if (object.didRegistry !== undefined && object.didRegistry !== null) {
      message.didRegistry = DidRegistry.fromPartial(object.didRegistry)
    } else {
      message.didRegistry = undefined
    }
    return message
  }
}

const baseRestQueryUserRequest: object = { mises_uid: '' }

export const RestQueryUserRequest = {
  encode(message: RestQueryUserRequest, writer: Writer = Writer.create()): Writer {
    if (message.mises_uid !== '') {
      writer.uint32(10).string(message.mises_uid)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): RestQueryUserRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseRestQueryUserRequest } as RestQueryUserRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.mises_uid = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): RestQueryUserRequest {
    const message = { ...baseRestQueryUserRequest } as RestQueryUserRequest
    if (object.mises_uid !== undefined && object.mises_uid !== null) {
      message.mises_uid = String(object.mises_uid)
    } else {
      message.mises_uid = ''
    }
    return message
  },

  toJSON(message: RestQueryUserRequest): unknown {
    const obj: any = {}
    message.mises_uid !== undefined && (obj.mises_uid = message.mises_uid)
    return obj
  },

  fromPartial(object: DeepPartial<RestQueryUserRequest>): RestQueryUserRequest {
    const message = { ...baseRestQueryUserRequest } as RestQueryUserRequest
    if (object.mises_uid !== undefined && object.mises_uid !== null) {
      message.mises_uid = object.mises_uid
    } else {
      message.mises_uid = ''
    }
    return message
  }
}

const baseRestQueryUserResponse: object = { version: 0 }

export const RestQueryUserResponse = {
  encode(message: RestQueryUserResponse, writer: Writer = Writer.create()): Writer {
    if (message.pub_info !== undefined) {
      PublicUserInfo.encode(message.pub_info, writer.uint32(10).fork()).ldelim()
    }
    if (message.pri_info !== undefined) {
      PrivateUserInfo.encode(message.pri_info, writer.uint32(18).fork()).ldelim()
    }
    if (message.version !== 0) {
      writer.uint32(24).uint64(message.version)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): RestQueryUserResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseRestQueryUserResponse } as RestQueryUserResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pub_info = PublicUserInfo.decode(reader, reader.uint32())
          break
        case 2:
          message.pri_info = PrivateUserInfo.decode(reader, reader.uint32())
          break
        case 3:
          message.version = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): RestQueryUserResponse {
    const message = { ...baseRestQueryUserResponse } as RestQueryUserResponse
    if (object.pub_info !== undefined && object.pub_info !== null) {
      message.pub_info = PublicUserInfo.fromJSON(object.pub_info)
    } else {
      message.pub_info = undefined
    }
    if (object.pri_info !== undefined && object.pri_info !== null) {
      message.pri_info = PrivateUserInfo.fromJSON(object.pri_info)
    } else {
      message.pri_info = undefined
    }
    if (object.version !== undefined && object.version !== null) {
      message.version = Number(object.version)
    } else {
      message.version = 0
    }
    return message
  },

  toJSON(message: RestQueryUserResponse): unknown {
    const obj: any = {}
    message.pub_info !== undefined && (obj.pub_info = message.pub_info ? PublicUserInfo.toJSON(message.pub_info) : undefined)
    message.pri_info !== undefined && (obj.pri_info = message.pri_info ? PrivateUserInfo.toJSON(message.pri_info) : undefined)
    message.version !== undefined && (obj.version = message.version)
    return obj
  },

  fromPartial(object: DeepPartial<RestQueryUserResponse>): RestQueryUserResponse {
    const message = { ...baseRestQueryUserResponse } as RestQueryUserResponse
    if (object.pub_info !== undefined && object.pub_info !== null) {
      message.pub_info = PublicUserInfo.fromPartial(object.pub_info)
    } else {
      message.pub_info = undefined
    }
    if (object.pri_info !== undefined && object.pri_info !== null) {
      message.pri_info = PrivateUserInfo.fromPartial(object.pri_info)
    } else {
      message.pri_info = undefined
    }
    if (object.version !== undefined && object.version !== null) {
      message.version = object.version
    } else {
      message.version = 0
    }
    return message
  }
}

const baseRestQueryUserRelationRequest: object = { mises_uid: '', filter: '' }

export const RestQueryUserRelationRequest = {
  encode(message: RestQueryUserRelationRequest, writer: Writer = Writer.create()): Writer {
    if (message.mises_uid !== '') {
      writer.uint32(10).string(message.mises_uid)
    }
    if (message.filter !== '') {
      writer.uint32(18).string(message.filter)
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(26).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): RestQueryUserRelationRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseRestQueryUserRelationRequest } as RestQueryUserRelationRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.mises_uid = reader.string()
          break
        case 2:
          message.filter = reader.string()
          break
        case 3:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): RestQueryUserRelationRequest {
    const message = { ...baseRestQueryUserRelationRequest } as RestQueryUserRelationRequest
    if (object.mises_uid !== undefined && object.mises_uid !== null) {
      message.mises_uid = String(object.mises_uid)
    } else {
      message.mises_uid = ''
    }
    if (object.filter !== undefined && object.filter !== null) {
      message.filter = String(object.filter)
    } else {
      message.filter = ''
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: RestQueryUserRelationRequest): unknown {
    const obj: any = {}
    message.mises_uid !== undefined && (obj.mises_uid = message.mises_uid)
    message.filter !== undefined && (obj.filter = message.filter)
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<RestQueryUserRelationRequest>): RestQueryUserRelationRequest {
    const message = { ...baseRestQueryUserRelationRequest } as RestQueryUserRelationRequest
    if (object.mises_uid !== undefined && object.mises_uid !== null) {
      message.mises_uid = object.mises_uid
    } else {
      message.mises_uid = ''
    }
    if (object.filter !== undefined && object.filter !== null) {
      message.filter = object.filter
    } else {
      message.filter = ''
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseMisesID: object = { mises_id: '', rel_type: '' }

export const MisesID = {
  encode(message: MisesID, writer: Writer = Writer.create()): Writer {
    if (message.mises_id !== '') {
      writer.uint32(10).string(message.mises_id)
    }
    if (message.rel_type !== '') {
      writer.uint32(18).string(message.rel_type)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MisesID {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMisesID } as MisesID
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.mises_id = reader.string()
          break
        case 2:
          message.rel_type = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MisesID {
    const message = { ...baseMisesID } as MisesID
    if (object.mises_id !== undefined && object.mises_id !== null) {
      message.mises_id = String(object.mises_id)
    } else {
      message.mises_id = ''
    }
    if (object.rel_type !== undefined && object.rel_type !== null) {
      message.rel_type = String(object.rel_type)
    } else {
      message.rel_type = ''
    }
    return message
  },

  toJSON(message: MisesID): unknown {
    const obj: any = {}
    message.mises_id !== undefined && (obj.mises_id = message.mises_id)
    message.rel_type !== undefined && (obj.rel_type = message.rel_type)
    return obj
  },

  fromPartial(object: DeepPartial<MisesID>): MisesID {
    const message = { ...baseMisesID } as MisesID
    if (object.mises_id !== undefined && object.mises_id !== null) {
      message.mises_id = object.mises_id
    } else {
      message.mises_id = ''
    }
    if (object.rel_type !== undefined && object.rel_type !== null) {
      message.rel_type = object.rel_type
    } else {
      message.rel_type = ''
    }
    return message
  }
}

const baseRestQueryUserRelationResponse: object = {}

export const RestQueryUserRelationResponse = {
  encode(message: RestQueryUserRelationResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.mises_list) {
      MisesID.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): RestQueryUserRelationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseRestQueryUserRelationResponse } as RestQueryUserRelationResponse
    message.mises_list = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.mises_list.push(MisesID.decode(reader, reader.uint32()))
          break
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): RestQueryUserRelationResponse {
    const message = { ...baseRestQueryUserRelationResponse } as RestQueryUserRelationResponse
    message.mises_list = []
    if (object.mises_list !== undefined && object.mises_list !== null) {
      for (const e of object.mises_list) {
        message.mises_list.push(MisesID.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: RestQueryUserRelationResponse): unknown {
    const obj: any = {}
    if (message.mises_list) {
      obj.mises_list = message.mises_list.map((e) => (e ? MisesID.toJSON(e) : undefined))
    } else {
      obj.mises_list = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<RestQueryUserRelationResponse>): RestQueryUserRelationResponse {
    const message = { ...baseRestQueryUserRelationResponse } as RestQueryUserRelationResponse
    message.mises_list = []
    if (object.mises_list !== undefined && object.mises_list !== null) {
      for (const e of object.mises_list) {
        message.mises_list.push(MisesID.fromPartial(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseRestQueryAppRequest: object = { mises_appid: '' }

export const RestQueryAppRequest = {
  encode(message: RestQueryAppRequest, writer: Writer = Writer.create()): Writer {
    if (message.mises_appid !== '') {
      writer.uint32(10).string(message.mises_appid)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): RestQueryAppRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseRestQueryAppRequest } as RestQueryAppRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.mises_appid = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): RestQueryAppRequest {
    const message = { ...baseRestQueryAppRequest } as RestQueryAppRequest
    if (object.mises_appid !== undefined && object.mises_appid !== null) {
      message.mises_appid = String(object.mises_appid)
    } else {
      message.mises_appid = ''
    }
    return message
  },

  toJSON(message: RestQueryAppRequest): unknown {
    const obj: any = {}
    message.mises_appid !== undefined && (obj.mises_appid = message.mises_appid)
    return obj
  },

  fromPartial(object: DeepPartial<RestQueryAppRequest>): RestQueryAppRequest {
    const message = { ...baseRestQueryAppRequest } as RestQueryAppRequest
    if (object.mises_appid !== undefined && object.mises_appid !== null) {
      message.mises_appid = object.mises_appid
    } else {
      message.mises_appid = ''
    }
    return message
  }
}

const baseRestQueryAppResponse: object = { version: 0 }

export const RestQueryAppResponse = {
  encode(message: RestQueryAppResponse, writer: Writer = Writer.create()): Writer {
    if (message.pub_info !== undefined) {
      PublicAppInfo.encode(message.pub_info, writer.uint32(10).fork()).ldelim()
    }
    if (message.version !== 0) {
      writer.uint32(16).uint64(message.version)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): RestQueryAppResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseRestQueryAppResponse } as RestQueryAppResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pub_info = PublicAppInfo.decode(reader, reader.uint32())
          break
        case 2:
          message.version = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): RestQueryAppResponse {
    const message = { ...baseRestQueryAppResponse } as RestQueryAppResponse
    if (object.pub_info !== undefined && object.pub_info !== null) {
      message.pub_info = PublicAppInfo.fromJSON(object.pub_info)
    } else {
      message.pub_info = undefined
    }
    if (object.version !== undefined && object.version !== null) {
      message.version = Number(object.version)
    } else {
      message.version = 0
    }
    return message
  },

  toJSON(message: RestQueryAppResponse): unknown {
    const obj: any = {}
    message.pub_info !== undefined && (obj.pub_info = message.pub_info ? PublicAppInfo.toJSON(message.pub_info) : undefined)
    message.version !== undefined && (obj.version = message.version)
    return obj
  },

  fromPartial(object: DeepPartial<RestQueryAppResponse>): RestQueryAppResponse {
    const message = { ...baseRestQueryAppResponse } as RestQueryAppResponse
    if (object.pub_info !== undefined && object.pub_info !== null) {
      message.pub_info = PublicAppInfo.fromPartial(object.pub_info)
    } else {
      message.pub_info = undefined
    }
    if (object.version !== undefined && object.version !== null) {
      message.version = object.version
    } else {
      message.version = 0
    }
    return message
  }
}

const baseRestQueryTxRequest: object = { txhash: '' }

export const RestQueryTxRequest = {
  encode(message: RestQueryTxRequest, writer: Writer = Writer.create()): Writer {
    if (message.txhash !== '') {
      writer.uint32(10).string(message.txhash)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): RestQueryTxRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseRestQueryTxRequest } as RestQueryTxRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.txhash = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): RestQueryTxRequest {
    const message = { ...baseRestQueryTxRequest } as RestQueryTxRequest
    if (object.txhash !== undefined && object.txhash !== null) {
      message.txhash = String(object.txhash)
    } else {
      message.txhash = ''
    }
    return message
  },

  toJSON(message: RestQueryTxRequest): unknown {
    const obj: any = {}
    message.txhash !== undefined && (obj.txhash = message.txhash)
    return obj
  },

  fromPartial(object: DeepPartial<RestQueryTxRequest>): RestQueryTxRequest {
    const message = { ...baseRestQueryTxRequest } as RestQueryTxRequest
    if (object.txhash !== undefined && object.txhash !== null) {
      message.txhash = object.txhash
    } else {
      message.txhash = ''
    }
    return message
  }
}

const baseRestTxResponse: object = { code: 0 }

export const RestTxResponse = {
  encode(message: RestTxResponse, writer: Writer = Writer.create()): Writer {
    if (message.tx_response !== undefined) {
      TxResponse.encode(message.tx_response, writer.uint32(10).fork()).ldelim()
    }
    if (message.code !== 0) {
      writer.uint32(16).uint32(message.code)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): RestTxResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseRestTxResponse } as RestTxResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.tx_response = TxResponse.decode(reader, reader.uint32())
          break
        case 2:
          message.code = reader.uint32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): RestTxResponse {
    const message = { ...baseRestTxResponse } as RestTxResponse
    if (object.tx_response !== undefined && object.tx_response !== null) {
      message.tx_response = TxResponse.fromJSON(object.tx_response)
    } else {
      message.tx_response = undefined
    }
    if (object.code !== undefined && object.code !== null) {
      message.code = Number(object.code)
    } else {
      message.code = 0
    }
    return message
  },

  toJSON(message: RestTxResponse): unknown {
    const obj: any = {}
    message.tx_response !== undefined && (obj.tx_response = message.tx_response ? TxResponse.toJSON(message.tx_response) : undefined)
    message.code !== undefined && (obj.code = message.code)
    return obj
  },

  fromPartial(object: DeepPartial<RestTxResponse>): RestTxResponse {
    const message = { ...baseRestTxResponse } as RestTxResponse
    if (object.tx_response !== undefined && object.tx_response !== null) {
      message.tx_response = TxResponse.fromPartial(object.tx_response)
    } else {
      message.tx_response = undefined
    }
    if (object.code !== undefined && object.code !== null) {
      message.code = object.code
    } else {
      message.code = 0
    }
    return message
  }
}

const baseRestQueryAppFeeGrantRequest: object = { mises_appid: '', mises_uid: '' }

export const RestQueryAppFeeGrantRequest = {
  encode(message: RestQueryAppFeeGrantRequest, writer: Writer = Writer.create()): Writer {
    if (message.mises_appid !== '') {
      writer.uint32(10).string(message.mises_appid)
    }
    if (message.mises_uid !== '') {
      writer.uint32(18).string(message.mises_uid)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): RestQueryAppFeeGrantRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseRestQueryAppFeeGrantRequest } as RestQueryAppFeeGrantRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.mises_appid = reader.string()
          break
        case 2:
          message.mises_uid = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): RestQueryAppFeeGrantRequest {
    const message = { ...baseRestQueryAppFeeGrantRequest } as RestQueryAppFeeGrantRequest
    if (object.mises_appid !== undefined && object.mises_appid !== null) {
      message.mises_appid = String(object.mises_appid)
    } else {
      message.mises_appid = ''
    }
    if (object.mises_uid !== undefined && object.mises_uid !== null) {
      message.mises_uid = String(object.mises_uid)
    } else {
      message.mises_uid = ''
    }
    return message
  },

  toJSON(message: RestQueryAppFeeGrantRequest): unknown {
    const obj: any = {}
    message.mises_appid !== undefined && (obj.mises_appid = message.mises_appid)
    message.mises_uid !== undefined && (obj.mises_uid = message.mises_uid)
    return obj
  },

  fromPartial(object: DeepPartial<RestQueryAppFeeGrantRequest>): RestQueryAppFeeGrantRequest {
    const message = { ...baseRestQueryAppFeeGrantRequest } as RestQueryAppFeeGrantRequest
    if (object.mises_appid !== undefined && object.mises_appid !== null) {
      message.mises_appid = object.mises_appid
    } else {
      message.mises_appid = ''
    }
    if (object.mises_uid !== undefined && object.mises_uid !== null) {
      message.mises_uid = object.mises_uid
    } else {
      message.mises_uid = ''
    }
    return message
  }
}

const baseAppFeeGrant: object = {}

export const AppFeeGrant = {
  encode(message: AppFeeGrant, writer: Writer = Writer.create()): Writer {
    if (message.spend_limit !== undefined) {
      Coin.encode(message.spend_limit, writer.uint32(10).fork()).ldelim()
    }
    if (message.period !== undefined) {
      Duration.encode(message.period, writer.uint32(18).fork()).ldelim()
    }
    if (message.expiration !== undefined) {
      Timestamp.encode(toTimestamp(message.expiration), writer.uint32(26).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): AppFeeGrant {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseAppFeeGrant } as AppFeeGrant
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.spend_limit = Coin.decode(reader, reader.uint32())
          break
        case 2:
          message.period = Duration.decode(reader, reader.uint32())
          break
        case 3:
          message.expiration = fromTimestamp(Timestamp.decode(reader, reader.uint32()))
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): AppFeeGrant {
    const message = { ...baseAppFeeGrant } as AppFeeGrant
    if (object.spend_limit !== undefined && object.spend_limit !== null) {
      message.spend_limit = Coin.fromJSON(object.spend_limit)
    } else {
      message.spend_limit = undefined
    }
    if (object.period !== undefined && object.period !== null) {
      message.period = Duration.fromJSON(object.period)
    } else {
      message.period = undefined
    }
    if (object.expiration !== undefined && object.expiration !== null) {
      message.expiration = fromJsonTimestamp(object.expiration)
    } else {
      message.expiration = undefined
    }
    return message
  },

  toJSON(message: AppFeeGrant): unknown {
    const obj: any = {}
    message.spend_limit !== undefined && (obj.spend_limit = message.spend_limit ? Coin.toJSON(message.spend_limit) : undefined)
    message.period !== undefined && (obj.period = message.period ? Duration.toJSON(message.period) : undefined)
    message.expiration !== undefined && (obj.expiration = message.expiration !== undefined ? message.expiration.toISOString() : null)
    return obj
  },

  fromPartial(object: DeepPartial<AppFeeGrant>): AppFeeGrant {
    const message = { ...baseAppFeeGrant } as AppFeeGrant
    if (object.spend_limit !== undefined && object.spend_limit !== null) {
      message.spend_limit = Coin.fromPartial(object.spend_limit)
    } else {
      message.spend_limit = undefined
    }
    if (object.period !== undefined && object.period !== null) {
      message.period = Duration.fromPartial(object.period)
    } else {
      message.period = undefined
    }
    if (object.expiration !== undefined && object.expiration !== null) {
      message.expiration = object.expiration
    } else {
      message.expiration = undefined
    }
    return message
  }
}

const baseRestQueryAppFeeGrantResponse: object = {}

export const RestQueryAppFeeGrantResponse = {
  encode(message: RestQueryAppFeeGrantResponse, writer: Writer = Writer.create()): Writer {
    if (message.grant !== undefined) {
      AppFeeGrant.encode(message.grant, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): RestQueryAppFeeGrantResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseRestQueryAppFeeGrantResponse } as RestQueryAppFeeGrantResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.grant = AppFeeGrant.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): RestQueryAppFeeGrantResponse {
    const message = { ...baseRestQueryAppFeeGrantResponse } as RestQueryAppFeeGrantResponse
    if (object.grant !== undefined && object.grant !== null) {
      message.grant = AppFeeGrant.fromJSON(object.grant)
    } else {
      message.grant = undefined
    }
    return message
  },

  toJSON(message: RestQueryAppFeeGrantResponse): unknown {
    const obj: any = {}
    message.grant !== undefined && (obj.grant = message.grant ? AppFeeGrant.toJSON(message.grant) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<RestQueryAppFeeGrantResponse>): RestQueryAppFeeGrantResponse {
    const message = { ...baseRestQueryAppFeeGrantResponse } as RestQueryAppFeeGrantResponse
    if (object.grant !== undefined && object.grant !== null) {
      message.grant = AppFeeGrant.fromPartial(object.grant)
    } else {
      message.grant = undefined
    }
    return message
  }
}

/** Rest defines the gRPC rest service. */
export interface RestQuery {
  /** query a did */
  QueryDid(request: RestQueryDidRequest): Promise<RestQueryDidResponse>
  /** query a user info */
  QueryUser(request: RestQueryUserRequest): Promise<RestQueryUserResponse>
  /** query user relations */
  QueryUserRelation(request: RestQueryUserRelationRequest): Promise<RestQueryUserRelationResponse>
  /** query app info */
  QueryApp(request: RestQueryAppRequest): Promise<RestQueryAppResponse>
  /** query app info */
  QueryAppFeeGrant(request: RestQueryAppFeeGrantRequest): Promise<RestQueryAppFeeGrantResponse>
  /** query a tx result */
  QueryTx(request: RestQueryTxRequest): Promise<RestTxResponse>
}

export class RestQueryClientImpl implements RestQuery {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  QueryDid(request: RestQueryDidRequest): Promise<RestQueryDidResponse> {
    const data = RestQueryDidRequest.encode(request).finish()
    const promise = this.rpc.request('misesid.misestm.v1beta1.RestQuery', 'QueryDid', data)
    return promise.then((data) => RestQueryDidResponse.decode(new Reader(data)))
  }

  QueryUser(request: RestQueryUserRequest): Promise<RestQueryUserResponse> {
    const data = RestQueryUserRequest.encode(request).finish()
    const promise = this.rpc.request('misesid.misestm.v1beta1.RestQuery', 'QueryUser', data)
    return promise.then((data) => RestQueryUserResponse.decode(new Reader(data)))
  }

  QueryUserRelation(request: RestQueryUserRelationRequest): Promise<RestQueryUserRelationResponse> {
    const data = RestQueryUserRelationRequest.encode(request).finish()
    const promise = this.rpc.request('misesid.misestm.v1beta1.RestQuery', 'QueryUserRelation', data)
    return promise.then((data) => RestQueryUserRelationResponse.decode(new Reader(data)))
  }

  QueryApp(request: RestQueryAppRequest): Promise<RestQueryAppResponse> {
    const data = RestQueryAppRequest.encode(request).finish()
    const promise = this.rpc.request('misesid.misestm.v1beta1.RestQuery', 'QueryApp', data)
    return promise.then((data) => RestQueryAppResponse.decode(new Reader(data)))
  }

  QueryAppFeeGrant(request: RestQueryAppFeeGrantRequest): Promise<RestQueryAppFeeGrantResponse> {
    const data = RestQueryAppFeeGrantRequest.encode(request).finish()
    const promise = this.rpc.request('misesid.misestm.v1beta1.RestQuery', 'QueryAppFeeGrant', data)
    return promise.then((data) => RestQueryAppFeeGrantResponse.decode(new Reader(data)))
  }

  QueryTx(request: RestQueryTxRequest): Promise<RestTxResponse> {
    const data = RestQueryTxRequest.encode(request).finish()
    const promise = this.rpc.request('misesid.misestm.v1beta1.RestQuery', 'QueryTx', data)
    return promise.then((data) => RestTxResponse.decode(new Reader(data)))
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>
}

declare var self: any | undefined
declare var window: any | undefined
var globalThis: any = (() => {
  if (typeof globalThis !== 'undefined') return globalThis
  if (typeof self !== 'undefined') return self
  if (typeof window !== 'undefined') return window
  if (typeof global !== 'undefined') return global
  throw 'Unable to locate global object'
})()

type Builtin = Date | Function | Uint8Array | string | number | undefined
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000
  const nanos = (date.getTime() % 1_000) * 1_000_000
  return { seconds, nanos }
}

function fromTimestamp(t: Timestamp): Date {
  let millis = t.seconds * 1_000
  millis += t.nanos / 1_000_000
  return new Date(millis)
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof Date) {
    return o
  } else if (typeof o === 'string') {
    return new Date(o)
  } else {
    return fromTimestamp(Timestamp.fromJSON(o))
  }
}

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER')
  }
  return long.toNumber()
}

if (util.Long !== Long) {
  util.Long = Long as any
  configure()
}
