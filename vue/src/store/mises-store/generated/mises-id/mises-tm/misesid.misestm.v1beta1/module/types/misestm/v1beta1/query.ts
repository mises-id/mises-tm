/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'
import { UserInfo } from '../../misestm/v1beta1/UserInfo'
import { PageRequest, PageResponse } from '../../cosmos/base/query/v1beta1/pagination'
import { UserRelation } from '../../misestm/v1beta1/UserRelation'
import { AppInfo } from '../../misestm/v1beta1/AppInfo'
import { DidRegistry } from '../../misestm/v1beta1/DidRegistry'

export const protobufPackage = 'misesid.misestm.v1beta1'

/** this line is used by starport scaffolding # 3 */
export interface QueryGetUserInfoRequest {
  id: number
}

export interface QueryGetUserInfoResponse {
  UserInfo: UserInfo | undefined
}

export interface QueryAllUserInfoRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllUserInfoResponse {
  UserInfo: UserInfo[]
  pagination: PageResponse | undefined
}

export interface QueryGetUserRelationRequest {
  id: number
}

export interface QueryGetUserRelationResponse {
  UserRelation: UserRelation | undefined
}

export interface QueryAllUserRelationRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllUserRelationResponse {
  UserRelation: UserRelation[]
  pagination: PageResponse | undefined
}

export interface QueryGetAppInfoRequest {
  id: number
}

export interface QueryGetAppInfoResponse {
  AppInfo: AppInfo | undefined
}

export interface QueryAllAppInfoRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllAppInfoResponse {
  AppInfo: AppInfo[]
  pagination: PageResponse | undefined
}

export interface QueryGetDidRegistryRequest {
  id: number
}

export interface QueryGetDidRegistryResponse {
  DidRegistry: DidRegistry | undefined
}

export interface QueryAllDidRegistryRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllDidRegistryResponse {
  DidRegistry: DidRegistry[]
  pagination: PageResponse | undefined
}

const baseQueryGetUserInfoRequest: object = { id: 0 }

export const QueryGetUserInfoRequest = {
  encode(message: QueryGetUserInfoRequest, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetUserInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetUserInfoRequest } as QueryGetUserInfoRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetUserInfoRequest {
    const message = { ...baseQueryGetUserInfoRequest } as QueryGetUserInfoRequest
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id)
    } else {
      message.id = 0
    }
    return message
  },

  toJSON(message: QueryGetUserInfoRequest): unknown {
    const obj: any = {}
    message.id !== undefined && (obj.id = message.id)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetUserInfoRequest>): QueryGetUserInfoRequest {
    const message = { ...baseQueryGetUserInfoRequest } as QueryGetUserInfoRequest
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id
    } else {
      message.id = 0
    }
    return message
  }
}

const baseQueryGetUserInfoResponse: object = {}

export const QueryGetUserInfoResponse = {
  encode(message: QueryGetUserInfoResponse, writer: Writer = Writer.create()): Writer {
    if (message.UserInfo !== undefined) {
      UserInfo.encode(message.UserInfo, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetUserInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetUserInfoResponse } as QueryGetUserInfoResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.UserInfo = UserInfo.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetUserInfoResponse {
    const message = { ...baseQueryGetUserInfoResponse } as QueryGetUserInfoResponse
    if (object.UserInfo !== undefined && object.UserInfo !== null) {
      message.UserInfo = UserInfo.fromJSON(object.UserInfo)
    } else {
      message.UserInfo = undefined
    }
    return message
  },

  toJSON(message: QueryGetUserInfoResponse): unknown {
    const obj: any = {}
    message.UserInfo !== undefined && (obj.UserInfo = message.UserInfo ? UserInfo.toJSON(message.UserInfo) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetUserInfoResponse>): QueryGetUserInfoResponse {
    const message = { ...baseQueryGetUserInfoResponse } as QueryGetUserInfoResponse
    if (object.UserInfo !== undefined && object.UserInfo !== null) {
      message.UserInfo = UserInfo.fromPartial(object.UserInfo)
    } else {
      message.UserInfo = undefined
    }
    return message
  }
}

const baseQueryAllUserInfoRequest: object = {}

export const QueryAllUserInfoRequest = {
  encode(message: QueryAllUserInfoRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllUserInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllUserInfoRequest } as QueryAllUserInfoRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllUserInfoRequest {
    const message = { ...baseQueryAllUserInfoRequest } as QueryAllUserInfoRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllUserInfoRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllUserInfoRequest>): QueryAllUserInfoRequest {
    const message = { ...baseQueryAllUserInfoRequest } as QueryAllUserInfoRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllUserInfoResponse: object = {}

export const QueryAllUserInfoResponse = {
  encode(message: QueryAllUserInfoResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.UserInfo) {
      UserInfo.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllUserInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllUserInfoResponse } as QueryAllUserInfoResponse
    message.UserInfo = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.UserInfo.push(UserInfo.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllUserInfoResponse {
    const message = { ...baseQueryAllUserInfoResponse } as QueryAllUserInfoResponse
    message.UserInfo = []
    if (object.UserInfo !== undefined && object.UserInfo !== null) {
      for (const e of object.UserInfo) {
        message.UserInfo.push(UserInfo.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllUserInfoResponse): unknown {
    const obj: any = {}
    if (message.UserInfo) {
      obj.UserInfo = message.UserInfo.map((e) => (e ? UserInfo.toJSON(e) : undefined))
    } else {
      obj.UserInfo = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllUserInfoResponse>): QueryAllUserInfoResponse {
    const message = { ...baseQueryAllUserInfoResponse } as QueryAllUserInfoResponse
    message.UserInfo = []
    if (object.UserInfo !== undefined && object.UserInfo !== null) {
      for (const e of object.UserInfo) {
        message.UserInfo.push(UserInfo.fromPartial(e))
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

const baseQueryGetUserRelationRequest: object = { id: 0 }

export const QueryGetUserRelationRequest = {
  encode(message: QueryGetUserRelationRequest, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetUserRelationRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetUserRelationRequest } as QueryGetUserRelationRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetUserRelationRequest {
    const message = { ...baseQueryGetUserRelationRequest } as QueryGetUserRelationRequest
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id)
    } else {
      message.id = 0
    }
    return message
  },

  toJSON(message: QueryGetUserRelationRequest): unknown {
    const obj: any = {}
    message.id !== undefined && (obj.id = message.id)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetUserRelationRequest>): QueryGetUserRelationRequest {
    const message = { ...baseQueryGetUserRelationRequest } as QueryGetUserRelationRequest
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id
    } else {
      message.id = 0
    }
    return message
  }
}

const baseQueryGetUserRelationResponse: object = {}

export const QueryGetUserRelationResponse = {
  encode(message: QueryGetUserRelationResponse, writer: Writer = Writer.create()): Writer {
    if (message.UserRelation !== undefined) {
      UserRelation.encode(message.UserRelation, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetUserRelationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetUserRelationResponse } as QueryGetUserRelationResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.UserRelation = UserRelation.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetUserRelationResponse {
    const message = { ...baseQueryGetUserRelationResponse } as QueryGetUserRelationResponse
    if (object.UserRelation !== undefined && object.UserRelation !== null) {
      message.UserRelation = UserRelation.fromJSON(object.UserRelation)
    } else {
      message.UserRelation = undefined
    }
    return message
  },

  toJSON(message: QueryGetUserRelationResponse): unknown {
    const obj: any = {}
    message.UserRelation !== undefined && (obj.UserRelation = message.UserRelation ? UserRelation.toJSON(message.UserRelation) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetUserRelationResponse>): QueryGetUserRelationResponse {
    const message = { ...baseQueryGetUserRelationResponse } as QueryGetUserRelationResponse
    if (object.UserRelation !== undefined && object.UserRelation !== null) {
      message.UserRelation = UserRelation.fromPartial(object.UserRelation)
    } else {
      message.UserRelation = undefined
    }
    return message
  }
}

const baseQueryAllUserRelationRequest: object = {}

export const QueryAllUserRelationRequest = {
  encode(message: QueryAllUserRelationRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllUserRelationRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllUserRelationRequest } as QueryAllUserRelationRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllUserRelationRequest {
    const message = { ...baseQueryAllUserRelationRequest } as QueryAllUserRelationRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllUserRelationRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllUserRelationRequest>): QueryAllUserRelationRequest {
    const message = { ...baseQueryAllUserRelationRequest } as QueryAllUserRelationRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllUserRelationResponse: object = {}

export const QueryAllUserRelationResponse = {
  encode(message: QueryAllUserRelationResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.UserRelation) {
      UserRelation.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllUserRelationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllUserRelationResponse } as QueryAllUserRelationResponse
    message.UserRelation = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.UserRelation.push(UserRelation.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllUserRelationResponse {
    const message = { ...baseQueryAllUserRelationResponse } as QueryAllUserRelationResponse
    message.UserRelation = []
    if (object.UserRelation !== undefined && object.UserRelation !== null) {
      for (const e of object.UserRelation) {
        message.UserRelation.push(UserRelation.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllUserRelationResponse): unknown {
    const obj: any = {}
    if (message.UserRelation) {
      obj.UserRelation = message.UserRelation.map((e) => (e ? UserRelation.toJSON(e) : undefined))
    } else {
      obj.UserRelation = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllUserRelationResponse>): QueryAllUserRelationResponse {
    const message = { ...baseQueryAllUserRelationResponse } as QueryAllUserRelationResponse
    message.UserRelation = []
    if (object.UserRelation !== undefined && object.UserRelation !== null) {
      for (const e of object.UserRelation) {
        message.UserRelation.push(UserRelation.fromPartial(e))
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

const baseQueryGetAppInfoRequest: object = { id: 0 }

export const QueryGetAppInfoRequest = {
  encode(message: QueryGetAppInfoRequest, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetAppInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetAppInfoRequest } as QueryGetAppInfoRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetAppInfoRequest {
    const message = { ...baseQueryGetAppInfoRequest } as QueryGetAppInfoRequest
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id)
    } else {
      message.id = 0
    }
    return message
  },

  toJSON(message: QueryGetAppInfoRequest): unknown {
    const obj: any = {}
    message.id !== undefined && (obj.id = message.id)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetAppInfoRequest>): QueryGetAppInfoRequest {
    const message = { ...baseQueryGetAppInfoRequest } as QueryGetAppInfoRequest
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id
    } else {
      message.id = 0
    }
    return message
  }
}

const baseQueryGetAppInfoResponse: object = {}

export const QueryGetAppInfoResponse = {
  encode(message: QueryGetAppInfoResponse, writer: Writer = Writer.create()): Writer {
    if (message.AppInfo !== undefined) {
      AppInfo.encode(message.AppInfo, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetAppInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetAppInfoResponse } as QueryGetAppInfoResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.AppInfo = AppInfo.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetAppInfoResponse {
    const message = { ...baseQueryGetAppInfoResponse } as QueryGetAppInfoResponse
    if (object.AppInfo !== undefined && object.AppInfo !== null) {
      message.AppInfo = AppInfo.fromJSON(object.AppInfo)
    } else {
      message.AppInfo = undefined
    }
    return message
  },

  toJSON(message: QueryGetAppInfoResponse): unknown {
    const obj: any = {}
    message.AppInfo !== undefined && (obj.AppInfo = message.AppInfo ? AppInfo.toJSON(message.AppInfo) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetAppInfoResponse>): QueryGetAppInfoResponse {
    const message = { ...baseQueryGetAppInfoResponse } as QueryGetAppInfoResponse
    if (object.AppInfo !== undefined && object.AppInfo !== null) {
      message.AppInfo = AppInfo.fromPartial(object.AppInfo)
    } else {
      message.AppInfo = undefined
    }
    return message
  }
}

const baseQueryAllAppInfoRequest: object = {}

export const QueryAllAppInfoRequest = {
  encode(message: QueryAllAppInfoRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllAppInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllAppInfoRequest } as QueryAllAppInfoRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllAppInfoRequest {
    const message = { ...baseQueryAllAppInfoRequest } as QueryAllAppInfoRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllAppInfoRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllAppInfoRequest>): QueryAllAppInfoRequest {
    const message = { ...baseQueryAllAppInfoRequest } as QueryAllAppInfoRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllAppInfoResponse: object = {}

export const QueryAllAppInfoResponse = {
  encode(message: QueryAllAppInfoResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.AppInfo) {
      AppInfo.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllAppInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllAppInfoResponse } as QueryAllAppInfoResponse
    message.AppInfo = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.AppInfo.push(AppInfo.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllAppInfoResponse {
    const message = { ...baseQueryAllAppInfoResponse } as QueryAllAppInfoResponse
    message.AppInfo = []
    if (object.AppInfo !== undefined && object.AppInfo !== null) {
      for (const e of object.AppInfo) {
        message.AppInfo.push(AppInfo.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllAppInfoResponse): unknown {
    const obj: any = {}
    if (message.AppInfo) {
      obj.AppInfo = message.AppInfo.map((e) => (e ? AppInfo.toJSON(e) : undefined))
    } else {
      obj.AppInfo = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllAppInfoResponse>): QueryAllAppInfoResponse {
    const message = { ...baseQueryAllAppInfoResponse } as QueryAllAppInfoResponse
    message.AppInfo = []
    if (object.AppInfo !== undefined && object.AppInfo !== null) {
      for (const e of object.AppInfo) {
        message.AppInfo.push(AppInfo.fromPartial(e))
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

const baseQueryGetDidRegistryRequest: object = { id: 0 }

export const QueryGetDidRegistryRequest = {
  encode(message: QueryGetDidRegistryRequest, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetDidRegistryRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetDidRegistryRequest } as QueryGetDidRegistryRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetDidRegistryRequest {
    const message = { ...baseQueryGetDidRegistryRequest } as QueryGetDidRegistryRequest
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id)
    } else {
      message.id = 0
    }
    return message
  },

  toJSON(message: QueryGetDidRegistryRequest): unknown {
    const obj: any = {}
    message.id !== undefined && (obj.id = message.id)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetDidRegistryRequest>): QueryGetDidRegistryRequest {
    const message = { ...baseQueryGetDidRegistryRequest } as QueryGetDidRegistryRequest
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id
    } else {
      message.id = 0
    }
    return message
  }
}

const baseQueryGetDidRegistryResponse: object = {}

export const QueryGetDidRegistryResponse = {
  encode(message: QueryGetDidRegistryResponse, writer: Writer = Writer.create()): Writer {
    if (message.DidRegistry !== undefined) {
      DidRegistry.encode(message.DidRegistry, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetDidRegistryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetDidRegistryResponse } as QueryGetDidRegistryResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.DidRegistry = DidRegistry.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetDidRegistryResponse {
    const message = { ...baseQueryGetDidRegistryResponse } as QueryGetDidRegistryResponse
    if (object.DidRegistry !== undefined && object.DidRegistry !== null) {
      message.DidRegistry = DidRegistry.fromJSON(object.DidRegistry)
    } else {
      message.DidRegistry = undefined
    }
    return message
  },

  toJSON(message: QueryGetDidRegistryResponse): unknown {
    const obj: any = {}
    message.DidRegistry !== undefined && (obj.DidRegistry = message.DidRegistry ? DidRegistry.toJSON(message.DidRegistry) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetDidRegistryResponse>): QueryGetDidRegistryResponse {
    const message = { ...baseQueryGetDidRegistryResponse } as QueryGetDidRegistryResponse
    if (object.DidRegistry !== undefined && object.DidRegistry !== null) {
      message.DidRegistry = DidRegistry.fromPartial(object.DidRegistry)
    } else {
      message.DidRegistry = undefined
    }
    return message
  }
}

const baseQueryAllDidRegistryRequest: object = {}

export const QueryAllDidRegistryRequest = {
  encode(message: QueryAllDidRegistryRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllDidRegistryRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllDidRegistryRequest } as QueryAllDidRegistryRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllDidRegistryRequest {
    const message = { ...baseQueryAllDidRegistryRequest } as QueryAllDidRegistryRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllDidRegistryRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllDidRegistryRequest>): QueryAllDidRegistryRequest {
    const message = { ...baseQueryAllDidRegistryRequest } as QueryAllDidRegistryRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllDidRegistryResponse: object = {}

export const QueryAllDidRegistryResponse = {
  encode(message: QueryAllDidRegistryResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.DidRegistry) {
      DidRegistry.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllDidRegistryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllDidRegistryResponse } as QueryAllDidRegistryResponse
    message.DidRegistry = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.DidRegistry.push(DidRegistry.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllDidRegistryResponse {
    const message = { ...baseQueryAllDidRegistryResponse } as QueryAllDidRegistryResponse
    message.DidRegistry = []
    if (object.DidRegistry !== undefined && object.DidRegistry !== null) {
      for (const e of object.DidRegistry) {
        message.DidRegistry.push(DidRegistry.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllDidRegistryResponse): unknown {
    const obj: any = {}
    if (message.DidRegistry) {
      obj.DidRegistry = message.DidRegistry.map((e) => (e ? DidRegistry.toJSON(e) : undefined))
    } else {
      obj.DidRegistry = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllDidRegistryResponse>): QueryAllDidRegistryResponse {
    const message = { ...baseQueryAllDidRegistryResponse } as QueryAllDidRegistryResponse
    message.DidRegistry = []
    if (object.DidRegistry !== undefined && object.DidRegistry !== null) {
      for (const e of object.DidRegistry) {
        message.DidRegistry.push(DidRegistry.fromPartial(e))
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

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a UserInfo by id. */
  UserInfo(request: QueryGetUserInfoRequest): Promise<QueryGetUserInfoResponse>
  /** Queries a list of UserInfo items. */
  UserInfoAll(request: QueryAllUserInfoRequest): Promise<QueryAllUserInfoResponse>
  /** Queries a UserRelation by id. */
  UserRelation(request: QueryGetUserRelationRequest): Promise<QueryGetUserRelationResponse>
  /** Queries a list of UserRelation items. */
  UserRelationAll(request: QueryAllUserRelationRequest): Promise<QueryAllUserRelationResponse>
  /** Queries a AppInfo by id. */
  AppInfo(request: QueryGetAppInfoRequest): Promise<QueryGetAppInfoResponse>
  /** Queries a list of AppInfo items. */
  AppInfoAll(request: QueryAllAppInfoRequest): Promise<QueryAllAppInfoResponse>
  /** Queries a DidRegistry by id. */
  DidRegistry(request: QueryGetDidRegistryRequest): Promise<QueryGetDidRegistryResponse>
  /** Queries a list of DidRegistry items. */
  DidRegistryAll(request: QueryAllDidRegistryRequest): Promise<QueryAllDidRegistryResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  UserInfo(request: QueryGetUserInfoRequest): Promise<QueryGetUserInfoResponse> {
    const data = QueryGetUserInfoRequest.encode(request).finish()
    const promise = this.rpc.request('misesid.misestm.v1beta1.Query', 'UserInfo', data)
    return promise.then((data) => QueryGetUserInfoResponse.decode(new Reader(data)))
  }

  UserInfoAll(request: QueryAllUserInfoRequest): Promise<QueryAllUserInfoResponse> {
    const data = QueryAllUserInfoRequest.encode(request).finish()
    const promise = this.rpc.request('misesid.misestm.v1beta1.Query', 'UserInfoAll', data)
    return promise.then((data) => QueryAllUserInfoResponse.decode(new Reader(data)))
  }

  UserRelation(request: QueryGetUserRelationRequest): Promise<QueryGetUserRelationResponse> {
    const data = QueryGetUserRelationRequest.encode(request).finish()
    const promise = this.rpc.request('misesid.misestm.v1beta1.Query', 'UserRelation', data)
    return promise.then((data) => QueryGetUserRelationResponse.decode(new Reader(data)))
  }

  UserRelationAll(request: QueryAllUserRelationRequest): Promise<QueryAllUserRelationResponse> {
    const data = QueryAllUserRelationRequest.encode(request).finish()
    const promise = this.rpc.request('misesid.misestm.v1beta1.Query', 'UserRelationAll', data)
    return promise.then((data) => QueryAllUserRelationResponse.decode(new Reader(data)))
  }

  AppInfo(request: QueryGetAppInfoRequest): Promise<QueryGetAppInfoResponse> {
    const data = QueryGetAppInfoRequest.encode(request).finish()
    const promise = this.rpc.request('misesid.misestm.v1beta1.Query', 'AppInfo', data)
    return promise.then((data) => QueryGetAppInfoResponse.decode(new Reader(data)))
  }

  AppInfoAll(request: QueryAllAppInfoRequest): Promise<QueryAllAppInfoResponse> {
    const data = QueryAllAppInfoRequest.encode(request).finish()
    const promise = this.rpc.request('misesid.misestm.v1beta1.Query', 'AppInfoAll', data)
    return promise.then((data) => QueryAllAppInfoResponse.decode(new Reader(data)))
  }

  DidRegistry(request: QueryGetDidRegistryRequest): Promise<QueryGetDidRegistryResponse> {
    const data = QueryGetDidRegistryRequest.encode(request).finish()
    const promise = this.rpc.request('misesid.misestm.v1beta1.Query', 'DidRegistry', data)
    return promise.then((data) => QueryGetDidRegistryResponse.decode(new Reader(data)))
  }

  DidRegistryAll(request: QueryAllDidRegistryRequest): Promise<QueryAllDidRegistryResponse> {
    const data = QueryAllDidRegistryRequest.encode(request).finish()
    const promise = this.rpc.request('misesid.misestm.v1beta1.Query', 'DidRegistryAll', data)
    return promise.then((data) => QueryAllDidRegistryResponse.decode(new Reader(data)))
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
