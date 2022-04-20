/* eslint-disable */
import * as Long from 'long'
import { util, configure, Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'misesid.misestm.v1beta1'

export interface UserRelation {
  creator: string
  id: number
  uidFrom: string
  uidTo: string
  isFollowing: boolean
  isBlocking: boolean
  isReferredBy: boolean
  version: number
}

const baseUserRelation: object = { creator: '', id: 0, uidFrom: '', uidTo: '', isFollowing: false, isBlocking: false, isReferredBy: false, version: 0 }

export const UserRelation = {
  encode(message: UserRelation, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id)
    }
    if (message.uidFrom !== '') {
      writer.uint32(26).string(message.uidFrom)
    }
    if (message.uidTo !== '') {
      writer.uint32(34).string(message.uidTo)
    }
    if (message.isFollowing === true) {
      writer.uint32(40).bool(message.isFollowing)
    }
    if (message.isBlocking === true) {
      writer.uint32(48).bool(message.isBlocking)
    }
    if (message.isReferredBy === true) {
      writer.uint32(56).bool(message.isReferredBy)
    }
    if (message.version !== 0) {
      writer.uint32(64).uint64(message.version)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): UserRelation {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseUserRelation } as UserRelation
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
          message.uidFrom = reader.string()
          break
        case 4:
          message.uidTo = reader.string()
          break
        case 5:
          message.isFollowing = reader.bool()
          break
        case 6:
          message.isBlocking = reader.bool()
          break
        case 7:
          message.isReferredBy = reader.bool()
          break
        case 8:
          message.version = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): UserRelation {
    const message = { ...baseUserRelation } as UserRelation
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
    if (object.uidFrom !== undefined && object.uidFrom !== null) {
      message.uidFrom = String(object.uidFrom)
    } else {
      message.uidFrom = ''
    }
    if (object.uidTo !== undefined && object.uidTo !== null) {
      message.uidTo = String(object.uidTo)
    } else {
      message.uidTo = ''
    }
    if (object.isFollowing !== undefined && object.isFollowing !== null) {
      message.isFollowing = Boolean(object.isFollowing)
    } else {
      message.isFollowing = false
    }
    if (object.isBlocking !== undefined && object.isBlocking !== null) {
      message.isBlocking = Boolean(object.isBlocking)
    } else {
      message.isBlocking = false
    }
    if (object.isReferredBy !== undefined && object.isReferredBy !== null) {
      message.isReferredBy = Boolean(object.isReferredBy)
    } else {
      message.isReferredBy = false
    }
    if (object.version !== undefined && object.version !== null) {
      message.version = Number(object.version)
    } else {
      message.version = 0
    }
    return message
  },

  toJSON(message: UserRelation): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.id !== undefined && (obj.id = message.id)
    message.uidFrom !== undefined && (obj.uidFrom = message.uidFrom)
    message.uidTo !== undefined && (obj.uidTo = message.uidTo)
    message.isFollowing !== undefined && (obj.isFollowing = message.isFollowing)
    message.isBlocking !== undefined && (obj.isBlocking = message.isBlocking)
    message.isReferredBy !== undefined && (obj.isReferredBy = message.isReferredBy)
    message.version !== undefined && (obj.version = message.version)
    return obj
  },

  fromPartial(object: DeepPartial<UserRelation>): UserRelation {
    const message = { ...baseUserRelation } as UserRelation
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
    if (object.uidFrom !== undefined && object.uidFrom !== null) {
      message.uidFrom = object.uidFrom
    } else {
      message.uidFrom = ''
    }
    if (object.uidTo !== undefined && object.uidTo !== null) {
      message.uidTo = object.uidTo
    } else {
      message.uidTo = ''
    }
    if (object.isFollowing !== undefined && object.isFollowing !== null) {
      message.isFollowing = object.isFollowing
    } else {
      message.isFollowing = false
    }
    if (object.isBlocking !== undefined && object.isBlocking !== null) {
      message.isBlocking = object.isBlocking
    } else {
      message.isBlocking = false
    }
    if (object.isReferredBy !== undefined && object.isReferredBy !== null) {
      message.isReferredBy = object.isReferredBy
    } else {
      message.isReferredBy = false
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
