package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuideGetCommissionGoodsList
// @id 1134
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询特殊商品提成明细)
func (client *Client) GuideGetCommissionGoodsList(request *GuideGetCommissionGoodsListRequest) (response *GuideGetCommissionGoodsListResponse, err error) {
	rpcResponse := CreateGuideGetCommissionGoodsListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuideGetCommissionGoodsListRequest struct {
	*api.BaseRequest

	PageNum     int    `json:"pageNum,omitempty"`
	PageSize    int    `json:"pageSize,omitempty"`
	GuiderWid   int64  `json:"guiderWid,omitempty"`
	CreateTime  string `json:"createTime,omitempty"`
	EndTime     string `json:"endTime,omitempty"`
	GuiderPhone string `json:"guiderPhone,omitempty"`
	JobNumber   string `json:"jobNumber,omitempty"`
}

type GuideGetCommissionGoodsListResponse struct {
}

func CreateGuideGetCommissionGoodsListRequest() (request *GuideGetCommissionGoodsListRequest) {
	request = &GuideGetCommissionGoodsListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "guide/getCommissionGoodsList", "api")
	request.Method = api.POST
	return
}

func CreateGuideGetCommissionGoodsListResponse() (response *api.BaseResponse[GuideGetCommissionGoodsListResponse]) {
	response = api.CreateResponse[GuideGetCommissionGoodsListResponse](&GuideGetCommissionGoodsListResponse{})
	return
}
