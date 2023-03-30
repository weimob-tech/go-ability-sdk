package weimob

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobCouponPresentGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobCouponPresentGetListRequest, PaasWeimobCouponPresentGetListResponse]
}

func (s PaasWeimobCouponPresentGetListService) NewRequest() spi.InvocationRequest[PaasWeimobCouponPresentGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobCouponPresentGetListRequest]{
		Params: &PaasWeimobCouponPresentGetListRequest{},
	}
}

func (s PaasWeimobCouponPresentGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobCouponPresentGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobCouponPresentGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobCouponPresentGetListRequest]))
}

type PaasWeimobCouponPresentGetListRequest struct {
	QueryParameter PaasWeimobCouponPresentGetListRequestQueryParameter `json:"queryParameter,omitempty"`
	PageNum        int64                                               `json:"pageNum,omitempty"`
	PageSize       int64                                               `json:"pageSize,omitempty"`
}
type PaasWeimobCouponPresentGetListRequestQueryParameter struct {
	SourceObjectList []PaasWeimobCouponPresentGetListRequestSourceObjectList `json:"sourceObjectList,omitempty"`
	CouponTemplateId int64                                                   `json:"couponTemplateId,omitempty"`
	Code             string                                                  `json:"code,omitempty"`
	FromUser         int64                                                   `json:"fromUser,omitempty"`
}
type PaasWeimobCouponPresentGetListRequestSourceObjectList struct {
	Source       int64  `json:"source,omitempty"`
	SourceOpenId string `json:"sourceOpenId,omitempty"`
	SourceAppId  string `json:"sourceAppId,omitempty"`
}

type PaasWeimobCouponPresentGetListResponse struct {
	PageList   []PaasWeimobCouponPresentGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                            `json:"pageNum,omitempty"`
	PageSize   int64                                            `json:"pageSize,omitempty"`
	TotalCount int64                                            `json:"totalCount,omitempty"`
}
type PaasWeimobCouponPresentGetListResponsePageList struct {
	Code             string `json:"code,omitempty"`
	Wid              int64  `json:"wid,omitempty"`
	WidBelongVid     int64  `json:"widBelongVid,omitempty"`
	WidBelongVidName string `json:"widBelongVidName,omitempty"`
	Nickname         string `json:"nickname,omitempty"`
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	UsedStoreId      int64  `json:"usedStoreId,omitempty"`
	UsedTime         int64  `json:"usedTime,omitempty"`
	ReceiveTime      int64  `json:"receiveTime,omitempty"`
	ExpireTime       int64  `json:"expireTime,omitempty"`
	OrderId          string `json:"orderId,omitempty"`
	Status           int64  `json:"status,omitempty"`
	UseChannel       int64  `json:"useChannel,omitempty"`
	Operator         int64  `json:"operator,omitempty"`
	OperatorPhone    string `json:"operatorPhone,omitempty"`
	UseScene         string `json:"useScene,omitempty"`
	UseVid           int64  `json:"useVid,omitempty"`
	UseVidName       string `json:"useVidName,omitempty"`
	ChannelType      int64  `json:"channelType,omitempty"`
	ChannelDesc      string `json:"channelDesc,omitempty"`
	AcquireSceneDesc string `json:"acquireSceneDesc,omitempty"`
	FromUser         int64  `json:"fromUser,omitempty"`
	FromUserName     string `json:"fromUserName,omitempty"`
}

func CreatePaasWeimobCouponPresentGetListResponse() *PaasWeimobCouponPresentGetListResponse {
	return &PaasWeimobCouponPresentGetListResponse{}
}

// OnPaasWeimobCouponPresentGetListServiceInvocation
// @id 1070
// @author WeimobCloud
// @create 2023-3-23
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1070?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=转赠记录查询)
func (s *Service) OnPaasWeimobCouponPresentGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobCouponPresentGetListRequest, PaasWeimobCouponPresentGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobCouponPresentGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobCouponPresentGetListRequest, PaasWeimobCouponPresentGetListResponse](invocation),
		"PaasWeimobCouponPresentGetListService",
		"weimob.coupon.present.getList",
	)
	return s
}
