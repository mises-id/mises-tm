/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

/**
* Event allows application developers to attach additional information to
ResponseBeginBlock, ResponseEndBlock, ResponseCheckTx and ResponseDeliverTx.
Later, transactions may be queried using these events.
*/
export interface AbciEvent {
  type?: string;
  attributes?: AbciEventAttribute[];
}

/**
 * EventAttribute is a single key-value pair, associated with an event.
 */
export interface AbciEventAttribute {
  /** @format byte */
  key?: string;

  /** @format byte */
  value?: string;
  index?: boolean;
}

/**
* `Any` contains an arbitrary serialized protocol buffer message along with a
URL that describes the type of the serialized message.

Protobuf library provides support to pack/unpack Any values in the form
of utility functions or additional generated methods of the Any type.

Example 1: Pack and unpack a message in C++.

    Foo foo = ...;
    Any any;
    any.PackFrom(foo);
    ...
    if (any.UnpackTo(&foo)) {
      ...
    }

Example 2: Pack and unpack a message in Java.

    Foo foo = ...;
    Any any = Any.pack(foo);
    ...
    if (any.is(Foo.class)) {
      foo = any.unpack(Foo.class);
    }

 Example 3: Pack and unpack a message in Python.

    foo = Foo(...)
    any = Any()
    any.Pack(foo)
    ...
    if any.Is(Foo.DESCRIPTOR):
      any.Unpack(foo)
      ...

 Example 4: Pack and unpack a message in Go

     foo := &pb.Foo{...}
     any, err := anypb.New(foo)
     if err != nil {
       ...
     }
     ...
     foo := &pb.Foo{}
     if err := any.UnmarshalTo(foo); err != nil {
       ...
     }

The pack methods provided by protobuf library will by default use
'type.googleapis.com/full.type.name' as the type URL and the unpack
methods only use the fully qualified type name after the last '/'
in the type URL, for example "foo.bar.com/x/y.z" will yield type
name "y.z".


JSON
====
The JSON representation of an `Any` value uses the regular
representation of the deserialized, embedded message, with an
additional field `@type` which contains the type URL. Example:

    package google.profile;
    message Person {
      string first_name = 1;
      string last_name = 2;
    }

    {
      "@type": "type.googleapis.com/google.profile.Person",
      "firstName": <string>,
      "lastName": <string>
    }

If the embedded message type is well-known and has a custom JSON
representation, that representation will be embedded adding a field
`value` which holds the custom JSON in addition to the `@type`
field. Example (for message [google.protobuf.Duration][]):

    {
      "@type": "type.googleapis.com/google.protobuf.Duration",
      "value": "1.212s"
    }
*/
export interface ProtobufAny {
  /**
   * A URL/resource name that uniquely identifies the type of the serialized
   * protocol buffer message. This string must contain at least
   * one "/" character. The last segment of the URL's path must represent
   * the fully qualified name of the type (as in
   * `path/google.protobuf.Duration`). The name should be in a canonical form
   * (e.g., leading "." is not accepted).
   *
   * In practice, teams usually precompile into the binary all types that they
   * expect it to use in the context of Any. However, for URLs which use the
   * scheme `http`, `https`, or no scheme, one can optionally set up a type
   * server that maps type URLs to message definitions as follows:
   *
   * * If no scheme is provided, `https` is assumed.
   * * An HTTP GET on the URL must yield a [google.protobuf.Type][]
   *   value in binary format, or produce an error.
   * * Applications are allowed to cache lookup results based on the
   *   URL, or have them precompiled into a binary to avoid any
   *   lookup. Therefore, binary compatibility needs to be preserved
   *   on changes to types. (Use versioned type names to manage
   *   breaking changes.)
   *
   * Note: this functionality is not currently available in the official
   * protobuf release, and it is not used for type URLs beginning with
   * type.googleapis.com.
   *
   * Schemes other than `http`, `https` (or the empty scheme) might be
   * used with implementation specific semantics.
   */
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

/**
 * ABCIMessageLog defines a structure containing an indexed tx ABCI message log.
 */
export interface V1Beta1ABCIMessageLog {
  /** @format int64 */
  msg_index?: number;
  log?: string;

  /**
   * Events contains a slice of Event objects that were emitted during some
   * execution.
   */
  events?: V1Beta1StringEvent[];
}

export interface V1Beta1AppFeeGrant {
  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  spend_limit?: V1Beta1Coin;
  period?: string;

  /** @format date-time */
  expiration?: string;
}

export interface V1Beta1AppInfo {
  creator?: string;

  /** @format uint64 */
  id?: string;
  appid?: string;
  pub_info?: V1Beta1PublicAppInfo;

  /** @format uint64 */
  version?: string;
}

/**
* Attribute defines an attribute wrapper where the key and value are
strings instead of raw bytes.
*/
export interface V1Beta1Attribute {
  key?: string;
  value?: string;
}

/**
* Coin defines a token with a denomination and an amount.

NOTE: The amount field is an Int which implements the custom method
signatures required by gogoproto.
*/
export interface V1Beta1Coin {
  denom?: string;
  amount?: string;
}

/**
* DenomUnit represents a struct that describes a given
denomination unit of the basic token.
*/
export interface V1Beta1DenomUnit {
  /** denom represents the string name of the given denom unit (e.g uatom). */
  denom?: string;

  /**
   * exponent represents power of 10 exponent that one must
   * raise the base_denom to in order to equal the given DenomUnit's denom
   * 1 denom = 1^exponent base_denom
   * (e.g. with a base_denom of uatom, one can create a DenomUnit of 'atom' with
   * exponent = 6, thus: 1 atom = 10^6 uatom).
   * @format int64
   */
  exponent?: number;
  aliases?: string[];
}

export interface V1Beta1DidRegistry {
  creator?: string;

  /** @format uint64 */
  id?: string;
  did?: string;
  pkeyDid?: string;
  pkeyType?: string;
  pkeyMultibase?: string;

  /** @format uint64 */
  version?: string;
}

/**
* Metadata represents a struct that describes
a basic token.
*/
export interface V1Beta1Metadata {
  description?: string;
  denom_units?: V1Beta1DenomUnit[];

  /** base represents the base denom (should be the DenomUnit with exponent = 0). */
  base?: string;

  /**
   * display indicates the suggested denom that should be
   * displayed in clients.
   */
  display?: string;

  /** Since: cosmos-sdk 0.43 */
  name?: string;

  /**
   * symbol is the token symbol usually shown on exchanges (eg: ATOM). This can
   * be the same as the display.
   *
   * Since: cosmos-sdk 0.43
   */
  symbol?: string;
}

export interface V1Beta1MisesID {
  mises_id?: string;
  rel_type?: string;
}

/**
 * MsgBurnNFTResponse defines the Msg/BurnNFT response type.
 */
export type V1Beta1MsgBurnNFTResponse = object;

export type V1Beta1MsgCreateDidRegistryResponse = object;

/**
 * MsgMintNFTResponse defines the Msg/MintNFT response type.
 */
export type V1Beta1MsgMintNFTResponse = object;

/**
 * MsgNewDenomResponse defines the MsgNewDenom response type.
 */
export type V1Beta1MsgNewDenomResponse = object;

/**
 * MsgNewNFTClassResponse defines the MsgNewNFTClass response type.
 */
export type V1Beta1MsgNewNFTClassResponse = object;

export type V1Beta1MsgUpdateAppInfoResponse = object;

/**
 * MsgUpdateNFTClassResponse defines the MsgUpdateNFTClass response type.
 */
export type V1Beta1MsgUpdateNFTClassResponse = object;

/**
 * MsgUpdateNFTResponse defines the MsgUpdateNFT response type.
 */
export type V1Beta1MsgUpdateNFTResponse = object;

export type V1Beta1MsgUpdateUserInfoResponse = object;

export type V1Beta1MsgUpdateUserRelationResponse = object;

/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
  /**
   * key is a value returned in PageResponse.next_key to begin
   * querying the next page most efficiently. Only one of offset or key
   * should be set.
   * @format byte
   */
  key?: string;

  /**
   * offset is a numeric offset that can be used when key is unavailable.
   * It is less efficient than using key. Only one of offset or key should
   * be set.
   * @format uint64
   */
  offset?: string;

  /**
   * limit is the total number of results to be returned in the result page.
   * If left empty it will default to a value to be set by each app.
   * @format uint64
   */
  limit?: string;

  /**
   * count_total is set to true  to indicate that the result set should include
   * a count of the total number of items available for pagination in UIs.
   * count_total is only respected when offset is used. It is ignored when key
   * is set.
   */
  count_total?: boolean;

  /**
   * reverse is set to true if results are to be returned in the descending order.
   *
   * Since: cosmos-sdk 0.43
   */
  reverse?: boolean;
}

/**
* PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }
*/
export interface V1Beta1PageResponse {
  /** @format byte */
  next_key?: string;

  /** @format uint64 */
  total?: string;
}

export interface V1Beta1PrivateUserInfo {
  enc_data?: string;
  iv?: string;
}

export interface V1Beta1PublicAppInfo {
  name?: string;
  domains?: string[];
  developer?: string;
  home_url?: string;
  icon_url?: string;
}

export interface V1Beta1PublicUserInfo {
  name?: string;
  gender?: string;
  avatar_url?: string;
  home_page_url?: string;
  emails?: string[];
  telephones?: string[];
  intro?: string;
}

export interface V1Beta1QueryAllAppInfoResponse {
  AppInfo?: V1Beta1AppInfo[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface V1Beta1QueryAllDidRegistryResponse {
  DidRegistry?: V1Beta1DidRegistry[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface V1Beta1QueryAllUserInfoResponse {
  UserInfo?: V1Beta1UserInfo[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface V1Beta1QueryAllUserRelationResponse {
  UserRelation?: V1Beta1UserRelation[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface V1Beta1QueryGetAppInfoResponse {
  AppInfo?: V1Beta1AppInfo;
}

export interface V1Beta1QueryGetDidRegistryResponse {
  DidRegistry?: V1Beta1DidRegistry;
}

export interface V1Beta1QueryGetUserInfoResponse {
  UserInfo?: V1Beta1UserInfo;
}

export interface V1Beta1QueryGetUserRelationResponse {
  UserRelation?: V1Beta1UserRelation;
}

export interface V1Beta1RestQueryAppFeeGrantResponse {
  grant?: V1Beta1AppFeeGrant;
}

export interface V1Beta1RestQueryAppResponse {
  pub_info?: V1Beta1PublicAppInfo;

  /** @format uint64 */
  version?: string;
}

export interface V1Beta1RestQueryDidResponse {
  didRegistry?: V1Beta1DidRegistry;
}

export interface V1Beta1RestQueryUserRelationResponse {
  mises_list?: V1Beta1MisesID[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface V1Beta1RestQueryUserResponse {
  pub_info?: V1Beta1PublicUserInfo;
  pri_info?: V1Beta1PrivateUserInfo;

  /** @format uint64 */
  version?: string;
}

export interface V1Beta1RestTxResponse {
  /** tx_response is the queried TxResponses. */
  tx_response?: V1Beta1TxResponse;

  /** @format int64 */
  code?: number;
}

/**
* StringEvent defines en Event object wrapper where all the attributes
contain key/value pairs that are strings instead of raw bytes.
*/
export interface V1Beta1StringEvent {
  type?: string;
  attributes?: V1Beta1Attribute[];
}

/**
* TxResponse defines a structure containing relevant tx data and metadata. The
tags are stringified and the log is JSON decoded.
*/
export interface V1Beta1TxResponse {
  /** @format int64 */
  height?: string;

  /** The transaction hash. */
  txhash?: string;
  codespace?: string;

  /**
   * Response code.
   * @format int64
   */
  code?: number;

  /** Result bytes, if any. */
  data?: string;

  /**
   * The output of the application's logger (raw string). May be
   * non-deterministic.
   */
  raw_log?: string;

  /** The output of the application's logger (typed). May be non-deterministic. */
  logs?: V1Beta1ABCIMessageLog[];

  /** Additional information. May be non-deterministic. */
  info?: string;

  /**
   * Amount of gas requested for transaction.
   * @format int64
   */
  gas_wanted?: string;

  /**
   * Amount of gas consumed by transaction.
   * @format int64
   */
  gas_used?: string;

  /** The request transaction bytes. */
  tx?: ProtobufAny;

  /**
   * Time of the previous block. For heights > 1, it's the weighted median of
   * the timestamps of the valid votes in the block.LastCommit. For height == 1,
   * it's genesis time.
   */
  timestamp?: string;

  /**
   * Events defines all the events emitted by processing a transaction. Note,
   * these events include those emitted by processing all the messages and those
   * emitted from the ante handler. Whereas Logs contains the events, with
   * additional metadata, emitted only by processing the messages.
   *
   * Since: cosmos-sdk 0.42.11, 0.44.5, 0.45
   */
  events?: AbciEvent[];
}

export interface V1Beta1UserInfo {
  creator?: string;

  /** @format uint64 */
  id?: string;
  uid?: string;
  pub_info?: V1Beta1PublicUserInfo;
  pri_info?: V1Beta1PrivateUserInfo;

  /** @format uint64 */
  version?: string;
}

export interface V1Beta1UserRelation {
  creator?: string;

  /** @format uint64 */
  id?: string;
  uidFrom?: string;
  uidTo?: string;
  isFollowing?: boolean;
  isBlocking?: boolean;
  isReferredBy?: boolean;

  /** @format uint64 */
  version?: string;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: keyof Omit<Body, "body" | "bodyUsed">;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "";
  private securityData: SecurityDataType = null as any;
  private securityWorker: null | ApiConfig<SecurityDataType>["securityWorker"] = null;
  private abortControllers = new Map<CancelToken, AbortController>();

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType) => {
    this.securityData = data;
  };

  private addQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];

    return (
      encodeURIComponent(key) +
      "=" +
      encodeURIComponent(Array.isArray(value) ? value.join(",") : typeof value === "number" ? value : `${value}`)
    );
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) =>
        typeof query[key] === "object" && !Array.isArray(query[key])
          ? this.toQueryString(query[key] as QueryParamsType)
          : this.addQueryParam(query, key),
      )
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((data, key) => {
        data.append(key, input[key]);
        return data;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  private mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format = "json",
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams = (secure && this.securityWorker && this.securityWorker(this.securityData)) || {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];

    return fetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      signal: cancelToken ? this.createAbortSignal(cancelToken) : void 0,
      body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = (null as unknown) as T;
      r.error = (null as unknown) as E;

      const data = await response[format]()
        .then((data) => {
          if (r.ok) {
            r.data = data;
          } else {
            r.error = data;
          }
          return r;
        })
        .catch((e) => {
          r.error = e;
          return r;
        });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title misestm/v1beta1/AppInfo.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryAppInfoAll
   * @summary Queries a list of AppInfo items.
   * @request GET:/mises-id/misestm/misestm/AppInfo
   */
  queryAppInfoAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<V1Beta1QueryAllAppInfoResponse, RpcStatus>({
      path: `/mises-id/misestm/misestm/AppInfo`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryAppInfo
   * @summary Queries a AppInfo by id.
   * @request GET:/mises-id/misestm/misestm/AppInfo/{id}
   */
  queryAppInfo = (id: string, params: RequestParams = {}) =>
    this.request<V1Beta1QueryGetAppInfoResponse, RpcStatus>({
      path: `/mises-id/misestm/misestm/AppInfo/${id}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryDidRegistryAll
   * @summary Queries a list of DidRegistry items.
   * @request GET:/mises-id/misestm/misestm/DidRegistry
   */
  queryDidRegistryAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<V1Beta1QueryAllDidRegistryResponse, RpcStatus>({
      path: `/mises-id/misestm/misestm/DidRegistry`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryDidRegistry
   * @summary Queries a DidRegistry by id.
   * @request GET:/mises-id/misestm/misestm/DidRegistry/{id}
   */
  queryDidRegistry = (id: string, params: RequestParams = {}) =>
    this.request<V1Beta1QueryGetDidRegistryResponse, RpcStatus>({
      path: `/mises-id/misestm/misestm/DidRegistry/${id}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryUserInfoAll
   * @summary Queries a list of UserInfo items.
   * @request GET:/mises-id/misestm/misestm/UserInfo
   */
  queryUserInfoAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<V1Beta1QueryAllUserInfoResponse, RpcStatus>({
      path: `/mises-id/misestm/misestm/UserInfo`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryUserInfo
   * @summary Queries a UserInfo by id.
   * @request GET:/mises-id/misestm/misestm/UserInfo/{id}
   */
  queryUserInfo = (id: string, params: RequestParams = {}) =>
    this.request<V1Beta1QueryGetUserInfoResponse, RpcStatus>({
      path: `/mises-id/misestm/misestm/UserInfo/${id}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryUserRelationAll
   * @summary Queries a list of UserRelation items.
   * @request GET:/mises-id/misestm/misestm/UserRelation
   */
  queryUserRelationAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<V1Beta1QueryAllUserRelationResponse, RpcStatus>({
      path: `/mises-id/misestm/misestm/UserRelation`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryUserRelation
   * @summary Queries a UserRelation by id.
   * @request GET:/mises-id/misestm/misestm/UserRelation/{id}
   */
  queryUserRelation = (id: string, params: RequestParams = {}) =>
    this.request<V1Beta1QueryGetUserRelationResponse, RpcStatus>({
      path: `/mises-id/misestm/misestm/UserRelation/${id}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags RestQuery
   * @name RestQueryQueryApp
   * @summary query app info
   * @request GET:/mises/app
   */
  restQueryQueryApp = (query?: { mises_appid?: string }, params: RequestParams = {}) =>
    this.request<V1Beta1RestQueryAppResponse, RpcStatus>({
      path: `/mises/app`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags RestQuery
   * @name RestQueryQueryAppFeeGrant
   * @summary query app info
   * @request GET:/mises/app/feegrant
   */
  restQueryQueryAppFeeGrant = (query?: { mises_appid?: string; mises_uid?: string }, params: RequestParams = {}) =>
    this.request<V1Beta1RestQueryAppFeeGrantResponse, RpcStatus>({
      path: `/mises/app/feegrant`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags RestQuery
   * @name RestQueryQueryDid
   * @summary query a did
   * @request GET:/mises/did
   */
  restQueryQueryDid = (query?: { mises_id?: string }, params: RequestParams = {}) =>
    this.request<V1Beta1RestQueryDidResponse, RpcStatus>({
      path: `/mises/did`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags RestQuery
   * @name RestQueryQueryTx
   * @summary query a tx result
   * @request GET:/mises/tx
   */
  restQueryQueryTx = (query?: { txhash?: string }, params: RequestParams = {}) =>
    this.request<V1Beta1RestTxResponse, RpcStatus>({
      path: `/mises/tx`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags RestQuery
   * @name RestQueryQueryUser
   * @summary query a user info
   * @request GET:/mises/user
   */
  restQueryQueryUser = (query?: { mises_uid?: string }, params: RequestParams = {}) =>
    this.request<V1Beta1RestQueryUserResponse, RpcStatus>({
      path: `/mises/user`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags RestQuery
   * @name RestQueryQueryUserRelation
   * @summary query user relations
   * @request GET:/mises/user/relation
   */
  restQueryQueryUserRelation = (
    query?: {
      mises_uid?: string;
      filter?: string;
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<V1Beta1RestQueryUserRelationResponse, RpcStatus>({
      path: `/mises/user/relation`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });
}
