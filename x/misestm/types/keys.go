package types

const (
	// ModuleName defines the module name
	ModuleName = "misestm"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_misestm"

	JsonDbKey = "jsondb_misestm"

	// this line is used by starport scaffolding # ibc/keys/name
)

// this line is used by starport scaffolding # ibc/keys/port

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	DidRegistryKey      = "DidRegistry-value-"
	DidRegistryCountKey = "DidRegistry-count-"
)

const (
	AppInfoKey      = "AppInfo-value-"
	AppInfoCountKey = "AppInfo-count-"
)

const (
	UserRelationKey      = "UserRelation-value-"
	UserRelationCountKey = "UserRelation-count-"
)

const (
	UserInfoKey      = "UserInfo-value-"
	UserInfoCountKey = "UserInfo-count-"
)
