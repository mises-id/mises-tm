/* eslint-disable */
import * as Long from 'long'
import { util, configure, Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'misesid.misestm.v1beta1'

export interface MisesAccount {
  misesID: string
  didRegistryID: number
  didType: number
  infoID: number
}

const baseMisesAccount: object = { misesID: '', didRegistryID: 0, didType: 0, infoID: 0 }

export const MisesAccount = {
  encode(message: MisesAccount, writer: Writer = Writer.create()): Writer {
    if (message.misesID !== '') {
      writer.uint32(10).string(message.misesID)
    }
    if (message.didRegistryID !== 0) {
      writer.uint32(16).uint64(message.didRegistryID)
    }
    if (message.didType !== 0) {
      writer.uint32(24).uint64(message.didType)
    }
    if (message.infoID !== 0) {
      writer.uint32(32).uint64(message.infoID)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MisesAccount {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMisesAccount } as MisesAccount
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.misesID = reader.string()
          break
        case 2:
          message.didRegistryID = longToNumber(reader.uint64() as Long)
          break
        case 3:
          message.didType = longToNumber(reader.uint64() as Long)
          break
        case 4:
          message.infoID = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MisesAccount {
    const message = { ...baseMisesAccount } as MisesAccount
    if (object.misesID !== undefined && object.misesID !== null) {
      message.misesID = String(object.misesID)
    } else {
      message.misesID = ''
    }
    if (object.didRegistryID !== undefined && object.didRegistryID !== null) {
      message.didRegistryID = Number(object.didRegistryID)
    } else {
      message.didRegistryID = 0
    }
    if (object.didType !== undefined && object.didType !== null) {
      message.didType = Number(object.didType)
    } else {
      message.didType = 0
    }
    if (object.infoID !== undefined && object.infoID !== null) {
      message.infoID = Number(object.infoID)
    } else {
      message.infoID = 0
    }
    return message
  },

  toJSON(message: MisesAccount): unknown {
    const obj: any = {}
    message.misesID !== undefined && (obj.misesID = message.misesID)
    message.didRegistryID !== undefined && (obj.didRegistryID = message.didRegistryID)
    message.didType !== undefined && (obj.didType = message.didType)
    message.infoID !== undefined && (obj.infoID = message.infoID)
    return obj
  },

  fromPartial(object: DeepPartial<MisesAccount>): MisesAccount {
    const message = { ...baseMisesAccount } as MisesAccount
    if (object.misesID !== undefined && object.misesID !== null) {
      message.misesID = object.misesID
    } else {
      message.misesID = ''
    }
    if (object.didRegistryID !== undefined && object.didRegistryID !== null) {
      message.didRegistryID = object.didRegistryID
    } else {
      message.didRegistryID = 0
    }
    if (object.didType !== undefined && object.didType !== null) {
      message.didType = object.didType
    } else {
      message.didType = 0
    }
    if (object.infoID !== undefined && object.infoID !== null) {
      message.infoID = object.infoID
    } else {
      message.infoID = 0
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
