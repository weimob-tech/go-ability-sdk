package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponTemplateSearch
// @id 1600
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1600?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=搜索优惠券)
func (client *Client) CouponTemplateSearch(request *CouponTemplateSearchRequest) (response *CouponTemplateSearchResponse, err error) {
	rpcResponse := CreateCouponTemplateSearchResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponTemplateSearchRequest struct {
	*api.BaseRequest

	CouponType       int64   `json:"couponType,omitempty"`
	Name             string  `json:"name,omitempty"`
	Status           int64   `json:"status,omitempty"`
	PlatformType     int64   `json:"platformType,omitempty"`
	PublishChannels  []int64 `json:"publishChannels,omitempty"`
	BelongVids       []int64 `json:"belongVids,omitempty"`
	AcceptVidType    int64   `json:"acceptVidType,omitempty"`
	AcceptVids       []int64 `json:"acceptVids,omitempty"`
	Paths            []int64 `json:"paths,omitempty"`
	QueryChannel     int64   `json:"queryChannel,omitempty"`
	NeedAppendRefer  bool    `json:"needAppendRefer,omitempty"`
	AcquireChannelId string  `json:"acquireChannelId,omitempty"`
	PageNum          int64   `json:"pageNum,omitempty"`
	PageSize         int64   `json:"pageSize,omitempty"`
	VidType          int64   `json:"vidType,omitempty"`
	Vid              int64   `json:"vid,omitempty"`
}

type CouponTemplateSearchResponse struct {
	PageList   []CouponTemplateSearchResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                  `json:"pageNum,omitempty"`
	PageSize   int64                                  `json:"pageSize,omitempty"`
	TotalCount int64                                  `json:"totalCount,omitempty"`
}

func CreateCouponTemplateSearchRequest() (request *CouponTemplateSearchRequest) {
	request = &CouponTemplateSearchRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "coupon/template/search", "apigw")
	request.Method = api.POST
	return
}

func CreateCouponTemplateSearchResponse() (response *api.BaseResponse[CouponTemplateSearchResponse]) {
	response = api.CreateResponse[CouponTemplateSearchResponse](&CouponTemplateSearchResponse{})
	return
}

type CouponTemplateSearchResponsePageList struct {
	ValidDate              CouponTemplateSearchResponseValidDate       `json:"validDate,omitempty"`
	AcceptTypeDTO          CouponTemplateSearchResponseAcceptTypeDTO   `json:"acceptTypeDTO,omitempty"`
	CouponExtDTO           CouponTemplateSearchResponseCouponExtDTO    `json:"couponExtDTO,omitempty"`
	Share                  CouponTemplateSearchResponseShare           `json:"share,omitempty"`
	UserReceiveInfo        CouponTemplateSearchResponseUserReceiveInfo `json:"userReceiveInfo,omitempty"`
	CouponTemplateId       int64                                       `json:"couponTemplateId,omitempty"`
	Name                   string                                      `json:"name,omitempty"`
	SubName                string                                      `json:"subName,omitempty"`
	CouponType             int64                                       `json:"couponType,omitempty"`
	Status                 int64                                       `json:"status,omitempty"`
	PublishChannels        []int64                                     `json:"publishChannels,omitempty"`
	PlatformType           int64                                       `json:"platformType,omitempty"`
	Stock                  int64                                       `json:"stock,omitempty"`
	Explain                string                                      `json:"explain,omitempty"`
	Logo                   string                                      `json:"logo,omitempty"`
	ReducePriceType        int64                                       `json:"reducePriceType,omitempty"`
	CashTicketAmt          string                                      `json:"cashTicketAmt,omitempty"`
	MinCashTicketAmt       string                                      `json:"minCashTicketAmt,omitempty"`
	MaxCashTicketAmt       string                                      `json:"maxCashTicketAmt,omitempty"`
	Discount               string                                      `json:"discount,omitempty"`
	HasUseThreshold        bool                                        `json:"hasUseThreshold,omitempty"`
	CanCashTicket          bool                                        `json:"canCashTicket,omitempty"`
	CanGoodsNumber         bool                                        `json:"canGoodsNumber,omitempty"`
	MinGoodsNumber         int64                                       `json:"minGoodsNumber,omitempty"`
	MaxGoodsNumber         int64                                       `json:"maxGoodsNumber,omitempty"`
	CashTicketCondition    string                                      `json:"cashTicketCondition,omitempty"`
	DiscountLimitType      int64                                       `json:"discountLimitType,omitempty"`
	DiscountLimitValue     string                                      `json:"discountLimitValue,omitempty"`
	UserTakeLimit          int64                                       `json:"userTakeLimit,omitempty"`
	CanStoreLaunch         bool                                        `json:"canStoreLaunch,omitempty"`
	IsEnable               bool                                        `json:"isEnable,omitempty"`
	IsExternalCoupon       bool                                        `json:"isExternalCoupon,omitempty"`
	Desc                   string                                      `json:"desc,omitempty"`
	RecommendStartTime     string                                      `json:"recommendStartTime,omitempty"`
	RecommendEndTime       string                                      `json:"recommendEndTime,omitempty"`
	ImageUrl               string                                      `json:"imageUrl,omitempty"`
	ShareImageUrl          string                                      `json:"shareImageUrl,omitempty"`
	IsPayment              bool                                        `json:"isPayment,omitempty"`
	CanGiveFriend          bool                                        `json:"canGiveFriend,omitempty"`
	CodeGenerateType       int64                                       `json:"codeGenerateType,omitempty"`
	CreateTime             string                                      `json:"createTime,omitempty"`
	PublishCount           int64                                       `json:"publishCount,omitempty"`
	IsInPackage            int64                                       `json:"isInPackage,omitempty"`
	RedirectType           int64                                       `json:"redirectType,omitempty"`
	BosId                  int64                                       `json:"bosId,omitempty"`
	Vid                    int64                                       `json:"vid,omitempty"`
	VidType                int64                                       `json:"vidType,omitempty"`
	SourceProdInstanceId   int64                                       `json:"sourceProdInstanceId,omitempty"`
	SourceProductId        int64                                       `json:"sourceProductId,omitempty"`
	SourceProductVersionId int64                                       `json:"sourceProductVersionId,omitempty"`
	StockUnlimited         bool                                        `json:"stockUnlimited,omitempty"`
}

type CouponTemplateSearchResponseValidDate struct {
	UseTimeType   int64  `json:"useTimeType,omitempty"`
	UseStartTime  string `json:"useStartTime,omitempty"`
	UseEndTime    string `json:"useEndTime,omitempty"`
	UseDelayDays  int64  `json:"useDelayDays,omitempty"`
	ValidDays     int64  `json:"validDays,omitempty"`
	ValidDateDesc string `json:"validDateDesc,omitempty"`
}

type CouponTemplateSearchResponseAcceptTypeDTO struct {
	AcceptStoreType   int64   `json:"acceptStoreType,omitempty"`
	AcceptStoreIds    []int64 `json:"acceptStoreIds,omitempty"`
	AcceptGoodsType   int64   `json:"acceptGoodsType,omitempty"`
	AcceptGoodsIds    []int64 `json:"acceptGoodsIds,omitempty"`
	ExistExcludeGoods bool    `json:"existExcludeGoods,omitempty"`
	ExcludeGoodsIds   []int64 `json:"excludeGoodsIds,omitempty"`
}

type CouponTemplateSearchResponseCouponExtDTO struct {
}

type CouponTemplateSearchResponseShare struct {
	ShareUrl CouponTemplateSearchResponseShareUrl `json:"shareUrl,omitempty"`
	Title    string                               `json:"title,omitempty"`
	Content  string                               `json:"content,omitempty"`
	ImageUrl string                               `json:"imageUrl,omitempty"`
}

type CouponTemplateSearchResponseShareUrl struct {
	H5Url   string `json:"h5Url,omitempty"`
	MiniUrl string `json:"miniUrl,omitempty"`
}

type CouponTemplateSearchResponseUserReceiveInfo struct {
	PublishCount  int64  `json:"publishCount,omitempty"`
	ReceiveStatus int64  `json:"receiveStatus,omitempty"`
	Condition     string `json:"condition,omitempty"`
	IsExpired     bool   `json:"isExpired,omitempty"`
}
