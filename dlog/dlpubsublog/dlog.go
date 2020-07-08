package dlpubsublog

import (
	"github.com/libp2p/go-libp2p-pubsub/util"
	"go.uber.org/zap"
)

var L *zap.Logger

func init() {
	L = util.GetXDebugLog("lpubsub")
}
