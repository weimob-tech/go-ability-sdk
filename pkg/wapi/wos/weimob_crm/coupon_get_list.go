package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponGetList
// @id 1613
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1613?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=批量查询券码列表)
func (client *Client) CouponGetList(request *CouponGetListRequest) (response *CouponGetListResponse, err error) {
	rpcResponse := CreateCouponGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponGetListRequest struct {
	*api.BaseRequest

	Codes []int64 `json:"codes,omitempty"`
	Wid   int64   `json:"wid,omitempty"`
}

type CouponGetListResponse struct {
	TextImageList      CouponGetListResponseTextImageList `json:"textImageList,omitempty"`
	ValidDate          CouponGetListResponseValidDate     `json:"validDate,omitempty"`
	CashTicketAmt      string                             `json:"cashTicketAmt,omitempty"`
	Code               string                             `json:"code,omitempty"`
	CouponTemplateId   int64                              `json:"couponTemplateId,omitempty"`
	CouponType         int64                              `json:"couponType,omitempty"`
	Desc               string                             `json:"desc,omitempty"`
	Discount           string                             `json:"discount,omitempty"`
	DiscountLimitType  int64                              `json:"discountLimitType,omitempty"`
	DiscountLimitValue string                             `json:"discountLimitValue,omitempty"`
	ExpireTime         string                             `json:"expireTime,omitempty"`
	Explain            string                             `json:"explain,omitempty"`
	MaxCashTicketAmt   string                             `json:"maxCashTicketAmt,omitempty"`
	MinCashTicketAmt   string                             `json:"minCashTicketAmt,omitempty"`
	Name               string                             `json:"name,omitempty"`
	PublishCount       int64                              `json:"publishCount,omitempty"`
	PublishTime        string                             `json:"publishTime,omitempty"`
	SendChannels       []int64                            `json:"sendChannels,omitempty"`
	Status             int64                              `json:"status,omitempty"`
	Stock              int64                              `json:"stock,omitempty"`
	UpdateTime         string                             `json:"updateTime,omitempty"`
	UseThreshold       string                             `json:"useThreshold,omitempty"`
	UserTakeLimit      int64                              `json:"userTakeLimit,omitempty"`
	Wid                int64                              `json:"wid,omitempty"`
	ConsumeDate        string                             `json:"consumeDate,omitempty"`
	OrderId            string                             `json:"orderId,omitempty"`
	AcceptGoodsType    int64                              `json:"acceptGoodsType,omitempty"`
}

func CreateCouponGetListRequest() (request *CouponGetListRequest) {
	request = &CouponGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "coupon/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateCouponGetListResponse() (response *api.BaseResponse[CouponGetListResponse]) {
	response = api.CreateResponse[CouponGetListResponse](&CouponGetListResponse{})
	return
}

type CouponGetListResponseTextImageList struct {
	ImageUrl       string `json:"imageUrl,omitempty"`
	WechatImageUrl string `json:"wechatImageUrl,omitempty"`
	Text           string `json:"text,omitempty"`
}

type CouponGetListResponseValidDate struct {
	UseDelayDays  int64  `json:"useDelayDays,omitempty"`
	UseEndTime    string `json:"useEndTime,omitempty"`
	UseStartTime  string `json:"useStartTime,omitempty"`
	UseTimeType   int64  `json:"useTimeType,omitempty"`
	ValidDateDesc string `json:"validDateDesc,omitempty"`
	ValidDays     int64  `json:"validDays,omitempty"`
}
