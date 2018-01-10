
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

// Contains the metrics collected by the downloader.

package downloader

import (
	"github.com/CoinMarketBrasil/cmbtoken/metrics"
)

var (
	headerInMeter      = metrics.NewMeter("cmbtoken/downloader/headers/in")
	headerReqTimer     = metrics.NewTimer("cmbtoken/downloader/headers/req")
	headerDropMeter    = metrics.NewMeter("cmbtoken/downloader/headers/drop")
	headerTimeoutMeter = metrics.NewMeter("cmbtoken/downloader/headers/timeout")

	bodyInMeter      = metrics.NewMeter("cmbtoken/downloader/bodies/in")
	bodyReqTimer     = metrics.NewTimer("cmbtoken/downloader/bodies/req")
	bodyDropMeter    = metrics.NewMeter("cmbtoken/downloader/bodies/drop")
	bodyTimeoutMeter = metrics.NewMeter("cmbtoken/downloader/bodies/timeout")

	receiptInMeter      = metrics.NewMeter("cmbtoken/downloader/receipts/in")
	receiptReqTimer     = metrics.NewTimer("cmbtoken/downloader/receipts/req")
	receiptDropMeter    = metrics.NewMeter("cmbtoken/downloader/receipts/drop")
	receiptTimeoutMeter = metrics.NewMeter("cmbtoken/downloader/receipts/timeout")

	stateInMeter   = metrics.NewMeter("cmbtoken/downloader/states/in")
	stateDropMeter = metrics.NewMeter("cmbtoken/downloader/states/drop")
)
