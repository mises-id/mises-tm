/* eslint-disable */
import * as Long from 'long'
import { util, configure, Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'misesid.misestm.v1beta1'

export interface PrivateUserInfo {
  enc_data: string
  iv: string
}

export interface PublicUserInfo {
  name: string
  gender: string
  avatar_url: string
  home_page_url: string
  emails: string[]
  telephones: string[]
  intro: string
}

export interface UserInfo {
  creator: string
  id: number
  uid: string
  pub_info: PublicUserInfo | undefined
  pri_info: PrivateUserInfo | undefined
  version: number
}

const basePrivateUserInfo: object = { enc_data: '', iv: '' }

export const PrivateUserInfo = {
  encode(message: PrivateUserInfo, writer: Writer = Writer.create()): Writer {
    if (message.enc_data !== '') {
      writer.uint32(10).string(message.enc_data)
    }
    if (message.iv !== '') {
      writer.uint32(18).string(message.iv)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): PrivateUserInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...basePrivateUserInfo } as PrivateUserInfo
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.enc_data = reader.string()
          break
        case 2:
          message.iv = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): PrivateUserInfo {
    const message = { ...basePrivateUserInfo } as PrivateUserInfo
    if (object.enc_data !== undefined && object.enc_data !== null) {
      message.enc_data = String(object.enc_data)
    } else {
      message.enc_data = ''
    }
    if (object.iv !== undefined && object.iv !== null) {
      message.iv = String(object.iv)
    } else {
      message.iv = ''
    }
    return message
  },

  toJSON(message: PrivateUserInfo): unknown {
    const obj: any = {}
    message.enc_data !== undefined && (obj.enc_data = message.enc_data)
    message.iv !== undefined && (obj.iv = message.iv)
    return obj
  },

  fromPartial(object: DeepPartial<PrivateUserInfo>): PrivateUserInfo {
    const message = { ...basePrivateUserInfo } as PrivateUserInfo
    if (object.enc_data !== undefined && object.enc_data !== null) {
      message.enc_data = object.enc_data
    } else {
      message.enc_data = ''
    }
    if (object.iv !== undefined && object.iv !== null) {
      message.iv = object.iv
    } else {
      message.iv = ''
    }
    return message
  }
}

const basePublicUserInfo: object = { name: '', gender: '', avatar_url: '', home_page_url: '', emails: '', telephones: '', intro: '' }

export const PublicUserInfo = {
  encode(message: PublicUserInfo, writer: Writer = Writer.create()): Writer {
    if (message.name !== '') {
      writer.uint32(10).string(message.name)
    }
    if (message.gender !== '') {
      writer.uint32(18).string(message.gender)
    }
    if (message.avatar_url !== '') {
      writer.uint32(26).string(message.avatar_url)
    }
    if (message.home_page_url !== '') {
      writer.uint32(34).string(message.home_page_url)
    }
    for (const v of message.emails) {
      writer.uint32(42).string(v!)
    }
    for (const v of message.telephones) {
      writer.uint32(50).string(v!)
    }
    if (message.intro !== '') {
      writer.uint32(58).string(message.intro)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): PublicUserInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...basePublicUserInfo } as PublicUserInfo
    message.emails = []
    message.telephones = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string()
          break
        case 2:
          message.gender = reader.string()
          break
        case 3:
          message.avatar_url = reader.string()
          break
        case 4:
          message.home_page_url = reader.string()
          break
        case 5:
          message.emails.push(reader.string())
          break
        case 6:
          message.telephones.push(reader.string())
          break
        case 7:
          message.intro = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): PublicUserInfo {
    const message = { ...basePublicUserInfo } as PublicUserInfo
    message.emails = []
    message.telephones = []
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name)
    } else {
      message.name = ''
    }
    if (object.gender !== undefined && object.gender !== null) {
      message.gender = String(object.gender)
    } else {
      message.gender = ''
    }
    if (object.avatar_url !== undefined && object.avatar_url !== null) {
      message.avatar_url = String(object.avatar_url)
    } else {
      message.avatar_url = ''
    }
    if (object.home_page_url !== undefined && object.home_page_url !== null) {
      message.home_page_url = String(object.home_page_url)
    } else {
      message.home_page_url = ''
    }
    if (object.emails !== undefined && object.emails !== null) {
      for (const e of object.emails) {
        message.emails.push(String(e))
      }
    }
    if (object.telephones !== undefined && object.telephones !== null) {
      for (const e of object.telephones) {
        message.telephones.push(String(e))
      }
    }
    if (object.intro !== undefined && object.intro !== null) {
      message.intro = String(object.intro)
    } else {
      message.intro = ''
    }
    return message
  },

  toJSON(message: PublicUserInfo): unknown {
    const obj: any = {}
    message.name !== undefined && (obj.name = message.name)
    message.gender !== undefined && (obj.gender = message.gender)
    message.avatar_url !== undefined && (obj.avatar_url = message.avatar_url)
    message.home_page_url !== undefined && (obj.home_page_url = message.home_page_url)
    if (message.emails) {
      obj.emails = message.emails.map((e) => e)
    } else {
      obj.emails = []
    }
    if (message.telephones) {
      obj.telephones = message.telephones.map((e) => e)
    } else {
      obj.telephones = []
    }
    message.intro !== undefined && (obj.intro = message.intro)
    return obj
  },

  fromPartial(object: DeepPartial<PublicUserInfo>): PublicUserInfo {
    const message = { ...basePublicUserInfo } as PublicUserInfo
    message.emails = []
    message.telephones = []
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name
    } else {
      message.name = ''
    }
    if (object.gender !== undefined && object.gender !== null) {
      message.gender = object.gender
    } else {
      message.gender = ''
    }
    if (object.avatar_url !== undefined && object.avatar_url !== null) {
      message.avatar_url = object.avatar_url
    } else {
      message.avatar_url = ''
    }
    if (object.home_page_url !== undefined && object.home_page_url !== null) {
      message.home_page_url = object.home_page_url
    } else {
      message.home_page_url = ''
    }
    if (object.emails !== undefined && object.emails !== null) {
      for (const e of object.emails) {
        message.emails.push(e)
      }
    }
    if (object.telephones !== undefined && object.telephones !== null) {
      for (const e of object.telephones) {
        message.telephones.push(e)
      }
    }
    if (object.intro !== undefined && object.intro !== null) {
      message.intro = object.intro
    } else {
      message.intro = ''
    }
    return message
  }
}

const baseUserInfo: object = { creator: '', id: 0, uid: '', version: 0 }

export const UserInfo = {
  encode(message: UserInfo, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id)
    }
    if (message.uid !== '') {
      writer.uint32(26).string(message.uid)
    }
    if (message.pub_info !== undefined) {
      PublicUserInfo.encode(message.pub_info, writer.uint32(34).fork()).ldelim()
    }
    if (message.pri_info !== undefined) {
      PrivateUserInfo.encode(message.pri_info, writer.uint32(42).fork()).ldelim()
    }
    if (message.version !== 0) {
      writer.uint32(48).uint64(message.version)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): UserInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseUserInfo } as UserInfo
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
          message.uid = reader.string()
          break
        case 4:
          message.pub_info = PublicUserInfo.decode(reader, reader.uint32())
          break
        case 5:
          message.pri_info = PrivateUserInfo.decode(reader, reader.uint32())
          break
        case 6:
          message.version = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): UserInfo {
    const message = { ...baseUserInfo } as UserInfo
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
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = String(object.uid)
    } else {
      message.uid = ''
    }
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

  toJSON(message: UserInfo): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.id !== undefined && (obj.id = message.id)
    message.uid !== undefined && (obj.uid = message.uid)
    message.pub_info !== undefined && (obj.pub_info = message.pub_info ? PublicUserInfo.toJSON(message.pub_info) : undefined)
    message.pri_info !== undefined && (obj.pri_info = message.pri_info ? PrivateUserInfo.toJSON(message.pri_info) : undefined)
    message.version !== undefined && (obj.version = message.version)
    return obj
  },

  fromPartial(object: DeepPartial<UserInfo>): UserInfo {
    const message = { ...baseUserInfo } as UserInfo
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
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = object.uid
    } else {
      message.uid = ''
    }
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
