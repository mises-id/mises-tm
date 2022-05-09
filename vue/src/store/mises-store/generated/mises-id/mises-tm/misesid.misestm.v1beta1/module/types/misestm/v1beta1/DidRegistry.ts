/* eslint-disable */
import * as Long from 'long'
import { util, configure, Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'misesid.misestm.v1beta1'

export interface DidRegistry {
  creator: string
  id: number
  did: string
  pkeyDid: string
  pkeyType: string
  pkeyMultibase: string
  version: number
}

const baseDidRegistry: object = { creator: '', id: 0, did: '', pkeyDid: '', pkeyType: '', pkeyMultibase: '', version: 0 }

export const DidRegistry = {
  encode(message: DidRegistry, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id)
    }
    if (message.did !== '') {
      writer.uint32(26).string(message.did)
    }
    if (message.pkeyDid !== '') {
      writer.uint32(34).string(message.pkeyDid)
    }
    if (message.pkeyType !== '') {
      writer.uint32(42).string(message.pkeyType)
    }
    if (message.pkeyMultibase !== '') {
      writer.uint32(50).string(message.pkeyMultibase)
    }
    if (message.version !== 0) {
      writer.uint32(56).uint64(message.version)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): DidRegistry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseDidRegistry } as DidRegistry
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
          message.did = reader.string()
          break
        case 4:
          message.pkeyDid = reader.string()
          break
        case 5:
          message.pkeyType = reader.string()
          break
        case 6:
          message.pkeyMultibase = reader.string()
          break
        case 7:
          message.version = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): DidRegistry {
    const message = { ...baseDidRegistry } as DidRegistry
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
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did)
    } else {
      message.did = ''
    }
    if (object.pkeyDid !== undefined && object.pkeyDid !== null) {
      message.pkeyDid = String(object.pkeyDid)
    } else {
      message.pkeyDid = ''
    }
    if (object.pkeyType !== undefined && object.pkeyType !== null) {
      message.pkeyType = String(object.pkeyType)
    } else {
      message.pkeyType = ''
    }
    if (object.pkeyMultibase !== undefined && object.pkeyMultibase !== null) {
      message.pkeyMultibase = String(object.pkeyMultibase)
    } else {
      message.pkeyMultibase = ''
    }
    if (object.version !== undefined && object.version !== null) {
      message.version = Number(object.version)
    } else {
      message.version = 0
    }
    return message
  },

  toJSON(message: DidRegistry): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.id !== undefined && (obj.id = message.id)
    message.did !== undefined && (obj.did = message.did)
    message.pkeyDid !== undefined && (obj.pkeyDid = message.pkeyDid)
    message.pkeyType !== undefined && (obj.pkeyType = message.pkeyType)
    message.pkeyMultibase !== undefined && (obj.pkeyMultibase = message.pkeyMultibase)
    message.version !== undefined && (obj.version = message.version)
    return obj
  },

  fromPartial(object: DeepPartial<DidRegistry>): DidRegistry {
    const message = { ...baseDidRegistry } as DidRegistry
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
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did
    } else {
      message.did = ''
    }
    if (object.pkeyDid !== undefined && object.pkeyDid !== null) {
      message.pkeyDid = object.pkeyDid
    } else {
      message.pkeyDid = ''
    }
    if (object.pkeyType !== undefined && object.pkeyType !== null) {
      message.pkeyType = object.pkeyType
    } else {
      message.pkeyType = ''
    }
    if (object.pkeyMultibase !== undefined && object.pkeyMultibase !== null) {
      message.pkeyMultibase = object.pkeyMultibase
    } else {
      message.pkeyMultibase = ''
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
