package ws

import (
	"github.com/mattermost/focalboard/server/model"
)

const (
	websocketActionAuth                 = "AUTH"
	websocketActionSubscribeWorkspace   = "SUBSCRIBE_WORKSPACE"
	websocketActionUnsubscribeWorkspace = "UNSUBSCRIBE_WORKSPACE"
	websocketActionSubscribeBlocks      = "SUBSCRIBE_BLOCKS"
	websocketActionUnsubscribeBlocks    = "UNSUBSCRIBE_BLOCKS"
	websocketActionUpdateBlock          = "UPDATE_BLOCK"
	websocketActionUpdateConfig         = "UPDATE_CLIENT_CONFIG"
)

type Adapter interface {
	BroadcastBlockChange(workspaceID string, block model.Block)
	BroadcastBlockDelete(workspaceID, blockID, parentID string)
	BroadcastConfigChange(clientConfig model.ClientConfig)
}
