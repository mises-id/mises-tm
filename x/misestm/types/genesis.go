package types

import (
	"fmt"
	// this line is used by starport scaffolding # ibc/genesistype/import
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		MisesAccountList: []*MisesAccount{},
		// this line is used by starport scaffolding # ibc/genesistype/default
		// this line is used by starport scaffolding # genesis/types/default
		UserInfoList:     []*UserInfo{},
		UserRelationList: []*UserRelation{},
		AppInfoList:      []*AppInfo{},
		DidRegistryList:  []*DidRegistry{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # ibc/genesistype/validate

	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated ID in UserInfo
	UserInfoIdMap := make(map[uint64]bool)

	for _, elem := range gs.UserInfoList {
		if _, ok := UserInfoIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for UserInfo")
		}
		UserInfoIdMap[elem.Id] = true
	}
	// Check for duplicated ID in UserRelation
	UserRelationIdMap := make(map[uint64]bool)

	for _, elem := range gs.UserRelationList {
		if _, ok := UserRelationIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for UserRelation")
		}
		UserRelationIdMap[elem.Id] = true
	}
	// Check for duplicated ID in AppInfo
	AppInfoIdMap := make(map[uint64]bool)

	for _, elem := range gs.AppInfoList {
		if _, ok := AppInfoIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for AppInfo")
		}
		AppInfoIdMap[elem.Id] = true
	}
	// Check for duplicated ID in DidRegistry
	DidRegistryIdMap := make(map[uint64]bool)

	for _, elem := range gs.DidRegistryList {
		if _, ok := DidRegistryIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for DidRegistry")
		}
		DidRegistryIdMap[elem.Id] = true
	}

	return nil
}
