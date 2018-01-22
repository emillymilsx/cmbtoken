// Copyright 2017 The CMBToken Authors
// This file is part of the cmbtoken library.
//
// The cmbtoken library is free software: you can redistribute it and/or modify
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
// along with the cmbtoken library. If not, see <http://www.gnu.org/licenses/>.

package core

import "github.com/CoinMarketBrasil/cmbtoken/common"

// BadHashes represent a set of manually tracked bad hashes (usually hard forks)
var BadHashes = map[common.Hash]bool{
	common.HexToHash("05bef30ef572270f654746da22639a7a0c97dd97a7050b9e252391996aaeb689"): true,
	common.HexToHash("7d05d08cbc596a2e5e4f13b80a743e53e09221b5323c3a61946b20873e58583f"): true,
}

//server connect pos
var BadHashes = map[common.Hash]bool{
	common.HexToHash("c7e616822f366fb1b5e0756af498cc11d2c0862edcb32ca65882f622ff39de1b"): true,
	common.HexToHash("4b0004613db7f12fe2176027efcf7358e6688a2bbf6de60d112b409817e425fd"): true,
}
