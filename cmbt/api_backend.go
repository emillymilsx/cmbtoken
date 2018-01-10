// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package cmbtoken

import (
	"context"
	"math/big"

	"github.com/CoinMarketBrasil/cmbtoken/accounts"
	"github.com/CoinMarketBrasil/cmbtoken/common"
	"github.com/CoinMarketBrasil/cmbtoken/common/math"
	"github.com/CoinMarketBrasil/cmbtoken/core"
	"github.com/CoinMarketBrasil/cmbtoken/core/bloombits"
	"github.com/CoinMarketBrasil/cmbtoken/core/state"
	"github.com/CoinMarketBrasil/cmbtoken/core/types"
	"github.com/CoinMarketBrasil/cmbtoken/core/vm"
	"github.com/CoinMarketBrasil/cmbtoken/eth/downloader"
	"github.com/CoinMarketBrasil/cmbtoken/eth/gasprice"
	"github.com/CoinMarketBrasil/cmbtoken/ethdb"
	"github.com/CoinMarketBrasil/cmbtoken/event"
	"github.com/CoinMarketBrasil/cmbtoken/params"
	"github.com/CoinMarketBrasil/cmbtoken/rpc"
)

// cmbtApiBackend implements cmbtapi.Backend for full nodes
type cmbtApiBackend struct {
	cmbt *CMBtoken
	gpo *gasprice.Oracle
}

func (b *cmbtApiBackend) ChainConfig() *params.ChainConfig {
	return b.cmbt.chainConfig
}

func (b *cmbtApiBackend) CurrentBlock() *types.Block {
	return b.cmbt.BlockChain.CurrentBlock()
}

func (b *cmbtApiBackend) SetHead(number uint64) {
	b.cmbt.protocolManager.downloader.Cancel()
	b.cmbt.BlockChain.SetHead(number)
}

func (b *cmbtApiBackend) HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error) {
	// Pending block is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block := b.cmbt.miner.PendingBlock()
		return block.Header(), nil
	}
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.cmbt.BlockChain.CurrentBlock().Header(), nil
	}
	return b.cmbt.BlockChain.GetHeaderByNumber(uint64(blockNr)), nil
}

func (b *cmbtApiBackend) BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error) {
	// Pending block is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block := b.cmbt.miner.PendingBlock()
		return block, nil
	}
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.cmbt.BlockChain.CurrentBlock(), nil
	}
	return b.cmbt.BlockChain.GetBlockByNumber(uint64(blockNr)), nil
}

func (b *cmbtApiBackend) StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*state.StateDB, *types.Header, error) {
	// Pending state is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block, state := b.cmbt.miner.Pending()
		return state, block.Header(), nil
	}
	// Otherwise resolve the block number and return its state
	header, err := b.HeaderByNumber(ctx, blockNr)
	if header == nil || err != nil {
		return nil, nil, err
	}
	stateDb, err := b.cmbt.BlockChain().StateAt(header.Root)
	return stateDb, header, err
}

func (b *cmbtApiBackend) GetBlock(ctx context.Context, blockHash common.Hash) (*types.Block, error) {
	return b.cmbt.blockchain.GetBlockByHash(blockHash), nil
}

func (b *cmbtApiBackend) GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error) {
	return core.GetBlockReceipts(b.cmbt.chainDb, blockHash, core.GetBlockNumber(b.cmbt.chainDb, blockHash)), nil
}

func (b *cmbtApiBackend) GetTd(blockHash common.Hash) *big.Int {
	return b.cmbt.blockchain.GetTdByHash(blockHash)
}

func (b *cmbtApiBackend) GetEVM(ctx context.Context, msg core.Message, state *state.StateDB, header *types.Header, vmCfg vm.Config) (*vm.EVM, func() error, error) {
	state.SetBalance(msg.From(), math.MaxBig256)
	vmError := func() error { return nil }

	context := core.NewEVMContext(msg, header, b.cmbt.BlockChain(), nil)
	return vm.NewEVM(context, state, b.cmbt.chainConfig, vmCfg), vmError, nil
}

func (b *cmbtApiBackend) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription {
	return b.cmbt.BlockChain().SubscribeRemovedLogsEvent(ch)
}

func (b *cmbtApiBackend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	return b.cmbt.BlockChain().SubscribeChainEvent(ch)
}

func (b *cmbtApiBackend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	return b.cmbt.BlockChain().SubscribeChainHeadEvent(ch)
}

func (b *cmbtApiBackend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	return b.cmbt.BlockChain().SubscribeChainSideEvent(ch)
}

func (b *cmbtApiBackend) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	return b.cmbt.BlockChain().SubscribeLogsEvent(ch)
}

func (b *cmbtApiBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	return b.cmbt.txPool.AddLocal(signedTx)
}

func (b *cmbtApiBackend) GetPoolTransactions() (types.Transactions, error) {
	pending, err := b.cmbt.txPool.Pending()
	if err != nil {
		return nil, err
	}
	var txs types.Transactions
	for _, batch := range pending {
		txs = append(txs, batch...)
	}
	return txs, nil
}

func (b *cmbtApiBackend) GetPoolTransaction(hash common.Hash) *types.Transaction {
	return b.cmbt.txPool.Get(hash)
}

func (b *cmbtApiBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	return b.cmbt.txPool.State().GetNonce(addr), nil
}

func (b *cmbtApiBackend) Stats() (pending int, queued int) {
	return b.cmbt.txPool.Stats()
}

func (b *cmbtApiBackend) TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	return b.cmbt.txPool().Content()
}

func (b *cmbtApiBackend) SubscribeTxPreEvent(ch chan<- core.TxPreEvent) event.Subscription {
	return b.cmbt.txPool().SubscribeTxPreEvent(ch)
}

func (b *cmbtApiBackend) Downloader() *downloader.Downloader {
	return b.cmbt.Downloader()
}

func (b *cmbtApiBackend) ProtocolVersion() int {
	return b.cmbt.CmbtVersion()
}

func (b *cmbtApiBackend) SuggestPrice(ctx context.Context) (*big.Int, error) {
	return b.gpo.SuggestPrice(ctx)
}

func (b *cmbtApiBackend) ChainDb() cmbtdb.Database {
	return b.cmbt.ChainDb()
}

func (b *cmbtApiBackend) EventMux() *event.TypeMux {
	return b.cmbt.EventMux()
}

func (b *cmbtApiBackend) AccountManager() *accounts.Manager {
	return b.cmbt.AccountManager()
}

func (b *cmbtApiBackend) BloomStatus() (uint64, uint64) {
	sections, _, _ := b.cmbt.bloomIndexer.Sections()
	return params.BloomBitsBlocks, sections
}

func (b *cmbtApiBackend) ServiceFilter(ctx context.Context, session *bloombits.MatcherSession) {
	for i := 0; i < bloomFilterThreads; i++ {
		go session.Multiplex(bloomRetrievalBatch, bloomRetrievalWait, b.cmbt.bloomRequests)
	}
}
