package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponCodeGetList
// @id 2212
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2212?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取优惠券码库列表)
func (client *Client) CouponCodeGetList(request *CouponCodeGetListRequest) (response *CouponCodeGetListResponse, err error) {
	rpcResponse := CreateCouponCodeGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponCodeGetListRequest struct {
	*api.BaseRequest

	Param    CouponCodeGetListRequestParam `json:"param,omitempty"`
	Mode     int64                         `json:"mode,omitempty"`
	StartId  int64                         `json:"startId,omitempty"`
	PageNum  int64                         `json:"pageNum,omitempty"`
	PageSize int64                         `json:"pageSize,omitempty"`
}

type CouponCodeGetListResponse struct {
	PageList    []CouponCodeGetListResponsePageList `json:"pageList,omitempty"`
	NextStartId int64                               `json:"nextStartId,omitempty"`
	PageNum     int64                               `json:"pageNum,omitempty"`
	PageSize    int64                               `json:"pageSize,omitempty"`
	TotalCount  int64                               `json:"totalCount,omitempty"`
}

func CreateCouponCodeGetListRequest() (request *CouponCodeGetListRequest) {
	request = &CouponCodeGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "coupon/code/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateCouponCodeGetListResponse() (response *api.BaseResponse[CouponCodeGetListResponse]) {
	response = api.CreateResponse[CouponCodeGetListResponse](&CouponCodeGetListResponse{})
	return
}

type CouponCodeGetListRequestParam struct {
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	Code             string `json:"code,omitempty"`
	Status           int64  `json:"status,omitempty"`
	OrderNo          string `json:"orderNo,omitempty"`
}

type CouponCodeGetListResponsePageList struct {
	CouponTemplateId int64  `json:"couponTemplateId,omitempty"`
	Code             string `json:"code,omitempty"`
	Wid              int64  `json:"wid,omitempty"`
	Status           int64  `json:"status,omitempty"`
	UsedStoreId      int64  `json:"usedStoreId,omitempty"`
	UsedTime         int64  `json:"usedTime,omitempty"`
	PublishTime      int64  `json:"publishTime,omitempty"`
	OrderNo          string `json:"orderNo,omitempty"`
	UseScene         int64  `json:"useScene,omitempty"`
	UpdateTime       int64  `json:"updateTime,omitempty"`
}
