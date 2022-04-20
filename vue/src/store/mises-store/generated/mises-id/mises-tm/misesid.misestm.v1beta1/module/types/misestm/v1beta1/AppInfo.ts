/* eslint-disable */
import * as Long from 'long'
import { util, configure, Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'misesid.misestm.v1beta1'

export interface PublicAppInfo {
  name: string
  domains: string[]
  developer: string
  home_url: string
  icon_url: string
}

export interface AppInfo {
  creator: string
  id: number
  appid: string
  pub_info: PublicAppInfo | undefined
  version: number
}

const basePublicAppInfo: object = { name: '', domains: '', developer: '', home_url: '', icon_url: '' }

export const PublicAppInfo = {
  encode(message: PublicAppInfo, writer: Writer = Writer.create()): Writer {
    if (message.name !== '') {
      writer.uint32(10).string(message.name)
    }
    for (const v of message.domains) {
      writer.uint32(18).string(v!)
    }
    if (message.developer !== '') {
      writer.uint32(26).string(message.developer)
    }
    if (message.home_url !== '') {
      writer.uint32(34).string(message.home_url)
    }
    if (message.icon_url !== '') {
      writer.uint32(42).string(message.icon_url)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): PublicAppInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...basePublicAppInfo } as PublicAppInfo
    message.domains = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string()
          break
        case 2:
          message.domains.push(reader.string())
          break
        case 3:
          message.developer = reader.string()
          break
        case 4:
          message.home_url = reader.string()
          break
        case 5:
          message.icon_url = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): PublicAppInfo {
    const message = { ...basePublicAppInfo } as PublicAppInfo
    message.domains = []
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name)
    } else {
      message.name = ''
    }
    if (object.domains !== undefined && object.domains !== null) {
      for (const e of object.domains) {
        message.domains.push(String(e))
      }
    }
    if (object.developer !== undefined && object.developer !== null) {
      message.developer = String(object.developer)
    } else {
      message.developer = ''
    }
    if (object.home_url !== undefined && object.home_url !== null) {
      message.home_url = String(object.home_url)
    } else {
      message.home_url = ''
    }
    if (object.icon_url !== undefined && object.icon_url !== null) {
      message.icon_url = String(object.icon_url)
    } else {
      message.icon_url = ''
    }
    return message
  },

  toJSON(message: PublicAppInfo): unknown {
    const obj: any = {}
    message.name !== undefined && (obj.name = message.name)
    if (message.domains) {
      obj.domains = message.domains.map((e) => e)
    } else {
      obj.domains = []
    }
    message.developer !== undefined && (obj.developer = message.developer)
    message.home_url !== undefined && (obj.home_url = message.home_url)
    message.icon_url !== undefined && (obj.icon_url = message.icon_url)
    return obj
  },

  fromPartial(object: DeepPartial<PublicAppInfo>): PublicAppInfo {
    const message = { ...basePublicAppInfo } as PublicAppInfo
    message.domains = []
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name
    } else {
      message.name = ''
    }
    if (object.domains !== undefined && object.domains !== null) {
      for (const e of object.domains) {
        message.domains.push(e)
      }
    }
    if (object.developer !== undefined && object.developer !== null) {
      message.developer = object.developer
    } else {
      message.developer = ''
    }
    if (object.home_url !== undefined && object.home_url !== null) {
      message.home_url = object.home_url
    } else {
      message.home_url = ''
    }
    if (object.icon_url !== undefined && object.icon_url !== null) {
      message.icon_url = object.icon_url
    } else {
      message.icon_url = ''
    }
    return message
  }
}

const baseAppInfo: object = { creator: '', id: 0, appid: '', version: 0 }

export const AppInfo = {
  encode(message: AppInfo, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id)
    }
    if (message.appid !== '') {
      writer.uint32(26).string(message.appid)
    }
    if (message.pub_info !== undefined) {
      PublicAppInfo.encode(message.pub_info, writer.uint32(34).fork()).ldelim()
    }
    if (message.version !== 0) {
      writer.uint32(40).uint64(message.version)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): AppInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseAppInfo } as AppInfo
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.id = longToNumber(reader.uint64() as Long)
          break
        case 3:
          message.appid = reader.string()
          break
        case 4:
          message.pub_info = PublicAppInfo.decode(reader, reader.uint32())
          break
        case 5:
          message.version = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): AppInfo {
    const message = { ...baseAppInfo } as AppInfo
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id)
    } else {
      message.id = 0
    }
    if (object.appid !== undefined && object.appid !== null) {
      message.appid = String(object.appid)
    } else {
      message.appid = ''
    }
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

  toJSON(message: AppInfo): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.id !== undefined && (obj.id = message.id)
    message.appid !== undefined && (obj.appid = message.appid)
    message.pub_info !== undefined && (obj.pub_info = message.pub_info ? PublicAppInfo.toJSON(message.pub_info) : undefined)
    message.version !== undefined && (obj.version = message.version)
    return obj
  },

  fromPartial(object: DeepPartial<AppInfo>): AppInfo {
    const message = { ...baseAppInfo } as AppInfo
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id
    } else {
      message.id = 0
    }
    if (object.appid !== undefined && object.appid !== null) {
      message.appid = object.appid
    } else {
      message.appid = ''
    }
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
