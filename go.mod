module github.com/mises-id/mises-tm

go 1.16

require (
	github.com/btcsuite/btcd v0.22.0-beta
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
	github.com/cosmos/cosmos-sdk v0.44.5
	github.com/cosmos/gravity-bridge/module v0.0.0-20210828152730-0195030967c3
	github.com/cosmos/ibc-go v1.2.4
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.2 // indirect
	github.com/multiformats/go-multibase v0.0.3
	github.com/rakyll/statik v0.1.7
	github.com/regen-network/cosmos-proto v0.3.1 // indirect
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.9.0
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/btcd v0.1.1
	github.com/tendermint/spm v0.1.8
	github.com/tendermint/tendermint v0.34.14
	github.com/tendermint/tm-db v0.6.6
	go.mongodb.org/mongo-driver v1.8.0
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.27.1
)

replace google.golang.org/grpc => google.golang.org/grpc v1.42.0

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/tendermint/tm-db => github.com/mises-id/tm-db v0.6.5-0.20210822095222-e1ff1e0dc734

replace github.com/cosmos/iavl => github.com/mises-id/iavl v0.17.4-0.20211207035003-f9d26e6150db

replace github.com/cosmos/cosmos-sdk => github.com/mises-id/cosmos-sdk v0.44.6-0.20211209094558-a7c9c77cfc17

replace github.com/tendermint/tendermint => github.com/mises-id/tendermint v0.34.15-0.20211207033151-1f29b59c0edf

replace github.com/cosmos/gravity-bridge/module => github.com/mises-id/gravity-bridge/module v0.0.0-20211122090520-c67b7ee2244e
