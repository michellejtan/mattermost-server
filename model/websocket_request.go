// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"github.com/mattermost/mattermost-server/v6/shared/i18n"
	"github.com/mattermost/mattermost-server/v6/shared/mlog"

	"github.com/vmihailenco/msgpack/v5"
)

// WebSocketRequest represents a request made to the server through a websocket.
type WebSocketRequest struct {
	// Client-provided fields
	Seq    int64                  `json:"seq" msgpack:"seq"`       // A counter which is incremented for every request made.
	Action string                 `json:"action" msgpack:"action"` // The action to perform for a request. For example: get_statuses, user_typing.
	Data   map[string]interface{} `json:"data" msgpack:"data"`     // The metadata for an action.

	// Server-provided fields
	Session Session            `json:"-" msgpack:"-"`
	T       i18n.TranslateFunc `json:"-" msgpack:"-"`
	Locale  string             `json:"-" msgpack:"-"`
	Logger  mlog.LoggerIFace   `json:"-" msgpack:"-"`
}

func (o *WebSocketRequest) Clone() (*WebSocketRequest, error) {
	buf, err := msgpack.Marshal(o)
	if err != nil {
		return nil, err
	}
	var ret WebSocketRequest
	err = msgpack.Unmarshal(buf, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
