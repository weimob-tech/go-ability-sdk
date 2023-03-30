package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponQueryCodeList
// @id 1454
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=根据模板id查询券码列表)
func (client *Client) CouponQueryCodeList(request *CouponQueryCodeListRequest) (response *CouponQueryCodeListResponse, err error) {
	rpcResponse := CreateCouponQueryCodeListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponQueryCodeListRequest struct {
	*api.BaseRequest

	QueryParameter CouponQueryCodeListRequestQueryParameter `json:"queryParameter,omitempty"`
	PageNum        int                                      `json:"pageNum,omitempty"`
	PageSize       int                                      `json:"pageSize,omitempty"`
	Mode           int                                      `json:"mode,omitempty"`
	StartId        int64                                    `json:"startId,omitempty"`
}

type CouponQueryCodeListResponse struct {
	PageList    CouponQueryCodeListResponsePageList `json:"pageList,omitempty"`
	TotalCount  int                                 `json:"totalCount,omitempty"`
	PageSize    int                                 `json:"pageSize,omitempty"`
	PageNum     int                                 `json:"pageNum,omitempty"`
	NextStartId int64                               `json:"nextStartId,omitempty"`
}

func CreateCouponQueryCodeListRequest() (request *CouponQueryCodeListRequest) {
	request = &CouponQueryCodeListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "coupon/queryCodeList", "api")
	request.Method = api.POST
	return
}

func CreateCouponQueryCodeListResponse() (response *api.BaseResponse[CouponQueryCodeListResponse]) {
	response = api.CreateResponse[CouponQueryCodeListResponse](&CouponQueryCodeListResponse{})
	return
}

type CouponQueryCodeListRequestQueryParameter struct {
	CardTemplateId int64  `json:"cardTemplateId,omitempty"`
	Status         int    `json:"status,omitempty"`
	Code           string `json:"code,omitempty"`
	OrderId        string `json:"orderId,omitempty"`
}

type CouponQueryCodeListResponsePageList struct {
	CardTemplateId int64  `json:"cardTemplateId,omitempty"`
	Code           string `json:"code,omitempty"`
	Wid            int64  `json:"wid,omitempty"`
	Status         int    `json:"status,omitempty"`
	UsedStoreId    int64  `json:"usedStoreId,omitempty"`
	UsedTime       int64  `json:"usedTime,omitempty"`
	UpdateTime     int64  `json:"updateTime,omitempty"`
	PublishTime    int64  `json:"publishTime,omitempty"`
	OrderId        string `json:"orderId,omitempty"`
	UseScene       int    `json:"useScene,omitempty"`
}
