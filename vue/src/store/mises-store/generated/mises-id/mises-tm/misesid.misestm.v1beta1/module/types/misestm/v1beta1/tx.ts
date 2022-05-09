/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'
import { PublicUserInfo, PrivateUserInfo } from '../../misestm/v1beta1/UserInfo'
import { Metadata } from '../../cosmos/bank/v1beta1/bank'
import { Any } from '../../google/protobuf/any'

export const protobufPackage = 'misesid.misestm.v1beta1'

export interface MsgUpdateUserInfo {
  creator: string
  uid: string
  pub_info: PublicUserInfo | undefined
  pri_info: PrivateUserInfo | undefined
  version: number
}

export interface MsgUpdateUserInfoResponse {}

export interface MsgUpdateUserRelation {
  creator: string
  uidFrom: string
  uidTo: string
  isFollowing: boolean
  isBlocking: boolean
  isReferredBy: boolean
  version: number
}

export interface MsgUpdateUserRelationResponse {}

export interface MsgUpdateAppInfo {
  creator: string
  appid: string
  name: string
  domains: string[]
  developer: string
  home_url: string
  icon_url: string
  version: number
}

export interface MsgUpdateAppInfoResponse {}

export interface MsgCreateDidRegistry {
  creator: string
  did: string
  pkeyDid: string
  pkeyType: string
  pkeyMultibase: string
  version: number
}

export interface MsgCreateDidRegistryResponse {}

/** MsgNewDenom defines an SDK message for creating a new denom. */
export interface MsgNewDenom {
  id: string
  amount: string
  denom_meta: Metadata | undefined
  sender: string
  recipient: string
}

/** MsgNewDenomResponse defines the MsgNewDenom response type. */
export interface MsgNewDenomResponse {}

/** MsgNewNFTClass defines an SDK message for creating a new NFTClass. */
export interface MsgNewNFTClass {
  id: string
  name: string
  uri: string
  schema: string
  symbol: string
  data: Any | undefined
  sender: string
}

/** MsgNewNFTClassResponse defines the MsgNewNFTClass response type. */
export interface MsgNewNFTClassResponse {}

/** MsgUpdateNFTClass defines an SDK message for editing a nft class. */
export interface MsgUpdateNFTClass {
  id: string
  class_id: string
  name: string
  uri: string
  data: Any | undefined
  sender: string
}

/** MsgUpdateNFTClassResponse defines the MsgUpdateNFTClass response type. */
export interface MsgUpdateNFTClassResponse {}

/** MsgMintNFT defines an SDK message for creating a new NFT. */
export interface MsgMintNFT {
  id: string
  class_id: string
  name: string
  uri: string
  data: Any | undefined
  sender: string
  recipient: string
}

/** MsgMintNFTResponse defines the Msg/MintNFT response type. */
export interface MsgMintNFTResponse {}

/** MsgUpdateNFT defines an SDK message for update a NFT. */
export interface MsgUpdateNFT {
  id: string
  class_id: string
  uri: string
  data: Any | undefined
  sender: string
}

/** MsgUpdateNFTResponse defines the MsgUpdateNFT response type. */
export interface MsgUpdateNFTResponse {}

/** MsgBurnNFT defines an SDK message for burning a NFT. */
export interface MsgBurnNFT {
  id: string
  class_id: string
  sender: string
}

/** MsgBurnNFTResponse defines the Msg/BurnNFT response type. */
export interface MsgBurnNFTResponse {}

const baseMsgUpdateUserInfo: object = { creator: '', uid: '', version: 0 }

export const MsgUpdateUserInfo = {
  encode(message: MsgUpdateUserInfo, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.uid !== '') {
      writer.uint32(18).string(message.uid)
    }
    if (message.pub_info !== undefined) {
      PublicUserInfo.encode(message.pub_info, writer.uint32(26).fork()).ldelim()
    }
    if (message.pri_info !== undefined) {
      PrivateUserInfo.encode(message.pri_info, writer.uint32(34).fork()).ldelim()
    }
    if (message.version !== 0) {
      writer.uint32(40).uint64(message.version)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateUserInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateUserInfo } as MsgUpdateUserInfo
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.uid = reader.string()
          break
        case 3:
          message.pub_info = PublicUserInfo.decode(reader, reader.uint32())
          break
        case 4:
          message.pri_info = PrivateUserInfo.decode(reader, reader.uint32())
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

  fromJSON(object: any): MsgUpdateUserInfo {
    const message = { ...baseMsgUpdateUserInfo } as MsgUpdateUserInfo
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
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

  toJSON(message: MsgUpdateUserInfo): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.uid !== undefined && (obj.uid = message.uid)
    message.pub_info !== undefined && (obj.pub_info = message.pub_info ? PublicUserInfo.toJSON(message.pub_info) : undefined)
    message.pri_info !== undefined && (obj.pri_info = message.pri_info ? PrivateUserInfo.toJSON(message.pri_info) : undefined)
    message.version !== undefined && (obj.version = message.version)
    return obj
  },

  fromPartial(object: DeepPartial<MsgUpdateUserInfo>): MsgUpdateUserInfo {
    const message = { ...baseMsgUpdateUserInfo } as MsgUpdateUserInfo
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
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

const baseMsgUpdateUserInfoResponse: object = {}

export const MsgUpdateUserInfoResponse = {
  encode(_: MsgUpdateUserInfoResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateUserInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateUserInfoResponse } as MsgUpdateUserInfoResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgUpdateUserInfoResponse {
    const message = { ...baseMsgUpdateUserInfoResponse } as MsgUpdateUserInfoResponse
    return message
  },

  toJSON(_: MsgUpdateUserInfoResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgUpdateUserInfoResponse>): MsgUpdateUserInfoResponse {
    const message = { ...baseMsgUpdateUserInfoResponse } as MsgUpdateUserInfoResponse
    return message
  }
}

const baseMsgUpdateUserRelation: object = { creator: '', uidFrom: '', uidTo: '', isFollowing: false, isBlocking: false, isReferredBy: false, version: 0 }

export const MsgUpdateUserRelation = {
  encode(message: MsgUpdateUserRelation, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.uidFrom !== '') {
      writer.uint32(18).string(message.uidFrom)
    }
    if (message.uidTo !== '') {
      writer.uint32(26).string(message.uidTo)
    }
    if (message.isFollowing === true) {
      writer.uint32(32).bool(message.isFollowing)
    }
    if (message.isBlocking === true) {
      writer.uint32(40).bool(message.isBlocking)
    }
    if (message.isReferredBy === true) {
      writer.uint32(48).bool(message.isReferredBy)
    }
    if (message.version !== 0) {
      writer.uint32(56).uint64(message.version)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateUserRelation {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateUserRelation } as MsgUpdateUserRelation
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.uidFrom = reader.string()
          break
        case 3:
          message.uidTo = reader.string()
          break
        case 4:
          message.isFollowing = reader.bool()
          break
        case 5:
          message.isBlocking = reader.bool()
          break
        case 6:
          message.isReferredBy = reader.bool()
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

  fromJSON(object: any): MsgUpdateUserRelation {
    const message = { ...baseMsgUpdateUserRelation } as MsgUpdateUserRelation
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
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

  toJSON(message: MsgUpdateUserRelation): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.uidFrom !== undefined && (obj.uidFrom = message.uidFrom)
    message.uidTo !== undefined && (obj.uidTo = message.uidTo)
    message.isFollowing !== undefined && (obj.isFollowing = message.isFollowing)
    message.isBlocking !== undefined && (obj.isBlocking = message.isBlocking)
    message.isReferredBy !== undefined && (obj.isReferredBy = message.isReferredBy)
    message.version !== undefined && (obj.version = message.version)
    return obj
  },

  fromPartial(object: DeepPartial<MsgUpdateUserRelation>): MsgUpdateUserRelation {
    const message = { ...baseMsgUpdateUserRelation } as MsgUpdateUserRelation
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
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

const baseMsgUpdateUserRelationResponse: object = {}

export const MsgUpdateUserRelationResponse = {
  encode(_: MsgUpdateUserRelationResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateUserRelationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateUserRelationResponse } as MsgUpdateUserRelationResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgUpdateUserRelationResponse {
    const message = { ...baseMsgUpdateUserRelationResponse } as MsgUpdateUserRelationResponse
    return message
  },

  toJSON(_: MsgUpdateUserRelationResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgUpdateUserRelationResponse>): MsgUpdateUserRelationResponse {
    const message = { ...baseMsgUpdateUserRelationResponse } as MsgUpdateUserRelationResponse
    return message
  }
}

const baseMsgUpdateAppInfo: object = { creator: '', appid: '', name: '', domains: '', developer: '', home_url: '', icon_url: '', version: 0 }

export const MsgUpdateAppInfo = {
  encode(message: MsgUpdateAppInfo, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.appid !== '') {
      writer.uint32(18).string(message.appid)
    }
    if (message.name !== '') {
      writer.uint32(26).string(message.name)
    }
    for (const v of message.domains) {
      writer.uint32(34).string(v!)
    }
    if (message.developer !== '') {
      writer.uint32(42).string(message.developer)
    }
    if (message.home_url !== '') {
      writer.uint32(50).string(message.home_url)
    }
    if (message.icon_url !== '') {
      writer.uint32(58).string(message.icon_url)
    }
    if (message.version !== 0) {
      writer.uint32(64).uint64(message.version)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateAppInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateAppInfo } as MsgUpdateAppInfo
    message.domains = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.appid = reader.string()
          break
        case 3:
          message.name = reader.string()
          break
        case 4:
          message.domains.push(reader.string())
          break
        case 5:
          message.developer = reader.string()
          break
        case 6:
          message.home_url = reader.string()
          break
        case 7:
          message.icon_url = reader.string()
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

  fromJSON(object: any): MsgUpdateAppInfo {
    const message = { ...baseMsgUpdateAppInfo } as MsgUpdateAppInfo
    message.domains = []
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.appid !== undefined && object.appid !== null) {
      message.appid = String(object.appid)
    } else {
      message.appid = ''
    }
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
    if (object.version !== undefined && object.version !== null) {
      message.version = Number(object.version)
    } else {
      message.version = 0
    }
    return message
  },

  toJSON(message: MsgUpdateAppInfo): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.appid !== undefined && (obj.appid = message.appid)
    message.name !== undefined && (obj.name = message.name)
    if (message.domains) {
      obj.domains = message.domains.map((e) => e)
    } else {
      obj.domains = []
    }
    message.developer !== undefined && (obj.developer = message.developer)
    message.home_url !== undefined && (obj.home_url = message.home_url)
    message.icon_url !== undefined && (obj.icon_url = message.icon_url)
    message.version !== undefined && (obj.version = message.version)
    return obj
  },

  fromPartial(object: DeepPartial<MsgUpdateAppInfo>): MsgUpdateAppInfo {
    const message = { ...baseMsgUpdateAppInfo } as MsgUpdateAppInfo
    message.domains = []
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.appid !== undefined && object.appid !== null) {
      message.appid = object.appid
    } else {
      message.appid = ''
    }
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
    if (object.version !== undefined && object.version !== null) {
      message.version = object.version
    } else {
      message.version = 0
    }
    return message
  }
}

const baseMsgUpdateAppInfoResponse: object = {}

export const MsgUpdateAppInfoResponse = {
  encode(_: MsgUpdateAppInfoResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateAppInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateAppInfoResponse } as MsgUpdateAppInfoResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgUpdateAppInfoResponse {
    const message = { ...baseMsgUpdateAppInfoResponse } as MsgUpdateAppInfoResponse
    return message
  },

  toJSON(_: MsgUpdateAppInfoResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgUpdateAppInfoResponse>): MsgUpdateAppInfoResponse {
    const message = { ...baseMsgUpdateAppInfoResponse } as MsgUpdateAppInfoResponse
    return message
  }
}

const baseMsgCreateDidRegistry: object = { creator: '', did: '', pkeyDid: '', pkeyType: '', pkeyMultibase: '', version: 0 }

export const MsgCreateDidRegistry = {
  encode(message: MsgCreateDidRegistry, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.did !== '') {
      writer.uint32(18).string(message.did)
    }
    if (message.pkeyDid !== '') {
      writer.uint32(26).string(message.pkeyDid)
    }
    if (message.pkeyType !== '') {
      writer.uint32(34).string(message.pkeyType)
    }
    if (message.pkeyMultibase !== '') {
      writer.uint32(42).string(message.pkeyMultibase)
    }
    if (message.version !== 0) {
      writer.uint32(48).uint64(message.version)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateDidRegistry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateDidRegistry } as MsgCreateDidRegistry
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.did = reader.string()
          break
        case 3:
          message.pkeyDid = reader.string()
          break
        case 4:
          message.pkeyType = reader.string()
          break
        case 5:
          message.pkeyMultibase = reader.string()
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

  fromJSON(object: any): MsgCreateDidRegistry {
    const message = { ...baseMsgCreateDidRegistry } as MsgCreateDidRegistry
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
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

  toJSON(message: MsgCreateDidRegistry): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.did !== undefined && (obj.did = message.did)
    message.pkeyDid !== undefined && (obj.pkeyDid = message.pkeyDid)
    message.pkeyType !== undefined && (obj.pkeyType = message.pkeyType)
    message.pkeyMultibase !== undefined && (obj.pkeyMultibase = message.pkeyMultibase)
    message.version !== undefined && (obj.version = message.version)
    return obj
  },

  fromPartial(object: DeepPartial<MsgCreateDidRegistry>): MsgCreateDidRegistry {
    const message = { ...baseMsgCreateDidRegistry } as MsgCreateDidRegistry
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
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

const baseMsgCreateDidRegistryResponse: object = {}

export const MsgCreateDidRegistryResponse = {
  encode(_: MsgCreateDidRegistryResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateDidRegistryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateDidRegistryResponse } as MsgCreateDidRegistryResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgCreateDidRegistryResponse {
    const message = { ...baseMsgCreateDidRegistryResponse } as MsgCreateDidRegistryResponse
    return message
  },

  toJSON(_: MsgCreateDidRegistryResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgCreateDidRegistryResponse>): MsgCreateDidRegistryResponse {
    const message = { ...baseMsgCreateDidRegistryResponse } as MsgCreateDidRegistryResponse
    return message
  }
}

const baseMsgNewDenom: object = { id: '', amount: '', sender: '', recipient: '' }

export const MsgNewDenom = {
  encode(message: MsgNewDenom, writer: Writer = Writer.create()): Writer {
    if (message.id !== '') {
      writer.uint32(10).string(message.id)
    }
    if (message.amount !== '') {
      writer.uint32(18).string(message.amount)
    }
    if (message.denom_meta !== undefined) {
      Metadata.encode(message.denom_meta, writer.uint32(26).fork()).ldelim()
    }
    if (message.sender !== '') {
      writer.uint32(34).string(message.sender)
    }
    if (message.recipient !== '') {
      writer.uint32(42).string(message.recipient)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgNewDenom {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgNewDenom } as MsgNewDenom
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string()
          break
        case 2:
          message.amount = reader.string()
          break
        case 3:
          message.denom_meta = Metadata.decode(reader, reader.uint32())
          break
        case 4:
          message.sender = reader.string()
          break
        case 5:
          message.recipient = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgNewDenom {
    const message = { ...baseMsgNewDenom } as MsgNewDenom
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id)
    } else {
      message.id = ''
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount)
    } else {
      message.amount = ''
    }
    if (object.denom_meta !== undefined && object.denom_meta !== null) {
      message.denom_meta = Metadata.fromJSON(object.denom_meta)
    } else {
      message.denom_meta = undefined
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender)
    } else {
      message.sender = ''
    }
    if (object.recipient !== undefined && object.recipient !== null) {
      message.recipient = String(object.recipient)
    } else {
      message.recipient = ''
    }
    return message
  },

  toJSON(message: MsgNewDenom): unknown {
    const obj: any = {}
    message.id !== undefined && (obj.id = message.id)
    message.amount !== undefined && (obj.amount = message.amount)
    message.denom_meta !== undefined && (obj.denom_meta = message.denom_meta ? Metadata.toJSON(message.denom_meta) : undefined)
    message.sender !== undefined && (obj.sender = message.sender)
    message.recipient !== undefined && (obj.recipient = message.recipient)
    return obj
  },

  fromPartial(object: DeepPartial<MsgNewDenom>): MsgNewDenom {
    const message = { ...baseMsgNewDenom } as MsgNewDenom
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id
    } else {
      message.id = ''
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount
    } else {
      message.amount = ''
    }
    if (object.denom_meta !== undefined && object.denom_meta !== null) {
      message.denom_meta = Metadata.fromPartial(object.denom_meta)
    } else {
      message.denom_meta = undefined
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender
    } else {
      message.sender = ''
    }
    if (object.recipient !== undefined && object.recipient !== null) {
      message.recipient = object.recipient
    } else {
      message.recipient = ''
    }
    return message
  }
}

const baseMsgNewDenomResponse: object = {}

export const MsgNewDenomResponse = {
  encode(_: MsgNewDenomResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgNewDenomResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgNewDenomResponse } as MsgNewDenomResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgNewDenomResponse {
    const message = { ...baseMsgNewDenomResponse } as MsgNewDenomResponse
    return message
  },

  toJSON(_: MsgNewDenomResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgNewDenomResponse>): MsgNewDenomResponse {
    const message = { ...baseMsgNewDenomResponse } as MsgNewDenomResponse
    return message
  }
}

const baseMsgNewNFTClass: object = { id: '', name: '', uri: '', schema: '', symbol: '', sender: '' }

export const MsgNewNFTClass = {
  encode(message: MsgNewNFTClass, writer: Writer = Writer.create()): Writer {
    if (message.id !== '') {
      writer.uint32(10).string(message.id)
    }
    if (message.name !== '') {
      writer.uint32(18).string(message.name)
    }
    if (message.uri !== '') {
      writer.uint32(26).string(message.uri)
    }
    if (message.schema !== '') {
      writer.uint32(34).string(message.schema)
    }
    if (message.symbol !== '') {
      writer.uint32(42).string(message.symbol)
    }
    if (message.data !== undefined) {
      Any.encode(message.data, writer.uint32(50).fork()).ldelim()
    }
    if (message.sender !== '') {
      writer.uint32(58).string(message.sender)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgNewNFTClass {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgNewNFTClass } as MsgNewNFTClass
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string()
          break
        case 2:
          message.name = reader.string()
          break
        case 3:
          message.uri = reader.string()
          break
        case 4:
          message.schema = reader.string()
          break
        case 5:
          message.symbol = reader.string()
          break
        case 6:
          message.data = Any.decode(reader, reader.uint32())
          break
        case 7:
          message.sender = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgNewNFTClass {
    const message = { ...baseMsgNewNFTClass } as MsgNewNFTClass
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id)
    } else {
      message.id = ''
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name)
    } else {
      message.name = ''
    }
    if (object.uri !== undefined && object.uri !== null) {
      message.uri = String(object.uri)
    } else {
      message.uri = ''
    }
    if (object.schema !== undefined && object.schema !== null) {
      message.schema = String(object.schema)
    } else {
      message.schema = ''
    }
    if (object.symbol !== undefined && object.symbol !== null) {
      message.symbol = String(object.symbol)
    } else {
      message.symbol = ''
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = Any.fromJSON(object.data)
    } else {
      message.data = undefined
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender)
    } else {
      message.sender = ''
    }
    return message
  },

  toJSON(message: MsgNewNFTClass): unknown {
    const obj: any = {}
    message.id !== undefined && (obj.id = message.id)
    message.name !== undefined && (obj.name = message.name)
    message.uri !== undefined && (obj.uri = message.uri)
    message.schema !== undefined && (obj.schema = message.schema)
    message.symbol !== undefined && (obj.symbol = message.symbol)
    message.data !== undefined && (obj.data = message.data ? Any.toJSON(message.data) : undefined)
    message.sender !== undefined && (obj.sender = message.sender)
    return obj
  },

  fromPartial(object: DeepPartial<MsgNewNFTClass>): MsgNewNFTClass {
    const message = { ...baseMsgNewNFTClass } as MsgNewNFTClass
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id
    } else {
      message.id = ''
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name
    } else {
      message.name = ''
    }
    if (object.uri !== undefined && object.uri !== null) {
      message.uri = object.uri
    } else {
      message.uri = ''
    }
    if (object.schema !== undefined && object.schema !== null) {
      message.schema = object.schema
    } else {
      message.schema = ''
    }
    if (object.symbol !== undefined && object.symbol !== null) {
      message.symbol = object.symbol
    } else {
      message.symbol = ''
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = Any.fromPartial(object.data)
    } else {
      message.data = undefined
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender
    } else {
      message.sender = ''
    }
    return message
  }
}

const baseMsgNewNFTClassResponse: object = {}

export const MsgNewNFTClassResponse = {
  encode(_: MsgNewNFTClassResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgNewNFTClassResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgNewNFTClassResponse } as MsgNewNFTClassResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgNewNFTClassResponse {
    const message = { ...baseMsgNewNFTClassResponse } as MsgNewNFTClassResponse
    return message
  },

  toJSON(_: MsgNewNFTClassResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgNewNFTClassResponse>): MsgNewNFTClassResponse {
    const message = { ...baseMsgNewNFTClassResponse } as MsgNewNFTClassResponse
    return message
  }
}

const baseMsgUpdateNFTClass: object = { id: '', class_id: '', name: '', uri: '', sender: '' }

export const MsgUpdateNFTClass = {
  encode(message: MsgUpdateNFTClass, writer: Writer = Writer.create()): Writer {
    if (message.id !== '') {
      writer.uint32(10).string(message.id)
    }
    if (message.class_id !== '') {
      writer.uint32(18).string(message.class_id)
    }
    if (message.name !== '') {
      writer.uint32(26).string(message.name)
    }
    if (message.uri !== '') {
      writer.uint32(34).string(message.uri)
    }
    if (message.data !== undefined) {
      Any.encode(message.data, writer.uint32(42).fork()).ldelim()
    }
    if (message.sender !== '') {
      writer.uint32(50).string(message.sender)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateNFTClass {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateNFTClass } as MsgUpdateNFTClass
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string()
          break
        case 2:
          message.class_id = reader.string()
          break
        case 3:
          message.name = reader.string()
          break
        case 4:
          message.uri = reader.string()
          break
        case 5:
          message.data = Any.decode(reader, reader.uint32())
          break
        case 6:
          message.sender = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgUpdateNFTClass {
    const message = { ...baseMsgUpdateNFTClass } as MsgUpdateNFTClass
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id)
    } else {
      message.id = ''
    }
    if (object.class_id !== undefined && object.class_id !== null) {
      message.class_id = String(object.class_id)
    } else {
      message.class_id = ''
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name)
    } else {
      message.name = ''
    }
    if (object.uri !== undefined && object.uri !== null) {
      message.uri = String(object.uri)
    } else {
      message.uri = ''
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = Any.fromJSON(object.data)
    } else {
      message.data = undefined
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender)
    } else {
      message.sender = ''
    }
    return message
  },

  toJSON(message: MsgUpdateNFTClass): unknown {
    const obj: any = {}
    message.id !== undefined && (obj.id = message.id)
    message.class_id !== undefined && (obj.class_id = message.class_id)
    message.name !== undefined && (obj.name = message.name)
    message.uri !== undefined && (obj.uri = message.uri)
    message.data !== undefined && (obj.data = message.data ? Any.toJSON(message.data) : undefined)
    message.sender !== undefined && (obj.sender = message.sender)
    return obj
  },

  fromPartial(object: DeepPartial<MsgUpdateNFTClass>): MsgUpdateNFTClass {
    const message = { ...baseMsgUpdateNFTClass } as MsgUpdateNFTClass
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id
    } else {
      message.id = ''
    }
    if (object.class_id !== undefined && object.class_id !== null) {
      message.class_id = object.class_id
    } else {
      message.class_id = ''
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name
    } else {
      message.name = ''
    }
    if (object.uri !== undefined && object.uri !== null) {
      message.uri = object.uri
    } else {
      message.uri = ''
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = Any.fromPartial(object.data)
    } else {
      message.data = undefined
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender
    } else {
      message.sender = ''
    }
    return message
  }
}

const baseMsgUpdateNFTClassResponse: object = {}

export const MsgUpdateNFTClassResponse = {
  encode(_: MsgUpdateNFTClassResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateNFTClassResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateNFTClassResponse } as MsgUpdateNFTClassResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgUpdateNFTClassResponse {
    const message = { ...baseMsgUpdateNFTClassResponse } as MsgUpdateNFTClassResponse
    return message
  },

  toJSON(_: MsgUpdateNFTClassResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgUpdateNFTClassResponse>): MsgUpdateNFTClassResponse {
    const message = { ...baseMsgUpdateNFTClassResponse } as MsgUpdateNFTClassResponse
    return message
  }
}

const baseMsgMintNFT: object = { id: '', class_id: '', name: '', uri: '', sender: '', recipient: '' }

export const MsgMintNFT = {
  encode(message: MsgMintNFT, writer: Writer = Writer.create()): Writer {
    if (message.id !== '') {
      writer.uint32(10).string(message.id)
    }
    if (message.class_id !== '') {
      writer.uint32(18).string(message.class_id)
    }
    if (message.name !== '') {
      writer.uint32(26).string(message.name)
    }
    if (message.uri !== '') {
      writer.uint32(34).string(message.uri)
    }
    if (message.data !== undefined) {
      Any.encode(message.data, writer.uint32(42).fork()).ldelim()
    }
    if (message.sender !== '') {
      writer.uint32(50).string(message.sender)
    }
    if (message.recipient !== '') {
      writer.uint32(58).string(message.recipient)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgMintNFT {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgMintNFT } as MsgMintNFT
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string()
          break
        case 2:
          message.class_id = reader.string()
          break
        case 3:
          message.name = reader.string()
          break
        case 4:
          message.uri = reader.string()
          break
        case 5:
          message.data = Any.decode(reader, reader.uint32())
          break
        case 6:
          message.sender = reader.string()
          break
        case 7:
          message.recipient = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgMintNFT {
    const message = { ...baseMsgMintNFT } as MsgMintNFT
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id)
    } else {
      message.id = ''
    }
    if (object.class_id !== undefined && object.class_id !== null) {
      message.class_id = String(object.class_id)
    } else {
      message.class_id = ''
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name)
    } else {
      message.name = ''
    }
    if (object.uri !== undefined && object.uri !== null) {
      message.uri = String(object.uri)
    } else {
      message.uri = ''
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = Any.fromJSON(object.data)
    } else {
      message.data = undefined
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender)
    } else {
      message.sender = ''
    }
    if (object.recipient !== undefined && object.recipient !== null) {
      message.recipient = String(object.recipient)
    } else {
      message.recipient = ''
    }
    return message
  },

  toJSON(message: MsgMintNFT): unknown {
    const obj: any = {}
    message.id !== undefined && (obj.id = message.id)
    message.class_id !== undefined && (obj.class_id = message.class_id)
    message.name !== undefined && (obj.name = message.name)
    message.uri !== undefined && (obj.uri = message.uri)
    message.data !== undefined && (obj.data = message.data ? Any.toJSON(message.data) : undefined)
    message.sender !== undefined && (obj.sender = message.sender)
    message.recipient !== undefined && (obj.recipient = message.recipient)
    return obj
  },

  fromPartial(object: DeepPartial<MsgMintNFT>): MsgMintNFT {
    const message = { ...baseMsgMintNFT } as MsgMintNFT
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id
    } else {
      message.id = ''
    }
    if (object.class_id !== undefined && object.class_id !== null) {
      message.class_id = object.class_id
    } else {
      message.class_id = ''
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name
    } else {
      message.name = ''
    }
    if (object.uri !== undefined && object.uri !== null) {
      message.uri = object.uri
    } else {
      message.uri = ''
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = Any.fromPartial(object.data)
    } else {
      message.data = undefined
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender
    } else {
      message.sender = ''
    }
    if (object.recipient !== undefined && object.recipient !== null) {
      message.recipient = object.recipient
    } else {
      message.recipient = ''
    }
    return message
  }
}

const baseMsgMintNFTResponse: object = {}

export const MsgMintNFTResponse = {
  encode(_: MsgMintNFTResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgMintNFTResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgMintNFTResponse } as MsgMintNFTResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgMintNFTResponse {
    const message = { ...baseMsgMintNFTResponse } as MsgMintNFTResponse
    return message
  },

  toJSON(_: MsgMintNFTResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgMintNFTResponse>): MsgMintNFTResponse {
    const message = { ...baseMsgMintNFTResponse } as MsgMintNFTResponse
    return message
  }
}

const baseMsgUpdateNFT: object = { id: '', class_id: '', uri: '', sender: '' }

export const MsgUpdateNFT = {
  encode(message: MsgUpdateNFT, writer: Writer = Writer.create()): Writer {
    if (message.id !== '') {
      writer.uint32(10).string(message.id)
    }
    if (message.class_id !== '') {
      writer.uint32(18).string(message.class_id)
    }
    if (message.uri !== '') {
      writer.uint32(26).string(message.uri)
    }
    if (message.data !== undefined) {
      Any.encode(message.data, writer.uint32(34).fork()).ldelim()
    }
    if (message.sender !== '') {
      writer.uint32(42).string(message.sender)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateNFT {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateNFT } as MsgUpdateNFT
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string()
          break
        case 2:
          message.class_id = reader.string()
          break
        case 3:
          message.uri = reader.string()
          break
        case 4:
          message.data = Any.decode(reader, reader.uint32())
          break
        case 5:
          message.sender = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgUpdateNFT {
    const message = { ...baseMsgUpdateNFT } as MsgUpdateNFT
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id)
    } else {
      message.id = ''
    }
    if (object.class_id !== undefined && object.class_id !== null) {
      message.class_id = String(object.class_id)
    } else {
      message.class_id = ''
    }
    if (object.uri !== undefined && object.uri !== null) {
      message.uri = String(object.uri)
    } else {
      message.uri = ''
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = Any.fromJSON(object.data)
    } else {
      message.data = undefined
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender)
    } else {
      message.sender = ''
    }
    return message
  },

  toJSON(message: MsgUpdateNFT): unknown {
    const obj: any = {}
    message.id !== undefined && (obj.id = message.id)
    message.class_id !== undefined && (obj.class_id = message.class_id)
    message.uri !== undefined && (obj.uri = message.uri)
    message.data !== undefined && (obj.data = message.data ? Any.toJSON(message.data) : undefined)
    message.sender !== undefined && (obj.sender = message.sender)
    return obj
  },

  fromPartial(object: DeepPartial<MsgUpdateNFT>): MsgUpdateNFT {
    const message = { ...baseMsgUpdateNFT } as MsgUpdateNFT
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id
    } else {
      message.id = ''
    }
    if (object.class_id !== undefined && object.class_id !== null) {
      message.class_id = object.class_id
    } else {
      message.class_id = ''
    }
    if (object.uri !== undefined && object.uri !== null) {
      message.uri = object.uri
    } else {
      message.uri = ''
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = Any.fromPartial(object.data)
    } else {
      message.data = undefined
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender
    } else {
      message.sender = ''
    }
    return message
  }
}

const baseMsgUpdateNFTResponse: object = {}

export const MsgUpdateNFTResponse = {
  encode(_: MsgUpdateNFTResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateNFTResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateNFTResponse } as MsgUpdateNFTResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgUpdateNFTResponse {
    const message = { ...baseMsgUpdateNFTResponse } as MsgUpdateNFTResponse
    return message
  },

  toJSON(_: MsgUpdateNFTResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgUpdateNFTResponse>): MsgUpdateNFTResponse {
    const message = { ...baseMsgUpdateNFTResponse } as MsgUpdateNFTResponse
    return message
  }
}

const baseMsgBurnNFT: object = { id: '', class_id: '', sender: '' }

export const MsgBurnNFT = {
  encode(message: MsgBurnNFT, writer: Writer = Writer.create()): Writer {
    if (message.id !== '') {
      writer.uint32(10).string(message.id)
    }
    if (message.class_id !== '') {
      writer.uint32(18).string(message.class_id)
    }
    if (message.sender !== '') {
      writer.uint32(26).string(message.sender)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgBurnNFT {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgBurnNFT } as MsgBurnNFT
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string()
          break
        case 2:
          message.class_id = reader.string()
          break
        case 3:
          message.sender = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgBurnNFT {
    const message = { ...baseMsgBurnNFT } as MsgBurnNFT
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id)
    } else {
      message.id = ''
    }
    if (object.class_id !== undefined && object.class_id !== null) {
      message.class_id = String(object.class_id)
    } else {
      message.class_id = ''
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender)
    } else {
      message.sender = ''
    }
    return message
  },

  toJSON(message: MsgBurnNFT): unknown {
    const obj: any = {}
    message.id !== undefined && (obj.id = message.id)
    message.class_id !== undefined && (obj.class_id = message.class_id)
    message.sender !== undefined && (obj.sender = message.sender)
    return obj
  },

  fromPartial(object: DeepPartial<MsgBurnNFT>): MsgBurnNFT {
    const message = { ...baseMsgBurnNFT } as MsgBurnNFT
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id
    } else {
      message.id = ''
    }
    if (object.class_id !== undefined && object.class_id !== null) {
      message.class_id = object.class_id
    } else {
      message.class_id = ''
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender
    } else {
      message.sender = ''
    }
    return message
  }
}

const baseMsgBurnNFTResponse: object = {}

export const MsgBurnNFTResponse = {
  encode(_: MsgBurnNFTResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgBurnNFTResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgBurnNFTResponse } as MsgBurnNFTResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgBurnNFTResponse {
    const message = { ...baseMsgBurnNFTResponse } as MsgBurnNFTResponse
    return message
  },

  toJSON(_: MsgBurnNFTResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgBurnNFTResponse>): MsgBurnNFTResponse {
    const message = { ...baseMsgBurnNFTResponse } as MsgBurnNFTResponse
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  /** NewDenom defines a method for create a new denom. */
  NewDenom(request: MsgNewDenom): Promise<MsgNewDenomResponse>
  NewNFTClass(request: MsgNewNFTClass): Promise<MsgNewNFTClassResponse>
  UpdateNFTClass(request: MsgUpdateNFTClass): Promise<MsgUpdateNFTClassResponse>
  /** MintNFT defines a method for mint a new nft */
  MintNFT(request: MsgMintNFT): Promise<MsgMintNFTResponse>
  /** UpdateNFT defines a method for editing a nft. */
  UpdateNFT(request: MsgUpdateNFT): Promise<MsgUpdateNFTResponse>
  /** BurnNFT defines a method for burning a nft. */
  BurnNFT(request: MsgBurnNFT): Promise<MsgBurnNFTResponse>
  UpdateUserInfo(request: MsgUpdateUserInfo): Promise<MsgUpdateUserInfoResponse>
  UpdateUserRelation(request: MsgUpdateUserRelation): Promise<MsgUpdateUserRelationResponse>
  UpdateAppInfo(request: MsgUpdateAppInfo): Promise<MsgUpdateAppInfoResponse>
  CreateDidRegistry(request: MsgCreateDidRegistry): Promise<MsgCreateDidRegistryResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  NewDenom(request: MsgNewDenom): Promise<MsgNewDenomResponse> {
    const data = MsgNewDenom.encode(request).finish()
    const promise = this.rpc.request('misesid.misestm.v1beta1.Msg', 'NewDenom', data)
    return promise.then((data) => MsgNewDenomResponse.decode(new Reader(data)))
  }

  NewNFTClass(request: MsgNewNFTClass): Promise<MsgNewNFTClassResponse> {
    const data = MsgNewNFTClass.encode(request).finish()
    const promise = this.rpc.request('misesid.misestm.v1beta1.Msg', 'NewNFTClass', data)
    return promise.then((data) => MsgNewNFTClassResponse.decode(new Reader(data)))
  }

  UpdateNFTClass(request: MsgUpdateNFTClass): Promise<MsgUpdateNFTClassResponse> {
    const data = MsgUpdateNFTClass.encode(request).finish()
    const promise = this.rpc.request('misesid.misestm.v1beta1.Msg', 'UpdateNFTClass', data)
    return promise.then((data) => MsgUpdateNFTClassResponse.decode(new Reader(data)))
  }

  MintNFT(request: MsgMintNFT): Promise<MsgMintNFTResponse> {
    const data = MsgMintNFT.encode(request).finish()
    const promise = this.rpc.request('misesid.misestm.v1beta1.Msg', 'MintNFT', data)
    return promise.then((data) => MsgMintNFTResponse.decode(new Reader(data)))
  }

  UpdateNFT(request: MsgUpdateNFT): Promise<MsgUpdateNFTResponse> {
    const data = MsgUpdateNFT.encode(request).finish()
    const promise = this.rpc.request('misesid.misestm.v1beta1.Msg', 'UpdateNFT', data)
    return promise.then((data) => MsgUpdateNFTResponse.decode(new Reader(data)))
  }

  BurnNFT(request: MsgBurnNFT): Promise<MsgBurnNFTResponse> {
    const data = MsgBurnNFT.encode(request).finish()
    const promise = this.rpc.request('misesid.misestm.v1beta1.Msg', 'BurnNFT', data)
    return promise.then((data) => MsgBurnNFTResponse.decode(new Reader(data)))
  }

  UpdateUserInfo(request: MsgUpdateUserInfo): Promise<MsgUpdateUserInfoResponse> {
    const data = MsgUpdateUserInfo.encode(request).finish()
    const promise = this.rpc.request('misesid.misestm.v1beta1.Msg', 'UpdateUserInfo', data)
    return promise.then((data) => MsgUpdateUserInfoResponse.decode(new Reader(data)))
  }

  UpdateUserRelation(request: MsgUpdateUserRelation): Promise<MsgUpdateUserRelationResponse> {
    const data = MsgUpdateUserRelation.encode(request).finish()
    const promise = this.rpc.request('misesid.misestm.v1beta1.Msg', 'UpdateUserRelation', data)
    return promise.then((data) => MsgUpdateUserRelationResponse.decode(new Reader(data)))
  }

  UpdateAppInfo(request: MsgUpdateAppInfo): Promise<MsgUpdateAppInfoResponse> {
    const data = MsgUpdateAppInfo.encode(request).finish()
    const promise = this.rpc.request('misesid.misestm.v1beta1.Msg', 'UpdateAppInfo', data)
    return promise.then((data) => MsgUpdateAppInfoResponse.decode(new Reader(data)))
  }

  CreateDidRegistry(request: MsgCreateDidRegistry): Promise<MsgCreateDidRegistryResponse> {
    const data = MsgCreateDidRegistry.encode(request).finish()
    const promise = this.rpc.request('misesid.misestm.v1beta1.Msg', 'CreateDidRegistry', data)
    return promise.then((data) => MsgCreateDidRegistryResponse.decode(new Reader(data)))
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
