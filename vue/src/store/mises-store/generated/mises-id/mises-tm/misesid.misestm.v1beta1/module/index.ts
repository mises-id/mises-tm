// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdateNFT } from "./types/misestm/v1beta1/tx";
import { MsgNewDenom } from "./types/misestm/v1beta1/tx";
import { MsgUpdateUserRelation } from "./types/misestm/v1beta1/tx";
import { MsgUpdateNFTClass } from "./types/misestm/v1beta1/tx";
import { MsgBurnNFT } from "./types/misestm/v1beta1/tx";
import { MsgUpdateAppInfo } from "./types/misestm/v1beta1/tx";
import { MsgNewNFTClass } from "./types/misestm/v1beta1/tx";
import { MsgUpdateUserInfo } from "./types/misestm/v1beta1/tx";
import { MsgMintNFT } from "./types/misestm/v1beta1/tx";
import { MsgCreateDidRegistry } from "./types/misestm/v1beta1/tx";


const types = [
  ["/misesid.misestm.v1beta1.MsgUpdateNFT", MsgUpdateNFT],
  ["/misesid.misestm.v1beta1.MsgNewDenom", MsgNewDenom],
  ["/misesid.misestm.v1beta1.MsgUpdateUserRelation", MsgUpdateUserRelation],
  ["/misesid.misestm.v1beta1.MsgUpdateNFTClass", MsgUpdateNFTClass],
  ["/misesid.misestm.v1beta1.MsgBurnNFT", MsgBurnNFT],
  ["/misesid.misestm.v1beta1.MsgUpdateAppInfo", MsgUpdateAppInfo],
  ["/misesid.misestm.v1beta1.MsgNewNFTClass", MsgNewNFTClass],
  ["/misesid.misestm.v1beta1.MsgUpdateUserInfo", MsgUpdateUserInfo],
  ["/misesid.misestm.v1beta1.MsgMintNFT", MsgMintNFT],
  ["/misesid.misestm.v1beta1.MsgCreateDidRegistry", MsgCreateDidRegistry],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgUpdateNFT: (data: MsgUpdateNFT): EncodeObject => ({ typeUrl: "/misesid.misestm.v1beta1.MsgUpdateNFT", value: MsgUpdateNFT.fromPartial( data ) }),
    msgNewDenom: (data: MsgNewDenom): EncodeObject => ({ typeUrl: "/misesid.misestm.v1beta1.MsgNewDenom", value: MsgNewDenom.fromPartial( data ) }),
    msgUpdateUserRelation: (data: MsgUpdateUserRelation): EncodeObject => ({ typeUrl: "/misesid.misestm.v1beta1.MsgUpdateUserRelation", value: MsgUpdateUserRelation.fromPartial( data ) }),
    msgUpdateNFTClass: (data: MsgUpdateNFTClass): EncodeObject => ({ typeUrl: "/misesid.misestm.v1beta1.MsgUpdateNFTClass", value: MsgUpdateNFTClass.fromPartial( data ) }),
    msgBurnNFT: (data: MsgBurnNFT): EncodeObject => ({ typeUrl: "/misesid.misestm.v1beta1.MsgBurnNFT", value: MsgBurnNFT.fromPartial( data ) }),
    msgUpdateAppInfo: (data: MsgUpdateAppInfo): EncodeObject => ({ typeUrl: "/misesid.misestm.v1beta1.MsgUpdateAppInfo", value: MsgUpdateAppInfo.fromPartial( data ) }),
    msgNewNFTClass: (data: MsgNewNFTClass): EncodeObject => ({ typeUrl: "/misesid.misestm.v1beta1.MsgNewNFTClass", value: MsgNewNFTClass.fromPartial( data ) }),
    msgUpdateUserInfo: (data: MsgUpdateUserInfo): EncodeObject => ({ typeUrl: "/misesid.misestm.v1beta1.MsgUpdateUserInfo", value: MsgUpdateUserInfo.fromPartial( data ) }),
    msgMintNFT: (data: MsgMintNFT): EncodeObject => ({ typeUrl: "/misesid.misestm.v1beta1.MsgMintNFT", value: MsgMintNFT.fromPartial( data ) }),
    msgCreateDidRegistry: (data: MsgCreateDidRegistry): EncodeObject => ({ typeUrl: "/misesid.misestm.v1beta1.MsgCreateDidRegistry", value: MsgCreateDidRegistry.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
