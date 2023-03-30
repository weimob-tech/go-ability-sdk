package Internationalized_software

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsGetRightsList
// @id 1198
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1198?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=售后订单列表)
func (client *Client) RightsGetRightsListV1_0(request *WeimobShopexpressRightsGetRightsListRequest) (response *WeimobShopexpressRightsGetRightsListResponse, err error) {
	rpcResponse := CreateWeimobShopexpressRightsGetRightsListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WeimobShopexpressRightsGetRightsListRequest struct {
	*api.BaseRequest

	StartTime     string `json:"startTime,omitempty"`
	EndTime       string `json:"endTime,omitempty"`
	AftersaleType int64  `json:"aftersaleType,omitempty"`
	QueryStatus   int64  `json:"queryStatus,omitempty"`
	SearchType    int64  `json:"searchType,omitempty"`
	Search        string `json:"search,omitempty"`
	PageSize      int64  `json:"pageSize,omitempty"`
	PageIndex     int64  `json:"pageIndex,omitempty"`
}

type WeimobShopexpressRightsGetRightsListResponse struct {
	DetailOutList []WeimobShopexpressRightsGetRightsListResponseDetailOutList `json:"detailOutList,omitempty"`
	AftersaleNo   string                                                      `json:"aftersaleNo,omitempty"`
	Uid           int64                                                       `json:"uid,omitempty"`
	CreateTime    string                                                      `json:"createTime,omitempty"`
	OrderNo       string                                                      `json:"orderNo,omitempty"`
	RefundType    int64                                                       `json:"refundType,omitempty"`
	TotalAmount   string                                                      `json:"totalAmount,omitempty"`
	RealAmount    string                                                      `json:"realAmount,omitempty"`
	NickName      string                                                      `json:"nickName,omitempty"`
	Status        int64                                                       `json:"status,omitempty"`
	StatusText    string                                                      `json:"statusText,omitempty"`
}

func CreateWeimobShopexpressRightsGetRightsListRequest() (request *WeimobShopexpressRightsGetRightsListRequest) {
	request = &WeimobShopexpressRightsGetRightsListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("Internationalized_software", "v1.0", "rights/getRightsList", "apigw")
	request.Method = api.POST
	return
}

func CreateWeimobShopexpressRightsGetRightsListResponse() (response *api.BaseResponse[WeimobShopexpressRightsGetRightsListResponse]) {
	response = api.CreateResponse[WeimobShopexpressRightsGetRightsListResponse](&WeimobShopexpressRightsGetRightsListResponse{})
	return
}

type WeimobShopexpressRightsGetRightsListResponseDetailOutList struct {
	Count             int64  `json:"count,omitempty"`
	GoodsName         string `json:"goodsName,omitempty"`
	Description       string `json:"description,omitempty"`
	Price             string `json:"price,omitempty"`
	GoodsPic          string `json:"goodsPic,omitempty"`
	RefundAmount      string `json:"refundAmount,omitempty"`
	OrderDetailAmount string `json:"orderDetailAmount,omitempty"`
}
