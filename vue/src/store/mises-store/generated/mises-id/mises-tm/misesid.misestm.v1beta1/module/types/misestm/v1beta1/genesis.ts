/* eslint-disable */
import * as Long from 'long'
import { util, configure, Writer, Reader } from 'protobufjs/minimal'
import { MisesAccount } from '../../misestm/v1beta1/MisesAccount'
import { UserInfo } from '../../misestm/v1beta1/UserInfo'
import { UserRelation } from '../../misestm/v1beta1/UserRelation'
import { AppInfo } from '../../misestm/v1beta1/AppInfo'
import { DidRegistry } from '../../misestm/v1beta1/DidRegistry'

export const protobufPackage = 'misesid.misestm.v1beta1'

/** GenesisState defines the misestm module's genesis state. */
export interface GenesisState {
  MisesAccountList: MisesAccount[]
  /** this line is used by starport scaffolding # genesis/proto/state */
  UserInfoList: UserInfo[]
  /** this line is used by starport scaffolding # genesis/proto/stateField */
  UserInfoCount: number
  /** this line is used by starport scaffolding # genesis/proto/stateField */
  UserRelationList: UserRelation[]
  /** this line is used by starport scaffolding # genesis/proto/stateField */
  UserRelationCount: number
  /** this line is used by starport scaffolding # genesis/proto/stateField */
  AppInfoList: AppInfo[]
  /** this line is used by starport scaffolding # genesis/proto/stateField */
  AppInfoCount: number
  /** this line is used by starport scaffolding # genesis/proto/stateField */
  DidRegistryList: DidRegistry[]
  /** this line is used by starport scaffolding # genesis/proto/stateField */
  DidRegistryCount: number
}

const baseGenesisState: object = { UserInfoCount: 0, UserRelationCount: 0, AppInfoCount: 0, DidRegistryCount: 0 }

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    for (const v of message.MisesAccountList) {
      MisesAccount.encode(v!, writer.uint32(74).fork()).ldelim()
    }
    for (const v of message.UserInfoList) {
      UserInfo.encode(v!, writer.uint32(58).fork()).ldelim()
    }
    if (message.UserInfoCount !== 0) {
      writer.uint32(64).uint64(message.UserInfoCount)
    }
    for (const v of message.UserRelationList) {
      UserRelation.encode(v!, writer.uint32(42).fork()).ldelim()
    }
    if (message.UserRelationCount !== 0) {
      writer.uint32(48).uint64(message.UserRelationCount)
    }
    for (const v of message.AppInfoList) {
      AppInfo.encode(v!, writer.uint32(26).fork()).ldelim()
    }
    if (message.AppInfoCount !== 0) {
      writer.uint32(32).uint64(message.AppInfoCount)
    }
    for (const v of message.DidRegistryList) {
      DidRegistry.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.DidRegistryCount !== 0) {
      writer.uint32(16).uint64(message.DidRegistryCount)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseGenesisState } as GenesisState
    message.MisesAccountList = []
    message.UserInfoList = []
    message.UserRelationList = []
    message.AppInfoList = []
    message.DidRegistryList = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 9:
          message.MisesAccountList.push(MisesAccount.decode(reader, reader.uint32()))
          break
        case 7:
          message.UserInfoList.push(UserInfo.decode(reader, reader.uint32()))
          break
        case 8:
          message.UserInfoCount = longToNumber(reader.uint64() as Long)
          break
        case 5:
          message.UserRelationList.push(UserRelation.decode(reader, reader.uint32()))
          break
        case 6:
          message.UserRelationCount = longToNumber(reader.uint64() as Long)
          break
        case 3:
          message.AppInfoList.push(AppInfo.decode(reader, reader.uint32()))
          break
        case 4:
          message.AppInfoCount = longToNumber(reader.uint64() as Long)
          break
        case 1:
          message.DidRegistryList.push(DidRegistry.decode(reader, reader.uint32()))
          break
        case 2:
          message.DidRegistryCount = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.MisesAccountList = []
    message.UserInfoList = []
    message.UserRelationList = []
    message.AppInfoList = []
    message.DidRegistryList = []
    if (object.MisesAccountList !== undefined && object.MisesAccountList !== null) {
      for (const e of object.MisesAccountList) {
        message.MisesAccountList.push(MisesAccount.fromJSON(e))
      }
    }
    if (object.UserInfoList !== undefined && object.UserInfoList !== null) {
      for (const e of object.UserInfoList) {
        message.UserInfoList.push(UserInfo.fromJSON(e))
      }
    }
    if (object.UserInfoCount !== undefined && object.UserInfoCount !== null) {
      message.UserInfoCount = Number(object.UserInfoCount)
    } else {
      message.UserInfoCount = 0
    }
    if (object.UserRelationList !== undefined && object.UserRelationList !== null) {
      for (const e of object.UserRelationList) {
        message.UserRelationList.push(UserRelation.fromJSON(e))
      }
    }
    if (object.UserRelationCount !== undefined && object.UserRelationCount !== null) {
      message.UserRelationCount = Number(object.UserRelationCount)
    } else {
      message.UserRelationCount = 0
    }
    if (object.AppInfoList !== undefined && object.AppInfoList !== null) {
      for (const e of object.AppInfoList) {
        message.AppInfoList.push(AppInfo.fromJSON(e))
      }
    }
    if (object.AppInfoCount !== undefined && object.AppInfoCount !== null) {
      message.AppInfoCount = Number(object.AppInfoCount)
    } else {
      message.AppInfoCount = 0
    }
    if (object.DidRegistryList !== undefined && object.DidRegistryList !== null) {
      for (const e of object.DidRegistryList) {
        message.DidRegistryList.push(DidRegistry.fromJSON(e))
      }
    }
    if (object.DidRegistryCount !== undefined && object.DidRegistryCount !== null) {
      message.DidRegistryCount = Number(object.DidRegistryCount)
    } else {
      message.DidRegistryCount = 0
    }
    return message
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {}
    if (message.MisesAccountList) {
      obj.MisesAccountList = message.MisesAccountList.map((e) => (e ? MisesAccount.toJSON(e) : undefined))
    } else {
      obj.MisesAccountList = []
    }
    if (message.UserInfoList) {
      obj.UserInfoList = message.UserInfoList.map((e) => (e ? UserInfo.toJSON(e) : undefined))
    } else {
      obj.UserInfoList = []
    }
    message.UserInfoCount !== undefined && (obj.UserInfoCount = message.UserInfoCount)
    if (message.UserRelationList) {
      obj.UserRelationList = message.UserRelationList.map((e) => (e ? UserRelation.toJSON(e) : undefined))
    } else {
      obj.UserRelationList = []
    }
    message.UserRelationCount !== undefined && (obj.UserRelationCount = message.UserRelationCount)
    if (message.AppInfoList) {
      obj.AppInfoList = message.AppInfoList.map((e) => (e ? AppInfo.toJSON(e) : undefined))
    } else {
      obj.AppInfoList = []
    }
    message.AppInfoCount !== undefined && (obj.AppInfoCount = message.AppInfoCount)
    if (message.DidRegistryList) {
      obj.DidRegistryList = message.DidRegistryList.map((e) => (e ? DidRegistry.toJSON(e) : undefined))
    } else {
      obj.DidRegistryList = []
    }
    message.DidRegistryCount !== undefined && (obj.DidRegistryCount = message.DidRegistryCount)
    return obj
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.MisesAccountList = []
    message.UserInfoList = []
    message.UserRelationList = []
    message.AppInfoList = []
    message.DidRegistryList = []
    if (object.MisesAccountList !== undefined && object.MisesAccountList !== null) {
      for (const e of object.MisesAccountList) {
        message.MisesAccountList.push(MisesAccount.fromPartial(e))
      }
    }
    if (object.UserInfoList !== undefined && object.UserInfoList !== null) {
      for (const e of object.UserInfoList) {
        message.UserInfoList.push(UserInfo.fromPartial(e))
      }
    }
    if (object.UserInfoCount !== undefined && object.UserInfoCount !== null) {
      message.UserInfoCount = object.UserInfoCount
    } else {
      message.UserInfoCount = 0
    }
    if (object.UserRelationList !== undefined && object.UserRelationList !== null) {
      for (const e of object.UserRelationList) {
        message.UserRelationList.push(UserRelation.fromPartial(e))
      }
    }
    if (object.UserRelationCount !== undefined && object.UserRelationCount !== null) {
      message.UserRelationCount = object.UserRelationCount
    } else {
      message.UserRelationCount = 0
    }
    if (object.AppInfoList !== undefined && object.AppInfoList !== null) {
      for (const e of object.AppInfoList) {
        message.AppInfoList.push(AppInfo.fromPartial(e))
      }
    }
    if (object.AppInfoCount !== undefined && object.AppInfoCount !== null) {
      message.AppInfoCount = object.AppInfoCount
    } else {
      message.AppInfoCount = 0
    }
    if (object.DidRegistryList !== undefined && object.DidRegistryList !== null) {
      for (const e of object.DidRegistryList) {
        message.DidRegistryList.push(DidRegistry.fromPartial(e))
      }
    }
    if (object.DidRegistryCount !== undefined && object.DidRegistryCount !== null) {
      message.DidRegistryCount = object.DidRegistryCount
    } else {
      message.DidRegistryCount = 0
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
