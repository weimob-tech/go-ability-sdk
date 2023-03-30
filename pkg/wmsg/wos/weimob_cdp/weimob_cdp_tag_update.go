package weimob_cdp

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCdpTagUpdateListener struct {
	handler msg.WosEventHandler[WeimobCdpTagUpdateEvent]
}

func (s WeimobCdpTagUpdateListener) NewMessage() msg.MsgRequest[WeimobCdpTagUpdateEvent] {
	return &msg.WosOpenMessage[WeimobCdpTagUpdateEvent]{
		MsgBody: &WeimobCdpTagUpdateEvent{},
	}
}

func (s WeimobCdpTagUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCdpTagUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCdpTagUpdateEvent]))
}

type WeimobCdpTagUpdateEvent struct {
	OpenAttrInfoVos   []WeimobCdpTagUpdateOpenAttrInfoVos `json:"openAttrInfoVos,omitempty"`
	ProductInstantId  int64                               `json:"productInstantId,omitempty"`
	TagId             int64                               `json:"tagId,omitempty"`
	TagType           int64                               `json:"tagType,omitempty"`
	TagName           string                              `json:"tagName,omitempty"`
	TagColor          string                              `json:"tagColor,omitempty"`
	AttType           int64                               `json:"attType,omitempty"`
	TagSource         string                              `json:"tagSource,omitempty"`
	TagDataUpdateTime string                              `json:"tagDataUpdateTime,omitempty"`
	BizViewFilterJson string                              `json:"bizViewFilterJson,omitempty"`
	TagVersion        int64                               `json:"tagVersion,omitempty"`
	UpdateTime        string                              `json:"updateTime,omitempty"`
	Tcode             string                              `json:"tcode,omitempty"`
	Event             string                              `json:"event,omitempty"`
	OperationSource   int64                               `json:"operationSource,omitempty"`
	ExtTagGId         string                              `json:"extTagGId,omitempty"`
	ChannelId         string                              `json:"channelId,omitempty"`
	BosId             int64                               `json:"bosId,omitempty"`
	Vid               int64                               `json:"vid,omitempty"`
	VidType           int64                               `json:"vidType,omitempty"`
}

type WeimobCdpTagUpdateOpenAttrInfoVos struct {
	AttrEvent string `json:"attrEvent,omitempty"`
	ExtTagId  string `json:"extTagId,omitempty"`
}

// OnWeimobCdpTagUpdateEvent
// @id 1266
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1266?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=标签变更)
func (s *Service) OnWeimobCdpTagUpdateEvent(handler msg.WosEventHandler[WeimobCdpTagUpdateEvent]) (service *Service) {
	var listener = &WeimobCdpTagUpdateListener{handler}
	s.Register(msg.WrapListener[WeimobCdpTagUpdateEvent](listener), "weimob_cdp.tag", "update")
	return s
}
