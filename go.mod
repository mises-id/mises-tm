module github.com/mises-id/mises-tm

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.43.0-rc2
	github.com/cosmos/ibc-go v1.0.0-beta1
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/rakyll/statik v0.1.7
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/spm v0.0.0-20210524110815-6d7452d2dc4a
	github.com/tendermint/tendermint v0.34.11
	github.com/tendermint/tm-db v0.6.4
	go.mongodb.org/mongo-driver v1.6.0
	google.golang.org/genproto v0.0.0-20210617175327-b9e0b3197ced
	google.golang.org/grpc v1.39.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

//replace github.com/tendermint/tm-db => github.com/mises-id/tm-db v0.6.5-0.20210711030038-fd67ffbb414b
replace github.com/tendermint/tm-db => ../tm-db

replace github.com/cosmos/iavl => ../iavl