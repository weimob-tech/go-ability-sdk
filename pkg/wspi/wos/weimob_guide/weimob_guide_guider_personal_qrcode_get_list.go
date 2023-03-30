package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuideGuiderPersonalQrcodeGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuideGuiderPersonalQrcodeGetListRequest, PaasWeimobGuideGuiderPersonalQrcodeGetListResponse]
}

func (s PaasWeimobGuideGuiderPersonalQrcodeGetListService) NewRequest() spi.InvocationRequest[PaasWeimobGuideGuiderPersonalQrcodeGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuideGuiderPersonalQrcodeGetListRequest]{
		Params: &PaasWeimobGuideGuiderPersonalQrcodeGetListRequest{},
	}
}

func (s PaasWeimobGuideGuiderPersonalQrcodeGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuideGuiderPersonalQrcodeGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobGuideGuiderPersonalQrcodeGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuideGuiderPersonalQrcodeGetListRequest]))
}

type PaasWeimobGuideGuiderPersonalQrcodeGetListRequest struct {
	GuiderList []PaasWeimobGuideGuiderPersonalQrcodeGetListRequestGuiderList `json:"guiderList,omitempty"`
}
type PaasWeimobGuideGuiderPersonalQrcodeGetListRequestGuiderList struct {
	GuiderWid   int64  `json:"guiderWid,omitempty"`
	GuiderPhone string `json:"guiderPhone,omitempty"`
}

type PaasWeimobGuideGuiderPersonalQrcodeGetListResponse struct {
	GuiderList []PaasWeimobGuideGuiderPersonalQrcodeGetListResponseGuiderList `json:"guiderList,omitempty"`
}
type PaasWeimobGuideGuiderPersonalQrcodeGetListResponseGuiderList struct {
	GuiderWid int64   `json:"guiderWid,omitempty"`
	QrCodes   []int64 `json:"qrCodes,omitempty"`
}

func CreatePaasWeimobGuideGuiderPersonalQrcodeGetListResponse() *PaasWeimobGuideGuiderPersonalQrcodeGetListResponse {
	return &PaasWeimobGuideGuiderPersonalQrcodeGetListResponse{}
}

// OnPaasWeimobGuideGuiderPersonalQrcodeGetListServiceInvocation
// @id 1375
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/1375?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=批量获取导购个人二维码)
func (s *Service) OnPaasWeimobGuideGuiderPersonalQrcodeGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuideGuiderPersonalQrcodeGetListRequest, PaasWeimobGuideGuiderPersonalQrcodeGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuideGuiderPersonalQrcodeGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuideGuiderPersonalQrcodeGetListRequest, PaasWeimobGuideGuiderPersonalQrcodeGetListResponse](invocation),
		"PaasWeimobGuideGuiderPersonalQrcodeGetListService",
		"weimobGuide.guider.personal.qrcode.getList",
	)
	return s
}
