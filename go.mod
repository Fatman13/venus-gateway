module github.com/ipfs-force-community/venus-gateway

go 1.16

require (
	github.com/filecoin-project/go-address v0.0.5
	github.com/filecoin-project/go-jsonrpc v0.1.4-0.20210217175800-45ea43ac2bec
	github.com/filecoin-project/go-state-types v0.1.1-0.20210506134452-99b279731c48
	github.com/filecoin-project/specs-actors/v5 v5.0.1
	github.com/filecoin-project/venus-auth v1.2.1
	github.com/filecoin-project/venus-wallet v1.2.0
	github.com/gbrlsnchs/jwt/v3 v3.0.0
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.8.0
	github.com/ipfs-force-community/metrics v0.0.0-20210716075100-f6c912bf4b47
	github.com/ipfs/go-cid v0.0.7
	github.com/ipfs/go-log/v2 v2.1.3
	github.com/multiformats/go-multiaddr v0.3.3
	github.com/pkg/errors v0.9.1
	github.com/urfave/cli/v2 v2.3.0
	go.opencensus.io v0.23.0
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
)

replace github.com/filecoin-project/go-jsonrpc => github.com/ipfs-force-community/go-jsonrpc v0.1.4-0.20210731021807-68e5207079bc

replace github.com/ipfs/go-ipfs-cmds => github.com/ipfs-force-community/go-ipfs-cmds v0.6.1-0.20210521090123-4587df7fa0ab

replace github.com/ipfs-force-community/venus-gateway => ./

replace github.com/filecoin-project/venus-auth => github.com/filecoin-project/venus-auth v1.2.2-0.20210719042318-b502a60e3b30
