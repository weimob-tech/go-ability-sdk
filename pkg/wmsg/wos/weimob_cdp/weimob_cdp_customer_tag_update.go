package weimob_cdp

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type WeimobCdpCustomerTagUpdateListener struct {
	handler msg.WosEventHandler[WeimobCdpCustomerTagUpdateEvent]
}

func (s WeimobCdpCustomerTagUpdateListener) NewMessage() msg.MsgRequest[WeimobCdpCustomerTagUpdateEvent] {
	return &msg.WosOpenMessage[WeimobCdpCustomerTagUpdateEvent]{
		MsgBody: &WeimobCdpCustomerTagUpdateEvent{},
	}
}

func (s WeimobCdpCustomerTagUpdateListener) OnMessage(ctx context.Context, message msg.MsgRequest[WeimobCdpCustomerTagUpdateEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.WosOpenMessage[WeimobCdpCustomerTagUpdateEvent]))
}

type WeimobCdpCustomerTagUpdateEvent struct {
	BosId            int64  `json:"bosId,omitempty"`
	Vid              int64  `json:"vid,omitempty"`
	VidType          int64  `json:"vidType,omitempty"`
	ProductInstantId int64  `json:"productInstantId,omitempty"`
	Wid              int64  `json:"wid,omitempty"`
	TagId            int64  `json:"tagId,omitempty"`
	AttId            int64  `json:"attId,omitempty"`
	AttValue         string `json:"attValue,omitempty"`
	TagVersion       int64  `json:"tagVersion,omitempty"`
	UpdateTime       string `json:"updateTime,omitempty"`
	Event            string `json:"event,omitempty"`
	Tcode            string `json:"tcode,omitempty"`
	OperationSource  int64  `json:"operationSource,omitempty"`
	Source           int64  `json:"source,omitempty"`
	TagType          int64  `json:"tagType,omitempty"`
	TagName          string `json:"tagName,omitempty"`
	IsManual         int64  `json:"isManual,omitempty"`
	ExtTagGId        string `json:"extTagGId,omitempty"`
	ChannelId        string `json:"channelId,omitempty"`
	ExtTagId         string `json:"extTagId,omitempty"`
}

// OnWeimobCdpCustomerTagUpdateEvent
// @id 1260
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/2/1260?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=用户标签关系变更（治理）)
func (s *Service) OnWeimobCdpCustomerTagUpdateEvent(handler msg.WosEventHandler[WeimobCdpCustomerTagUpdateEvent]) (service *Service) {
	var listener = &WeimobCdpCustomerTagUpdateListener{handler}
	s.Register(msg.WrapListener[WeimobCdpCustomerTagUpdateEvent](listener), "weimob_cdp.customer.tag", "update")
	return s
}
