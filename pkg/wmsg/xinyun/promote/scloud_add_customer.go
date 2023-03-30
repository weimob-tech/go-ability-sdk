package promote

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

type ScloudAddCustomerListener struct {
	handler msg.XinyunEventHandler[ScloudAddCustomerEvent]
}

func (s ScloudAddCustomerListener) NewMessage() msg.MsgRequest[ScloudAddCustomerEvent] {
	return &msg.XinyunOpenMessage[ScloudAddCustomerEvent]{
		MsgBody: &ScloudAddCustomerEvent{},
	}
}

func (s ScloudAddCustomerListener) OnMessage(ctx context.Context, message msg.MsgRequest[ScloudAddCustomerEvent]) (x.Result, error) {
	return s.handler(ctx, message.(*msg.XinyunOpenMessage[ScloudAddCustomerEvent]))
}

type ScloudAddCustomerEvent struct {
	Tag              []map[string]any `json:"tag,omitempty"`
	Id               int64            `json:"id,omitempty"`
	CustomerWid      int64            `json:"customerWid,omitempty"`
	Nickname         string           `json:"nickname,omitempty"`
	AvatarUrl        string           `json:"avatarUrl,omitempty"`
	RealName         string           `json:"realName,omitempty"`
	Phone            string           `json:"phone,omitempty"`
	CreateTime       int64            `json:"createTime,omitempty"`
	LastActiveTime   int64            `json:"lastActiveTime,omitempty"`
	WxId             string           `json:"wxId,omitempty"`
	Birthday         int64            `json:"birthday,omitempty"`
	City             string           `json:"city,omitempty"`
	CorpName         string           `json:"corpName,omitempty"`
	Job              string           `json:"job,omitempty"`
	Remark           string           `json:"remark,omitempty"`
	Source           string           `json:"source,omitempty"`
	FollowStage      string           `json:"followStage,omitempty"`
	TurnoverRate     byte             `json:"turnoverRate,omitempty"`
	EmployeeWid      int64            `json:"employeeWid,omitempty"`
	EmployeeName     string           `json:"employeeName,omitempty"`
	EmployeeUserId   string           `json:"employeeUserId,omitempty"`
	FirstContactTime int64            `json:"firstContactTime,omitempty"`
}

// OnScloudAddCustomerEvent
// @id 1269
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=新增客户消息
func (s *Service) OnScloudAddCustomerEvent(handler msg.XinyunEventHandler[ScloudAddCustomerEvent]) (service *Service) {
	var listener = &ScloudAddCustomerListener{handler}
	s.Register(msg.WrapListener[ScloudAddCustomerEvent](listener), "scloud", "addCustomer")
	return s
}
