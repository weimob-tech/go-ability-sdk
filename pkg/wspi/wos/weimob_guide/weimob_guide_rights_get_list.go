package weimob_guide

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
)

type PaasWeimobGuideRightsGetListService struct {
	handler spi.WosInvocationHandler[PaasWeimobGuideRightsGetListRequest, PaasWeimobGuideRightsGetListResponse]
}

func (s PaasWeimobGuideRightsGetListService) NewRequest() spi.InvocationRequest[PaasWeimobGuideRightsGetListRequest] {
	return &spi.WosInvocationRequest[PaasWeimobGuideRightsGetListRequest]{
		Params: &PaasWeimobGuideRightsGetListRequest{},
	}
}

func (s PaasWeimobGuideRightsGetListService) Invoke(
	ctx context.Context,
	request spi.InvocationRequest[PaasWeimobGuideRightsGetListRequest],
) (
	spi.InvocationResponse[PaasWeimobGuideRightsGetListResponse],
	error,
) {
	return s.handler(ctx, request.(*spi.WosInvocationRequest[PaasWeimobGuideRightsGetListRequest]))
}

type PaasWeimobGuideRightsGetListRequest struct {
	RightsOrderNos []int64 `json:"rightsOrderNos,omitempty"`
}

type PaasWeimobGuideRightsGetListResponse struct {
	RightsList []PaasWeimobGuideRightsGetListResponseRightsList `json:"rightsList,omitempty"`
}
type PaasWeimobGuideRightsGetListResponseRightsList struct {
	AgreeTime                         PaasWeimobGuideRightsGetListResponseAgreeTime            `json:"agreeTime,omitempty"`
	ConfirmReceivedTime               PaasWeimobGuideRightsGetListResponseConfirmReceivedTime  `json:"confirmReceivedTime,omitempty"`
	CreateTime                        PaasWeimobGuideRightsGetListResponseCreateTime           `json:"createTime,omitempty"`
	DeliveryRefundPoints              PaasWeimobGuideRightsGetListResponseDeliveryRefundPoints `json:"deliveryRefundPoints,omitempty"`
	FinishTime                        PaasWeimobGuideRightsGetListResponseFinishTime           `json:"finishTime,omitempty"`
	ItemList                          []PaasWeimobGuideRightsGetListResponseItemList           `json:"itemList,omitempty"`
	OrderNo                           PaasWeimobGuideRightsGetListResponseOrderNo              `json:"orderNo,omitempty"`
	RefundPaySuccessTime              PaasWeimobGuideRightsGetListResponseRefundPaySuccessTime `json:"refundPaySuccessTime,omitempty"`
	RefundPoints                      PaasWeimobGuideRightsGetListResponseRefundPoints         `json:"refundPoints,omitempty"`
	ReturnGoodsTime                   PaasWeimobGuideRightsGetListResponseReturnGoodsTime      `json:"returnGoodsTime,omitempty"`
	RightsId                          PaasWeimobGuideRightsGetListResponseRightsId             `json:"rightsId,omitempty"`
	RightsType                        PaasWeimobGuideRightsGetListResponseRightsType           `json:"rightsType,omitempty"`
	UpdateTime                        PaasWeimobGuideRightsGetListResponseUpdateTime           `json:"updateTime,omitempty"`
	AmountReturnFailReason            string                                                   `json:"amountReturnFailReason,omitempty"`
	BosId                             int64                                                    `json:"bosId,omitempty"`
	CancelType                        int64                                                    `json:"cancelType,omitempty"`
	CustomRightsReason                string                                                   `json:"customRightsReason,omitempty"`
	DeliveryRefundAmount              string                                                   `json:"deliveryRefundAmount,omitempty"`
	DeliveryRefundBalance             string                                                   `json:"deliveryRefundBalance,omitempty"`
	DeliveryRefundPointDiscountAmount string                                                   `json:"deliveryRefundPointDiscountAmount,omitempty"`
	FreightTemplateId                 int64                                                    `json:"freightTemplateId,omitempty"`
	ReasonImageUrlList                string                                                   `json:"reasonImageUrlList,omitempty"`
	ReferOrderType                    int64                                                    `json:"referOrderType,omitempty"`
	RefundAmount                      string                                                   `json:"refundAmount,omitempty"`
	RefundBalance                     string                                                   `json:"refundBalance,omitempty"`
	RefundDeliveryAmount              string                                                   `json:"refundDeliveryAmount,omitempty"`
	RefundInvoiceTexAmount            int64                                                    `json:"refundInvoiceTexAmount,omitempty"`
	RefundMethod                      int64                                                    `json:"refundMethod,omitempty"`
	RefundPointDiscountAmount         string                                                   `json:"refundPointDiscountAmount,omitempty"`
	RefundStorageCard                 int64                                                    `json:"refundStorageCard,omitempty"`
	RefundType                        int64                                                    `json:"refundType,omitempty"`
	RefusedCode                       string                                                   `json:"refusedCode,omitempty"`
	RefusedReason                     string                                                   `json:"refusedReason,omitempty"`
	RightsCauseType                   int64                                                    `json:"rightsCauseType,omitempty"`
	RightsReason                      string                                                   `json:"rightsReason,omitempty"`
	RightsStatus                      int64                                                    `json:"rightsStatus,omitempty"`
	UserNickName                      string                                                   `json:"userNickName,omitempty"`
	Vid                               int64                                                    `json:"vid,omitempty"`
	VidType                           int64                                                    `json:"vidType,omitempty"`
	Wid                               int64                                                    `json:"wid,omitempty"`
}
type PaasWeimobGuideRightsGetListResponseAgreeTime struct {
}
type PaasWeimobGuideRightsGetListResponseConfirmReceivedTime struct {
}
type PaasWeimobGuideRightsGetListResponseCreateTime struct {
}
type PaasWeimobGuideRightsGetListResponseDeliveryRefundPoints struct {
}
type PaasWeimobGuideRightsGetListResponseFinishTime struct {
}
type PaasWeimobGuideRightsGetListResponseItemList struct {
	BalanceDiscountAmount             string `json:"balanceDiscountAmount,omitempty"`
	CostPrice                         int64  `json:"costPrice,omitempty"`
	DeliveryRefundAmount              int64  `json:"deliveryRefundAmount,omitempty"`
	DeliveryRefundBalance             string `json:"deliveryRefundBalance,omitempty"`
	DeliveryRefundPointDiscountAmount string `json:"deliveryRefundPointDiscountAmount,omitempty"`
	DeliveryRefundPoints              int64  `json:"deliveryRefundPoints,omitempty"`
	GoodsId                           int64  `json:"goodsId,omitempty"`
	GoodsTitle                        string `json:"goodsTitle,omitempty"`
	ImageUrl                          string `json:"imageUrl,omitempty"`
	OrderItemId                       int64  `json:"orderItemId,omitempty"`
	PaymentAmount                     string `json:"paymentAmount,omitempty"`
	Price                             string `json:"price,omitempty"`
	RefundAmount                      string `json:"refundAmount,omitempty"`
	RefundBalance                     string `json:"refundBalance,omitempty"`
	RefundDeliveryAmount              string `json:"refundDeliveryAmount,omitempty"`
	RefundInvoiceTexAmount            string `json:"refundInvoiceTexAmount,omitempty"`
	RefundPointDiscountAmount         string `json:"refundPointDiscountAmount,omitempty"`
	RefundPoints                      int64  `json:"refundPoints,omitempty"`
	RefundStorageCard                 string `json:"refundStorageCard,omitempty"`
	SkuId                             int64  `json:"skuId,omitempty"`
	SkuName                           string `json:"skuName,omitempty"`
	SkuNum                            int64  `json:"skuNum,omitempty"`
}
type PaasWeimobGuideRightsGetListResponseOrderNo struct {
}
type PaasWeimobGuideRightsGetListResponseRefundPaySuccessTime struct {
}
type PaasWeimobGuideRightsGetListResponseRefundPoints struct {
}
type PaasWeimobGuideRightsGetListResponseReturnGoodsTime struct {
}
type PaasWeimobGuideRightsGetListResponseRightsId struct {
}
type PaasWeimobGuideRightsGetListResponseRightsType struct {
}
type PaasWeimobGuideRightsGetListResponseUpdateTime struct {
}

func CreatePaasWeimobGuideRightsGetListResponse() *PaasWeimobGuideRightsGetListResponse {
	return &PaasWeimobGuideRightsGetListResponse{}
}

// OnPaasWeimobGuideRightsGetListServiceInvocation
// @id 734
// @author WeimobCloud
// @create 2023-3-22
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/3/734?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询售后列表)
func (s *Service) OnPaasWeimobGuideRightsGetListServiceInvocation(
	handler spi.WosInvocationHandler[PaasWeimobGuideRightsGetListRequest, PaasWeimobGuideRightsGetListResponse],
) (service *Service) {
	var invocation = &PaasWeimobGuideRightsGetListService{handler}
	s.Register(
		spi.WrapInvoker[PaasWeimobGuideRightsGetListRequest, PaasWeimobGuideRightsGetListResponse](invocation),
		"PaasWeimobGuideRightsGetListService",
		"weimobGuide.rights.getList",
	)
	return s
}
