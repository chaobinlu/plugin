package init

import (
	_ "gitlab.33.cn/chain33/plugin/consensus/para"
	_ "gitlab.33.cn/chain33/plugin/consensus/pbft"
	_ "gitlab.33.cn/chain33/plugin/consensus/raft"
	_ "gitlab.33.cn/chain33/plugin/consensus/tendermint"
	_ "gitlab.33.cn/chain33/plugin/consensus/ticket"
)