package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponTemplateGetList
// @id 1601
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1601?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=批量获取优惠券列表)
func (client *Client) CouponTemplateGetList(request *CouponTemplateGetListRequest) (response *CouponTemplateGetListResponse, err error) {
	rpcResponse := CreateCouponTemplateGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponTemplateGetListRequest struct {
	*api.BaseRequest

	CouponTemplateIds []int64 `json:"couponTemplateIds,omitempty"`
	Useful            bool    `json:"useful,omitempty"`
	Vid               int64   `json:"vid,omitempty"`
	NeedAcceptInfo    bool    `json:"needAcceptInfo,omitempty"`
}

type CouponTemplateGetListResponse struct {
	AcceptTypeDTO          CouponTemplateGetListResponseAcceptTypeDTO   `json:"acceptTypeDTO,omitempty"`
	CouponExtDTO           CouponTemplateGetListResponseCouponExtDTO    `json:"couponExtDTO,omitempty"`
	PublishCount           CouponTemplateGetListResponsePublishCount    `json:"publishCount,omitempty"`
	Share                  CouponTemplateGetListResponseShare           `json:"share,omitempty"`
	UserReceiveInfo        CouponTemplateGetListResponseUserReceiveInfo `json:"userReceiveInfo,omitempty"`
	ValidDate              CouponTemplateGetListResponseValidDate       `json:"validDate,omitempty"`
	VidType                CouponTemplateGetListResponseVidType         `json:"vidType,omitempty"`
	BosId                  int64                                        `json:"bosId,omitempty"`
	CanCashTicket          bool                                         `json:"canCashTicket,omitempty"`
	CanGiveFriend          bool                                         `json:"canGiveFriend,omitempty"`
	CanGoodsNumber         bool                                         `json:"canGoodsNumber,omitempty"`
	CanStoreLaunch         bool                                         `json:"canStoreLaunch,omitempty"`
	CashTicketAmt          string                                       `json:"cashTicketAmt,omitempty"`
	CashTicketCondition    string                                       `json:"cashTicketCondition,omitempty"`
	CodeGenerateType       int64                                        `json:"codeGenerateType,omitempty"`
	CouponTemplateId       int64                                        `json:"couponTemplateId,omitempty"`
	CouponType             int64                                        `json:"couponType,omitempty"`
	CreateTime             string                                       `json:"createTime,omitempty"`
	Desc                   string                                       `json:"desc,omitempty"`
	Discount               string                                       `json:"discount,omitempty"`
	DiscountLimitType      int64                                        `json:"discountLimitType,omitempty"`
	DiscountLimitValue     string                                       `json:"discountLimitValue,omitempty"`
	Explain                string                                       `json:"explain,omitempty"`
	FreightReduceType      int64                                        `json:"freightReduceType,omitempty"`
	HasUseThreshold        bool                                         `json:"hasUseThreshold,omitempty"`
	ImageUrl               string                                       `json:"imageUrl,omitempty"`
	IsEnable               bool                                         `json:"isEnable,omitempty"`
	IsExternalCoupon       bool                                         `json:"isExternalCoupon,omitempty"`
	IsPayment              bool                                         `json:"isPayment,omitempty"`
	Logo                   string                                       `json:"logo,omitempty"`
	MaxCashTicketAmt       string                                       `json:"maxCashTicketAmt,omitempty"`
	MaxGoodsNumber         int64                                        `json:"maxGoodsNumber,omitempty"`
	MerchantId             int64                                        `json:"merchantId,omitempty"`
	MinCashTicketAmt       string                                       `json:"minCashTicketAmt,omitempty"`
	MinGoodsNumber         int64                                        `json:"minGoodsNumber,omitempty"`
	Name                   string                                       `json:"name,omitempty"`
	Pid                    int64                                        `json:"pid,omitempty"`
	PlatformType           int64                                        `json:"platformType,omitempty"`
	PublishChannels        []int64                                      `json:"publishChannels,omitempty"`
	RecommendEndTime       string                                       `json:"recommendEndTime,omitempty"`
	RecommendStartTime     string                                       `json:"recommendStartTime,omitempty"`
	RedirectType           int64                                        `json:"redirectType,omitempty"`
	ReducePriceType        int64                                        `json:"reducePriceType,omitempty"`
	ShareImageUrl          string                                       `json:"shareImageUrl,omitempty"`
	SourceProdInstanceId   int64                                        `json:"sourceProdInstanceId,omitempty"`
	SourceProductId        int64                                        `json:"sourceProductId,omitempty"`
	SourceProductVersionId int64                                        `json:"sourceProductVersionId,omitempty"`
	Status                 int64                                        `json:"status,omitempty"`
	Stock                  int64                                        `json:"stock,omitempty"`
	SubName                string                                       `json:"subName,omitempty"`
	ThirdTemplateId        string                                       `json:"thirdTemplateId,omitempty"`
	UserTakeLimit          int64                                        `json:"userTakeLimit,omitempty"`
	Vid                    int64                                        `json:"vid,omitempty"`
}

func CreateCouponTemplateGetListRequest() (request *CouponTemplateGetListRequest) {
	request = &CouponTemplateGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "coupon/template/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateCouponTemplateGetListResponse() (response *api.BaseResponse[CouponTemplateGetListResponse]) {
	response = api.CreateResponse[CouponTemplateGetListResponse](&CouponTemplateGetListResponse{})
	return
}

type CouponTemplateGetListResponseAcceptTypeDTO struct {
	AcceptGoodsIds    []int64 `json:"acceptGoodsIds,omitempty"`
	AcceptGoodsType   int64   `json:"acceptGoodsType,omitempty"`
	AcceptStoreIds    []int64 `json:"acceptStoreIds,omitempty"`
	AcceptStoreType   int64   `json:"acceptStoreType,omitempty"`
	ExcludeGoodsIds   []int64 `json:"excludeGoodsIds,omitempty"`
	ExistExcludeGoods bool    `json:"existExcludeGoods,omitempty"`
}

type CouponTemplateGetListResponseCouponExtDTO struct {
}

type CouponTemplateGetListResponsePublishCount struct {
}

type CouponTemplateGetListResponseShare struct {
	ShareUrl CouponTemplateGetListResponseShareUrl `json:"shareUrl,omitempty"`
	Title    string                                `json:"title,omitempty"`
	Content  string                                `json:"content,omitempty"`
	ImageUrl string                                `json:"imageUrl,omitempty"`
}

type CouponTemplateGetListResponseShareUrl struct {
	H5Url   string `json:"h5Url,omitempty"`
	MiniUrl string `json:"miniUrl,omitempty"`
}

type CouponTemplateGetListResponseUserReceiveInfo struct {
	PublishCount  int64  `json:"publishCount,omitempty"`
	ReceiveStatus int64  `json:"receiveStatus,omitempty"`
	Condition     string `json:"condition,omitempty"`
	IsExpired     bool   `json:"isExpired,omitempty"`
}

type CouponTemplateGetListResponseValidDate struct {
	UseDelayDays   int64  `json:"useDelayDays,omitempty"`
	UseEndTime     string `json:"useEndTime,omitempty"`
	UseStartTime   string `json:"useStartTime,omitempty"`
	UseTimeType    int64  `json:"useTimeType,omitempty"`
	ValidDateDesc  string `json:"validDateDesc,omitempty"`
	ValidDays      int64  `json:"validDays,omitempty"`
	SubUseTimeType int64  `json:"subUseTimeType,omitempty"`
}

type CouponTemplateGetListResponseVidType struct {
}
