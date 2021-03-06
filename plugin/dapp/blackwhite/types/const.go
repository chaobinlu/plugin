// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"reflect"

	"github.com/33cn/chain33/common/log/log15"
	"github.com/33cn/chain33/types"
)

// status
const (
	BlackwhiteStatusCreate = iota + 1
	BlackwhiteStatusPlay
	BlackwhiteStatusShow
	BlackwhiteStatusTimeout
	BlackwhiteStatusDone
)

const (
	// log for blackwhite game
	TyLogBlackwhiteCreate   = 750
	TyLogBlackwhitePlay     = 751
	TyLogBlackwhiteShow     = 752
	TyLogBlackwhiteTimeout  = 753
	TyLogBlackwhiteDone     = 754
	TyLogBlackwhiteLoopInfo = 755
)

const (
	GetBlackwhiteRoundInfo       = "GetBlackwhiteRoundInfo"
	GetBlackwhiteByStatusAndAddr = "GetBlackwhiteByStatusAndAddr"
	GetBlackwhiteloopResult      = "GetBlackwhiteloopResult"
)

var (
	BlackwhiteX = "blackwhite"
	glog        = log15.New("module", BlackwhiteX)
	//GRPCName         = "chain33.blackwhite"
	JRPCName         = "Blackwhite"
	ExecerBlackwhite = []byte(BlackwhiteX)
	actionName       = map[string]int32{
		"Create":      BlackwhiteActionCreate,
		"Play":        BlackwhiteActionPlay,
		"Show":        BlackwhiteActionShow,
		"TimeoutDone": BlackwhiteActionTimeoutDone,
	}
	logInfo = map[int64]*types.LogInfo{
		TyLogBlackwhiteCreate:   {reflect.TypeOf(ReceiptBlackwhite{}), "LogBlackwhiteCreate"},
		TyLogBlackwhitePlay:     {reflect.TypeOf(ReceiptBlackwhite{}), "LogBlackwhitePlay"},
		TyLogBlackwhiteShow:     {reflect.TypeOf(ReceiptBlackwhite{}), "LogBlackwhiteShow"},
		TyLogBlackwhiteTimeout:  {reflect.TypeOf(ReceiptBlackwhite{}), "LogBlackwhiteTimeout"},
		TyLogBlackwhiteDone:     {reflect.TypeOf(ReceiptBlackwhite{}), "LogBlackwhiteDone"},
		TyLogBlackwhiteLoopInfo: {reflect.TypeOf(ReplyLoopResults{}), "LogBlackwhiteLoopInfo"},
	}
)

func init() {
	types.AllowUserExec = append(types.AllowUserExec, ExecerBlackwhite)
}
