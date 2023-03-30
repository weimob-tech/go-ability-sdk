package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponCustomerGetList
// @id 1988
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1988?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询用户券列表)
func (client *Client) CouponCustomerGetList(request *CouponCustomerGetListRequest) (response *CouponCustomerGetListResponse, err error) {
	rpcResponse := CreateCouponCustomerGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponCustomerGetListRequest struct {
	*api.BaseRequest

	Wid         int64   `json:"wid,omitempty"`
	PageSize    int64   `json:"pageSize,omitempty"`
	StatusRange []int64 `json:"statusRange,omitempty"`
	PageNum     int64   `json:"pageNum,omitempty"`
}

type CouponCustomerGetListResponse struct {
	PageList   []CouponCustomerGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    string                                  `json:"pageNum,omitempty"`
	PageSize   string                                  `json:"pageSize,omitempty"`
	TotalCount string                                  `json:"totalCount,omitempty"`
}

func CreateCouponCustomerGetListRequest() (request *CouponCustomerGetListRequest) {
	request = &CouponCustomerGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "coupon/customer/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateCouponCustomerGetListResponse() (response *api.BaseResponse[CouponCustomerGetListResponse]) {
	response = api.CreateResponse[CouponCustomerGetListResponse](&CouponCustomerGetListResponse{})
	return
}

type CouponCustomerGetListResponsePageList struct {
	ValidDate          CouponCustomerGetListResponseValidDate     `json:"validDate,omitempty"`
	AcceptTypeDTO      CouponCustomerGetListResponseAcceptTypeDTO `json:"acceptTypeDTO,omitempty"`
	Code               string                                     `json:"code,omitempty"`
	PublishTime        int64                                      `json:"publishTime,omitempty"`
	AcquireVid         int64                                      `json:"acquireVid,omitempty"`
	PlatformType       int64                                      `json:"platformType,omitempty"`
	CouponSubType      int64                                      `json:"couponSubType,omitempty"`
	CouponTemplateId   int64                                      `json:"couponTemplateId,omitempty"`
	CouponType         int64                                      `json:"couponType,omitempty"`
	ReducePriceType    int64                                      `json:"reducePriceType,omitempty"`
	DiscountLimitType  int64                                      `json:"discountLimitType,omitempty"`
	DiscountLimitValue string                                     `json:"discountLimitValue,omitempty"`
	Status             int64                                      `json:"status,omitempty"`
	Name               string                                     `json:"name,omitempty"`
	DiscountAmount     string                                     `json:"discountAmount,omitempty"`
	UseThreshold       string                                     `json:"useThreshold,omitempty"`
	MaxGoodsCount      int64                                      `json:"maxGoodsCount,omitempty"`
	MinGoodsCount      int64                                      `json:"minGoodsCount,omitempty"`
	CanGift            bool                                       `json:"canGift,omitempty"`
	CanShare           bool                                       `json:"canShare,omitempty"`
	Desc               string                                     `json:"desc,omitempty"`
	Explain            string                                     `json:"explain,omitempty"`
	IsAllUseScene      bool                                       `json:"isAllUseScene,omitempty"`
	AllSceneList       []int64                                    `json:"allSceneList,omitempty"`
	UserTakeLimit      int64                                      `json:"userTakeLimit,omitempty"`
	UpdateTime         string                                     `json:"updateTime,omitempty"`
}

type CouponCustomerGetListResponseValidDate struct {
	UseStartTime string `json:"useStartTime,omitempty"`
	UseEndTime   string `json:"useEndTime,omitempty"`
}

type CouponCustomerGetListResponseAcceptTypeDTO struct {
	AcceptStoreType   int64   `json:"acceptStoreType,omitempty"`
	AcceptStoreIds    []int64 `json:"acceptStoreIds,omitempty"`
	AcceptGoodsType   int64   `json:"acceptGoodsType,omitempty"`
	AcceptGoodsIds    []int64 `json:"acceptGoodsIds,omitempty"`
	ExistExcludeGoods bool    `json:"existExcludeGoods,omitempty"`
	ExcludeGoodsIds   []int64 `json:"excludeGoodsIds,omitempty"`
	ExistExcludeStore bool    `json:"existExcludeStore,omitempty"`
	ExcludeStoreIds   []int64 `json:"excludeStoreIds,omitempty"`
	IncludeStoreGoods bool    `json:"includeStoreGoods,omitempty"`
}
