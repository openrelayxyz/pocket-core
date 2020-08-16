module github.com/pokt-network/pocket-core

go 1.13

require (
	github.com/go-kit/kit v0.10.0
	github.com/gogo/protobuf v1.3.1
	github.com/hashicorp/golang-lru v0.5.4
	github.com/julienschmidt/httprouter v1.3.0
	github.com/magiconair/properties v1.8.1
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.5.1
	github.com/regen-network/cosmos-proto v0.3.0 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/src-d/proteus v1.3.3 // indirect
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/go-amino v0.15.0
	github.com/tendermint/iavl v0.14.0
	github.com/tendermint/tendermint v0.33.7
	github.com/tendermint/tm-db v0.5.1
	github.com/willf/bitset v1.1.10 // indirect
	github.com/willf/bloom v2.0.3+incompatible
	golang.org/x/crypto v0.0.0-20200429183012-4b2356b1ed79
	google.golang.org/grpc v1.30.0 // indirect
	google.golang.org/protobuf v1.25.0
	gopkg.in/h2non/gock.v1 v1.0.15
	gopkg.in/src-d/go-parse-utils.v1 v1.1.2 // indirect
	gopkg.in/src-d/proteus.v1 v1.3.3 // indirect
	gopkg.in/yaml.v2 v2.3.0
)

replace github.com/tendermint/tendermint => github.com/pokt-network/tendermint v0.32.11-0.20200813180226-426e34f7a488
