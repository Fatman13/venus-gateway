package main

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"

	proof5 "github.com/filecoin-project/specs-actors/v5/actors/runtime/proof"

	"github.com/ipfs-force-community/venus-gateway/types/wallet"

	"github.com/ipfs-force-community/venus-gateway/proofevent"
	"github.com/ipfs-force-community/venus-gateway/walletevent"
)

type IGatewayPushAPI interface {
	proofevent.IProofEvent
	walletevent.IWalletEvent
}

type IGatewayAPI interface {
	proofevent.IProofEventAPI
	walletevent.IWalletEventAPI
	IGatewayPushAPI
}

type GatewayStruct struct {
	ComputeProof           func(ctx context.Context, miner address.Address, sectorInfos []proof5.SectorInfo, rand abi.PoStRandomness) ([]proof5.PoStProof, error)
	WalletHas              func(ctx context.Context, supportAccount string, addr address.Address) (bool, error)
	WalletSign             func(ctx context.Context, account string, addr address.Address, toSign []byte, meta wallet.MsgMeta) (*crypto.Signature, error)
	ListConnectedMiners    func(ctx context.Context) ([]address.Address, error)
	ListMinerConnection    func(ctx context.Context, addr address.Address) (*proofevent.MinerState, error)
	ListWalletInfo         func(ctx context.Context) ([]*walletevent.WalletDetail, error)
	ListWalletInfoByWallet func(ctx context.Context, wallet string) (*walletevent.WalletDetail, error)
}

var _ IGatewayAPI = (*GatewayAPI)(nil)

type GatewayAPI struct {
	proofevent.IProofEventAPI
	walletevent.IWalletEventAPI
	pe *proofevent.ProofEventStream
	we *walletevent.WalletEventStream
}

func NewGatewayAPI(pe *proofevent.ProofEventStream, we *walletevent.WalletEventStream) *GatewayAPI {
	return &GatewayAPI{
		IProofEventAPI:  proofevent.NewProofEventAPI(pe),
		IWalletEventAPI: walletevent.NewWalletEventAPI(we),
		pe:              pe,
		we:              we,
	}
}

func (g *GatewayAPI) ComputeProof(ctx context.Context, miner address.Address, sectorInfos []proof5.SectorInfo, rand abi.PoStRandomness) ([]proof5.PoStProof, error) {
	return g.pe.ComputeProof(ctx, miner, sectorInfos, rand)
}

func (g *GatewayAPI) WalletHas(ctx context.Context, supportAccount string, addr address.Address) (bool, error) {
	return g.we.WalletHas(ctx, supportAccount, addr)
}

func (g *GatewayAPI) WalletSign(ctx context.Context, account string, addr address.Address, toSign []byte, meta wallet.MsgMeta) (*crypto.Signature, error) {
	return g.we.WalletSign(ctx, account, addr, toSign, meta)
}

func (g *GatewayAPI) ListConnectedMiners(ctx context.Context) ([]address.Address, error) {
	return g.pe.ListConnectedMiners(ctx)
}

func (g *GatewayAPI) ListMinerConnection(ctx context.Context, addr address.Address) (*proofevent.MinerState, error) {
	return g.pe.ListMinerConnection(ctx, addr)
}

func (g *GatewayAPI) ListWalletInfo(ctx context.Context) ([]*walletevent.WalletDetail, error) {
	return g.we.ListWalletInfo(ctx)
}

func (g *GatewayAPI) ListWalletInfoByWallet(ctx context.Context, wallet string) (*walletevent.WalletDetail, error) {
	return g.we.ListWalletInfoByWallet(ctx, wallet)
}
