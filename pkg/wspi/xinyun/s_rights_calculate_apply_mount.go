package xinyun

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasRightsCalculateApplyMountService struct {
	handler spi.XinyunInvocationHandler[PaasRightsCalculateApplyMountRequest, PaasRightsCalculateApplyMountResponse]
}

func (s PaasRightsCalculateApplyMountService) NewRequest() spi.InvocationRequest[PaasRightsCalculateApplyMountRequest] {
	return &spi.XinyunInvocationRequest[PaasRightsCalculateApplyMountRequest]{
		Params: &PaasRightsCalculateApplyMountRequest{},
	}
}

func (s PaasRightsCalculateApplyMountService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasRightsCalculateApplyMountRequest],
) (
	spi.InvocationResponse[PaasRightsCalculateApplyMountResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.XinyunInvocationRequest[PaasRightsCalculateApplyMountRequest]))
}

type PaasRightsCalculateApplyMountRequest struct {
	ApplyItems []PaasRightsCalculateApplyMountRequestApplyItems `json:"applyItems,omitempty"`
	Wid        int64                                            `json:"wid,omitempty"`
	Pid        int64                                            `json:"pid,omitempty"`
	StoreId    int64                                            `json:"storeId,omitempty"`
	OrderNo    int64                                            `json:"orderNo,omitempty"`
}
type PaasRightsCalculateApplyMountRequestApplyItems struct {
	ItemId           int64  `json:"itemId,omitempty"`
	SkuCode          string `json:"skuCode,omitempty"`
	ApplyAmount      int64  `json:"applyAmount,omitempty"`
	ApplyGoodsPoints int64  `json:"applyGoodsPoints,omitempty"`
	ApplyNum         int64  `json:"applyNum,omitempty"`
}

type PaasRightsCalculateApplyMountResponse struct {
	RightsDetails    []PaasRightsCalculateApplyMountResponseRightsDetails `json:"rightsDetails,omitempty"`
	OrderNo          int64                                                `json:"orderNo,omitempty"`
	TotalApplyAmount int64                                                `json:"totalApplyAmount,omitempty"`
	AmountDesc       string                                               `json:"amountDesc,omitempty"`
	Success          bool                                                 `json:"success,omitempty"`
}
type PaasRightsCalculateApplyMountResponseRightsDetails struct {
	ItemId              int64  `json:"itemId,omitempty"`
	OriginalApplyAmount int64  `json:"originalApplyAmount,omitempty"`
	ApplyAmount         int64  `json:"applyAmount,omitempty"`
	AmountDesc          string `json:"amountDesc,omitempty"`
}

func CreatePaasRightsCalculateApplyMountResponse() *PaasRightsCalculateApplyMountResponse {
	return &PaasRightsCalculateApplyMountResponse{}
}

// OnPaasRightsCalculateApplyMountServiceInvocation
// @id 3883
// @author WeimobCloud
// @create 2023-3-23
// @doc [](https://doc.weimobcloud.com/search?key=售后申请金额批价)
func (s *Service) OnPaasRightsCalculateApplyMountServiceInvocation(
	handler spi.XinyunInvocationHandler[PaasRightsCalculateApplyMountRequest, PaasRightsCalculateApplyMountResponse],
) (service *Service) {
	var invocation = &PaasRightsCalculateApplyMountService{handler}
	s.Register(
		spi.WrapInvoker[PaasRightsCalculateApplyMountRequest, PaasRightsCalculateApplyMountResponse](invocation),
		"PaasRightsCalculateApplyMountService",
		"sRightsCalculateApplyMount",
	)
	return s
}
