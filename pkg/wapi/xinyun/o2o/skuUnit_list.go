package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SkuUnitList
// @id 267
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取菜品规格列表)
func (client *Client) SkuUnitList(request *SkuUnitListRequest) (response *SkuUnitListResponse, err error) {
	rpcResponse := CreateSkuUnitListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SkuUnitListRequest struct {
	*api.BaseRequest

	PageNum    int   `json:"pageNum,omitempty"`
	PageSize   int   `json:"pageSize,omitempty"`
	MerchantId int64 `json:"merchantId,omitempty"`
}

type SkuUnitListResponse struct {
}

func CreateSkuUnitListRequest() (request *SkuUnitListRequest) {
	request = &SkuUnitListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "skuUnit/list", "api")
	request.Method = api.POST
	return
}

func CreateSkuUnitListResponse() (response *api.BaseResponse[SkuUnitListResponse]) {
	response = api.CreateResponse[SkuUnitListResponse](&SkuUnitListResponse{})
	return
}
