import { txClient, queryClient, MissingWalletError , registry} from './module'

import { PublicAppInfo } from "./module/types/misestm/v1beta1/AppInfo"
import { AppInfo } from "./module/types/misestm/v1beta1/AppInfo"
import { DidRegistry } from "./module/types/misestm/v1beta1/DidRegistry"
import { MisesAccount } from "./module/types/misestm/v1beta1/MisesAccount"
import { PrivateUserInfo } from "./module/types/misestm/v1beta1/UserInfo"
import { PublicUserInfo } from "./module/types/misestm/v1beta1/UserInfo"
import { UserInfo } from "./module/types/misestm/v1beta1/UserInfo"
import { UserRelation } from "./module/types/misestm/v1beta1/UserRelation"
import { MisesID } from "./module/types/misestm/v1beta1/rest_query"
import { AppFeeGrant } from "./module/types/misestm/v1beta1/rest_query"


export { PublicAppInfo, AppInfo, DidRegistry, MisesAccount, PrivateUserInfo, PublicUserInfo, UserInfo, UserRelation, MisesID, AppFeeGrant };

async function initTxClient(vuexGetters) {
	return await txClient(vuexGetters['common/wallet/signer'], {
		addr: vuexGetters['common/env/apiTendermint']
	})
}

async function initQueryClient(vuexGetters) {
	return await queryClient({
		addr: vuexGetters['common/env/apiCosmos']
	})
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

function getStructure(template) {
	let structure = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field: any = {}
		field.name = key
		field.type = typeof value
		structure.fields.push(field)
	}
	return structure
}

const getDefaultState = () => {
	return {
				UserInfo: {},
				UserInfoAll: {},
				UserRelation: {},
				UserRelationAll: {},
				AppInfo: {},
				AppInfoAll: {},
				DidRegistry: {},
				DidRegistryAll: {},
				QueryDid: {},
				QueryUser: {},
				QueryUserRelation: {},
				QueryApp: {},
				QueryAppFeeGrant: {},
				QueryTx: {},
				
				_Structure: {
						PublicAppInfo: getStructure(PublicAppInfo.fromPartial({})),
						AppInfo: getStructure(AppInfo.fromPartial({})),
						DidRegistry: getStructure(DidRegistry.fromPartial({})),
						MisesAccount: getStructure(MisesAccount.fromPartial({})),
						PrivateUserInfo: getStructure(PrivateUserInfo.fromPartial({})),
						PublicUserInfo: getStructure(PublicUserInfo.fromPartial({})),
						UserInfo: getStructure(UserInfo.fromPartial({})),
						UserRelation: getStructure(UserRelation.fromPartial({})),
						MisesID: getStructure(MisesID.fromPartial({})),
						AppFeeGrant: getStructure(AppFeeGrant.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getUserInfo: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.UserInfo[JSON.stringify(params)] ?? {}
		},
				getUserInfoAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.UserInfoAll[JSON.stringify(params)] ?? {}
		},
				getUserRelation: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.UserRelation[JSON.stringify(params)] ?? {}
		},
				getUserRelationAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.UserRelationAll[JSON.stringify(params)] ?? {}
		},
				getAppInfo: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.AppInfo[JSON.stringify(params)] ?? {}
		},
				getAppInfoAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.AppInfoAll[JSON.stringify(params)] ?? {}
		},
				getDidRegistry: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.DidRegistry[JSON.stringify(params)] ?? {}
		},
				getDidRegistryAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.DidRegistryAll[JSON.stringify(params)] ?? {}
		},
				getQueryDid: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.QueryDid[JSON.stringify(params)] ?? {}
		},
				getQueryUser: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.QueryUser[JSON.stringify(params)] ?? {}
		},
				getQueryUserRelation: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.QueryUserRelation[JSON.stringify(params)] ?? {}
		},
				getQueryApp: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.QueryApp[JSON.stringify(params)] ?? {}
		},
				getQueryAppFeeGrant: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.QueryAppFeeGrant[JSON.stringify(params)] ?? {}
		},
				getQueryTx: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.QueryTx[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: misesid.misestm.v1beta1 initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryUserInfo({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryUserInfo( key.id)).data
				
					
				commit('QUERY', { query: 'UserInfo', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryUserInfo', payload: { options: { all }, params: {...key},query }})
				return getters['getUserInfo']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryUserInfo API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryUserInfoAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryUserInfoAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryUserInfoAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'UserInfoAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryUserInfoAll', payload: { options: { all }, params: {...key},query }})
				return getters['getUserInfoAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryUserInfoAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryUserRelation({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryUserRelation( key.id)).data
				
					
				commit('QUERY', { query: 'UserRelation', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryUserRelation', payload: { options: { all }, params: {...key},query }})
				return getters['getUserRelation']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryUserRelation API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryUserRelationAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryUserRelationAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryUserRelationAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'UserRelationAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryUserRelationAll', payload: { options: { all }, params: {...key},query }})
				return getters['getUserRelationAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryUserRelationAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryAppInfo({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryAppInfo( key.id)).data
				
					
				commit('QUERY', { query: 'AppInfo', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryAppInfo', payload: { options: { all }, params: {...key},query }})
				return getters['getAppInfo']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryAppInfo API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryAppInfoAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryAppInfoAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryAppInfoAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'AppInfoAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryAppInfoAll', payload: { options: { all }, params: {...key},query }})
				return getters['getAppInfoAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryAppInfoAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDidRegistry({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryDidRegistry( key.id)).data
				
					
				commit('QUERY', { query: 'DidRegistry', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDidRegistry', payload: { options: { all }, params: {...key},query }})
				return getters['getDidRegistry']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDidRegistry API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDidRegistryAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryDidRegistryAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryDidRegistryAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'DidRegistryAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDidRegistryAll', payload: { options: { all }, params: {...key},query }})
				return getters['getDidRegistryAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDidRegistryAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async RestQueryQueryDid({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.restQueryQueryDid(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.restQueryQueryDid({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'QueryDid', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'RestQueryQueryDid', payload: { options: { all }, params: {...key},query }})
				return getters['getQueryDid']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:RestQueryQueryDid API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async RestQueryQueryUser({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.restQueryQueryUser(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.restQueryQueryUser({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'QueryUser', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'RestQueryQueryUser', payload: { options: { all }, params: {...key},query }})
				return getters['getQueryUser']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:RestQueryQueryUser API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async RestQueryQueryUserRelation({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.restQueryQueryUserRelation(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.restQueryQueryUserRelation({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'QueryUserRelation', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'RestQueryQueryUserRelation', payload: { options: { all }, params: {...key},query }})
				return getters['getQueryUserRelation']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:RestQueryQueryUserRelation API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async RestQueryQueryApp({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.restQueryQueryApp(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.restQueryQueryApp({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'QueryApp', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'RestQueryQueryApp', payload: { options: { all }, params: {...key},query }})
				return getters['getQueryApp']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:RestQueryQueryApp API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async RestQueryQueryAppFeeGrant({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.restQueryQueryAppFeeGrant(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.restQueryQueryAppFeeGrant({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'QueryAppFeeGrant', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'RestQueryQueryAppFeeGrant', payload: { options: { all }, params: {...key},query }})
				return getters['getQueryAppFeeGrant']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:RestQueryQueryAppFeeGrant API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async RestQueryQueryTx({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.restQueryQueryTx(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.restQueryQueryTx({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'QueryTx', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'RestQueryQueryTx', payload: { options: { all }, params: {...key},query }})
				return getters['getQueryTx']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:RestQueryQueryTx API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgUpdateNFT({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateNFT(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateNFT:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateNFT:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgNewDenom({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgNewDenom(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgNewDenom:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgNewDenom:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateUserRelation({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateUserRelation(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateUserRelation:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateUserRelation:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateNFTClass({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateNFTClass(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateNFTClass:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateNFTClass:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgBurnNFT({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgBurnNFT(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBurnNFT:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgBurnNFT:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateAppInfo({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateAppInfo(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateAppInfo:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateAppInfo:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgNewNFTClass({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgNewNFTClass(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgNewNFTClass:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgNewNFTClass:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateUserInfo({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateUserInfo(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateUserInfo:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateUserInfo:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgMintNFT({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgMintNFT(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgMintNFT:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgMintNFT:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateDidRegistry({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateDidRegistry(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateDidRegistry:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateDidRegistry:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgUpdateNFT({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateNFT(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateNFT:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateNFT:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgNewDenom({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgNewDenom(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgNewDenom:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgNewDenom:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateUserRelation({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateUserRelation(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateUserRelation:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateUserRelation:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateNFTClass({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateNFTClass(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateNFTClass:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateNFTClass:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgBurnNFT({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgBurnNFT(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBurnNFT:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgBurnNFT:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateAppInfo({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateAppInfo(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateAppInfo:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateAppInfo:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgNewNFTClass({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgNewNFTClass(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgNewNFTClass:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgNewNFTClass:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateUserInfo({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateUserInfo(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateUserInfo:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateUserInfo:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgMintNFT({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgMintNFT(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgMintNFT:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgMintNFT:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateDidRegistry({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateDidRegistry(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateDidRegistry:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateDidRegistry:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
