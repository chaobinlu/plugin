// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

type CreateCallTx struct {
	Amount   uint64 `json:"amount"`
	Code     string `json:"code"`
	GasLimit uint64 `json:"gasLimit"`
	GasPrice uint32 `json:"gasPrice"`
	Note     string `json:"note"`
	Alias    string `json:"alias"`
	Fee      int64  `json:"fee"`
	Name     string `json:"name"`
	IsCreate bool   `json:"isCreate"`
}
